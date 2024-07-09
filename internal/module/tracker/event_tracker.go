package tracker

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/rs/zerolog"

	"github.com/bnb-chain/token-recover-app/internal/abi/tokenhub"
	"github.com/bnb-chain/token-recover-app/internal/abi/tokenrecoverportal"
	"github.com/bnb-chain/token-recover-app/internal/config"
	"github.com/bnb-chain/token-recover-app/internal/store"
	"github.com/bnb-chain/token-recover-app/pkg/keymanager"
	"github.com/bnb-chain/token-recover-app/pkg/util"
)

var (
	TokenRecoveryContractAddress = common.HexToAddress("0x0000000000000000000000000000000000003000")
	TokenHubContractAddress      = common.HexToAddress("0x0000000000000000000000000000000000001004")
)

type EventTracker struct {
	config *config.Config
	km     keymanager.KeyManager
	store  store.GeneralStore

	logger   *zerolog.Logger
	stopChan chan struct{}
}

func NewEventTracker(
	config *config.Config,
	km keymanager.KeyManager,
	store store.GeneralStore,
	logger *zerolog.Logger,
) *EventTracker {
	return &EventTracker{
		config:   config,
		km:       km,
		store:    store,
		stopChan: make(chan struct{}),
		logger:   logger,
	}
}

func (tracker *EventTracker) StartListeningTokenRecoverEvent() error {
	rpcClient, err := rpc.DialContext(context.Background(), tracker.config.BSC.URL)
	if err != nil {
		return err
	}
	ethClient, err := ethclient.Dial(tracker.config.BSC.URL)
	if err != nil {
		return err
	}
	tokenHubAbi, err := abi.JSON(strings.NewReader(tokenhub.TokenhubABI))
	if err != nil {
		return err
	}
	tokenRecoverPortalAbi, err := abi.JSON(strings.NewReader(tokenrecoverportal.TokenrecoverportalABI))
	if err != nil {
		return err
	}

	go func() {
		ticker := time.NewTicker(time.Duration(tracker.config.BSC.BlockInterval) * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				processedNumber, err := tracker.store.BscBlockStore().GetProcessedBlockNumber()
				if err != nil {
					tracker.logger.Err(err).Msg("failed to get processed block number")
					continue
				}
				if processedNumber.Cmp(common.Big0) == 0 {
					processedNumber = big.NewInt(int64(tracker.config.BSC.StartHeight))
				}
				tracker.logger.Info().Int64("start_height", processedNumber.Int64()).Msg("start listening token recover event")
				ctx, cancel := context.WithTimeout(context.Background(), time.Duration(tracker.config.BSC.BlockInterval)*time.Second)
				defer cancel()
				var head *types.Header
				err = rpcClient.CallContext(ctx, &head, "eth_getBlockByNumber", "finalized", false)
				if err == nil && head == nil {
					tracker.logger.Error().Msg("failed to get finalized block")
					continue
				}
				tracker.logger.Info().Uint64("block_number", head.Number.Uint64()).Msg("get finalized block")
				finalizedNumber := head.Number
				if processedNumber.Cmp(finalizedNumber) >= 0 {
					continue
				}
				targetNumber := new(big.Int).Add(processedNumber, big.NewInt(tracker.config.BSC.ProcessLimit))
				if targetNumber.Cmp(finalizedNumber) > 0 {
					targetNumber = finalizedNumber
				}

				// filter logs
				tracker.logger.Info().Uint64("from", processedNumber.Uint64()).Uint64("to", targetNumber.Uint64()).Msg("filter logs")
				query := ethereum.FilterQuery{
					FromBlock: processedNumber,
					ToBlock:   targetNumber,
					Addresses: []common.Address{TokenRecoveryContractAddress, TokenHubContractAddress},
					Topics: [][]common.Hash{
						{
							tokenRecoverPortalAbi.Events["TokenRecoverRequested"].ID,
							tokenHubAbi.Events["TokenRecoverLocked"].ID,
							tokenHubAbi.Events["WithdrawUnlockedToken"].ID,
							tokenHubAbi.Events["CancelTokenRecoverLock"].ID,
						},
					},
				}
				eventLogs, err := ethClient.FilterLogs(ctx, query)
				if err != nil {
					tracker.logger.Err(err).Msg("failed to filter logs")
					continue
				}

				tokenRecoverEvents := make([]*store.TokenRecoverEvent, 0, len(eventLogs))
				tokenRecoverRequestedResults := make(map[string]tokenrecoverportal.TokenRecoverRequestedEvent, len(eventLogs))
				tokenRecoverLockedResults := make(map[string]tokenhub.TokenRecoverLockedEvent, len(eventLogs))
				withdrawUnlockedTokenResults := make([]tokenhub.WithdrawUnlockedTokenEvent, 0, len(eventLogs))
				cancelTokenRecoverLockResults := make([]tokenhub.CancelTokenRecoverLockEvent, 0, len(eventLogs))
				for _, vLog := range eventLogs {
					tracker.logger.Debug().
						Str("block_hash", vLog.BlockHash.Hex()).
						Uint64("block_number", vLog.BlockNumber).
						Str("tx_hash", vLog.TxHash.Hex()).
						Interface("topics", vLog.Topics).
						Str("contract_address", vLog.Address.Hex()).
						Msg("found a log")

					if vLog.Address == TokenRecoveryContractAddress {
						var tokenRecoverRequestEvent tokenrecoverportal.TokenRecoverRequestedEvent
						err = tokenRecoverPortalAbi.UnpackIntoInterface(&tokenRecoverRequestEvent, "TokenRecoverRequested", vLog.Data)
						if err != nil {
							tracker.logger.Debug().Err(err).Msg("failed to unpack TokenRecoverRequested event")
						} else {
							tokenRecoverRequestedResults[vLog.TxHash.Hex()] = tokenRecoverRequestEvent
						}
					}

					if vLog.Address == TokenHubContractAddress {
						if vLog.Topics[0] == tokenHubAbi.Events["TokenRecoverLocked"].ID {
							var tokenRecoverLockedEvent tokenhub.TokenRecoverLockedEvent
							err = tokenHubAbi.UnpackIntoInterface(&tokenRecoverLockedEvent, "TokenRecoverLocked", vLog.Data)
							if err != nil {
								tracker.logger.Debug().Err(err).Msg("failed to unpack TokenRecoverLocked event")
							} else {
								// Get indexed parameters from topics
								tokenSymbol := common.BytesToHash(vLog.Topics[1].Bytes())
								tokenAddr := common.BytesToAddress(vLog.Topics[2].Bytes())
								recipient := common.BytesToAddress(vLog.Topics[3].Bytes())
								tokenRecoverLockedEvent.TokenSymbol = tokenSymbol
								tokenRecoverLockedEvent.TokenAddr = tokenAddr
								tokenRecoverLockedEvent.Recipient = recipient
								tokenRecoverLockedResults[vLog.TxHash.Hex()] = tokenRecoverLockedEvent
							}
						}

						if vLog.Topics[0] == tokenHubAbi.Events["WithdrawUnlockedToken"].ID {
							var withdrawUnlockedTokenEvent tokenhub.WithdrawUnlockedTokenEvent
							err = tokenHubAbi.UnpackIntoInterface(&withdrawUnlockedTokenEvent, "WithdrawUnlockedToken", vLog.Data)
							if err != nil {
								tracker.logger.Debug().Err(err).Msg("failed to unpack WithdrawUnlockedToken event")
							} else {
								// Get indexed parameters from topics
								tokenAddr := common.BytesToAddress(vLog.Topics[1].Bytes())
								recipient := common.BytesToAddress(vLog.Topics[2].Bytes())
								withdrawUnlockedTokenEvent.TokenAddr = tokenAddr
								withdrawUnlockedTokenEvent.Recipient = recipient
								withdrawUnlockedTokenResults = append(withdrawUnlockedTokenResults, withdrawUnlockedTokenEvent)
								fmt.Println("DEBUGGGG!!! WithdrawUnlockedToken", fmt.Sprintf("%+v", withdrawUnlockedTokenEvent))
							}
						}

						if vLog.Topics[0] == tokenHubAbi.Events["CancelTokenRecoverLock"].ID {
							var cancelTokenRecoverLockEvent tokenhub.CancelTokenRecoverLockEvent
							err = tokenHubAbi.UnpackIntoInterface(&cancelTokenRecoverLockEvent, "CancelTokenRecoverLock", vLog.Data)
							if err != nil {
								tracker.logger.Debug().Err(err).Msg("failed to unpack CancelTokenRecoverLock event")
							} else {
								// Get indexed parameters from topics
								tokenSymbol := common.BytesToHash(vLog.Topics[1].Bytes())
								tokenAddr := common.BytesToAddress(vLog.Topics[2].Bytes())
								Attacker := common.BytesToAddress(vLog.Topics[3].Bytes())
								cancelTokenRecoverLockEvent.TokenSymbol = tokenSymbol
								cancelTokenRecoverLockEvent.TokenAddr = tokenAddr
								cancelTokenRecoverLockEvent.Attacker = Attacker
								cancelTokenRecoverLockResults = append(cancelTokenRecoverLockResults, cancelTokenRecoverLockEvent)
								fmt.Println("DEBUGGGG!!! CancelTokenRecoverLock", fmt.Sprintf("%+v", cancelTokenRecoverLockEvent))
							}
						}
					}
				}

				for txHash, tokenRecoverRequestedEvent := range tokenRecoverRequestedResults {
					tokenRecoverLockedEvent, ok := tokenRecoverLockedResults[txHash]
					if !ok {
						continue
					}

					tokenOwner := sdk.AccAddress(tokenRecoverRequestedEvent.OwnerAddress[:])
					tokenSymbol := util.DecodeBytesToSymbol(tokenRecoverRequestedEvent.TokenSymbol[:])
					claimAddress := tokenRecoverRequestedEvent.Account
					amount := tokenRecoverLockedEvent.Amount
					tokenAddress := tokenRecoverLockedEvent.TokenAddr
					unlockAt := tokenRecoverLockedEvent.UnlockAt
					status := store.Locked

					event := &store.TokenRecoverEvent{
						TokenOwner:           tokenOwner,
						TokenContractAddress: common.Address(tokenAddress),
						Denom:                tokenSymbol,
						Amount:               amount,
						ClaimAddress:         common.Address(claimAddress),
						UnlockAt:             unlockAt.Int64(),
						Status:               status,
					}

					tracker.logger.Info().
						Str("token_owner", tokenOwner.String()).
						Str("token_symbol", tokenSymbol).
						Str("claim_address", common.Address(claimAddress).Hex()).
						Str("token_address", common.Address(tokenAddress).Hex()).
						Int64("amount", amount.Int64()).
						Int64("unlock_at", unlockAt.Int64()).
						Msg("found a token recover event")
					tokenRecoverEvents = append(tokenRecoverEvents, event)
				}

				for _, withdrawUnlockedTokenEvent := range withdrawUnlockedTokenResults {
					tokenAddress := withdrawUnlockedTokenEvent.TokenAddr
					claimAddress := withdrawUnlockedTokenEvent.Recipient
					amount := withdrawUnlockedTokenEvent.Amount

					event, err := tracker.store.TokenRecoverEventStore().GetTokenRecoverEvent(store.TokenRecoverEvent{
						TokenContractAddress: common.Address(tokenAddress),
						ClaimAddress:         common.Address(claimAddress),
						Amount:               amount,
					})
					if err != nil {
						tracker.logger.Err(err).
							Str("token_address", common.Address(tokenAddress).Hex()).
							Str("claim_address", common.Address(claimAddress).Hex()).
							Int64("amount", amount.Int64()).
							Msg("failed to get token recover event for withdraw event")
					}

					if err == nil &&
						event.Status == store.Locked &&
						event.Amount.Cmp(amount) == 0 {
						event.Status = store.Unlocked
						tokenRecoverEvents = append(tokenRecoverEvents, event)
					}
				}
				for _, cancelTokenRecoverLockEvent := range cancelTokenRecoverLockResults {
					tokenSymbol := util.DecodeBytesToSymbol(cancelTokenRecoverLockEvent.TokenSymbol[:])
					claimAddress := common.Address(cancelTokenRecoverLockEvent.Attacker)
					amount := cancelTokenRecoverLockEvent.Amount

					event, err := tracker.store.TokenRecoverEventStore().GetTokenRecoverEvent(store.TokenRecoverEvent{
						Denom:        tokenSymbol,
						ClaimAddress: claimAddress,
						Amount:       amount,
					})
					if err != nil {
						tracker.logger.Err(err).
							Str("token_symbol", tokenSymbol).
							Str("claim_address", claimAddress.Hex()).
							Msg("failed to get token recover event for cancel event")
					}

					if err == nil {
						event.Status = store.Cancelled
						tokenRecoverEvents = append(tokenRecoverEvents, event)
					}
				}

				if len(tokenRecoverEvents) > 0 {
					err = tracker.store.TokenRecoverEventStore().BatchSaveTokenRecoverEvent(tokenRecoverEvents)
					if err != nil {
						tracker.logger.Err(err).Msg("failed to save token recover events")
					}
				}

				err = tracker.store.BscBlockStore().SaveProcessedBlockNumber(targetNumber)
				if err != nil {
					tracker.logger.Err(err).Uint64("block_number", targetNumber.Uint64()).Msg("failed to save processed block number")
				}
				tracker.logger.Info().Uint64("block_number", targetNumber.Uint64()).Msg("save processed block number")

			case <-tracker.stopChan:
				return
			}
		}
	}()

	return nil
}

func (tracker *EventTracker) StartAutoWithdrawBot() error {
	return nil
}

func (tracker *EventTracker) Stop() error {
	close(tracker.stopChan)
	return nil
}

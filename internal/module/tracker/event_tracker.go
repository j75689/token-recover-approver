package tracker

import (
	"context"
	"math/big"
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
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
				if err != nil || head == nil {
					tracker.logger.Error().Err(err).Msg("failed to get finalized block")
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

				tokenRecoverRequestedResults := make(map[string]tokenrecoverportal.TokenRecoverRequestedEvent, len(eventLogs))
				tokenRecoverLockedResults := make(map[string]tokenhub.TokenRecoverLockedEvent, len(eventLogs))
				withdrawUnlockedTokenResults := make(map[string]tokenhub.WithdrawUnlockedTokenEvent, len(eventLogs))
				cancelTokenRecoverLockResults := make(map[string]tokenhub.CancelTokenRecoverLockEvent, len(eventLogs))
				txToBlockNumber := make(map[string]uint64, len(eventLogs))
				for _, vLog := range eventLogs {
					tracker.logger.Debug().
						Str("block_hash", vLog.BlockHash.Hex()).
						Uint64("block_number", vLog.BlockNumber).
						Str("tx_hash", vLog.TxHash.Hex()).
						Interface("topics", vLog.Topics).
						Str("contract_address", vLog.Address.Hex()).
						Msg("found a log")

					if vLog.Address == TokenRecoveryContractAddress &&
						vLog.Topics[0] == tokenRecoverPortalAbi.Events["TokenRecoverRequested"].ID {
						var tokenRecoverRequestEvent tokenrecoverportal.TokenRecoverRequestedEvent
						err = tokenRecoverPortalAbi.UnpackIntoInterface(&tokenRecoverRequestEvent, "TokenRecoverRequested", vLog.Data)
						if err != nil {
							tracker.logger.Debug().Err(err).Msg("failed to unpack TokenRecoverRequested event")
						} else {
							tokenRecoverRequestedResults[vLog.TxHash.Hex()] = tokenRecoverRequestEvent
							txToBlockNumber[vLog.TxHash.Hex()] = vLog.BlockNumber
						}
					}

					if vLog.Address == TokenHubContractAddress {
						if vLog.Topics[0] == tokenHubAbi.Events["TokenRecoverLocked"].ID {
							var tokenRecoverLockedEvent tokenhub.TokenRecoverLockedEvent
							err = tokenHubAbi.UnpackIntoInterface(&tokenRecoverLockedEvent, "TokenRecoverLocked", vLog.Data)
							if err != nil {
								tracker.logger.Debug().Err(err).Msg("failed to unpack TokenRecoverLocked event")
							} else {
								tokenSymbol := common.BytesToHash(vLog.Topics[1].Bytes())
								tokenAddr := common.BytesToAddress(vLog.Topics[2].Bytes())
								recipient := common.BytesToAddress(vLog.Topics[3].Bytes())
								tokenRecoverLockedEvent.TokenSymbol = tokenSymbol
								tokenRecoverLockedEvent.TokenAddr = tokenAddr
								tokenRecoverLockedEvent.Recipient = recipient
								tokenRecoverLockedResults[vLog.TxHash.Hex()] = tokenRecoverLockedEvent
								txToBlockNumber[vLog.TxHash.Hex()] = vLog.BlockNumber
							}
						}

						if vLog.Topics[0] == tokenHubAbi.Events["WithdrawUnlockedToken"].ID {
							var withdrawUnlockedTokenEvent tokenhub.WithdrawUnlockedTokenEvent
							err = tokenHubAbi.UnpackIntoInterface(&withdrawUnlockedTokenEvent, "WithdrawUnlockedToken", vLog.Data)
							if err != nil {
								tracker.logger.Debug().Err(err).Msg("failed to unpack WithdrawUnlockedToken event")
							} else {
								tokenAddr := common.BytesToAddress(vLog.Topics[1].Bytes())
								recipient := common.BytesToAddress(vLog.Topics[2].Bytes())
								withdrawUnlockedTokenEvent.TokenAddr = tokenAddr
								withdrawUnlockedTokenEvent.Recipient = recipient
								withdrawUnlockedTokenResults[vLog.TxHash.Hex()] = withdrawUnlockedTokenEvent
								txToBlockNumber[vLog.TxHash.Hex()] = vLog.BlockNumber
							}
						}

						if vLog.Topics[0] == tokenHubAbi.Events["CancelTokenRecoverLock"].ID {
							var cancelTokenRecoverLockEvent tokenhub.CancelTokenRecoverLockEvent
							err = tokenHubAbi.UnpackIntoInterface(&cancelTokenRecoverLockEvent, "CancelTokenRecoverLock", vLog.Data)
							if err != nil {
								tracker.logger.Debug().Err(err).Msg("failed to unpack CancelTokenRecoverLock event")
							} else {
								tokenSymbol := common.BytesToHash(vLog.Topics[1].Bytes())
								tokenAddr := common.BytesToAddress(vLog.Topics[2].Bytes())
								Attacker := common.BytesToAddress(vLog.Topics[3].Bytes())
								cancelTokenRecoverLockEvent.TokenSymbol = tokenSymbol
								cancelTokenRecoverLockEvent.TokenAddr = tokenAddr
								cancelTokenRecoverLockEvent.Attacker = Attacker
								cancelTokenRecoverLockResults[vLog.TxHash.Hex()] = cancelTokenRecoverLockEvent
								txToBlockNumber[vLog.TxHash.Hex()] = vLog.BlockNumber
							}
						}
					}
				}

				tokenRecoverEvents := make([]*store.TokenRecoverEvent, 0, len(eventLogs))
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
						RecoveredBlockNumber: txToBlockNumber[txHash],
						RecoveredTxHash:      common.HexToHash(txHash),
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
				if len(tokenRecoverEvents) > 0 {
					err = tracker.store.TokenRecoverEventStore().BatchSaveTokenRecoverEvent(tokenRecoverEvents)
					if err != nil {
						tracker.logger.Err(err).Msg("failed to save token recover events")
					}
				}

				withdrawTokenEvents := make([]*store.TokenRecoverEvent, 0, len(eventLogs))
				for txHash, withdrawUnlockedTokenEvent := range withdrawUnlockedTokenResults {
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

					if err == nil {
						event.Status = store.Unlocked
						event.WithdrawTxHash = common.HexToHash(txHash)
						withdrawTokenEvents = append(withdrawTokenEvents, event)
					}
				}
				if len(withdrawTokenEvents) > 0 {
					err = tracker.store.TokenRecoverEventStore().BatchSaveTokenRecoverEvent(withdrawTokenEvents)
					if err != nil {
						tracker.logger.Err(err).Msg("failed to save token withdraw events")
					}
				}

				cancelTokenEvents := make([]*store.TokenRecoverEvent, 0, len(eventLogs))
				for txHash, cancelTokenRecoverLockEvent := range cancelTokenRecoverLockResults {
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
						event.CancelledTxHash = common.HexToHash(txHash)
						cancelTokenEvents = append(cancelTokenEvents, event)
					}
				}
				if len(cancelTokenEvents) > 0 {
					err = tracker.store.TokenRecoverEventStore().BatchSaveTokenRecoverEvent(cancelTokenEvents)
					if err != nil {
						tracker.logger.Err(err).Msg("failed to save canceled token events")
					}
				}

				err = tracker.store.BscBlockStore().SaveProcessedBlockNumber(targetNumber)
				if err != nil {
					tracker.logger.Err(err).Uint64("block_number", targetNumber.Uint64()).Msg("failed to save processed block number")
				}
				tracker.logger.Info().Uint64("block_number", targetNumber.Uint64()).Msg("save processed block number")

			case <-tracker.stopChan:
				tracker.logger.Info().Msg("stop listening token recover event")
				return
			}
		}
	}()

	return nil
}

func (tracker *EventTracker) StartAutoWithdrawBot() error {
	client, err := ethclient.Dial(tracker.config.BSC.URL)
	if err != nil {
		return err
	}
	tokenHubContract, err := tokenhub.NewTokenhub(TokenHubContractAddress, client)
	if err != nil {
		return err
	}
	go func() {
		ticker := time.NewTicker(time.Duration(tracker.config.BSC.BlockInterval) * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				events, _, err := tracker.store.TokenRecoverEventStore().GetTokenRecoverEvents(store.TokenRecoverEvent{},
					store.Pagination{Offset: 0, Limit: int(tracker.config.BSC.WithdrawLimit)}, &store.ExtraCondition{AllowUnlocked: true})
				if err != nil {
					tracker.logger.Err(err).Msg("failed to get token recover events")
					continue
				}
				if len(events) == 0 {
					continue
				}
				func() {
					ctx, cancel := context.WithTimeout(context.Background(), time.Duration(tracker.config.BSC.BlockInterval)*time.Second)
					defer cancel()
					balance, err := client.BalanceAt(ctx, tracker.km.Address(), nil)
					if err != nil {
						tracker.logger.Err(err).Msg("failed to get balance")
						return
					}
					tracker.logger.Info().Str("address", tracker.km.Address().Hex()).Str("balance", balance.String()).Msg("balance of bot account")

					gasPrice, err := client.SuggestGasPrice(ctx)
					if err != nil {
						tracker.logger.Err(err).Msg("failed to get gas price")
						return
					}
					nonce, err := client.PendingNonceAt(context.Background(), tracker.km.Address())
					if err != nil {
						tracker.logger.Err(err).Str("address", tracker.km.Address().Hex()).Msg("failed to get nonce")
						return
					}
					chainId, err := client.ChainID(ctx)
					if err != nil {
						tracker.logger.Err(err).Msg("failed to get chain Id")
						return
					}
					transactOpts, err := bind.NewKeyedTransactorWithChainID(tracker.km.PrivKey(), chainId)
					if err != nil {
						tracker.logger.Err(err).Msg("failed to bind tx opts")
						return
					}
					transactOpts.Nonce = big.NewInt(int64(nonce))
					transactOpts.Value = common.Big0
					transactOpts.GasLimit = tracker.config.BSC.GasLimit
					transactOpts.GasPrice = gasPrice

					waitingTx := make([]*types.Transaction, 0, len(events))
					txToEvents := make(map[common.Hash]*store.TokenRecoverEvent, len(events))
					for _, event := range events {
						tokenContractAddress := common.HexToAddress("0x0000000000000000000000000000000000000000")
						if event.Denom != "BNB" {
							tokenContractAddress = event.TokenContractAddress
						}
						recipientAddress := event.ClaimAddress
						tx, err := tokenHubContract.WithdrawUnlockedToken(transactOpts, tokenContractAddress, recipientAddress)
						if err != nil {
							tracker.logger.Error().Err(err).
								Interface("event", event).
								Msg("failed to send tx to chain")
							break
						}
						tracker.logger.Info().Interface("event", event).Str("tx_hash", tx.Hash().Hex()).Msg("send a withdraw tx")
						waitingTx = append(waitingTx, tx)
						txToEvents[tx.Hash()] = event

						transactOpts.Nonce = new(big.Int).Add(transactOpts.Nonce, common.Big1)
					}

					confirmedEvents := make([]*store.TokenRecoverEvent, 0, len(events))
					for _, tx := range waitingTx {
						receipt, err := bind.WaitMined(context.Background(), client, tx)
						if err != nil {
							continue
						}
						if receipt.Status != 1 {
							tracker.logger.Error().Str("tx_hash", tx.Hash().Hex()).Msg("tx execution fail")
							continue
						}

						event := txToEvents[tx.Hash()]
						event.Status = store.Unlocked
						event.WithdrawTxHash = tx.Hash()
						confirmedEvents = append(confirmedEvents, event)
					}

					if len(confirmedEvents) > 0 {
						err = tracker.store.TokenRecoverEventStore().BatchSaveTokenRecoverEvent(confirmedEvents)
						if err != nil {
							tracker.logger.Err(err).Int("events_num", len(confirmedEvents)).Msg("failed to save confirmed events")
						}
					}

					tracker.logger.Info().Int("events_num", len(events)).Int("confirmed_events_num", len(confirmedEvents)).Msg("success to withdraw events")
				}()

			case <-tracker.stopChan:
				tracker.logger.Info().Msg("stop auto withdraw bot")
				return
			}
		}
	}()
	return nil
}

func (tracker *EventTracker) Stop() error {
	close(tracker.stopChan)
	return nil
}

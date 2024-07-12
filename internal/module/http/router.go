package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math"
	"math/big"
	"net/http"
	"strconv"

	"github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/julienschmidt/httprouter"

	"github.com/bnb-chain/token-recover-app/internal/module/approval"
	"github.com/bnb-chain/token-recover-app/internal/module/tracker"
	"github.com/bnb-chain/token-recover-app/internal/store"
)

func (server *HttpServer) Ping(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	server.Response(w, Success, http.StatusOK, nil, "pong", nil)
}

func (server *HttpServer) GetTokenRecoverApproval(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		server.Response(w, InvalidRequest, http.StatusOK, nil, nil, err)
		return
	}
	defer r.Body.Close()
	req := &approval.GetTokenRecoverApprovalRequest{}
	err = json.Unmarshal(body, req)
	if err != nil {
		server.Response(w, InvalidRequest, http.StatusOK, nil, nil, err)
		return
	}
	server.logger.Info().Interface("request", req).Msg("GetTokenRecoverApproval")

	err = req.Validate()
	if err != nil {
		server.Response(w, InvalidRequest, http.StatusOK, nil, nil, err)
		return
	}

	resp, err := server.approvalService.GetTokenRecoverApproval(req)
	if err != nil {
		server.Response(w, InvalidRequest, http.StatusOK, nil, nil, err)
		return
	}

	server.Response(w, Success, http.StatusOK, nil, resp, nil)
}

func (server *HttpServer) GetTokenRecoverEvents(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	owner := ps.ByName("owner")
	if len(owner) == 0 {
		server.Response(w, Success, http.StatusOK, nil, []struct{}{}, nil)
		return
	}
	ownerAddr, err := types.AccAddressFromBech32(owner)
	if err != nil {
		server.Response(w, InvalidRequest, http.StatusOK, nil, []struct{}{}, err)
		return
	}
	offsetStr := r.URL.Query().Get("offset")
	limitStr := r.URL.Query().Get("limit")
	offset := 0
	limit := 10
	if newOffset, err := strconv.Atoi(offsetStr); err == nil {
		offset = newOffset
	}
	if newLimit, err := strconv.Atoi(limitStr); err == nil {
		limit = newLimit
	}
	proofs, count, err := server.store.ProofStore().GetAccountAssetProofs(ownerAddr, store.Pagination{Offset: offset, Limit: limit})
	if err != nil {
		server.Response(w, InvalidRequest, http.StatusInternalServerError, nil, nil, err)
		return
	}
	tokenRecoverEvents, _, err := server.store.TokenRecoverEventStore().GetTokenRecoverEvents(
		store.TokenRecoverEvent{
			TokenOwner: ownerAddr,
		},
		store.Pagination{
			Offset: 0,
			Limit:  math.MaxInt,
		},
		nil,
	)
	if err != nil {
		server.Response(w, InvalidRequest, http.StatusInternalServerError, nil, nil, err)
		return
	}
	tokenRecoverEventMap := make(map[string]store.TokenRecoverEvent)
	for _, event := range tokenRecoverEvents {
		tokenRecoverEventMap[event.Denom] = *event
	}
	resp := make([]tracker.TokenRecoverEventResponse, 0, len(proofs))
	for _, proof := range proofs {
		tokenRecoverEventResponse := tracker.TokenRecoverEventResponse{
			Symbol: proof.Denom,
			Amount: big.NewInt(proof.Amount),
			Status: store.Pending,
		}
		tokenInfo, ok := server.tokenList[proof.Denom]
		if ok {
			tokenRecoverEventResponse.Name = tokenInfo.Name
			tokenRecoverEventResponse.ContractAddress = common.HexToAddress(tokenInfo.ContractAddress)
			if tokenInfo.ContractAddress == "" {
				tokenRecoverEventResponse.Status = store.NotBounded
			}
		} else {
			tokenRecoverEventResponse.Status = store.NotBounded
		}
		tokenRecoverEvent, ok := tokenRecoverEventMap[proof.Denom]
		if ok {
			tokenRecoverEventResponse.Status = tokenRecoverEvent.Status
			tokenRecoverEventResponse.ContractAddress = tokenRecoverEvent.TokenContractAddress
			tokenRecoverEventResponse.RecipientAddress = tokenRecoverEvent.ClaimAddress
			tokenRecoverEventResponse.UnlockAt = tokenRecoverEvent.UnlockAt
		}

		resp = append(resp, tokenRecoverEventResponse)
	}
	server.Response(w, Success, http.StatusOK, count, resp, nil)
}

func (server *HttpServer) GetTokenRecoverEvent(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	owner := ps.ByName("owner")
	symbol := r.URL.Query().Get("symbol")

	if len(owner) == 0 || len(symbol) == 0 {
		server.Response(w, InvalidRequest, http.StatusNotFound, nil, nil, fmt.Errorf("invalid params"))
		return
	}
	ownerAddr, err := types.AccAddressFromBech32(owner)
	if err != nil {
		server.Response(w, InvalidRequest, http.StatusNotFound, nil, nil, err)
		return
	}
	proof, err := server.store.ProofStore().GetAccountAssetProof(ownerAddr, symbol)
	if errors.Is(err, store.ErrRecordNotFound) {
		server.Response(w, InvalidRequest, http.StatusNotFound, nil, nil, err)
		return
	}
	if err != nil {
		server.Response(w, InvalidRequest, http.StatusInternalServerError, nil, nil, err)
		return
	}
	tokenInfo, ok := server.tokenList[symbol]
	if !ok {
		server.Response(w, InvalidRequest, http.StatusNotFound, nil, nil, fmt.Errorf("token not found"))
		return
	}
	tokenRecoverEvent, err := server.store.TokenRecoverEventStore().GetTokenRecoverEvent(
		store.TokenRecoverEvent{
			TokenOwner: ownerAddr,
			Denom:      symbol,
		})
	if err != nil && !errors.Is(err, store.ErrRecordNotFound) {
		server.Response(w, InvalidRequest, http.StatusInternalServerError, nil, nil, err)
		return
	}
	status := store.Pending
	contractAddress := common.Address{}
	recipientAddress := common.Address{}
	unlockAt := int64(0)
	if tokenInfo.ContractAddress == "" {
		status = store.NotBounded
	} else {
		contractAddress = common.HexToAddress(tokenInfo.ContractAddress)
	}
	if tokenRecoverEvent != nil {
		status = tokenRecoverEvent.Status
		contractAddress = tokenRecoverEvent.TokenContractAddress
		recipientAddress = tokenRecoverEvent.ClaimAddress
		unlockAt = tokenRecoverEvent.UnlockAt
	}
	server.Response(w, Success, http.StatusOK, nil, tracker.TokenRecoverEventResponse{
		Name:             tokenInfo.Name,
		Symbol:           proof.Denom,
		Amount:           big.NewInt(proof.Amount),
		Status:           status,
		UnlockAt:         unlockAt,
		RecipientAddress: recipientAddress,
		ContractAddress:  contractAddress,
	}, nil)
}

func (server *HttpServer) Response(w http.ResponseWriter, code ResponseCode, statusCode int, count interface{}, data interface{}, err error) {
	resp := Response{
		Code:  code,
		Count: count,
		Data:  data,
	}
	if err != nil {
		resp.Error = err.Error()
	}
	w.WriteHeader(statusCode)
	fmt.Fprint(w, resp.Marshal())
}

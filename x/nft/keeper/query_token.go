package keeper

import (
	"context"

	"cosmossdk.io/store/prefix"
	"divine/x/nft/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) TokenAll(ctx context.Context, req *types.QueryAllTokenRequest) (*types.QueryAllTokenResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var tokens []types.Token

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	tokenStore := prefix.NewStore(store, types.KeyPrefix(types.TokenKeyPrefix))

	pageRes, err := query.Paginate(tokenStore, req.Pagination, func(key []byte, value []byte) error {
		var token types.Token
		if err := k.cdc.Unmarshal(value, &token); err != nil {
			return err
		}

		tokens = append(tokens, token)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllTokenResponse{Token: tokens, Pagination: pageRes}, nil
}

func (k Keeper) Token(ctx context.Context, req *types.QueryGetTokenRequest) (*types.QueryGetTokenResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, found := k.GetToken(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetTokenResponse{Token: val}, nil
}

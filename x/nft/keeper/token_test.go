package keeper_test

import (
	"context"
	"strconv"
	"testing"

	keepertest "divine/testutil/keeper"
	"divine/testutil/nullify"
	"divine/x/nft/keeper"
	"divine/x/nft/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNToken(keeper keeper.Keeper, ctx context.Context, n int) []types.Token {
	items := make([]types.Token, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetToken(ctx, items[i])
	}
	return items
}

func TestTokenGet(t *testing.T) {
	keeper, ctx := keepertest.NftKeeper(t)
	items := createNToken(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetToken(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestTokenRemove(t *testing.T) {
	keeper, ctx := keepertest.NftKeeper(t)
	items := createNToken(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveToken(ctx,
			item.Index,
		)
		_, found := keeper.GetToken(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestTokenGetAll(t *testing.T) {
	keeper, ctx := keepertest.NftKeeper(t)
	items := createNToken(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllToken(ctx)),
	)
}

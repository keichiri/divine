package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "divine/testutil/keeper"
	"divine/x/nft/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.NftKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}

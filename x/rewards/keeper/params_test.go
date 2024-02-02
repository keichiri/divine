package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "divine/testutil/keeper"
	"divine/x/rewards/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.RewardsKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}

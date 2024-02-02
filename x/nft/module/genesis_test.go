package nft_test

import (
	"testing"

	keepertest "divine/testutil/keeper"
	"divine/testutil/nullify"
	"divine/x/nft/module"
	"divine/x/nft/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		TokenList: []types.Token{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.NftKeeper(t)
	nft.InitGenesis(ctx, k, genesisState)
	got := nft.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.TokenList, got.TokenList)
	// this line is used by starport scaffolding # genesis/test/assert
}

package nft

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"divine/testutil/sample"
	nftsimulation "divine/x/nft/simulation"
	"divine/x/nft/types"
)

// avoid unused import issue
var (
	_ = nftsimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgCreateToken = "op_weight_msg_token"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateToken int = 100

	opWeightMsgUpdateToken = "op_weight_msg_token"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateToken int = 100

	opWeightMsgDeleteToken = "op_weight_msg_token"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteToken int = 100

	opWeightMsgTransferToken = "op_weight_msg_transfer_token"
	// TODO: Determine the simulation weight value
	defaultWeightMsgTransferToken int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	nftGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		TokenList: []types.Token{
			{
				Creator: sample.AccAddress(),
				Index:   "0",
			},
			{
				Creator: sample.AccAddress(),
				Index:   "1",
			},
		},
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&nftGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// ProposalContents doesn't return any content functions for governance proposals.
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateToken int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateToken, &weightMsgCreateToken, nil,
		func(_ *rand.Rand) {
			weightMsgCreateToken = defaultWeightMsgCreateToken
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateToken,
		nftsimulation.SimulateMsgCreateToken(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	//var weightMsgUpdateToken int
	//simState.AppParams.GetOrGenerate(opWeightMsgUpdateToken, &weightMsgUpdateToken, nil,
	//	func(_ *rand.Rand) {
	//		weightMsgUpdateToken = defaultWeightMsgUpdateToken
	//	},
	//)
	//operations = append(operations, simulation.NewWeightedOperation(
	//	weightMsgUpdateToken,
	//	nftsimulation.SimulateMsgUpdateToken(am.accountKeeper, am.bankKeeper, am.keeper),
	//))

	//var weightMsgDeleteToken int
	//simState.AppParams.GetOrGenerate(opWeightMsgDeleteToken, &weightMsgDeleteToken, nil,
	//	func(_ *rand.Rand) {
	//		weightMsgDeleteToken = defaultWeightMsgDeleteToken
	//	},
	//)
	//operations = append(operations, simulation.NewWeightedOperation(
	//	weightMsgDeleteToken,
	//	nftsimulation.SimulateMsgDeleteToken(am.accountKeeper, am.bankKeeper, am.keeper),
	//))

	var weightMsgTransferToken int
	simState.AppParams.GetOrGenerate(opWeightMsgTransferToken, &weightMsgTransferToken, nil,
		func(_ *rand.Rand) {
			weightMsgTransferToken = defaultWeightMsgTransferToken
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgTransferToken,
		nftsimulation.SimulateMsgTransferToken(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateToken,
			defaultWeightMsgCreateToken,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				nftsimulation.SimulateMsgCreateToken(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		//simulation.NewWeightedProposalMsg(
		//	opWeightMsgUpdateToken,
		//	defaultWeightMsgUpdateToken,
		//	func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
		//		nftsimulation.SimulateMsgUpdateToken(am.accountKeeper, am.bankKeeper, am.keeper)
		//		return nil
		//	},
		//),
		//simulation.NewWeightedProposalMsg(
		//	opWeightMsgDeleteToken,
		//	defaultWeightMsgDeleteToken,
		//	func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
		//		nftsimulation.SimulateMsgDeleteToken(am.accountKeeper, am.bankKeeper, am.keeper)
		//		return nil
		//	},
		//),
		simulation.NewWeightedProposalMsg(
			opWeightMsgTransferToken,
			defaultWeightMsgTransferToken,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				nftsimulation.SimulateMsgTransferToken(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}

package simulation

import (
	"math/rand"
	"strconv"

	"divine/x/nft/keeper"
	"divine/x/nft/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	moduletestutil "github.com/cosmos/cosmos-sdk/types/module/testutil"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func SimulateMsgCreateToken(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)

		i := r.Int()
		msg := &types.MsgCreateToken{
			Creator: simAccount.Address.String(),
			Index:   strconv.Itoa(i),
		}

		_, found := k.GetToken(ctx, msg.Index)
		if found {
			return simtypes.NoOpMsg(types.ModuleName, sdk.MsgTypeURL(msg), "Token already exist"), nil, nil
		}

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           moduletestutil.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			Context:         ctx,
			SimAccount:      simAccount,
			ModuleName:      types.ModuleName,
			CoinsSpentInMsg: sdk.NewCoins(),
			AccountKeeper:   ak,
			Bankkeeper:      bk,
		}
		return simulation.GenAndDeliverTxWithRandFees(txCtx)
	}
}

//func SimulateMsgUpdateToken(
//	ak types.AccountKeeper,
//	bk types.BankKeeper,
//	k keeper.Keeper,
//) simtypes.Operation {
//	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
//	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
//		var (
//			simAccount = simtypes.Account{}
//			token      = types.Token{}
//			msg        = &types.MsgUpdateToken{}
//			allToken   = k.GetAllToken(ctx)
//			found      = false
//		)
//		for _, obj := range allToken {
//			simAccount, found = FindAccount(accs, obj.Creator)
//			if found {
//				token = obj
//				break
//			}
//		}
//		if !found {
//			return simtypes.NoOpMsg(types.ModuleName, sdk.MsgTypeURL(msg), "token creator not found"), nil, nil
//		}
//		msg.Creator = simAccount.Address.String()
//
//		msg.Index = token.Index
//
//		txCtx := simulation.OperationInput{
//			R:               r,
//			App:             app,
//			TxGen:           moduletestutil.MakeTestEncodingConfig().TxConfig,
//			Cdc:             nil,
//			Msg:             msg,
//			Context:         ctx,
//			SimAccount:      simAccount,
//			ModuleName:      types.ModuleName,
//			CoinsSpentInMsg: sdk.NewCoins(),
//			AccountKeeper:   ak,
//			Bankkeeper:      bk,
//		}
//		return simulation.GenAndDeliverTxWithRandFees(txCtx)
//	}
//}
//
//func SimulateMsgDeleteToken(
//	ak types.AccountKeeper,
//	bk types.BankKeeper,
//	k keeper.Keeper,
//) simtypes.Operation {
//	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
//	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
//		var (
//			simAccount = simtypes.Account{}
//			token      = types.Token{}
//			msg        = &types.MsgUpdateToken{}
//			allToken   = k.GetAllToken(ctx)
//			found      = false
//		)
//		for _, obj := range allToken {
//			simAccount, found = FindAccount(accs, obj.Creator)
//			if found {
//				token = obj
//				break
//			}
//		}
//		if !found {
//			return simtypes.NoOpMsg(types.ModuleName, sdk.MsgTypeURL(msg), "token creator not found"), nil, nil
//		}
//		msg.Creator = simAccount.Address.String()
//
//		msg.Index = token.Index
//
//		txCtx := simulation.OperationInput{
//			R:               r,
//			App:             app,
//			TxGen:           moduletestutil.MakeTestEncodingConfig().TxConfig,
//			Cdc:             nil,
//			Msg:             msg,
//			Context:         ctx,
//			SimAccount:      simAccount,
//			ModuleName:      types.ModuleName,
//			CoinsSpentInMsg: sdk.NewCoins(),
//			AccountKeeper:   ak,
//			Bankkeeper:      bk,
//		}
//		return simulation.GenAndDeliverTxWithRandFees(txCtx)
//	}
//}

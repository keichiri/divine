package nft

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "divine/api/divine/nft"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod: "TokenAll",
					Use:       "list-token",
					Short:     "List all token",
				},
				{
					RpcMethod:      "Token",
					Use:            "show-token [id]",
					Short:          "Shows a token",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "CreateToken",
					Use:            "create-token [index] [owner] [dataDescriptor]",
					Short:          "Create a new token",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "owner"}, {ProtoField: "dataDescriptor"}, {ProtoField: "fee"}},
				},
				//{
				//	RpcMethod:      "UpdateToken",
				//	Use:            "update-token [index] [owner] [dataDescriptor]",
				//	Short:          "Update token",
				//	PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "owner"}, {ProtoField: "dataDescriptor"}},
				//},
				//{
				//	RpcMethod:      "DeleteToken",
				//	Use:            "delete-token [index]",
				//	Short:          "Delete token",
				//	PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				//},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}

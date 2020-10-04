package cli

import (
	"bufio"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/earth2378/logistic/x/logistic/types"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd(cdc *codec.Codec) *cobra.Command {
	logisticTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	logisticTxCmd.AddCommand(flags.PostCommands(
		GetCmdInitDeal(cdc),
		GetCmdTransport(cdc),
		GetCmdUpdateTmp(cdc),
		GetCmdReceive(cdc),
		GetCmdReject(cdc),
		GetCmdClearance(cdc),
	)...)

	return logisticTxCmd
}

func GetCmdInitDeal(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "initDeal [customer] [price] [maxTmp] [minTmp]",
		Short: "Init a new deal",
		Args:  cobra.ExactArgs(4), // Does your request require arguments
		RunE: func(cmd *cobra.Command, args []string) error {

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			price, err := sdk.ParseCoins(args[1])
			if err != nil {
				return err
			}

			maxTmp, err := strconv.Atoi(args[2])
			if err != nil {
				return err
			}
			minTmp, err := strconv.Atoi(args[3])
			if err != nil {
				return err
			}
			msg := types.NewMsgInitDeal(cliCtx.GetFromAddress(), sdk.AccAddress(args[0]), price, maxTmp, minTmp)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

func GetCmdTransport(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "transport [transporter]",
		Short: "select transporter",
		Args:  cobra.ExactArgs(1), // Does your request require arguments
		RunE: func(cmd *cobra.Command, args []string) error {

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			msg := types.NewMsgTransport(cliCtx.GetFromAddress(), sdk.AccAddress(args[0]))
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

func GetCmdUpdateTmp(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "updateTmp [tmp]",
		Short: "update current tmp",
		Args:  cobra.ExactArgs(1), // Does your request require arguments
		RunE: func(cmd *cobra.Command, args []string) error {

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			tmp, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateTmp(cliCtx.GetFromAddress(), tmp)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

func GetCmdReceive(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "receive",
		Short: "receive product",
		Args:  cobra.ExactArgs(0), // Does your request require arguments
		RunE: func(cmd *cobra.Command, args []string) error {

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			msg := types.NewMsgReceive(cliCtx.GetFromAddress())
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

func GetCmdReject(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "reject",
		Short: "reject product",
		Args:  cobra.ExactArgs(0), // Does your request require arguments
		RunE: func(cmd *cobra.Command, args []string) error {

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			msg := types.NewMsgReject(cliCtx.GetFromAddress())
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

func GetCmdClearance(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "clearance",
		Short: "clear product when customer receive product",
		Args:  cobra.ExactArgs(0), // Does your request require arguments
		RunE: func(cmd *cobra.Command, args []string) error {

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			msg := types.NewMsgClearance(cliCtx.GetFromAddress())
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

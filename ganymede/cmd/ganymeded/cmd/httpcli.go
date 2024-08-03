package cmd

import (
	"log"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	httpcli "github.com/jack139/ganymede/ganymede/cmd/http"
	httphelper "github.com/jack139/ganymede/ganymede/cmd/http/helper"
	myclient "github.com/jack139/ganymede/ganymede/cmd/client"
)

func HttpCliCmd() *cobra.Command {
	cmd := &cobra.Command{ // 启动http服务
		Use:   "http",
		Short: "start http service",
		RunE: func(cmd *cobra.Command, args []string) error {
			// 取得参数
			yaml, err := cmd.Flags().GetString("yaml")
			if err!=nil {
				log.Fatal(err)
			}
			// 载入配置文件，初始化
			httphelper.ReadSettings(yaml)

			// 保存 cmd
			httphelper.HttpCmd = cmd


			// 设置 --from 参数
			fromAddr, err := myclient.GetAddrStr(cmd, httphelper.Settings.Chain.NodeUser)
			if err != nil {
				log.Fatal(err)
			}
			log.Println("NodeUser: ", fromAddr)

			cmd.Flags().Set(flags.FlagFrom, fromAddr)  // 设置 --from
			cmd.Flags().Set(flags.FlagChainID, httphelper.Settings.Chain.ChainID)  // 设置 --chain-id

			// 启动 http 服务
			httpcli.RunServer()
			// 不会返回
			return nil
		},
	}

	cmd.Flags().String(flags.FlagChainID, "", "network chain ID")
	//cmd.Flags().String(flags.FlagKeyringDir, "", "The client Keyring directory; if omitted, the default 'home' directory will be used")
	//cmd.Flags().String(flags.FlagKeyringBackend, flags.DefaultKeyringBackend, "Select keyring's backend (os|file|kwallet|pass|test|memory)")
	cmd.Flags().String(flags.FlagFrom, "", "Name or address of private key with which to sign")
	//cmd.Flags().String(flags.FlagNode, "tcp://localhost:26657", "<host>:<port> to tendermint rpc interface for this chain")
	cmd.Flags().BoolP(flags.FlagSkipConfirmation, "y", true, "Skip tx broadcasting prompt confirmation")
	cmd.Flags().Float64(flags.FlagGasAdjustment, 1.5, "adjustment gas")
	cmd.Flags().String(flags.FlagGas, "auto", "gas limit to set per-transaction")
	cmd.Flags().String("yaml", "config/settings.yaml", "yaml config file path")

	return cmd
}

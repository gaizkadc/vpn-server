/*
 * Copyright (C) 2019 Nalej - All Rights Reserved
 */

package commands

import (
	"github.com/nalej/vpn-server/internal/pkg/config"
	"github.com/nalej/vpn-server/internal/pkg/server"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var cfg = config.Config{}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Launch the Edge Controller API",
	Long:  `Launch the Edge Controller API`,
	Run: func(cmd *cobra.Command, args []string) {
		SetupLogging()
		log.Info().Msg("Launching API!")
		cfg.Debug = debugLevel
		server := server.NewService(cfg)
		server.Run()
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().IntVar(&cfg.Port, "port", 5555, "Port to receive management communications")
	runCmd.PersistentFlags().StringVar(&cfg.VPNServerAddress, "vpnServerAddress", "localhost", "VPN Server Address")
}
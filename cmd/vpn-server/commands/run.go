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
	Short: "Launch the VPN Server",
	Long:  `Launch the VPN Server`,
	Run: func(cmd *cobra.Command, args []string) {
		SetupLogging()
		log.Info().Msg("Launching gRPC VPN Server!")
		cfg.Debug = debugLevel

		cfg.Print()
		err := cfg.Validate()
		if err != nil {
			log.Fatal().Err(err)
		}

		server := server.NewService(cfg)
		server.Run()
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().IntVar(&cfg.VPNServerPort, "port", 5666, "Port to launch the gRPC server")
	runCmd.PersistentFlags().StringVar(&cfg.VPNServerAddress, "vpnServerAddress", "localhost", "VPN Server Address")
}
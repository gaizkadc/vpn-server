/*
 * Copyright 2019 Nalej
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
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

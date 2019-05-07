/*
 * Copyright (C) 2019 Nalej - All Rights Reserved
 */

package config

import (
	"github.com/nalej/derrors"
	"github.com/nalej/vpn-server/version"
	"github.com/rs/zerolog/log"
)

type Config struct {
	// Debug level is active.
	Debug bool
	// Port where the gRPC API service will listen requests
	VPNServerPort int
	// URL of the VPN server
	VPNServerAddress string
	// VPNServerPassword
	VPNServerPassword string
}

func (conf *Config) Validate() derrors.Error {

	if conf.VPNServerPort <= 0 {
		return derrors.NewInvalidArgumentError("port must be valid")
	}

	if conf.VPNServerAddress == "" {
		return derrors.NewInvalidArgumentError("VPNServerAddress must be set")
	}

	if conf.VPNServerPassword == "" {
		return derrors.NewInvalidArgumentError("VPNServerPassword must be set")
	}

	return nil
}

func (conf *Config) Print() {
	log.Info().Str("app", version.AppVersion).Str("commit", version.Commit).Msg("Version")
	log.Info().Int("port", conf.VPNServerPort).Msg("gRPC port")
	log.Info().Str("URL", conf.VPNServerAddress).Msg("VPN Server")
	log.Info().Str("password", conf.VPNServerPassword).Msg("VPN Server Password")
}
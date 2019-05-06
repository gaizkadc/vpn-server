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
	// Port where the gRPC API service will listen requests.
	Port int
	// TODO extend description
	// VPNServerAddress
	VPNServerAddress string
	// VPNServerPassword
	VPNServerPassword string
	// TODO Remove
	// Username
	Username string
}

func (conf *Config) Validate() derrors.Error {

	if conf.Port <= 0 {
		return derrors.NewInvalidArgumentError("port must be valid")
	}

	if conf.VPNServerAddress == "" {
		return derrors.NewInvalidArgumentError("VPNServerAddress must be set")
	}

	// TODO Add extra validation

	return nil
}

func (conf *Config) Print() {
	log.Info().Str("app", version.AppVersion).Str("commit", version.Commit).Msg("Version")
	log.Info().Int("port", conf.Port).Msg("gRPC port")
	log.Info().Str("URL", conf.VPNServerAddress).Msg("VPN Server")

	// TODO Print all parameters in the config
}
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
}

func (conf *Config) Validate() derrors.Error {

	if conf.VPNServerPort <= 0 {
		return derrors.NewInvalidArgumentError("port must be valid")
	}

	if conf.VPNServerAddress == "" {
		return derrors.NewInvalidArgumentError("VPNServerAddress must be set")
	}

	return nil
}

func (conf *Config) Print() {
	log.Info().Str("app", version.AppVersion).Str("commit", version.Commit).Msg("Version")
	log.Info().Int("port", conf.VPNServerPort).Msg("gRPC port")
	log.Info().Str("URL", conf.VPNServerAddress).Msg("VPN Server")
}

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

package vpn

import (
	"github.com/nalej/derrors"
	"github.com/rs/zerolog/log"
	"os/exec"
)

type vpnHelper struct {
}

func NewVPNHelper() *vpnHelper {
	return &vpnHelper{}
}

// CreateDefaultUser creates a user and password by default for the EIP to connect to it
func (helper *vpnHelper) CreateDefaultUser() derrors.Error {
	//Create user
	cmd := exec.Command(command, cmdMode, defaultVPNServer, hub, cmdCmd, userCreateCmd, defaultVPNUser, group, realName, note)
	log.Debug().Str("Server", defaultVPNServer).Str("Username", defaultVPNUser).Msg("Default user created in VPN Server")

	err := cmd.Run()
	if err != nil {
		return derrors.NewGenericError("error executing UserCreate command", err)
	}

	// Execute UserPasswordSet command for Username
	cmd = exec.Command(command, cmdMode, defaultVPNServer, hub, cmdCmd, userPasswordSetCmd, defaultVPNUser, userPassword+defaultVPNPassword)
	log.Debug().Str("Server", defaultVPNServer).Str("Username", defaultVPNUser).Msg("Password for default user created")

	err = cmd.Run()
	if err != nil {
		return derrors.NewGenericError("error executing UserDelete userPasswordSet command", err)
	}

	return nil
}

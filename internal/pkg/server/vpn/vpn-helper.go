/*
 * Copyright (C) 2019 Nalej - All Rights Reserved
 */

package vpn

import (
	"github.com/nalej/derrors"
	"github.com/rs/zerolog/log"
	"os/exec"
)

type vpnHelper struct {

}

func NewVPNHelper () *vpnHelper {
	return &vpnHelper{}
}

// CreateDefaultUser creates a user and password by default for the EIP to connect to it
func (helper * vpnHelper) CreateDefaultUser () derrors.Error {
	//Create user
	cmd := exec.Command(command, cmdMode, defaultVPNServer, hub, cmdCmd,  userCreateCmd, defaultVPNUser, group, realName, note)
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
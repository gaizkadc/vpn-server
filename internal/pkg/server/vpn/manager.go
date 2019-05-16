/*
 * Copyright (C) 2019 Nalej - All Rights Reserved
 */

package vpn

import (
	"bytes"
	"github.com/google/uuid"
	"github.com/nalej/derrors"
	"github.com/nalej/grpc-common-go"
	"github.com/nalej/grpc-vpn-server-go"
	"github.com/nalej/vpn-server/internal/pkg/config"
	"github.com/rs/zerolog/log"
	"os/exec"
	"strings"
)


const (
	command = "vpncmd"
	cmdMode = "/Server"
	hub = "/adminhub:DEFAULT"
	cmdCmd = "/cmd"
	userCreateCmd = "UserCreate"
	userPasswordSetCmd = "UserPasswordSet"
	userDeleteCmd = "UserDelete"
	userListCmd = "UserList"
	group = "/GROUP:none"
	realName = "/REALNAME:none"
	note = "/NOTE:none"
	userPassword = "/password:"
)

// Manager structure with the entities involved in the management of VPN users
type Manager struct {
	config config.Config
}

func NewManager(config config.Config) Manager{
	return Manager{
		config: config,
	}
}

// AddVPNUser adds a user to the VPN Server
func (m * Manager) AddVPNUser (addUserRequest grpc_vpn_server_go.AddVPNUserRequest) (*grpc_vpn_server_go.VPNUser, derrors.Error) {
	// Check if server is up, return an error if it's not
	// Check if username exists; return an error if it does

	// Execute command:
	//vpncmd /server Server Name /password:password /adminhub:DEFAULT /cmd UserCreate ABC /GROUP:none /REALNAME:none /NOTE:none
	cmd := exec.Command(command, cmdMode, m.config.VPNServerAddress, hub, cmdCmd,  userCreateCmd, addUserRequest.Username, group, realName, note)
	log.Debug().Str("Server", m.config.VPNServerAddress).Str("Username", addUserRequest.Username).Msg("User created in VPN Server")

	err := cmd.Run()
	if err != nil {
		return nil, derrors.NewGenericError("error executing UserCreate command", err)
	}

	// Create a password
	rawPassword, err := uuid.NewUUID()
	if err != nil {
		return nil, derrors.NewGenericError("error creating password", err)
	}

	password := rawPassword.String()

	// Execute UserPasswordSet command for Username
	cmd = exec.Command(command, cmdMode, m.config.VPNServerAddress, hub, cmdCmd, userPasswordSetCmd, addUserRequest.Username, userPassword+password)
	log.Debug().Str("Server", m.config.VPNServerAddress).Str("Username", addUserRequest.Username).Msg("Password for user created")

	err = cmd.Run()
	if err != nil {
		return nil, derrors.NewGenericError("error executing UserDelete userPasswordSet command", err)
	}

	// Return user
	return &grpc_vpn_server_go.VPNUser{
		Username: addUserRequest.Username,
		Password: password,
	}, nil
}

// DeleteVPNUser adds a user to the VPN Server
func (m * Manager) DeleteVPNUser (deleteUserRequest grpc_vpn_server_go.DeleteVPNUserRequest) (*grpc_common_go.Success, derrors.Error) {
	// Check if server is up, return an error if it's not

	// Check if username exists; return an error if it doesn't

	// Execute command
	cmd := exec.Command(command, cmdMode, m.config.VPNServerAddress, hub, cmdCmd, userDeleteCmd, deleteUserRequest.Username)
	log.Debug().Str("Server", m.config.VPNServerAddress).Str("Username", deleteUserRequest.Username).Msg("User deleted from VPN Server")

	err := cmd.Run()
	if err != nil {
		return nil, derrors.NewGenericError("error executing UserDelete command", err)
	}

	return &grpc_common_go.Success {}, nil
}

// ListVPNUsers list current users from a VPN server
func (m * Manager) ListVPNUsers (listUsersRequest grpc_vpn_server_go.GetVPNUserListRequest) (*grpc_vpn_server_go.VPNUserList, derrors.Error) {
	// Check if server is up, return an error if it's not

	// Execute command
	log.Debug().Str("Server", m.config.VPNServerAddress).Msg("Retrieving user list from VPN Server")
	cmd := exec.Command(command, cmdMode, m.config.VPNServerAddress, hub, cmdCmd, userListCmd)

	var outbuf bytes.Buffer
	cmd.Stdout = &outbuf

	err := cmd.Run()
	if err != nil {
		return nil, derrors.NewGenericError("error executing UserList command", err)
	}

	rawUserList := outbuf.String()
	userList := m.parseRawUserList (rawUserList)

	if err != nil {
		return nil, derrors.NewGenericError("error when parsing user list", err)
	}

	grpcUserList := grpc_vpn_server_go.VPNUserList{
		OrganizationId: listUsersRequest.OrganizationId,
		Usernames: userList,
		}

	return &grpcUserList, nil
}

// parseRawUserList parses the output of the listusers command and returns a clean list of usernames
func (m * Manager) parseRawUserList (raw string) []string {
	lines := strings.Split(raw,"\n")
	userList := make([]string, 0)

	for _, line := range lines {
		if strings.Contains(line, "User Name") {
			tokens := strings.Split(line, "|")
			if len(tokens) == 2 {
				username := tokens[1]
				userList = append(userList, strings.TrimSpace(username))
			}
		}
	}

	return userList
}
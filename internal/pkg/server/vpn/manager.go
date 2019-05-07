/*
 * Copyright (C) 2019 Nalej - All Rights Reserved
 */

package vpn

import (
	"fmt"
	"github.com/nalej/grpc-common-go"
	"github.com/nalej/grpc-vpn-server-go"
	"github.com/nalej/derrors"
	"github.com/nalej/vpn-server/internal/pkg/config"
	"github.com/rs/zerolog/log"
	"os/exec"
	"github.com/google/uuid"
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
	vpnServer := fmt.Sprintf(addUserRequest.Url.Hostname,addUserRequest.Url.Port)
	// Check if server is up, return an error if it's not
	// Check if username exists; return an error if it does

	// Execute command:
	//vpncmd /server Server Name /password:password /adminhub:DEFAULT /cmd UserCreate ABC /GROUP:none /REALNAME:none /NOTE:none
	exec.Command("vpncmd", "/Server", vpnServer, "/password", m.config.VPNServerPassword, "/adminhub:DEFAULT", "/cmd UserCreate", addUserRequest.Username, "/GROUP:none /REALNAME:none /NOTE:none")
	log.Debug().Str("Server", vpnServer).Str("Username", addUserRequest.Username).Msg("User created in VPN Server")

	// Create a password
	rawPassword, err := uuid.NewUUID()
	if err != nil {
		log.Fatal().Errs("could not create new password", []error{err})
	}

	password := rawPassword.String()

	// Execute UserPasswordSet command for Username
	cmd := exec.Command("vpncmd", "/Server", vpnServer, "/password", m.config.VPNServerPassword, "/adminhub:DEFAULT", "/cmd UserPasswordSet", addUserRequest.Username, "/password", password)
	log.Debug().Str("Server", vpnServer).Str("Username", addUserRequest.Username).Msg("Password for user created")

	err = cmd.Run()
	if err != nil {
		log.Fatal().Errs("cmd.Run() failed with %s\n", []error{err})
	}

	// Return user
	return &grpc_vpn_server_go.VPNUser{
		Username: addUserRequest.Username,
		Password: password,
	}, nil
}

// DeleteVPNUser adds a user to the VPN Server
func (m * Manager) DeleteVPNUser (deleteUserRequest grpc_vpn_server_go.DeleteVPNUserRequest) (*grpc_common_go.Success, derrors.Error) {
	vpnServer := fmt.Sprintf(deleteUserRequest.Url.Hostname,deleteUserRequest.Url.Port)

	// Check if server is up, return an error if it's not

	// Check if username exists; return an error if it doesn't

	// Execute command
	cmd := exec.Command("vpncmd", "/Server", vpnServer, "/password", m.config.VPNServerPassword, "/adminhub:DEFAULT", "/cmd UserDelete", deleteUserRequest.Username)
	log.Debug().Str("Server", vpnServer).Str("Username", deleteUserRequest.Username).Msg("User deleted from VPN Server")

	err := cmd.Run()
	if err != nil {
		log.Fatal().Errs("cmd.Run() failed with %s\n", []error{err})
	}

	return &grpc_common_go.Success {}, nil
}

// ListVPNUsers list current users from a VPN server
func (m * Manager) ListVPNUsers (listUsersRequest grpc_vpn_server_go.GetVPNUserListRequest) (*grpc_vpn_server_go.VPNUserList, derrors.Error) {
	vpnServer := fmt.Sprintf(listUsersRequest.Url.Hostname,listUsersRequest.Url.Port)
	// Check if server is up, return an error if it's not

	// Execute command
	log.Debug().Str("Server", vpnServer).Msg("Retrieving user list from VPN Server")
	cmd := exec.Command("vpncmd", "/Server", vpnServer, "/password", m.config.VPNServerPassword, "/adminhub:DEFAULT", "/cmd UserList")

	err := cmd.Run()
	if err != nil {
		log.Fatal().Errs("cmd.Run() failed with %s\n", []error{err})
	}

	// TODO: this is still not returning anything
	userList := grpc_vpn_server_go.VPNUserList{}
	return &userList, nil
}

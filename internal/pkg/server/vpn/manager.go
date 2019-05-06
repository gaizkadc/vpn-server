/*
 * Copyright (C) 2019 Nalej - All Rights Reserved
 */

package vpn

import (
	"github.com/nalej/grpc-common-go"
	"github.com/nalej/grpc-vpn-server-go"
	"github.com/nalej/derrors"
	"github.com/nalej/vpn-server/internal/pkg/config"
	"github.com/rs/zerolog/log"
	"math/rand"
	"os/exec"
	"strings"
	"time"
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
	// TODO Refactor to fmt.Sprintf
	vpnServer := strings.Join([]string{m.config.VPNServerAddress, ":", string(m.config.Port)}, "")
	// Check if server is up, return an error if it's not
	// Check if username exists; return an error if it does

	// Execute command:
	//vpncmd /server Server Name /password:password /adminhub:DEFAULT /cmd UserCreate ABC /GROUP:none /REALNAME:none /NOTE:none
	// TODO The user to be create is on addUserRequest not m.config.Username
	exec.Command("vpncmd", "/Server", vpnServer, "/password", m.config.VPNServerPassword, "/adminhub:DEFAULT", "/cmd UserCreate", m.config.Username, "/GROUP:none /REALNAME:none /NOTE:none")
	log.Debug().Str("Server", vpnServer).Str("Username", m.config.Username).Msg("User created in VPN Server")

	// Create a password
	// TODO Use uuid.NewV4().String()
	password := CreateRandomPassword ()

	// Execute UserPasswordSet command for Username
	// TODO The user to be create is on addUserRequest not m.config.Username
	exec.Command("vpncmd", "/Server", vpnServer, "/password", m.config.VPNServerPassword, "/adminhub:DEFAULT", "/cmd UserPasswordSet", m.config.Username, "/password", password)
	log.Debug().Str("Server", vpnServer).Str("Username", m.config.Username).Msg("Password for user created")

	// TODO Execute command

	// Return user
	// TODO The user to be create is on addUserRequest not m.config.Username
	return &grpc_vpn_server_go.VPNUser{
		Username: m.config.Username,
		Password: password,
	}, nil
}

// DeleteVPNUser adds a user to the VPN Server
func (m * Manager) DeleteVPNUser (deleteUserRequest grpc_vpn_server_go.DeleteVPNUserRequest) (*grpc_common_go.Success, derrors.Error) {
	// TODO Refactor to fmt.Sprintf
	vpnServer := strings.Join([]string{m.config.VPNServerAddress, ":", string(m.config.Port)}, "")

	// Check if server is up, return an error if it's not

	// Check if username exists; return an error if it doesn't

	// Execute command
	// TODO The user to be create is on deleteUserRequest not m.config.Username
	exec.Command("vpncmd", "/Server", vpnServer, "/password", m.config.VPNServerPassword, "/adminhub:DEFAULT", "/cmd UserDelete", m.config.Username)
	log.Debug().Str("Server", vpnServer).Str("Username", m.config.Username).Msg("User deleted from VPN Server")

	// TODO Execute command

	return &grpc_common_go.Success {}, nil
}

// ListVPNUsers list current users from a VPN server
func (m * Manager) ListVPNUsers (listUsersRequest grpc_vpn_server_go.GetVPNUserListRequest) (*grpc_vpn_server_go.VPNUserList, derrors.Error) {
	// TODO Refactor to fmt.Sprintf
	vpnServer := strings.Join([]string{m.config.VPNServerAddress, ":", string(m.config.Port)}, "")
	// Check if server is up, return an error if it's not

	// Execute command
	log.Debug().Str("Server", vpnServer).Msg("Retrieving user list from VPN Server")
	exec.Command("vpncmd", "/Server", vpnServer, "/password", m.config.VPNServerPassword, "/adminhub:DEFAULT", "/cmd UserList")

	// TODO Execute command

	// TODO: this is still not returning anything
	userList := grpc_vpn_server_go.VPNUserList{}
	return &userList, nil
}

// TODO Remove
func CreateRandomPassword () (string) {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	rand.Seed(time.Now().UnixNano())

	pw := make([]rune, 10)

	for i := range pw {
		pw[i] = letters[rand.Intn(len(letters))]
	}

	return string(pw)
}

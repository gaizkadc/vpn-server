/*
 * Copyright (C) 2019 Nalej - All Rights Reserved
 */

package entities

import (
	"github.com/nalej/derrors"
	"github.com/nalej/grpc-vpn-server-go"
)


func ValidAddVPNUserRequest(request *grpc_vpn_server_go.AddVPNUserRequest) derrors.Error{
	// TODO validate organization_id is present and username is not empty
	if request.Url.Hostname == ""{
		return derrors.NewInvalidArgumentError("vpn server hostname must not be empty")
	}
	return nil
}

func ValidDeleteVPNUserRequest(request *grpc_vpn_server_go.DeleteVPNUserRequest) derrors.Error{
	// TODO validate organization_id is present and username is not empty
	if request.Url.Hostname == ""{
		return derrors.NewInvalidArgumentError("vpn server hostname must not be empty")
	}
	return nil
}

func ValidGetVPNUserListRequest(request *grpc_vpn_server_go.GetVPNUserListRequest) derrors.Error{
	// TODO validate organization_id is present
	if request.Url.Hostname == ""{
		return derrors.NewInvalidArgumentError("vpn server hostname must not be empty")
	}
	return nil
}
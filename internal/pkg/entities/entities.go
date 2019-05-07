/*
 * Copyright (C) 2019 Nalej - All Rights Reserved
 */

package entities

import (
	"github.com/nalej/derrors"
	"github.com/nalej/grpc-vpn-server-go"
)


func ValidAddVPNUserRequest(request *grpc_vpn_server_go.AddVPNUserRequest) derrors.Error{
	if request.OrganizationId == ""{
		return derrors.NewInvalidArgumentError("organization_id must not be empty")
	}

	if request.Username == "" {
		return derrors.NewInvalidArgumentError("username must not be empty")
	}

	return nil
}

func ValidDeleteVPNUserRequest(request *grpc_vpn_server_go.DeleteVPNUserRequest) derrors.Error{
	if request.OrganizationId == ""{
		return derrors.NewInvalidArgumentError("organization_id must not be empty")
	}

	if request.Username == "" {
		return derrors.NewInvalidArgumentError("username must not be empty")
	}

	return nil
}

func ValidGetVPNUserListRequest(request *grpc_vpn_server_go.GetVPNUserListRequest) derrors.Error{
	if request.OrganizationId == ""{
		return derrors.NewInvalidArgumentError("organization_id must not be empty")
	}
	return nil
}
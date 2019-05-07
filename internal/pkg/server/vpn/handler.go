/*
 * Copyright (C) 2019 Nalej - All Rights Reserved
 */

package vpn

import (
	"github.com/nalej/grpc-common-go"
	"github.com/nalej/grpc-vpn-server-go"
	"github.com/nalej/vpn-server/internal/pkg/entities"
	"github.com/nalej/grpc-utils/pkg/conversions"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/context"
)

type Handler struct {
	manager Manager
}

func NewHandler(manager Manager) *Handler {
	return &Handler{
		manager,
	}
}


func (h *Handler) AddVPNUser(ctx context.Context, request *grpc_vpn_server_go.AddVPNUserRequest) (*grpc_vpn_server_go.VPNUser, error) {

	verr := entities.ValidAddVPNUserRequest(request)
	if verr != nil {
		return nil, conversions.ToGRPCError(verr)
	}
	// TODO Fix log entry, show organization_id, username
	log.Debug().Str("username", request.Url.Hostname).Msg("add user vpn request")
	response, err := h.manager.AddVPNUser(*request)
	if err != nil{
		return nil, err
	}
	// TODO show organization_id
	log.Debug().Str("username", response.Username).Msg("user has been created")
	return response, nil
}

func (h *Handler) DeleteVPNUser(ctx context.Context, request *grpc_vpn_server_go.DeleteVPNUserRequest) (*grpc_common_go.Success, error) {

	verr := entities.ValidDeleteVPNUserRequest(request)
	if verr != nil {
		return nil, conversions.ToGRPCError(verr)
	}
	// TODO Fix log entry, show organization_id, username
	log.Debug().Str("username", request.Username).Msg("delete user vpn request")
	response, err := h.manager.DeleteVPNUser(*request)
	if err != nil{
		return nil, err
	}
	// TODO Fix log entry, show organization_id, username
	log.Debug().Interface("response", response).Msg("user has been deleted")
	return response, nil
}

func (h *Handler) ListVPNUsers(ctx context.Context, request *grpc_vpn_server_go.GetVPNUserListRequest) (*grpc_vpn_server_go.VPNUserList, error) {

	verr := entities.ValidGetVPNUserListRequest(request)
	if verr != nil {
		return nil, conversions.ToGRPCError(verr)
	}
	// TODO Fix log entry, show organization_id
	log.Debug().Str("vpnServerHostname", request.Url.Hostname).Str("port", request.Url.Port).Msg("list vpn users request")
	response, err := h.manager.ListVPNUsers(*request)
	if err != nil{
		return nil, err
	}

	for i := range response.Username {
		log.Debug().Str("username", response.Username [i]).Msg("username")
	}
	// TODO change log organization_id, len(response.Usernames)
	return response, nil
}



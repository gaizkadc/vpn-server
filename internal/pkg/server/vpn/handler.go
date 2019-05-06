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
	log.Debug().Str("vpnServerHostname", request.Url.Hostname).Str("port", request.Url.Port).Msg("add vpn request")

	verr := entities.ValidAddVPNUserRequest(request)
	if verr != nil {
		return nil, conversions.ToGRPCError(verr)
	}
	response, err := h.manager.AddVPNUser(*request)
	if err != nil{
		return nil, err
	}
	log.Debug().Str("username", response.Username).Msg("user has been created")
	return response, nil
}

func (h *Handler) DeleteVPNUser(ctx context.Context, request *grpc_vpn_server_go.DeleteVPNUserRequest) (*grpc_common_go.Success, error) {
	log.Debug().Str("vpnServerHostname", request.Url.Hostname).Str("port", request.Url.Port).Msg("add vpn request")

	verr := entities.ValidDeleteVPNUserRequest(request)
	if verr != nil {
		return nil, conversions.ToGRPCError(verr)
	}
	response, err := h.manager.DeleteVPNUser(*request)
	if err != nil{
		return nil, err
	}
	log.Debug().Interface("response", response).Msg("user has been deleted")
	return response, nil
}

func (h *Handler) ListVPNUsers(ctx context.Context, request *grpc_vpn_server_go.GetVPNUserListRequest) (*grpc_vpn_server_go.VPNUserList, error) {
	log.Debug().Str("vpnServerHostname", request.Url.Hostname).Str("port", request.Url.Port).Msg("add vpn request")

	verr := entities.ValidGetVPNUserListRequest(request)
	if verr != nil {
		return nil, conversions.ToGRPCError(verr)
	}
	response, err := h.manager.ListVPNUsers(*request)
	if err != nil{
		return nil, err
	}

	for i := range response.Username {
		log.Debug().Str("username", response.Username [i]).Msg("username")
	}
	return response, nil
}



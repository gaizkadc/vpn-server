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

package server

import (
	"fmt"
	"github.com/nalej/grpc-vpn-server-go"
	"github.com/nalej/vpn-server/internal/pkg/config"
	"github.com/nalej/vpn-server/internal/pkg/server/vpn"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

type Service struct {
	Configuration config.Config
}

// NewService creates a new service.
func NewService(conf config.Config) *Service {
	return &Service{
		conf,
	}
}

// Run the service, launch the REST service handler.
func (s *Service) Run() error {
	cErr := s.Configuration.Validate()
	if cErr != nil {
		log.Fatal().Str("err", cErr.DebugReport()).Msg("invalid configuration")
	}
	s.Configuration.Print()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.Configuration.VPNServerPort))
	if err != nil {
		log.Fatal().Errs("failed to listen: %v", []error{err})
	}

	err = vpn.NewVPNHelper().CreateDefaultUser()
	if err != nil {
		log.Fatal().Errs("failed to create default user", []error{err})
	}

	// Create handlers
	manager := vpn.NewManager(s.Configuration)
	handler := vpn.NewHandler(manager)

	// gRPC Server
	grpcServer := grpc.NewServer()

	grpc_vpn_server_go.RegisterVPNServerServer(grpcServer, handler)

	if s.Configuration.Debug {
		log.Info().Msg("Enabling gRPC server reflection")
		// Register reflection service on gRPC server.
		reflection.Register(grpcServer)
	}
	log.Info().Int("port", s.Configuration.VPNServerPort).Msg("Launching gRPC server")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal().Errs("failed to serve: %v", []error{err})
	}

	return nil
}

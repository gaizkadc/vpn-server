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

package entities

import (
	"github.com/nalej/derrors"
	"github.com/nalej/grpc-vpn-server-go"
)

func ValidAddVPNUserRequest(request *grpc_vpn_server_go.AddVPNUserRequest) derrors.Error {
	if request.OrganizationId == "" {
		return derrors.NewInvalidArgumentError("organization_id must not be empty")
	}

	if request.Username == "" {
		return derrors.NewInvalidArgumentError("username must not be empty")
	}

	return nil
}

func ValidDeleteVPNUserRequest(request *grpc_vpn_server_go.DeleteVPNUserRequest) derrors.Error {
	if request.OrganizationId == "" {
		return derrors.NewInvalidArgumentError("organization_id must not be empty")
	}

	if request.Username == "" {
		return derrors.NewInvalidArgumentError("username must not be empty")
	}

	return nil
}

func ValidGetVPNUserListRequest(request *grpc_vpn_server_go.GetVPNUserListRequest) derrors.Error {
	if request.OrganizationId == "" {
		return derrors.NewInvalidArgumentError("organization_id must not be empty")
	}
	return nil
}

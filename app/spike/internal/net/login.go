//    \\ SPIKE: Secure your secrets with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package net

import (
	"encoding/json"
	"errors"
	"github.com/spiffe/go-spiffe/v2/workloadapi"
	"github.com/spiffe/spike/internal/config"
	"github.com/spiffe/spike/internal/entity/v1/reqres"
	"github.com/spiffe/spike/internal/net"
)

func Login(source *workloadapi.X509Source, password string) (string, error) {
	r := reqres.AdminLoginRequest{
		Password: password,
	}
	mr, err := json.Marshal(r)
	if err != nil {
		return "", errors.Join(
			errors.New("login: I am having problem generating the payload"),
			err,
		)
	}

	client, err := net.CreateMtlsClient(source, config.CanTalkToPilot)

	body, err := net.Post(client, urlAdminLogin, mr)
	if err != nil {
		return "", errors.Join(
			errors.New("login: I am having problem sending the request"), err)
	}

	var res reqres.AdminLoginResponse
	err = json.Unmarshal(body, &res)
	if err != nil {
		return "", errors.Join(
			errors.New("login: Problem parsing response body"),
			err,
		)
	}

	return res.Token, nil
}

/*
 * Copyright (C) 2017. Huawei Technologies Co., LTD. All rights reserved.
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of Apache License, Version 2.0.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * Apache License, Version 2.0 for more details.
 */
package auth

import (
	"smn-sdk-go/commom"
	"smn-sdk-go/util"
	"net/http"
	"strings"
	"io/ioutil"
	"encoding/json"
	"time"
	"sync"
	"fmt"
)

var mu sync.Mutex

const (
	xSubjectToken         = "X-Subject-Token"
	expiredInterval int64 = 30 * 60
	id                    = "id"
	expiresAt             = "expires_at"
)
/**
 */
type Auth struct {
	smnConfiguration *commom.SmnConfiguration
	httpClient       *http.Client
	projectId        string
	authToken        string
	expiresTime      int64
}

func NewAuth(smnConfiguration *commom.SmnConfiguration, client *http.Client) *Auth {
	auth := new(Auth)
	auth.smnConfiguration = smnConfiguration
	auth.httpClient = client
	return auth
}

func (auth *Auth) GetTokenAndProject() (token, projectId string, err error) {
	if auth.authToken == "" || auth.isExpired() {
		mu.Lock()
		defer mu.Unlock()
		if auth.authToken == "" || auth.isExpired() {
			if err := auth.postForToken(); err != nil {
				return "", "", err
			}
		}

	}
	projectId = auth.projectId
	token = auth.authToken
	return
}

func (auth *Auth) isExpired() bool {
	return auth.expiresTime < time.Now().UTC().Unix()
}

func (auth *Auth) getTokenUrl() (url string, err error) {
	url = util.HttpsPrefix + util.Iam + "." + auth.smnConfiguration.RegionName + "." + util.Endpoint + "/v3/auth/tokens"
	return
}

func (auth *Auth) postForToken() (err error) {
	reqBody, err := auth.getTokenRequestBody()
	tokenUrl, err := auth.getTokenUrl()

	postReq, err := http.NewRequest("POST",
		tokenUrl,
		strings.NewReader(reqBody)) //post内容

	if err != nil {
		return err
	}

	postReq.Header.Set("Content-Type", "application/json; charset=UTF-8")

	resp, err := auth.httpClient.Do(postReq)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if (resp.StatusCode >= 200 && resp.StatusCode < 300) == false {
		return fmt.Errorf("get %s %s", tokenUrl, body)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return err
	}

	tokenMap := result["token"].(map[string]interface{})
	projectMap := tokenMap["project"].(map[string]interface{})

	auth.authToken = string(resp.Header.Get(xSubjectToken))
	auth.projectId = projectMap[id].(string)

	expiresAt := tokenMap[expiresAt].(string)
	if unixTime, err := util.StringToTimestamp(expiresAt); err == nil {
		auth.expiresTime = unixTime - expiredInterval
	} else {
		return err
	}
	return
}

func (auth *Auth) getTokenRequestBody() (body string, err error) {
	body = `{
		    "auth": {
		        "identity": {
		            "methods": [
		                "password"
		            ],
		            "password": {
		                "user": {
		                    "name": "` + auth.smnConfiguration.Username + "\"," + `
		                    "password":"` + auth.smnConfiguration.Password + "\"," + `
		                    "domain": {
		                        "name": "` + auth.smnConfiguration.DomainName + `"
		                    }
		                }
		            }
		        },
		        "scope": {
		            "project": {
		                "name":"` + auth.smnConfiguration.RegionName + `"
		            }
		        }
		    }
		}`
	return
}
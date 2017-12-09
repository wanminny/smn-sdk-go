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
package client

import (
	"fmt"
	"smn-sdk-go/util"
	"io"
)

type ListTopicAttributesRequest struct {
	*BaseRequest
	TopicUrn string `json:"-"`
	Name     string `json:"name"`
}

type ListTopicAttributesResponse struct {
	*BaseResponse
	Attributes map[string]interface{} `json:"attributes"`
}

func (client *SmnClient) ListTopicAttributes(request *ListTopicAttributesRequest) (response *ListTopicAttributesResponse, err error) {
	response = &ListTopicAttributesResponse{
		BaseResponse: &BaseResponse{},
	}
	err = client.sendRequest(request, response)
	return
}

func (client *SmnClient) NewListTopicAttributesRequest() (request *ListTopicAttributesRequest) {
	request = &ListTopicAttributesRequest{
		BaseRequest: &BaseRequest{Headers: make(map[string]string)},
	}
	return
}

func (request *ListTopicAttributesRequest) GetUrl() (urlStr string, err error) {
	if request.TopicUrn == "" {
		return "", fmt.Errorf("topic urn is null")
	}
	urlStr = request.BaseRequest.GetSmnServiceUrl() + util.V2Version + util.UrlDelimiter + request.projectId +
		util.UrlDelimiter + util.Notifications + util.UrlDelimiter + util.Topics +
		util.UrlDelimiter + request.TopicUrn + util.UrlDelimiter +
		util.Attributes

	urlEncoded, err := util.GetQueryParams(request)
	if err != nil {
		return
	}
	if urlEncoded != "" {
		urlStr += "?" + urlEncoded
	}
	return
}

func (request *ListTopicAttributesRequest) GetMethod() string {
	return util.GET
}

func (request *ListTopicAttributesRequest) GetBodyReader() (reader io.Reader, err error) {
	return nil, nil
}
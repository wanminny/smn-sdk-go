package client

import (
	"smn-sdk-go/util"
	"io"
	"fmt"
)

type UpdateTopicRequest struct {
	*BaseRequest
	TopicUrn    string `json:"-"`
	DisplayName string `json:"display_name"`
}

type UpdateTopicResponse struct {
	*BaseResponse
}

func (client *SmnClient) UpdateTopic(request *UpdateTopicRequest) (response *UpdateTopicResponse, err error) {
	response = &UpdateTopicResponse{
		BaseResponse: &BaseResponse{},
	}
	err = client.sendRequest(request, response)
	return
}

func (client *SmnClient) NewUpdateTopicRequest() (request *UpdateTopicRequest) {
	request = &UpdateTopicRequest{
		BaseRequest: &BaseRequest{Headers: make(map[string]string)},
	}
	return
}

func (request *UpdateTopicRequest) GetUrl() (string, error) {
	if request.TopicUrn == "" {
		return "", fmt.Errorf("topic urn is null")
	}

	if !util.ValidateTopicDiplayName(request.DisplayName) {
		return "", fmt.Errorf("topic displayName is overSize")
	}

	return request.BaseRequest.GetSmnServiceUrl() + util.V2Version + util.UrlDelimiter + request.projectId +
		util.UrlDelimiter + util.Notifications + util.UrlDelimiter + util.Topics +
		util.UrlDelimiter + request.TopicUrn, nil
}

func (request *UpdateTopicRequest) GetMethod() string {
	return util.PUT
}

func (request *UpdateTopicRequest) GetBodyReader() (reader io.Reader, err error) {
	return util.GetBodyParams(request)
}

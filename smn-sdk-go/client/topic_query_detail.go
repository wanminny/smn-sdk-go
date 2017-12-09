package client

import (
	"smn-sdk-go/util"
	"io"
	"fmt"
)

type QueryTopicDetailRequest struct {
	*BaseRequest
	TopicUrn string
}

type QueryTopicDetailResponse struct {
	*BaseResponse
	TopicUrn    string `json:"topic_urn"`
	DisplayName string `json:"display_name"`
	Name        string `json:"name"`
	PushPolicy  int    `json:"push_policy"`
	CreateTime  string `json:"create_time"`
	UpdateTime  string `json:"update_time"`
}

func (client *SmnClient) QueryTopicDetail(request *QueryTopicDetailRequest) (response *QueryTopicDetailResponse, err error) {
	response = &QueryTopicDetailResponse{
		BaseResponse: &BaseResponse{},
	}
	err = client.sendRequest(request, response)
	return
}

func (client *SmnClient) NewQueryTopicDetailRequest() (request *QueryTopicDetailRequest) {
	request = &QueryTopicDetailRequest{
		BaseRequest: &BaseRequest{Headers: make(map[string]string)},
	}
	return
}

func (request *QueryTopicDetailRequest) GetUrl() (string, error) {
	if request.TopicUrn == "" {
		return "", fmt.Errorf("topic urn is null")
	}

	return request.BaseRequest.GetSmnServiceUrl() + util.V2Version + util.UrlDelimiter + request.projectId +
		util.UrlDelimiter + util.Notifications + util.UrlDelimiter + util.Topics +
		util.UrlDelimiter + request.TopicUrn, nil
}

func (request *QueryTopicDetailRequest) GetMethod() string {
	return util.GET
}

func (request *QueryTopicDetailRequest) GetBodyReader() (reader io.Reader, err error) {
	return nil, nil
}

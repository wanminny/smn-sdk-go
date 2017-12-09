package util

import (
	"regexp"
)

const (
	MaxTopicDisplayName = 192
	PhoneMatchPattern   = "^\\+?[0-9]{1,31}"
	TopicNamePattern    = "^[a-zA-Z0-9]{1}[-_a-zA-Z0-9]{0,255}$"
)

func ValidatePhone(phone string) bool {
	reg := regexp.MustCompile(PhoneMatchPattern)
	return reg.MatchString(phone)
}

func ValidateSmsEventType(eventType string) bool {
	return eventType == SmsFailEvent || eventType == SmsReplyEvent ||
		eventType == SmsSuccessEvent
}

func ValidateTopicName(topicName string) bool {
	if topicName == "" {
		return false
	}
	reg := regexp.MustCompile(TopicNamePattern)
	return reg.MatchString(topicName)
}

func ValidateTopicDiplayName(displayName string) bool {
	if displayName == "" {
		return true
	}
	bytes := []byte(displayName)
	return len(bytes) < MaxTopicDisplayName
}
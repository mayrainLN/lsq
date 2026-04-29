package service

import "ai-wordbook/config"

type DeepSeekProvider struct{}

func (p *DeepSeekProvider) QueryWord(word string) (*WordResult, error) {
	return callChatAPI(
		config.AppConfig.DeepSeekBaseURL,
		config.AppConfig.DeepSeekAPIKey,
		"deepseek-chat",
		word,
	)
}

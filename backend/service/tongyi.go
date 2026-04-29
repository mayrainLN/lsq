package service

import "ai-wordbook/config"

type TongyiProvider struct{}

func (p *TongyiProvider) QueryWord(word string) (*WordResult, error) {
	return callChatAPI(
		config.AppConfig.TongyiBaseURL,
		config.AppConfig.TongyiAPIKey,
		"qwen-plus",
		word,
	)
}

package service

import "fmt"

type WordResult struct {
	Definition string           `json:"definition"`
	Sentences  []SentenceResult `json:"sentences"`
}

type SentenceResult struct {
	English string `json:"english"`
	Chinese string `json:"chinese"`
}

type AIProvider interface {
	QueryWord(word string) (*WordResult, error)
}

func GetProvider(name string) (AIProvider, error) {
	switch name {
	case "deepseek":
		return &DeepSeekProvider{}, nil
	case "tongyi":
		return &TongyiProvider{}, nil
	default:
		return nil, fmt.Errorf("不支持的 AI 提供商: %s", name)
	}
}

package postapi

import (
	"github.com/dev-baris/ai-bot/vectordb"

	cache "github.com/dev-baris/ai-bot/cache"
	openai "github.com/sashabaranov/go-openai"
)

type HandlerContext struct {
	openAIClient *openai.Client
	cache        *cache.Cache
	vectorDB     vectordb.VectorDB
}

func NewHandlerContext(openAIClient *openai.Client, vectorDB vectordb.VectorDB) *HandlerContext {
	return &HandlerContext{
		openAIClient: openAIClient,
		cache:        cache.New(cache.NoExpiration, cache.NoExpiration),
		vectorDB:     vectorDB,
	}
}

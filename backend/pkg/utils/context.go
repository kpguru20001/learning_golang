package utils

import (
	"context"
	"fmt"
	"shadow-docs/configs"

	"github.com/gin-gonic/gin"
)

func StandardContextTimeout() context.Context {
	ctx, cancel := context.WithTimeout(context.Background(), configs.Constants.StandardContextTimeout)
	defer cancel()

	return ctx
}

func GinContextFromContext(ctx context.Context) (*gin.Context, error) {
	ginContext := ctx.Value(configs.Constants.GinContextKey)
	if ginContext == nil {
		err := fmt.Errorf("could not retrieve gin.Context")
		return nil, err
	}

	gc, ok := ginContext.(*gin.Context)
	if !ok {
		err := fmt.Errorf("gin.Context has wrong type")
		return nil, err
	}
	return gc, nil
}

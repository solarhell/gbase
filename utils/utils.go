package utils

import (
	stdcontext "context"

	"go.uber.org/zap"

	"github.com/solarhell/gbase/context"
)

func SimpleContext() context.Context {
	var logger, _ = zap.NewDevelopment()
	return context.New(
		stdcontext.Background(),
		context.NewEnv(),
		context.NewLogger("", "", logger.Sugar(), nil, nil),
	)
}

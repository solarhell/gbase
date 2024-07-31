package gbase

import (
	stdcontext "context"
	"strconv"
	"sync/atomic"

	"go.uber.org/zap"

	"github.com/solarhell/gbase/context"
)

// type alias
type (
	Context = context.Context
)

// SessionNameGenerator used to generate session name automatically,
// you can replace it
var SessionNameGenerator = func() func() string {
	var counter uint64
	return func() string {
		// return bootid related sequential name
		const h = "000000000000"
		var seq = atomic.AddUint64(&counter, 1)
		var s = strconv.FormatUint(seq, 10)
		return BootID[:24] + h[:12-len(s)] + s
	}
}()

// SimpleContext return new context without name
func SimpleContext() Context {
	return context.Simple()
}

// SessionContext return new context use auto session name that
// generated by SessionNameGenerator
func SessionContext() Context {
	return NamedContext(SessionNameGenerator())
}

// NamedContext return new context use specified session name
func NamedContext(name string) Context {
	return context.New(stdcontext.Background(), context.NewEnv(),
		context.NewLogger(name, "", zap.S(), nil, nil),
	)
}

// ToSimpleContext return new context use specified official context without name
func ToSimpleContext(gctx stdcontext.Context) Context {
	return ToNamedContext(gctx, "")
}

// ToSessionContext return new context use specified official context with auto session name that
// generated by SessionNameGenerator
func ToSessionContext(gctx stdcontext.Context) Context {
	return ToNamedContext(gctx, SessionNameGenerator())
}

// ToNamedContext return new context use specified official context and name
func ToNamedContext(gctx stdcontext.Context, name string) Context {
	return context.New(gctx, context.NewEnv(),
		context.NewLogger(name, "", zap.S(), nil, nil),
	)
}

// SetSession use to set real session name
func SetSession(ctx Context, session string) {
	context.SetSession(ctx, session)
}

// GetSession use to get session name if it has real session, otherwise return context name instead
func GetSession(ctx Context) string {
	return context.GetSession(ctx)
}

// GetRealSession use to get real session name
func GetRealSession(ctx Context) string {
	return context.GetRealSession(ctx)
}

package gbase

import (
	"github.com/google/uuid"
)

var (
	BootID = uuid.New().String()
)

func init() {
	initZapLogger()
}

// initZapLogger should used when your process startup,
// and you should replace it with a new logger policy after parsing your config.
// If you do not like it's style, just use zap.ReplaceGlobals() by you self, before using Context
func initZapLogger() {
	var _, err = ReplaceZapLogger("debug", "stderr", "console", false)
	if err != nil {
		panic(err)
	}
}

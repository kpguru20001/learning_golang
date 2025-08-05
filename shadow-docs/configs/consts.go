package configs

import "time"

type ConstantsStruct struct {
	StandardContextTimeout time.Duration
	GinContextKey          string
}

var Constants = ConstantsStruct{
	StandardContextTimeout: 10 * time.Second,
	GinContextKey:          "GinCtx",
}

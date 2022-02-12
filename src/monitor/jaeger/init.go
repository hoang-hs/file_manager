package jaeger

import (
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegerCfg "github.com/uber/jaeger-client-go/config"
)

func InitJaeger() {
	cfg := jaegerCfg.Configuration{
		ServiceName: "file_manager_jaeger",
		Sampler: &jaegerCfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegerCfg.ReporterConfig{
			LogSpans: true,
		},
	}
	tracer, _, err := cfg.NewTracer(
		jaegerCfg.Logger(jaeger.StdLogger),
	)
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}
	opentracing.SetGlobalTracer(tracer)
}

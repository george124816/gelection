package otel

import (
	"context"
	"io"
	defaultLog "log"
	"os"
	"time"

	"github.com/george124816/gelection/internal/configs"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploghttp"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp"
	otellog "go.opentelemetry.io/otel/log"
	"go.opentelemetry.io/otel/log/global"
	"go.opentelemetry.io/otel/sdk/log"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.37.0"
)

type otelWriter struct {
	logger otellog.Logger
	level  otellog.Severity
}

func (ow *otelWriter) Write(p []byte) (n int, err error) {
	ctx := context.Background()

	rec := otellog.Record{}
	rec.SetBody(otellog.StringValue(string(p)))
	rec.SetSeverity(otellog.SeverityInfo)
	rec.SetSeverityText("INFO")

	ow.logger.Emit(ctx, rec)
	return len(p), nil
}

func StartMetrics() error {
	ctx := context.Background()
	exp, err := otlpmetrichttp.New(ctx,
		otlpmetrichttp.WithEndpoint(configs.GetOtelConfig().String()),
		otlpmetrichttp.WithURLPath("/v1/metrics"),
		otlpmetrichttp.WithInsecure(),
	)

	if err != nil {
		return err
	}

	res, err := resource.Merge(resource.Default(),
		resource.NewWithAttributes(semconv.SchemaURL,
			semconv.ServiceName("gelection"),
			semconv.ServiceVersion("0.0.1"),
		))

	if err != nil {
		return err
	}

	meterProvider := metric.NewMeterProvider(
		metric.WithResource(res),
		metric.WithReader(
			metric.NewPeriodicReader(exp,
				metric.WithInterval(1*time.Second),
			)))

	otel.SetMeterProvider(meterProvider)

	return nil
}

func StartLogs() error {
	ctx := context.Background()

	res, err := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName("gelection"),
			semconv.ServiceVersion("0.0.1"),
		),
	)
	if err != nil {
		return err
	}

	exporter, err := otlploghttp.New(ctx,
		otlploghttp.WithEndpoint(configs.GetOtelConfig().String()),
		otlploghttp.WithInsecure(),
		otlploghttp.WithURLPath("/v1/logs"),
	)
	if err != nil {
		return err
	}

	processor := log.NewBatchProcessor(exporter)
	provider := log.NewLoggerProvider(
		log.WithResource(res),
		log.WithProcessor(processor),
	)

	global.SetLoggerProvider(provider)

	RedirectStdLogToOtel()

	return nil
}

func RedirectStdLogToOtel() {
	otelLogger := global.GetLoggerProvider().Logger("stdlib")
	ow := &otelWriter{logger: otelLogger}
	mw := io.MultiWriter(os.Stderr, ow)
	defaultLog.SetOutput(mw)
}

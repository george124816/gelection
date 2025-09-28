package otel

import (
	"context"
	"log/slog"
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

type OtelHandler struct {
	logger otellog.Logger
	level  slog.Level
}

type MultiHandler struct {
	handlers []slog.Handler
}

func (h *MultiHandler) Enabled(ctx context.Context, lvl slog.Level) bool {
	for _, hh := range h.handlers {
		if hh.Enabled(ctx, lvl) {
			return true
		}
	}
	return false
}

func (h *MultiHandler) Handle(ctx context.Context, r slog.Record) error {
	for _, hh := range h.handlers {
		_ = hh.Handle(ctx, r)
	}
	return nil
}

func (h *MultiHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	newHandlers := make([]slog.Handler, len(h.handlers))
	for i, hh := range h.handlers {
		newHandlers[i] = hh.WithAttrs(attrs)
	}
	return &MultiHandler{handlers: newHandlers}
}

func (h *MultiHandler) WithGroup(name string) slog.Handler {
	newHandlers := make([]slog.Handler, len(h.handlers))
	for i, hh := range h.handlers {
		newHandlers[i] = hh.WithGroup(name)
	}
	return &MultiHandler{handlers: newHandlers}
}

func NewOtelHandler(logger otellog.Logger, level slog.Level) *OtelHandler {
	return &OtelHandler{logger: logger, level: level}
}

func (h *OtelHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return h
}

func (h *OtelHandler) WithGroup(name string) slog.Handler {
	return h
}

func (h *OtelHandler) Enabled(ctx context.Context, lvl slog.Level) bool {
	return lvl >= h.level
}
func (h *OtelHandler) Handle(ctx context.Context, r slog.Record) error {
	rec := otellog.Record{}
	rec.SetTimestamp(r.Time)
	rec.SetSeverityText(r.Level.String())
	rec.SetSeverity(mapSlogLevelToOTEL(r.Level))

	// build body
	msg := r.Message
	rec.SetBody(otellog.StringValue(msg))

	// add attributes as OTEL attributes
	r.Attrs(func(a slog.Attr) bool {
		// a.Key and a.Value
		rec.AddAttributes(otellog.KeyValue{
			Key:   a.Key,
			Value: otellog.StringValue(a.Value.String()),
		})
		return true
	})

	h.logger.Emit(ctx, rec)
	return nil
}

func mapSlogLevelToOTEL(level slog.Level) otellog.Severity {
	switch level {
	case slog.LevelDebug:
		return otellog.SeverityDebug
	case slog.LevelInfo:
		return otellog.SeverityInfo
	case slog.LevelWarn:
		return otellog.SeverityWarn
	case slog.LevelError:
		return otellog.SeverityError
	default:
		return otellog.SeverityInfo
	}
}

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

	InitSlogWithOtel()

	return nil
}

func InitSlogWithOtel() {
	otelLogger := global.GetLoggerProvider().Logger("main")
	otelHandler := NewOtelHandler(otelLogger, slog.LevelInfo)

	stdoutHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})

	// combine both
	handler := &MultiHandler{handlers: []slog.Handler{stdoutHandler, otelHandler}}
	slog.SetDefault(slog.New(handler))
}

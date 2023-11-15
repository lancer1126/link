package conf

import (
	"github.com/alimy/tryst/cfg"
	"github.com/getsentry/sentry-go"
	"link/internal/version"
	"time"
)

func initSentry() {
	cfg.Be("Sentry", func() {
		opts := sentry.ClientOptions{
			Dsn:              sentrySetting.Dsn,
			Debug:            sentrySetting.Debug,
			AttachStacktrace: sentrySetting.AttachStacktrace,
			TracesSampleRate: sentrySetting.TracesSampleRate,
		}

		_ = sentry.Init(opts)
		if sentrySetting.AttachLogrus {
			setupSentryLogrus(opts)
		}
		sentry.WithScope(func(scope *sentry.Scope) {
			scope.SetExtras(map[string]any{
				"version": version.Info(),
				"time":    time.Now().Local(),
			})
			sentry.CaptureMessage("link sentry works!")
		})
	})
}

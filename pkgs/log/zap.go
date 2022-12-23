package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func getConfig() (c zap.Config) {
	c = zap.Config{
		Level:             zap.NewAtomicLevelAt(zap.DebugLevel),
		Development:       true,
		DisableCaller:     false,
		DisableStacktrace: true,
		Sampling:          nil,
		Encoding:          "console",
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:          "M",
			LevelKey:            "L",
			TimeKey:             "T",
			NameKey:             "N",
			CallerKey:           "C",
			FunctionKey:         "F",
			StacktraceKey:       "S",
			SkipLineEnding:      false,
			LineEnding:          zapcore.DefaultLineEnding,
			EncodeLevel:         zapcore.CapitalColorLevelEncoder,
			EncodeTime:          zapcore.RFC3339TimeEncoder,
			EncodeDuration:      zapcore.StringDurationEncoder,
			EncodeCaller:        zapcore.FullCallerEncoder,
			EncodeName:          nil,
			NewReflectedEncoder: nil,
			ConsoleSeparator:    "",
		},
		InitialFields:    nil,
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}

	if !c.Development {
		c.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("")
		c.EncoderConfig.EncodeCaller = nil
		c.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	}

	return
}

var cfg = getConfig()

func getLogger() (z *zap.Logger) {
	z, err := cfg.Build()

	if err != nil {
		panic(err)
	}
	return
}

func Get() *zap.SugaredLogger {
	return getLogger().Sugar()
}

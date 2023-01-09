package logger

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLogger(t *testing.T) {
	cases := map[string]struct {
		level    LogLevel
		expected []string
	}{
		"ErrorLog": {
			level:    Error,
			expected: []string{"error"},
		},
		"WarnLog": {
			level:    Warn,
			expected: []string{"error", "warn"},
		},
		"InfoLog": {
			level:    Info,
			expected: []string{"error", "warn", "info"},
		},
		"DebugLog": {
			level:    Debug,
			expected: []string{"error", "warn", "info", "debug"},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			b := make([]byte, 0, 1024)
			buf := bytes.NewBuffer(b)

			defaultLogger.output = buf
			SetLevel(tc.level)
			defaultLogger.Error("error")
			defaultLogger.Warn("warn")
			defaultLogger.Info("info")
			defaultLogger.Debug("debug")

			fmt.Fprint(os.Stderr, buf.String())
			for _, s := range tc.expected {
				require.Contains(t, buf.String(), s)
			}
		})
	}
}

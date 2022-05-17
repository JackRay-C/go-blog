package logger

import (
	"blog/internal/config"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSimple(t *testing.T)  {
	l, err := NewSimpleLogger(&config.App{AppLogType: SimpleLog, Logs: &config.Logs{Simple: &config.Simple{
		Level:        LevelInfo.String(),
		LogInConsole: true,
		Directory:    "logs",
		FileName:     "debug.log",
	}}} )
	if err != nil {
		t.Log(err)
	}

	l.Infof("infof")
	l.Debug("debugf")
	l.Warnf("warnf")

}

func Test(t *testing.T) {
	testCases := []struct {
		name string
		args *config.App
		want Logger
		err  error
	}{
		{
			name: "config logger new simple logger",
			args: &config.App{
				AppLogType: SimpleLog,
				Logs: &config.Logs{
					Simple: &config.Simple{
						Level:        LevelInfo.String(),
						LogInConsole: true,
						Directory:    "logs",
						FileName:     "debug.log",
					},
					Zap: nil,
				},
			},
			want: &SimpleLogger{},
			err:  nil,
		},
		{
			name: "config zap new zap logger",
			args: &config.App{
				AppLogType: ZapLog,
				Logs: &config.Logs{
					Simple:nil,
					Zap: &config.Zap{
						Level:         "info",
						Format:        "console",
						Directory:     "logs",
						LinkName:      "debug-zap.log",
						ShowLine:      true,
						LogInConsole:  true,
						LogMaxSize:    100,
						LogMaxAge:     30,
						LogMaxBackups: 5,
					},
				},
			},
			want: &ZapLogger{},
			err:  nil,
		},
		{
			name: "un support config type",
			args: &config.App{},
			want: nil,
			err: errors.New("un support logger type. "),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			want := tc.want
			l, err := New(tc.args)

			if err != nil {
				if !assert.Equal(t, tc.err, err) {
					t.Errorf("expect nil error, get %s", err)
				}
			}
			if err == nil {
				l.Infof("infof")
				l.Errorf("errorf")
				l.Warnf("warnf")
				if !assert.Equal(t, tc.want, want) {
					t.Errorf("expect %v, got %v", tc.want, l)
				}
			}
		})
	}
}


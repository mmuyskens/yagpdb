package common

import (
	"bytes"
	"net/http"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/jonas747/discordgo"
	"github.com/sirupsen/logrus"
)

type ContextHook struct{}

func (hook ContextHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (hook ContextHook) Fire(entry *logrus.Entry) error {
	// Skip if already provided
	if _, ok := entry.Data["stck"]; ok {
		return nil
	}

	pc := make([]uintptr, 3)
	cnt := runtime.Callers(6, pc)

	for i := 0; i < cnt; i++ {
		fu := runtime.FuncForPC(pc[i] - 1)
		name := fu.Name()
		if !strings.Contains(name, "github.com/sirupsen/logrus") {
			file, line := fu.FileLine(pc[i] - 1)

			entry.Data["stck"] = filepath.Base(name) + ":" + filepath.Base(file) + ":" + strconv.Itoa(line)
			break
		}
	}
	return nil
}

type STDLogProxy struct{}

func (p *STDLogProxy) Write(b []byte) (n int, err error) {

	// Check to see if this is from discordgo
	isDiscord := bytes.HasPrefix(b, []byte("[DG"))
	discordLogLevel := 2
	if isDiscord && len(b) > 4 {

		parsedLevel, err := strconv.Atoi(string(b[3]))
		if err == nil {
			discordLogLevel = parsedLevel
		}
	}

	n = len(b)

	data := make(logrus.Fields)

	if !isDiscord {
		pc := make([]uintptr, 3)
		runtime.Callers(4, pc)

		fu := runtime.FuncForPC(pc[0] - 1)
		name := fu.Name()
		file, line := fu.FileLine(pc[0] - 1)

		data["stck"] = filepath.Base(name) + ":" + filepath.Base(file) + ":" + strconv.Itoa(line)
	} else {
		data["stck"] = "" // prevent upstream from adding it
	}

	logLine := string(b)
	if strings.HasSuffix(logLine, "\n") {
		logLine = strings.TrimSuffix(logLine, "\n")
	}

	f := logrus.WithFields(data)

	switch discordLogLevel {
	case 0:
		f.Error(logLine)
	case 1:
		f.Warn(logLine)
	default:
		f.Info(logLine)
	}

	return
}

type GORMLogger struct {
}

func (g *GORMLogger) Print(v ...interface{}) {
	logrus.WithField("stck", "...").Error(v...)
}

type LoggingTransport struct {
	Inner http.RoundTripper
}

var numberRemover = strings.NewReplacer(
	"0", "",
	"1", "",
	"2", "",
	"3", "",
	"4", "",
	"5", "",
	"6", "",
	"7", "",
	"8", "",
	"9", "")

func (t *LoggingTransport) RoundTrip(request *http.Request) (*http.Response, error) {

	bucketI := request.Context().Value(discordgo.CtxKeyRatelimitBucket)
	var rlBucket *discordgo.Bucket
	if bucketI != nil {
		rlBucket = bucketI.(*discordgo.Bucket)
	}

	inner := t.Inner
	if inner == nil {
		inner = http.DefaultTransport
	}

	started := time.Now()

	code := 0
	resp, err := inner.RoundTrip(request)
	if resp != nil {
		code = resp.StatusCode
	}

	since := time.Since(started).Seconds() * 1000
	go func() {
		path := request.URL.Path
		if rlBucket != nil {
			path = rlBucket.Key
		}

		path = numberRemover.Replace(path)

		if Statsd != nil {
			Statsd.Incr("discord.num_requests", []string{"method:" + request.Method, "resp_code:" + strconv.Itoa(code), "path:" + request.Method + "-" + path}, 1)
			Statsd.Gauge("discord.http_latency", since, nil, 1)
			if code == 429 {
				Statsd.Incr("discord.requests.429", []string{"method:" + request.Method, "path:" + request.Method + "-" + path}, 1)
			}
		}

		if since > 5000 {
			logrus.WithField("path", request.URL.Path).WithField("ms", since).WithField("method", request.Method).Warn("Request took longer than 5 seconds to complete!")
		}

		// Statsd.Incr("discord.response.code."+strconv.Itoa(floored), nil, 1)
		// Statsd.Incr("discord.request.method."+request.Method, nil, 1)
	}()

	return resp, err
}

func AddLogHook(hook logrus.Hook) {
	logrus.AddHook(hook)
}

func SetLoggingLevel(level logrus.Level) {
	logrus.SetLevel(level)
}

func SetLogFormatter(formatter logrus.Formatter) {
	logrus.SetFormatter(formatter)
}

func GetPluginLogger(plugin Plugin) *logrus.Entry {
	info := plugin.PluginInfo()
	return logrus.WithField("p", info.SysName)
}

func GetFixedPrefixLogger(prefix string) *logrus.Entry {
	return logrus.WithField("p", prefix)
}

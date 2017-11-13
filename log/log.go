package log

import (
	"context"
	"fmt"
	"io"

	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

func Setup(writter io.Writer) {
	logger.Out = writter
	logger.Level = logrus.DebugLevel
	logger.Formatter = &logrus.TextFormatter{FullTimestamp: true}
}

func getContextString(ctx context.Context, key string) string {
	value, ok := ctx.Value(key).(string)
	if !ok {
		return ""
	}
	return value
}

func withContextFields(ctx context.Context, fields []interface{}) logrus.Fields {
	size := len(fields)
	if size%2 != 0 {
		fields = append(fields, nil)
		size = size + 1
	}
	data := make(logrus.Fields, (size/2)+1)
	ID := getContextString(ctx, "ID")
	data["ID"] = ID
	for i := 0; i < size/2; i++ {
		key, ok := fields[2*i].(string)
		if !ok {
			key = fmt.Sprintf("unknown%d", i)
		}
		data[key] = fields[2*i+1]
	}
	return data
}

func Debug(ctx context.Context, message string, fields ...interface{}) {
	fieldMap := withContextFields(ctx, fields)
	logger.WithFields(fieldMap).Debug(message)
}

func Info(ctx context.Context, message string, fields ...interface{}) {
	fieldMap := withContextFields(ctx, fields)
	logger.WithFields(fieldMap).Info(message)
}

func Warn(ctx context.Context, message string, fields ...interface{}) {
	fieldMap := withContextFields(ctx, fields)
	logger.WithFields(fieldMap).Warn(message)
}

func Error(ctx context.Context, message string, fields ...interface{}) {
	fieldMap := withContextFields(ctx, fields)
	logger.WithFields(fieldMap).Error(message)
}

func Fatal(ctx context.Context, message string, fields ...interface{}) {
	fieldMap := withContextFields(ctx, fields)
	logger.WithFields(fieldMap).Fatal(message)
}

func Panic(ctx context.Context, message string, fields ...interface{}) {
	fieldMap := withContextFields(ctx, fields)
	logger.WithFields(fieldMap).Panic(message)
}

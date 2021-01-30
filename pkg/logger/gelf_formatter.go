package logger

import (
	"encoding/json"

	"github.com/sirupsen/logrus"
)

type fields map[string]interface{}

type GelfFormatter struct {
	fields fields
}

// Format formats the log entry to GELF JSON.
func (f *GelfFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	data := make(fields, len(entry.Data)+6)
	blacklist := []string{"_id", "id", "timestamp", "version", "level"}

	for k, v := range entry.Data {
		if _, ok := f.fields[k]; ok {
			continue
		}

		if contains(k, blacklist) {
			continue
		}

		switch v := v.(type) {
		case error:
			data["_"+k] = v.Error()
		default:
			data["_"+k] = v
		}
	}

	for k, v := range f.fields {
		data[k] = v
	}

	data["version"] = "1.1"
	data["short_message"] = entry.Message
	data["full_message"] = entry.Message
	data["timestamp"] = entry.Time.UnixNano() / 1000000
	data["level"] = entry.Level

	serialized, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return append(serialized, '\n'), nil
}

func contains(key string, blacklist []string) bool {
	for _, v := range blacklist {
		if key == v {
			return true
		}
	}

	return false
}

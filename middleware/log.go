package middleware

import (
	"bytes"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

type logMessage struct {
	ServiceId  string `json:"service_id,omitempty"`
	LogType    string `json:"log_type"`
	TimeStamp  int64  `json:"timestamp"`
	Latency    int64  `json:"latency"`
	ClientIP   string `json:"client_ip"`
	Method     string `json:"method"`
	Path       string `json:"path"`
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
	BodySize   int    `json:"body_size"`
}

func Log(serviceId ...string) gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		sid := ""
		if len(serviceId) > 0 {
			sid = serviceId[0]
		}

		buffer := &bytes.Buffer{}
		encoder := json.NewEncoder(buffer)
		encoder.SetEscapeHTML(false)
		_ = encoder.Encode(logMessage{
			ServiceId:  sid,
			LogType:    "access",
			TimeStamp:  params.TimeStamp.Unix(),
			Latency:    params.Latency.Milliseconds(),
			ClientIP:   params.ClientIP,
			Method:     params.Method,
			Path:       params.Path,
			Message:    params.ErrorMessage,
			StatusCode: params.StatusCode,
			BodySize:   params.BodySize,
		})
		return buffer.String()
	})
}

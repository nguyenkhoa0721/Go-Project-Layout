package log

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"runtime"
	"strings"
)

func init() {
	logrus.SetReportCaller(true)
	formatter := &logrus.TextFormatter{
		TimestampFormat:        "2006-01-02 15:04:05", // the "time" field configuratiom
		FullTimestamp:          true,
		DisableLevelTruncation: true, // log level field configuration
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			return "", fmt.Sprintf(" %s:%d", formatFilePath(f.File), f.Line)
		},
	}
	logrus.SetFormatter(formatter)
}
func formatFilePath(path string) string {
	arr := strings.Split(path, "/")
	return arr[len(arr)-1]
}

func GrpcLogger(
	result interface{},
	err error,
) (interface{}, error) {
	statusCode := codes.Unknown
	if st, ok := status.FromError(err); !ok {
		statusCode = st.Code()
	}

	logger := logrus.WithFields(logrus.Fields{
		"status code": int(statusCode),
		"status text": statusCode.String(),
	})
	if err != nil {
		logger.Error(err)
	} else {
		logger.Info("received a gRPC request")
	}

	return result, err
}

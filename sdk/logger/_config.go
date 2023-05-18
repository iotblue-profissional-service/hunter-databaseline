package logger

import "os"

var (
	logFilesTimeInDays = os.Getenv("log_file_time_in_days")
	enableLogs         = os.Getenv("enable_logs")
)

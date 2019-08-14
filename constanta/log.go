package constanta

// LogFormat is global constanta for logging
type LogFormat string

// Add function to get value in string for LogFormat
func (lf LogFormat) String() string {
	return string(lf)
}

// This constanta util for LogFormat
const (
	LFDebug   LogFormat = "debuginfo"
	LFValuer  LogFormat = "valuer"
	LFContext LogFormat = "context"
)

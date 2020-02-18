package logging

// Option is option for logging
type Option struct {
	LogLevel Level
	Format   FormatLog
	Ouput    OutputLog
}

// IsEmpty is checking struct option is empty or not
func (et Option) IsEmpty() bool {
	var opt Option
	if et == opt {
		return true
	}
	return false
}
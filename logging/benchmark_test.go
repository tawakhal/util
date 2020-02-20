package logging

import (
	"testing"
)

func BenchmarkConsoleLog(b *testing.B) {
	opt := Option{
		LogLevel: LevelALL,
		Format:   FormatFMT,
		Ouput:    Console,
	}

	logger := NewConsoleLog(opt)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Info(i, i)
	}
}

func BenchmarkLogger(b *testing.B) {
	opt := Option{
		LogLevel: LevelALL,
		Format:   FormatFMT,
		Ouput:    Console,
	}

	logger := NewLogger(opt)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Info(i, i)
	}
}

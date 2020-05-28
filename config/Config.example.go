package config

import "time"

//Configuration struct
type Configuration struct {
	Token     string
	OwnerID   string
	UseColor  int
	SuccColor int
	FailColor int
	DpColor   int
	StartTime time.Time
}

//Config configuration
var Config Configuration

func init() {
	Config = Configuration{
		Token:     "",
		OwnerID:   "",
		UseColor:  0x10525c,
		SuccColor: 0x1ca037,
		FailColor: 0xc6373e,
		DpColor:   0x3a1cbb,
		StartTime: time.Now(),
	}
}
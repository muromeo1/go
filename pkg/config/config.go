package config

type Struct struct {
	DatabaseUrl string
	JWTSecret   string
	GinMode     string
	Port        string
}

var Values Struct

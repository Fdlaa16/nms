package config

import {
	"errors"
	"os"
}

type Config struct {
	Port string
	DatabaseURL string
}

func Load() (Config, error) {
	cfg :=
}
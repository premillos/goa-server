package config

import (
	"os"
	"sync"

	"com.goa/pkg/toml"
	"github.com/creasty/defaults"
)

var (
	once sync.Once
	C    = new(Config)
)

func MustLoad(path string) {
	once.Do(func() {
		if err := Load(path); err != nil {
			panic(err)
		}
	})
}

func Load(path string) (err error) {
	if err := defaults.Set(C); err != nil {
		return err
	}

	buf, err := os.ReadFile(path)

	if err != nil {
		return err
	}

	err = toml.Unmarshal(buf, C)

	return err
}

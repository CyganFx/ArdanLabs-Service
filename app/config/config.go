// Package config contains web api settings
package config

import (
	"fmt"
	"github.com/ardanlabs/conf"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

type Config struct {
	conf.Version
	Web struct {
		APIHost         string        `conf:"default:0.0.0.0:3000"`
		DebugHost       string        `conf:"default:0.0.0.0:4000"`
		ReadTimeout     time.Duration `conf:"default:5s"`
		WriteTimeout    time.Duration `conf:"default:5s"`
		ShutdownTimeout time.Duration `conf:"default:5s"`
	}
	Auth struct {
		KeyID          string `conf:"default:54bb2165-71e1-41a6-af3e-7da4a0e1e2c1"`
		PrivateKeyFile string
		Algorithm      string `conf:"default:RS256"`
	}
	DB struct {
		User       string `conf:"default:postgres"`
		Password   string `conf:"default:postgres,noprint" env:"DB_PASS"`
		Host       string `conf:"default:0.0.0.0"`
		Name       string `conf:"default:postgres" yaml:"name"`
		DisableTLS bool   `conf:"default:true"`
	} `yaml:"db"`
}

func Init(cfg *Config, configsDir, prefix, build string) error {

	//You should be in root directory in order to run this successfully
	privateKeyFilePath, err := filepath.Abs("./private.pem")
	if err != nil {
		return errors.Wrap(err, "getting private key path")
	}

	cfg.Auth.PrivateKeyFile = privateKeyFilePath
	cfg.Version.SVN = build
	cfg.Version.Desc = "copyright information here"

	if err := conf.Parse(os.Args[1:], prefix, cfg); err != nil {
		switch err {
		case conf.ErrHelpWanted:
			usage, err := conf.Usage(prefix, cfg)
			if err != nil {
				return errors.Wrap(err, "generating config usage")
			}
			fmt.Println(usage)
			return nil
		case conf.ErrVersionWanted:
			version, err := conf.VersionString(prefix, cfg)
			if err != nil {
				return errors.Wrap(err, "generating config version")
			}
			fmt.Println(version)
			return nil
		}
		return errors.Wrap(err, "parsing config")
	}

	content, err := ioutil.ReadFile(configsDir)
	if err != nil {
		return errors.Wrap(err, "reading file")
	}

	if err = yaml.Unmarshal(content, cfg); err != nil {
		return errors.Wrap(err, "unmarshalling config")
	}

	if err = cleanenv.ReadEnv(cfg); err != nil {
		return errors.Wrap(err, "reading env")
	}

	return nil
}

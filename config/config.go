package config

import (
	"gopkg.in/ini.v1"
	"log"
	"os"
	"time"
)

type ConfigList struct {
	ApiKey      string
	ApiSecret   string
	LogFile     string
	ProductCode string

	TradeDuration time.Duration
	Durations     map[string]time.Duration
	DbName        string
	SQLDriver     string
	Port          int

	BackTest         bool
	UsePercent       float64
	DataLimit        int
	StopLimitPercent float64
	NumRanking       int
}

var Config ConfigList

func init() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Printf("Failed to read file: %v", err)
		os.Exit(1)
	}

	durations := map[string]time.Duration{
		"1s": time.Second,
		"1m": time.Minute,
		"1h": time.Hour,
	}

	Config = ConfigList{
		ApiKey:           cfg.Section("bitflyer").Key("api_key").String(),
		ApiSecret:        cfg.Section("bitflyer").Key("api_secret").String(),
		LogFile:          cfg.Section("go-trading").Key("log_file").String(),
		ProductCode:      cfg.Section("go-trading").Key("product_code").String(),
		Durations:        durations,
		TradeDuration:    durations[cfg.Section("go-trading").Key("trade_duration").String()],
		DbName:           cfg.Section("db").Key("name").String(),
		SQLDriver:        cfg.Section("db").Key("driver").String(),
		Port:             cfg.Section("web").Key("port").MustInt(),
		BackTest:         cfg.Section("go-trading").Key("back_test").MustBool(),
		UsePercent:       cfg.Section("go-trading").Key("use_percent").MustFloat64(),
		DataLimit:        cfg.Section("go-trading").Key("data_limit").MustInt(),
		StopLimitPercent: cfg.Section("go-trading").Key("stop_limit_percent").MustFloat64(),
		NumRanking:       cfg.Section("go-trading").Key("num_ranking").MustInt(),
	}
}

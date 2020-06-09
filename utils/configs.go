package utils

import (
	"fmt"
	"github.com/go-ini/ini"
	"os"
	"strings"
)

type Config struct {
	Country string
	OriginId string
	DestinationId string
	DepartureDate string
	MinDate string
	Days string
	TelegramApiKey string
	TelegramChannelId string
}

func LoadConfig() (Config)  {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var config Config
	config.Country = strings.ToLower(cfg.Section("megabus").Key("country").String())
	config.OriginId = cfg.Section("megabus").Key("originId").String()
	config.DestinationId = cfg.Section("megabus").Key("destinationId").String()
	config.DepartureDate = cfg.Section("megabus").Key("departureDate").String()
	config.MinDate = cfg.Section("megabus").Key("minDate").String()
	config.Days = cfg.Section("megabus").Key("days").String()
	config.TelegramApiKey = cfg.Section("telegram").Key("api_key").String()
	config.TelegramChannelId = cfg.Section("telegram").Key("channel_id").String()
	return config
}

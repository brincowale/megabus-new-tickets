package main

import (
	"encoding/json"
	"fmt"
	"github.com/brincowale/go-telegram-sender"
	"github.com/parnurzeal/gorequest"
	"megabus-new-tickets/utils"
	"net/http"
	"os"
	"time"
)

type Tickets struct {
	Dates []struct {
		Price     interface{} `json:"price"`
		Date      string      `json:"date"`
		Available bool        `json:"available"`
	} `json:"dates"`
}

func main() {
	configs := utils.LoadConfig()
	t := telegram.New(configs.TelegramApiKey)
	URL := generateURL(configs)
	tickets := getTickets(URL)
	sendToTelegram(tickets, t, configs)
}

func generateURL(configs utils.Config) string {
	return "https://" + configs.Country + ".megabus.com/journey-planner/api/journeys/prices" +
		"?originId=" + configs.OriginId + "&destinationId=" + configs.DestinationId +
		"&departureDate=" + configs.DepartureDate + "&minDate=" + configs.MinDate +
		"&days=" + configs.Days +
		"&totalPassengers=1&concessionCount=0&nusCount=0&otherDisabilityCount=0&wheelchairSeated=0&pcaCount=0"
}

func getTickets(URL string) Tickets {
	request := gorequest.New().Timeout(30*time.Second).Retry(3, 5*time.Second, http.StatusInternalServerError)
	_, body, _ := request.Get(URL).End()
	var tickets Tickets
	err := json.Unmarshal([]byte(body), &tickets)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return tickets
}

func sendToTelegram(tickets Tickets, t telegram.Client, configs utils.Config) {
	for _, ticket := range tickets.Dates {
		if ticket.Price != nil {
			err := telegram.SendMessage(t, telegram.Message{
				ChatId: configs.TelegramChannelId,
				Text:   fmt.Sprintf("Date: %v\nPrice: %v", ticket.Date, ticket.Price),
			})
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
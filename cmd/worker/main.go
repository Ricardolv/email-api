package main

import (
	"email-api/internal/domain/campaign"
	"email-api/internal/infrastructure/mail"
	"email-api/internal/infrastructure/persistence"
	"fmt"
	"log"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("started worker")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := persistence.NewDb()
	repository := persistence.CampaignRepository{Db: db}

	campaignService := campaign.ServiceImp{
		Repository: &persistence.CampaignRepository{Db: db},
		SendMail:   mail.SendMail,
	}

	campaigns, _ := repository.GetCampaignsToBeSent()

	for {
		for _, campaign := range campaigns {
			campaignService.SendEmailAndUpdateStatus(&campaign)
		}

		time.Sleep(10 * time.Second)
	}

}

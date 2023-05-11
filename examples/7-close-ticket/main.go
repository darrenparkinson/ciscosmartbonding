package main

import (
	"context"
	"log"
	"os"

	"github.com/darrenparkinson/ciscosmartbonding"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	var u, p string
	var ok bool
	if u, ok = os.LookupEnv("SMART_BONDING_CLIENTID"); !ok {
		log.Fatal("missing client id")
	}
	if p, ok = os.LookupEnv("SMART_BONDING_SECRET"); !ok {
		log.Fatal("missing secret")
	}

	c := ciscosmartbonding.NewClient(u, p, nil)
	c.RestyClient.Debug = true

	// You can use the PushUpdate function...

	// Or use the helper function:
	ctr := ciscosmartbonding.CloseTicketRequest{
		ProblemDescription: "Updated description of the ticket",
		CustomerSymptom:    "Faulty power socket",
		HardwareProductID:  "495266",
		SoftwareProductID:  "888410",
		ResolutionSummary:  "faulty power socket changed",
		Complexity:         ciscosmartbonding.Level0Procedural,
		ResolutionCode:     ciscosmartbonding.ResolutionCode_CustomerEducation,
		UnderlyingCause:    ciscosmartbonding.UnderlyingCause_Licensing,
	}

	res, err := c.CloseTicket(context.Background(), "PartnerTicket15", ctr)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res.StatusCode())

}

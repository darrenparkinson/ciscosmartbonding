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
	mustLoadEnv(&u, &p)

	c := ciscosmartbonding.NewClient(u, p, nil)
	c.RestyClient.Debug = true

	// You can use the PushUpdate function... (change these details as required)
	res, err := c.PushUpdate(context.Background(), &ciscosmartbonding.CallData{
		Calls: &ciscosmartbonding.InboundCallsHolder{
			CustCallID: ciscosmartbonding.String("PartnerTicket121"),
			Remarks:    ciscosmartbonding.String("Ticket Update - Escalation"),
		},
		CallStates: &ciscosmartbonding.CallSystemCodesHolder{
			ShortName: ciscosmartbonding.String("Update"),
		},
		Priorities: &ciscosmartbonding.CallSystemCodesHolder{
			ShortName: ciscosmartbonding.String("Escalated"),
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res.StatusCode())
}

func mustLoadEnv(u, p *string) {
	var ok bool
	if *u, ok = os.LookupEnv("SMART_BONDING_USERNAME"); !ok {
		log.Fatal("missing SMART_BONDING_USERNAME variable")
	}
	if *p, ok = os.LookupEnv("SMART_BONDING_PASSWORD"); !ok {
		log.Fatal("missing SMART_BONDING_PASSWORD variable")
	}
}

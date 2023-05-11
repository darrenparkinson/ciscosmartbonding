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
	var clientID, secret string
	mustLoadEnv(&clientID, &secret)

	c := ciscosmartbonding.NewClient(clientID, secret, nil)
	c.RestyClient.Debug = true

	// You can use the PushUpdate function... (change these details as required)
	res, err := c.PushUpdate(context.Background(), &ciscosmartbonding.CallData{
		Calls: &ciscosmartbonding.InboundCallsHolder{
			CustCallID: ciscosmartbonding.String("PartnerTicket16"),
			// Remarks:    ciscosmartbonding.String("Ticket Update - Escalation"),
			Remarks: &ciscosmartbonding.StringOrSliceOfErrors{
				RemarkString: "Ticket Update - Escalation",
			},
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

func mustLoadEnv(clientID, secret *string) {
	var ok bool
	if *clientID, ok = os.LookupEnv("SMART_BONDING_CLIENTID"); !ok {
		log.Fatal("missing SMART_BONDING_CLIENTID variable")
	}
	if *secret, ok = os.LookupEnv("SMART_BONDING_SECRET"); !ok {
		log.Fatal("missing SMART_BONDING_SECRET variable")
	}
}

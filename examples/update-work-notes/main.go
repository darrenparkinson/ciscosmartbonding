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
	if u, ok = os.LookupEnv("SMART_BONDING_USERNAME"); !ok {
		log.Fatal("missing username")
	}
	if p, ok = os.LookupEnv("SMART_BONDING_PASSWORD"); !ok {
		log.Fatal("missing password")
	}

	c := ciscosmartbonding.NewClient(u, p, nil)
	c.RestyClient.Debug = true

	// You can use the PushUpdate function...
	// res, err := c.PushUpdate(context.Background(), &ciscosmartbonding.CallData{
	// 	Calls: &ciscosmartbonding.InboundCallsHolder{
	// 		CustCallID: ciscosmartbonding.String("PartnerTicket001"),
	// 		Remarks:    ciscosmartbonding.String("some ticket updates. All notes / text updates go in here"),
	// 	},
	// 	CallStates: &ciscosmartbonding.CallSystemCodesHolder{
	// 		ShortName: ciscosmartbonding.String("Update"),
	// 	},
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println(res.StatusCode())

	// Or use the helper function:
	res, err := c.UpdateTicketWithWorkNotes(context.Background(), "PartnerTicket001", "some ticket updates. All notes / text updates go in here")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res.StatusCode())

}

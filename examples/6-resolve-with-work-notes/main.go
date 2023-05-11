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
	// res, err := c.PushUpdate(context.Background(), &ciscosmartbonding.CallData{
	// 	Calls: &ciscosmartbonding.InboundCallsHolder{
	// 		CustCallID: ciscosmartbonding.String(ticketID),
	// 		Remarks: &ciscosmartbonding.StringOrSliceOfErrors{RemarkString: remarks},
	// 	},
	// 	CallStates: &ciscosmartbonding.CallSystemCodesHolder{
	// 		ShortName: ciscosmartbonding.String("Resolved"),
	// 	},
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println(res.StatusCode())

	// Or use the helper function:
	res, err := c.ResolveTicketWithWorkNotes(context.Background(), "PartnerTicket14", "some ticket updates. All notes / text updates go in here - update resolve")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res.StatusCode())

}

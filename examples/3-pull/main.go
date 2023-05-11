package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"strings"

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
	for {
		res, resp, err := c.PullUpdate(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		// they don't send us this any more, but just in case...
		if resp.StatusCode() == http.StatusNoContent {
			log.Println("no content")
			break
		}
		// they may just send us a message and status...
		if res.Status == "200" && strings.ToLower(res.Message) == "no messages available to send" {
			log.Println("no messages to retrieve")
			break
		}
		// log.Printf("%+v\n", res)
		log.Println("Service Grid Call ID:", res.GetCalls().GetSDCallID())
		log.Println("    Customer Call ID:", res.GetCalls().GetCustCallID())
		log.Println("  Cisco TAC CSOne ID:", res.GetCalls().GetSPCallID())
		log.Println("          Call State:", res.GetCallStates().GetShortName())
		// Remarks now contains either a string or a slice of errors!
		remarks := res.GetCalls().GetRemarks()
		log.Println("             Remarks:", remarks.RemarkString)
		if remarks.RemarkString == "" && len(remarks.Errors) > 0 {
			for _, e := range res.GetCalls().GetRemarks().Errors {
				log.Println("               Error:", e.ErrorCode, e.ErrorMessage)
			}
		}
	}

}

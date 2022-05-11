package main

import (
	"context"
	"log"
	"net/http"
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
	// c.RestyClient.Debug = true
	for {
		res, resp, err := c.PullUpdate(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("%+v\n", res)
		if resp.StatusCode() == http.StatusNoContent {
			break
		}
	}

}

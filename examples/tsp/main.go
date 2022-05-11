package main

import (
	"context"
	"fmt"
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
	// c.RestyClient.Debug = true

	res, err := c.GetTSPCodes(context.Background(), &ciscosmartbonding.ListParams{Limit: ciscosmartbonding.Int64(5)})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(len(res.TspCodes))
	for _, code := range res.TspCodes {
		fmt.Printf("%s\t%s\t%s\n", code.GetTechName(), code.GetSubTechName(), code.GetProblemCodeDescription())
	}

	res2, err := c.GetAllTSPCodes(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	log.Println(len(res2))
}

package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
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

	savejson := flag.Bool("save", false, "save to tsp-codes.json")
	flag.Parse()

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

	if *savejson {
		saveJSON("tsp-codes.json", res2)
	}
}

func saveJSON(filename string, data interface{}) error {
	file, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, file, 0644)
}

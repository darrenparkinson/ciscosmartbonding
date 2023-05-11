package main

import (
	"context"
	"encoding/csv"
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
	if u, ok = os.LookupEnv("SMART_BONDING_CLIENTID"); !ok {
		log.Fatal("missing client id")
	}
	if p, ok = os.LookupEnv("SMART_BONDING_SECRET"); !ok {
		log.Fatal("missing secret")
	}

	savejson := flag.Bool("save", false, "save to tsp-codes.json")
	savecsv := flag.Bool("csv", false, "save to tsp-codes.csv")
	flag.Parse()

	c := ciscosmartbonding.NewClient(u, p, nil)
	// c.RestyClient.Debug = true

	res2, err := c.GetAllTSPCodes(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Retrieved", len(res2), "codes")

	if *savejson {
		fmt.Println("Saving to tsp-codes.json...")
		saveJSON("tsp-codes.json", res2)
	}

	if *savecsv {
		fmt.Println("Saving to tsp-codes.csv...")
		saveCSV("tsp-codes.csv", res2)
	}
}

func saveJSON(filename string, data interface{}) error {
	file, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, file, 0644)
}

func saveCSV(filename string, codes []ciscosmartbonding.TspCode) error {
	data := [][]string{}
	data = append(data, []string{"Technology", "SubTechnology", "ProblemCode"})
	for _, code := range codes {
		data = append(data, []string{code.GetTechName(), code.GetSubTechName(), code.GetProblemCodeDescription()})
	}
	return csvExport(filename, data)
}

func csvExport(filename string, data [][]string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, value := range data {
		if err := writer.Write(value); err != nil {
			return err
		}
	}
	return nil
}

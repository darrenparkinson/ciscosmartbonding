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
	var u, p, cv, cev, first, last, ccoid, tel, email, productID, serialNo, contractID, siteID string
	mustLoadEnv(&u, &p, &cv, &cev, &first, &last, &ccoid, &tel, &email, &productID, &serialNo, &contractID, &siteID)

	c := ciscosmartbonding.NewClient(u, p, nil)
	c.RestyClient.Debug = true

	// You can use the PushUpdate function... (change these details as required)
	res, err := c.PushUpdate(context.Background(), &ciscosmartbonding.CallData{
		Calls: &ciscosmartbonding.InboundCallsHolder{
			CustCallID:              ciscosmartbonding.String("PartnerTicket123"),
			ShortDescription:        ciscosmartbonding.String("Title of Ticket"),
			Description:             ciscosmartbonding.String("Test Ticket - long description of the issue"),
			CustomerReasonCategory1: ciscosmartbonding.String("Voice - Communications Manager"),
			CustomerReasonCategory2: ciscosmartbonding.String("Business Edition 6000 series / 7000 series with UCM"),
			CustomerReasonCategory3: ciscosmartbonding.String("INSTLL_UNSTLL_UPGRD"),
			Caller: &ciscosmartbonding.PersonsHolder{
				LastName:  ciscosmartbonding.String("Sampler"),
				FirstName: ciscosmartbonding.String("William"),
				Title:     ciscosmartbonding.String("ACME Inc."),
				EMail:     ciscosmartbonding.String("wsampler@company.com"),
			},
			CHD: &ciscosmartbonding.PersonsHolder{
				LastName:  ciscosmartbonding.String(last),
				FirstName: ciscosmartbonding.String(first),
				PIN:       ciscosmartbonding.String(ccoid),
				Sign:      ciscosmartbonding.String("B2B"),
				Tel:       ciscosmartbonding.String(tel),
				EMail:     ciscosmartbonding.String(email),
			},
			MainComp: &ciscosmartbonding.ComponentsHolder{
				Room:      ciscosmartbonding.String(productID),
				SerNrProv: ciscosmartbonding.String(serialNo),
				InvNr:     ciscosmartbonding.String(contractID),
				Location:  ciscosmartbonding.String(siteID),
			},
			SubComp: &ciscosmartbonding.ComponentsHolder{
				LocationName:     ciscosmartbonding.String("Suite 112/Gate 09"),
				LocationCategory: ciscosmartbonding.String("Data Center"),
				LocationZip:      ciscosmartbonding.String("12345"),
				LocationCity:     ciscosmartbonding.String("Capital City"),
				LocationStreet:   ciscosmartbonding.String("12345 Main Street"),
				LocationProvince: ciscosmartbonding.String(""),
			},
		},
		CallStates: &ciscosmartbonding.CallSystemCodesHolder{
			ShortName: ciscosmartbonding.String("New"),
		},
		Contracts: &ciscosmartbonding.InboundContractsHolder{
			ShortName: ciscosmartbonding.String(cv),
		},
		ContractElements: &ciscosmartbonding.InboundContractElementsHolder{
			ShortName: ciscosmartbonding.String(cev),
		},
		Priorities: &ciscosmartbonding.CallSystemCodesHolder{
			ShortName: ciscosmartbonding.String("Escalated"),
		},
		Severities: &ciscosmartbonding.CallSystemCodesHolder{
			ShortName: ciscosmartbonding.String("1"),
		},
		ExtTableValues: &ciscosmartbonding.CallExtensionsHolder{
			Field11:  ciscosmartbonding.String("Jane Doe, Reception"),
			Field14:  ciscosmartbonding.String("2022-04-01 10:20:00"),
			Field21:  ciscosmartbonding.String("None"),
			Field25:  ciscosmartbonding.String("United States"),
			Field26:  ciscosmartbonding.String("1"),
			Field31:  ciscosmartbonding.String("2022-04-01 10:20:00"),
			Field104: ciscosmartbonding.String("2022-04-01 10:20:00"),
			Field108: ciscosmartbonding.String("VendINC123"),
			Field109: ciscosmartbonding.String("Vendor"),
			Field111: ciscosmartbonding.String("Router Cat9K"),
			Field112: ciscosmartbonding.String("Power"),
			Field113: ciscosmartbonding.String("Simon"),
			Field114: ciscosmartbonding.String("Smith"),
			Field115: ciscosmartbonding.String("1-855-000-0000"),
			Field116: ciscosmartbonding.String("simon.smith@vendor.com"),
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res.StatusCode())
}

func mustLoadEnv(u, p, cv, cev, first, last, ccoid, tel, email, productID, serialNo, contractID, siteID *string) {
	var ok bool
	if *u, ok = os.LookupEnv("SMART_BONDING_USERNAME"); !ok {
		log.Fatal("missing SMART_BONDING_USERNAME variable")
	}
	if *p, ok = os.LookupEnv("SMART_BONDING_PASSWORD"); !ok {
		log.Fatal("missing SMART_BONDING_PASSWORD variable")
	}
	if *cv, ok = os.LookupEnv("SMART_BONDING_CONTRACT_VALUE"); !ok {
		log.Fatal("missing SMART_BONDING_CONTRACT_VALUE variable")
	}
	if *cev, ok = os.LookupEnv("SMART_BONDING_CONTRACT_ELEMENT_VALUE"); !ok {
		log.Fatal("missing SMART_BONDING_CONTRACT_ELEMENT_VALUE variable")
	}
	if *first, ok = os.LookupEnv("SMART_BONDING_FIRSTNAME"); !ok {
		log.Fatal("missing SMART_BONDING_FIRSTNAME variable")
	}
	if *last, ok = os.LookupEnv("SMART_BONDING_LASTNAME"); !ok {
		log.Fatal("missing SMART_BONDING_LASTNAME variable")
	}
	if *ccoid, ok = os.LookupEnv("SMART_BONDING_CCOID"); !ok {
		log.Fatal("missing SMART_BONDING_CCOID variable")
	}
	if *tel, ok = os.LookupEnv("SMART_BONDING_TELEPHONE"); !ok {
		log.Fatal("missing SMART_BONDING_TELEPHONE variable")
	}
	if *email, ok = os.LookupEnv("SMART_BONDING_EMAIL"); !ok {
		log.Fatal("missing SMART_BONDING_EMAIL variable")
	}

	if *productID, ok = os.LookupEnv("SMART_BONDING_PRODUCTID"); !ok {
		log.Fatal("missing SMART_BONDING_PRODUCTID variable")
	}
	if *serialNo, ok = os.LookupEnv("SMART_BONDING_SERIALNO"); !ok {
		log.Fatal("missing SMART_BONDING_SERIALNO variable")
	}
	if *contractID, ok = os.LookupEnv("SMART_BONDING_CONTRACTID"); !ok {
		log.Fatal("missing SMART_BONDING_CONTRACTID variable")
	}
	if *siteID, ok = os.LookupEnv("SMART_BONDING_INSTALLSITEID"); !ok {
		log.Fatal("missing SMART_BONDING_INSTALLSITEID variable")
	}

}

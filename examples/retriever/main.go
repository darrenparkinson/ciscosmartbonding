package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/darrenparkinson/ciscosmartbonding"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
)

const servicename = "cisco-smartbonding-retriever" // used by tracing

var (
	env = flag.String("env", "production", "Environment [ production | development ]")
)

type retrieverService struct {
	username string // Cisco Smart Bonding Username
	password string // Cisco Smart Bonding Password
	// contractValue        string // Cisco Smart Bonding Contract Value
	// contractElementValue string // Cisco Smart Bonding Contract Element Value

	logger *zerolog.Logger

	sbc *ciscosmartbonding.Client // Cisco Smart Bonding Client for making requests.
}

func main() {
	flag.Parse()
	godotenv.Load()
	svc := retrieverService{}
	mustMapEnv(&svc.username, "SMART_BONDING_USERNAME")
	mustMapEnv(&svc.password, "SMART_BONDING_PASSWORD")
	// mustMapEnv(&svc.contractValue, "SMART_BONDING_CONTRACT_VALUE")
	// mustMapEnv(&svc.contractElementValue, "SMART_BONDING_CONTRACT_ELEMENT_VALUE")

	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	if *env == "development" {
		logger = logger.Output(zerolog.ConsoleWriter{Out: os.Stderr})
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
	log.SetFlags(0)
	log.SetOutput(logger)
	svc.logger = &logger

	svc.sbc = ciscosmartbonding.NewClient(svc.username, svc.password, nil)

	go func() {

		for {
			time.Sleep(time.Second * 30)
			logger.Info().Msg("making pull update request")
			res, resp, err := svc.sbc.PullUpdate(context.Background())
			if err != nil {
				logger.Error().Err(err).Msg("error pulling update")
				continue
			}
			if resp.StatusCode() == http.StatusNoContent {
				logger.Info().Msg("no content received: sleeping")
				continue
			}
			// TODO: Send this somewhere
			// TODO: Add to WG so we can wait on quit to ensure dealt with
			fmt.Println("****** NEW UPDATE ******")
			fmt.Println("Service Grid Call ID:", fmt.Sprintf("%0.f", res.GetCalls().GetSDCallID()))
			fmt.Println("    CS One Ticket ID:", res.GetCalls().GetSPCallID())
			fmt.Println("    Customer Call ID:", res.GetCalls().GetCustCallID())
			fmt.Println("   Short Description:", res.GetCalls().GetShortDescription())
			fmt.Println("          Call State:", res.GetCallStates().GetShortName())
			fmt.Println("             Remarks:", res.GetCalls().GetRemarks())
			fmt.Println("****** END NEW UPDATE ******")

		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt) // via SIGINT/SIGKILL/SIGQUIT/SIGTERM
	<-c                            // block until signal received
	logger.Info().Msg("received exit signal")
	os.Exit(0)
}

func mustMapEnv(target *string, envKey string) {
	v := os.Getenv(envKey)
	if v == "" {
		panic(fmt.Sprintf("environment variable %q not set", envKey))
	}
	*target = v
}

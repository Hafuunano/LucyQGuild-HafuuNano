package main

import (
	"github.com/FloatTech/floatbox/process"
	nano "github.com/fumiama/NanoBot"
	"github.com/joho/godotenv"
	_ "github.com/moyoez/HafuuNano/plugins/bind"
	_ "github.com/moyoez/HafuuNano/plugins/chun"
	_ "github.com/moyoez/HafuuNano/plugins/fortune"
	_ "github.com/moyoez/HafuuNano/plugins/mai"
	_ "github.com/moyoez/HafuuNano/plugins/ping"
	_ "github.com/moyoez/HafuuNano/plugins/reborn"
	_ "github.com/moyoez/HafuuNano/plugins/sign"
	_ "github.com/moyoez/HafuuNano/plugins/status"
	"log"
	"os"
	//	_ "github.com/moyoez/HafuuNano/plugins/tarap"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	//nano.OpenAPI = nano.SandboxAPI
	nano.Run(process.GlobalInitMutex.Unlock, &nano.Bot{
		AppID:      os.Getenv("appid"),
		Token:      os.Getenv("token"),
		Secret:     os.Getenv("secret"),
		Intents:    nano.IntentAllGuild,
		SuperUsers: []string{os.Getenv("master")},
	})
}

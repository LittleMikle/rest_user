package main

import (
	"github.com/LittleMikle/rest_user/pkg/config"
	"github.com/LittleMikle/rest_user/pkg/routes"
	"github.com/LittleMikle/rest_user/storage"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func main() {
	conf, err := config.CreateConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}

	_, err = storage.NewConnection(conf)
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}

	router := gin.Default()

	routes.UserRoute(router)

	router.Run("localhost:8081")
}

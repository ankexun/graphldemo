package main

import (
	"fmt"
	"graphqldemo/models"
	"graphqldemo/schema"
	"graphqldemo/utils"
	"log"
	"net/http"

	"graphqldemo/service/config"

	"github.com/rs/cors"
)

// 初始化配置
func init() {
	config.Setup()
	models.Setup()
}

// main
func main() {
	mux := http.NewServeMux()

	h := schema.Register()

	// 跨域
	c := cors.New(cors.Options{
		// AllowedOrigins:   []string{"http://localhost:8080"},
		// AllowedHeaders: []string{"*"},
		AllowedOrigins:   config.ServerSetting.AllowedOrigins,
		AllowedHeaders:   config.ServerSetting.AllowedHeaders,
		AllowCredentials: true,
		// Debug:            true,
	})

	mux.Handle("/graphql", c.Handler(h))
	mux.HandleFunc("/login", utils.CreateTokenEndpoint)

	readTimeout := config.ServerSetting.ReadTimeout
	writeTimeout := config.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf("%s:%d", config.ServerSetting.Host, config.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        mux,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	log.Fatal(server.ListenAndServe())
}

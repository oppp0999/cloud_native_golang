package main

import (
	"flag"
	"log"

	"/Users/kimheejae/Desktop/project/cloud_native_golang/3. 마이크로서비스 보안/3. go에서 HTTPS 서버 구축/Chapter03/myevents/src/eventsservice/rest"
	"/Users/kimheejae/Desktop/project/cloud_native_golang/3. 마이크로서비스 보안/3. go에서 HTTPS 서버 구축/Chapter03/myevents/src/lib/configuration"
	"/Users/kimheejae/Desktop/project/cloud_native_golang/3. 마이크로서비스 보안/3. go에서 HTTPS 서버 구축/Chapter03/myevents/src/lib/persistence/dblayer"
)

func main() {

	confPath := flag.String("conf", `.\configuration\config.json`, "flag to set the path to the configuration json file")
	flag.Parse()
	//extract configuration
	config, _ := configuration.ExtractConfiguration(*confPath)

	log.Println("Connecting to database")
	dbhandler, err := dblayer.NewPersistenceLayer(config.Databasetype, config.DBConnection)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Database connection successful... ")
	//RESTful API start
	httpErrChan, httptlsErrChan := rest.ServeAPI(config.RestfulEndpoint, config.RestfulTLSEndPint, dbhandler)
	select {
	case err := <-httpErrChan:
		log.Fatal("HTTP Error: ", err)
	case err := <-httptlsErrChan:
		log.Fatal("HTTPS Error: ", err)
	}
}

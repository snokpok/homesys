package env

import (
	"fmt"
	"homesys/apps/online-serv/constants"
	"log"
	"os"
	"strconv"
)

type Env struct {
	Port                int64
	InfluxDBURL         string
	RaspberryPiIPv4Addr string
}

func GetEnv() Env {
	portpure := os.Getenv("PORT")
	// port undefined -> default
	if len(portpure) == 0 {
		portpure = fmt.Sprint(constants.DEFAULT_APP_PORT)
	}
	// can't be negative
	if portpure[0] == '-' {
		log.Fatalf("Port must not be a negative number")
	}
	// try parsing and default if can't
	port, err := strconv.ParseInt(portpure, 0, 64)
	if err != nil {
		port = int64(constants.DEFAULT_APP_PORT)
	}
	return Env{
		Port:                port,
		InfluxDBURL:         os.Getenv("INFLUXDB_URL"),
		RaspberryPiIPv4Addr: os.Getenv("RASPBERRYPI_IPV4_ADDR"),
	}
}

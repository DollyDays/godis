package main

import (
	"log"
	"os"
)

type CmdType = byte

const (
	COMMAND_UNKONW CmdType = 0x00
	COMMAND_INLINE CmdType = 0x01
	COMMAND_BULK   CmdType = 0x02
)

const (
	GODIS_IO_BUF     int = 1024 * 16
	GODIS_MAX_BULK   int = 1024 * 4
	GODIS_MAX_INLINE int = 1024 * 4
)

type GodisDB struct {
	data   *Dict
	expire *Dict
}

type GodisServer struct {
	fd        int
	port      int
	db        *GodisDB
	clients   map[int]*GodisClient
	eventLoop *EventLoop
}

type GodisClient struct {
}

func initServer(config *Config) error {
	server.port = config.Port

}

func main() {
	path := os.Args[1]
	config, err := LoadConfig(path)
	if err != nil {
		log.Printf("config error : %v\n", err)
	}
	err = initServer(config)
	if err != nil {
		log.Printf("init server error: %v\n", err)
	}
}

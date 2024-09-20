package main

import (
	"fmt"
	"ideas/cmd/api"
	"ideas/db"
	"ideas/utils"
)

func main() {
	if err := MountWebServer(); err != nil {
		panic(fmt.Sprintf("Erro montando o webserver -> %s", err))
	}
}

func MountWebServer() error {
	db, _db := db.MountDB()

	if _db != nil {
		return fmt.Errorf("%s", _db)
	}

	if _ok := db.Migrate(); _ok != nil {
		return fmt.Errorf("%s", _ok)
	}

	port, _port := utils.GetEnv("HTTP_PORT")

	if _port != nil {
		return fmt.Errorf("%s", _port)
	}

	server := api.Serve(fmt.Sprintf(":%s", port), db)

	if err := server.Run(); err != nil {
		return fmt.Errorf("%s", err)
	}

	return nil
}
package main

import (
	"fmt"
	"server/cmd/api"
	"server/db"
	"server/internal/auth"
	"server/utils"
)

func main() {
	if err := MountWebserver(); err != nil {
		panic(fmt.Sprintf("error mounting webserver -> %s", err))
	}
}

func MountWebserver() error {
	db, _db := db.MountDatabase()

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

	secret, _secret := utils.GetEnv("JWT_SECRET")

	if _secret != nil {
		return fmt.Errorf("%s", _secret)
	}

	jwt := auth.DefaultOpts([]byte(secret))

	server := api.Serve(fmt.Sprintf(":%s", port), db, jwt)

	if err := server.Run(); err != nil {
		return fmt.Errorf("%s", err)
	}

	return nil
}

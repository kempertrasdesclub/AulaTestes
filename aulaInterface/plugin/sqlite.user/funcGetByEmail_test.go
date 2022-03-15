package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"test/aulaInterface/constants"
	"test/aulaInterface/dataformat"
)

func ExampleSQLiteUser_GetByEmail() {
	var err error
	var user dataformat.User
	var userAsByte []byte

	var sqlUser = SQLiteUser{}
	err = sqlUser.Connect(constants.KSQLiteConnectionString)
	if err != nil {
		log.Fatalf("sqlUser.Connect().error: %v", err.Error())
	}

	err = sqlUser.Install()
	if err != nil {
		log.Fatalf("sqlUser.Install().error: %v", err.Error())
	}

	var found bool
	found, err = sqlUser.MailExists("helmut.kemper@gmail.com")
	if err != nil {
		log.Fatalf("sqlUser.MailExists().error: %v", err.Error())
	}

	if found == false {
		log.Fatal("sqlUser.MailExists().found: false")
	}

	user, err = sqlUser.GetByEmail("helmut.kemper@gmail.com")
	if err != nil {
		log.Fatalf("sqlUser.GetMainMenu().error: %v", err.Error())
	}

	userAsByte, err = json.Marshal(&user)
	if err != nil {
		log.Fatalf("json.Marshal().error: %v", err.Error())
	}

	fmt.Printf("%s", userAsByte)

	err = os.Remove(constants.KSQLiteConnectionString)
	if err != nil {
		log.Fatalf("os.Remove().error: %v", err.Error())
	}

	// Output:
	// {"id":"24867707-3a21-4368-8058-3d7b1ddb8c06","manuId":"5996b891-9d3c-4038-af37-cb07f5f0f72d","admin":1,"name":"Helmut Kemper","nickname":"Kemper","mail":"helmut.kemper@gmail.com"}
}

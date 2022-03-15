package main

import (
	"fmt"
	"github.com/google/uuid"
)

type UID struct {
}

func (e *UID) Get() (uID string) {
	id := uuid.New()
	return id.String()
}

func main() {
	var uid = UID{}
	fmt.Printf("%v", uid.Get())
}

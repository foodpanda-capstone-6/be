package main

import (
	"fmt"
	"os"
)

func GetServerIngressPort() string {
	return fmt.Sprintf("%s%s", ":", os.Getenv("SERVER_INGRESS_PORT"))
}

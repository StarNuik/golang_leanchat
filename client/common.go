package client

import "log"

func checkClientError(err error, where string) {
	if err != nil {
		log.Fatalf("%s error: %v\n", where, err)
	}
}

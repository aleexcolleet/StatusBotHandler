package main

import (
	"MST_FV/config"
	"context"
	"fmt"
)

/*
MicroService -> Telegram Message Requester

Things to do:

- Be more specific with errors for a better debug
- Use context.Context (to improve cancellations and deadline)
- The code could use goroutines to manage requests simultaneously
- Mock example for tests

*/

func main() {
	cfg, err := config.GetConfig(context.Background())
	if err != nil {
		fmt.Errorf("Error getting config: %v", err)
	}

}

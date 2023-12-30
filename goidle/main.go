// Command goidle runs forever but does nothing. This is useful to keep
// a docker container in running state.
package main

import (
	"time"
)

func main() {
	for {
		time.Sleep(24 * time.Hour)
	}
}

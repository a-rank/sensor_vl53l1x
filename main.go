package main

import (
	"github.com/d2r2/go-i2c"
	"log"
)

func main() {
	i2c_bus, err := i2c.NewI2C(0x29, 0)
	if err != nil {
		log.Fatalf("Failed to connect %s", err)
	}

	defer func() {
		cerr := i2c_bus.Close()
		if err == nil {
			err = cerr
		}
	}()
}

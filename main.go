package main

import (
	"github.com/d2r2/go-i2c"
	"github.com/d2r2/go-vl53l0x"
	"log"
)

func main() {
	i2cBus, err := i2c.NewI2C(0x29, 1)
	if err != nil {
		log.Fatalf("Failed to connect %s", err)
	}
	defer i2cBus.Close()
	
	sensor := vl53l0x.NewVl53l0x()
	err = sensor.Reset(i2cBus)
	if err != nil {
		log.Fatalf("Failed to reset %s", err)
	}
	err = sensor.Init(i2cBus)
	if err != nil {
		log.Fatalf("Failed to init %s", err)
	}
	rng, err := sensor.ReadRangeSingleMillimeters(i2cBus)
	if err != nil {
		log.Fatalf("Failed to read %s", err)
	}
	log.Printf("Measured range = %v mm", rng)
}

package main

import (
	"fmt"

	"github.com/dagger/hello-dagger/internal/shipments"
)

func main() {
	for _, shipment := range shipments.PendingShipments() {
		fmt.Printf("processing %s with %s\n", shipment.ID, shipment.Carrier)
	}
}

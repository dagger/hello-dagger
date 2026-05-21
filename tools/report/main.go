package main

import (
	"fmt"

	"github.com/dagger/hello-dagger/internal/catalog"
)

func main() {
	products := catalog.FeaturedProducts()
	fmt.Printf("%d products, total value $%.2f\n", len(products), float64(catalog.TotalInventoryValue(products))/100)
}

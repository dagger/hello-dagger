package catalog

import "testing"

func TestFeaturedProducts(t *testing.T) {
	products := FeaturedProducts()
	if len(products) == 0 {
		t.Fatal("expected featured products")
	}
	for _, product := range products {
		if product.ID == "" || product.Name == "" || product.Price <= 0 {
			t.Fatalf("invalid product: %#v", product)
		}
	}
}

func TestTotalInventoryValue(t *testing.T) {
	got := TotalInventoryValue([]Product{
		{Price: 2500},
		{Price: 7500},
	})
	if got != 10000 {
		t.Fatalf("expected value 10000, got %d", got)
	}
}

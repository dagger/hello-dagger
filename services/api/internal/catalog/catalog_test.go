package catalog

import "testing"

func TestFeaturedProducts(t *testing.T) {
	products := FeaturedProducts()
	if len(products) != 3 {
		t.Fatalf("expected 3 featured products, got %d", len(products))
	}

	for _, product := range products {
		if product.ID == "" || product.Name == "" || product.Price <= 0 {
			t.Fatalf("invalid product: %#v", product)
		}
	}
}

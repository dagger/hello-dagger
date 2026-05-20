package catalog

type Product struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Price    int    `json:"price"`
}

func FeaturedProducts() []Product {
	return []Product{
		{ID: "sku_001", Name: "Trace Hoodie", Category: "apparel", Price: 6400},
		{ID: "sku_002", Name: "Cloud Runner Mug", Category: "home", Price: 1800},
		{ID: "sku_003", Name: "Pipeline Notebook", Category: "office", Price: 1200},
	}
}

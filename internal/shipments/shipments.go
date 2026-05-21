package shipments

type Shipment struct {
	ID      string
	Carrier string
	Status  string
}

func PendingShipments() []Shipment {
return []Shipment{
		{ID: "ship_1001", Carrier: "ground", Status: "queued"},
		{ID: "ship_1002", Carrier: "air", Status: "label_created"},
}
}

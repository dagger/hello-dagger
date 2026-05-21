package shipments

import "testing"

func TestPendingShipments(t *testing.T) {
	shipments := PendingShipments()
	if len(shipments) != 2 {
		t.Fatalf("expected 2 pending shipments, got %d", len(shipments))
	}
}

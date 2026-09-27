package analytics

import (
	"reflect"
	"testing"
)

func TestProperties(t *testing.T) {
	p := NewProperties()
	if p == nil {
		t.Fatal("expected non-nil properties")
	}

	prod1 := Product{ID: "p1", SKU: "sku1", Name: "name1", Price: 10.5}

	p.SetRevenue(100.5).
		SetCurrency("USD").
		SetValue(50.0).
		SetPath("/home").
		SetReferrer("https://google.com").
		SetTitle("Home Page").
		SetURL("https://example.com/home").
		SetName("login_event").
		SetCategory("auth").
		SetSKU("sku-123").
		SetPrice(19.99).
		SetProductId("prod-456").
		SetOrderId("order-789").
		SetTotal(200.0).
		SetSubtotal(180.0).
		SetShipping(10.0).
		SetTax(10.0).
		SetDiscount(5.0).
		SetCoupon("SAVE5").
		SetProducts(prod1).
		SetRepeat(true)

	tests := []struct {
		key  string
		want any
	}{
		{"revenue", 100.5},
		{"currency", "USD"},
		{"value", 50.0},
		{"path", "/home"},
		{"referrer", "https://google.com"},
		{"title", "Home Page"},
		{"url", "https://example.com/home"},
		{"name", "login_event"},
		{"category", "auth"},
		{"sku", "sku-123"},
		{"price", 19.99},
		{"id", "prod-456"},
		{"orderId", "order-789"},
		{"total", 200.0},
		{"subtotal", 180.0},
		{"shipping", 10.0},
		{"tax", 10.0},
		{"discount", 5.0},
		{"coupon", "SAVE5"},
		{"products", []Product{prod1}},
		{"repeat", true},
	}

	for _, tc := range tests {
		got, ok := p[tc.key]
		if !ok {
			t.Errorf("key %q not found", tc.key)
		} else if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("key %q = %v, want %v", tc.key, got, tc.want)
		}
	}
}

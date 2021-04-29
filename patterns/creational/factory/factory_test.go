package creational

import (
	"strings"
	"testing"
)

func TestCreatePaymentMethodCash(t *testing.T) {
	payment, err := GetPaymentMethod(Cash)
	if err != nil {
		t.Fatal("el metodo de pago Cash debe existir")
	}

	msg := payment.Pay(10.30)
	if !strings.Contains(msg, "pagado con efectivo") {
		t.Error("El metodo de pago con efectivo no es correcto")
	}
	t.Log("LOG:", msg)
}

func TestGetPaymentMethodDebitCard(t *testing.T) {
	payment, err := GetPaymentMethod(DebitCard)
	if err != nil {
		t.Fatal("el metodo de pago Debito debe existir")
	}

	msg := payment.Pay(22.30)
	if !strings.Contains(msg, "pagado con debito") {
		t.Error("El metodo de pago con debito no es correcto")
	}
	t.Log("LOG:", msg)
}

func TestGetPaymentMethodNonExistent(t *testing.T) {
	_, err := GetPaymentMethod(20)
	if err == nil {
		t.Error("el metodo de pago con id 20 debe devolver un error")
	}
	t.Log("LOG:", err)
}

package main

import "fmt"

type PaymentMethod interface {
	pay(amount float64)
}

type Bkash struct {
	apikey string
}
type Naggad struct {
	apikey string
}
type PaymentService struct {
	method PaymentMethod
}

func NewPaymentService (method PaymentMethod)PaymentService{
  return PaymentService{
	method: method,
  }
}

func (bk Bkash) pay(amount float64) {
	fmt.Printf("paying %.2f tk with bkash\n", amount)
}

func (ng Naggad) pay(amount float64) {
	fmt.Printf("paying %.2f tk with naggad", amount)
}

func (ps PaymentService) checkout() {

	ps.method.pay(100.00)
}

func main() {
	bkash := Bkash{apikey: "343453"}
	naggad :=Naggad{apikey: "34345343"}
	paymentService := NewPaymentService(bkash)
	paymentService1 := NewPaymentService(naggad)

	paymentService.checkout()
	paymentService1.checkout()

}

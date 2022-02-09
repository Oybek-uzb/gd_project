package main

func main() {
	product := "vehicle"
	payWay := 3

	var payment Payment
	switch payWay {
	case 1:
		payment = NewCardPayment("23123123234", "any cvv")
	case 2:
		payment = NewPayPalPayment("351873632", "active")
	case 3:
		payment = NewQIWIPayment("436245236734535", "inactive")
	}

	processOrder(product, payment)
}

func processOrder(product string, payment Payment) {
	// implementation
	err := payment.Pay()
	if err != nil {
		return
	}
}

// Payment ...
type Payment interface {
	Pay() error
}

type cardPayment struct {
	cardNumber, cvv string
}

func NewCardPayment(cardNumber, cvv string) Payment {
	return &cardPayment{
		cardNumber: cardNumber,
		cvv:        cvv,
	}
}

func (p *cardPayment) Pay() error {
	// implementation
	return nil
}

type payPalPayment struct {
	payPalId, status string
}

func NewPayPalPayment(payPalId, status string) Payment {
	return &payPalPayment{
		payPalId: payPalId,
		status:   status,
	}
}

func (p *payPalPayment) Pay() error {
	// implementation
	return nil
}

type qiwiPayment struct {
	qiwiId, status string
}

func NewQIWIPayment(qiwiId, status string) Payment {
	return &qiwiPayment{
		qiwiId: qiwiId,
		status: status,
	}
}

func (p *qiwiPayment) Pay() error {
	// implementation
	return nil
}

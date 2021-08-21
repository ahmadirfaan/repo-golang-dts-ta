package main

import (
	"fmt"
)

//interface
type Payment interface {
	Pay()
}

type Wallect struct {
	PhoneNumber string
}

func (w Wallect) Pay() {
	fmt.Println("Pembayaran menggunakan Wallet dengan no", w.PhoneNumber)
}

type BankTransfer struct {
	AccountNumber string
}

func (b BankTransfer) Pay() {
	fmt.Println("Pembayaran menggunakan Bank Transfer dengan no", b.AccountNumber)
}

type CreditCard struct {
	CardNumber string
}

func (c CreditCard) Pay() {
	fmt.Println("Pembayaran dengan kartu kredit dengan no", c.CardNumber)
}

//Factory
func NewPayment(paymentType string, identifier string) (Payment, error) {
	switch paymentType{
	case "wallet":
		//lakukan validasi identifier misal validatePhoneNumber()
		return Wallect{PhoneNumber: identifier}, nil
	case "bank_transfer":
		//lakukan validasi identifier misal validateAccountNumber()
		return BankTransfer{AccountNumber: identifier}, nil
	case "cc":
		//lakukan validasi identifier misal validateCreditCard()
		return CreditCard{CardNumber: identifier}, nil
	}
	return nil, fmt.Errorf("Payment %s not supporterd", paymentType)
}

func main() {
	wallet, err := NewPayment("wallet", "081216839593")
	if err != nil {
		fmt.Println(err.Error())
	}
	bankTransfer, err := NewPayment("bank_transfer", "12345678")
	if err != nil {
		fmt.Println(err.Error())
	}
	cc, err := NewPayment("cc", "3115-1612-5156-1235")
	if err != nil {
		fmt.Println(err.Error())
	}
	_, err = NewPayment("ngutang", "ngutang dulu bro")
	if err != nil {
		fmt.Println(err.Error())
	}

	wallet.Pay()
	bankTransfer.Pay()
	cc.Pay()
}
// Patron de diseño Factory Method
package main

import "fmt"

type PayMethod interface {
	Pay()
}

type Paypal struct{}

func (p *Paypal) Pay() {
	fmt.Println("Pagado con Paypal")
}

type Cash struct{}

func (c *Cash) Pay() {
	fmt.Println("Pagado con Efectivo")
}

type CreditCard struct{}

func (c *CreditCard) Pay() {
	fmt.Println("Pagado con Tarjeta de crédito")
}

// Se encarga de crear la instancia de la clase que se necesite
// Se puede devolver una interfaz
func Factory(method uint) PayMethod {
	switch method {
	case 1:
		return &Paypal{}
	case 2:
		return &Cash{}
	case 3:
		return &CreditCard{}
	default:
		return nil
	}
}

func main() {
	var method uint

	fmt.Println("Seleccione un método de pago")
	fmt.Println("1: Paypal")
	fmt.Println("2: Efectivo")
	fmt.Println("3: Tarjeta de crédito")
	_, err := fmt.Scanln(&method)
	if err != nil || method > 3 {
		panic("Debe seleccionar un método valido")
	}

	payMethod := Factory(method)
	payMethod.Pay()
}

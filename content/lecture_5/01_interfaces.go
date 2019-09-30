package lecture_5

import (
	"fmt"
	"time"
)

// имитация платежной системы

type Wallet struct {
	funds int
}

type CreditCard struct {
	funds      int
	owner      string
	expireTime time.Time
	bonuses    []Bonus
}

type Bonus struct {
	name        string
	description string
}

type Payer interface {
	//if not enough funds - return error
	Pay(int) error
}

// все интерфейсы должны заканчиваться на -er | очень редко -able
type Funder interface {
	GetFunds() int
}

type PayFunder interface {
	Payer
	Funder
}

//todo : выяснить почему после return происходит вычисление остатка
func (w *Wallet) Pay(amount int) error {
	// проверить достаточно ли средств
	// если нет - вернуть ошибку
	if amount > w.funds {
		return fmt.Errorf("недостаточно средств")
	}

	// вычесть средства
	// вернуть nil
	w.funds -= amount
	return nil
}

func (c *CreditCard) Pay(amount int) error {
	if amount > c.funds {
		return fmt.Errorf("недостаточно средств")
	}

	c.funds -= amount
	return nil
}

func Buy(p Payer, amount int) error {
	err := p.Pay(amount)
	if err != nil {
		return err
	}
	return nil
}

func CheckAndBuy(p PayFunder, amount int) error {
	if p.GetFunds() <= 0 {
		fmt.Println("пополните счет")
		return nil
	}

	err := p.Pay(amount)
	if err != nil {
		return err
	}
	return nil
}

func (w *Wallet) GetFunds() int {
	return w.funds
}

func CheckPaymentType(p Payer) interface{} {
	// добавить еще один платежный метод
	// переписать блок if's на switch
	if wallet, ok := p.(*Wallet); ok {
		fmt.Println("ты пользуешься наличкой")
		return wallet.funds
	}

	if card, ok := p.(*CreditCard); ok {
		fmt.Println("ты пользуешься кредитной картой | владелец карты : ", card.owner)
		return card.expireTime
	}
	return nil
}

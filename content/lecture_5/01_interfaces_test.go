package lecture_5

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

// Alt + Insert

//название теста должно начинаться с Test
func TestPayByWallet(t *testing.T) {
	amount := 50
	payment := 30
	w := &Wallet{amount}

	w.Pay(payment)

	//было 50, заплатили 30 = ждем остаток в 20
	assert.Equal(t, w.funds, -1)
}

func TestPayByWalletGivesError(t *testing.T) {
	amount := 50
	payment := 60
	w := &Wallet{amount}

	err := w.Pay(payment)
	assert.Error(t, err)
}

func TestBuyByWallet(t *testing.T) {
	amount := 50
	payment := 60
	w := &Wallet{amount}

	_ = Buy(w, payment)
	//забудьте про fmt.Println()
	//if err != nil {
	//	panic(err)
	//}

	assert.Equal(t, w.funds, amount-payment)
}

func TestBuyByCreditCard(t *testing.T) {
	amount := 100
	payment := 50

	c := &CreditCard{amount, "Zhannur", time.Now(), nil}

	_ = Buy(c, payment)

	assert.Equal(t, c.funds, amount-payment)
}

func TestCheckAndBuy(t *testing.T) {
	amount := 50
	w := &Wallet{amount}

	CheckAndBuy(w, amount)

	assert.Equal(t, w.funds, 0)
}

func TestCheckPaymentType(t *testing.T) {
	amount := 100

	c := &CreditCard{amount, "Zhannur", time.Now(), nil}

	s := CheckPaymentType(c)

	fmt.Println(s)
}

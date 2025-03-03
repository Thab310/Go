package crypto

type Wallet struct {
	balance int
}

func (w Wallet) Deposit(amount int) {

}

func (w Wallet) Balance() float64 {
	return float64(w.balance)
}

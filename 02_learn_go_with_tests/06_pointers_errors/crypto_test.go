package crypto

import (
	"fmt"
	"testing"
)

func TestCrypto(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposit(10)

	got := wallet.Balance()
	want := 10
	fmt.Printf("address of balance in test is %p \n", &wallet.balance)

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

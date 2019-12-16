package pointer

import (
	"testing"
)

func TestWallet(t *testing.T) {
	assertBalance := func(wallet Wallet, want Bitcoin, t *testing.T) {
		got := wallet.Balance()
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}
	t.Run("testing deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		want := Bitcoin(10)

		assertBalance(wallet, want, t)
	})

	t.Run("testing withdraw", func(t *testing.T) {
		Wallet := Wallet{balance: Bitcoin(27)}
		Wallet.Withdraw(Bitcoin(6))
		want := Bitcoin(21)

		got := Wallet.Balance()
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

}

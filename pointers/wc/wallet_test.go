package pointer

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("testing deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		want := Bitcoin(10)

		assertBalance(wallet, want, t)
	})

	t.Run("testing withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(27)}
		err := wallet.Withdraw(Bitcoin(6))
		want := Bitcoin(21)

		assertBalance(wallet, want, t)
		assertNoError(t, err)

	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(27)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(108))
		assertBalance(wallet, startingBalance, t)
		assertError(t, err, ErrInsufficientFunds)

	})

}

func assertNoError(t *testing.T, err error) {
	t.Helper()

	if err != nil {
		t.Fatal("got an error, but didn't want one")
	}
}
func assertError(t *testing.T, got error, want string) {
	t.Helper()

	if got == nil {
		t.Fatal("wanted error but got none")
	}
	if got.Error() != want {
		t.Errorf("Got %q want %q", got, want)
	}
}

func assertBalance(wallet Wallet, want Bitcoin, t *testing.T) {
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

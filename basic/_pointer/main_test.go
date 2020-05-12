package main

import (
	"errors"
	"testing"
)

var InsufficientFundsError = errors.New("cannot withdraw, insufficient funds")

type Wallet struct {
	balance  int
	balance2 int
}

func (w Wallet) Deposit(amount int) {
	w.balance += amount
}

func (w *Wallet) Withdraw(amount int) error {

	if amount > w.balance {
		return InsufficientFundsError
	}

	w.balance -= amount
	return nil
}

func (w Wallet) Balance() int {
	return w.balance
}

func (w *Wallet) DepositPoniter(amount int) {
	// 修改结构体的值时，需要注意:
	// 在 Go 中，当调用一个函数或方法时，参数会被复制。
	// 当调用 func (w Wallet) Deposit(amount int) 时，w 是来自我们调用方法的副本。
	// 不需要太过计算机化，当你创建一个值，例如一个 wallet，它就会被存储在内存的某处。你可以用 &myval 找到那块内存的地址。
	w.balance2 += amount
}

func (w Wallet) Balance2WithoutPoniter() int {
	return w.balance2
}

func (w *Wallet) BalancePoniter() int {
	return w.balance2
}

func assertError(t *testing.T, got error, want error) {
	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}

	if got != want {
		t.Errorf("got '%s', want '%s'", got, want)
	}
}

func TestWallet(t *testing.T) {

	// 基于复制的值设置struct成员变量
	wallet := Wallet{}
	wallet.Deposit(10)
	got := wallet.Balance()
	want := 10
	if got == want {
		t.Errorf("got %d want %d", got, want)
	}

	// 基于指针设置struct成员变量
	wallet2 := Wallet{}
	wallet2.DepositPoniter(10)
	got2 := wallet2.BalancePoniter()
	got22 := wallet2.Balance2WithoutPoniter()
	want2 := 10
	if got2 != want2 {
		t.Errorf("got %d want %d", got, want)
	}
	if got2 != got22 {
		t.Errorf("got %d want %d", got, want)
	}

	// 断言异常
	err := wallet.Withdraw(100)
	assertError(t, err, InsufficientFundsError)
}

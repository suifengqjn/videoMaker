package account

var Acc *Account
type Account struct {

}

func NewAccount() *Account  {
	acc := &Account{}
	return acc
}

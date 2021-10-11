package accountService

import "log"

type Account struct{}

func CreateAccount(account Account) {
    log.Println("AccountService: CreateAccount")
}

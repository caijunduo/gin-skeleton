package accountService

import "log"

type Permission struct{}

func CreatePermission(permission Permission) {
    log.Println("AccountService: CreatePermission")
}

package mysqlc

import (
    "github.com/gohouse/gorose/v2"
    "skeleton/pkg/mysqlx"
)

func Default() gorose.IOrm {
    return mysqlx.Connection("default")
}
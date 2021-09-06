package db

import (
    "upper.io/db.v3/lib/sqlbuilder"
    "upper.io/db.v3/mysql"
    "upper.io/db.v3/sqlite"
)

func NewMySQL(builder *sqlbuilder.Database, settings mysql.ConnectionURL) error {
    database, err := mysql.Open(settings)
    if err != nil {
        return err
    }
    *builder = database
    return nil
}

func NewSQLite(builder *sqlbuilder.Database, settings sqlite.ConnectionURL) error {
    database, err := sqlite.Open(settings)
    if err != nil {
        return err
    }
    *builder = database
    return nil
}

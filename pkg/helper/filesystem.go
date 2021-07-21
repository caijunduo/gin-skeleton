package helper

import (
    "log"
    "os"
)

type filesystem struct{}

func (fs *filesystem) IsExists(path string) (os.FileInfo, bool) {
    f, err := os.Stat(path)
    return f, err == nil || os.IsExist(err)
}

func (fs *filesystem) IsDir(path string) (os.FileInfo, bool) {
    f, flag := fs.IsExists(path)
    return f, flag && f.IsDir()
}

func (fs *filesystem) IsFile(path string) (os.FileInfo, bool) {
    f, flag := fs.IsExists(path)
    return f, flag && !f.IsDir()
}

func (fs *filesystem) OpenFile(path string, flag int, perm os.FileMode) (*os.File, error) {
    var (
        f   *os.File
        b   bool
        err error
    )
    _, b = fs.IsFile(path)
    if b {
        f, _ = os.OpenFile(path, flag, perm)
    } else {
        f, err = os.Create(path)
    }
    return f, err
}

func (fs *filesystem) WriteFile(path string, str string) error {
    var (
        f   *os.File
        err error
    )
    f, err = fs.OpenFile(path, os.O_APPEND, os.ModePerm)

    defer func() {
        if err = f.Close(); err != nil {
            log.Println(err)
        }
    }()

    if err != nil {
        return err
    }

    if _, err = f.WriteString(str); err != nil {
        return err
    }

    return nil
}

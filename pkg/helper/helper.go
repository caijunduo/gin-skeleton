package helper

var (
    Crypto      *crypto
    Filesystem  *filesystem
    Transaction *transaction
)

func New() {
    Crypto = &crypto{}
    Filesystem = &filesystem{}
    Transaction = &transaction{}
}

package signaturex

const (
    Md5 = "md5"
    RSA = "rsa"
)

type Option struct {
    AppKey         string
    AppSecret      string
    Expires        int
    PublicKeyPath  string
    PrivateKeyPath string
}

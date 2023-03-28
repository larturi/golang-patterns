package strategy

import "fmt"

type PasswordProtector struct {
	user          string
	passwordName  string
	hashAlgorithm HashAlgorithm
}

type HashAlgorithm interface {
	Hash(p *PasswordProtector)
}

func NewPasswordProtector(
	user string,
	passwordName string,
	hash HashAlgorithm,
) *PasswordProtector {
	return &PasswordProtector{
		user:          user,
		passwordName:  passwordName,
		hashAlgorithm: hash,
	}
}

func (p *PasswordProtector) SetHashAlgorithm(hash HashAlgorithm) {
	p.hashAlgorithm = hash
}

func (p *PasswordProtector) Hash() {
	p.hashAlgorithm.Hash(p)
}

type SHA struct{}

func (SHA) Hash(p *PasswordProtector) {
	fmt.Printf("Hashing using SHA for %s\n", p.passwordName)
}

type MD5 struct{}

func (MD5) Hash(p *PasswordProtector) {
	fmt.Printf("Hashing using MD5 for %s\n", p.passwordName)
}

func MainStrategy(run bool) {

	if !run {
		return
	}

	sha := &SHA{}
	md5 := &MD5{}

	PasswordProtector := NewPasswordProtector("Cande", "Gmail Pass", sha)
	PasswordProtector.Hash()
	PasswordProtector.SetHashAlgorithm(md5)
	PasswordProtector.Hash()

}

package authentication

type JwtService struct {
	Secret string
	Issuer string
}

func (service *JwtService) SetSecret(secret string) {
	service.Secret = secret
}

func (service *JwtService) SetIssuer(issuer string) {
	service.Issuer = issuer
}

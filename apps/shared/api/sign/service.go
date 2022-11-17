package sign

import (
	"crypto/hmac"
	"crypto/sha256"
	"net/http"
)

type Service interface {
	Sign(request *http.Request) []byte
	Verify(request *http.Request, sign []byte) bool
}

type service struct {
	signKey []byte
}

func NewService(signKey string) Service {
	return &service{signKey: []byte(signKey)}
}

func (s *service) Sign(request *http.Request) []byte {
	mac := hmac.New(sha256.New, s.signKey)

	// todo
	data := request.Host + request.URL.Path
	if request.URL.RawQuery != "" {
		data += "?" + request.URL.RawQuery
	}

	return mac.Sum([]byte(data))
}

func (s *service) Verify(request *http.Request, sign []byte) bool {
	expected := s.Sign(request)
	return hmac.Equal(expected, sign)
}

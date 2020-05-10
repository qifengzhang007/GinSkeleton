package HttpClient

import (
	"github.com/qifengzhang007/goz"
)

func CreateClient() *goz.Request {
	return goz.NewClient()
}

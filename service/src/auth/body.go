package auth

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

type Body struct {
	Code uint16
	Data gin.H
}

func Make_Body(code uint16) Body {
	return Body{
		code,
		gin.H{},
	}
}

func (self *Body) Set_data(k string, v interface{}) {
	self.Data[k] = v
}

func (self Body) MarshalJSON() ([]byte, error) {
    body := gin.H {
        "code": self.Code,
    }

    if len(self.Data) > 0 {
        body["data"] = self.Data
    }

    return json.Marshal(body)
}

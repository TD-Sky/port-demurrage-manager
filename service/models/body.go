package models

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
)

type Body struct {
	Code uint16
	Msg  string
	Data gin.H
}

func Make_Body(code uint16) Body {
	return Body{
		Code: code,
		Msg:  "",
		Data: gin.H{},
	}
}

func (self *Body) Set_message(msg string) *Body {
	self.Msg = msg

	return self
}

func (self *Body) Set_data(k string, v interface{}) *Body {
	self.Data[k] = v

	return self
}

func (self Body) MarshalJSON() ([]byte, error) {
	body := gin.H{
		"code": self.Code,
	}

	if self.Msg != "" {
		body["msg"] = self.Msg
	}

	if len(self.Data) > 0 {
		body["data"] = self.Data
	}

	return json.Marshal(body)
}

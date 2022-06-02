package auth

import "github.com/gin-gonic/gin"

type Body struct {
	code uint16
	data gin.H
}

func Make_Body(code uint16) Body {
	return Body{
		code,
		gin.H{},
	}
}

func (self *Body) Set_data(k string, v interface{}) {
	self.data[k] = v
}

func (self *Body) To_json() gin.H {
	res := gin.H{
		"code": self.code,
	}

	if len(self.data) > 0 {
		res["data"] = self.data
	}

	return res
}

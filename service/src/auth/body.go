package auth

type Body struct {
	code uint16
	msg  string
	data map[string]interface{}
}

func Make_Body(code uint16, msg string) Body {
	return Body{
		code,
		msg,
		map[string]interface{}{},
	}
}

func (self *Body) Set_data(k string, v interface{}) {
	self.data[k] = v
}

func (self *Body) To_json() map[string]interface{} {
	res := map[string]interface{}{
		"code": self.code,
		"msg":  self.msg,
	}

	if len(self.data) > 0 {
		res["data"] = self.data
	}

	return res
}

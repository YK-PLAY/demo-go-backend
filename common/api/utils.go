package commonapi

import (
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
)

func MakeResponse(c *gin.Context, status int, message string) {
	m := gin.H{}
	m["status"] = status
	m["message"] = message

	c.JSON(200, m)
}

func MakeResponseWithBody(c *gin.Context, status int, message string, body Response) {
	m := gin.H{}
	m["status"] = status
	m["message"] = message

	if body != nil {
		v := reflect.ValueOf(body)
		t := reflect.TypeOf(body)

		for i := 0; i < v.NumField(); i++ {
			field := t.Field(i)
			jsonTag := field.Tag.Get("json")
			key := strings.TrimSpace(strings.Split(jsonTag, ",")[0])
			value := v.FieldByName(field.Name)
			vi := value.Interface()
			if vi != nil {
				m[key] = vi
			}
		}
	}

	c.JSON(200, m)
}

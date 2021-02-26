package commonapi

import (
	"encoding/json"
	"io/ioutil"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
)

func MakeResponse(c *gin.Context, status int, message string) {
	m := gin.H{}
	writeCommonResField(m, status, message)

	c.JSON(200, m)
}

func MakeResponseWithBody(c *gin.Context, status int, message string, body Response) {
	m := gin.H{}
	writeCommonResField(m, status, message)

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

func ReadRequest(c *gin.Context, req interface{}) error {
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return err
	}

	json.Unmarshal(jsonData, &req)

	return nil
}

func writeCommonResField(m map[string]interface{}, status int, message string) {
	if m != nil {
		m["status"] = status
		if len(message) > 0 {
			m["message"] = message
		}
	}
}

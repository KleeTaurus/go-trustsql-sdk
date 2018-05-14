package trustsql

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

// Lint 封装了请求餐具数据拼接要求
// 1.参数名ASCII码从小到大排序（字典序）；
// 2.如果参数的值为空不参与签名；
// 3.参数名区分大小写；
func Lint(v interface{}, c interface{}) string {
	signMap := make(map[string]string)
	// getCheckString(&signMap, reflect.ValueOf(v))
	getCheckString(&signMap, reflect.ValueOf(c))
	var keys []string
	for k := range signMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	lintString := ""
	first := true
	for k := range keys {
		// check if value is empty
		if signMap[keys[k]] == "" {
			continue
		}
		if keys[k] == "mch_sign" {
			continue
		}

		if !first {
			lintString = lintString + "&" + keys[k] + "=" + signMap[keys[k]]
		} else {
			lintString = keys[k] + "=" + signMap[keys[k]]
		}
		first = false
	}
	log.Printf("lintString is %s", lintString)
	return lintString
}

func getCheckString(m *map[string]string, v reflect.Value) {
	for i := 0; i < v.NumField(); i++ {
		if "sign" == v.Type().Field(i).Tag.Get("json") {
			continue
		}
		tagArr := v.Type().Field(i).Tag.Get("json")
		tag := strings.Split(tagArr, ",")[0]

		// TODO 没有覆盖全类型
		switch v.Field(i).Kind() {
		case reflect.Int64:
			{
				(*m)[tag] = strconv.FormatInt(v.Field(i).Interface().(int64), 10)
				continue
			}
		case reflect.Int:
			{
				(*m)[tag] = strconv.Itoa(v.Field(i).Interface().(int))
				continue
			}
		case reflect.Map:
			{
				mapField := v.Field(i).Interface().(map[string]interface{})
				if mapField == nil {
					continue
				}
				value, err := json.Marshal(mapField)
				if err != nil {
					// TODO
					fmt.Println(err)
				}
				(*m)[tag] = string(value)
				continue
			}
		}

		(*m)[tag] = v.Field(i).Interface().(string)
	}
}

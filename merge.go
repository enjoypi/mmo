package ext

import (
	"reflect"
)

func MergeMapStruct(src map[string]interface{}, dst interface{}) error {
	structValue := reflect.ValueOf(dst).Elem()
	for k, v := range src {
		structFieldValue := structValue.FieldByName(k)

		if !structFieldValue.IsValid() {
			continue
		}

		if !structFieldValue.CanSet() {
			continue
		}

		structFieldType := structFieldValue.Type()
		if structFieldType.Kind() == reflect.Ptr {
			//fmt.Println(structFieldType, structFieldValue)
			//fmt.Println(reflect.Indirect(structFieldValue))
			////val := structFieldValue
			////if val.IsNil()  {
			//val := reflect.New(structFieldType).Elem()
			////}
			//fmt.Println(structFieldValue.IsNil(), val)
			////if err := MergeMapStruct(v.(map[string]interface{}), val); err != nil {
			////	return nil
			////}
			//structFieldValue.Set(val)
		} else {
			val := reflect.ValueOf(v)
			if structFieldType != val.Type() {
				continue
			}

			structFieldValue.Set(val)
		}
	}

	return nil
}

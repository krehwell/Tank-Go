package utils

import "reflect"

// ex:
// a := User { Id: 0, Name: "Yuza", Age: 20 }
// b := UserBody { Name: "Yakuza" }
// result - a = User{ Id: 0, Name: "Yakuza", Age: 20 }
func MergeInPlaceStructWithPartialStruct(obj, partialObj interface{}) {
	objValue := reflect.ValueOf(obj).Elem()
	partialObjValue := reflect.ValueOf(partialObj)

	for i := 0; i < objValue.NumField(); i++ {
		objField := objValue.Field(i)
		if !objField.CanSet() {
			continue
		}

		partialObjField := partialObjValue.FieldByName(objValue.Type().Field(i).Name)
		if partialObjField.IsValid() {
			objField.Set(partialObjField)
		}
	}
}

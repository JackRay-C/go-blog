package transform

import (
	"errors"
	"fmt"
	"reflect"
)

// Transition 将ref的数据填充到target
//  ref {"id": 1, "username": "renhj", "age": 11, "password": "password"}
//  transition after
// 	target {"id": 1, "username": "renhj"}
func Transition(ref interface{}, target interface{}) error {
	tt := reflect.TypeOf(target)  // target 类型
	tv := reflect.ValueOf(target) // target 值
	tk := tt.Kind()               // target 种类

	if tk != reflect.Ptr {
		return errors.New("target must be ptr type! ")
	}
	tfn := tt.Elem().NumField() // target fields 数量



	rt := reflect.TypeOf(ref)  // ref 类型
	rv := reflect.ValueOf(ref) // ref 值
	rk := rt.Kind()            // ref 种类
	rkm := map[string]interface{}{}
	rvm := map[string]interface{}{}

	// 2、遍历ref，放到map里
	if rk == reflect.Ptr {
		rfn := rt.Elem().NumField()
		for i := 0; i < rfn; i++ {
			name := rt.Elem().Field(i).Name
			kind := rv.Elem().Field(i).Kind()
			value := rv.Elem().Field(i)
			rkm[name] = kind
			rvm[name] = value
		}
	} else {
		rfn := rt.NumField()
		for i := 0; i < rfn; i++ {
			name := rt.Field(i).Name
			kind := rv.Field(i).Kind()
			value := rv.Field(i)
			rkm[name] = kind
			rvm[name] = value
		}
	}

	// 3、遍历target，获取map里同名的值
	for i := 0; i < tfn; i++ {
		if tv.Elem().Field(i).Kind() != rkm[tt.Elem().Field(i).Name] {
			// 	判断同名的成员属性类型是否一致
			return errors.New(fmt.Sprintf("[%s:%s] mismatch type [%s:%s]", tt.Name(), tt.Elem().Field(i).Name, rt.Name(), rt.Field(i).Name))
		}
		tv.Elem().Field(i).Set(rvm[tt.Elem().Field(i).Name].(reflect.Value))
	}

	return nil
}

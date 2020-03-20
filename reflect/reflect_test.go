package reflect_test

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTypeAndValue(t *testing.T) {
	var f int64 = 10
	t.Log(reflect.TypeOf(f), reflect.ValueOf(f))
	t.Log(reflect.ValueOf(f).Type())
	/*
		PS C:\Users\vct\Desktop\github\Golang_demos\reflect> go test -v .\reflect_test.go
		=== RUN   TestTypeAndValue
		--- PASS: TestTypeAndValue (0.00s)
		    reflect_test.go:10: int64 10
		    reflect_test.go:11: int64
		PASS
		ok      command-line-arguments  0.204s
	*/
}

func CheckType(v interface{}) {
	t := reflect.TypeOf(v)
	switch t.Kind() {
	case reflect.Float32, reflect.Float64:
		fmt.Println("Float")
	case reflect.Int, reflect.Int32, reflect.Int64:
		fmt.Println("Integer")
	default:
		fmt.Println("Unknown", t)
	}
}

func TestBasicType(t *testing.T) {
	var f float64 = 12
	CheckType(f) // Float
}

type Employee struct {
	EmployeeID string
	Name       string `format:"normal"`
	Age        int
}

func (e *Employee) UpdateAge(newVal int) {
	e.Age = newVal
}

func TestInvokeByName(t *testing.T) {
	e := &Employee{"1", "Mike", 30}
	//按名字获取成员
	//注意TypeOf(*e)，而ValueOf(e)
	t.Logf("Name: value(%[1]v),Type(%[1]T) ", reflect.ValueOf(*e).FieldByName("Name"))
	if nameField, ok := reflect.TypeOf(*e).FieldByName("Name"); !ok {
		t.Error("Failed to get 'Name' field.")
	} else {
		t.Log("Tag:format", nameField.Tag.Get("format"))
	}

	//按名字进行函数调用
	reflect.ValueOf(e).MethodByName("UpdateAge").Call([]reflect.Value{reflect.ValueOf(1)})
	t.Log("Updated Age:", e)

	/*
	   === RUN   TestInvokeByName
	   --- PASS: TestInvokeByName (0.00s)
	       reflect_test.go:59: Name: value(Mike),Type(reflect.Value)
	       reflect_test.go:63: Tag:format normal
	       reflect_test.go:66: Updated Age: &{1 Mike 1}
	   PASS
	*/

}

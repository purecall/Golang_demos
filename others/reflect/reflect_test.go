package reflect_test

import (
	"errors"
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

func TestDeepEqual(t *testing.T) {
	a := map[int]string{1: "one", 2: "two", 3: "three"}
	b := map[int]string{1: "one", 2: "two", 3: "three"}
	// t.Log(a == b) // invalid operation
	t.Log(reflect.DeepEqual(a, b))

	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	s3 := []int{2, 3, 1}
	t.Log("s1==s2?", reflect.DeepEqual(s1, s2))
	t.Log("s1==s3?", reflect.DeepEqual(s1, s3))
	/*
		=== RUN   TestDeepEqual
		--- PASS: TestDeepEqual (0.00s)
		    reflect_test.go:86: true
		    reflect_test.go:91: s1==s2? true
		    reflect_test.go:92: s1==s3? false
		PASS
	*/
}

type Customer struct {
	CookieID string
	Name     string
	Age      int
}

//两个结构体，都有Name和Age字段
//写一个通用的函数，使其能够对不同结构体的相同字段进行赋值
func fillBySettings(st interface{}, settings map[string]interface{}) error {
	// func(v Value) Elem() Value
	// Elem() returns the value that the interface v contains or that the pointer points to
	// It panics if v's Kind is not Interface of Ptr
	// Ite returns the zero Value if v is nil

	if reflect.TypeOf(st).Elem().Kind() != reflect.Ptr {
		// 传入的是指针类型，Elem() 获取指针指向的具体结构
		if reflect.TypeOf(st).Elem().Kind() != reflect.Struct {
			return errors.New("the first param should be a pointer to the struct type.")
		}
	}

	if settings == nil {
		return errors.New("settings is nil.")
	}

	var (
		field reflect.StructField
		ok    bool
	)

	for k, v := range settings {
		//遍历map的key，在结构体中找相同的名字
		if field, ok = (reflect.ValueOf(st)).Elem().Type().FieldByName(k); !ok {
			continue
		}
		if field.Type == reflect.TypeOf(v) {
			vstr := reflect.ValueOf(st)
			vstr = vstr.Elem()
			vstr.FieldByName(k).Set(reflect.ValueOf(v))
		}
	}
	return nil
}

func TestFillNameAndAge(t *testing.T) {
	settings := map[string]interface{}{"Name": "Mike", "Age": 20}
	e := Employee{}
	if err := fillBySettings(&e, settings); err != nil {
		t.Fatal(err)
	}
	t.Log(e)

	c := new(Customer)
	if err := fillBySettings(c, settings); err != nil {
		t.Fatal(err)
	}
	t.Log(*c)
	/*
		=== RUN   TestFillNameAndAge
		--- PASS: TestFillNameAndAge (0.00s)
		    reflect_test.go:147: { Mike 20}
		    reflect_test.go:153: { Mike 20}
		PASS
	*/
}

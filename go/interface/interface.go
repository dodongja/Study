package main

/*type Stringer interface {
	String() string
}

type Student struct {
	Name string
	Age  int
}

func (s Student) String() string {
	return fmt.Sprintf("안녕! 나는 %d살 %s 라고 해", s.Age, s.Name)
}

func main() {
	student := Student{"철수", 12}
	var stringer Stringer

	stringer = student

	fmt.Printf("%s\n", stringer.String())

}*/

/*func PrintVal(v interface{}) {
	switch t := v.(type) {
	case int:
		fmt.Printf("v is int %d\n", int(t))
	case float64:
		fmt.Printf("v is float64 %f\n", float64(t))
	case string:
		fmt.Printf("v is string %s\n", string(t))
	default:
		fmt.Printf("Not supported type: %T:%v\n", t, t)
	}
}

type Student struct {
	Age int
}

func main() {
	PrintVal(10)
	PrintVal(3.14)
	PrintVal("Hello")
	PrintVal(Student{15})
}*/

/*func TotalTime(db Database) int {
	db.Get()
	db.Set()
}

type Database interface {
	Get()
	Set()
}

type BDatabse struct {}
type CDatabse struct {}

func (b BDatabse) Get() {}
func (b BDatabse) Set() {}

func (c CDatabse) Get() {}
func (c CDatabse) Set() {}

func Compare() {
	BDB := &BDatabse{}
	CDB := &CDatabse{}

	if TotalTime(BDB) < TotalTime(CDB) {
		fmt.Println("B회사 제품이 더 빠름")
	} else {
		fmt.Println("C회사 제품이 더 빠름")
	}
	t := 10;
	fmt.Printf("%d", int(t))
}*/

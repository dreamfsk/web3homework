package main

import "fmt"

// 使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，
// 再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。
// 为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息

type Person struct {
	Name string
	Age  int
}

func (p *Person) GetInfo() string {
	return fmt.Sprintf("%s is %d years old", p.Name, p.Age)
}

type Employee struct {
	person     Person
	EmployeeID int
}

func (em Employee) PrintInfo() string {
	return fmt.Sprintf(" ID: %d, %s", em.EmployeeID, em.person.GetInfo())
}

func main() {

	u := Employee{
		person:     Person{Name: "张三", Age: 18},
		EmployeeID: 1,
	}
	fmt.Println(u.PrintInfo())
}

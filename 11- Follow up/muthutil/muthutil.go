// package main
package muthutil
// import (
// 	"errors"
// 	"fmt"
// )


// type Payable interface{
// 	CalculatePay() float64
// }

// type Notifier interface{
// 	Notify(message string)
// }
// type Worker interface{
// 	Work() string 
// 	GetSalary() float64
// }





// type Person struct{
// 	Name string
// 	Age int
// }
// func New(name string, age int)Person{
// 	return Person{
// 		Name: name,
// 		Age: age,
// 	}
// }
// func(p Person) Greet(){
// 	fmt.Printf("Hello, my name is %s\n", p.Name)
// }
// func (p * Person) IsAdult()bool{
// 	return p.Age >= 18
// }



// type Employee struct{
// 	Person
// 	Salary float64
// }
// func (e *Employee) Promote(){
// 	if e.Age > 25{
// 		e.Salary += 1000
// 	}
// }
// func (e *Employee) RaiseSalary(percent float64){
// 	e.Salary = e.Salary + e.Salary * percent
// }
// func (emp  Employee) Work()string{
// 	return fmt.Sprintf("Employee Name: %v, Age: %v, Salary: %v\n", emp.Name, emp.Age, emp.Salary)
// }
// func (emp  Employee) GetSalary()float64{
// 	return emp.Salary
// }
// func (e Employee) Notify(message string){
// 	fmt.Println("Sending email to employee ", e.Name)
// 	fmt.Println(` Dear our valued Employee. `, message)
// }

// func (e Employee) CalculatePay()float64{
// 	return e.Salary
// }

// func (e *Employee) RaiseSalarySafe(percent float64)error{
// 	if percent <= 0{
// 		return errors.New("raise percentage can't be zero or negative")
// 	}
// 	e.RaiseSalary(percent)
// 	return nil

// }




// type Manager struct{
// 	Person
// 	Bonus float64
// }
// func (man Manager) Work()string{
// 		return fmt.Sprintf("Employee Name: %v, Age: %v, Salary: %v\n", man.Name, man.Age, man.Bonus)

// }
// func (man Manager) GetSalary()float64{
// 	return man.Bonus
// }
// func (man Manager) Notify(message string){
// 	fmt.Println("Sending sms message to our valued ", man.Name)
// 	fmt.Print("Dear our valued manager, hope this sms finds you will. ", message)
// }
// func (e Manager) CalculatePay()float64{
// 	return e.Bonus
// }


// func main() {
// 	// Task 1
// 	fmt.Println("Person info.....")
// 	p := New("nsralla", 23)
// 	p.Greet()
// 	fmt.Println("am I adult? ", p.IsAdult())

// 	fmt.Println("Employee info...")
// 	emp := Employee{
// 		Person: Person{
// 			Name:"nsralla",
// 			Age:23,
// 		},
// 		Salary: 1200,
// 	}
// 	fmt.Println("Employee name: ", emp.Name)
// 	fmt.Println("Employee age: ", emp.Age)
// 	fmt.Println("Employee salary: ", emp.Salary)
// 	emp.RaiseSalary(0.2)
// 	fmt.Println("employee salary after raise: ", emp.Salary)
// 	emp.Promote()
// 	fmt.Println("employee salary after promotion: ", emp.Salary)
// 	fmt.Println("----------------------------")

// 	// Task 2
// 	employees := make([]Employee,3)
// 	employees[0] = emp
// 	employees[1] = Employee{
// 		Person:Person{
// 			Name:"Ali",
// 			Age:44,
// 		},
// 		Salary: 2200,
// 	}
// 	employees[2] = Employee{
// 	Person:Person{
// 		Name:"Wael",
// 		Age:41,
// 	},
// 	Salary: 2100,
// 	}

// 	for _, emp := range employees{
// 		PrintWorkerInfo(emp)
// 		emp.RaiseSalary(0.1)
// 		fmt.Println("After update:")
// 		PrintWorkerInfo(emp)
// 		fmt.Println("##########")
// 	}

// 	fmt.Println("Task 3 ------------------------")
// 	workers := make([]Worker, 3)
// 	workers[0] = Employee{
// 		Person:Person{
// 			Name:"Ali",
// 			Age:44,
// 		},
// 		Salary: 2200,
// 	}
// 	workers[1] = Employee{
// 		Person:Person{
// 			Name:"Mohammad",
// 			Age:24,
// 		},
// 		Salary: 1200,
// 	}
// 	workers[2] = Manager{
// 		Person:Person{
// 			Name:"Mansour",
// 			Age:65,
// 		},
// 		Bonus: 8700,
// 	}


// 	for _, worker := range workers{
// 		PrintWorkerInfo(worker)
// 	}

// 	fmt.Println("Task 5-------------------")
// 	notifiers := make([]Notifier, 3)
// 	notifiers[0] = Employee{
// 		Person:Person{
// 			Name:"Ali",
// 			Age:44,
// 		},
// 		Salary: 2200,
// 	}
// 	notifiers[1] = Employee{
// 		Person:Person{
// 			Name:"Mohammad",
// 			Age:24,
// 		},
// 		Salary: 1200,
// 	}
// 	notifiers[2] = Manager{
// 		Person:Person{
// 			Name:"Mansour",
// 			Age:65,
// 		},
// 		Bonus: 8700,
// 	}
// 	SendNotification(notifiers, "None clay lurked for himnot revellers nor light, from lone mine feel crime feere misery name her. Moths had his few drowsy with they steel prose gild, sight way hall.")

// 	fmt.Println("Task 6 .................")
// 	err := emp.RaiseSalarySafe(0.0)
// 	if err != nil{
// 		fmt.Println(err)
// 	}
// 	fmt.Println("My Salary after raise: ", emp.Salary)
// 	fmt.Println("Task 7 ------------------")
// 	fmt.Println("Task 5-------------------")
// 	payables := make([]Payable, 3)
// 	payables[0] = Employee{
// 		Person:Person{
// 			Name:"Ali",
// 			Age:44,
// 		},
// 		Salary: 2200,
// 	}
// 	payables[1] = Employee{
// 		Person:Person{
// 			Name:"Mohammad",
// 			Age:24,
// 		},
// 		Salary: 1200,
// 	}
// 	payables[2] = Manager{
// 		Person:Person{
// 			Name:"Mansour",
// 			Age:65,
// 		},
// 		Bonus: 8700,
// 	}
// 	fmt.Println(TotalPayroll(payables))
// }




// func TotalPayroll(p []Payable)(total float64){
// 	if len(p) == 0{
// 		return 0.0
// 	}
// 	for _,payable := range p{
// 		total += payable.CalculatePay()
// 	}
// 	return
// }





// func SendNotification(notifiers []Notifier, msg string){
// 	for _, notifier := range notifiers{
// 		notifier.Notify(msg)
// 	}
// }



// // PrintWorkerInfo displays information about a Worker by calling its Work method.
// func PrintWorkerInfo(w Worker){
// 	fmt.Println(w.Work())
// }

// // 
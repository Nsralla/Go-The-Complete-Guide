package main

import (
	"fmt"

)

// type Student struct{
// 	Name string
// 	Score int
// }

// func risky(n int){
// 	defer func(){
// 		r := recover()
// 		if r != nil{
// 			fmt.Println("program recovered")
// 		}
// 	}()
// 	if n == 0{
// 		panic("some error happened")
// 	}
// 	fmt.Println(10 / n)
	
// }


type Animal interface{
	Speak() string
	Breed() string
}

type BaseAnimal struct{
	Name string 
	Sound string
}

func (b BaseAnimal) Speak() string{
	return fmt.Sprintf("%s makes %s sound \n", b.Name, b.Sound)
}

type Dog struct{
	BaseAnimal
	breed string
}
func (d Dog) Breed()string{
	return fmt.Sprintf("%s is a %s", d.Name, d.breed)
}

func NewDog(name, sound, breed string) Dog {
	return Dog{
		BaseAnimal: BaseAnimal{Name: name, Sound: sound},
		breed:      breed,
	}
}

func main() {



	
	// risky(0)

	// fmt.Println("continue execution despite panic!!")
	// wordsCount := make(map[string]int)
	// WordFrequency(wordsCount,"Hi all, nsr is my name , nsr is my prefered name")
	// fmt.Println(wordsCount)

	// unique := Unique([]int{9,2,2,1,77,7,7,4,2,9})
	// fmt.Println(unique)

	// ReverseSlice([]int{1,2,3,4,5})

	// m := MergeMaps(map[string]int{
	// 	"a":1,
	// 	"b":2,
	// 	"c":3,
	// }, map[string]int{
	// 	"a":7,
	// 	"b":3,
	// 	"f":1,
	// 	"g":2,
	// })

	// fmt.Println(m)
	// students := make(map[string]Student)

	// for i:=0; i< 3;i++{
	// 	fmt.Println("Enter name then score: ")
	// 	var name string
	// 	var score int
	// 	fmt.Scanln(&name)
	// 	fmt.Scanln(&score)
	// 	students[name] = Student{Name:name, Score:score}
	// }

	// fmt.Println(students)
	// students := []map[string]any{
	// 	{
	// 		"name":"nsr",
	// 		"score":90,
	// 	},
	// 	{
	// 		"name":"wael",
	// 		"score":99,
	// 	},
	// 	{
	// 		"name":"mohsen",
	// 		"score":71,
	// 	},
	// }



	// studentsCopy := make([]map[string]any, len(students))

	// for idx, student := range students{
	// 	stCopy := make(map[string]any)
	// 	stCopy["name"] = student["name"]
	// 	stCopy["score"] = student["score"]
	// 	studentsCopy[idx] = stCopy
	// }

	// studentsCopy[0]["score"] = 100
	// fmt.Println(students)

	// for i := range students{
	// 	for j := range students{
	// 		if students[i]["score"].(int) < students[j]["score"].(int){
	// 			students[i], students[j] = students[j], students[i]
	// 		} 
	// 	} 
	// }
	// sort.Slice(students, func(i, j int) bool{
	// 	return students[i]["score"].(int) > students[j]["score"].(int)
	// })
	// fmt.Println(students)
	// fmt.Println(GroupByLength([]string{"go", "c", "java", "rust", "python","c++","html"}))


	
}

// func WordFrequency(wordsCount map[string]int, text string){
// 	words := strings.Split(text," ")
// 	for _, word := range words{
// 		wordsCount[word] ++
// 	}
// }
// func Unique(nums []int)[]int{
// 	seen := make(map[int]bool)
// 	unqiueNums := []int{}
// 	for _, num := range nums{
// 		if !seen[num]{
// 			seen[num] = true
// 			unqiueNums = append(unqiueNums, num)
// 		}
// 	}
// 	return unqiueNums

// }
// func ReverseSlice(s []int){
// 	for idx := range (len(s) / 2){
// 		s[idx], s[len(s) -1 - idx] = s[len(s) -1 - idx], s[idx]
// 	}
// 	fmt.Println(s)
// }

// func MergeMaps(m1, m2 map[string]int) map[string]int{
// 	merged := make(map[string]int)
// 	for key, value := range m1{
// 		_, ok1 := m1[key]
// 		_, ok2 := m2[key]
// 		if ok1 && ok2{
// 			merged[key] = m1[key] + m2[key]
// 			continue
// 		}
// 		merged[key] = value
// 	}
// 	return merged
// }

// func GroupByLength(words []string) map[int][]string{
// 	langs := make(map[int][]string)
// 	for _, word := range words{
// 		n := len(word)
// 		_, ok := langs[n]
// 		if !ok{
// 			langs[n] = []string{}
// 			langs[n] = append(langs[n], word)
// 			continue
// 		}
// 		langs[n] = append(langs[n], word)
// 	}
// 	return langs
// }

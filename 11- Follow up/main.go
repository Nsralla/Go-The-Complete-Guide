package main

import (
	"errors"
	"fmt"
)

type User struct{
	Name string
	PhoneNumber int
}

func getUserMap(names []string, phoneNumbers []int) (map[string]User,error) {
	if len(names) != len(phoneNumbers){
		return nil, errors.New("names and phone numbers must has the same length") 
	}

	info := make(map[string]User)
	for index, name := range names{
		info[name] = User{
			Name: name,
			PhoneNumber: phoneNumbers[index],
		}
	}
	return info, nil
}

func main() {
	info, err := getUserMap([]string{"nsr","hasan","ahmad"}, []int{1113,03231, 3129})
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println(info)

	deleted, err := deleteIfNecessary(info, "nsr")
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println("was delelted? ",deleted)
	fmt.Println(info)
	// _,ok := info["hasan"]
	// if !ok{
	// 	fmt.Println("hasan doesn't exists")
	// }

	fmt.Println(getCounts([]string{"2","2","4","44","5","5"}))

	fmt.Println(getNameCounts([]string{"nsr","ahmad","ali","nael","nsr","ali"}))
}
//  Maps are passed by reference, so like slices
func deleteIfNecessary(users map[string]User, name string)(deleted bool, err error){
	
	if  _, ok := users[name]; !ok{
		fmt.Println("User with such name doesn't exist")
		return false, errors.New("User doesn't exist")
	}
	delete(users, name)
	return true, nil
}
func getCounts(userIDs []string)map[string]int{
	idFrequency := make(map[string]int)

	for _, id :=range userIDs{
		frequency := idFrequency[id]
		fmt.Println("freq: ",frequency)
		frequency ++
		idFrequency[id] = frequency
	}
	return idFrequency
}


func getNameCounts(names []string)map[rune]map[string]int{
	namesFreqByChar := make(map[rune]map[string]int)
	for _, name := range names{
		firChar := []rune(name)[0]
		if len(namesFreqByChar[firChar]) == 0{
			namesFreqByChar[firChar] = map[string]int{name: 1}
		}else{
		 	namesFreqByChar[firChar][name] ++
		}
		
	}
	return  namesFreqByChar
}
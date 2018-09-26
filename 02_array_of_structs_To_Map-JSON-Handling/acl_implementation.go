package main

import ("fmt"
		"encoding/json"
		)

type FunctionConditions struct {
    FunctionName string
	ConditionsList []ACL
 }

var fnCondArray [] FunctionConditions

type ACL struct {
                 Org string
				 OU string
				 Role string
         }
var aclArray []ACL

var FunctionConditionsMap =  make(map[string][]ACL)

func main() {

accessList := `[{   "FunctionName" : "issue1" , 
					"ConditionsList": [ {"Org":"HILTON","OU":"Bangalore","Role":"ADMIN"}, 
										{"Org":"ACCOR", "OU":"Bangalore","Role":"ADMIN"},
										{"Org":"LUFTHANSA", "OU":"Bangalore","Role":"ADMIN"}] } ,
				{   "FunctionName" : "issue2" , 
					"ConditionsList": [ {"Org":"HILTON", "OU":"Bangalore","Role":"ADMIN"}, 
										{"Org":"ACCOR", "OU":"Bangalore","Role":"ADMIN"}]}
				]`
	
	json.Unmarshal([]byte(accessList), &fnCondArray)

	for _,result := range fnCondArray {
			FunctionConditionsMap[result.FunctionName] = result.ConditionsList
          }
	
	fmt.Println(FunctionConditionsMap["issue1"][0])
	fmt.Println(FunctionConditionsMap["issue1"][0].Org)
	
}
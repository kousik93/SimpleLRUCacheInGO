package main

import (
	"container/list"
	"fmt"
)

func push(l *list.List, val int){
	l.PushFront(val)
}


func pop(l *list.List){
	l.Remove(l.Back())
}

func mostRecent(l *list.List, val int){

	for e:=l.Front(); e != nil; e = e.Next() {
		if e.Value == val{
			l.MoveToFront(e)
			break
		}
	}
}


func set(m map[string]int,l *list.List,key string,size int){

		var val int 
		fmt.Println("Enter value\n")
		fmt.Scan(&val)
		if l.Len()>=size{

			e:=l.Back()
			for key, value := range m {
    			if value==e.Value{
    				delete(m, key)
    				break

    			}
			}

			pop(l)

		}
		m[key]=val
		l.PushFront(val)
		fmt.Println("Key and Value added")
}



func get(m map[string]int,l *list.List,key string,size int){
	
	i:=m[key]
	if i==0 {
		fmt.Println("Key value not found. Will Insert")
		set(m,l,key,size)
	}

	if i!=0{

		fmt.Println("Value is found. It is:")
		fmt.Println(i)
		mostRecent(l,i)

	}
	
}

func printList(l *list.List){
	fmt.Println("Contents of List:")
	//Iterate through list and print its contents.
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

}

func printMap(m map[string]int){
	fmt.Println("Contents of map:")
	for key, value := range m {
    		fmt.Println("Key:", key, "Value:", value)
		}
}

func main() {
	
	l := list.New()
	var m map[string]int
	m = make(map[string]int)
	var size int
	var choise int
	i:=0

	fmt.Println("What is the size of the cache")
	fmt.Scanf("%d",&size)

	for i==0{
		fmt.Println("\n\n")
		fmt.Println("Choose option")
		fmt.Println("1. Get a value")
		fmt.Println("2. Insert a value")
		fmt.Println("3. Print list and map")
		fmt.Println("4. Quit")
		fmt.Println("\n\n")

		fmt.Scanf("%d",&choise)
		if choise==4{
			i=1
		}

		if(choise==1){
			var key string
			fmt.Println("Enter Key")
			fmt.Scanf("%s",&key)
			get(m,l,key,size)
		}
		if(choise==2){
			var key string
			fmt.Println("Enter the Key")
			fmt.Scanf("%s",&key)
			ii:=m[key]
			if ii==0{
				set(m,l,key,size)
			}
			if ii!=0{
				fmt.Println("Error. Key already exists")
			}
		}
		if(choise==3){
			printList(l)
			printMap(m)
		}
	}

			
}
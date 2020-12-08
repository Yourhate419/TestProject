package main

import (
	"fmt"
	"os"
	"io"
	"encoding/json"
	"io/ioutil"
	"encoding/csv"
	"log"
)

type User struct {
	Id string
	Fullname string
	Email string
}

type Data struct {
	Users []User
}

func readCSV(c chan<- *User) {
	file, err := os.Open(os.Getenv("FILENAMECSV"))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	 
	reader := csv.NewReader(file)
	reader.Comma = ';'
 
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		data := User{Id : record[0], Fullname : record[1], Email : record[2]}
		c <- &data 
	}
}

func writeJSON(c <-chan *User) {
	var users Data

    for {
		record := *(<- c)


		records, err := ioutil.ReadFile(os.Getenv("FILENAMEJSON"))
		if err != nil {
		  log.Fatal(err)
		}

		err = json.Unmarshal(records, &users)
		
		if err != nil {
			log.Fatal(err)
		}

		users.Users = append(users.Users, record)

		file, _ := json.MarshalIndent(users, "", " ")
 
		_ = ioutil.WriteFile(os.Getenv("FILENAMEJSON"), file, 0644)
	}
}

func main() {
	var c chan *User = make(chan *User)
	setenv()

    go readCSV(c)
	go writeJSON(c)


	fmt.Println("Нажмите Enter")
	var input string
    fmt.Scanln(&input)
}
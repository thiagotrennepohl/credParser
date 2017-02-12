package main

import (

	"os"
	"log"
	 "crypto/md5"
	"encoding/hex"
	"encoding/json"
	"encoding/csv"
)

type User struct {
	Username string `json:"username"`
	Password Passwords  `json:"passwords"`
}

type Passwords struct {
	Pwd string `json:"password"`
	PwdFormat string `json:"password_format"`
}

func checkErr(err error){
	if err != nil{
		log.Fatal("Oops, it seems somthing went wrong", "error: ", err)
	}
}

func checkArgs() bool{
	if len(os.Args) == 5 {
		return true
	}
	return false
}



func main() {
	if checkArgs() {
		filePath := os.Args[2]
		output := os.Args[4]
		log.Println("loading file", filePath)
		csvFile, err := os.Open(filePath)
		checkErr(err)
		defer csvFile.Close()

		reader := csv.NewReader(csvFile)
		reader.Comma = ' '

		csvData, err := reader.ReadAll()
		checkErr(err)
		log.Println("Parsing json...")
		var users []User
		var user User
		var passwords Passwords
		for _, each := range csvData {
			log.Println(each[0], " ", each[1], " ", each[2])

			username := each[0] + "@" + each[2]
			encryptedPwd := md5Hash(each[1])
			passwords.Pwd = encryptedPwd

			passwords.PwdFormat = "md5"
			user.Username = username
			user.Password = passwords

			users = append(users, user)
		}

		marshalData(users, output)
	} else {
		help()
	}

}

func marshalData(data []User, output string) {
	log.Print("Identing Json and Writing to ", output)
	jsonData, err := json.MarshalIndent(data, "", "    ")
	checkErr(err)
	datafile, err := os.Create(output)
	checkErr(err)
	datafile.Write(jsonData)
	datafile.Close()
}

func md5Hash(text string) string {
	digest := md5.New()
	digest.Write([]byte(text))
	return hex.EncodeToString(digest.Sum(nil))
}

func help(){
	log.Fatal("Usage:\n ./credparser -f /path/to/creds.txt -o /path/to/results.json\n\n e.g: ./credparser -f ./file.txt -o ./data.json\n")
}
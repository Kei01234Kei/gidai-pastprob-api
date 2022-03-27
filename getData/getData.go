package main

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

type QueryParams struct {
	Faculty    string `json:"faculty"`
	Department string `json:"department"`
	Course     string `json:"course"`
}

type ProbInformation struct {
	Id            int       `json:"id" gorm:"column:id"`
	ImageURL      string    `json:"imageURL" gorm:"column:imageURL"`
	Faculty       string    `json:"faculty" gorm:"column:faculty"`
	Department    string    `json:"department" gorm:"column:department"`
	Course        string    `json:"course" gorm:"column:course"`
	SubjectName   string    `json:"subjectName" gorm:"column:subjectName"`
	YearOfStudent int       `json:"yearOfStudent" gorm:"column:yearOfStudent"`
	YearOfTest    int       `json:"yearOfTest" gorm:"column:yearOfTest"`
	Semester      string    `json:"semester" gorm:"column:semester"`
	Teacher       string    `json:"teacher" gorm:"column:teacher"`
	CreatedAt     time.Time `json:"createdAt" gorm:"column:createdAt;type:DATETIME"`
	UpdatedAt     time.Time `json:"updatedAt" gorm:"column:updatedAt;type:DATETIME"`
}

type Response struct {
	Data string `json:"body"`
}

func Handler(request events.APIGatewayProxyRequest) (Response, error) {
	switch request.HTTPMethod {
	case "GET":
		return ParseURL(request)
	default:
		log.Fatal("error")
		return Response{
			Data: "Unsupported method",
		}, nil
	}
}

func ParseURL(request events.APIGatewayProxyRequest) (Response, error) {
	queryParams := QueryParams{
		Faculty:    request.QueryStringParameters["faculty"],
		Department: request.QueryStringParameters["department"],
		Course:     request.QueryStringParameters["course"],
	}
	return GetDataFromDB(queryParams)
}

func GetDataFromDB(queryParams QueryParams) (Response, error) {
	if f, err := os.Stat(".env"); !(os.IsNotExist(err) || f.IsDir()) {
		err := godotenv.Load()
		if err != nil {
			log.Fatal(err)
		}
	}
	db, err := gorm.Open("mysql", os.Getenv("DB_ROLE")+":"+os.Getenv("DB_PASSWORD")+"@tcp("+os.Getenv("DB_HOST")+":3306)/"+os.Getenv("DB_NAME")+"?charset=utf8&parseTime=true&loc=Asia%2FTokyo")
	if err != nil {
		log.Fatal(err)
	}
	db.SingularTable(true)
	defer db.Close()
	probInformation := []*ProbInformation{}
	error := db.Where("faculty = ? AND department = ? AND course = ?", queryParams.Faculty, queryParams.Department, queryParams.Course).Find(&probInformation).Error
	if error != nil {
		log.Fatal(error)
	}
	data, _ := json.Marshal(probInformation)
	return Response{
		Data: string(data),
	}, nil
}

func main() {
	lambda.Start(Handler)
}

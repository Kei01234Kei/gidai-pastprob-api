package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type QueryParams struct {
	Faculty    string `json:"faculty"`
	Department string `json:"department"`
	Course     string `json:"course"`
}

type SubjectInformation struct {
	SubjectName   string    `json:"subjectName" gorm:"column:subjectName"`
	YearOfStudent int       `json:"yearOfStudent" gorm:"column:yearOfStudent"`
	Semester      string    `json:"semester" gorm:"column:semester"`
	Teacher       string    `json:"teacher" gorm:"column:teacher"`
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
	return GetSubjectDataFromDB(queryParams)
}

func GetSubjectDataFromDB(queryParams QueryParams) (Response, error) {
	if f, err := os.Stat(".env"); !(os.IsNotExist(err) || f.IsDir()) {
		err := godotenv.Load()
		if err != nil {
			log.Fatal(err)
		}
	}
	dsn := os.Getenv("DB_ROLE") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ":3306)/" + os.Getenv("DB_NAME") + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	subjectInformation := []*SubjectInformation{}
	error := db.Table("prob_information").Distinct("subjectName", "yearOfStudent", "semester", "teacher").Where("faculty = ? AND department = ? AND course = ?", queryParams.Faculty, queryParams.Department, queryParams.Course).Find(&subjectInformation).Error
	if error != nil {
		log.Fatal(error)
	}
	data, _ := json.Marshal(subjectInformation)
	return Response{
		Data: string(data),
	}, nil
}

func main() {
	lambda.Start(Handler)
}

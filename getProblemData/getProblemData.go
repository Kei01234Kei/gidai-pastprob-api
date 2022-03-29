package main

import (
	"encoding/json"
	"log"
	"os"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type QueryParams struct {
	SubjectName   string `json:"subjectName"`
	YearOfStudent int    `json:"yearOfStudent"`
	Semester      string `json:"semester"`
	Teacher       string `json:"teacher"`
}

type ProblemData struct {
	ImageURL   string `json:"imageURL" gorm:"column:imageURL"`
	YearOfTest int    `json:"yearOfTest" gorm:"column:yearOfTest"`
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
	yearOfStudent, err := strconv.Atoi(request.QueryStringParameters["yearOfStudent"])
	if err != nil {
		log.Fatal(err)
	}
	queryParams := QueryParams{
		SubjectName:   request.QueryStringParameters["subjectName"],
		YearOfStudent: yearOfStudent,
		Semester:      request.QueryStringParameters["semester"],
		Teacher:       request.QueryStringParameters["teacher"],
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
	problemData := []*ProblemData{}
	error := db.Table("prob_information").Select("imageURL", "yearOfTest").Where("subjectName = ? AND yearOfStudent = ? AND semester = ? AND teacher = ?", queryParams.SubjectName, queryParams.YearOfStudent, queryParams.Semester, queryParams.Teacher).Find(&problemData).Error
	if error != nil {
		log.Fatal(error)
	}
	data, _ := json.Marshal(problemData)
	return Response{
		Data: string(data),
	}, nil
}

func main() {
	lambda.Start(Handler)
}

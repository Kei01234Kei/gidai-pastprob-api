# gidai-pastprob-api

## プログラムについて

`getData.go`ではfaculty(学部)、department(学科)、course(コース)の3つのクエリパラメータを使ってDBを検索し、テーブル情報全てを返す。

`getSubjectData.go`ではfaculty(学部)、department(学科)、course(コース)の3つのクエリパラメータを使ってDBを検索し、subjectName(教科名)、yearOfStudent(学年)、semester(学期)、teacher(先生)の4つのデータを返す。

`getProblemData.go`ではsubjectName(教科名)、yearOfStudent(学年)、semester(学期)、teacher(先生)の4つのクエリパラメータを使ってDBを検索し、iamgeURL(画像のオブジェクトURL)、yearOfTest(テストの実施年)の2つのデータを返す。int型に変換不可能なパラメータをyearOfStudentに入れるとエラーコード(502)が帰ってくることに注意。

## DBテーブルについて

gidai-pastprob-apiを作成するにあたって、画像の保存にはAWS S3を、画像の基本情報とS3のオブジェクトURLの紐付けにはAWS RDS(My SQL)を使用する。
下の表がRDSのテーブル構造である。

| id                          | imageURL      | faculty       | department    | course        | subjectName   | yearOfStudent | yearOfTest   | semester      | teacher       | createdAt                                    | updatedAt                                                                |
| :-------------------------: | :-----------: | :-----------: | :-----------: | :-----------: | :-----------: | :-----------: | :----------: | :-----------: | :-----------: | :------------------------------------------: | :----------------------------------------------------------------------: |
| INT NOT NULL AUTO_INCREMENT | TEXT NOT NULL | TEXT NOT NULL | TEXT NOT NULL | TEXT NOT NULL | TEXT NOT NULL | INT NOT NULL  | INT NOT NULL | TEXT NOT NULL | TEXT NOT NULL | TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP | TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP |

# gidai-pastprob-api

gidai-pastprob-apiを作成するにあたって、画像の保存にはAWS S3を、画像の基本情報とS3のオブジェクトURLの紐付けにはAWS RDS(My SQL)を使用する。
下の表がRDSのテーブル構造である。

| id                          | imageURL      | faculty       | department    | course | subjectName   | yearOfStudent | yearOfTest   | semester      | teacher       | createdAt                                    | updatedAt                                                                |
| :-------------------------: | :-----------: | :-----------: | :-----------: | :----: | :-----------: | :-----------: | :----------: | :-----------: | :-----------: | :------------------------------------------: | :----------------------------------------------------------------------: |
| INT NOT NULL AUTO_INCREMENT | TEXT NOT NULL | TEXT NOT NULL | TEXT NOT NULL | TEXT   | TEXT NOT NULL | INT NOT NULL  | INT NOT NULL | TEXT NOT NULL | TEXT NOT NULL | TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP | TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP |

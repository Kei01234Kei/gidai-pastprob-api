# gidai-pastprob-api

gidai-pastprob-apiを作成するにあたって、画像の保存にはAWS S3を、画像の基本情報とS3のオブジェクトURLの紐付けにはAWS RDS(My SQL)を使用する。
下の表がRDSのテーブル構造である。

| id                          | imageURL      | faculty       | department    | course | subject       | studentYear  | year         | teacher       |
| :-------------------------: | :-----------: | :-----------: | :-----------: | :----: | :-----------: | :----------: | :----------: | :-----------: |
| int NOT NULL AUTO_INCREMENT | text NOT NULL | text NOT NULL | text NOT NULL | text   | text NOT NULL | int NOT NULL | int NOT NULL | text NOT NULL |

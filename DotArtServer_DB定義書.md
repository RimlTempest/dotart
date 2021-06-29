# MySql に保管する内容

ユーザーデータ
ユーザーの保存したキャンパスデータ
ユーザーの保存したパレットデータ
汎用パレットデータ

## ユーザーテーブル

テーブル名：users

```SQL
id PRIMARY KEY NOT NULL AUTO_INCREMENT
user_id varchar(256) PRIMARY KEY NOT NULL UNIQUE
display_name varchar(256) NOT NULL UNIQUE
e_mail varchar(256) NOT NULL UNIQUE
pass_word varchar(256) NOT NULL
createdAt datetime
updatedAt datetime
```

## キャンパステーブル

テーブル名：canvases

```SQL
id PRIMARY KEY NOT NULL AUTO_INCREMENT
canvas_id varchar(256) PRIMARY KEY NOT NULL UNIQUE
user_id varchar(256) NOT NULL
canvas_name varchar(256) NOT NULL
range varchar(256)
pallet text
index_data text
createdAt datetime
updatedAt datetime
FOREIGN KEY (user_id) REFERENCES users(user_id)
ON DELETE SET NULL ON UPDATE CASCADE
```

## 基本パレットテーブル

テーブル名：basic_pallet

```SQL
id PRIMARY KEY NOT NULL AUTO_INCREMENT
pallet_id varchar(256) PRIMARY KEY NOT NULL
pallet_name varchar(256) NOT NULL
summary varchar(256)
pallet_data text
createdAt datetime
```

## ユーザー個人用パレットテーブル

テーブル名：users_pallet

```SQL
id PRIMARY KEY NOT NULL AUTO_INCREMENT
user_id varchar(256) NOT NULL
pallet_id varchar(256) PRIMARY KEY NOT NULL
pallet_name varchar(256)
pallet_data text
createdAt datetime
updatedAt datetime
FOREIGN KEY (user_id) REFERENCES users(user_id)
ON DELETE SET NULL ON UPDATE CASCADE
```

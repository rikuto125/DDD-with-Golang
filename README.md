# DDD-with-Golang

This is a sample project to demonstrate how to implement DDD with Golang.

## How to run

```bash
$ mysql -u root -p
$ create database test;
```

```bash
$ go run main.go
```
mysqlでtestデータベースを作成ごtable作成がめんどくさかったら，
database.goの中のコメントアウトを外してください．(実行後に消してください)
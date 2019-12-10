## Instalasi

#### Package
```bash
go get -u github.com/labstack/echo/...
go get github.com/casbin/casbin
go get github.com/dgrijalva/jwt-go
go get github.com/go-sql-driver/mysql
go get github.com/spf13/viper
```

#### Jalanlan Aplikasi
```bash
go run casbin.go
```
#### Atau Build
```bash
go build casbin.go
casbin.exe or ./casbin
```
#### Desain Api
```bash
GET http://localhost:8080/login (All Group)
GET http://localhost:8080/person (All Group)
POST http://localhost:8080/person (Admin Only)
```
#### User

|USER|PASSWORD|GRUP|
|-----|:-----|-----|
|faiz|faiz|admin|
|nur|fatimah|user|

#### Testing
```curl
#Login With User Group
curl -i http://nur:fatimah@localhost:8080/login

#Response
HTTP/1.1 200 OK
Content-Type: application/json
X-Token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU5NTI5NDQsImlzcyI6IkdvbGFuZyBSQkFDIHdpdGggQ2FzYmluIiwidXNlcm5hbWUiOiJudXIiLCJyb2xlIjoidXNlciJ9.dFFexVwYbf4xGdQ7c4zqR0BUriWzQYCcvrT7cPoivPw

Date: Tue, 10 Dec 2019 03:42:24 GMT
Content-Length: 27

{"message":"Sukses Login"}


#Get Person With User
curl -X GET -H "x-token:eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU5NTI5NDQsImlzcyI6IkdvbGFuZyBSQkFDIHdpdGggQ2FzYmluIiwidXNlcm5hbWUiOiJudXIiLCJyb2xlIjoidXNlciJ9.dFFexVwYbf4xGdQ7c4zqR0BUriWzQYCcvrT7cPoivPw" http://localhost:8080/person

#Response
{"name":"Faizul Amaly","age":26}


#Post Person With User
curl -X POST -H "x-token:eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU5NTI5NDQsImlzcyI6IkdvbGFuZyBSQkFDIHdpdGggQ2FzYmluIiwidXNlcm5hbWUiOiJudXIiLCJyb2xlIjoidXNlciJ9.dFFexVwYbf4xGdQ7c4zqR0BUriWzQYCcvrT7cPoivPw" http://localhost:8080/person

#Response
{"message":"Forbidden"}
```
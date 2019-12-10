## Instalasi

#### Jalanlan Aplikasi
```golang
go run casbin.go
```
#### Atau Build
```golang
go build casbin.go
casbin.exe or ./casbin
```
#### Desain Api
|METHOD|PATH|
|------|:----|
|GET|/login|
|GET/POST|/person|

#### User
|USER|PASSWORD|GRUP|
|-----|:-----|-----|
|faiz|faiz|admin|
|nur|fatimah|user|

#### Database Sederhana
```sql
CREATE DATABASE IF NOT EXISTS `casbin`;
CREATE TABLE `person` (
  `id` int(11) NOT NULL,
  `nama` varchar(50) NOT NULL,
  `alamat` varchar(100) NOT NULL,
  `umur` int(3) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

INSERT INTO `person` (`id`, `nama`, `alamat`, `umur`) VALUES
(1, 'Faizul Amaly', 'Kambingan Barat, Madura', 26),
(2, 'Nur Fatimah', 'Sarasa, Sulawesi Barat', 21);

CREATE TABLE `user` (
  `username` varchar(25) NOT NULL,
  `password` varchar(50) NOT NULL,
  `role` varchar(15) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

INSERT INTO `user` (`username`, `password`, `role`) VALUES
('faiz', 'faiz', 'admin'),
('nur', 'fatimah', 'user');

ALTER TABLE `person`
  ADD PRIMARY KEY (`id`);

ALTER TABLE `user`
  ADD PRIMARY KEY (`username`);

ALTER TABLE `person`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;
COMMIT;
```

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


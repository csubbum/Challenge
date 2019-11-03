# Challenge

<br>
This is a file uploader/query api using golang.

### Application Requirements:
- Golang
- Curl or Postman to test the API

### Building the application
```
go build -o challenge *.go
```

### Running unit test and integration test
```
go test -spellCheckKey=<Replace it with provided key>
```

### Checking coverage for test
```
go test -coverprofile=coverage.out -spellCheckKey=<Replace it with provided key>
go tool cover -html=coverage.out
```
At this moment the code coverage is 86.1%. Main.go is not covered.

### Running the application
The application needs a spellcheck api key. The api key will be sent in email.
```
challenge -spellCheckKey=<Replace it with provided key>
```

##### Calling upload API
```
curl -X POST \
  http://127.0.0.1:8080/file/upload \
  -H 'content-type: multipart/form-data' \
  -F myFile=@<Your file name>
```

##### Calling query API
The spellcheck and trim are optional
```
curl -X GET \
  'http://127.0.0.1:8080/file/query?fileName=<Your file name>&spellCheck=<true|false>&trim=<word to trim>'
```

### Future design/improvement consideration

* Current design stores the data in memory. A database like dynamodb should be used to persist the data.
Only the dao layer needs to be modified for persisting the data in the database
* API's should be secured

### Reference
* https://matt.aimonetti.net/posts/2013-07-golang-multipart-file-upload-example/
* https://medium.com/@masnun/making-http-requests-in-golang-dd123379efe7
* https://marcofranssen.nl/go-webserver-with-gracefull-shutdown/
* https://gobyexample.com/signals 

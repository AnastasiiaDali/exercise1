
## Learning Golang with Salt / exercise1

"Add" program takes file names (files are in data folder) passed through command line<br> 
OR <br>
string of numbers passed through command line and returns a sum


"Math" program spins up the server and returns the sum of numbers passed in URL.
Once you spinned up the server, use your postman to send POST request
```bash
http://localhost:8082/add?num=76&num=65
```
Should return 141

### To build Add program go inside cmd folder and run
```bash
go build -o add cmd/add/main.go
```

### To build Math program go inside cmd folder and run
```bash
go build -o math cmd/math/main.go
```

### Run Add Locally

```bash
 ./add --input-file="input2.csv" --input-file="input1.txt"
```
OR
```bash
 ./add --input-numbers="1,2,3" 
```

### Run Math Locally
```bash
./math
```

### Running Tests
To run tests, run the following command inside exercise1 directory:

```bash
 go test ./...
```


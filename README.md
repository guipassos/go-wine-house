# go-wine-house
"Go Winehouse" is a simple API written in Golang to assist in the maintenance of wineries.

## CONFIG
* Set up your database in the "api\config" folder

## RUN
* The main file is at path "\api\src"
```bash
go run main.go
```

## Examples / Usage - Routes

### 1. POST - Create Wine
#### Method: POST
* /wines
* Example: http://localhost:3000/wines
#### Headers: 
* Content-Type: application/json
#### Body:  
* {"name": "wine-name","brand": "wine-brand","description":"...","year":2019 ,"country":"France","quantity":10,"status":true}

### 2. GET - Get All Wines
#### Method: GET
* /wines
* Example: http://localhost:3000/wines

### 3. GET - Get Wine By Id
#### Method: GET
* /wines/1
* Example: http://localhost:3000/wines/1

### 4. PUT - Edit Wine By Id
#### Method: PUT
* /wines/1
* Example: http://localhost:3000/wines/1
#### Headers: 
* Content-Type: application/json
#### Body:  
* {"name": "wine-name","brand": "wine-brand","description":"...","year":2010 ,"country":"Brazil","quantity":20,"status":true}

### 5. PUT - Enable or Disable Wine By Id
#### Method: PUT
* /wines/1/ENABLE_OR_DISABLE
* Example: http://localhost:3000/wines/1/enable

### 6. DELETE - Delete Wine By Id
#### Method: DELETE
* /wines/1
* Example: http://localhost:3000/wines/1

## Contributing
Pull requests are welcome. 

## License
[MIT](https://choosealicense.com/licenses/mit/)

### Project Structure
```
project/
├── controllers/
│   ├── customerController.go
│   ├── transferController.go
├── data/
│   ├── customerData.json
│   ├── historyData.json
│   ├── tokenBlacklistData.json
├── helpers/
│   ├── jsonFileHelper.go
│   ├── passwordHelper.go
│   ├── tokenHelper.go
├── middlewares/
│   ├── authMiddleware.go
├── models/
│   ├── customerModel.go
│   ├── historyModel.go
│   ├── tokenBlacklistModel.go
├── repositories/
│   ├── customerRepository.go
│   ├── historyRepository.go
│   ├── tokenBlacklistRepository.go
├── routes/
│   ├── customerRouter.go
│   ├── transferRouter.go
├── main.go
```
### Features
- gin-gonic
- golang-jwt
- json-iterator
### How to use the project
```
go mod download
```
```
go run main.go
```
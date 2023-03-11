
`go get go.mongodb.org/mongo-driver/mongo`

`mongodb://<username>:<password>@<host>:<port>`

```go

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// MongoDB 서버 URL
	url := "mongodb://localhost:27017"

	// MongoDB 클라이언트 생성
	client, err := mongo.NewClient(options.Client().ApplyURI(url))
	if err != nil {
		log.Fatal(err)
	}

	// MongoDB 서버에 연결
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// 클라이언트 연결 해제
	defer func() {
		if err = client.Disconnect(context.Background()); err != nil {
			log.Fatal(err)
		}
	}()

	// 연결 확인
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDB connection success!")
}

```

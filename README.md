
usage

```
go get -u github.com/msd-rs/msd-sdk-go
```

```go
import (
    "github.com/msd-rs/msd-sdk-go"
    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials/insecure"
)

func main() {
    host := "127.0.0.1:50051"
    conn, err := grpc.Dial(host, grpc.WithTransportCredentials(insecure.NewCredentials()))
    if err != nil {
        fmt.Println("dail msd server failed", err)
        return
    }
    client := msd.NewApiV1Client(conn)
    req := &msd.GetDataFrameRequest{
        Sql: "select * from kline1d.SH000001"
    }
    resp, err := client.Get(req)
    if err != nil {
        fmt.Println("query failed", err)
        return
    }
    fmt.Println(resp)
    conn.Close()
}
```


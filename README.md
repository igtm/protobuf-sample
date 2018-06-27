# protobuf-sample
Protocol BufferをGo/Echoで試してみた。

# Usage

```
# パッケージインストール
dep ensure

# サーバ起動
go run main.go

# クライアント実行(Protobuf)
go run client.go
# クライアント実行(Json)
go run client-json.go

Content-Lengthに２倍ほどの差があることがわかって
protobufすごいってなる。
```

# 実行結果例

```
protobuf-sample $ go run client.go 
[status] 200
[header] Content-Length: 49
[header] Content-Type: application/protobuf
[header] Date: Wed, 27 Jun 2018 15:43:38 GMT
[body] [10 23 8 1 18 9 229 140 151 230 181 183 233 129 147 26 8 104 111 107 107 97 105 100 111 10 22 8 47 18 9 230 178 150 231 184 132 231 156 140 26 7 111 107 105 110 97 119 97]
1, 北海道, hokkaido
47, 沖縄県, okinawa

protobuf-sample $ go run client-json.go 
[status] 200
[header] Content-Type: application/json; charset=UTF-8
[header] Date: Wed, 27 Jun 2018 15:43:50 GMT
[header] Content-Length: 97
[body] [{"Id":1,"Name":"北海道","Romaji":"hokkaido"},{"Id":47,"Name":"沖縄県","Romaji":"okinawa"}]

```
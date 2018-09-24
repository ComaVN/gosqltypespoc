# PoC for unexpected behaviour of field types in Go (My)SQL library

```
$ docker-compose up --build go
(...)
Attaching to gosqltypespoc_go_1
go_1  | 2018/09/24 13:16:29 Started
go_1  | 2018/09/24 13:16:29 SELECT 123 WHERE ? = 1
go_1  | 2018/09/24 13:16:29 int64: 123
go_1  | 2018/09/24 13:16:29 SELECT 123 WHERE 1 = 1
go_1  | 2018/09/24 13:16:29 []uint8: []byte{0x31, 0x32, 0x33}
go_1  | 2018/09/24 13:16:29 Ready
```

package common

import "sync"

/*
sync.WaitGroup 用于等待多个 Goroutine 完成。
*/
var Wg sync.WaitGroup

var JwtSecretKey = []byte("JWT_SECRET")

package main

import "fmt"

func main() {
	//аллоцировали память для структуры ConnectionPool
	connectionPool := ConnectionPool{}

	//проинициализировали его пятью соединениями
	connectionPool.init(5)

	userId1 := "qwe123"
	userId2 := "qwe124"
	userId3 := "qwe125"

	//3 наших юзера с разными userId забрали по соединению
	connectionPool.getConnection(userId1)
	connectionPool.getConnection(userId2)
	connectionPool.getConnection(userId3)

	//print должен показывать 2 свободных и 3 занятых соединения
	fmt.Println(connectionPool)

	connectionPool.returnConnection(userId1)

	//print должен показывать 3 свободных и 2 занятых соединения
	fmt.Println(connectionPool)
}

type ConnectionPool struct {
	//connection -> user_id
	connections map[Connection]string
	// Connectionx1 -> "qwe123"
	// Connectionx2 -> "qwe124"
	// Connectionx3 -> "qwe125"
}

type Connection struct {
	Id string

	//"50" - 50 секунд
	Timeout string
}

func (c *ConnectionPool) getConnection(userId string) Connection {
	panic("implement me!")
}

func (c *ConnectionPool) returnConnection(userId string) {
	panic("implement me!")
}

func (c *ConnectionPool) init(size int) {
	panic("implement me!")
}

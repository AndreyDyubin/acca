package main

func main() {
	// - внутренний grpc АПИ
	// - хандлер для платежек

	/*
		Входящая операция падает в общую очередь
		Колторая обрабатывается в горутине
		Все состояния персистятся в PG
		В случае падения процесса очередь воссоздается из БД (то есть сохранять состояния команд?)
	*/
}

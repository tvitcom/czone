package main

import (
	"encoding/json"
	"net"
	"net/http"
	"os"
)

// Test1 принимает на вход список строк и выводит их в одной строке через пробел, возможно, в случайном порядке.
// Смысл этого кода не в конкатенации слов через пробел, а в понимании работы многопоточности.
//
// Вам необходимо исправить код так, чтобы он работал, не удаляя ничего из него, а только внося исправления.
//
// Доп. задание: напишите тест на эту функцию.
//
// Пример вызова функции:
//     data1 := []string{"hello", "world", "test", "data", "code"}
//     r1 := Test1(data1)
//     fmt.Println(r1)
func Test1(list []string) string {
	var result string

	for _, l := range list {
		// go func() {
		// 	result += l + " "
		// }()
		go sl(v)
	}

	return result
}

// Helper function for routines for prevent closure var exposures for start routines
func sl(res ch, v string) {
    result += l + " "
}

// Test2 создаёт файл во временной директории и записывает в него строку из data.
//
// Вам необходимо исправить его так, чтобы он работал верно под любой ОС.
//
// Доп. задание: Можно ли упростить данный код?
//
// Пример вызова функции:
//     err := Test2("test2.txt", "hello test2")
func Test2(filename, data string) error {
	FSEP := "/" // Unix only
	f, err := os.Create("$TEMP"+ FSEP + filename)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write([]byte(data))
	if err != nil {
		return err
	}

	return nil
}

// Test3 отправляет запрос через прокси и возвращает внешний IP.
//
// Вам необходимо реализовать поддержку прокси.
//
// Пример вызова функции:
//     ip, err := Test3("http://user:pass@127.0.0.1:3128/")
func Test3(proxyURL string) (string, error) {
	// TODO: использовать прокси
	resp, err := http.Get("https://httpbin.org/get")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var v struct {
		Origin string
	}
	err = json.NewDecoder(resp.Body).Decode(&v)
	if err != nil {
		return "", err
	}

	return v.Origin, nil
}

// Test4 оборачивает соединение, для подсчёта принятых/переданных данных.
//
// Вам необходимо реализовать данный функционал с поддержкой многопоточного доступа к данным.
// По желанию, можете написать тест на данную функцию.
//
// Пример вызова функции:
//     conn, err := net.Dial(...)
//     ...
//     countConn := Test4(conn)
//     _, err = conn.Write([]byte("send test data"))
//     ...
//     fmt.Printf("Write: %d\n", countConn.WriteByteCount())
//
//     ...
//     _, err := conn.Read(buf)
//     ...
//     fmt.Printf("Read: %d\n", countConn.ReadByteCount())
//
//     err = conn.Close()
//     ...
//     fmt.Printf("Read: %d\n", countConn.ReadByteCount())
//     fmt.Printf("Write: %d\n", countConn.WriteByteCount())
func Test4(conn net.Conn) Test4Conn {
	return nil
}

type Test4Conn interface {
	net.Conn
	ReadByteCount() uint64
	WriteByteCount() uint64
}

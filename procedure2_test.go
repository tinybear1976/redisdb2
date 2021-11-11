package redisdb2

import "testing"

func Test_SetEx(t *testing.T) {
	New("test", "127.0.0.1:6379", "", 0)
	conn, _ := Connect("test")
	err := SETPX(conn, "aaa", "go_test", 20000)
	t.Log(err)
	Diconnect(conn)
	Destroy()
}
func Test_EXEC(t *testing.T) {
	New("test", "127.0.0.1:6379", "", 0)
	conn, _ := Connect("test")
	repl, err := EXEC(conn, "set", "k1", "go_test", "ex", "30")
	t.Log(repl, err)
	Diconnect(conn)
	Destroy()
}

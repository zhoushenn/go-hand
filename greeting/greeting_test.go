package greeting

import "testing"

//_test.go 结尾的文件会被认为是测试文件. 运行go test 进行测试

func TestHello(t *testing.T) {

	Hello()

	//t.Log("hello world")

	t.Fatal("测试不通过")
}

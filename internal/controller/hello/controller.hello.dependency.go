package hello

//go:generate mockgen -source=./controller.hello.dependency.go -destination=./controller.hello.dependency.mock.go -package=hello
type HelloRepositoryInterface interface {
	GetHelloMessage() (string, error)
	GetHelloFromMongo()
}

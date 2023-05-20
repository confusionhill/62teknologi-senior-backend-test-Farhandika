package hello

func (r HelloRepository) GetHelloMessage() (string, error) {
	return "hello", nil
}

func (r *HelloRepository) GetHelloFromMongo() {
}

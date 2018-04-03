swaggergen:
	go get -u github.com/go-swagger/go-swagger/cmd/swagger
	swagger generate model --spec=model/model.yaml --model-package=model --target=.

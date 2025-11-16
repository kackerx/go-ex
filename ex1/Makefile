.PHONY: mock
mock:
	@mockgen -source=domain/aggregate/repository/user_repository.go -destination=domain/aggregate/repository/mock/user_repository_gen.go -package=repomock
	@go mod tidy
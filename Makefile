testuser:
	go test ./features/user... -coverprofile=cover.out && go tool cover -html=cover.out

testclass:
	go test ./features/class... -coverprofile=cover.out && go tool cover -html=cover.out

testlog:
	go test ./features/log... -coverprofile=cover.out && go tool cover -html=cover.out

run:
	go run main.go
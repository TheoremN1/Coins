cd database_service
start cmd.exe /k "dotnet run --project database_api"

cd ../users_service
start cmd.exe /k "go run cmd/main.go"

cd ../products_service
start cmd.exe /k "go run cmd/main.go"

cd ../requests_service
start cmd.exe /k "go run cmd/main.go"

cd ../react_app
start cmd.exe /k "npm run dev"


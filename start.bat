cd database_service
start cmd.exe /k "dotnet run --nodocker"

cd ../users_service
start cmd.exe /k "go run cmd/main.go --nodocker"

cd ../products_service
start cmd.exe /k "go run cmd/main.go --nodocker"

cd ../requests_service
start cmd.exe /k "go run cmd/main.go --nodocker"

cd ../react_app
start cmd.exe /k "npm run dev"

cd ../nginx/nodocker
start cmd.exe /k "nginx"
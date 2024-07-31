FROM mcr.microsoft.com/dotnet/aspnet:8.0 AS base
WORKDIR /app

FROM mcr.microsoft.com/dotnet/sdk:8.0 AS build
WORKDIR /src
COPY ["database_service.csproj", "."]
RUN dotnet restore "database_service.csproj"
COPY . .
RUN dotnet build "database_service.csproj" -c Release -o /app/build

FROM build AS publish
RUN dotnet publish "database_service.csproj" -c Release -o /app/publish

FROM base AS final
WORKDIR /app
COPY --from=publish /app/publish .
ENTRYPOINT ["dotnet", "database_service.dll"]
# control2-backend

## Setup

Go 1.18+

```shell
go mod download
```

## Development

Start the development server on http://localhost:5222

```shell
go build
./control2

# or

go run main.go
```

## Production

Generate static web and move it to backend static embed:

```shell
go generate
```

## ToDo

- Customised recovery middleware.
- More cmdlets.
- Better and more formal cmdlet norm.
- Rename API.
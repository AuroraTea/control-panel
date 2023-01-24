# control2-backend

## Setup

```shell
go mod download
```

## Development

Start the development server on http://localhost:5222

```shell
go build
./control2
```

## Production

Generate static web and move it to backend static embed:

```shell
rsrc -manifest .\control2.exe.manifest -o control2.exe.syso
go build
./control2
```

## ToDo

- Customised recovery middleware.
- More cmdlets.
- Better and more formal cmdlet norm.
- Rename API.
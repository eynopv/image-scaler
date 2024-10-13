# Image Scaler API

This API is designed for image storage and retrieval in different resolutions while maintaining
aspect ratio of original image.

## Project structure

```
├── Makefile
├── README.md
├── cmd
│   └── api
│       ├── errors.go
│       ├── helpers.go
│       ├── helpers_test.go
│       ├── image.go
│       ├── main.go
│       └── routes.go
├── go.mod
├── go.sum
├── internal
│   ├── data
│   │   ├── image.go
│   │   ├── image_test.go
│   │   ├── size.go
│   │   └── size_test.go
│   └── validator
│       ├── validator.go
│       └── validator_test.go
└── pkg
    ├── storage
    │   ├── file_system.go
    │   └── file_system_test.go
    └── transforms
        ├── scale.go
        └── scale_test.go
```

## Getting Started

1. Download source code

```sh
git clone http://github.com/eynopv/image-scaler
```

2. Go to directory

```sh
cd image-scaler
```

3. Run for production

**Using `make`**

```sh
make run
```

**Manually**

```sh
go build -ldflags="-s -w" -o=./bin/api ./cmd/api
./bin/api
```

## Commands

| Command | Code         | Description                                        |
| ------- | ------------ | -------------------------------------------------- |
| run     | `make run`   | Run for production.                                |
| dev     | `make dev`   | Run for development.                               |
| tidy    | `make tidy`  | Format code and validate modules.                  |
| build   | `make build` | Build binary.                                      |
| test    | `make test`  | Run tests.                                         |
| clean   | `make clean` | Clean up environment (removes `bin` and `uploads`) |

## Endpoints

| Method | Endpoint     | Description     |
| ------ | ------------ | --------------- |
| `POST` | `/image`     | Store image.    |
| `GET`  | `/image/:id` | Retrieve image. |

### Retrieving image

| Query Parameter | Required | Default Value | Description                       |
| --------------- | -------- | ------------- | --------------------------------- |
| `width`         | false    | 0             | Maximum width of returned image.  |
| `height`        | false    | 0             | Maximum height of returned image. |

If no width and height passed then orginal image will be returned.

## Architectural desicions

- UUID is used for image file name to avoid name collision.
- Scaled images are stored along side original image, this increases storage space but avoids
  scaling same images repeatedly. This can be removed if CDN or other caching solution is used
  infront.
- Width and Height are added to the name of image file.

## Further development

- User authentication
- Access permissions
- e2e testing

# WEATHER APP GO

This is a test app that consists of 2 endpoints for getting and putting weather by inputed city

## Notes

- Mongo Atlas was used as database choice
- Open Weather API was used to retrieve weather related info by given city

## Usage

1. Download this project
2. Download all required dependecies
3. Create '.env' file and fill it by given '.env.example' file
4. Run app:
```console
    $ go run ./cmd/api
```

## Endpoints

- **GET: /weather?city=Name** - retrieving weather by given city name
- **PUT: /weather?city=Name** - updating (inserting if doesn't exist) weather by given city name
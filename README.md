# Test task for the What to Farm company

Create a Go application that includes server and client side.

The application must be implemented as a CLI based on https://github.com/spf13/cobra

## Server-side

Must provide data provided from Binance public API https://api.binance.com/api/v3/ticker/price

Two methods must be implemented

### GET

```curl
curl http://localhost:3001/api/v1/rates?pairs=BTC-USDT,ETH-USDT
```

```json
{ "ETH-USDT": 1780.123, "BTC-USDT": 46956.45 }
```

### POST

```curl
curl -X POST --data '{ "pairs": ["BTC-USDT", "ETH-USDT"] }' [http://localhost:3001/api/v1/rates](http://localhost:3001/api/v1/rates)
```

```json
{ "ETH-USDT": 1780.123, "BTC-USDT": 46956.45 }
```

The server-side application must run as follows

```curl
$ go run . server
Listening at port 3001...
```

## Client-side

Makes a GET request to the server and prints the result in the terminal

```curl
$ go run . rate ---pair=ETH-USDT
1780.123
```

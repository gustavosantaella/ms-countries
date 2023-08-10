
# Country service

This API provide differents endpoint to get information about countries, states, cities, and currencies. Also you can do currencies conversions.
## Demo

You can test the api here [Test api here](https://ms-countries.onrender.com/)
## API Reference

#### Get all countries

```http
  GET /countries/
```
#### Get item

```http
  GET /countries/${iso3}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `iso3`      | `string` | **Required**. Id of item to fetch |


### List currencies 
```http
  GET /currencies
```

#### Currency conversion
Get aprox currency conversion

```http
  GET /currencies/conversion?base=${base}&qoute=${qoute}&amount=${amount}
```

| Query Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `base`      | `string` | **Required**. currency base |
| `qoute`      | `string` | **Required**. currency to convert |
| `amount`      | `string` | **Required**. amount to convert |




## Tech Documentation

- [Gin framework](https://github.com/gin-gonic/gin)
- [Supabase](supabase.com)
- [Go](https://github.com/go-pg/pg)


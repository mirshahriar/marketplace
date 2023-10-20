# List Products

## Description

We have the following parameters for listing products:

- `page` - page number. Default is `1`.
- `limit` - number of products per page. Default is `10`.
- `sort_by` - sort by field of the products. Default is `id`.
- `sort_direction` - sort direction of the products. Default is `asc`.

## Example

#### Request

For Local
```http
GET /products HTTP/1.1
Host: localhost:8080
Content-Type: application/json
```

#### Response

> Please [check this](./../../docs/guide/types.md#type-page) for more details about the response body.

```json
{
  "data": [
    {
      "id": 1,
      "name": "product 1",
      "description": "This is a test product",
      "price": 10,
      "CreatedAt": "2023-10-20T04:51:45.626+06:00",
      "UpdatedAt": "2023-10-20T04:51:45.626+06:00"
    },
    {
      "id": 2,
      "name": "product 2",
      "description": "This is a test product",
      "price": 10,
      "CreatedAt": "2023-10-20T04:51:45.626+06:00",
      "UpdatedAt": "2023-10-20T04:51:45.626+06:00"
    }
  ],
  "page": 1,
  "size": 2,
  "total": 2
}
```

## Errors


| Code | Message                   | Description                                                     |
|------|---------------------------|-----------------------------------------------------------------|
| 400  | `Invalid request`         | Error Message                                                   |
| 500  | `Internal database error` | Error Message                                                   |


#### Links

- [Create](./create.md) a new Product.
- [Get](./get.md) a Product.
- [Update](./update.md) a Product.
- [Delete](./delete.md) a Product.

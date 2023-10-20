# Get a Product

## Description

We can get a product by its `id`.

## Example

#### Request

For Local
```http
GET /products/:id HTTP/1.1
Host: localhost:8080
Content-Type: application/json
```

Here, `:id` is the id of the product.

#### Response

> Please [check this](./../../docs/guide/types.md#type-productresponse) for more details about the response body.

```json
{
    "id": 1,
    "name": "product 1",
    "description": "This is a test product",
    "price": 10.0,
    "created_at": "2021-10-20T04:04:30.000Z",
    "updated_at": "2021-10-20T04:04:30.000Z"
}
```

## Errors

| Code | Message                       | Description                                                     |
|------|-------------------------------|-----------------------------------------------------------------|
| 400  | `Invalid request`             | Error Message                                                   |
| 400  | `Requested product not found` |                                                                 |
| 500  | `Internal database error`     | Error Message                                                   |


#### Links

- [Create](./create.md) a new Product.
- [List](./list.md) all Products.
- [Update](./update.md) a Product.
- [Delete](./delete.md) a Product.

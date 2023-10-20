# Create a Product

## Description

We have the following parameters for creating a product:

- `name` - name of the product. This is required. Maximum length is `100`.
- `description` - description of the product. This is required. Maximum length is `200`.
- `price` - price of the product. Price must be a positive number.

> Note: Product name can be duplicate


## Example

#### Request

For Local
```http
POST /products/ HTTP/1.1
Host: localhost:8080
Content-Type: application/json
Authorization: Bearer CONExgMnrPpNBnZm
```

Body

> Please [check this](./../../docs/guide/types.md#type-productbody) for more details about the request body.

```json
{
    "name": "product 1",
    "description": "This is a test product",
    "price": 10.0
}
```

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

| Code | Message                   | Description                                                     |
|------|---------------------------|-----------------------------------------------------------------|
| 401  | `Unauthorized`            |                                                                 |
| 400  | `Invalid request`         | Error Message                                                   |
| 400  | `Invalid input data`      | You must provide the name                                       |
| 400  | `Invalid input data`      | You must provide a valid name with maximum length of 100        |
| 400  | `Invalid input data`      | You must provide the description                                |
| 400  | `Invalid input data`      | You must provide a valid description with maximum length of 200 |
| 400  | `Invalid input data`      | You must provide a valid price greater than or equal to 0       |
| 500  | `Internal database error` | Error Message                                                   |


#### Links 

- [List](./list.md) all Products.
- [Get](./get.md) a Product.
- [Update](./update.md) a Product.
- [Delete](./delete.md) a Product.

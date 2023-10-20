# Delete a Product 

## Description

We can delete a product by its `id`.

## Example

#### Request

For Local
```http
DELETE /products/:id HTTP/1.1
Host: localhost:8080
Content-Type: application/json
Authorization: Bearer CONExgMnrPpNBnZm
```

Here, `:id` is the id of the product.

## Errors


| Code | Message                       | Description                                                     |
|------|-------------------------------|-----------------------------------------------------------------|
| 401  | `Unauthorized`                |                                                                 |
| 400  | `Invalid request`             | Error Message                                                   |
| 400  | `Requested product not found` |                                                                 |
| 500  | `Internal database error`     | Error Message                                                   |


#### Links

- [Create](./create.md) a new Product.
- [List](./list.md) all Products.
- [Get](./get.md) a Product.
- [Update](./update.md) a Product.

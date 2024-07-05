# Микросервис для управления пользователями, товарами, заказами и платежами

## Introduction

Этот проект представляет собой REST API для управления пользователями, товарами, заказами и платежами. Ниже приведены доступные эндпоинты и их функционал.

## Installation

1. Clone the repository:

```sh
git clone https://github.com/togzhanzhakhani/projects.git
cd projects
```

2. Build the project:

```sh
make build
```

3. Run the server:

```sh
make run
```

## Deployed on RENDER:

Base URL: https://projects-j02i.onrender.com

# API Endpoints
## Users
#### GET /users: Get a list of all users.
#### Пример ответа:
```sh
[
    {
        "id": 1,
        "name": "John Doe",
        "email": "john.doe@example.com",
        "registration_date": "2024-07-06T12:00:00Z",
        "role": "user"
    },
    {
        "id": 2,
        "name": "Jane Smith",
        "email": "jane.smith@example.com",
        "registration_date": "2024-07-06T13:00:00Z",
        "role": "admin"
    }
]
```
#### POST /users: Create a new user.
### Request Body:

```sh
{
    "name": "John Doe",
    "email": "johndoe@example.com",
    "role": "admin"
}
```
#### GET /users/{id}: Get details of a specific user.
### Пример ответа:
```sh
{
    "id": 3,
    "name": "John Doe",
    "email": "johndoe@example.com",
    "registration_date": "2024-07-06T14:00:00Z",
    "role": "admin"
}
```
#### PUT /users/{id}: Update details of a specific user.
### Пример запроса:
```sh
{
    "id": 3,
    "name": "Updated Name",
    "email": "johndoe@example.com",
    "registration_date": "2024-07-06T14:00:00Z",
    "role": "admin"
}
```
#### DELETE /users/{id}: Delete a specific user.
#### Пример ответа:
```sh
HTTP 204 No Content
```
#### GET /users/search?name={name}: Find users by name.
#### Пример запроса:
```sh
GET /users/search?name=John
```
#### Пример ответа:
```sh
[
    {
        "id": 1,
        "name": "John Doe",
        "email": "john.doe@example.com",
        "registration_date": "2024-07-06T12:00:00Z",
        "role": "user"
    }
]
```
#### GET /users/search?email={email}: Find users by email.
#### Пример запроса:
```sh
GET /users/search?email=jdoe@mple.com
```
#### Пример ответа:
```sh
[
    {
        "error": "No users found"
    }
]
```
## Товары
#### GET /products: Get a list of all products.
```sh
[
    {
        "id": 1,
        "name": "Product 1",
        "description": "Description of Product 1",
        "price": 19.99,
        "category": "Category A"
    },
    {
        "id": 2,
        "name": "Product 2",
        "description": "Description of Product 2",
        "price": 29.99,
        "category": "Category B"
    }
]
```
#### POST /products: Create a new products.
### Request Body:

```sh
{
    "name": "New Product",
    "description": "Description of New Product",
    "price": 24.99,
    "category": "Category C"
}
```
#### GET /products/{id}: Get details of a specific product.
### Пример ответа:

```sh
{
    "id": 3,
    "name": "New Product",
    "description": "Description of New Product",
    "price": 24.99,
    "category": "Category C"
}
```
#### PUT /products/{id}: Update details of a specific product.
### Request Body:

```sh
{
    "id": 3,
    "name": "Updated Product Name",
    "description": "Description of New Product",
    "price": 19.99,
    "category": "Category A"
}
```
#### DELETE /products/{id}: Delete a specific product.
### Пример ответа:
```sh
HTTP 204 No Content
```
#### GET /products/search?name={name}
### Пример запроса:
```sh
GET /products/search?name=Product
```
### Пример ответа:
```sh
[
    {
        "id": 1,
        "name": "Product 1",
        "description": "Description of Product 1",
        "price": 19.99,
        "category": "Category A"
    }
]
```
#### GET /products/search?category={category}
### Пример запроса:
```sh
GET /products/search?category=Category A
```
## Заказы
#### GET /orders
```sh
[
    {
        "id": 1,
        "user_id": 1,
        "product_ids": [1, 2],
        "total_price": 69.97,
        "order_date": "2024-07-06T12:00:00Z",
        "status": "new"
    }
]
```
#### POST /orders: Create a new order.
### Request Body:

```sh
{
    "user_id": 1,
    "product_ids": [1, 2],
    "total_price": 69.97,
    "status": "new"
}
```
#### GET /orders/{id}
### Пример ответа:

```sh
{
    "id": 1,
    "user_id": 1,
    "product_ids": [1, 2],
    "total_price": 69.97,
    "order_date": "2024-07-06T12:00:00Z",
    "status": "new"
}
```
#### PUT /orders/{id}
### Request Body:

```sh
{
    "user_id": 1,
    "product_ids": [1, 2],
    "total_price": 69.97,
    "status": "completed"
}
```
#### DELETE /orders/{id}: Delete a specific order.
### Пример ответа:
```sh
HTTP 204 No Content
```
#### GET /orders/search?user={userId}
### Пример запроса:
```sh
GET /orders/search?user=1
```
### Пример ответа:
```sh
[
    {
        "id": 1,
        "user_id": 1,
        "product_ids": [1, 2],
        "total_price": 69.97,
        "order_date": "2024-07-06T12:00:00Z",
        "status": "new"
    }
]
```
#### GET /orders/search?status={status}
### Пример запроса:
```sh
GET /orders/search?status=completed
```
### Пример ответа:
```sh
[
    {
        "id": 1,
        "user_id": 1,
        "product_ids": [1, 2],
        "total_price": 69.97,
        "order_date": "2024-07-06T12:00:00Z",
        "status": "completed"
    }
]
```
## Заказы
#### GET /payments
```sh
[
    {
        "id": 1,
        "user_id": 1,
        "order_id": 1,
        "amount": 69.97,
        "payment_date": "2024-07-06T15:00:00Z",
        "payment_status": "success"
    }
]
```
#### POST /payments: Create a new payment.
### Request Body:

```sh
 {
    "user_id": 1,
    "order_id": 1,
    "amount": 69.97,
}
```
#### GET /payments/{id}
### Пример ответа:

```sh
{
    "id": 1,
    "user_id": 1,
    "order_id": 1,
    "amount": 69.97,
    "payment_date": "2024-07-06T15:00:00Z",
    "payment_status": "success"
}
```
#### PUT /payments/{id}
### Request Body:

```sh
 {
    "user_id": 1,
    "order_id": 1,
    "amount": 99.97,
}
```
#### DELETE /payments/{id}: Delete a specific payment.
### Пример ответа:
```sh
HTTP 204 No Content
```
#### GET /payments/search?user={userId}
### Пример запроса:
```sh
GET /payments/search?user=1
```
### Пример ответа:
```sh
[
    {
        "id": 1,
        "user_id": 1,
        "order_id": 1,
        "amount": 69.97,
        "payment_date": "2024-07-06T15:00:00Z",
        "payment_status": "success"
    }
]
```
#### GET /payments/search?order={orderId}
### Пример запроса:
```sh
GET GET /payments/search?order=1
```
### Пример ответа:
```sh
[
    {
        "id": 1,
        "user_id": 1,
        "order_id": 1,
        "amount": 69.97,
        "payment_date": "2024-07-06T15:00:00Z",
        "payment_status": "success"
    }
]
```
#### GET /payments/search?status={status}
### Пример запроса:
```sh
GET /payments/search?status=failed
```
### Пример ответа:
```sh
[
    {
        "error": "Payments not found"
    }
]

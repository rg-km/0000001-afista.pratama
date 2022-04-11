# Cashier App

## Goals
- Learn how to read and write to a CSV file (database concept)
- Learn CRUD operations
- Learn how to render data to the frontend
- Learn how to make and call APIs
- Learn code reuse: common `repository` code is reused by different frontend (terminal & API)



## 3 modes

### Terminal Mode

Video (click the image):

[![Terminal](https://img.youtube.com/vi/iZOPT3axoi4/maxresdefault.jpg)](https://youtu.be/iZOPT3axoi4)

### TView Mode

Video (click the image):

[![TView](https://img.youtube.com/vi/e1yX1O8pKYs/maxresdefault.jpg)](https://youtu.be/e1yX1O8pKYs)


### API Mode

Video (click the image):

[![API](https://img.youtube.com/vi/lI-RO36El08/maxresdefault.jpg)](https://youtu.be/lI-RO36El08)

Available APIs:
- `/api/user/login?username=<username>&password=<password>`

success :
```json
{
    "username": "aditira"
}
```

fail :
```json
{
    "error": "Login Failed"
}       
```

- `/api/user/logout?username=<username>`

success : kosong

error :
```json
{
    "error": "no user is logged in"
}
```

- `/api/dashboard?cash=<cash>`

success : (response code 200)
```json
{
    "username": "aditira", // login username
    "cart_items": null,
    "total_price": 0,
    "money_changes": 20000
}
```

error :
```json
{
    "error": "no user is logged in"
}
```

- `/api/products`

```json
{
    "products": [
        {
        "name": "Orange",
        "price": 5000,
        "category": "Fruits"
        },
        {
        "name": "Apple",
        "price": 2000,
        "category": "Fruits"
        },
        {
        "name": "Melon",
        "price": 4000,
        "category": "Fruits"
        },
        {
        "name": "Watermelon",
        "price": 10000,
        "category": "Fruits"
        },
        {
        "name": "Banana",
        "price": 4000,
        "category": "Fruits"
        },
        {
        "name": "Carrot",
        "price": 2000,
        "category": "Vegetables"
        },
        {
        "name": "Broccoli",
        "price": 5200,
        "category": "Vegetables"
        },
        {
        "name": "Cucumber",
        "price": 3400,
        "category": "Vegetables"
        },
        {
        "name": "Potato",
        "price": 6500,
        "category": "Vegetables"
        },
        {
        "name": "Tomato",
        "price": 2200,
        "category": "Vegetables"
        },
        {
        "name": "Coffee",
        "price": 4300,
        "category": "Drink"
        },
        {
        "name": "Milk",
        "price": 4000,
        "category": "Drink"
        },
        {
        "name": "Tea",
        "price": 2700,
        "category": "Drink"
        }
    ]
}
```

- `/api/cart/add?product_name=<product_name>`

```json
{
    "name": "Apple",
    "price": 2000,
    "category": "Fruits"
}
```

dashboard:

```json
{
    "username": "aditira",
        "cart_items": [
        {
            "product_name": "Apple",
            "category": "Fruits",
            "price": 2000,
            "quantity": 2
        },
        {
            "product_name": "Banana",
            "category": "Fruits",
            "price": 4000,
            "quantity": 2
        }
    ],
    "total_price": 12000,
    "money_changes": -12000
}
```

- `/api/cart/clear`

success : kosong (response code 200)

dashboard: 
```json
{
    "username": "aditira",
    "cart_items": null,
    "total_price": 0,
    "money_changes": 0
}
```


- `/api/carts`

kalau di dashboard isi `cart_items` itu `null`

```json
{
    "cart_items": []
}
```

kalau tidak kosong, tampilkan list carts dari dashboard

```json
{
    "cart_items": [
        {
            "Category": "Fruits",
            "ProductName": "Apple",
            "Price": 2000,
            "Quantity": 1
        }
    ]
}
```

For simplicity, we only use HTTP GET here.
The API also doesn't support multiple concurrent sessions. We will learn about this later :)
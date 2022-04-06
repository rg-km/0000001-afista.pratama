package repository

import (
	"strconv"

	"github.com/ruang-guru/playground/backend/basic-golang/cashier-app/db"
)

type CartItemRepository struct {
	db db.DB // konektornya dari si db
}

func NewCartItemRepository(db db.DB) CartItemRepository {
	return CartItemRepository{db}
}

func (u *CartItemRepository) LoadOrCreate() ([]CartItem, error) {
	records, err := u.db.Load("cart_items")
	if err != nil {
		records = [][]string{
			{"category", "product_name", "price", "quantity"},
		}
		if err := u.db.Save("cart_items", records); err != nil {
			return nil, err
		}
	}

	result := make([]CartItem, 0)
	for i := 1; i < len(records); i++ {
		price, err := strconv.Atoi(records[i][2])
		if err != nil {
			return nil, err
		}

		qty, err := strconv.Atoi(records[i][3])
		if err != nil {
			return nil, err
		}

		cartItem := CartItem{
			Category:    records[i][0],
			ProductName: records[i][1],
			Price:       price,
			Quantity:    qty,
		}
		result = append(result, cartItem)
	}

	return result, nil
}

func (u *CartItemRepository) Save(cartItems []CartItem) error {
	records := [][]string{
		{"category", "product_name", "price", "quantity"},
	}
	for i := 0; i < len(cartItems); i++ {
		records = append(records, []string{
			cartItems[i].Category,
			cartItems[i].ProductName,
			strconv.Itoa(cartItems[i].Price),
			strconv.Itoa(cartItems[i].Quantity),
		})
	}
	return u.db.Save("cart_items", records)
}

func (u *CartItemRepository) SelectAll() ([]CartItem, error) {
	//return []CartItem{}, nil
	return u.LoadOrCreate()
}

func (u *CartItemRepository) Add(product Product) error {
	carts, err := u.LoadOrCreate()
	if err != nil {
		return err
	}

	flag := false

	for i := 0; i < len(carts); i++ {
		if carts[i].ProductName == product.ProductName {
			flag = true
			carts[i].Quantity++
			return u.Save(carts)
		}
	}

	if flag == false {
		carts = append(carts, CartItem{
			Category:    product.Category,
			ProductName: product.ProductName,
			Price:       product.Price,
			Quantity:    1,
		})
	}

	return u.Save(carts)
}

func (u *CartItemRepository) ResetCartItems() error {
	return nil // TODO: replace this
}

func (u *CartItemRepository) TotalPrice() (int, error) {
	// get total price

	// tinggal di total
	// 2 barang
	// a 1000 2 = 2000
	// b 5000 3 = 15000
	// 17000

	carts, err := u.LoadOrCreate()
	if err != nil {
		return 0, err
	}

	totalPrice := 0

	for _, cart := range carts {
		totalPrice += (cart.Quantity * cart.Price)
	}

	return totalPrice, nil
}

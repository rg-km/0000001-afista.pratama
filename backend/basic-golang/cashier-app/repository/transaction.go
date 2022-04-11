package repository

type TransactionRepository struct {
	cartItemRepository CartItemRepository
}

func NewTransactionRepository(cartItemRepository CartItemRepository) TransactionRepository {
	return TransactionRepository{cartItemRepository}
}

func (u *TransactionRepository) Pay(amount int) (int, error) {
	// TODO: replace this

	// selesaikan card dulu
	// total pricenya

	// kembalian = didapat dari amount - total 20.000 - 17.000 = 3.000

	total, err := u.cartItemRepository.TotalPrice()
	if err != nil {
		return 0, err
	}

	return amount - total, nil
}

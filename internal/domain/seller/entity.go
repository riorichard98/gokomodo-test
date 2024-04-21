package seller

// Seller represents the seller table in the database.
type Seller struct {
	ID           string `db:"id"`
	Email        string `db:"email"`
	Name         string `db:"name"`
	Password     string `db:"password"`
	AlamatPickup string `db:"alamat_pickup"`
}

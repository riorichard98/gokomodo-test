package buyer

type Buyer struct {
	ID               string `db:"id"`
	Email            string `db:"email"`
	Name             string `db:"name"`
	Password         string `db:"password"`
	AlamatPengiriman string `db:"alamat_pengiriman"`
}

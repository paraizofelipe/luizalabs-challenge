package repository

import (
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/paraizofelipe/luizalabs-challenge/product/domain"
	"github.com/paraizofelipe/luizalabs-challenge/storage"
)

type repository struct {
	storage storage.PostgresStorage
}

func NewPostgreRepository(db *sqlx.DB) ProductRepository {
	return &repository{
		storage: storage.NewPostgres(db),
	}
}

func (r repository) FindAll() (listProduct []domain.Product, err error) {
	statement := `
			SELECT
					id,
					title,
					image,
					price,
					brand,
					review_score,
					created_at,
					updated_at
			FROM
					products;
	`
	err = r.storage.Find(statement, &listProduct, nil)
	return
}

func (r repository) FindByID(id string) (product domain.Product, err error) {
	statement := `
			SELECT
					id,
					title,
					image,
					price,
					brand,
					review_score,
					created_at,
					updated_at
			FROM
					products
			WHERE
					id = $1;
	`
	err = r.storage.Find(statement, &product, nil)
	return
}

func (r repository) Add(product domain.Product) (err error) {
	statement := `
			INSERT INTO product (
				id,
				title,
				image,
				price,
				brand,
				review_score,
				created_at,
				updated_at
			) VALUES (
				$1,
				$2,
				$3,
				$4,
				$5
			);
	`
	err = r.storage.Exec(statement,
		uuid.New(),
		product.Title,
		product.Image,
		product.Price,
		product.Brand,
		product.ReviewScore,
		time.Now().UTC(),
		time.Now().UTC(),
	)
	return
}

func (r repository) Update(product domain.Product) (err error) {
	statement := `
			UPDATE
				product
			SET
				title = $1,
				image = $2,
				price = $3,
				brand = $4,
				review_score = $5,
			WHERE
				id = $6
        `
	return r.storage.Exec(
		statement,
		product.Title,
		product.Image,
		product.Price,
		product.Brand,
		product.ReviewScore,
		time.Now().UTC(),
		product.Id,
	)
}

func (r repository) RemoveByID(id string) (err error) {
	return r.storage.Exec(`DELETE FROM product WHERE email = $1`, id)
}

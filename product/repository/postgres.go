package repository

import (
	"database/sql"
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

func (r repository) ListByPage(page int) (listProduct []domain.Product, err error) {
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
			ORDER BY title
			OFFSET $1
			LIMIT 20;
	`
	err = r.storage.FindAll(statement, &listProduct, page)
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
	err = r.storage.Find(statement, &product, id)
	return
}

func (r repository) FindByTitleAndBrand(brand string, title string) (product domain.Product, err error) {
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
					brand = $1
					AND
					title = $2;
	`
	err = r.storage.Find(statement, &product, brand, title)
	if err == sql.ErrNoRows {
		err = nil
	}
	return
}

func (r repository) Add(product domain.Product) (err error) {
	statement := `
			INSERT INTO products (
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
				$5,
				$6,
				$7,
				$8
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
				products
			SET
				title = $1,
				image = $2,
				price = $3,
				brand = $4,
				review_score = $5,
				updated_at = $6
			WHERE
				id = $7
        `
	return r.storage.Exec(
		statement,
		product.Title,
		product.Image,
		product.Price,
		product.Brand,
		product.ReviewScore,
		time.Now().UTC(),
		product.ID,
	)
}

func (r repository) RemoveByID(id string) (err error) {
	return r.storage.Exec(`DELETE FROM products WHERE id = $1`, id)
}

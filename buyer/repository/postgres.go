package repository

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/paraizofelipe/luizalabs-challenge/buyer/domain"
	"github.com/paraizofelipe/luizalabs-challenge/storage"
)

type repository struct {
	storage storage.PostgresStorage
}

func NewPostgreRepository(db *sqlx.DB) BuyerRepository {
	return &repository{
		storage: storage.NewPostgres(db),
	}
}

func (r repository) FindAll() (listBuyer []domain.Buyer, err error) {
	statement := `
			SELECT
				id,
				name,
				email,
				created_at,
				updated_at
			FROM
				buyers;
	`
	err = r.storage.Find(statement, &listBuyer, nil)
	return
}

func (r repository) FindByID(id string) (buyer domain.Buyer, err error) {
	statement := `
			SELECT
				id,
				name,
				email,
				created_at,
				updated_at
			FROM
				buyers
			WHERE
				id = $1;
	`
	err = r.storage.Find(statement, &buyer, id)
	if err == sql.ErrNoRows {
		err = nil
	}
	return
}

func (r repository) FindByEmail(email string) (buyer domain.Buyer, err error) {
	statement := `
			SELECT
				id,
				name,
				email,
				created_at,
				updated_at
			FROM
				buyers
			WHERE
				email = $1;
	`
	err = r.storage.Find(statement, &buyer, email)
	if err == sql.ErrNoRows {
		err = nil
	}
	return
}

func (r repository) Add(buyer domain.Buyer) (err error) {
	statement := `
			INSERT INTO buyers (
				id,
				name,
				email,
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
		buyer.Name,
		buyer.Email,
		time.Now().UTC(),
		time.Now().UTC(),
	)
	return
}

func (r repository) AddFavoriteProduct(id string, productID string) (err error) {
	statement := `
			INSERT INTO product_to_buyer (
				buyer_id
				product_id,
				created_at,
				updated_at
			) VALUES (
				$1,
				$2,
				$3,
				$4
			);
	`
	err = r.storage.Exec(statement,
		id,
		productID,
		time.Now().UTC(),
		time.Now().UTC(),
	)
	return
}

func (r repository) Update(buyer domain.Buyer) (err error) {
	statement := `
			UPDATE
				buyers
			SET
				name = $1,
				email = $2,
				updated_at = $3
			WHERE
				id = $4
	`
	return r.storage.Exec(
		statement,
		buyer.Name,
		buyer.Email,
		time.Now().UTC(),
		buyer.ID,
	)
}

func (r repository) RemoveByID(id string) (err error) {
	return r.storage.Exec(`DELETE FROM buyers WHERE id = $1`, id)
}

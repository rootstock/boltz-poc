package storage

import (
	"database/sql"

	"github.com/jmoiron/sqlx"

	log "github.com/sirupsen/logrus"
)

type Payment struct {
	PreimageHash string `db:"PreimageHash"`
	Preimage     string `db:"Preimage"`
	Invoice      string `db:"Invoice"`
	Amount       int    `db:"Amount"`
	ProductId    int    `db:"ProductId"`
	Status       string `db:"Status"`
	Tx           string `db:"Tx"`
}

type PaymentRequest struct {
	ProductId int `json:"productId"`
}

const createPaymentTable = `
CREATE TABLE IF NOT EXISTS payment (
	PreimageHash TEXT PRIMARY KEY,
	Preimage TEXT,
	Invoice TEXT,
	Amount NUMBER,
	ProductId NUMBER,
	Status TEXT,
	Tx TEXT
)
`

const insertPayment = `
INSERT INTO payment (
    PreimageHash,
	Preimage,
	Invoice,
	Amount,
	ProductId,
	Status,
	Tx
)
VALUES (
    :PreimageHash,
	:Preimage,
	:Invoice,
	:Amount,
	:ProductId,
	:Status,
	:Tx
)
`

const updatePayments = `
UPDATE payment SET  Status = :Status,
					Tx = :Tx
Where PreimageHash = :PreimageHash`

const getPayments = `
SELECT *
FROM payment`

const getPayment = `
SELECT *
FROM payment
WHERE PreimageHash = ?
LIMIT 1`

func (db *DB) SavePayment(payment *Payment) error {
	exist, err := db.GetPayment(payment.PreimageHash)
	if err != nil {
		log.Error("Error saving payment", err)
		return err
	}
	var queryString string
	if (exist != Payment{}) {
		log.Debug("updating payment{", payment.PreimageHash, "}", ": ", payment)
		queryString = updatePayments
	} else {
		log.Debug("inserting payment{", payment.PreimageHash, "}", ": ", payment)
		queryString = insertPayment
	}
	query, args, _ := sqlx.Named(queryString, payment)
	result, err := db.db.Exec(query, args...)
	if err != nil {
		log.Error("Error saving payment", err)
		return err
	}
	log.Debug("success saving payment", result)
	return nil
}

func (db *DB) GetPayments() ([]*Payment, error) {
	log.Debug("retrieving payments... ")
	var payments []*Payment
	rows, err := db.db.Queryx(getPayments)
	if err == sql.ErrNoRows {
		return payments, nil
	}
	defer rows.Close()
	for rows.Next() {
		var payment Payment

		if err := rows.StructScan(&payment); err != nil {
			return nil, err
		}

		payments = append(payments, &payment)
	}
	return payments, nil
}

func (db *DB) GetPayment(preimageHash string) (result Payment, err error) {
	log.Debugf("retrieving payment: %s ", preimageHash)
	rows, err := db.db.Queryx(getPayment, preimageHash)
	if err != nil {
		log.Fatalf("Error retrieving payment: %v", err)
		return
	}
	defer rows.Close()
	if rows.Next() {
		if err = rows.StructScan(&result); err != nil {
			log.Fatalf("Error mapping payment: %v", err)
			return
		}
	}
	return

}

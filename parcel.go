package main

import (
	"database/sql"
)

type ParcelStore struct {
	db *sql.DB
}

func NewParcelStore(db *sql.DB) ParcelStore {
	return ParcelStore{db: db}
}

func (s ParcelStore) Add(p Parcel) (int, error) {
	result, err := s.db.Exec(
		`INSERT INTO parcel (client, status, address, created_at) VALUES (?, ?, ?, ?)`, // реализуйте добавление строки в таблицу parcel, используйте данные из переменной p
		p.Client, p.Status, p.Address, p.CreatedAt,
	)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
} // верните идентификатор последней добавленной записи

func (s ParcelStore) Get(number int) (Parcel, error) {
	row := s.db.QueryRow(`SELECT * FROM parcel WHERE number = ?`, number)
	p := Parcel{}
	err := row.Scan(&p.Number, &p.Client, &p.Status, &p.Address, &p.CreatedAt)
	return p, err
}

func (s ParcelStore) GetByClient(client int) ([]Parcel, error) {
	rows, err := s.db.Query(`SELECT * FROM parcel WHERE client = ?`, client)
	if err != nil {
		return nil, err
	}
	defer rows.Close() // реализуйте чтение строк из таблицы parcel по заданному client
	// здесь из таблицы может вернуться несколько строк

	// заполните срез Parcel данными из таблицы
	var res []Parcel
	for rows.Next() {
		p := Parcel{}
		err := rows.Scan(&p.Number, &p.Client, &p.Status, &p.Address, &p.CreatedAt)
		if err != nil {
			return nil, err
		}
		res = append(res, p)
	}

	return res, nil
}

func (s ParcelStore) SetStatus(number int, status string) error {
	_, err := s.db.Exec(`UPDATE parcel SET status = ? WHERE number = ?`, status, number)
	return err // реализуйте обновление статуса в таблице parcel

}

func (s ParcelStore) SetAddress(number int, address string) error {
	_, err := s.db.Exec(`UPDATE parcel SET address = ? WHERE number = ? AND status = ?`, address, number, ParcelStatusRegistered)
	return err // реализуйте обновление адреса в таблице parcel
	// менять адрес можно только если значение статуса registered

}

func (s ParcelStore) Delete(number int) error {
	_, err := s.db.Exec(`DELETE FROM parcel WHERE number = ? AND status = ?`, number, ParcelStatusRegistered)
	return err // реализуйте удаление строки из таблицы parcel
	// удалять строку можно только если значение статуса registered

}

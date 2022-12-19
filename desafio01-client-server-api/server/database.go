package server

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func SaveCotacao(cotacao Cotacao) (err error) {
	// db, err := sql.Open("sqlite3", "./db/cotacao.db")
	// if err != nil {
	// 	return err
	// }
	// defer db.Close()
	// err = persist(db, &cotacao)
	// if err != nil {
	// 	return err
	// }
	return nil
}

func persist(db *sql.DB, cotacao *Cotacao) (err error) {
	stmt, err := db.Prepare("insert into cotacoes(id, code, codein, name, high, low, var_bid, pct_change, bid, ask, timestamp, create_date) values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(cotacao.Id, cotacao.Code, cotacao.CodeIn, cotacao.Name, cotacao.High, cotacao.Low, cotacao.VarBid, cotacao.PctChange, cotacao.Bid, cotacao.Ask, cotacao.Timestamp, cotacao.CreateDate)
	if err != nil {
		return err
	}
	return nil
}
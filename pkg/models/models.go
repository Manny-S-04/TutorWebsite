package models

import (
	"database/sql"
	"errors"
)

var ErrNoRecord = errors.New("No matching record found")

type ReviewModel struct{
    DB *sql.DB
}

type Review struct{
    ID int
    Name string
    Content string
    Stars string
} 

func (m *ReviewModel) Insert(name,content,stars string) (int, error) {
    stmt := `INSERT INTO user_reviews (name, content, stars) 
		VALUES(?,?,?)`

	result, err := m.DB.Exec(stmt, name, content, stars)
	if err != nil {
		return 0, err
}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (m *ReviewModel) Get(id int) (*Review, error) {
    review := &Review{}

	err := m.DB.QueryRow(`SELECT id, name, content, stars FROM user_reviews
	WHERE id = ?`, id).Scan(&review.ID, &review.Name, &review.Content, &review.Stars)

	if err == sql.ErrNoRows {
		return nil, ErrNoRecord
	} else if err != nil {
		return nil, err
	}
	return review, nil
}

func (m *ReviewModel) GetAll() ([]*Review, error){
    stmt := `SELECT * FROM user_reviews ORDER BY stars DESC`
    rows, err := m.DB.Query(stmt)
    if err != nil{
        return nil, err
    }
    defer rows.Close()

    reviews := []*Review{}

    for rows.Next(){
        r := &Review{}
        err = rows.Scan(&r.ID, &r.Name, &r.Content, &r.Stars)

        if err != nil{
            return nil, err
        }
        reviews = append(reviews, r)
    }
    if err = rows.Err(); err != nil{
        return nil, err
    }
    return reviews, nil
}

func (m *ReviewModel) Latest() ([]*Review, error) {
	stmt := `SELECT id, name, content, stars FROM user_reviews
	ORDER BY id DESC`
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	reviews := []*Review{}

	for rows.Next() {
		r := &Review{}
		err = rows.Scan(&r.ID, &r.Name, &r.Content, &r.Stars)

		if err != nil {
			return nil, err
		}
        reviews = append(reviews, r)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return reviews, nil
}

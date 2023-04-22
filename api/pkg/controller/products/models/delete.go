package products

func (p *ProductRepository) DeleteBy(id int64) (int64, error) {
	res, err := p.Conn.Exec("DELETE FROM product WHERE id=$1", id)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

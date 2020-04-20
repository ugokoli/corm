package corm

func (d *DB) runMigrateQuery(query string) error {
	return d.session.Query(query).Exec()
}

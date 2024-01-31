package postgres

import (
	"bw-erp/helper"
	"bw-erp/models"
)

func (stg *Postgres) GetWorkVolumeList(companyID string) ([]models.WorkVolume, error) {
	var arr []interface{}
	var workVolumes []models.WorkVolume
	params := make(map[string]interface{})
	query := `SELECT 
		sum(width*height) meter_square, 
		washed_at::date, 
		type 
		FROM order_items`

	filter := " WHERE true"
	order := " ORDER BY washed_at"
	arrangement := " DESC"
	group := " group by washed_at::date, type"

	q := query + filter + group + order + arrangement

	q, arr = helper.ReplaceQueryParams(q, params)
	rows, err := stg.db.Query(q, arr...)
	if err != nil {
		return workVolumes, err
	}
	defer rows.Close()

	for rows.Next() {
		var workVolume models.WorkVolume
		err = rows.Scan(
			&workVolume.MeterSquare,
			&workVolume.WashedAt,
			&workVolume.Type)
		if err != nil {
			return nil, err
		}
		workVolumes = append(workVolumes, workVolume)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return workVolumes, nil
}
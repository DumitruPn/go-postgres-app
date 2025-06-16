package notification

import (
	"database/sql"
	"maps"
	"slices"
)

func Insert(db *sql.DB, notification CreateNotificationRequest) (int, error) {
	var id int
	err := db.QueryRow(`
		INSERT INTO data.notifications (value)
		VALUES ($1) RETURNING id
	`, notification.Value).Scan(&id)

	return id, err
}

func GetAll(db *sql.DB) ([]Notification, error) {
	rows, err := db.Query(`SELECT id, value FROM data.notifications`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	notificationsMap := make(map[int]Notification)

	for rows.Next() {
		var dto Dto
		if err := rows.Scan(&dto.ID, &dto.Value); err != nil {
			return nil, err
		}
		notificationsMap[dto.ID] = Notification{
			ID:    dto.ID,
			Value: dto.Value,
		}
	}

	return slices.Collect(maps.Values(notificationsMap)), nil
}

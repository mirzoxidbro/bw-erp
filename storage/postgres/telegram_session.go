package postgres

import (
	"bw-erp/models"
	"fmt"
)

func (stg *Postgres) CreateTelegramSessionModel(entity models.TelegramSessionModel) error {
	_, err := stg.db.Exec(`INSERT INTO telegram_sessions(
		order_id,
		bot_id,
		chat_id
	) VALUES (
		$1,
		$2,
		$3
	)`,
		entity.OrderID,
		entity.BotID,
		entity.ChatID,
	)

	if err != nil {
		fmt.Print("\n session create error", err)
		return err
	}

	return nil
}

func (stg *Postgres) DeleteTelegramSession(ID int) (rowsAffected int64, err error) {
	query := `DELETE FROM "telegram_sessions" WHERE id = $1`

	result, err := stg.db.Exec(query, ID)
	if err != nil {
		return 0, err
	}

	rowsAffected, err = result.RowsAffected()

	if err != nil {
		return 0, err
	}

	return rowsAffected, err
}

func (stg *Postgres) GetTelegramSessionByChatIDBotID(ChatID int64, BotID int64) (models.TelegramSessionModel, error) {
	var session models.TelegramSessionModel
	err := stg.db.QueryRow(`select id, bot_id, chat_id, order_id from telegram_sessions where bot_id = $1 and chat_id = $2`, BotID, ChatID).Scan(
		&session.ID,
		&session.BotID,
		&session.ChatID,
		&session.OrderID,
	)
	if err != nil {
		return session, err
	}

	return session, nil
}
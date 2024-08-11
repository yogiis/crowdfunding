package transaction

import "crowdfunding/user"

type GetCampaignTransactionsInput struct {
	ID   int `uri:"id" binding:"required"`
	User user.User
}

type CreateTransactionInput struct {
	Amount     int `uri:"amount" binding:"required"`
	CampaignID int `uri:"campaign_id" binding:"required"`
	User       user.User
}

package transaction

type service struct {
	repository Repository
}

type Service interface {
	GetTransactionByCampaignByID(input GetCampaignTransactionsInput) ([]Transaction, error)
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetTransactionByCampaignByID(input GetCampaignTransactionsInput) ([]Transaction, error) {
	transaction, err := s.repository.GetCampaignById(input.ID)
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

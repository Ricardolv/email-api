package campaign

type Repository interface {
	Create(campaign *Campaign) error
	FindAll() []Campaign
}

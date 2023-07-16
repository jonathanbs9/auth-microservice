package infra

import (
	"auth-microservice/domain"
	"context"

	"gorm.io/gorm"
)

type postgresStorage struct {
	*gorm.DB
}

type admin struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     string `gorm:"unique"`
	Password  string
	AuthToken domain.AuthToken `gorm:"embedded;embeddedPrefix:auth_token_"`
}

func toAdmin(a *admin) domain.Admin {
	return domain.Admin{
		ID:           domain.AdminID(a.ID),
		FirstName:    a.FirstName,
		Lastname:     a.LastName,
		Email:        a.Email,
		PasswordHash: a.Password,
		AuthToken:    a.AuthToken,
	}
}

func fromAdmin(a domain.Admin) *admin {
	return &admin{
		FirstName: a.FirstName,
		LastName:  a.Lastname,
		Email:     a.Email,
		Password:  a.PasswordHash,
		AuthToken: a.AuthToken,
	}
}

func NewPostgresStorage(db *gorm.DB) (domain.Storage, error) {
	if err := db.AutoMigrate(&admin{}); err != nil {
		return nil, err
	}

	return &postgresStorage{db}, nil
}

func (s postgresStorage) Save(ctx context.Context, a domain.Admin) error {
	row := fromAdmin(a)
	if a.ID == domain.AdminID(0) {
		return s.Create(row).Error
	}

	return s.WithContext(ctx).Model(&admin{}).Where("id =?", a.ID).
		Updates(row).
		Error
}

func (s postgresStorage) FindAll(ctx context.Context) ([]domain.Admin, error) {
	var rows []admin
	var admins []domain.Admin

	tx := s.WithContext(ctx).Find(&rows)
	if tx.Error != nil {
		return admins, tx.Error
	}

	for _, row := range rows {
		admins = append(admins, toAdmin(&row))
	}

	return admins, nil
}

func (s postgresStorage) FindByID(ctx context.Context, id domain.AdminID) (domain.Admin, error) {
	var row admin
	tx := s.First(ctx, &row, id)
	if tx.Error != nil {
		return toAdmin(&row), tx.Error
	}

	return toAdmin(&row), nil
}

func (s postgresStorage) FindByName(ctx context.Context, name string) (domain.Admin, error) {
	var row admin

	tx := s.WithContext(ctx).First(&row, "first_name = ?", name)
	if tx.Error != nil {
		return toAdmin(&row), tx.Error
	}

	return toAdmin(&row), nil
}

func (s postgresStorage) FindByEmail(ctx context.Context, email string) (domain.Admin, error) {
	var row admin
	tx := s.WithContext(ctx).First(&row, "email = ?", email)
	if tx.Error != nil {
		return toAdmin(&row), tx.Error
	}

	return toAdmin(&row), nil
}

func (s postgresStorage) FindByAuthTokenID(ctx context.Context, id domain.AuthTokenID) (domain.Admin, error) {
	var row admin

	tx := s.WithContext(ctx).First(&row, "auth_token_id = ?", id)
	if tx.Error != nil {
		return toAdmin(&row), tx.Error
	}

	return toAdmin(&row), nil
}

func (s postgresStorage) DeleteByID(ctx context.Context, id domain.AdminID) error {
	return s.WithContext(ctx).Delete(&admin{}, id).Error
}

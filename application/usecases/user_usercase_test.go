package usecases_test

import (
	"github.com/bxcodec/faker/v3"
	"github.com/acampagnaro/go-grpc/application/repositories"
	"github.com/acampagnaro/go-grpc/application/usecases"
	"github.com/acampagnaro/go-grpc/domain"
	"github.com/acampagnaro/go-grpc/framework/utils"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/require"
	"testing"
)

var db *gorm.DB

func TestLogin_Auth(t *testing.T) {
	db = utils.ConnectDB("test")

	repo := repositories.UserRepositoryDb{Db: db}
	email := faker.Email()
	password := faker.Password()

	newUser, err := domain.NewUser(faker.Name(), email, password)
	repo.Insert(newUser)

	userUserCase := usecases.UserUseCase{UserRepository: repo}
	_, err = userUserCase.Auth(email, password)
	require.Nil(t, err)
}

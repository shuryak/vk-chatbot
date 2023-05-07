package usecase

import (
	"fmt"
	"github.com/shuryak/vk-chatbot/internal/models"
	"github.com/shuryak/vk-chatbot/pkg/vkapi"
	"github.com/shuryak/vk-chatbot/pkg/vkapi/params"
	"strconv"
	"time"
)

type VKUserManager struct {
	*vkapi.VKAPI
}

func NewVKUserManager(vkapi *vkapi.VKAPI) *VKUserManager {
	return &VKUserManager{vkapi}
}

// Check for implementation
var _ UserManager = (*VKUserManager)(nil)

func (vk *VKUserManager) GetUser(ID int) (*models.User, error) {
	builder := params.NewUsersGetBuilder().
		UserIDs([]string{strconv.Itoa(ID)}).
		Fields([]string{"photo_id, city, bdate"})

	users, err := vk.UsersGet(builder.Params)
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, fmt.Errorf("user %d not recieved", ID)
	}

	user := models.User{
		ID:        users[0].ID,
		PhotoID:   users[0].PhotoId,
		Name:      users[0].FirstName,
		BirthDate: nil,
		City:      users[0].City.Title,
	}

	t, err := time.Parse("2.1.2006", users[0].Bdate)
	if err == nil {
		user.BirthDate = &t
	}

	return &user, nil
}

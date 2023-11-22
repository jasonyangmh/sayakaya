package cron

import (
	"context"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/jasonyangmh/sayakaya/config"
	"github.com/jasonyangmh/sayakaya/model"
	"github.com/jasonyangmh/sayakaya/shared"
	"github.com/jasonyangmh/sayakaya/usecase"
)

type Cron struct {
	config       *config.Config
	userUsecase  usecase.UserUsecase
	promoUsecase usecase.PromoUsecase
}

func New(cfg *config.Config, uu usecase.UserUsecase, pu usecase.PromoUsecase) *Cron {
	return &Cron{
		config:       cfg,
		userUsecase:  uu,
		promoUsecase: pu,
	}
}

func (c *Cron) Run() {
	s := gocron.NewScheduler(time.UTC)

	s.Every(1).Day().At(c.config.Schedule).Do(func() {
		users, _ := c.userUsecase.FindUsers(context.Background())

		for _, user := range users {
			if isBirthday := c.userUsecase.CheckUserBirthday(context.Background(), &user); isBirthday {
				promo, err := c.promoUsecase.FindLatestPromo(context.Background())
				if err != nil {
					return
				}

				code := shared.GenerateCode(c.config.CodeLen)

				userPromo := &model.UserPromo{
					Code:    code,
					UserID:  user.Model.ID,
					PromoID: promo.Model.ID,
				}

				userPromo, err = c.promoUsecase.CreatePromoForUser(context.Background(), userPromo)
				if err != nil {
					return
				}

				userPromo, err = c.promoUsecase.FindPromoByCode(context.Background(), userPromo)
				if err != nil {
					return
				}

				shared.SendMail(c.config, userPromo)
			}

		}
	})

	s.StartAsync()
}

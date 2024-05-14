package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	FirstName     string             `bson:"firstName" validate:"alpha"`
	LastName      string             `bson:"lastName" validate:"alpha"`
	Email         string             `bson:"email" validate:"required,email"`
	Password      string             `bson:"password" validate:"required"`
	RememberToken string             `bson:"rememberToken,omitempty"`
	Phone         string             `bson:"phone" validate:"e164"`
	Address       string             `bson:"address"`
	Birthday      time.Time          `bson:"birthday,omitempty" validate:"omitempty,datetime"`
	Gender        string             `bson:"gender,omitempty" validate:"omitempty,oneof=male female other"`
	Nationality   string             `bson:"nationality,omitempty"`
	Bio           string             `bson:"bio,omitempty"`
	CreatedAt     time.Time          `bson:"createdAt"`
	UpdatedAt     time.Time          `bson:"updatedAt"`
	DeletedAt     *time.Time         `bson:"deletedAt,omitempty"`
	// RoleID        primitive.ObjectID `bson:"roleId" validate:"required"` // currently not used
}

// SoftDelete sets the DeletedAt field to the current time to mark an entry as deleted.
func (u *User) SoftDelete() {
	now := time.Now()
	u.DeletedAt = &now
}

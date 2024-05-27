package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	FirstName      string             `bson:"firstName" json:"firstName"`
	LastName       string             `bson:"lastName" json:"lastName"`
	Email          string             `bson:"email" json:"email"`
	Password       string             `bson:"password" json:"password"`
	RememberToken  string             `bson:"rememberToken,omitempty" json:"rememberToken,omitempty"`
	Phone          string             `bson:"phone" json:"phone"`
	Birthday       time.Time          `bson:"birthday,omitempty" json:"birthday,omitempty"`
	Gender         string             `bson:"gender,omitempty" json:"gender,omitempty"`
	Nationality    string             `bson:"nationality,omitempty" json:"nationality,omitempty"`
	Bio            string             `bson:"bio,omitempty" json:"bio,omitempty"`
	ProfilePicture string             `bson:"profilePicture,omitempty" json:"profilePicture,omitempty"`
	Address        Address            `bson:"address,omitempty" json:"address,omitempty"`

	MusicalInfoID primitive.ObjectID `bson:"musical_info_id,omitempty" json:"musical_info_id,omitempty"` // Reference to MusicalInformation
	CoursePrefs   CoursePreferences  `bson:"course_preferences" json:"course_preferences"`               // Using embedded approach for now, will be change when Instructor model is implemented

	CreatedAt time.Time  `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time  `bson:"updatedAt" json:"updatedAt"`
	DeletedAt *time.Time `bson:"deletedAt,omitempty" json:"deletedAt,omitempty"`
	// RoleID        primitive.ObjectID `bson:"roleId" validate:"required"` // currently not used
}

type Address struct {
	Street     string `bson:"street,omitempty" json:"street,omitempty"`
	City       string `bson:"city" json:"city"`
	State      string `bson:"state" json:"state"`
	PostalCode string `bson:"postalCode" json:"postalCode"`
	Country    string `bson:"country" json:"country"`
}

// SoftDelete sets the DeletedAt field to the current time to mark an entry as deleted.
func (u *User) SoftDelete() {
	now := time.Now()
	u.DeletedAt = &now
}

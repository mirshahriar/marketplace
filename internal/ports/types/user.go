package types

// User is the DB model for users table
type User struct {
	// ID is the primary key
	ID uint `json:"id" gorm:"primary_key;auto_increment:true"`
	// Username is the username of the user. Must be unique.
	Username string `json:"username" gorm:"unique;NOT NULL"`
	// Password is the password of the user.
	// Store the hashed password.
	Password string `json:"password" gorm:"NOT NULL"`
	// Email is the email of the user. Must be unique.
	// +optional
	Email string `json:"email" gorm:"unique"`
}

// TableName returns the table name for the User model
func (User) TableName() string {
	return "users"
}

// LoggedInUser is the response with user id after verifying the token
type LoggedInUser struct {
	ID uint `json:"id"`
}

// Token is the DB model for tokens table
type Token struct {
	// UserID is the foreign key of users table
	UserID uint `json:"user_id" gorm:"NOT NULL"`
	User   User `json:"user" gorm:"foreignKey:UserID;references:ID"`
	// Token is the token of the user
	// Store the hashed token.
	Token string `json:"token" gorm:"NOT NULL"`
}

// TableName returns the table name for the Token model
func (Token) TableName() string {
	return "tokens"
}

// the database model for a user account
package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

// Id is the primary key which is auto-incremented
// Password contains hexadecimal representation of PBKDF2 hash of the plaintext password
// PBKDF2 is simply a way to securely store a user’s password in the database
// Reg_key contains the UUID string that is used for account verification (via email)
// UUID is a unique identifier that can be used to ensure the identity of the user for verification purposes
// Reg_date is the timestamp indicationg the time of registration
// orm:"Unique" means Email must be a unique key
// Reg_date will be automatically assigned the date and time of database insertion
type AuthUser struct {
	Id       int
	First    string
	Last     string
	Email    string `orm:"unique"`
	Password string
	Reg_key  string
	Reg_date time.Time `orm:"auto_now_add;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(AuthUser))
}

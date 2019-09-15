package jwt

import (
	"time"

	"github.com/go-chi/jwtauth"
	"github.com/go-pg/pg/orm"
)

// Token holds refresh jwt information.
type Token struct {
	ID        int       `json:"id,omitempty"`
	CreatedDate time.Time `json:"createdDate,omitempty"`
	ModifiedDate time.Time `json:"modifiedDate,omitempty"`
	AccountID int       `json:"-"`

	Token      string    `json:"-"`
	Expiry     time.Time `json:"-"`
	Mobile     bool      `sql:",notnull" json:"mobile"`
	Identifier string    `json:"identifier,omitempty"`
}

// BeforeInsert hook executed before database insert operation.
func (t *Token) BeforeInsert(db orm.DB) error {
	now := time.Now()
	if t.CreatedDate.IsZero() {
		t.CreatedDate = now
		t.ModifiedDate = now
	}
	return nil
}

// BeforeUpdate hook executed before database update operation.
func (t *Token) BeforeUpdate(db orm.DB) error {
	t.ModifiedDate = time.Now()
	return nil
}

// Claims returns the token claims to be signed
func (t *Token) Claims() jwtauth.Claims {
	return jwtauth.Claims{
		"id":    t.ID,
		"token": t.Token,
	}
}

package responses

import (
	"time"

	pjson "github.com/increase/increase-go/core/json"
	"github.com/increase/increase-go/pagination"
)

type Card struct {
	// The card identifier.
	ID string `json:"id,required"`
	// The identifier for the account this card belongs to.
	AccountID string `json:"account_id,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Card was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The card's description for display purposes.
	Description string `json:"description,required,nullable"`
	// The last 4 digits of the Card's Primary Account Number.
	Last4 string `json:"last4,required"`
	// The month the card expires in M format (e.g., August is 8).
	ExpirationMonth int64 `json:"expiration_month,required"`
	// The year the card expires in YYYY format (e.g., 2025).
	ExpirationYear int64 `json:"expiration_year,required"`
	// This indicates if payments can be made with the card.
	Status CardStatus `json:"status,required"`
	// The Card's billing address.
	BillingAddress CardBillingAddress `json:"billing_address,required"`
	// The contact information used in the two-factor steps for digital wallet card
	// creation. At least one field must be present to complete the digital wallet
	// steps.
	DigitalWallet CardDigitalWallet `json:"digital_wallet,required,nullable"`
	// A constant representing the object's type. For this resource it will always be
	// `card`.
	Type CardType `json:"type,required"`
	JSON CardJSON
}

type CardJSON struct {
	ID              pjson.Metadata
	AccountID       pjson.Metadata
	CreatedAt       pjson.Metadata
	Description     pjson.Metadata
	Last4           pjson.Metadata
	ExpirationMonth pjson.Metadata
	ExpirationYear  pjson.Metadata
	Status          pjson.Metadata
	BillingAddress  pjson.Metadata
	DigitalWallet   pjson.Metadata
	Type            pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into Card using the internal pjson
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *Card) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type CardStatus string

const (
	CardStatusActive   CardStatus = "active"
	CardStatusDisabled CardStatus = "disabled"
	CardStatusCanceled CardStatus = "canceled"
)

type CardBillingAddress struct {
	// The first line of the billing address.
	Line1 string `json:"line1,required,nullable"`
	// The second line of the billing address.
	Line2 string `json:"line2,required,nullable"`
	// The city of the billing address.
	City string `json:"city,required,nullable"`
	// The US state of the billing address.
	State string `json:"state,required,nullable"`
	// The postal code of the billing address.
	PostalCode string `json:"postal_code,required,nullable"`
	JSON       CardBillingAddressJSON
}

type CardBillingAddressJSON struct {
	Line1      pjson.Metadata
	Line2      pjson.Metadata
	City       pjson.Metadata
	State      pjson.Metadata
	PostalCode pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into CardBillingAddress using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *CardBillingAddress) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type CardDigitalWallet struct {
	// An email address that can be used to verify the cardholder via one-time passcode
	// over email.
	Email string `json:"email,required,nullable"`
	// A phone number that can be used to verify the cardholder via one-time passcode
	// over SMS.
	Phone string `json:"phone,required,nullable"`
	// The card profile assigned to this digital card. Card profiles may also be
	// assigned at the program level.
	CardProfileID string `json:"card_profile_id,required,nullable"`
	JSON          CardDigitalWalletJSON
}

type CardDigitalWalletJSON struct {
	Email         pjson.Metadata
	Phone         pjson.Metadata
	CardProfileID pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into CardDigitalWallet using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *CardDigitalWallet) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type CardType string

const (
	CardTypeCard CardType = "card"
)

type CardDetails struct {
	// The identifier for the Card for which sensitive details have been returned.
	CardID string `json:"card_id,required"`
	// The card number.
	PrimaryAccountNumber string `json:"primary_account_number,required"`
	// The month the card expires in M format (e.g., August is 8).
	ExpirationMonth int64 `json:"expiration_month,required"`
	// The year the card expires in YYYY format (e.g., 2025).
	ExpirationYear int64 `json:"expiration_year,required"`
	// The three-digit verification code for the card. It's also known as the Card
	// Verification Code (CVC), the Card Verification Value (CVV), or the Card
	// Identification (CID).
	VerificationCode string `json:"verification_code,required"`
	// A constant representing the object's type. For this resource it will always be
	// `card_details`.
	Type CardDetailsType `json:"type,required"`
	JSON CardDetailsJSON
}

type CardDetailsJSON struct {
	CardID               pjson.Metadata
	PrimaryAccountNumber pjson.Metadata
	ExpirationMonth      pjson.Metadata
	ExpirationYear       pjson.Metadata
	VerificationCode     pjson.Metadata
	Type                 pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into CardDetails using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *CardDetails) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type CardDetailsType string

const (
	CardDetailsTypeCardDetails CardDetailsType = "card_details"
)

type CardList struct {
	// The contents of the list.
	Data []Card `json:"data,required"`
	// A pointer to a place in the list.
	NextCursor string `json:"next_cursor,required,nullable"`
	JSON       CardListJSON
}

type CardListJSON struct {
	Data       pjson.Metadata
	NextCursor pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into CardList using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *CardList) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type CardsPage struct {
	*pagination.Page[Card]
}

func (r *CardsPage) Card() *Card {
	return r.Current()
}

func (r *CardsPage) NextPage() (*CardsPage, error) {
	if page, err := r.Page.NextPage(); err != nil {
		return nil, err
	} else {
		return &CardsPage{page}, nil
	}
}
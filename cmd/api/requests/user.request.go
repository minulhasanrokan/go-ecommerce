package requests

type RegisterUserRequest struct {
	FirstName string `json:"first_name" validate:"required,min=2,max=200"`
	LastName  string `json:"last_name" validate:"required,min=2,max=200"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,max=20,min=6"`
	Mobile    string `json:"mobile" validate:"required,min=11,max=20"`
}

type LoginUserRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,max=20,min=6"`
}

type VerificationCodeRequest struct {
	Code string `json:"code" validate:"required,max=200,min=1"`
}

type ProfileInputRequest struct {
	FirstName    string              `json:"first_name" validate:"required,min=2,max=200"`
	LastName     string              `json:"last_name" validate:"required,min=2,max=200"`
	AddressInput AddressInputRequest `json:"address"`
}

type AddressInputRequest struct {
	AddressLine1 string `json:"address_line1" validate:"required,min=2,max=200"`
	AddressLine2 string `json:"address_line2" validate:"required,min=2,max=200"`
	City         string `json:"city" validate:"required,min=2,max=200"`
	PostCode     uint   `json:"post_code" validate:"required,min=2,max=9999"`
	Country      string `json:"country" validate:"required,min=2,max=200"`
}

type ProfileUpdateRequest struct {
	FirstName    string               `json:"first_name" validate:"required,min=2,max=200"`
	LastName     string               `json:"last_name" validate:"required,min=2,max=200"`
	AddressInput AddressUpdateRequest `json:"address"`
}

type AddressUpdateRequest struct {
	AddressLine1 string `json:"address_line1" validate:"required,min=2,max=200"`
	AddressLine2 string `json:"address_line2" validate:"required,min=2,max=200"`
	City         string `json:"city" validate:"required,min=2,max=200"`
	PostCode     uint   `json:"post_code" validate:"required,min=2,max=9999"`
	Country      string `json:"country" validate:"required,min=2,max=200"`
	AddressId    uint   `json:"address_id" validate:"required"`
}

type SellerInputRequest struct {
	FirstName         string `json:"first_name" validate:"required,min=2,max=200"`
	LastName          string `json:"last_name" validate:"required,min=2,max=200"`
	PhoneNumber       string `json:"phone_number" validate:"required,min=11,max=20"`
	BankAccountNumber string `json:"bankAccountNumber" validate:"required,min=11,max=20"`
	SwiftCode         string `json:"swiftCode" validate:"required,min=1,max=20"`
	PaymentType       string `json:"paymentType" validate:"required,min=1,max=20"`
}

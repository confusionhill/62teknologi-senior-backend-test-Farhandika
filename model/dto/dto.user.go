package dto

type UserInformationDTO struct {
	Id       uint64 `json:"id"`
	Email    string `json:"email"`
	FullName string `json:"full_name"`
	Access   uint32 `json:"access"`
}

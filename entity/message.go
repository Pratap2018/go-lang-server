package entity

type UserRedicrectionMessageType struct {
	AppId          string `json:"appId"`
	ExternalUserId string `json:"externalUserId"`
	EventId        string `json:"eventId"`
	IsEmail        bool   `json:"isEmail"`
	Iat            int64  `json:"iat"`
	Exp            int64  `json:"exp"`
}

type RequestBodyRedirection struct {
	Message     UserRedicrectionMessageType `json:"message"`
	Fyresign    string                      `json:"fyresign"`
	MessageHash string                      `json:"messageHash"`
}

package entities

type OtpMobile struct {
	Mobile string `json:"mobile,omitempty" bson:"mobile,omitempty"`
	Otp    string `json:"otp,omitempty" bson:"otp,omitempty"`
}

type ResponseBody struct {
	Status     string `json:"status"`
	StatusCode string `json:"statusCode"`
	Reason     string `json:"reason"`
}

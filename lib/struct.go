package lib

type Server struct {
	Success     bool        `json:"success"`
	Message     string      `json:"message"`
	ResultsInfo PageInfo    `json:"ResultInfo,omitempty"`
	Results     interface{} `json:"result,omitempty"`
}
type FormRegister struct {
	FullName        string `form:"fullName"`
	Email           string `form:"email"`
	Password        string `form:"password"`
	ConfirmPassword string `form:"confirmPassword" binding:"eqfield=Password"`
}

type PageInfo struct {
	TotalData int `json:"TotalData,omitempty"`
	TotalPage int `json:"TotalPage,omitempty"`
	Page      int `json:"Page,omitempty"`
	Limit     int `json:"Limit,omitempty"`
	Next      int `json:"Next,omitempty"`
	Prev      int `json:"Prev,omitempty"`
}

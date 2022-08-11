package data

const (
	MessageOk                      = "ok"
	MessageErrorInvalidEP          = "invalid email or password"
	MessageErrorTokenInvalid       = "token is invalid"
	AuthorHeader                   = "authorization"
	Bearer                         = "bearer"
	MessageErrorNewDevice          = "invalid new device information"
	MessageErrorCreateNewDevice    = "can not create new device"
	LostFieldInput                 = "Lost input field json"
	ErrorUserFromContext           = "Get user from context error"
	ErrorDateInput                 = "Date Input Error"
	ErrorInsertBorrowRequest       = "Insert Borrow Request Failed"
	ErrorGetProfileIT              = "can not get profile it"
	ErrorNotFoundID                = "can not found id of user"
	ErrorUserUpdatePassword        = "Error user update password"
	ErrorUserConfirmPassword       = "Confirm Password must be the same Password as above"
	ErrorUserCurrentPassword       = "Error user confirm current password"
	ErrorConfirmPassword           = "Confirm password must be the same password as above"
	ErrorGetProfileUser            = "can not get profile user"
	ErrorUserUpdateAccount         = "user update account failed"
	ErrorUserUpdateAccountBindJSON = "error user update account bind json"
	ErrorBindQuery                 = "Listen query device error"
	MessageErrorGetDeviceList      = "can not get device list"
	ErrorGetDeviceDetailUser       = "can not get device detail user"
	ErrorGetNameCategory           = "can not get name category"
	ErrorGetIdFromRouter           = "error fetched id from router"
	ErrorUpdateRequest             = "Error update request"
	ErrorFindUserID                = "Error find UserId"
	ErrorUserPermission            = "Error User doesn't have permission to handle this"
	ErrotRequestStatus             = "Error request status not valid"
	ErrorFindDeviceStatus          = "Error find device status"
	ErrorFindRequestStatus         = "Error find request status"
	ErrorRequestStatus             = "Error request modified status"
	ErrorParamQuery                = "Listen param error"
	ErrorGetListCategory           = "can not get list requests"
	ErrorDetailRequestForbiden     = "can not get detail request"
)

const (
	CodeSuccess                     = 100
	CodeErrorBind                   = 101
	CodeErrorInvalidEmailPassword   = 102
	CodeErrorInvalidToken           = 103
	CodeErrorToken                  = 104
	CodeErrorCreateNewDevice        = 105
	CodeErrorGetUserId              = 118
	CodeErrorDate                   = 119
	CodeErrorCreateNewBorrowRequest = 120
	CodeInternalError               = 106
	CodeBadRequest                  = 107
	CodeErrorGetDeviceList          = 121
	CodeErrorParam                  = 122
	CodeErrorFindDeviceStatus       = 123
	CodeErrorFindRequestStatus      = 124
	CodeErrorRequestIT              = 125
	CodeErrorUpdateRequest          = 126
	CodeErrorCancelRequest          = 127
	CodeErrorPermission             = 128
	CodeErrotRequestStatus          = 129
	CodeErrorForbiden               = 108
)

//EmptyObj object is used when data doesnt want to be null on json
type EmptyObj struct{}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type PageInfo struct {
	Total        int         `json:"total"`
	CurrentTotal int         `json:"current_total"`
	CurrentPage  int         `json:"current_page"`
	Data         interface{} `json:"data"`
}

type PagingResponse struct {
	Code    int      `json:"code"`
	Message string   `json:"message"`
	Data    PageInfo `json:"data"`
}

func NewResponse(Code int, Message string, Data interface{}) *Response {
	return &Response{Code: Code, Message: Message, Data: Data}
}

func NewPagingResponse(Code int, Message string, Total int, CurrentTotal int, CurrentPage int, Data interface{}) *PagingResponse {
	return &PagingResponse{
		Code:    Code,
		Message: Message,
		Data: PageInfo{
			Total:        Total,
			CurrentTotal: CurrentTotal,
			CurrentPage:  CurrentPage,
			Data:         Data,
		},
	}
}

func ErrorResponse(Code int, Message string) *Response {
	return &Response{Code: Code, Message: Message, Data: nil}
}
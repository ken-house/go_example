package errorAssets

type Level int

// 提示错误显示效果
const (
	LevelToast Level = 1 // toast提示
	LevelPopup Level = 2 // 弹窗提示
)

type ErrorNo interface {
	GetTitle() string
	GetCode() int
	ToastError() errorNo
	PopupError(popupTitle, popupContent string, popupStyle int) errorNo
}

type errorNo struct {
	Error error `json:"error"`
}

type error struct {
	Code         int    `json:"code,string"`   // 业务编码
	Level        Level  `json:"level,string"`  // 弹窗提示类型
	Message      string `json:"message"`       // level=1 toast标题
	PopupTitle   string `json:"popup_title"`   // level=2 弹框标题
	PopupContent string `json:"popup_content"` // level=2 弹框内容
}

func NewError(code int, title string) ErrorNo {
	return &errorNo{
		Error: error{
			Code:    code,
			Level:   LevelToast,
			Message: title,
		},
	}
}

func (err *errorNo) GetTitle() string {
	return err.Error.Message
}

func (err *errorNo) GetCode() int {
	return err.Error.Code
}

// ToastError toast提示返回结构
func (err *errorNo) ToastError() errorNo {
	return errorNo{
		error{
			Code:    err.Error.Code,
			Level:   LevelToast,
			Message: err.Error.Message,
		},
	}
}

// PopupError 弹窗返回结构
func (err *errorNo) PopupError(popupTitle, popupContent string, popupStyle int) errorNo {
	return errorNo{
		error{
			Code:         err.Error.Code,
			Level:        LevelPopup,
			PopupTitle:   popupTitle,
			PopupContent: popupContent,
		},
	}
}

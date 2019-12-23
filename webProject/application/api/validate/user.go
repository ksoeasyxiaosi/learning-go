package validate

type Insert struct {
	UserName       string `form:"UserName" binding:"required,gt=1,lt=6"`
	Password       string `form:"Password" binding:"required,gt=6,lt=12,eqfield=PasswordRepeat"`
	PasswordRepeat string `form:"PasswordRepeat" binding:"required,gt=6,lt=12,eqfield=Password"`
	NickName       string `form:"NickName" binding:"required,gt=1,lt=12"`
}

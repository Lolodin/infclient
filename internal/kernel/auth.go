package kernel

type AuthData struct {
	Login, Password, Token string
}

func (a *AuthData) GetToken() string {
	return a.Token
}

func (a *AuthData) GetPassword() string {
	return a.Password
}

func (a *AuthData) SetLogin(str string) {
	a.Login = str
}

func (a *AuthData) SetPassword(str string) {
	a.Password = str
}

func (a *AuthData) SetToken(str string) {
	a.Token = str
}

func (a AuthData) GetLogin() string {
	return a.Login
}

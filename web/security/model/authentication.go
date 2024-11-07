package model

type Authentication struct {
	Principal     interface{}            // 认证通过前，存放username；认证通过后，存放User信息
	Credentials   interface{}            // 认证通过前，存放password；认证通过后，清空，返回nil
	Details       map[string]interface{} // 存放额外的信息，选填，比如client Ip
	Authenticated bool                   // 是否已经通过认证
	AuthMethod    string                 // 认证方式，比如jwt
}

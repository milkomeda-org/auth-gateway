package serializer

// Jwt Json Web Token
type Jwt struct {
}

// IsLegal jwt是否合法
func IsLegal(jwt map[string]interface{}) bool {
	return nil != jwt["exp"] && nil != jwt["iat"] && nil != jwt["user_name"] && nil != jwt["user_id"]
}

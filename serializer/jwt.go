package serializer

type Jwt struct {
}

func IsLegal(jwt map[string]interface{}) bool {
	return nil != jwt["exp"] && nil != jwt["iat"] && nil != jwt["user_name"] && nil != jwt["user_id"]
}

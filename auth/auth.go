package auth

type user struct {
	name string
	pass string
}

func Auth(u user) bool {
	utest := user{"kauê", "123"}
	if u.name == utest.name && u.pass == utest.pass {
		return true
	} else {
		return false
	}
}

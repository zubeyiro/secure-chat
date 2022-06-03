package server

func auth(name string) bool {
	if len(name) < 3 {
		return false
	}
	_, ok := userMap[name]

	return !ok
}

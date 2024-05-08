package metrics

func BoolStr(value bool) string {
	if value {
		return "true"
	}

	return "false"
}

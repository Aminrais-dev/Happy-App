package helper

func FailedResponseHelper(msg interface{}) map[string]interface{} {
	return map[string]interface{}{
		"status":  "failed",
		"message": msg,
	}
}

func SuccessResponseHelper(msg interface{}) map[string]interface{} {
	return map[string]interface{}{
		"status":  "success",
		"message": msg,
	}
}

func SuccessDataResponseHelper(msg, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"status":  "success",
		"message": msg,
		"data":    data,
	}
}

func SuccessFeedResponseHelper(msg, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"status":  "success",
		"message": msg,
		"Feed":    data,
	}
}

func SuccessCartResponseHelper(msg, data interface{}, data2 interface{}) map[string]interface{} {
	return map[string]interface{}{
		"status":    "success",
		"message":   msg,
		"community": data,
		"listcarts": data2,
	}
}

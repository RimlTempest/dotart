package middleware

// Methods 許可するメソッド
func Methods() []string {
	methods := []string{
		"POST",
		"GET",
		"PUT",
		"DELETE",
		"OPTIONS",
	}
	return methods
}

// Headers 許可するヘッダー
func Headers() []string {
	headers := []string{
		"Access-Control-Allow-Origin",
		"Access-Control-Allow-Headers",
		"Content-Type",
		"Content-Length",
		"Connection",
		"Host",
		"Accept",
		"Accept-Encoding",
		"X-CSRF-Token",
		"X-Requested-With",
		"Origin",
		"Authorization",
		"User-Agent",
		"Date",
		"access-token",
		"client",
		"uid",
		"token-type",
	}
	return headers
}

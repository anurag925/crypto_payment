package logger

type ContextKey string
type ContextValue map[ContextKey]any

const (
	ContextKeyValues    ContextKey = "values"
	ContextKeyRequestID ContextKey = "request_id"
	ContextKeyAccountID ContextKey = "account_id"
)

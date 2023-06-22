package form

import "github.com/dev-baris/ai-bot/errorlist"

type Form interface {
	Validate() errorlist.Errors
	String() string
}

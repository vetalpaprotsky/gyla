package testutils

import (
	"fmt"
)

func TestErr(gotErr, wantErr error) string {
	if gotErr == nil && wantErr == nil {
		return ""
	}

	if gotErr != nil && wantErr == nil {
		return fmt.Sprintf("got error %s; want error %s", gotErr.Error(), "nil")
	}

	if gotErr == nil && wantErr != nil {
		return fmt.Sprintf("got error %s; want error %s", "nil", wantErr.Error())
	}

	if gotErr.Error() != wantErr.Error() {
		return fmt.Sprintf("got error %s; want error %s", gotErr.Error(), wantErr.Error())
	}

	return ""
}

func TestGotPtr[T comparable](got, want *T) string {
	if got == nil && want == nil {
		return ""
	}

	if got != nil && want == nil {
		return fmt.Sprintf("got %+v; want %s", *got, "nil")
	}

	if got == nil && want != nil {
		return fmt.Sprintf("got %s; want %+v", "nil", *want)
	}

	if *got != *want {
		return fmt.Sprintf("got %+v; want %+v", *got, *want)
	}

	return ""
}

func TestGot[T comparable](got, want T) string {
	if got != want {
		return fmt.Sprintf("got %+v; want %+v", got, want)
	}

	return ""
}

func FunctionCallName(funcName string, args ...any) string {
	result := fmt.Sprintf("%s(", funcName)

	for i, v := range args {
		result = result + fmt.Sprintf("%+v %T", v, v)
		if i+1 != len(args) {
			result += ","
		}
	}

	return result + ")"
}

func MethodCallName[T any](obj T, funcName string, args ...any) string {
	return fmt.Sprintf("%T%+v.%s", obj, obj, FunctionCallName(funcName, args...))
}

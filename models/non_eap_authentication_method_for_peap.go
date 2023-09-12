package models
import (
    "errors"
)
// Non-EAP methods for authentication when PEAP is the selected EAP type.
type NonEapAuthenticationMethodForPeap int

const (
    // None.
    NONE_NONEAPAUTHENTICATIONMETHODFORPEAP NonEapAuthenticationMethodForPeap = iota
    // Microsoft CHAP Version 2 (MS-CHAP v2).
    MICROSOFTCHAPVERSIONTWO_NONEAPAUTHENTICATIONMETHODFORPEAP
)

func (i NonEapAuthenticationMethodForPeap) String() string {
    return []string{"none", "microsoftChapVersionTwo"}[i]
}
func ParseNonEapAuthenticationMethodForPeap(v string) (any, error) {
    result := NONE_NONEAPAUTHENTICATIONMETHODFORPEAP
    switch v {
        case "none":
            result = NONE_NONEAPAUTHENTICATIONMETHODFORPEAP
        case "microsoftChapVersionTwo":
            result = MICROSOFTCHAPVERSIONTWO_NONEAPAUTHENTICATIONMETHODFORPEAP
        default:
            return 0, errors.New("Unknown NonEapAuthenticationMethodForPeap value: " + v)
    }
    return &result, nil
}
func SerializeNonEapAuthenticationMethodForPeap(values []NonEapAuthenticationMethodForPeap) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i NonEapAuthenticationMethodForPeap) isMultiValue() bool {
    return false
}

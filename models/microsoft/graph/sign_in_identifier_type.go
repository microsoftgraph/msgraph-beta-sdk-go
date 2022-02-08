package graph
import (
    "strings"
    "errors"
)
// 
type SignInIdentifierType int

const (
    USERPRINCIPALNAME_SIGNINIDENTIFIERTYPE SignInIdentifierType = iota
    PHONENUMBER_SIGNINIDENTIFIERTYPE
    PROXYADDRESS_SIGNINIDENTIFIERTYPE
    QRCODE_SIGNINIDENTIFIERTYPE
    ONPREMISESUSERPRINCIPALNAME_SIGNINIDENTIFIERTYPE
    UNKNOWNFUTUREVALUE_SIGNINIDENTIFIERTYPE
)

func (i SignInIdentifierType) String() string {
    return []string{"USERPRINCIPALNAME", "PHONENUMBER", "PROXYADDRESS", "QRCODE", "ONPREMISESUSERPRINCIPALNAME", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseSignInIdentifierType(v string) (interface{}, error) {
    result := USERPRINCIPALNAME_SIGNINIDENTIFIERTYPE
    switch strings.ToUpper(v) {
        case "USERPRINCIPALNAME":
            result = USERPRINCIPALNAME_SIGNINIDENTIFIERTYPE
        case "PHONENUMBER":
            result = PHONENUMBER_SIGNINIDENTIFIERTYPE
        case "PROXYADDRESS":
            result = PROXYADDRESS_SIGNINIDENTIFIERTYPE
        case "QRCODE":
            result = QRCODE_SIGNINIDENTIFIERTYPE
        case "ONPREMISESUSERPRINCIPALNAME":
            result = ONPREMISESUSERPRINCIPALNAME_SIGNINIDENTIFIERTYPE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_SIGNINIDENTIFIERTYPE
        default:
            return 0, errors.New("Unknown SignInIdentifierType value: " + v)
    }
    return &result, nil
}
func SerializeSignInIdentifierType(values []SignInIdentifierType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

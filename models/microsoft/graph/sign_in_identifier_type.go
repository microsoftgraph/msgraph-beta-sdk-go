package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "USERPRINCIPALNAME":
            return USERPRINCIPALNAME_SIGNINIDENTIFIERTYPE, nil
        case "PHONENUMBER":
            return PHONENUMBER_SIGNINIDENTIFIERTYPE, nil
        case "PROXYADDRESS":
            return PROXYADDRESS_SIGNINIDENTIFIERTYPE, nil
        case "QRCODE":
            return QRCODE_SIGNINIDENTIFIERTYPE, nil
        case "ONPREMISESUSERPRINCIPALNAME":
            return ONPREMISESUSERPRINCIPALNAME_SIGNINIDENTIFIERTYPE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_SIGNINIDENTIFIERTYPE, nil
    }
    return 0, errors.New("Unknown SignInIdentifierType value: " + v)
}
func SerializeSignInIdentifierType(values []SignInIdentifierType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

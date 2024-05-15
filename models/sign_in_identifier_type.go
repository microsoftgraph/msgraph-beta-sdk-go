package models
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
    return []string{"userPrincipalName", "phoneNumber", "proxyAddress", "qrCode", "onPremisesUserPrincipalName", "unknownFutureValue"}[i]
}
func ParseSignInIdentifierType(v string) (any, error) {
    result := USERPRINCIPALNAME_SIGNINIDENTIFIERTYPE
    switch v {
        case "userPrincipalName":
            result = USERPRINCIPALNAME_SIGNINIDENTIFIERTYPE
        case "phoneNumber":
            result = PHONENUMBER_SIGNINIDENTIFIERTYPE
        case "proxyAddress":
            result = PROXYADDRESS_SIGNINIDENTIFIERTYPE
        case "qrCode":
            result = QRCODE_SIGNINIDENTIFIERTYPE
        case "onPremisesUserPrincipalName":
            result = ONPREMISESUSERPRINCIPALNAME_SIGNINIDENTIFIERTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_SIGNINIDENTIFIERTYPE
        default:
            return nil, nil
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
func (i SignInIdentifierType) isMultiValue() bool {
    return false
}

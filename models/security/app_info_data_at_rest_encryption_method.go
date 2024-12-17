package security
type AppInfoDataAtRestEncryptionMethod int

const (
    AES_APPINFODATAATRESTENCRYPTIONMETHOD AppInfoDataAtRestEncryptionMethod = iota
    BITLOCKER_APPINFODATAATRESTENCRYPTIONMETHOD
    BLOWFISH_APPINFODATAATRESTENCRYPTIONMETHOD
    DES3_APPINFODATAATRESTENCRYPTIONMETHOD
    DES_APPINFODATAATRESTENCRYPTIONMETHOD
    RC4_APPINFODATAATRESTENCRYPTIONMETHOD
    RSA_APPINFODATAATRESTENCRYPTIONMETHOD
    NOTSUPPORTED_APPINFODATAATRESTENCRYPTIONMETHOD
    UNKNOWN_APPINFODATAATRESTENCRYPTIONMETHOD
    UNKNOWNFUTUREVALUE_APPINFODATAATRESTENCRYPTIONMETHOD
)

func (i AppInfoDataAtRestEncryptionMethod) String() string {
    return []string{"aes", "bitLocker", "blowfish", "des3", "des", "rc4", "rsA", "notSupported", "unknown", "unknownFutureValue"}[i]
}
func ParseAppInfoDataAtRestEncryptionMethod(v string) (any, error) {
    result := AES_APPINFODATAATRESTENCRYPTIONMETHOD
    switch v {
        case "aes":
            result = AES_APPINFODATAATRESTENCRYPTIONMETHOD
        case "bitLocker":
            result = BITLOCKER_APPINFODATAATRESTENCRYPTIONMETHOD
        case "blowfish":
            result = BLOWFISH_APPINFODATAATRESTENCRYPTIONMETHOD
        case "des3":
            result = DES3_APPINFODATAATRESTENCRYPTIONMETHOD
        case "des":
            result = DES_APPINFODATAATRESTENCRYPTIONMETHOD
        case "rc4":
            result = RC4_APPINFODATAATRESTENCRYPTIONMETHOD
        case "rsA":
            result = RSA_APPINFODATAATRESTENCRYPTIONMETHOD
        case "notSupported":
            result = NOTSUPPORTED_APPINFODATAATRESTENCRYPTIONMETHOD
        case "unknown":
            result = UNKNOWN_APPINFODATAATRESTENCRYPTIONMETHOD
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_APPINFODATAATRESTENCRYPTIONMETHOD
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeAppInfoDataAtRestEncryptionMethod(values []AppInfoDataAtRestEncryptionMethod) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AppInfoDataAtRestEncryptionMethod) isMultiValue() bool {
    return false
}

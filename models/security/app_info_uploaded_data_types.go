package security
type AppInfoUploadedDataTypes int

const (
    DOCUMENTS_APPINFOUPLOADEDDATATYPES AppInfoUploadedDataTypes = iota
    MEDIAFILES_APPINFOUPLOADEDDATATYPES
    CODINGFILES_APPINFOUPLOADEDDATATYPES
    CREDITCARDS_APPINFOUPLOADEDDATATYPES
    DATABASEFILES_APPINFOUPLOADEDDATATYPES
    NONE_APPINFOUPLOADEDDATATYPES
    UNKNOWN_APPINFOUPLOADEDDATATYPES
    UNKNOWNFUTUREVALUE_APPINFOUPLOADEDDATATYPES
)

func (i AppInfoUploadedDataTypes) String() string {
    return []string{"documents", "mediaFiles", "codingFiles", "creditCards", "databaseFiles", "none", "unknown", "unknownFutureValue"}[i]
}
func ParseAppInfoUploadedDataTypes(v string) (any, error) {
    result := DOCUMENTS_APPINFOUPLOADEDDATATYPES
    switch v {
        case "documents":
            result = DOCUMENTS_APPINFOUPLOADEDDATATYPES
        case "mediaFiles":
            result = MEDIAFILES_APPINFOUPLOADEDDATATYPES
        case "codingFiles":
            result = CODINGFILES_APPINFOUPLOADEDDATATYPES
        case "creditCards":
            result = CREDITCARDS_APPINFOUPLOADEDDATATYPES
        case "databaseFiles":
            result = DATABASEFILES_APPINFOUPLOADEDDATATYPES
        case "none":
            result = NONE_APPINFOUPLOADEDDATATYPES
        case "unknown":
            result = UNKNOWN_APPINFOUPLOADEDDATATYPES
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_APPINFOUPLOADEDDATATYPES
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeAppInfoUploadedDataTypes(values []AppInfoUploadedDataTypes) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AppInfoUploadedDataTypes) isMultiValue() bool {
    return false
}

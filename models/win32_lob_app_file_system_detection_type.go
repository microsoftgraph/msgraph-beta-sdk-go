// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package models
// Contains all supported file system detection type.
type Win32LobAppFileSystemDetectionType int

const (
    // Not configured.
    NOTCONFIGURED_WIN32LOBAPPFILESYSTEMDETECTIONTYPE Win32LobAppFileSystemDetectionType = iota
    // Whether the specified file or folder exists.
    EXISTS_WIN32LOBAPPFILESYSTEMDETECTIONTYPE
    // Last modified date.
    MODIFIEDDATE_WIN32LOBAPPFILESYSTEMDETECTIONTYPE
    // Created date.
    CREATEDDATE_WIN32LOBAPPFILESYSTEMDETECTIONTYPE
    // Version value type.
    VERSION_WIN32LOBAPPFILESYSTEMDETECTIONTYPE
    // Size detection type.
    SIZEINMB_WIN32LOBAPPFILESYSTEMDETECTIONTYPE
    // The specified file or folder does not exist.
    DOESNOTEXIST_WIN32LOBAPPFILESYSTEMDETECTIONTYPE
)

func (i Win32LobAppFileSystemDetectionType) String() string {
    return []string{"notConfigured", "exists", "modifiedDate", "createdDate", "version", "sizeInMB", "doesNotExist"}[i]
}
func ParseWin32LobAppFileSystemDetectionType(v string) (any, error) {
    result := NOTCONFIGURED_WIN32LOBAPPFILESYSTEMDETECTIONTYPE
    switch v {
        case "notConfigured":
            result = NOTCONFIGURED_WIN32LOBAPPFILESYSTEMDETECTIONTYPE
        case "exists":
            result = EXISTS_WIN32LOBAPPFILESYSTEMDETECTIONTYPE
        case "modifiedDate":
            result = MODIFIEDDATE_WIN32LOBAPPFILESYSTEMDETECTIONTYPE
        case "createdDate":
            result = CREATEDDATE_WIN32LOBAPPFILESYSTEMDETECTIONTYPE
        case "version":
            result = VERSION_WIN32LOBAPPFILESYSTEMDETECTIONTYPE
        case "sizeInMB":
            result = SIZEINMB_WIN32LOBAPPFILESYSTEMDETECTIONTYPE
        case "doesNotExist":
            result = DOESNOTEXIST_WIN32LOBAPPFILESYSTEMDETECTIONTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeWin32LobAppFileSystemDetectionType(values []Win32LobAppFileSystemDetectionType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i Win32LobAppFileSystemDetectionType) isMultiValue() bool {
    return false
}

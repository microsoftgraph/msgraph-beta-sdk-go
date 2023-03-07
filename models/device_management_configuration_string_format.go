package models
import (
    "errors"
)
// 
type DeviceManagementConfigurationStringFormat int

const (
    // Indicates a string with no well-defined format expected.
    NONE_DEVICEMANAGEMENTCONFIGURATIONSTRINGFORMAT DeviceManagementConfigurationStringFormat = iota
    // Indicates a string that is expected to be a valid email address.
    EMAIL_DEVICEMANAGEMENTCONFIGURATIONSTRINGFORMAT
    // Indicates a string that is expected to be a valid GUID.
    GUID_DEVICEMANAGEMENTCONFIGURATIONSTRINGFORMAT
    // Indicates a string that is expected to be a valid IP address.
    IP_DEVICEMANAGEMENTCONFIGURATIONSTRINGFORMAT
    // Indicates a string that is base64 encoded.
    BASE64_DEVICEMANAGEMENTCONFIGURATIONSTRINGFORMAT
    // Indicates a string that is expected to be a valid URL.
    URL_DEVICEMANAGEMENTCONFIGURATIONSTRINGFORMAT
    // Indicates a string that should refer to a version.
    VERSION_DEVICEMANAGEMENTCONFIGURATIONSTRINGFORMAT
    // Indicates a string that is expected to be a valid XML.
    XML_DEVICEMANAGEMENTCONFIGURATIONSTRINGFORMAT
    // Indicates a string that is expected to be a valid date.
    DATE_DEVICEMANAGEMENTCONFIGURATIONSTRINGFORMAT
    // Indicates a string that is expected to be a valid time.
    TIME_DEVICEMANAGEMENTCONFIGURATIONSTRINGFORMAT
    BINARY_DEVICEMANAGEMENTCONFIGURATIONSTRINGFORMAT
    // Indicates a string that is expected to be a valid Regex string.
    REGEX_DEVICEMANAGEMENTCONFIGURATIONSTRINGFORMAT
    // Indicates a string that is expected to be a valid JSON string.
    JSON_DEVICEMANAGEMENTCONFIGURATIONSTRINGFORMAT
    // Indicates a string that is expected to be a valid Datetime.
    DATETIME_DEVICEMANAGEMENTCONFIGURATIONSTRINGFORMAT
    // Indicates a Windows SKU applicability value that maps to Intune.
    SURFACEHUB_DEVICEMANAGEMENTCONFIGURATIONSTRINGFORMAT
    // String whose value is a bash script
    BASHSCRIPT_DEVICEMANAGEMENTCONFIGURATIONSTRINGFORMAT
    // Sentinel member for cases where the client cannot handle the new enum values.
    UNKNOWNFUTUREVALUE_DEVICEMANAGEMENTCONFIGURATIONSTRINGFORMAT
)

func (i DeviceManagementConfigurationStringFormat) String() string {
    return []string{"none", "email", "guid", "ip", "base64", "url", "version", "xml", "date", "time", "binary", "regEx", "json", "dateTime", "surfaceHub", "bashScript", "unknownFutureValue"}[i]
}
func ParseDeviceManagementConfigurationStringFormat(v string) (any, error) {
    result := NONE_DEVICEMANAGEMENTCONFIGURATIONSTRINGFORMAT
    switch v {
        case "none":
            result = NONE_DEVICEMANAGEMENTCONFIGURATIONSTRINGFORMAT
        case "email":
            result = EMAIL_DEVICEMANAGEMENTCONFIGURATIONSTRINGFORMAT
        case "guid":
            result = GUID_DEVICEMANAGEMENTCONFIGURATIONSTRINGFORMAT
        case "ip":
            result = IP_DEVICEMANAGEMENTCONFIGURATIONSTRINGFORMAT
        case "base64":
            result = BASE64_DEVICEMANAGEMENTCONFIGURATIONSTRINGFORMAT
        case "url":
            result = URL_DEVICEMANAGEMENTCONFIGURATIONSTRINGFORMAT
        case "version":
            result = VERSION_DEVICEMANAGEMENTCONFIGURATIONSTRINGFORMAT
        case "xml":
            result = XML_DEVICEMANAGEMENTCONFIGURATIONSTRINGFORMAT
        case "date":
            result = DATE_DEVICEMANAGEMENTCONFIGURATIONSTRINGFORMAT
        case "time":
            result = TIME_DEVICEMANAGEMENTCONFIGURATIONSTRINGFORMAT
        case "binary":
            result = BINARY_DEVICEMANAGEMENTCONFIGURATIONSTRINGFORMAT
        case "regEx":
            result = REGEX_DEVICEMANAGEMENTCONFIGURATIONSTRINGFORMAT
        case "json":
            result = JSON_DEVICEMANAGEMENTCONFIGURATIONSTRINGFORMAT
        case "dateTime":
            result = DATETIME_DEVICEMANAGEMENTCONFIGURATIONSTRINGFORMAT
        case "surfaceHub":
            result = SURFACEHUB_DEVICEMANAGEMENTCONFIGURATIONSTRINGFORMAT
        case "bashScript":
            result = BASHSCRIPT_DEVICEMANAGEMENTCONFIGURATIONSTRINGFORMAT
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_DEVICEMANAGEMENTCONFIGURATIONSTRINGFORMAT
        default:
            return 0, errors.New("Unknown DeviceManagementConfigurationStringFormat value: " + v)
    }
    return &result, nil
}
func SerializeDeviceManagementConfigurationStringFormat(values []DeviceManagementConfigurationStringFormat) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

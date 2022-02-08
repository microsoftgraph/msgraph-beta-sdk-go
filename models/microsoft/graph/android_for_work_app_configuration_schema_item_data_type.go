package graph
import (
    "strings"
    "errors"
)
// 
type AndroidForWorkAppConfigurationSchemaItemDataType int

const (
    BOOL_ANDROIDFORWORKAPPCONFIGURATIONSCHEMAITEMDATATYPE AndroidForWorkAppConfigurationSchemaItemDataType = iota
    INTEGER_ANDROIDFORWORKAPPCONFIGURATIONSCHEMAITEMDATATYPE
    STRING_ANDROIDFORWORKAPPCONFIGURATIONSCHEMAITEMDATATYPE
    CHOICE_ANDROIDFORWORKAPPCONFIGURATIONSCHEMAITEMDATATYPE
    MULTISELECT_ANDROIDFORWORKAPPCONFIGURATIONSCHEMAITEMDATATYPE
    BUNDLE_ANDROIDFORWORKAPPCONFIGURATIONSCHEMAITEMDATATYPE
    BUNDLEARRAY_ANDROIDFORWORKAPPCONFIGURATIONSCHEMAITEMDATATYPE
    HIDDEN_ANDROIDFORWORKAPPCONFIGURATIONSCHEMAITEMDATATYPE
)

func (i AndroidForWorkAppConfigurationSchemaItemDataType) String() string {
    return []string{"BOOL", "INTEGER", "STRING", "CHOICE", "MULTISELECT", "BUNDLE", "BUNDLEARRAY", "HIDDEN"}[i]
}
func ParseAndroidForWorkAppConfigurationSchemaItemDataType(v string) (interface{}, error) {
    result := BOOL_ANDROIDFORWORKAPPCONFIGURATIONSCHEMAITEMDATATYPE
    switch strings.ToUpper(v) {
        case "BOOL":
            result = BOOL_ANDROIDFORWORKAPPCONFIGURATIONSCHEMAITEMDATATYPE
        case "INTEGER":
            result = INTEGER_ANDROIDFORWORKAPPCONFIGURATIONSCHEMAITEMDATATYPE
        case "STRING":
            result = STRING_ANDROIDFORWORKAPPCONFIGURATIONSCHEMAITEMDATATYPE
        case "CHOICE":
            result = CHOICE_ANDROIDFORWORKAPPCONFIGURATIONSCHEMAITEMDATATYPE
        case "MULTISELECT":
            result = MULTISELECT_ANDROIDFORWORKAPPCONFIGURATIONSCHEMAITEMDATATYPE
        case "BUNDLE":
            result = BUNDLE_ANDROIDFORWORKAPPCONFIGURATIONSCHEMAITEMDATATYPE
        case "BUNDLEARRAY":
            result = BUNDLEARRAY_ANDROIDFORWORKAPPCONFIGURATIONSCHEMAITEMDATATYPE
        case "HIDDEN":
            result = HIDDEN_ANDROIDFORWORKAPPCONFIGURATIONSCHEMAITEMDATATYPE
        default:
            return 0, errors.New("Unknown AndroidForWorkAppConfigurationSchemaItemDataType value: " + v)
    }
    return &result, nil
}
func SerializeAndroidForWorkAppConfigurationSchemaItemDataType(values []AndroidForWorkAppConfigurationSchemaItemDataType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

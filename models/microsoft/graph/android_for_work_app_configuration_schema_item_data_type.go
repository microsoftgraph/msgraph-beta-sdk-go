package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "BOOL":
            return BOOL_ANDROIDFORWORKAPPCONFIGURATIONSCHEMAITEMDATATYPE, nil
        case "INTEGER":
            return INTEGER_ANDROIDFORWORKAPPCONFIGURATIONSCHEMAITEMDATATYPE, nil
        case "STRING":
            return STRING_ANDROIDFORWORKAPPCONFIGURATIONSCHEMAITEMDATATYPE, nil
        case "CHOICE":
            return CHOICE_ANDROIDFORWORKAPPCONFIGURATIONSCHEMAITEMDATATYPE, nil
        case "MULTISELECT":
            return MULTISELECT_ANDROIDFORWORKAPPCONFIGURATIONSCHEMAITEMDATATYPE, nil
        case "BUNDLE":
            return BUNDLE_ANDROIDFORWORKAPPCONFIGURATIONSCHEMAITEMDATATYPE, nil
        case "BUNDLEARRAY":
            return BUNDLEARRAY_ANDROIDFORWORKAPPCONFIGURATIONSCHEMAITEMDATATYPE, nil
        case "HIDDEN":
            return HIDDEN_ANDROIDFORWORKAPPCONFIGURATIONSCHEMAITEMDATATYPE, nil
    }
    return 0, errors.New("Unknown AndroidForWorkAppConfigurationSchemaItemDataType value: " + v)
}
func SerializeAndroidForWorkAppConfigurationSchemaItemDataType(values []AndroidForWorkAppConfigurationSchemaItemDataType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the deviceManagement singleton.
type AndroidManagedStoreAppConfigurationSchemaItemDataType int

const (
    BOOL_ANDROIDMANAGEDSTOREAPPCONFIGURATIONSCHEMAITEMDATATYPE AndroidManagedStoreAppConfigurationSchemaItemDataType = iota
    INTEGER_ANDROIDMANAGEDSTOREAPPCONFIGURATIONSCHEMAITEMDATATYPE
    STRING_ANDROIDMANAGEDSTOREAPPCONFIGURATIONSCHEMAITEMDATATYPE
    CHOICE_ANDROIDMANAGEDSTOREAPPCONFIGURATIONSCHEMAITEMDATATYPE
    MULTISELECT_ANDROIDMANAGEDSTOREAPPCONFIGURATIONSCHEMAITEMDATATYPE
    BUNDLE_ANDROIDMANAGEDSTOREAPPCONFIGURATIONSCHEMAITEMDATATYPE
    BUNDLEARRAY_ANDROIDMANAGEDSTOREAPPCONFIGURATIONSCHEMAITEMDATATYPE
    HIDDEN_ANDROIDMANAGEDSTOREAPPCONFIGURATIONSCHEMAITEMDATATYPE
)

func (i AndroidManagedStoreAppConfigurationSchemaItemDataType) String() string {
    return []string{"BOOL", "INTEGER", "STRING", "CHOICE", "MULTISELECT", "BUNDLE", "BUNDLEARRAY", "HIDDEN"}[i]
}
func ParseAndroidManagedStoreAppConfigurationSchemaItemDataType(v string) (interface{}, error) {
    result := BOOL_ANDROIDMANAGEDSTOREAPPCONFIGURATIONSCHEMAITEMDATATYPE
    switch strings.ToUpper(v) {
        case "BOOL":
            result = BOOL_ANDROIDMANAGEDSTOREAPPCONFIGURATIONSCHEMAITEMDATATYPE
        case "INTEGER":
            result = INTEGER_ANDROIDMANAGEDSTOREAPPCONFIGURATIONSCHEMAITEMDATATYPE
        case "STRING":
            result = STRING_ANDROIDMANAGEDSTOREAPPCONFIGURATIONSCHEMAITEMDATATYPE
        case "CHOICE":
            result = CHOICE_ANDROIDMANAGEDSTOREAPPCONFIGURATIONSCHEMAITEMDATATYPE
        case "MULTISELECT":
            result = MULTISELECT_ANDROIDMANAGEDSTOREAPPCONFIGURATIONSCHEMAITEMDATATYPE
        case "BUNDLE":
            result = BUNDLE_ANDROIDMANAGEDSTOREAPPCONFIGURATIONSCHEMAITEMDATATYPE
        case "BUNDLEARRAY":
            result = BUNDLEARRAY_ANDROIDMANAGEDSTOREAPPCONFIGURATIONSCHEMAITEMDATATYPE
        case "HIDDEN":
            result = HIDDEN_ANDROIDMANAGEDSTOREAPPCONFIGURATIONSCHEMAITEMDATATYPE
        default:
            return 0, errors.New("Unknown AndroidManagedStoreAppConfigurationSchemaItemDataType value: " + v)
    }
    return &result, nil
}
func SerializeAndroidManagedStoreAppConfigurationSchemaItemDataType(values []AndroidManagedStoreAppConfigurationSchemaItemDataType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

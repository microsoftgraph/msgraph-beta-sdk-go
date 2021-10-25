package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "BOOL":
            return BOOL_ANDROIDMANAGEDSTOREAPPCONFIGURATIONSCHEMAITEMDATATYPE, nil
        case "INTEGER":
            return INTEGER_ANDROIDMANAGEDSTOREAPPCONFIGURATIONSCHEMAITEMDATATYPE, nil
        case "STRING":
            return STRING_ANDROIDMANAGEDSTOREAPPCONFIGURATIONSCHEMAITEMDATATYPE, nil
        case "CHOICE":
            return CHOICE_ANDROIDMANAGEDSTOREAPPCONFIGURATIONSCHEMAITEMDATATYPE, nil
        case "MULTISELECT":
            return MULTISELECT_ANDROIDMANAGEDSTOREAPPCONFIGURATIONSCHEMAITEMDATATYPE, nil
        case "BUNDLE":
            return BUNDLE_ANDROIDMANAGEDSTOREAPPCONFIGURATIONSCHEMAITEMDATATYPE, nil
        case "BUNDLEARRAY":
            return BUNDLEARRAY_ANDROIDMANAGEDSTOREAPPCONFIGURATIONSCHEMAITEMDATATYPE, nil
        case "HIDDEN":
            return HIDDEN_ANDROIDMANAGEDSTOREAPPCONFIGURATIONSCHEMAITEMDATATYPE, nil
    }
    return 0, errors.New("Unknown AndroidManagedStoreAppConfigurationSchemaItemDataType value: " + v)
}
func SerializeAndroidManagedStoreAppConfigurationSchemaItemDataType(values []AndroidManagedStoreAppConfigurationSchemaItemDataType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

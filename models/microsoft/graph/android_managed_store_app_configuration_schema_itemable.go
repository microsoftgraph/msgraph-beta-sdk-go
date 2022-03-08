package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AndroidManagedStoreAppConfigurationSchemaItemable 
type AndroidManagedStoreAppConfigurationSchemaItemable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetDataType()(*AndroidManagedStoreAppConfigurationSchemaItemDataType)
    GetDefaultBoolValue()(*bool)
    GetDefaultIntValue()(*int32)
    GetDefaultStringArrayValue()([]string)
    GetDefaultStringValue()(*string)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetIndex()(*int32)
    GetParentIndex()(*int32)
    GetSchemaItemKey()(*string)
    GetSelections()([]KeyValuePairable)
    SetDataType(value *AndroidManagedStoreAppConfigurationSchemaItemDataType)()
    SetDefaultBoolValue(value *bool)()
    SetDefaultIntValue(value *int32)()
    SetDefaultStringArrayValue(value []string)()
    SetDefaultStringValue(value *string)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetIndex(value *int32)()
    SetParentIndex(value *int32)()
    SetSchemaItemKey(value *string)()
    SetSelections(value []KeyValuePairable)()
}

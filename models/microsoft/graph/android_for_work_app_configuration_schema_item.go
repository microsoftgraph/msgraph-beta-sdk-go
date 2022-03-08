package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AndroidForWorkAppConfigurationSchemaItem provides operations to manage the deviceManagement singleton.
type AndroidForWorkAppConfigurationSchemaItem struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The type of value this item describes. Possible values are: bool, integer, string, choice, multiselect, bundle, bundleArray, hidden.
    dataType *AndroidForWorkAppConfigurationSchemaItemDataType;
    // Default value for boolean type items, if specified by the app developer
    defaultBoolValue *bool;
    // Default value for integer type items, if specified by the app developer
    defaultIntValue *int32;
    // Default value for string array type items, if specified by the app developer
    defaultStringArrayValue []string;
    // Default value for string type items, if specified by the app developer
    defaultStringValue *string;
    // Description of what the item controls within the application
    description *string;
    // Human readable name
    displayName *string;
    // Unique key the application uses to identify the item
    schemaItemKey *string;
    // List of human readable name/value pairs for the valid values that can be set for this item (Choice and Multiselect items only)
    selections []KeyValuePairable;
}
// NewAndroidForWorkAppConfigurationSchemaItem instantiates a new androidForWorkAppConfigurationSchemaItem and sets the default values.
func NewAndroidForWorkAppConfigurationSchemaItem()(*AndroidForWorkAppConfigurationSchemaItem) {
    m := &AndroidForWorkAppConfigurationSchemaItem{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAndroidForWorkAppConfigurationSchemaItemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidForWorkAppConfigurationSchemaItemFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewAndroidForWorkAppConfigurationSchemaItem(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AndroidForWorkAppConfigurationSchemaItem) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDataType gets the dataType property value. The type of value this item describes. Possible values are: bool, integer, string, choice, multiselect, bundle, bundleArray, hidden.
func (m *AndroidForWorkAppConfigurationSchemaItem) GetDataType()(*AndroidForWorkAppConfigurationSchemaItemDataType) {
    if m == nil {
        return nil
    } else {
        return m.dataType
    }
}
// GetDefaultBoolValue gets the defaultBoolValue property value. Default value for boolean type items, if specified by the app developer
func (m *AndroidForWorkAppConfigurationSchemaItem) GetDefaultBoolValue()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.defaultBoolValue
    }
}
// GetDefaultIntValue gets the defaultIntValue property value. Default value for integer type items, if specified by the app developer
func (m *AndroidForWorkAppConfigurationSchemaItem) GetDefaultIntValue()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.defaultIntValue
    }
}
// GetDefaultStringArrayValue gets the defaultStringArrayValue property value. Default value for string array type items, if specified by the app developer
func (m *AndroidForWorkAppConfigurationSchemaItem) GetDefaultStringArrayValue()([]string) {
    if m == nil {
        return nil
    } else {
        return m.defaultStringArrayValue
    }
}
// GetDefaultStringValue gets the defaultStringValue property value. Default value for string type items, if specified by the app developer
func (m *AndroidForWorkAppConfigurationSchemaItem) GetDefaultStringValue()(*string) {
    if m == nil {
        return nil
    } else {
        return m.defaultStringValue
    }
}
// GetDescription gets the description property value. Description of what the item controls within the application
func (m *AndroidForWorkAppConfigurationSchemaItem) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDisplayName gets the displayName property value. Human readable name
func (m *AndroidForWorkAppConfigurationSchemaItem) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidForWorkAppConfigurationSchemaItem) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["dataType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidForWorkAppConfigurationSchemaItemDataType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDataType(val.(*AndroidForWorkAppConfigurationSchemaItemDataType))
        }
        return nil
    }
    res["defaultBoolValue"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultBoolValue(val)
        }
        return nil
    }
    res["defaultIntValue"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultIntValue(val)
        }
        return nil
    }
    res["defaultStringArrayValue"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetDefaultStringArrayValue(res)
        }
        return nil
    }
    res["defaultStringValue"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultStringValue(val)
        }
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["schemaItemKey"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSchemaItemKey(val)
        }
        return nil
    }
    res["selections"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateKeyValuePairFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]KeyValuePairable, len(val))
            for i, v := range val {
                res[i] = v.(KeyValuePairable)
            }
            m.SetSelections(res)
        }
        return nil
    }
    return res
}
// GetSchemaItemKey gets the schemaItemKey property value. Unique key the application uses to identify the item
func (m *AndroidForWorkAppConfigurationSchemaItem) GetSchemaItemKey()(*string) {
    if m == nil {
        return nil
    } else {
        return m.schemaItemKey
    }
}
// GetSelections gets the selections property value. List of human readable name/value pairs for the valid values that can be set for this item (Choice and Multiselect items only)
func (m *AndroidForWorkAppConfigurationSchemaItem) GetSelections()([]KeyValuePairable) {
    if m == nil {
        return nil
    } else {
        return m.selections
    }
}
func (m *AndroidForWorkAppConfigurationSchemaItem) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *AndroidForWorkAppConfigurationSchemaItem) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetDataType() != nil {
        cast := (*m.GetDataType()).String()
        err := writer.WriteStringValue("dataType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("defaultBoolValue", m.GetDefaultBoolValue())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("defaultIntValue", m.GetDefaultIntValue())
        if err != nil {
            return err
        }
    }
    if m.GetDefaultStringArrayValue() != nil {
        err := writer.WriteCollectionOfStringValues("defaultStringArrayValue", m.GetDefaultStringArrayValue())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("defaultStringValue", m.GetDefaultStringValue())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("schemaItemKey", m.GetSchemaItemKey())
        if err != nil {
            return err
        }
    }
    if m.GetSelections() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSelections()))
        for i, v := range m.GetSelections() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("selections", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AndroidForWorkAppConfigurationSchemaItem) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDataType sets the dataType property value. The type of value this item describes. Possible values are: bool, integer, string, choice, multiselect, bundle, bundleArray, hidden.
func (m *AndroidForWorkAppConfigurationSchemaItem) SetDataType(value *AndroidForWorkAppConfigurationSchemaItemDataType)() {
    if m != nil {
        m.dataType = value
    }
}
// SetDefaultBoolValue sets the defaultBoolValue property value. Default value for boolean type items, if specified by the app developer
func (m *AndroidForWorkAppConfigurationSchemaItem) SetDefaultBoolValue(value *bool)() {
    if m != nil {
        m.defaultBoolValue = value
    }
}
// SetDefaultIntValue sets the defaultIntValue property value. Default value for integer type items, if specified by the app developer
func (m *AndroidForWorkAppConfigurationSchemaItem) SetDefaultIntValue(value *int32)() {
    if m != nil {
        m.defaultIntValue = value
    }
}
// SetDefaultStringArrayValue sets the defaultStringArrayValue property value. Default value for string array type items, if specified by the app developer
func (m *AndroidForWorkAppConfigurationSchemaItem) SetDefaultStringArrayValue(value []string)() {
    if m != nil {
        m.defaultStringArrayValue = value
    }
}
// SetDefaultStringValue sets the defaultStringValue property value. Default value for string type items, if specified by the app developer
func (m *AndroidForWorkAppConfigurationSchemaItem) SetDefaultStringValue(value *string)() {
    if m != nil {
        m.defaultStringValue = value
    }
}
// SetDescription sets the description property value. Description of what the item controls within the application
func (m *AndroidForWorkAppConfigurationSchemaItem) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDisplayName sets the displayName property value. Human readable name
func (m *AndroidForWorkAppConfigurationSchemaItem) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetSchemaItemKey sets the schemaItemKey property value. Unique key the application uses to identify the item
func (m *AndroidForWorkAppConfigurationSchemaItem) SetSchemaItemKey(value *string)() {
    if m != nil {
        m.schemaItemKey = value
    }
}
// SetSelections sets the selections property value. List of human readable name/value pairs for the valid values that can be set for this item (Choice and Multiselect items only)
func (m *AndroidForWorkAppConfigurationSchemaItem) SetSelections(value []KeyValuePairable)() {
    if m != nil {
        m.selections = value
    }
}

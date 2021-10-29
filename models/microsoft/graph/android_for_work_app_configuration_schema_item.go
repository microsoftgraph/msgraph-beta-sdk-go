package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
    selections []KeyValuePair;
}
// Instantiates a new androidForWorkAppConfigurationSchemaItem and sets the default values.
func NewAndroidForWorkAppConfigurationSchemaItem()(*AndroidForWorkAppConfigurationSchemaItem) {
    m := &AndroidForWorkAppConfigurationSchemaItem{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AndroidForWorkAppConfigurationSchemaItem) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the dataType property value. The type of value this item describes. Possible values are: bool, integer, string, choice, multiselect, bundle, bundleArray, hidden.
func (m *AndroidForWorkAppConfigurationSchemaItem) GetDataType()(*AndroidForWorkAppConfigurationSchemaItemDataType) {
    if m == nil {
        return nil
    } else {
        return m.dataType
    }
}
// Gets the defaultBoolValue property value. Default value for boolean type items, if specified by the app developer
func (m *AndroidForWorkAppConfigurationSchemaItem) GetDefaultBoolValue()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.defaultBoolValue
    }
}
// Gets the defaultIntValue property value. Default value for integer type items, if specified by the app developer
func (m *AndroidForWorkAppConfigurationSchemaItem) GetDefaultIntValue()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.defaultIntValue
    }
}
// Gets the defaultStringArrayValue property value. Default value for string array type items, if specified by the app developer
func (m *AndroidForWorkAppConfigurationSchemaItem) GetDefaultStringArrayValue()([]string) {
    if m == nil {
        return nil
    } else {
        return m.defaultStringArrayValue
    }
}
// Gets the defaultStringValue property value. Default value for string type items, if specified by the app developer
func (m *AndroidForWorkAppConfigurationSchemaItem) GetDefaultStringValue()(*string) {
    if m == nil {
        return nil
    } else {
        return m.defaultStringValue
    }
}
// Gets the description property value. Description of what the item controls within the application
func (m *AndroidForWorkAppConfigurationSchemaItem) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the displayName property value. Human readable name
func (m *AndroidForWorkAppConfigurationSchemaItem) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the schemaItemKey property value. Unique key the application uses to identify the item
func (m *AndroidForWorkAppConfigurationSchemaItem) GetSchemaItemKey()(*string) {
    if m == nil {
        return nil
    } else {
        return m.schemaItemKey
    }
}
// Gets the selections property value. List of human readable name/value pairs for the valid values that can be set for this item (Choice and Multiselect items only)
func (m *AndroidForWorkAppConfigurationSchemaItem) GetSelections()([]KeyValuePair) {
    if m == nil {
        return nil
    } else {
        return m.selections
    }
}
// The deserialization information for the current model
func (m *AndroidForWorkAppConfigurationSchemaItem) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["dataType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidForWorkAppConfigurationSchemaItemDataType)
        if err != nil {
            return err
        }
        cast := val.(AndroidForWorkAppConfigurationSchemaItemDataType)
        m.SetDataType(&cast)
        return nil
    }
    res["defaultBoolValue"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetDefaultBoolValue(val)
        return nil
    }
    res["defaultIntValue"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetDefaultIntValue(val)
        return nil
    }
    res["defaultStringArrayValue"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetDefaultStringArrayValue(res)
        return nil
    }
    res["defaultStringValue"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDefaultStringValue(val)
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDescription(val)
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["schemaItemKey"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSchemaItemKey(val)
        return nil
    }
    res["selections"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewKeyValuePair() })
        if err != nil {
            return err
        }
        res := make([]KeyValuePair, len(val))
        for i, v := range val {
            res[i] = *(v.(*KeyValuePair))
        }
        m.SetSelections(res)
        return nil
    }
    return res
}
func (m *AndroidForWorkAppConfigurationSchemaItem) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *AndroidForWorkAppConfigurationSchemaItem) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetDataType() != nil {
        cast := m.GetDataType().String()
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
    {
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
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSelections()))
        for i, v := range m.GetSelections() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *AndroidForWorkAppConfigurationSchemaItem) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the dataType property value. The type of value this item describes. Possible values are: bool, integer, string, choice, multiselect, bundle, bundleArray, hidden.
// Parameters:
//  - value : Value to set for the dataType property.
func (m *AndroidForWorkAppConfigurationSchemaItem) SetDataType(value *AndroidForWorkAppConfigurationSchemaItemDataType)() {
    m.dataType = value
}
// Sets the defaultBoolValue property value. Default value for boolean type items, if specified by the app developer
// Parameters:
//  - value : Value to set for the defaultBoolValue property.
func (m *AndroidForWorkAppConfigurationSchemaItem) SetDefaultBoolValue(value *bool)() {
    m.defaultBoolValue = value
}
// Sets the defaultIntValue property value. Default value for integer type items, if specified by the app developer
// Parameters:
//  - value : Value to set for the defaultIntValue property.
func (m *AndroidForWorkAppConfigurationSchemaItem) SetDefaultIntValue(value *int32)() {
    m.defaultIntValue = value
}
// Sets the defaultStringArrayValue property value. Default value for string array type items, if specified by the app developer
// Parameters:
//  - value : Value to set for the defaultStringArrayValue property.
func (m *AndroidForWorkAppConfigurationSchemaItem) SetDefaultStringArrayValue(value []string)() {
    m.defaultStringArrayValue = value
}
// Sets the defaultStringValue property value. Default value for string type items, if specified by the app developer
// Parameters:
//  - value : Value to set for the defaultStringValue property.
func (m *AndroidForWorkAppConfigurationSchemaItem) SetDefaultStringValue(value *string)() {
    m.defaultStringValue = value
}
// Sets the description property value. Description of what the item controls within the application
// Parameters:
//  - value : Value to set for the description property.
func (m *AndroidForWorkAppConfigurationSchemaItem) SetDescription(value *string)() {
    m.description = value
}
// Sets the displayName property value. Human readable name
// Parameters:
//  - value : Value to set for the displayName property.
func (m *AndroidForWorkAppConfigurationSchemaItem) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the schemaItemKey property value. Unique key the application uses to identify the item
// Parameters:
//  - value : Value to set for the schemaItemKey property.
func (m *AndroidForWorkAppConfigurationSchemaItem) SetSchemaItemKey(value *string)() {
    m.schemaItemKey = value
}
// Sets the selections property value. List of human readable name/value pairs for the valid values that can be set for this item (Choice and Multiselect items only)
// Parameters:
//  - value : Value to set for the selections property.
func (m *AndroidForWorkAppConfigurationSchemaItem) SetSelections(value []KeyValuePair)() {
    m.selections = value
}

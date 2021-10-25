package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type AndroidManagedStoreAppConfigurationSchemaItem struct {
    additionalData map[string]interface{};
    dataType *AndroidManagedStoreAppConfigurationSchemaItemDataType;
    defaultBoolValue *bool;
    defaultIntValue *int32;
    defaultStringArrayValue []string;
    defaultStringValue *string;
    description *string;
    displayName *string;
    index *int32;
    parentIndex *int32;
    schemaItemKey *string;
    selections []KeyValuePair;
}
func NewAndroidManagedStoreAppConfigurationSchemaItem()(*AndroidManagedStoreAppConfigurationSchemaItem) {
    m := &AndroidManagedStoreAppConfigurationSchemaItem{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *AndroidManagedStoreAppConfigurationSchemaItem) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *AndroidManagedStoreAppConfigurationSchemaItem) GetDataType()(*AndroidManagedStoreAppConfigurationSchemaItemDataType) {
    if m == nil {
        return nil
    } else {
        return m.dataType
    }
}
func (m *AndroidManagedStoreAppConfigurationSchemaItem) GetDefaultBoolValue()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.defaultBoolValue
    }
}
func (m *AndroidManagedStoreAppConfigurationSchemaItem) GetDefaultIntValue()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.defaultIntValue
    }
}
func (m *AndroidManagedStoreAppConfigurationSchemaItem) GetDefaultStringArrayValue()([]string) {
    if m == nil {
        return nil
    } else {
        return m.defaultStringArrayValue
    }
}
func (m *AndroidManagedStoreAppConfigurationSchemaItem) GetDefaultStringValue()(*string) {
    if m == nil {
        return nil
    } else {
        return m.defaultStringValue
    }
}
func (m *AndroidManagedStoreAppConfigurationSchemaItem) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
func (m *AndroidManagedStoreAppConfigurationSchemaItem) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
func (m *AndroidManagedStoreAppConfigurationSchemaItem) GetIndex()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.index
    }
}
func (m *AndroidManagedStoreAppConfigurationSchemaItem) GetParentIndex()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.parentIndex
    }
}
func (m *AndroidManagedStoreAppConfigurationSchemaItem) GetSchemaItemKey()(*string) {
    if m == nil {
        return nil
    } else {
        return m.schemaItemKey
    }
}
func (m *AndroidManagedStoreAppConfigurationSchemaItem) GetSelections()([]KeyValuePair) {
    if m == nil {
        return nil
    } else {
        return m.selections
    }
}
func (m *AndroidManagedStoreAppConfigurationSchemaItem) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["dataType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidManagedStoreAppConfigurationSchemaItemDataType)
        if err != nil {
            return err
        }
        cast := val.(AndroidManagedStoreAppConfigurationSchemaItemDataType)
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
    res["index"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetIndex(val)
        return nil
    }
    res["parentIndex"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetParentIndex(val)
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
func (m *AndroidManagedStoreAppConfigurationSchemaItem) IsNil()(bool) {
    return m == nil
}
func (m *AndroidManagedStoreAppConfigurationSchemaItem) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        err := writer.WriteInt32Value("index", m.GetIndex())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("parentIndex", m.GetParentIndex())
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
func (m *AndroidManagedStoreAppConfigurationSchemaItem) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *AndroidManagedStoreAppConfigurationSchemaItem) SetDataType(value *AndroidManagedStoreAppConfigurationSchemaItemDataType)() {
    m.dataType = value
}
func (m *AndroidManagedStoreAppConfigurationSchemaItem) SetDefaultBoolValue(value *bool)() {
    m.defaultBoolValue = value
}
func (m *AndroidManagedStoreAppConfigurationSchemaItem) SetDefaultIntValue(value *int32)() {
    m.defaultIntValue = value
}
func (m *AndroidManagedStoreAppConfigurationSchemaItem) SetDefaultStringArrayValue(value []string)() {
    m.defaultStringArrayValue = value
}
func (m *AndroidManagedStoreAppConfigurationSchemaItem) SetDefaultStringValue(value *string)() {
    m.defaultStringValue = value
}
func (m *AndroidManagedStoreAppConfigurationSchemaItem) SetDescription(value *string)() {
    m.description = value
}
func (m *AndroidManagedStoreAppConfigurationSchemaItem) SetDisplayName(value *string)() {
    m.displayName = value
}
func (m *AndroidManagedStoreAppConfigurationSchemaItem) SetIndex(value *int32)() {
    m.index = value
}
func (m *AndroidManagedStoreAppConfigurationSchemaItem) SetParentIndex(value *int32)() {
    m.parentIndex = value
}
func (m *AndroidManagedStoreAppConfigurationSchemaItem) SetSchemaItemKey(value *string)() {
    m.schemaItemKey = value
}
func (m *AndroidManagedStoreAppConfigurationSchemaItem) SetSelections(value []KeyValuePair)() {
    m.selections = value
}

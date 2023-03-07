package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// AndroidManagedStoreAppConfigurationSchemaItem single configuration item inside an Android application's custom configuration schema.
type AndroidManagedStoreAppConfigurationSchemaItem struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewAndroidManagedStoreAppConfigurationSchemaItem instantiates a new androidManagedStoreAppConfigurationSchemaItem and sets the default values.
func NewAndroidManagedStoreAppConfigurationSchemaItem()(*AndroidManagedStoreAppConfigurationSchemaItem) {
    m := &AndroidManagedStoreAppConfigurationSchemaItem{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAndroidManagedStoreAppConfigurationSchemaItemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidManagedStoreAppConfigurationSchemaItemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidManagedStoreAppConfigurationSchemaItem(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AndroidManagedStoreAppConfigurationSchemaItem) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetBackingStore gets the backingStore property value. Stores model information.
func (m *AndroidManagedStoreAppConfigurationSchemaItem) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDataType gets the dataType property value. Data type for a configuration item inside an Android application's custom configuration schema
func (m *AndroidManagedStoreAppConfigurationSchemaItem) GetDataType()(*AndroidManagedStoreAppConfigurationSchemaItemDataType) {
    val, err := m.GetBackingStore().Get("dataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AndroidManagedStoreAppConfigurationSchemaItemDataType)
    }
    return nil
}
// GetDefaultBoolValue gets the defaultBoolValue property value. Default value for boolean type items, if specified by the app developer
func (m *AndroidManagedStoreAppConfigurationSchemaItem) GetDefaultBoolValue()(*bool) {
    val, err := m.GetBackingStore().Get("defaultBoolValue")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetDefaultIntValue gets the defaultIntValue property value. Default value for integer type items, if specified by the app developer
func (m *AndroidManagedStoreAppConfigurationSchemaItem) GetDefaultIntValue()(*int32) {
    val, err := m.GetBackingStore().Get("defaultIntValue")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetDefaultStringArrayValue gets the defaultStringArrayValue property value. Default value for string array type items, if specified by the app developer
func (m *AndroidManagedStoreAppConfigurationSchemaItem) GetDefaultStringArrayValue()([]string) {
    val, err := m.GetBackingStore().Get("defaultStringArrayValue")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetDefaultStringValue gets the defaultStringValue property value. Default value for string type items, if specified by the app developer
func (m *AndroidManagedStoreAppConfigurationSchemaItem) GetDefaultStringValue()(*string) {
    val, err := m.GetBackingStore().Get("defaultStringValue")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDescription gets the description property value. Description of what the item controls within the application
func (m *AndroidManagedStoreAppConfigurationSchemaItem) GetDescription()(*string) {
    val, err := m.GetBackingStore().Get("description")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDisplayName gets the displayName property value. Human readable name
func (m *AndroidManagedStoreAppConfigurationSchemaItem) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidManagedStoreAppConfigurationSchemaItem) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["dataType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidManagedStoreAppConfigurationSchemaItemDataType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDataType(val.(*AndroidManagedStoreAppConfigurationSchemaItemDataType))
        }
        return nil
    }
    res["defaultBoolValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultBoolValue(val)
        }
        return nil
    }
    res["defaultIntValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultIntValue(val)
        }
        return nil
    }
    res["defaultStringArrayValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
    res["defaultStringValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultStringValue(val)
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["index"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIndex(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["parentIndex"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParentIndex(val)
        }
        return nil
    }
    res["schemaItemKey"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSchemaItemKey(val)
        }
        return nil
    }
    res["selections"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
// GetIndex gets the index property value. Unique index the application uses to maintain nested schema items
func (m *AndroidManagedStoreAppConfigurationSchemaItem) GetIndex()(*int32) {
    val, err := m.GetBackingStore().Get("index")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AndroidManagedStoreAppConfigurationSchemaItem) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetParentIndex gets the parentIndex property value. Index of parent schema item to track nested schema items
func (m *AndroidManagedStoreAppConfigurationSchemaItem) GetParentIndex()(*int32) {
    val, err := m.GetBackingStore().Get("parentIndex")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetSchemaItemKey gets the schemaItemKey property value. Unique key the application uses to identify the item
func (m *AndroidManagedStoreAppConfigurationSchemaItem) GetSchemaItemKey()(*string) {
    val, err := m.GetBackingStore().Get("schemaItemKey")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSelections gets the selections property value. List of human readable name/value pairs for the valid values that can be set for this item (Choice and Multiselect items only)
func (m *AndroidManagedStoreAppConfigurationSchemaItem) GetSelections()([]KeyValuePairable) {
    val, err := m.GetBackingStore().Get("selections")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]KeyValuePairable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AndroidManagedStoreAppConfigurationSchemaItem) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteInt32Value("index", m.GetIndex())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
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
    if m.GetSelections() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSelections()))
        for i, v := range m.GetSelections() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
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
func (m *AndroidManagedStoreAppConfigurationSchemaItem) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *AndroidManagedStoreAppConfigurationSchemaItem) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDataType sets the dataType property value. Data type for a configuration item inside an Android application's custom configuration schema
func (m *AndroidManagedStoreAppConfigurationSchemaItem) SetDataType(value *AndroidManagedStoreAppConfigurationSchemaItemDataType)() {
    err := m.GetBackingStore().Set("dataType", value)
    if err != nil {
        panic(err)
    }
}
// SetDefaultBoolValue sets the defaultBoolValue property value. Default value for boolean type items, if specified by the app developer
func (m *AndroidManagedStoreAppConfigurationSchemaItem) SetDefaultBoolValue(value *bool)() {
    err := m.GetBackingStore().Set("defaultBoolValue", value)
    if err != nil {
        panic(err)
    }
}
// SetDefaultIntValue sets the defaultIntValue property value. Default value for integer type items, if specified by the app developer
func (m *AndroidManagedStoreAppConfigurationSchemaItem) SetDefaultIntValue(value *int32)() {
    err := m.GetBackingStore().Set("defaultIntValue", value)
    if err != nil {
        panic(err)
    }
}
// SetDefaultStringArrayValue sets the defaultStringArrayValue property value. Default value for string array type items, if specified by the app developer
func (m *AndroidManagedStoreAppConfigurationSchemaItem) SetDefaultStringArrayValue(value []string)() {
    err := m.GetBackingStore().Set("defaultStringArrayValue", value)
    if err != nil {
        panic(err)
    }
}
// SetDefaultStringValue sets the defaultStringValue property value. Default value for string type items, if specified by the app developer
func (m *AndroidManagedStoreAppConfigurationSchemaItem) SetDefaultStringValue(value *string)() {
    err := m.GetBackingStore().Set("defaultStringValue", value)
    if err != nil {
        panic(err)
    }
}
// SetDescription sets the description property value. Description of what the item controls within the application
func (m *AndroidManagedStoreAppConfigurationSchemaItem) SetDescription(value *string)() {
    err := m.GetBackingStore().Set("description", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. Human readable name
func (m *AndroidManagedStoreAppConfigurationSchemaItem) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetIndex sets the index property value. Unique index the application uses to maintain nested schema items
func (m *AndroidManagedStoreAppConfigurationSchemaItem) SetIndex(value *int32)() {
    err := m.GetBackingStore().Set("index", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AndroidManagedStoreAppConfigurationSchemaItem) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetParentIndex sets the parentIndex property value. Index of parent schema item to track nested schema items
func (m *AndroidManagedStoreAppConfigurationSchemaItem) SetParentIndex(value *int32)() {
    err := m.GetBackingStore().Set("parentIndex", value)
    if err != nil {
        panic(err)
    }
}
// SetSchemaItemKey sets the schemaItemKey property value. Unique key the application uses to identify the item
func (m *AndroidManagedStoreAppConfigurationSchemaItem) SetSchemaItemKey(value *string)() {
    err := m.GetBackingStore().Set("schemaItemKey", value)
    if err != nil {
        panic(err)
    }
}
// SetSelections sets the selections property value. List of human readable name/value pairs for the valid values that can be set for this item (Choice and Multiselect items only)
func (m *AndroidManagedStoreAppConfigurationSchemaItem) SetSelections(value []KeyValuePairable)() {
    err := m.GetBackingStore().Set("selections", value)
    if err != nil {
        panic(err)
    }
}
// AndroidManagedStoreAppConfigurationSchemaItemable 
type AndroidManagedStoreAppConfigurationSchemaItemable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetDataType()(*AndroidManagedStoreAppConfigurationSchemaItemDataType)
    GetDefaultBoolValue()(*bool)
    GetDefaultIntValue()(*int32)
    GetDefaultStringArrayValue()([]string)
    GetDefaultStringValue()(*string)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetIndex()(*int32)
    GetOdataType()(*string)
    GetParentIndex()(*int32)
    GetSchemaItemKey()(*string)
    GetSelections()([]KeyValuePairable)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetDataType(value *AndroidManagedStoreAppConfigurationSchemaItemDataType)()
    SetDefaultBoolValue(value *bool)()
    SetDefaultIntValue(value *int32)()
    SetDefaultStringArrayValue(value []string)()
    SetDefaultStringValue(value *string)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetIndex(value *int32)()
    SetOdataType(value *string)()
    SetParentIndex(value *int32)()
    SetSchemaItemKey(value *string)()
    SetSelections(value []KeyValuePairable)()
}

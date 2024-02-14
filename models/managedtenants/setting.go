package managedtenants

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type Setting struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewSetting instantiates a new Setting and sets the default values.
func NewSetting()(*Setting) {
    m := &Setting{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSettingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSettingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSetting(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Setting) GetAdditionalData()(map[string]any) {
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
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *Setting) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDisplayName gets the displayName property value. The display name for the setting. Required. Read-only.
// returns a *string when successful
func (m *Setting) GetDisplayName()(*string) {
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
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Setting) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["jsonValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJsonValue(val)
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
    res["overwriteAllowed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOverwriteAllowed(val)
        }
        return nil
    }
    res["settingId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingId(val)
        }
        return nil
    }
    res["valueType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagementParameterValueType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValueType(val.(*ManagementParameterValueType))
        }
        return nil
    }
    return res
}
// GetJsonValue gets the jsonValue property value. The value for the setting serialized as string of JSON. Required. Read-only.
// returns a *string when successful
func (m *Setting) GetJsonValue()(*string) {
    val, err := m.GetBackingStore().Get("jsonValue")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *Setting) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOverwriteAllowed gets the overwriteAllowed property value. A flag indicating whether the setting can be override existing configurations when applied. Required. Read-only.
// returns a *bool when successful
func (m *Setting) GetOverwriteAllowed()(*bool) {
    val, err := m.GetBackingStore().Get("overwriteAllowed")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetSettingId gets the settingId property value. The settingId property
// returns a *string when successful
func (m *Setting) GetSettingId()(*string) {
    val, err := m.GetBackingStore().Get("settingId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetValueType gets the valueType property value. The valueType property
// returns a *ManagementParameterValueType when successful
func (m *Setting) GetValueType()(*ManagementParameterValueType) {
    val, err := m.GetBackingStore().Get("valueType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ManagementParameterValueType)
    }
    return nil
}
// Serialize serializes information the current object
func (m *Setting) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("jsonValue", m.GetJsonValue())
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
        err := writer.WriteBoolValue("overwriteAllowed", m.GetOverwriteAllowed())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("settingId", m.GetSettingId())
        if err != nil {
            return err
        }
    }
    if m.GetValueType() != nil {
        cast := (*m.GetValueType()).String()
        err := writer.WriteStringValue("valueType", &cast)
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Setting) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *Setting) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDisplayName sets the displayName property value. The display name for the setting. Required. Read-only.
func (m *Setting) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetJsonValue sets the jsonValue property value. The value for the setting serialized as string of JSON. Required. Read-only.
func (m *Setting) SetJsonValue(value *string)() {
    err := m.GetBackingStore().Set("jsonValue", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *Setting) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetOverwriteAllowed sets the overwriteAllowed property value. A flag indicating whether the setting can be override existing configurations when applied. Required. Read-only.
func (m *Setting) SetOverwriteAllowed(value *bool)() {
    err := m.GetBackingStore().Set("overwriteAllowed", value)
    if err != nil {
        panic(err)
    }
}
// SetSettingId sets the settingId property value. The settingId property
func (m *Setting) SetSettingId(value *string)() {
    err := m.GetBackingStore().Set("settingId", value)
    if err != nil {
        panic(err)
    }
}
// SetValueType sets the valueType property value. The valueType property
func (m *Setting) SetValueType(value *ManagementParameterValueType)() {
    err := m.GetBackingStore().Set("valueType", value)
    if err != nil {
        panic(err)
    }
}
type Settingable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetDisplayName()(*string)
    GetJsonValue()(*string)
    GetOdataType()(*string)
    GetOverwriteAllowed()(*bool)
    GetSettingId()(*string)
    GetValueType()(*ManagementParameterValueType)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetDisplayName(value *string)()
    SetJsonValue(value *string)()
    SetOdataType(value *string)()
    SetOverwriteAllowed(value *bool)()
    SetSettingId(value *string)()
    SetValueType(value *ManagementParameterValueType)()
}

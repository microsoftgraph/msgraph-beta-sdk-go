package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type TodoSettings struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewTodoSettings instantiates a new TodoSettings and sets the default values.
func NewTodoSettings()(*TodoSettings) {
    m := &TodoSettings{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTodoSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTodoSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTodoSettings(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TodoSettings) GetAdditionalData()(map[string]any) {
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
func (m *TodoSettings) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TodoSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["isExternalJoinEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsExternalJoinEnabled(val)
        }
        return nil
    }
    res["isExternalShareEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsExternalShareEnabled(val)
        }
        return nil
    }
    res["isPushNotificationEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsPushNotificationEnabled(val)
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
    return res
}
// GetIsExternalJoinEnabled gets the isExternalJoinEnabled property value. Controls whether users can join lists from users external to your organization.
// returns a *bool when successful
func (m *TodoSettings) GetIsExternalJoinEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("isExternalJoinEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIsExternalShareEnabled gets the isExternalShareEnabled property value. Controls whether users can share lists with external users.
// returns a *bool when successful
func (m *TodoSettings) GetIsExternalShareEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("isExternalShareEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIsPushNotificationEnabled gets the isPushNotificationEnabled property value. Controls whether push notifications are enabled for your users.
// returns a *bool when successful
func (m *TodoSettings) GetIsPushNotificationEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("isPushNotificationEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *TodoSettings) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *TodoSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("isExternalJoinEnabled", m.GetIsExternalJoinEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isExternalShareEnabled", m.GetIsExternalShareEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isPushNotificationEnabled", m.GetIsPushNotificationEnabled())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TodoSettings) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *TodoSettings) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetIsExternalJoinEnabled sets the isExternalJoinEnabled property value. Controls whether users can join lists from users external to your organization.
func (m *TodoSettings) SetIsExternalJoinEnabled(value *bool)() {
    err := m.GetBackingStore().Set("isExternalJoinEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetIsExternalShareEnabled sets the isExternalShareEnabled property value. Controls whether users can share lists with external users.
func (m *TodoSettings) SetIsExternalShareEnabled(value *bool)() {
    err := m.GetBackingStore().Set("isExternalShareEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetIsPushNotificationEnabled sets the isPushNotificationEnabled property value. Controls whether push notifications are enabled for your users.
func (m *TodoSettings) SetIsPushNotificationEnabled(value *bool)() {
    err := m.GetBackingStore().Set("isPushNotificationEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *TodoSettings) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
type TodoSettingsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetIsExternalJoinEnabled()(*bool)
    GetIsExternalShareEnabled()(*bool)
    GetIsPushNotificationEnabled()(*bool)
    GetOdataType()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetIsExternalJoinEnabled(value *bool)()
    SetIsExternalShareEnabled(value *bool)()
    SetIsPushNotificationEnabled(value *bool)()
    SetOdataType(value *string)()
}

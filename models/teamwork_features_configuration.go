package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// TeamworkFeaturesConfiguration 
type TeamworkFeaturesConfiguration struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewTeamworkFeaturesConfiguration instantiates a new teamworkFeaturesConfiguration and sets the default values.
func NewTeamworkFeaturesConfiguration()(*TeamworkFeaturesConfiguration) {
    m := &TeamworkFeaturesConfiguration{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTeamworkFeaturesConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamworkFeaturesConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamworkFeaturesConfiguration(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamworkFeaturesConfiguration) GetAdditionalData()(map[string]any) {
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
func (m *TeamworkFeaturesConfiguration) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetEmailToSendLogsAndFeedback gets the emailToSendLogsAndFeedback property value. Email address to send logs and feedback.
func (m *TeamworkFeaturesConfiguration) GetEmailToSendLogsAndFeedback()(*string) {
    val, err := m.GetBackingStore().Get("emailToSendLogsAndFeedback")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamworkFeaturesConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["emailToSendLogsAndFeedback"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmailToSendLogsAndFeedback(val)
        }
        return nil
    }
    res["isAutoScreenShareEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsAutoScreenShareEnabled(val)
        }
        return nil
    }
    res["isBluetoothBeaconingEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsBluetoothBeaconingEnabled(val)
        }
        return nil
    }
    res["isHideMeetingNamesEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsHideMeetingNamesEnabled(val)
        }
        return nil
    }
    res["isSendLogsAndFeedbackEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsSendLogsAndFeedbackEnabled(val)
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
// GetIsAutoScreenShareEnabled gets the isAutoScreenShareEnabled property value. True if auto screen shared is enabled.
func (m *TeamworkFeaturesConfiguration) GetIsAutoScreenShareEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("isAutoScreenShareEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIsBluetoothBeaconingEnabled gets the isBluetoothBeaconingEnabled property value. True if Bluetooth beaconing is enabled.
func (m *TeamworkFeaturesConfiguration) GetIsBluetoothBeaconingEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("isBluetoothBeaconingEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIsHideMeetingNamesEnabled gets the isHideMeetingNamesEnabled property value. True if hiding meeting names is enabled.
func (m *TeamworkFeaturesConfiguration) GetIsHideMeetingNamesEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("isHideMeetingNamesEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIsSendLogsAndFeedbackEnabled gets the isSendLogsAndFeedbackEnabled property value. True if sending logs and feedback is enabled.
func (m *TeamworkFeaturesConfiguration) GetIsSendLogsAndFeedbackEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("isSendLogsAndFeedbackEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *TeamworkFeaturesConfiguration) GetOdataType()(*string) {
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
func (m *TeamworkFeaturesConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("emailToSendLogsAndFeedback", m.GetEmailToSendLogsAndFeedback())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isAutoScreenShareEnabled", m.GetIsAutoScreenShareEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isBluetoothBeaconingEnabled", m.GetIsBluetoothBeaconingEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isHideMeetingNamesEnabled", m.GetIsHideMeetingNamesEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isSendLogsAndFeedbackEnabled", m.GetIsSendLogsAndFeedbackEnabled())
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamworkFeaturesConfiguration) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *TeamworkFeaturesConfiguration) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetEmailToSendLogsAndFeedback sets the emailToSendLogsAndFeedback property value. Email address to send logs and feedback.
func (m *TeamworkFeaturesConfiguration) SetEmailToSendLogsAndFeedback(value *string)() {
    err := m.GetBackingStore().Set("emailToSendLogsAndFeedback", value)
    if err != nil {
        panic(err)
    }
}
// SetIsAutoScreenShareEnabled sets the isAutoScreenShareEnabled property value. True if auto screen shared is enabled.
func (m *TeamworkFeaturesConfiguration) SetIsAutoScreenShareEnabled(value *bool)() {
    err := m.GetBackingStore().Set("isAutoScreenShareEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetIsBluetoothBeaconingEnabled sets the isBluetoothBeaconingEnabled property value. True if Bluetooth beaconing is enabled.
func (m *TeamworkFeaturesConfiguration) SetIsBluetoothBeaconingEnabled(value *bool)() {
    err := m.GetBackingStore().Set("isBluetoothBeaconingEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetIsHideMeetingNamesEnabled sets the isHideMeetingNamesEnabled property value. True if hiding meeting names is enabled.
func (m *TeamworkFeaturesConfiguration) SetIsHideMeetingNamesEnabled(value *bool)() {
    err := m.GetBackingStore().Set("isHideMeetingNamesEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetIsSendLogsAndFeedbackEnabled sets the isSendLogsAndFeedbackEnabled property value. True if sending logs and feedback is enabled.
func (m *TeamworkFeaturesConfiguration) SetIsSendLogsAndFeedbackEnabled(value *bool)() {
    err := m.GetBackingStore().Set("isSendLogsAndFeedbackEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *TeamworkFeaturesConfiguration) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// TeamworkFeaturesConfigurationable 
type TeamworkFeaturesConfigurationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetEmailToSendLogsAndFeedback()(*string)
    GetIsAutoScreenShareEnabled()(*bool)
    GetIsBluetoothBeaconingEnabled()(*bool)
    GetIsHideMeetingNamesEnabled()(*bool)
    GetIsSendLogsAndFeedbackEnabled()(*bool)
    GetOdataType()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetEmailToSendLogsAndFeedback(value *string)()
    SetIsAutoScreenShareEnabled(value *bool)()
    SetIsBluetoothBeaconingEnabled(value *bool)()
    SetIsHideMeetingNamesEnabled(value *bool)()
    SetIsSendLogsAndFeedbackEnabled(value *bool)()
    SetOdataType(value *string)()
}

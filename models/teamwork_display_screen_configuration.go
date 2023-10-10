package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// TeamworkDisplayScreenConfiguration 
type TeamworkDisplayScreenConfiguration struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewTeamworkDisplayScreenConfiguration instantiates a new teamworkDisplayScreenConfiguration and sets the default values.
func NewTeamworkDisplayScreenConfiguration()(*TeamworkDisplayScreenConfiguration) {
    m := &TeamworkDisplayScreenConfiguration{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTeamworkDisplayScreenConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamworkDisplayScreenConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamworkDisplayScreenConfiguration(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamworkDisplayScreenConfiguration) GetAdditionalData()(map[string]any) {
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
func (m *TeamworkDisplayScreenConfiguration) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetBacklightBrightness gets the backlightBrightness property value. The brightness level on the device (0-100). Not applicable for Microsoft Teams Rooms devices.
func (m *TeamworkDisplayScreenConfiguration) GetBacklightBrightness()(*int32) {
    val, err := m.GetBackingStore().Get("backlightBrightness")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetBacklightTimeout gets the backlightTimeout property value. Timeout for backlight (30-3600 secs). Not applicable for Teams Rooms devices.
func (m *TeamworkDisplayScreenConfiguration) GetBacklightTimeout()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    val, err := m.GetBackingStore().Get("backlightTimeout")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamworkDisplayScreenConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["backlightBrightness"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBacklightBrightness(val)
        }
        return nil
    }
    res["backlightTimeout"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBacklightTimeout(val)
        }
        return nil
    }
    res["isHighContrastEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsHighContrastEnabled(val)
        }
        return nil
    }
    res["isScreensaverEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsScreensaverEnabled(val)
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
    res["screensaverTimeout"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScreensaverTimeout(val)
        }
        return nil
    }
    return res
}
// GetIsHighContrastEnabled gets the isHighContrastEnabled property value. True if high contrast mode is enabled. Not applicable for Teams Rooms devices.
func (m *TeamworkDisplayScreenConfiguration) GetIsHighContrastEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("isHighContrastEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIsScreensaverEnabled gets the isScreensaverEnabled property value. True if screensaver is enabled. Not applicable for Teams Rooms devices.
func (m *TeamworkDisplayScreenConfiguration) GetIsScreensaverEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("isScreensaverEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *TeamworkDisplayScreenConfiguration) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetScreensaverTimeout gets the screensaverTimeout property value. Screensaver timeout from 30 to 3600 secs. Not applicable for Teams Rooms devices.
func (m *TeamworkDisplayScreenConfiguration) GetScreensaverTimeout()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    val, err := m.GetBackingStore().Get("screensaverTimeout")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    }
    return nil
}
// Serialize serializes information the current object
func (m *TeamworkDisplayScreenConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("backlightBrightness", m.GetBacklightBrightness())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteISODurationValue("backlightTimeout", m.GetBacklightTimeout())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isHighContrastEnabled", m.GetIsHighContrastEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isScreensaverEnabled", m.GetIsScreensaverEnabled())
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
        err := writer.WriteISODurationValue("screensaverTimeout", m.GetScreensaverTimeout())
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
func (m *TeamworkDisplayScreenConfiguration) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *TeamworkDisplayScreenConfiguration) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetBacklightBrightness sets the backlightBrightness property value. The brightness level on the device (0-100). Not applicable for Microsoft Teams Rooms devices.
func (m *TeamworkDisplayScreenConfiguration) SetBacklightBrightness(value *int32)() {
    err := m.GetBackingStore().Set("backlightBrightness", value)
    if err != nil {
        panic(err)
    }
}
// SetBacklightTimeout sets the backlightTimeout property value. Timeout for backlight (30-3600 secs). Not applicable for Teams Rooms devices.
func (m *TeamworkDisplayScreenConfiguration) SetBacklightTimeout(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    err := m.GetBackingStore().Set("backlightTimeout", value)
    if err != nil {
        panic(err)
    }
}
// SetIsHighContrastEnabled sets the isHighContrastEnabled property value. True if high contrast mode is enabled. Not applicable for Teams Rooms devices.
func (m *TeamworkDisplayScreenConfiguration) SetIsHighContrastEnabled(value *bool)() {
    err := m.GetBackingStore().Set("isHighContrastEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetIsScreensaverEnabled sets the isScreensaverEnabled property value. True if screensaver is enabled. Not applicable for Teams Rooms devices.
func (m *TeamworkDisplayScreenConfiguration) SetIsScreensaverEnabled(value *bool)() {
    err := m.GetBackingStore().Set("isScreensaverEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *TeamworkDisplayScreenConfiguration) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetScreensaverTimeout sets the screensaverTimeout property value. Screensaver timeout from 30 to 3600 secs. Not applicable for Teams Rooms devices.
func (m *TeamworkDisplayScreenConfiguration) SetScreensaverTimeout(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    err := m.GetBackingStore().Set("screensaverTimeout", value)
    if err != nil {
        panic(err)
    }
}
// TeamworkDisplayScreenConfigurationable 
type TeamworkDisplayScreenConfigurationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetBacklightBrightness()(*int32)
    GetBacklightTimeout()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    GetIsHighContrastEnabled()(*bool)
    GetIsScreensaverEnabled()(*bool)
    GetOdataType()(*string)
    GetScreensaverTimeout()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetBacklightBrightness(value *int32)()
    SetBacklightTimeout(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)()
    SetIsHighContrastEnabled(value *bool)()
    SetIsScreensaverEnabled(value *bool)()
    SetOdataType(value *string)()
    SetScreensaverTimeout(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)()
}

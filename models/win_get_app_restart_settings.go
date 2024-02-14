package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// WinGetAppRestartSettings contains properties describing restart coordination following an app installation.
type WinGetAppRestartSettings struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewWinGetAppRestartSettings instantiates a new WinGetAppRestartSettings and sets the default values.
func NewWinGetAppRestartSettings()(*WinGetAppRestartSettings) {
    m := &WinGetAppRestartSettings{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateWinGetAppRestartSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateWinGetAppRestartSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWinGetAppRestartSettings(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *WinGetAppRestartSettings) GetAdditionalData()(map[string]any) {
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
func (m *WinGetAppRestartSettings) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetCountdownDisplayBeforeRestartInMinutes gets the countdownDisplayBeforeRestartInMinutes property value. The number of minutes before the restart time to display the countdown dialog for pending restarts.
// returns a *int32 when successful
func (m *WinGetAppRestartSettings) GetCountdownDisplayBeforeRestartInMinutes()(*int32) {
    val, err := m.GetBackingStore().Get("countdownDisplayBeforeRestartInMinutes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *WinGetAppRestartSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["countdownDisplayBeforeRestartInMinutes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCountdownDisplayBeforeRestartInMinutes(val)
        }
        return nil
    }
    res["gracePeriodInMinutes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGracePeriodInMinutes(val)
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
    res["restartNotificationSnoozeDurationInMinutes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRestartNotificationSnoozeDurationInMinutes(val)
        }
        return nil
    }
    return res
}
// GetGracePeriodInMinutes gets the gracePeriodInMinutes property value. The number of minutes to wait before restarting the device after an app installation.
// returns a *int32 when successful
func (m *WinGetAppRestartSettings) GetGracePeriodInMinutes()(*int32) {
    val, err := m.GetBackingStore().Get("gracePeriodInMinutes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *WinGetAppRestartSettings) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRestartNotificationSnoozeDurationInMinutes gets the restartNotificationSnoozeDurationInMinutes property value. The number of minutes to snooze the restart notification dialog when the snooze button is selected.
// returns a *int32 when successful
func (m *WinGetAppRestartSettings) GetRestartNotificationSnoozeDurationInMinutes()(*int32) {
    val, err := m.GetBackingStore().Get("restartNotificationSnoozeDurationInMinutes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// Serialize serializes information the current object
func (m *WinGetAppRestartSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("countdownDisplayBeforeRestartInMinutes", m.GetCountdownDisplayBeforeRestartInMinutes())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("gracePeriodInMinutes", m.GetGracePeriodInMinutes())
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
        err := writer.WriteInt32Value("restartNotificationSnoozeDurationInMinutes", m.GetRestartNotificationSnoozeDurationInMinutes())
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
func (m *WinGetAppRestartSettings) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *WinGetAppRestartSettings) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetCountdownDisplayBeforeRestartInMinutes sets the countdownDisplayBeforeRestartInMinutes property value. The number of minutes before the restart time to display the countdown dialog for pending restarts.
func (m *WinGetAppRestartSettings) SetCountdownDisplayBeforeRestartInMinutes(value *int32)() {
    err := m.GetBackingStore().Set("countdownDisplayBeforeRestartInMinutes", value)
    if err != nil {
        panic(err)
    }
}
// SetGracePeriodInMinutes sets the gracePeriodInMinutes property value. The number of minutes to wait before restarting the device after an app installation.
func (m *WinGetAppRestartSettings) SetGracePeriodInMinutes(value *int32)() {
    err := m.GetBackingStore().Set("gracePeriodInMinutes", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *WinGetAppRestartSettings) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetRestartNotificationSnoozeDurationInMinutes sets the restartNotificationSnoozeDurationInMinutes property value. The number of minutes to snooze the restart notification dialog when the snooze button is selected.
func (m *WinGetAppRestartSettings) SetRestartNotificationSnoozeDurationInMinutes(value *int32)() {
    err := m.GetBackingStore().Set("restartNotificationSnoozeDurationInMinutes", value)
    if err != nil {
        panic(err)
    }
}
type WinGetAppRestartSettingsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetCountdownDisplayBeforeRestartInMinutes()(*int32)
    GetGracePeriodInMinutes()(*int32)
    GetOdataType()(*string)
    GetRestartNotificationSnoozeDurationInMinutes()(*int32)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetCountdownDisplayBeforeRestartInMinutes(value *int32)()
    SetGracePeriodInMinutes(value *int32)()
    SetOdataType(value *string)()
    SetRestartNotificationSnoozeDurationInMinutes(value *int32)()
}

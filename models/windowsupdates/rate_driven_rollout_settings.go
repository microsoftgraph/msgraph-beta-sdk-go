package windowsupdates

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RateDrivenRolloutSettings 
type RateDrivenRolloutSettings struct {
    GradualRolloutSettings
}
// NewRateDrivenRolloutSettings instantiates a new rateDrivenRolloutSettings and sets the default values.
func NewRateDrivenRolloutSettings()(*RateDrivenRolloutSettings) {
    m := &RateDrivenRolloutSettings{
        GradualRolloutSettings: *NewGradualRolloutSettings(),
    }
    odataTypeValue := "#microsoft.graph.windowsUpdates.rateDrivenRolloutSettings"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateRateDrivenRolloutSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRateDrivenRolloutSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRateDrivenRolloutSettings(), nil
}
// GetDevicesPerOffer gets the devicesPerOffer property value. Specifies the number of devices that are offered at the same time. When not set, all devices in the deployment are offered content at the same time.
func (m *RateDrivenRolloutSettings) GetDevicesPerOffer()(*int32) {
    val, err := m.GetBackingStore().Get("devicesPerOffer")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RateDrivenRolloutSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.GradualRolloutSettings.GetFieldDeserializers()
    res["devicesPerOffer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDevicesPerOffer(val)
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
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *RateDrivenRolloutSettings) GetOdataType()(*string) {
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
func (m *RateDrivenRolloutSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.GradualRolloutSettings.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("devicesPerOffer", m.GetDevicesPerOffer())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDevicesPerOffer sets the devicesPerOffer property value. Specifies the number of devices that are offered at the same time. When not set, all devices in the deployment are offered content at the same time.
func (m *RateDrivenRolloutSettings) SetDevicesPerOffer(value *int32)() {
    err := m.GetBackingStore().Set("devicesPerOffer", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *RateDrivenRolloutSettings) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// RateDrivenRolloutSettingsable 
type RateDrivenRolloutSettingsable interface {
    GradualRolloutSettingsable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDevicesPerOffer()(*int32)
    GetOdataType()(*string)
    SetDevicesPerOffer(value *int32)()
    SetOdataType(value *string)()
}

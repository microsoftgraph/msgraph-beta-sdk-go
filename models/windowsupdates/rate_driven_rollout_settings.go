package windowsupdates

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RateDrivenRolloutSettings struct {
    GradualRolloutSettings
}
// NewRateDrivenRolloutSettings instantiates a new RateDrivenRolloutSettings and sets the default values.
func NewRateDrivenRolloutSettings()(*RateDrivenRolloutSettings) {
    m := &RateDrivenRolloutSettings{
        GradualRolloutSettings: *NewGradualRolloutSettings(),
    }
    odataTypeValue := "#microsoft.graph.windowsUpdates.rateDrivenRolloutSettings"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateRateDrivenRolloutSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRateDrivenRolloutSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRateDrivenRolloutSettings(), nil
}
// GetDevicesPerOffer gets the devicesPerOffer property value. Specifies the number of devices that are offered at the same time. When not set, all devices in the deployment are offered content at the same time.
// returns a *int32 when successful
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
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
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
    return res
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
    return nil
}
// SetDevicesPerOffer sets the devicesPerOffer property value. Specifies the number of devices that are offered at the same time. When not set, all devices in the deployment are offered content at the same time.
func (m *RateDrivenRolloutSettings) SetDevicesPerOffer(value *int32)() {
    err := m.GetBackingStore().Set("devicesPerOffer", value)
    if err != nil {
        panic(err)
    }
}
type RateDrivenRolloutSettingsable interface {
    GradualRolloutSettingsable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDevicesPerOffer()(*int32)
    SetDevicesPerOffer(value *int32)()
}

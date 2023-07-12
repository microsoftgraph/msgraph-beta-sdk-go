package windowsupdates

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DurationDrivenRolloutSettings 
type DurationDrivenRolloutSettings struct {
    GradualRolloutSettings
    // The OdataType property
    OdataType *string
}
// NewDurationDrivenRolloutSettings instantiates a new durationDrivenRolloutSettings and sets the default values.
func NewDurationDrivenRolloutSettings()(*DurationDrivenRolloutSettings) {
    m := &DurationDrivenRolloutSettings{
        GradualRolloutSettings: *NewGradualRolloutSettings(),
    }
    odataTypeValue := "#microsoft.graph.windowsUpdates.durationDrivenRolloutSettings"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateDurationDrivenRolloutSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDurationDrivenRolloutSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDurationDrivenRolloutSettings(), nil
}
// GetDurationUntilDeploymentEnd gets the durationUntilDeploymentEnd property value. The target duration of the rollout. Given durationBetweenOffers and durationUntilDeploymentEnd, the system will automatically calculate how many devices are in each offering.
func (m *DurationDrivenRolloutSettings) GetDurationUntilDeploymentEnd()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    val, err := m.GetBackingStore().Get("durationUntilDeploymentEnd")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DurationDrivenRolloutSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.GradualRolloutSettings.GetFieldDeserializers()
    res["durationUntilDeploymentEnd"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDurationUntilDeploymentEnd(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *DurationDrivenRolloutSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.GradualRolloutSettings.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteISODurationValue("durationUntilDeploymentEnd", m.GetDurationUntilDeploymentEnd())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDurationUntilDeploymentEnd sets the durationUntilDeploymentEnd property value. The target duration of the rollout. Given durationBetweenOffers and durationUntilDeploymentEnd, the system will automatically calculate how many devices are in each offering.
func (m *DurationDrivenRolloutSettings) SetDurationUntilDeploymentEnd(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    err := m.GetBackingStore().Set("durationUntilDeploymentEnd", value)
    if err != nil {
        panic(err)
    }
}
// DurationDrivenRolloutSettingsable 
type DurationDrivenRolloutSettingsable interface {
    GradualRolloutSettingsable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDurationUntilDeploymentEnd()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    SetDurationUntilDeploymentEnd(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)()
}

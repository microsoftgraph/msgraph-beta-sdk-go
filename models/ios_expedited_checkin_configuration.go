package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IosExpeditedCheckinConfiguration experimental profile to increase the rate of device check-ins per day of iOS devices. This profile type is deprecated.
type IosExpeditedCheckinConfiguration struct {
    AppleExpeditedCheckinConfigurationBase
}
// NewIosExpeditedCheckinConfiguration instantiates a new iosExpeditedCheckinConfiguration and sets the default values.
func NewIosExpeditedCheckinConfiguration()(*IosExpeditedCheckinConfiguration) {
    m := &IosExpeditedCheckinConfiguration{
        AppleExpeditedCheckinConfigurationBase: *NewAppleExpeditedCheckinConfigurationBase(),
    }
    odataTypeValue := "#microsoft.graph.iosExpeditedCheckinConfiguration"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateIosExpeditedCheckinConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIosExpeditedCheckinConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIosExpeditedCheckinConfiguration(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IosExpeditedCheckinConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AppleExpeditedCheckinConfigurationBase.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *IosExpeditedCheckinConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AppleExpeditedCheckinConfigurationBase.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// IosExpeditedCheckinConfigurationable 
type IosExpeditedCheckinConfigurationable interface {
    AppleExpeditedCheckinConfigurationBaseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}

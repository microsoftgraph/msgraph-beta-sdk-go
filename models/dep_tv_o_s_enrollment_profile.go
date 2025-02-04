package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DepTvOSEnrollmentProfile the depTvOSEnrollmentProfile resource represents an Apple Device Enrollment Program (DEP) enrollment profile specific to Apple TV device configuration. This type of profile must be assigned to Apple TV devices before the devices can enroll via DEP. However, This entity type will only be used as a navigation property to fetch the display name of the profile while getting the exitsing depOnboardingSetting entity, it won't support any operations, as the new entity is supported in device configuration(DCV2) graph calls
type DepTvOSEnrollmentProfile struct {
    EnrollmentProfile
}
// NewDepTvOSEnrollmentProfile instantiates a new DepTvOSEnrollmentProfile and sets the default values.
func NewDepTvOSEnrollmentProfile()(*DepTvOSEnrollmentProfile) {
    m := &DepTvOSEnrollmentProfile{
        EnrollmentProfile: *NewEnrollmentProfile(),
    }
    odataTypeValue := "#microsoft.graph.depTvOSEnrollmentProfile"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateDepTvOSEnrollmentProfileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDepTvOSEnrollmentProfileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDepTvOSEnrollmentProfile(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DepTvOSEnrollmentProfile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.EnrollmentProfile.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *DepTvOSEnrollmentProfile) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.EnrollmentProfile.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
type DepTvOSEnrollmentProfileable interface {
    EnrollmentProfileable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}

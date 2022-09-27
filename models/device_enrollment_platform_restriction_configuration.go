package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceEnrollmentPlatformRestrictionConfiguration 
type DeviceEnrollmentPlatformRestrictionConfiguration struct {
    DeviceEnrollmentConfiguration
    // Restrictions based on platform, platform operating system version, and device ownership
    platformRestriction DeviceEnrollmentPlatformRestrictionable
    // This enum indicates the platform type for which the enrollment restriction applies.
    platformType *EnrollmentRestrictionPlatformType
}
// NewDeviceEnrollmentPlatformRestrictionConfiguration instantiates a new DeviceEnrollmentPlatformRestrictionConfiguration and sets the default values.
func NewDeviceEnrollmentPlatformRestrictionConfiguration()(*DeviceEnrollmentPlatformRestrictionConfiguration) {
    m := &DeviceEnrollmentPlatformRestrictionConfiguration{
        DeviceEnrollmentConfiguration: *NewDeviceEnrollmentConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.deviceEnrollmentPlatformRestrictionConfiguration";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateDeviceEnrollmentPlatformRestrictionConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceEnrollmentPlatformRestrictionConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceEnrollmentPlatformRestrictionConfiguration(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceEnrollmentPlatformRestrictionConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceEnrollmentConfiguration.GetFieldDeserializers()
    res["platformRestriction"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateDeviceEnrollmentPlatformRestrictionFromDiscriminatorValue , m.SetPlatformRestriction)
    res["platformType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseEnrollmentRestrictionPlatformType , m.SetPlatformType)
    return res
}
// GetPlatformRestriction gets the platformRestriction property value. Restrictions based on platform, platform operating system version, and device ownership
func (m *DeviceEnrollmentPlatformRestrictionConfiguration) GetPlatformRestriction()(DeviceEnrollmentPlatformRestrictionable) {
    return m.platformRestriction
}
// GetPlatformType gets the platformType property value. This enum indicates the platform type for which the enrollment restriction applies.
func (m *DeviceEnrollmentPlatformRestrictionConfiguration) GetPlatformType()(*EnrollmentRestrictionPlatformType) {
    return m.platformType
}
// Serialize serializes information the current object
func (m *DeviceEnrollmentPlatformRestrictionConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceEnrollmentConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("platformRestriction", m.GetPlatformRestriction())
        if err != nil {
            return err
        }
    }
    if m.GetPlatformType() != nil {
        cast := (*m.GetPlatformType()).String()
        err = writer.WriteStringValue("platformType", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetPlatformRestriction sets the platformRestriction property value. Restrictions based on platform, platform operating system version, and device ownership
func (m *DeviceEnrollmentPlatformRestrictionConfiguration) SetPlatformRestriction(value DeviceEnrollmentPlatformRestrictionable)() {
    m.platformRestriction = value
}
// SetPlatformType sets the platformType property value. This enum indicates the platform type for which the enrollment restriction applies.
func (m *DeviceEnrollmentPlatformRestrictionConfiguration) SetPlatformType(value *EnrollmentRestrictionPlatformType)() {
    m.platformType = value
}

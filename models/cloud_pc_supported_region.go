package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CloudPcSupportedRegion 
type CloudPcSupportedRegion struct {
    Entity
    // The name for the supported region. Read-only.
    displayName *string
    // The regionGroup property
    regionGroup *CloudPcRegionGroup
    // The status of the supported region. Possible values are: available, restricted, unavailable, unknownFutureValue. Read-only.
    regionStatus *CloudPcSupportedRegionStatus
    // The supportedSolution property
    supportedSolution *CloudPcManagementService
}
// NewCloudPcSupportedRegion instantiates a new CloudPcSupportedRegion and sets the default values.
func NewCloudPcSupportedRegion()(*CloudPcSupportedRegion) {
    m := &CloudPcSupportedRegion{
        Entity: *NewEntity(),
    }
    return m
}
// CreateCloudPcSupportedRegionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCloudPcSupportedRegionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudPcSupportedRegion(), nil
}
// GetDisplayName gets the displayName property value. The name for the supported region. Read-only.
func (m *CloudPcSupportedRegion) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CloudPcSupportedRegion) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["regionGroup"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseCloudPcRegionGroup , m.SetRegionGroup)
    res["regionStatus"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseCloudPcSupportedRegionStatus , m.SetRegionStatus)
    res["supportedSolution"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseCloudPcManagementService , m.SetSupportedSolution)
    return res
}
// GetRegionGroup gets the regionGroup property value. The regionGroup property
func (m *CloudPcSupportedRegion) GetRegionGroup()(*CloudPcRegionGroup) {
    return m.regionGroup
}
// GetRegionStatus gets the regionStatus property value. The status of the supported region. Possible values are: available, restricted, unavailable, unknownFutureValue. Read-only.
func (m *CloudPcSupportedRegion) GetRegionStatus()(*CloudPcSupportedRegionStatus) {
    return m.regionStatus
}
// GetSupportedSolution gets the supportedSolution property value. The supportedSolution property
func (m *CloudPcSupportedRegion) GetSupportedSolution()(*CloudPcManagementService) {
    return m.supportedSolution
}
// Serialize serializes information the current object
func (m *CloudPcSupportedRegion) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetRegionGroup() != nil {
        cast := (*m.GetRegionGroup()).String()
        err = writer.WriteStringValue("regionGroup", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetRegionStatus() != nil {
        cast := (*m.GetRegionStatus()).String()
        err = writer.WriteStringValue("regionStatus", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetSupportedSolution() != nil {
        cast := (*m.GetSupportedSolution()).String()
        err = writer.WriteStringValue("supportedSolution", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDisplayName sets the displayName property value. The name for the supported region. Read-only.
func (m *CloudPcSupportedRegion) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetRegionGroup sets the regionGroup property value. The regionGroup property
func (m *CloudPcSupportedRegion) SetRegionGroup(value *CloudPcRegionGroup)() {
    m.regionGroup = value
}
// SetRegionStatus sets the regionStatus property value. The status of the supported region. Possible values are: available, restricted, unavailable, unknownFutureValue. Read-only.
func (m *CloudPcSupportedRegion) SetRegionStatus(value *CloudPcSupportedRegionStatus)() {
    m.regionStatus = value
}
// SetSupportedSolution sets the supportedSolution property value. The supportedSolution property
func (m *CloudPcSupportedRegion) SetSupportedSolution(value *CloudPcManagementService)() {
    m.supportedSolution = value
}

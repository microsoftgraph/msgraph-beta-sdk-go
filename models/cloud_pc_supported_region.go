package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CloudPcSupportedRegion 
type CloudPcSupportedRegion struct {
    Entity
}
// NewCloudPcSupportedRegion instantiates a new cloudPcSupportedRegion and sets the default values.
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
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CloudPcSupportedRegion) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["regionGroup"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcRegionGroup)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRegionGroup(val.(*CloudPcRegionGroup))
        }
        return nil
    }
    res["regionStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcSupportedRegion_regionStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRegionStatus(val.(*CloudPcSupportedRegion_regionStatus))
        }
        return nil
    }
    res["supportedSolution"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcManagementService)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSupportedSolution(val.(*CloudPcManagementService))
        }
        return nil
    }
    return res
}
// GetRegionGroup gets the regionGroup property value. The regionGroup property
func (m *CloudPcSupportedRegion) GetRegionGroup()(*CloudPcRegionGroup) {
    val, err := m.GetBackingStore().Get("regionGroup")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudPcRegionGroup)
    }
    return nil
}
// GetRegionStatus gets the regionStatus property value. The status of the supported region. Possible values are: available, restricted, unavailable, unknownFutureValue. Read-only.
func (m *CloudPcSupportedRegion) GetRegionStatus()(*CloudPcSupportedRegion_regionStatus) {
    val, err := m.GetBackingStore().Get("regionStatus")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudPcSupportedRegion_regionStatus)
    }
    return nil
}
// GetSupportedSolution gets the supportedSolution property value. The supportedSolution property
func (m *CloudPcSupportedRegion) GetSupportedSolution()(*CloudPcManagementService) {
    val, err := m.GetBackingStore().Get("supportedSolution")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudPcManagementService)
    }
    return nil
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
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetRegionGroup sets the regionGroup property value. The regionGroup property
func (m *CloudPcSupportedRegion) SetRegionGroup(value *CloudPcRegionGroup)() {
    err := m.GetBackingStore().Set("regionGroup", value)
    if err != nil {
        panic(err)
    }
}
// SetRegionStatus sets the regionStatus property value. The status of the supported region. Possible values are: available, restricted, unavailable, unknownFutureValue. Read-only.
func (m *CloudPcSupportedRegion) SetRegionStatus(value *CloudPcSupportedRegion_regionStatus)() {
    err := m.GetBackingStore().Set("regionStatus", value)
    if err != nil {
        panic(err)
    }
}
// SetSupportedSolution sets the supportedSolution property value. The supportedSolution property
func (m *CloudPcSupportedRegion) SetSupportedSolution(value *CloudPcManagementService)() {
    err := m.GetBackingStore().Set("supportedSolution", value)
    if err != nil {
        panic(err)
    }
}
// CloudPcSupportedRegionable 
type CloudPcSupportedRegionable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDisplayName()(*string)
    GetRegionGroup()(*CloudPcRegionGroup)
    GetRegionStatus()(*CloudPcSupportedRegion_regionStatus)
    GetSupportedSolution()(*CloudPcManagementService)
    SetDisplayName(value *string)()
    SetRegionGroup(value *CloudPcRegionGroup)()
    SetRegionStatus(value *CloudPcSupportedRegion_regionStatus)()
    SetSupportedSolution(value *CloudPcManagementService)()
}

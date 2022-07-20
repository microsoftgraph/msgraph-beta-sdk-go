package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CloudPcSupportedRegion 
type CloudPcSupportedRegion struct {
    Entity
    // The name for the supported region. Read-only.
    displayName *string
    // The regionStatus property
    regionStatus *CloudPcSupportedRegionStatus
}
// NewCloudPcSupportedRegion instantiates a new CloudPcSupportedRegion and sets the default values.
func NewCloudPcSupportedRegion()(*CloudPcSupportedRegion) {
    m := &CloudPcSupportedRegion{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.cloudPcSupportedRegion";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateCloudPcSupportedRegionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCloudPcSupportedRegionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudPcSupportedRegion(), nil
}
// GetDisplayName gets the displayName property value. The name for the supported region. Read-only.
func (m *CloudPcSupportedRegion) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
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
    res["regionStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcSupportedRegionStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRegionStatus(val.(*CloudPcSupportedRegionStatus))
        }
        return nil
    }
    return res
}
// GetRegionStatus gets the regionStatus property value. The regionStatus property
func (m *CloudPcSupportedRegion) GetRegionStatus()(*CloudPcSupportedRegionStatus) {
    if m == nil {
        return nil
    } else {
        return m.regionStatus
    }
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
    if m.GetRegionStatus() != nil {
        cast := (*m.GetRegionStatus()).String()
        err = writer.WriteStringValue("regionStatus", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDisplayName sets the displayName property value. The name for the supported region. Read-only.
func (m *CloudPcSupportedRegion) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetRegionStatus sets the regionStatus property value. The regionStatus property
func (m *CloudPcSupportedRegion) SetRegionStatus(value *CloudPcSupportedRegionStatus)() {
    if m != nil {
        m.regionStatus = value
    }
}

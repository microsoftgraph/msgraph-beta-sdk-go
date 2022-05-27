package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CloudPcSupportedRegion provides operations to manage the deviceManagement singleton.
type CloudPcSupportedRegion struct {
    Entity
    // The name for the supported region. Read-only.
    displayName *string
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
    return res
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
    return nil
}
// SetDisplayName sets the displayName property value. The name for the supported region. Read-only.
func (m *CloudPcSupportedRegion) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}

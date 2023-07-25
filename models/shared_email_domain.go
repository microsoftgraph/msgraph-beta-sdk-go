package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SharedEmailDomain 
type SharedEmailDomain struct {
    Entity
}
// NewSharedEmailDomain instantiates a new sharedEmailDomain and sets the default values.
func NewSharedEmailDomain()(*SharedEmailDomain) {
    m := &SharedEmailDomain{
        Entity: *NewEntity(),
    }
    return m
}
// CreateSharedEmailDomainFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSharedEmailDomainFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSharedEmailDomain(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SharedEmailDomain) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["provisioningStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProvisioningStatus(val)
        }
        return nil
    }
    return res
}
// GetProvisioningStatus gets the provisioningStatus property value. The provisioningStatus property
func (m *SharedEmailDomain) GetProvisioningStatus()(*string) {
    val, err := m.GetBackingStore().Get("provisioningStatus")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *SharedEmailDomain) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("provisioningStatus", m.GetProvisioningStatus())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetProvisioningStatus sets the provisioningStatus property value. The provisioningStatus property
func (m *SharedEmailDomain) SetProvisioningStatus(value *string)() {
    err := m.GetBackingStore().Set("provisioningStatus", value)
    if err != nil {
        panic(err)
    }
}
// SharedEmailDomainable 
type SharedEmailDomainable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetProvisioningStatus()(*string)
    SetProvisioningStatus(value *string)()
}

package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SharedEmailDomain provides operations to manage the collection of accessReview entities.
type SharedEmailDomain struct {
    Entity
    // The provisioningStatus property
    provisioningStatus *string
}
// NewSharedEmailDomain instantiates a new sharedEmailDomain and sets the default values.
func NewSharedEmailDomain()(*SharedEmailDomain) {
    m := &SharedEmailDomain{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.sharedEmailDomain";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateSharedEmailDomainFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSharedEmailDomainFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSharedEmailDomain(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SharedEmailDomain) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["provisioningStatus"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetProvisioningStatus)
    return res
}
// GetProvisioningStatus gets the provisioningStatus property value. The provisioningStatus property
func (m *SharedEmailDomain) GetProvisioningStatus()(*string) {
    return m.provisioningStatus
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
    m.provisioningStatus = value
}

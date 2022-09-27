package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CommunicationsUserIdentity 
type CommunicationsUserIdentity struct {
    Identity
    // The user's tenant ID.
    tenantId *string
}
// NewCommunicationsUserIdentity instantiates a new CommunicationsUserIdentity and sets the default values.
func NewCommunicationsUserIdentity()(*CommunicationsUserIdentity) {
    m := &CommunicationsUserIdentity{
        Identity: *NewIdentity(),
    }
    odataTypeValue := "#microsoft.graph.communicationsUserIdentity";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateCommunicationsUserIdentityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCommunicationsUserIdentityFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCommunicationsUserIdentity(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CommunicationsUserIdentity) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Identity.GetFieldDeserializers()
    res["tenantId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetTenantId)
    return res
}
// GetTenantId gets the tenantId property value. The user's tenant ID.
func (m *CommunicationsUserIdentity) GetTenantId()(*string) {
    return m.tenantId
}
// Serialize serializes information the current object
func (m *CommunicationsUserIdentity) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Identity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("tenantId", m.GetTenantId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetTenantId sets the tenantId property value. The user's tenant ID.
func (m *CommunicationsUserIdentity) SetTenantId(value *string)() {
    m.tenantId = value
}

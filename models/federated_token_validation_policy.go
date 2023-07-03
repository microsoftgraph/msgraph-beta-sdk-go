package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// FederatedTokenValidationPolicy 
type FederatedTokenValidationPolicy struct {
    DirectoryObject
}
// NewFederatedTokenValidationPolicy instantiates a new FederatedTokenValidationPolicy and sets the default values.
func NewFederatedTokenValidationPolicy()(*FederatedTokenValidationPolicy) {
    m := &FederatedTokenValidationPolicy{
        DirectoryObject: *NewDirectoryObject(),
    }
    odataTypeValue := "#microsoft.graph.federatedTokenValidationPolicy"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateFederatedTokenValidationPolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateFederatedTokenValidationPolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFederatedTokenValidationPolicy(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *FederatedTokenValidationPolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DirectoryObject.GetFieldDeserializers()
    res["validatingDomains"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateValidatingDomainsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValidatingDomains(val.(ValidatingDomainsable))
        }
        return nil
    }
    return res
}
// GetValidatingDomains gets the validatingDomains property value. The validatingDomains property
func (m *FederatedTokenValidationPolicy) GetValidatingDomains()(ValidatingDomainsable) {
    val, err := m.GetBackingStore().Get("validatingDomains")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ValidatingDomainsable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *FederatedTokenValidationPolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DirectoryObject.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("validatingDomains", m.GetValidatingDomains())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValidatingDomains sets the validatingDomains property value. The validatingDomains property
func (m *FederatedTokenValidationPolicy) SetValidatingDomains(value ValidatingDomainsable)() {
    err := m.GetBackingStore().Set("validatingDomains", value)
    if err != nil {
        panic(err)
    }
}
// FederatedTokenValidationPolicyable 
type FederatedTokenValidationPolicyable interface {
    DirectoryObjectable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetValidatingDomains()(ValidatingDomainsable)
    SetValidatingDomains(value ValidatingDomainsable)()
}

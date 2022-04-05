package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TenantRelationship provides operations to manage the tenantRelationship singleton.
type TenantRelationship struct {
    Entity
    // The customer who has a delegated admin relationship with a Microsoft partner.
    delegatedAdminCustomers []DelegatedAdminCustomerable;
    // The details of the delegated administrative privileges that a Microsoft partner has in a customer tenant.
    delegatedAdminRelationships []DelegatedAdminRelationshipable;
}
// NewTenantRelationship instantiates a new tenantRelationship and sets the default values.
func NewTenantRelationship()(*TenantRelationship) {
    m := &TenantRelationship{
        Entity: *NewEntity(),
    }
    return m
}
// CreateTenantRelationshipFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTenantRelationshipFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTenantRelationship(), nil
}
// GetDelegatedAdminCustomers gets the delegatedAdminCustomers property value. The customer who has a delegated admin relationship with a Microsoft partner.
func (m *TenantRelationship) GetDelegatedAdminCustomers()([]DelegatedAdminCustomerable) {
    if m == nil {
        return nil
    } else {
        return m.delegatedAdminCustomers
    }
}
// GetDelegatedAdminRelationships gets the delegatedAdminRelationships property value. The details of the delegated administrative privileges that a Microsoft partner has in a customer tenant.
func (m *TenantRelationship) GetDelegatedAdminRelationships()([]DelegatedAdminRelationshipable) {
    if m == nil {
        return nil
    } else {
        return m.delegatedAdminRelationships
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TenantRelationship) GetFieldDeserializers()(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["delegatedAdminCustomers"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDelegatedAdminCustomerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DelegatedAdminCustomerable, len(val))
            for i, v := range val {
                res[i] = v.(DelegatedAdminCustomerable)
            }
            m.SetDelegatedAdminCustomers(res)
        }
        return nil
    }
    res["delegatedAdminRelationships"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDelegatedAdminRelationshipFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DelegatedAdminRelationshipable, len(val))
            for i, v := range val {
                res[i] = v.(DelegatedAdminRelationshipable)
            }
            m.SetDelegatedAdminRelationships(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *TenantRelationship) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetDelegatedAdminCustomers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDelegatedAdminCustomers()))
        for i, v := range m.GetDelegatedAdminCustomers() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("delegatedAdminCustomers", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDelegatedAdminRelationships() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDelegatedAdminRelationships()))
        for i, v := range m.GetDelegatedAdminRelationships() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("delegatedAdminRelationships", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDelegatedAdminCustomers sets the delegatedAdminCustomers property value. The customer who has a delegated admin relationship with a Microsoft partner.
func (m *TenantRelationship) SetDelegatedAdminCustomers(value []DelegatedAdminCustomerable)() {
    if m != nil {
        m.delegatedAdminCustomers = value
    }
}
// SetDelegatedAdminRelationships sets the delegatedAdminRelationships property value. The details of the delegated administrative privileges that a Microsoft partner has in a customer tenant.
func (m *TenantRelationship) SetDelegatedAdminRelationships(value []DelegatedAdminRelationshipable)() {
    if m != nil {
        m.delegatedAdminRelationships = value
    }
}

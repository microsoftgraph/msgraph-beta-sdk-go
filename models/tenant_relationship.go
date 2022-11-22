package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TenantRelationship 
type TenantRelationship struct {
    Entity
    // The customer who has a delegated admin relationship with a Microsoft partner.
    delegatedAdminCustomers []DelegatedAdminCustomerable
    // The details of the delegated administrative privileges that a Microsoft partner has in a customer tenant.
    delegatedAdminRelationships []DelegatedAdminRelationshipable
}
// NewTenantRelationship instantiates a new TenantRelationship and sets the default values.
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
    return m.delegatedAdminCustomers
}
// GetDelegatedAdminRelationships gets the delegatedAdminRelationships property value. The details of the delegated administrative privileges that a Microsoft partner has in a customer tenant.
func (m *TenantRelationship) GetDelegatedAdminRelationships()([]DelegatedAdminRelationshipable) {
    return m.delegatedAdminRelationships
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TenantRelationship) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["delegatedAdminCustomers"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDelegatedAdminCustomerFromDiscriminatorValue , m.SetDelegatedAdminCustomers)
    res["delegatedAdminRelationships"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDelegatedAdminRelationshipFromDiscriminatorValue , m.SetDelegatedAdminRelationships)
    return res
}
// Serialize serializes information the current object
func (m *TenantRelationship) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetDelegatedAdminCustomers() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetDelegatedAdminCustomers())
        err = writer.WriteCollectionOfObjectValues("delegatedAdminCustomers", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDelegatedAdminRelationships() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetDelegatedAdminRelationships())
        err = writer.WriteCollectionOfObjectValues("delegatedAdminRelationships", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDelegatedAdminCustomers sets the delegatedAdminCustomers property value. The customer who has a delegated admin relationship with a Microsoft partner.
func (m *TenantRelationship) SetDelegatedAdminCustomers(value []DelegatedAdminCustomerable)() {
    m.delegatedAdminCustomers = value
}
// SetDelegatedAdminRelationships sets the delegatedAdminRelationships property value. The details of the delegated administrative privileges that a Microsoft partner has in a customer tenant.
func (m *TenantRelationship) SetDelegatedAdminRelationships(value []DelegatedAdminRelationshipable)() {
    m.delegatedAdminRelationships = value
}

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DelegatedAdminCustomer provides operations to manage the collection of activityStatistics entities.
type DelegatedAdminCustomer struct {
    Entity
    // The Azure AD display name of the customer tenant. Read-only. Supports $orderBy.
    displayName *string
    // Contains the management details of a service in the customer tenant that's managed by delegated administration.
    serviceManagementDetails []DelegatedAdminServiceManagementDetailable
    // The Azure AD-assigned tenant ID of the customer. Read-only.
    tenantId *string
}
// NewDelegatedAdminCustomer instantiates a new delegatedAdminCustomer and sets the default values.
func NewDelegatedAdminCustomer()(*DelegatedAdminCustomer) {
    m := &DelegatedAdminCustomer{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.delegatedAdminCustomer";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateDelegatedAdminCustomerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDelegatedAdminCustomerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDelegatedAdminCustomer(), nil
}
// GetDisplayName gets the displayName property value. The Azure AD display name of the customer tenant. Read-only. Supports $orderBy.
func (m *DelegatedAdminCustomer) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DelegatedAdminCustomer) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["serviceManagementDetails"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDelegatedAdminServiceManagementDetailFromDiscriminatorValue , m.SetServiceManagementDetails)
    res["tenantId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetTenantId)
    return res
}
// GetServiceManagementDetails gets the serviceManagementDetails property value. Contains the management details of a service in the customer tenant that's managed by delegated administration.
func (m *DelegatedAdminCustomer) GetServiceManagementDetails()([]DelegatedAdminServiceManagementDetailable) {
    return m.serviceManagementDetails
}
// GetTenantId gets the tenantId property value. The Azure AD-assigned tenant ID of the customer. Read-only.
func (m *DelegatedAdminCustomer) GetTenantId()(*string) {
    return m.tenantId
}
// Serialize serializes information the current object
func (m *DelegatedAdminCustomer) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    if m.GetServiceManagementDetails() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetServiceManagementDetails())
        err = writer.WriteCollectionOfObjectValues("serviceManagementDetails", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("tenantId", m.GetTenantId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDisplayName sets the displayName property value. The Azure AD display name of the customer tenant. Read-only. Supports $orderBy.
func (m *DelegatedAdminCustomer) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetServiceManagementDetails sets the serviceManagementDetails property value. Contains the management details of a service in the customer tenant that's managed by delegated administration.
func (m *DelegatedAdminCustomer) SetServiceManagementDetails(value []DelegatedAdminServiceManagementDetailable)() {
    m.serviceManagementDetails = value
}
// SetTenantId sets the tenantId property value. The Azure AD-assigned tenant ID of the customer. Read-only.
func (m *DelegatedAdminCustomer) SetTenantId(value *string)() {
    m.tenantId = value
}

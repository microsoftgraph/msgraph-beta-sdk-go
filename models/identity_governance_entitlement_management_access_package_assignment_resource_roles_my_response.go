package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRolesMyResponse provides operations to call the My method.
type IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRolesMyResponse struct {
    BaseCollectionPaginationCountResponse
    // The value property
    value []AccessPackageAssignmentResourceRoleable
}
// NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRolesMyResponse instantiates a new IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRolesMyResponse and sets the default values.
func NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRolesMyResponse()(*IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRolesMyResponse) {
    m := &IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRolesMyResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateIdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRolesMyResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRolesMyResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRolesMyResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRolesMyResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAccessPackageAssignmentResourceRoleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessPackageAssignmentResourceRoleable, len(val))
            for i, v := range val {
                res[i] = v.(AccessPackageAssignmentResourceRoleable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRolesMyResponse) GetValue()([]AccessPackageAssignmentResourceRoleable) {
    return m.value
}
// Serialize serializes information the current object
func (m *IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRolesMyResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseCollectionPaginationCountResponse.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetValue() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetValue()))
        for i, v := range m.GetValue() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("value", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The value property
func (m *IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRolesMyResponse) SetValue(value []AccessPackageAssignmentResourceRoleable)() {
    m.value = value
}

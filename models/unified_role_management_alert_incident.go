package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UnifiedRoleManagementAlertIncident 
type UnifiedRoleManagementAlertIncident struct {
    Entity
}
// NewUnifiedRoleManagementAlertIncident instantiates a new unifiedRoleManagementAlertIncident and sets the default values.
func NewUnifiedRoleManagementAlertIncident()(*UnifiedRoleManagementAlertIncident) {
    m := &UnifiedRoleManagementAlertIncident{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUnifiedRoleManagementAlertIncidentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUnifiedRoleManagementAlertIncidentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.invalidLicenseAlertIncident":
                        return NewInvalidLicenseAlertIncident(), nil
                    case "#microsoft.graph.noMfaOnRoleActivationAlertIncident":
                        return NewNoMfaOnRoleActivationAlertIncident(), nil
                    case "#microsoft.graph.redundantAssignmentAlertIncident":
                        return NewRedundantAssignmentAlertIncident(), nil
                    case "#microsoft.graph.rolesAssignedOutsidePrivilegedIdentityManagementAlertIncident":
                        return NewRolesAssignedOutsidePrivilegedIdentityManagementAlertIncident(), nil
                    case "#microsoft.graph.sequentialActivationRenewalsAlertIncident":
                        return NewSequentialActivationRenewalsAlertIncident(), nil
                    case "#microsoft.graph.staleSignInAlertIncident":
                        return NewStaleSignInAlertIncident(), nil
                    case "#microsoft.graph.tooManyGlobalAdminsAssignedToTenantAlertIncident":
                        return NewTooManyGlobalAdminsAssignedToTenantAlertIncident(), nil
                }
            }
        }
    }
    return NewUnifiedRoleManagementAlertIncident(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UnifiedRoleManagementAlertIncident) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *UnifiedRoleManagementAlertIncident) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// UnifiedRoleManagementAlertIncidentable 
type UnifiedRoleManagementAlertIncidentable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}

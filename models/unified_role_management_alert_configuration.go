package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UnifiedRoleManagementAlertConfiguration 
type UnifiedRoleManagementAlertConfiguration struct {
    Entity
}
// NewUnifiedRoleManagementAlertConfiguration instantiates a new unifiedRoleManagementAlertConfiguration and sets the default values.
func NewUnifiedRoleManagementAlertConfiguration()(*UnifiedRoleManagementAlertConfiguration) {
    m := &UnifiedRoleManagementAlertConfiguration{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUnifiedRoleManagementAlertConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUnifiedRoleManagementAlertConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.invalidLicenseAlertConfiguration":
                        return NewInvalidLicenseAlertConfiguration(), nil
                    case "#microsoft.graph.noMfaOnRoleActivationAlertConfiguration":
                        return NewNoMfaOnRoleActivationAlertConfiguration(), nil
                    case "#microsoft.graph.redundantAssignmentAlertConfiguration":
                        return NewRedundantAssignmentAlertConfiguration(), nil
                    case "#microsoft.graph.rolesAssignedOutsidePrivilegedIdentityManagementAlertConfiguration":
                        return NewRolesAssignedOutsidePrivilegedIdentityManagementAlertConfiguration(), nil
                    case "#microsoft.graph.sequentialActivationRenewalsAlertConfiguration":
                        return NewSequentialActivationRenewalsAlertConfiguration(), nil
                    case "#microsoft.graph.staleSignInAlertConfiguration":
                        return NewStaleSignInAlertConfiguration(), nil
                    case "#microsoft.graph.tooManyGlobalAdminsAssignedToTenantAlertConfiguration":
                        return NewTooManyGlobalAdminsAssignedToTenantAlertConfiguration(), nil
                }
            }
        }
    }
    return NewUnifiedRoleManagementAlertConfiguration(), nil
}
// GetAlertDefinition gets the alertDefinition property value. The definition of the alert that contains its description, impact, and measures to mitigate or prevent it. Supports $expand.
func (m *UnifiedRoleManagementAlertConfiguration) GetAlertDefinition()(UnifiedRoleManagementAlertDefinitionable) {
    val, err := m.GetBackingStore().Get("alertDefinition")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(UnifiedRoleManagementAlertDefinitionable)
    }
    return nil
}
// GetAlertDefinitionId gets the alertDefinitionId property value. The identifier of an alert definition. Supports $filter (eq, ne).
func (m *UnifiedRoleManagementAlertConfiguration) GetAlertDefinitionId()(*string) {
    val, err := m.GetBackingStore().Get("alertDefinitionId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UnifiedRoleManagementAlertConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["alertDefinition"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUnifiedRoleManagementAlertDefinitionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAlertDefinition(val.(UnifiedRoleManagementAlertDefinitionable))
        }
        return nil
    }
    res["alertDefinitionId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAlertDefinitionId(val)
        }
        return nil
    }
    res["isEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEnabled(val)
        }
        return nil
    }
    res["scopeId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScopeId(val)
        }
        return nil
    }
    res["scopeType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScopeType(val)
        }
        return nil
    }
    return res
}
// GetIsEnabled gets the isEnabled property value. true if the alert is enabled. Setting it to false disables PIM scanning the tenant to identify instances that trigger the alert.
func (m *UnifiedRoleManagementAlertConfiguration) GetIsEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("isEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetScopeId gets the scopeId property value. The identifier of the scope to which the alert is related. Only / is supported to represent the tenant scope. Supports $filter (eq, ne).
func (m *UnifiedRoleManagementAlertConfiguration) GetScopeId()(*string) {
    val, err := m.GetBackingStore().Get("scopeId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetScopeType gets the scopeType property value. The type of scope where the alert is created. DirectoryRole is the only currently supported scope type for Azure AD roles.
func (m *UnifiedRoleManagementAlertConfiguration) GetScopeType()(*string) {
    val, err := m.GetBackingStore().Get("scopeType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *UnifiedRoleManagementAlertConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("alertDefinition", m.GetAlertDefinition())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("alertDefinitionId", m.GetAlertDefinitionId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isEnabled", m.GetIsEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("scopeId", m.GetScopeId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("scopeType", m.GetScopeType())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAlertDefinition sets the alertDefinition property value. The definition of the alert that contains its description, impact, and measures to mitigate or prevent it. Supports $expand.
func (m *UnifiedRoleManagementAlertConfiguration) SetAlertDefinition(value UnifiedRoleManagementAlertDefinitionable)() {
    err := m.GetBackingStore().Set("alertDefinition", value)
    if err != nil {
        panic(err)
    }
}
// SetAlertDefinitionId sets the alertDefinitionId property value. The identifier of an alert definition. Supports $filter (eq, ne).
func (m *UnifiedRoleManagementAlertConfiguration) SetAlertDefinitionId(value *string)() {
    err := m.GetBackingStore().Set("alertDefinitionId", value)
    if err != nil {
        panic(err)
    }
}
// SetIsEnabled sets the isEnabled property value. true if the alert is enabled. Setting it to false disables PIM scanning the tenant to identify instances that trigger the alert.
func (m *UnifiedRoleManagementAlertConfiguration) SetIsEnabled(value *bool)() {
    err := m.GetBackingStore().Set("isEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetScopeId sets the scopeId property value. The identifier of the scope to which the alert is related. Only / is supported to represent the tenant scope. Supports $filter (eq, ne).
func (m *UnifiedRoleManagementAlertConfiguration) SetScopeId(value *string)() {
    err := m.GetBackingStore().Set("scopeId", value)
    if err != nil {
        panic(err)
    }
}
// SetScopeType sets the scopeType property value. The type of scope where the alert is created. DirectoryRole is the only currently supported scope type for Azure AD roles.
func (m *UnifiedRoleManagementAlertConfiguration) SetScopeType(value *string)() {
    err := m.GetBackingStore().Set("scopeType", value)
    if err != nil {
        panic(err)
    }
}
// UnifiedRoleManagementAlertConfigurationable 
type UnifiedRoleManagementAlertConfigurationable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAlertDefinition()(UnifiedRoleManagementAlertDefinitionable)
    GetAlertDefinitionId()(*string)
    GetIsEnabled()(*bool)
    GetScopeId()(*string)
    GetScopeType()(*string)
    SetAlertDefinition(value UnifiedRoleManagementAlertDefinitionable)()
    SetAlertDefinitionId(value *string)()
    SetIsEnabled(value *bool)()
    SetScopeId(value *string)()
    SetScopeType(value *string)()
}

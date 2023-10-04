package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UnifiedRoleManagementAlertDefinition 
type UnifiedRoleManagementAlertDefinition struct {
    Entity
}
// NewUnifiedRoleManagementAlertDefinition instantiates a new unifiedRoleManagementAlertDefinition and sets the default values.
func NewUnifiedRoleManagementAlertDefinition()(*UnifiedRoleManagementAlertDefinition) {
    m := &UnifiedRoleManagementAlertDefinition{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUnifiedRoleManagementAlertDefinitionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUnifiedRoleManagementAlertDefinitionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUnifiedRoleManagementAlertDefinition(), nil
}
// GetDescription gets the description property value. The description of the alert.
func (m *UnifiedRoleManagementAlertDefinition) GetDescription()(*string) {
    val, err := m.GetBackingStore().Get("description")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDisplayName gets the displayName property value. The friendly display name that renders in Privileged Identity Management (PIM) alerts in the Microsoft Entra admin center.
func (m *UnifiedRoleManagementAlertDefinition) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UnifiedRoleManagementAlertDefinition) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["howToPrevent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHowToPrevent(val)
        }
        return nil
    }
    res["isConfigurable"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsConfigurable(val)
        }
        return nil
    }
    res["isRemediatable"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsRemediatable(val)
        }
        return nil
    }
    res["mitigationSteps"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMitigationSteps(val)
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
    res["securityImpact"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecurityImpact(val)
        }
        return nil
    }
    res["severityLevel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAlertSeverity)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSeverityLevel(val.(*AlertSeverity))
        }
        return nil
    }
    return res
}
// GetHowToPrevent gets the howToPrevent property value. Long-form text that indicates the ways to prevent the alert from being triggered in your tenant.
func (m *UnifiedRoleManagementAlertDefinition) GetHowToPrevent()(*string) {
    val, err := m.GetBackingStore().Get("howToPrevent")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetIsConfigurable gets the isConfigurable property value. true if the alert configuration can be customized in the tenant, and false otherwise. For example, the number and percentage thresholds of the 'There are too many global administrators' alert can be configured by users, while the 'This organization doesn't have Azure AD Premium P2' can't be configured, because the criteria are restricted.
func (m *UnifiedRoleManagementAlertDefinition) GetIsConfigurable()(*bool) {
    val, err := m.GetBackingStore().Get("isConfigurable")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIsRemediatable gets the isRemediatable property value. true if the alert can be remediated, and false otherwise.
func (m *UnifiedRoleManagementAlertDefinition) GetIsRemediatable()(*bool) {
    val, err := m.GetBackingStore().Get("isRemediatable")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetMitigationSteps gets the mitigationSteps property value. The methods to mitigate the alert when it's triggered in the tenant. For example, to mitigate the 'There are too many global administrators', you could remove redundant privileged role assignments.
func (m *UnifiedRoleManagementAlertDefinition) GetMitigationSteps()(*string) {
    val, err := m.GetBackingStore().Get("mitigationSteps")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetScopeId gets the scopeId property value. The identifier of the scope where the alert is related. / is the only supported one for the tenant. Supports $filter (eq, ne).
func (m *UnifiedRoleManagementAlertDefinition) GetScopeId()(*string) {
    val, err := m.GetBackingStore().Get("scopeId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetScopeType gets the scopeType property value. The type of scope where the alert is created. DirectoryRole is the only currently supported scope type for Azure AD Roles.
func (m *UnifiedRoleManagementAlertDefinition) GetScopeType()(*string) {
    val, err := m.GetBackingStore().Get("scopeType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSecurityImpact gets the securityImpact property value. Security impact of the alert. For example, it could be information leaks or unauthorized access.
func (m *UnifiedRoleManagementAlertDefinition) GetSecurityImpact()(*string) {
    val, err := m.GetBackingStore().Get("securityImpact")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSeverityLevel gets the severityLevel property value. Severity level of the alert. The possible values are: unknown, informational, low, medium, high, unknownFutureValue.
func (m *UnifiedRoleManagementAlertDefinition) GetSeverityLevel()(*AlertSeverity) {
    val, err := m.GetBackingStore().Get("severityLevel")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AlertSeverity)
    }
    return nil
}
// Serialize serializes information the current object
func (m *UnifiedRoleManagementAlertDefinition) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("howToPrevent", m.GetHowToPrevent())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isConfigurable", m.GetIsConfigurable())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isRemediatable", m.GetIsRemediatable())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("mitigationSteps", m.GetMitigationSteps())
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
    {
        err = writer.WriteStringValue("securityImpact", m.GetSecurityImpact())
        if err != nil {
            return err
        }
    }
    if m.GetSeverityLevel() != nil {
        cast := (*m.GetSeverityLevel()).String()
        err = writer.WriteStringValue("severityLevel", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDescription sets the description property value. The description of the alert.
func (m *UnifiedRoleManagementAlertDefinition) SetDescription(value *string)() {
    err := m.GetBackingStore().Set("description", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. The friendly display name that renders in Privileged Identity Management (PIM) alerts in the Microsoft Entra admin center.
func (m *UnifiedRoleManagementAlertDefinition) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetHowToPrevent sets the howToPrevent property value. Long-form text that indicates the ways to prevent the alert from being triggered in your tenant.
func (m *UnifiedRoleManagementAlertDefinition) SetHowToPrevent(value *string)() {
    err := m.GetBackingStore().Set("howToPrevent", value)
    if err != nil {
        panic(err)
    }
}
// SetIsConfigurable sets the isConfigurable property value. true if the alert configuration can be customized in the tenant, and false otherwise. For example, the number and percentage thresholds of the 'There are too many global administrators' alert can be configured by users, while the 'This organization doesn't have Azure AD Premium P2' can't be configured, because the criteria are restricted.
func (m *UnifiedRoleManagementAlertDefinition) SetIsConfigurable(value *bool)() {
    err := m.GetBackingStore().Set("isConfigurable", value)
    if err != nil {
        panic(err)
    }
}
// SetIsRemediatable sets the isRemediatable property value. true if the alert can be remediated, and false otherwise.
func (m *UnifiedRoleManagementAlertDefinition) SetIsRemediatable(value *bool)() {
    err := m.GetBackingStore().Set("isRemediatable", value)
    if err != nil {
        panic(err)
    }
}
// SetMitigationSteps sets the mitigationSteps property value. The methods to mitigate the alert when it's triggered in the tenant. For example, to mitigate the 'There are too many global administrators', you could remove redundant privileged role assignments.
func (m *UnifiedRoleManagementAlertDefinition) SetMitigationSteps(value *string)() {
    err := m.GetBackingStore().Set("mitigationSteps", value)
    if err != nil {
        panic(err)
    }
}
// SetScopeId sets the scopeId property value. The identifier of the scope where the alert is related. / is the only supported one for the tenant. Supports $filter (eq, ne).
func (m *UnifiedRoleManagementAlertDefinition) SetScopeId(value *string)() {
    err := m.GetBackingStore().Set("scopeId", value)
    if err != nil {
        panic(err)
    }
}
// SetScopeType sets the scopeType property value. The type of scope where the alert is created. DirectoryRole is the only currently supported scope type for Azure AD Roles.
func (m *UnifiedRoleManagementAlertDefinition) SetScopeType(value *string)() {
    err := m.GetBackingStore().Set("scopeType", value)
    if err != nil {
        panic(err)
    }
}
// SetSecurityImpact sets the securityImpact property value. Security impact of the alert. For example, it could be information leaks or unauthorized access.
func (m *UnifiedRoleManagementAlertDefinition) SetSecurityImpact(value *string)() {
    err := m.GetBackingStore().Set("securityImpact", value)
    if err != nil {
        panic(err)
    }
}
// SetSeverityLevel sets the severityLevel property value. Severity level of the alert. The possible values are: unknown, informational, low, medium, high, unknownFutureValue.
func (m *UnifiedRoleManagementAlertDefinition) SetSeverityLevel(value *AlertSeverity)() {
    err := m.GetBackingStore().Set("severityLevel", value)
    if err != nil {
        panic(err)
    }
}
// UnifiedRoleManagementAlertDefinitionable 
type UnifiedRoleManagementAlertDefinitionable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetHowToPrevent()(*string)
    GetIsConfigurable()(*bool)
    GetIsRemediatable()(*bool)
    GetMitigationSteps()(*string)
    GetScopeId()(*string)
    GetScopeType()(*string)
    GetSecurityImpact()(*string)
    GetSeverityLevel()(*AlertSeverity)
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetHowToPrevent(value *string)()
    SetIsConfigurable(value *bool)()
    SetIsRemediatable(value *bool)()
    SetMitigationSteps(value *string)()
    SetScopeId(value *string)()
    SetScopeType(value *string)()
    SetSecurityImpact(value *string)()
    SetSeverityLevel(value *AlertSeverity)()
}

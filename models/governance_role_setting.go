package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type GovernanceRoleSetting struct {
    Entity
}
// NewGovernanceRoleSetting instantiates a new GovernanceRoleSetting and sets the default values.
func NewGovernanceRoleSetting()(*GovernanceRoleSetting) {
    m := &GovernanceRoleSetting{
        Entity: *NewEntity(),
    }
    return m
}
// CreateGovernanceRoleSettingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGovernanceRoleSettingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGovernanceRoleSetting(), nil
}
// GetAdminEligibleSettings gets the adminEligibleSettings property value. The rule settings that are evaluated when an administrator tries to add an eligible role assignment.
// returns a []GovernanceRuleSettingable when successful
func (m *GovernanceRoleSetting) GetAdminEligibleSettings()([]GovernanceRuleSettingable) {
    val, err := m.GetBackingStore().Get("adminEligibleSettings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]GovernanceRuleSettingable)
    }
    return nil
}
// GetAdminMemberSettings gets the adminMemberSettings property value. The rule settings that are evaluated when an administrator tries to add a direct member role assignment.
// returns a []GovernanceRuleSettingable when successful
func (m *GovernanceRoleSetting) GetAdminMemberSettings()([]GovernanceRuleSettingable) {
    val, err := m.GetBackingStore().Get("adminMemberSettings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]GovernanceRuleSettingable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GovernanceRoleSetting) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["adminEligibleSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateGovernanceRuleSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GovernanceRuleSettingable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(GovernanceRuleSettingable)
                }
            }
            m.SetAdminEligibleSettings(res)
        }
        return nil
    }
    res["adminMemberSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateGovernanceRuleSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GovernanceRuleSettingable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(GovernanceRuleSettingable)
                }
            }
            m.SetAdminMemberSettings(res)
        }
        return nil
    }
    res["isDefault"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDefault(val)
        }
        return nil
    }
    res["lastUpdatedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastUpdatedBy(val)
        }
        return nil
    }
    res["lastUpdatedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastUpdatedDateTime(val)
        }
        return nil
    }
    res["resource"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGovernanceResourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResource(val.(GovernanceResourceable))
        }
        return nil
    }
    res["resourceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceId(val)
        }
        return nil
    }
    res["roleDefinition"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGovernanceRoleDefinitionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoleDefinition(val.(GovernanceRoleDefinitionable))
        }
        return nil
    }
    res["roleDefinitionId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoleDefinitionId(val)
        }
        return nil
    }
    res["userEligibleSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateGovernanceRuleSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GovernanceRuleSettingable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(GovernanceRuleSettingable)
                }
            }
            m.SetUserEligibleSettings(res)
        }
        return nil
    }
    res["userMemberSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateGovernanceRuleSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GovernanceRuleSettingable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(GovernanceRuleSettingable)
                }
            }
            m.SetUserMemberSettings(res)
        }
        return nil
    }
    return res
}
// GetIsDefault gets the isDefault property value. Read-only. Indicate if the roleSetting is a default roleSetting
// returns a *bool when successful
func (m *GovernanceRoleSetting) GetIsDefault()(*bool) {
    val, err := m.GetBackingStore().Get("isDefault")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetLastUpdatedBy gets the lastUpdatedBy property value. Read-only. The display name of the administrator who last updated the roleSetting.
// returns a *string when successful
func (m *GovernanceRoleSetting) GetLastUpdatedBy()(*string) {
    val, err := m.GetBackingStore().Get("lastUpdatedBy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetLastUpdatedDateTime gets the lastUpdatedDateTime property value. Read-only. The time when the role setting was last updated. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
// returns a *Time when successful
func (m *GovernanceRoleSetting) GetLastUpdatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastUpdatedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetResource gets the resource property value. Read-only. The associated resource for this role setting.
// returns a GovernanceResourceable when successful
func (m *GovernanceRoleSetting) GetResource()(GovernanceResourceable) {
    val, err := m.GetBackingStore().Get("resource")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(GovernanceResourceable)
    }
    return nil
}
// GetResourceId gets the resourceId property value. Required. The id of the resource that the role setting is associated with.
// returns a *string when successful
func (m *GovernanceRoleSetting) GetResourceId()(*string) {
    val, err := m.GetBackingStore().Get("resourceId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRoleDefinition gets the roleDefinition property value. Read-only. The role definition that is enforced with this role setting.
// returns a GovernanceRoleDefinitionable when successful
func (m *GovernanceRoleSetting) GetRoleDefinition()(GovernanceRoleDefinitionable) {
    val, err := m.GetBackingStore().Get("roleDefinition")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(GovernanceRoleDefinitionable)
    }
    return nil
}
// GetRoleDefinitionId gets the roleDefinitionId property value. Required. The id of the role definition that the role setting is associated with.
// returns a *string when successful
func (m *GovernanceRoleSetting) GetRoleDefinitionId()(*string) {
    val, err := m.GetBackingStore().Get("roleDefinitionId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetUserEligibleSettings gets the userEligibleSettings property value. The rule settings that are evaluated when a user tries to add an eligible role assignment. The setting is not supported for now.
// returns a []GovernanceRuleSettingable when successful
func (m *GovernanceRoleSetting) GetUserEligibleSettings()([]GovernanceRuleSettingable) {
    val, err := m.GetBackingStore().Get("userEligibleSettings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]GovernanceRuleSettingable)
    }
    return nil
}
// GetUserMemberSettings gets the userMemberSettings property value. The rule settings that are evaluated when a user tries to activate his role assignment.
// returns a []GovernanceRuleSettingable when successful
func (m *GovernanceRoleSetting) GetUserMemberSettings()([]GovernanceRuleSettingable) {
    val, err := m.GetBackingStore().Get("userMemberSettings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]GovernanceRuleSettingable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *GovernanceRoleSetting) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAdminEligibleSettings() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAdminEligibleSettings()))
        for i, v := range m.GetAdminEligibleSettings() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("adminEligibleSettings", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAdminMemberSettings() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAdminMemberSettings()))
        for i, v := range m.GetAdminMemberSettings() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("adminMemberSettings", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isDefault", m.GetIsDefault())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("lastUpdatedBy", m.GetLastUpdatedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastUpdatedDateTime", m.GetLastUpdatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("resource", m.GetResource())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("resourceId", m.GetResourceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("roleDefinition", m.GetRoleDefinition())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("roleDefinitionId", m.GetRoleDefinitionId())
        if err != nil {
            return err
        }
    }
    if m.GetUserEligibleSettings() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserEligibleSettings()))
        for i, v := range m.GetUserEligibleSettings() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("userEligibleSettings", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserMemberSettings() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserMemberSettings()))
        for i, v := range m.GetUserMemberSettings() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("userMemberSettings", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdminEligibleSettings sets the adminEligibleSettings property value. The rule settings that are evaluated when an administrator tries to add an eligible role assignment.
func (m *GovernanceRoleSetting) SetAdminEligibleSettings(value []GovernanceRuleSettingable)() {
    err := m.GetBackingStore().Set("adminEligibleSettings", value)
    if err != nil {
        panic(err)
    }
}
// SetAdminMemberSettings sets the adminMemberSettings property value. The rule settings that are evaluated when an administrator tries to add a direct member role assignment.
func (m *GovernanceRoleSetting) SetAdminMemberSettings(value []GovernanceRuleSettingable)() {
    err := m.GetBackingStore().Set("adminMemberSettings", value)
    if err != nil {
        panic(err)
    }
}
// SetIsDefault sets the isDefault property value. Read-only. Indicate if the roleSetting is a default roleSetting
func (m *GovernanceRoleSetting) SetIsDefault(value *bool)() {
    err := m.GetBackingStore().Set("isDefault", value)
    if err != nil {
        panic(err)
    }
}
// SetLastUpdatedBy sets the lastUpdatedBy property value. Read-only. The display name of the administrator who last updated the roleSetting.
func (m *GovernanceRoleSetting) SetLastUpdatedBy(value *string)() {
    err := m.GetBackingStore().Set("lastUpdatedBy", value)
    if err != nil {
        panic(err)
    }
}
// SetLastUpdatedDateTime sets the lastUpdatedDateTime property value. Read-only. The time when the role setting was last updated. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *GovernanceRoleSetting) SetLastUpdatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastUpdatedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetResource sets the resource property value. Read-only. The associated resource for this role setting.
func (m *GovernanceRoleSetting) SetResource(value GovernanceResourceable)() {
    err := m.GetBackingStore().Set("resource", value)
    if err != nil {
        panic(err)
    }
}
// SetResourceId sets the resourceId property value. Required. The id of the resource that the role setting is associated with.
func (m *GovernanceRoleSetting) SetResourceId(value *string)() {
    err := m.GetBackingStore().Set("resourceId", value)
    if err != nil {
        panic(err)
    }
}
// SetRoleDefinition sets the roleDefinition property value. Read-only. The role definition that is enforced with this role setting.
func (m *GovernanceRoleSetting) SetRoleDefinition(value GovernanceRoleDefinitionable)() {
    err := m.GetBackingStore().Set("roleDefinition", value)
    if err != nil {
        panic(err)
    }
}
// SetRoleDefinitionId sets the roleDefinitionId property value. Required. The id of the role definition that the role setting is associated with.
func (m *GovernanceRoleSetting) SetRoleDefinitionId(value *string)() {
    err := m.GetBackingStore().Set("roleDefinitionId", value)
    if err != nil {
        panic(err)
    }
}
// SetUserEligibleSettings sets the userEligibleSettings property value. The rule settings that are evaluated when a user tries to add an eligible role assignment. The setting is not supported for now.
func (m *GovernanceRoleSetting) SetUserEligibleSettings(value []GovernanceRuleSettingable)() {
    err := m.GetBackingStore().Set("userEligibleSettings", value)
    if err != nil {
        panic(err)
    }
}
// SetUserMemberSettings sets the userMemberSettings property value. The rule settings that are evaluated when a user tries to activate his role assignment.
func (m *GovernanceRoleSetting) SetUserMemberSettings(value []GovernanceRuleSettingable)() {
    err := m.GetBackingStore().Set("userMemberSettings", value)
    if err != nil {
        panic(err)
    }
}
type GovernanceRoleSettingable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdminEligibleSettings()([]GovernanceRuleSettingable)
    GetAdminMemberSettings()([]GovernanceRuleSettingable)
    GetIsDefault()(*bool)
    GetLastUpdatedBy()(*string)
    GetLastUpdatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetResource()(GovernanceResourceable)
    GetResourceId()(*string)
    GetRoleDefinition()(GovernanceRoleDefinitionable)
    GetRoleDefinitionId()(*string)
    GetUserEligibleSettings()([]GovernanceRuleSettingable)
    GetUserMemberSettings()([]GovernanceRuleSettingable)
    SetAdminEligibleSettings(value []GovernanceRuleSettingable)()
    SetAdminMemberSettings(value []GovernanceRuleSettingable)()
    SetIsDefault(value *bool)()
    SetLastUpdatedBy(value *string)()
    SetLastUpdatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetResource(value GovernanceResourceable)()
    SetResourceId(value *string)()
    SetRoleDefinition(value GovernanceRoleDefinitionable)()
    SetRoleDefinitionId(value *string)()
    SetUserEligibleSettings(value []GovernanceRuleSettingable)()
    SetUserMemberSettings(value []GovernanceRuleSettingable)()
}

package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type EntitlementManagementSettings struct {
    Entity
    // If externalUserLifecycleAction is BlockSignInAndDelete, the number of days after an external user is blocked from sign in before their account is deleted.
    daysUntilExternalUserDeletedAfterBlocked *int32;
    // One of None, BlockSignIn, or BlockSignInAndDelete.
    externalUserLifecycleAction *string;
}
// Instantiates a new entitlementManagementSettings and sets the default values.
func NewEntitlementManagementSettings()(*EntitlementManagementSettings) {
    m := &EntitlementManagementSettings{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the daysUntilExternalUserDeletedAfterBlocked property value. If externalUserLifecycleAction is BlockSignInAndDelete, the number of days after an external user is blocked from sign in before their account is deleted.
func (m *EntitlementManagementSettings) GetDaysUntilExternalUserDeletedAfterBlocked()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.daysUntilExternalUserDeletedAfterBlocked
    }
}
// Gets the externalUserLifecycleAction property value. One of None, BlockSignIn, or BlockSignInAndDelete.
func (m *EntitlementManagementSettings) GetExternalUserLifecycleAction()(*string) {
    if m == nil {
        return nil
    } else {
        return m.externalUserLifecycleAction
    }
}
// The deserialization information for the current model
func (m *EntitlementManagementSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["daysUntilExternalUserDeletedAfterBlocked"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDaysUntilExternalUserDeletedAfterBlocked(val)
        }
        return nil
    }
    res["externalUserLifecycleAction"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalUserLifecycleAction(val)
        }
        return nil
    }
    return res
}
func (m *EntitlementManagementSettings) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *EntitlementManagementSettings) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("daysUntilExternalUserDeletedAfterBlocked", m.GetDaysUntilExternalUserDeletedAfterBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("externalUserLifecycleAction", m.GetExternalUserLifecycleAction())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the daysUntilExternalUserDeletedAfterBlocked property value. If externalUserLifecycleAction is BlockSignInAndDelete, the number of days after an external user is blocked from sign in before their account is deleted.
// Parameters:
//  - value : Value to set for the daysUntilExternalUserDeletedAfterBlocked property.
func (m *EntitlementManagementSettings) SetDaysUntilExternalUserDeletedAfterBlocked(value *int32)() {
    m.daysUntilExternalUserDeletedAfterBlocked = value
}
// Sets the externalUserLifecycleAction property value. One of None, BlockSignIn, or BlockSignInAndDelete.
// Parameters:
//  - value : Value to set for the externalUserLifecycleAction property.
func (m *EntitlementManagementSettings) SetExternalUserLifecycleAction(value *string)() {
    m.externalUserLifecycleAction = value
}

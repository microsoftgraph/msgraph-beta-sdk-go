package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type PrivilegedRoleSettings struct {
    Entity
    // true if the approval is required when activate the role. false if the approval is not required when activate the role.
    approvalOnElevation *bool;
    // List of Approval ids, if approval is required for activation.
    approverIds []string;
    // The duration when the role is activated.
    elevationDuration *string;
    // true if mfaOnElevation is configurable. false if mfaOnElevation is not configurable.
    isMfaOnElevationConfigurable *bool;
    // Internal used only.
    lastGlobalAdmin *bool;
    // Maximal duration for the activated role.
    maxElavationDuration *string;
    // true if MFA is required to activate the role. false if MFA is not required to activate the role.
    mfaOnElevation *bool;
    // Minimal duration for the activated role.
    minElevationDuration *string;
    // true if send notification to the end user when the role is activated. false if do not send notification when the role is activated.
    notificationToUserOnElevation *bool;
    // true if the ticketing information is required when activate the role. false if the ticketing information is not required when activate the role.
    ticketingInfoOnElevation *bool;
}
// Instantiates a new privilegedRoleSettings and sets the default values.
func NewPrivilegedRoleSettings()(*PrivilegedRoleSettings) {
    m := &PrivilegedRoleSettings{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the approvalOnElevation property value. true if the approval is required when activate the role. false if the approval is not required when activate the role.
func (m *PrivilegedRoleSettings) GetApprovalOnElevation()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.approvalOnElevation
    }
}
// Gets the approverIds property value. List of Approval ids, if approval is required for activation.
func (m *PrivilegedRoleSettings) GetApproverIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.approverIds
    }
}
// Gets the elevationDuration property value. The duration when the role is activated.
func (m *PrivilegedRoleSettings) GetElevationDuration()(*string) {
    if m == nil {
        return nil
    } else {
        return m.elevationDuration
    }
}
// Gets the isMfaOnElevationConfigurable property value. true if mfaOnElevation is configurable. false if mfaOnElevation is not configurable.
func (m *PrivilegedRoleSettings) GetIsMfaOnElevationConfigurable()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isMfaOnElevationConfigurable
    }
}
// Gets the lastGlobalAdmin property value. Internal used only.
func (m *PrivilegedRoleSettings) GetLastGlobalAdmin()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.lastGlobalAdmin
    }
}
// Gets the maxElavationDuration property value. Maximal duration for the activated role.
func (m *PrivilegedRoleSettings) GetMaxElavationDuration()(*string) {
    if m == nil {
        return nil
    } else {
        return m.maxElavationDuration
    }
}
// Gets the mfaOnElevation property value. true if MFA is required to activate the role. false if MFA is not required to activate the role.
func (m *PrivilegedRoleSettings) GetMfaOnElevation()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.mfaOnElevation
    }
}
// Gets the minElevationDuration property value. Minimal duration for the activated role.
func (m *PrivilegedRoleSettings) GetMinElevationDuration()(*string) {
    if m == nil {
        return nil
    } else {
        return m.minElevationDuration
    }
}
// Gets the notificationToUserOnElevation property value. true if send notification to the end user when the role is activated. false if do not send notification when the role is activated.
func (m *PrivilegedRoleSettings) GetNotificationToUserOnElevation()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.notificationToUserOnElevation
    }
}
// Gets the ticketingInfoOnElevation property value. true if the ticketing information is required when activate the role. false if the ticketing information is not required when activate the role.
func (m *PrivilegedRoleSettings) GetTicketingInfoOnElevation()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.ticketingInfoOnElevation
    }
}
// The deserialization information for the current model
func (m *PrivilegedRoleSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["approvalOnElevation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetApprovalOnElevation(val)
        return nil
    }
    res["approverIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetApproverIds(res)
        return nil
    }
    res["elevationDuration"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetElevationDuration(val)
        return nil
    }
    res["isMfaOnElevationConfigurable"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsMfaOnElevationConfigurable(val)
        return nil
    }
    res["lastGlobalAdmin"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetLastGlobalAdmin(val)
        return nil
    }
    res["maxElavationDuration"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetMaxElavationDuration(val)
        return nil
    }
    res["mfaOnElevation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetMfaOnElevation(val)
        return nil
    }
    res["minElevationDuration"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetMinElevationDuration(val)
        return nil
    }
    res["notificationToUserOnElevation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetNotificationToUserOnElevation(val)
        return nil
    }
    res["ticketingInfoOnElevation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetTicketingInfoOnElevation(val)
        return nil
    }
    return res
}
func (m *PrivilegedRoleSettings) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *PrivilegedRoleSettings) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("approvalOnElevation", m.GetApprovalOnElevation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("approverIds", m.GetApproverIds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("elevationDuration", m.GetElevationDuration())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isMfaOnElevationConfigurable", m.GetIsMfaOnElevationConfigurable())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("lastGlobalAdmin", m.GetLastGlobalAdmin())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("maxElavationDuration", m.GetMaxElavationDuration())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("mfaOnElevation", m.GetMfaOnElevation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("minElevationDuration", m.GetMinElevationDuration())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("notificationToUserOnElevation", m.GetNotificationToUserOnElevation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("ticketingInfoOnElevation", m.GetTicketingInfoOnElevation())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the approvalOnElevation property value. true if the approval is required when activate the role. false if the approval is not required when activate the role.
// Parameters:
//  - value : Value to set for the approvalOnElevation property.
func (m *PrivilegedRoleSettings) SetApprovalOnElevation(value *bool)() {
    m.approvalOnElevation = value
}
// Sets the approverIds property value. List of Approval ids, if approval is required for activation.
// Parameters:
//  - value : Value to set for the approverIds property.
func (m *PrivilegedRoleSettings) SetApproverIds(value []string)() {
    m.approverIds = value
}
// Sets the elevationDuration property value. The duration when the role is activated.
// Parameters:
//  - value : Value to set for the elevationDuration property.
func (m *PrivilegedRoleSettings) SetElevationDuration(value *string)() {
    m.elevationDuration = value
}
// Sets the isMfaOnElevationConfigurable property value. true if mfaOnElevation is configurable. false if mfaOnElevation is not configurable.
// Parameters:
//  - value : Value to set for the isMfaOnElevationConfigurable property.
func (m *PrivilegedRoleSettings) SetIsMfaOnElevationConfigurable(value *bool)() {
    m.isMfaOnElevationConfigurable = value
}
// Sets the lastGlobalAdmin property value. Internal used only.
// Parameters:
//  - value : Value to set for the lastGlobalAdmin property.
func (m *PrivilegedRoleSettings) SetLastGlobalAdmin(value *bool)() {
    m.lastGlobalAdmin = value
}
// Sets the maxElavationDuration property value. Maximal duration for the activated role.
// Parameters:
//  - value : Value to set for the maxElavationDuration property.
func (m *PrivilegedRoleSettings) SetMaxElavationDuration(value *string)() {
    m.maxElavationDuration = value
}
// Sets the mfaOnElevation property value. true if MFA is required to activate the role. false if MFA is not required to activate the role.
// Parameters:
//  - value : Value to set for the mfaOnElevation property.
func (m *PrivilegedRoleSettings) SetMfaOnElevation(value *bool)() {
    m.mfaOnElevation = value
}
// Sets the minElevationDuration property value. Minimal duration for the activated role.
// Parameters:
//  - value : Value to set for the minElevationDuration property.
func (m *PrivilegedRoleSettings) SetMinElevationDuration(value *string)() {
    m.minElevationDuration = value
}
// Sets the notificationToUserOnElevation property value. true if send notification to the end user when the role is activated. false if do not send notification when the role is activated.
// Parameters:
//  - value : Value to set for the notificationToUserOnElevation property.
func (m *PrivilegedRoleSettings) SetNotificationToUserOnElevation(value *bool)() {
    m.notificationToUserOnElevation = value
}
// Sets the ticketingInfoOnElevation property value. true if the ticketing information is required when activate the role. false if the ticketing information is not required when activate the role.
// Parameters:
//  - value : Value to set for the ticketingInfoOnElevation property.
func (m *PrivilegedRoleSettings) SetTicketingInfoOnElevation(value *bool)() {
    m.ticketingInfoOnElevation = value
}

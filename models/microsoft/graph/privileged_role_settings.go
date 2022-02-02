package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PrivilegedRoleSettings 
type PrivilegedRoleSettings struct {
    Entity
    // true if the approval is required when activate the role. false if the approval is not required when activate the role.
    approvalOnElevation *bool;
    // List of Approval ids, if approval is required for activation.
    approverIds []string;
    // The duration when the role is activated.
    elevationDuration *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration;
    // true if mfaOnElevation is configurable. false if mfaOnElevation is not configurable.
    isMfaOnElevationConfigurable *bool;
    // Internal used only.
    lastGlobalAdmin *bool;
    // Maximal duration for the activated role.
    maxElavationDuration *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration;
    // true if MFA is required to activate the role. false if MFA is not required to activate the role.
    mfaOnElevation *bool;
    // Minimal duration for the activated role.
    minElevationDuration *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration;
    // true if send notification to the end user when the role is activated. false if do not send notification when the role is activated.
    notificationToUserOnElevation *bool;
    // true if the ticketing information is required when activate the role. false if the ticketing information is not required when activate the role.
    ticketingInfoOnElevation *bool;
}
// NewPrivilegedRoleSettings instantiates a new privilegedRoleSettings and sets the default values.
func NewPrivilegedRoleSettings()(*PrivilegedRoleSettings) {
    m := &PrivilegedRoleSettings{
        Entity: *NewEntity(),
    }
    return m
}
// GetApprovalOnElevation gets the approvalOnElevation property value. true if the approval is required when activate the role. false if the approval is not required when activate the role.
func (m *PrivilegedRoleSettings) GetApprovalOnElevation()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.approvalOnElevation
    }
}
// GetApproverIds gets the approverIds property value. List of Approval ids, if approval is required for activation.
func (m *PrivilegedRoleSettings) GetApproverIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.approverIds
    }
}
// GetElevationDuration gets the elevationDuration property value. The duration when the role is activated.
func (m *PrivilegedRoleSettings) GetElevationDuration()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration) {
    if m == nil {
        return nil
    } else {
        return m.elevationDuration
    }
}
// GetIsMfaOnElevationConfigurable gets the isMfaOnElevationConfigurable property value. true if mfaOnElevation is configurable. false if mfaOnElevation is not configurable.
func (m *PrivilegedRoleSettings) GetIsMfaOnElevationConfigurable()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isMfaOnElevationConfigurable
    }
}
// GetLastGlobalAdmin gets the lastGlobalAdmin property value. Internal used only.
func (m *PrivilegedRoleSettings) GetLastGlobalAdmin()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.lastGlobalAdmin
    }
}
// GetMaxElavationDuration gets the maxElavationDuration property value. Maximal duration for the activated role.
func (m *PrivilegedRoleSettings) GetMaxElavationDuration()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration) {
    if m == nil {
        return nil
    } else {
        return m.maxElavationDuration
    }
}
// GetMfaOnElevation gets the mfaOnElevation property value. true if MFA is required to activate the role. false if MFA is not required to activate the role.
func (m *PrivilegedRoleSettings) GetMfaOnElevation()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.mfaOnElevation
    }
}
// GetMinElevationDuration gets the minElevationDuration property value. Minimal duration for the activated role.
func (m *PrivilegedRoleSettings) GetMinElevationDuration()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration) {
    if m == nil {
        return nil
    } else {
        return m.minElevationDuration
    }
}
// GetNotificationToUserOnElevation gets the notificationToUserOnElevation property value. true if send notification to the end user when the role is activated. false if do not send notification when the role is activated.
func (m *PrivilegedRoleSettings) GetNotificationToUserOnElevation()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.notificationToUserOnElevation
    }
}
// GetTicketingInfoOnElevation gets the ticketingInfoOnElevation property value. true if the ticketing information is required when activate the role. false if the ticketing information is not required when activate the role.
func (m *PrivilegedRoleSettings) GetTicketingInfoOnElevation()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.ticketingInfoOnElevation
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PrivilegedRoleSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["approvalOnElevation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApprovalOnElevation(val)
        }
        return nil
    }
    res["approverIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetApproverIds(res)
        }
        return nil
    }
    res["elevationDuration"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetElevationDuration(val)
        }
        return nil
    }
    res["isMfaOnElevationConfigurable"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsMfaOnElevationConfigurable(val)
        }
        return nil
    }
    res["lastGlobalAdmin"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastGlobalAdmin(val)
        }
        return nil
    }
    res["maxElavationDuration"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxElavationDuration(val)
        }
        return nil
    }
    res["mfaOnElevation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMfaOnElevation(val)
        }
        return nil
    }
    res["minElevationDuration"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinElevationDuration(val)
        }
        return nil
    }
    res["notificationToUserOnElevation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotificationToUserOnElevation(val)
        }
        return nil
    }
    res["ticketingInfoOnElevation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTicketingInfoOnElevation(val)
        }
        return nil
    }
    return res
}
func (m *PrivilegedRoleSettings) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
    if m.GetApproverIds() != nil {
        err = writer.WriteCollectionOfStringValues("approverIds", m.GetApproverIds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteISODurationValue("elevationDuration", m.GetElevationDuration())
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
        err = writer.WriteISODurationValue("maxElavationDuration", m.GetMaxElavationDuration())
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
        err = writer.WriteISODurationValue("minElevationDuration", m.GetMinElevationDuration())
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
// SetApprovalOnElevation sets the approvalOnElevation property value. true if the approval is required when activate the role. false if the approval is not required when activate the role.
func (m *PrivilegedRoleSettings) SetApprovalOnElevation(value *bool)() {
    if m != nil {
        m.approvalOnElevation = value
    }
}
// SetApproverIds sets the approverIds property value. List of Approval ids, if approval is required for activation.
func (m *PrivilegedRoleSettings) SetApproverIds(value []string)() {
    if m != nil {
        m.approverIds = value
    }
}
// SetElevationDuration sets the elevationDuration property value. The duration when the role is activated.
func (m *PrivilegedRoleSettings) SetElevationDuration(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)() {
    if m != nil {
        m.elevationDuration = value
    }
}
// SetIsMfaOnElevationConfigurable sets the isMfaOnElevationConfigurable property value. true if mfaOnElevation is configurable. false if mfaOnElevation is not configurable.
func (m *PrivilegedRoleSettings) SetIsMfaOnElevationConfigurable(value *bool)() {
    if m != nil {
        m.isMfaOnElevationConfigurable = value
    }
}
// SetLastGlobalAdmin sets the lastGlobalAdmin property value. Internal used only.
func (m *PrivilegedRoleSettings) SetLastGlobalAdmin(value *bool)() {
    if m != nil {
        m.lastGlobalAdmin = value
    }
}
// SetMaxElavationDuration sets the maxElavationDuration property value. Maximal duration for the activated role.
func (m *PrivilegedRoleSettings) SetMaxElavationDuration(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)() {
    if m != nil {
        m.maxElavationDuration = value
    }
}
// SetMfaOnElevation sets the mfaOnElevation property value. true if MFA is required to activate the role. false if MFA is not required to activate the role.
func (m *PrivilegedRoleSettings) SetMfaOnElevation(value *bool)() {
    if m != nil {
        m.mfaOnElevation = value
    }
}
// SetMinElevationDuration sets the minElevationDuration property value. Minimal duration for the activated role.
func (m *PrivilegedRoleSettings) SetMinElevationDuration(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)() {
    if m != nil {
        m.minElevationDuration = value
    }
}
// SetNotificationToUserOnElevation sets the notificationToUserOnElevation property value. true if send notification to the end user when the role is activated. false if do not send notification when the role is activated.
func (m *PrivilegedRoleSettings) SetNotificationToUserOnElevation(value *bool)() {
    if m != nil {
        m.notificationToUserOnElevation = value
    }
}
// SetTicketingInfoOnElevation sets the ticketingInfoOnElevation property value. true if the ticketing information is required when activate the role. false if the ticketing information is not required when activate the role.
func (m *PrivilegedRoleSettings) SetTicketingInfoOnElevation(value *bool)() {
    if m != nil {
        m.ticketingInfoOnElevation = value
    }
}

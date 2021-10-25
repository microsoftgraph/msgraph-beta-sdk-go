package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type PrivilegedRoleSettings struct {
    Entity
    approvalOnElevation *bool;
    approverIds []string;
    elevationDuration *string;
    isMfaOnElevationConfigurable *bool;
    lastGlobalAdmin *bool;
    maxElavationDuration *string;
    mfaOnElevation *bool;
    minElevationDuration *string;
    notificationToUserOnElevation *bool;
    ticketingInfoOnElevation *bool;
}
func NewPrivilegedRoleSettings()(*PrivilegedRoleSettings) {
    m := &PrivilegedRoleSettings{
        Entity: *NewEntity(),
    }
    return m
}
func (m *PrivilegedRoleSettings) GetApprovalOnElevation()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.approvalOnElevation
    }
}
func (m *PrivilegedRoleSettings) GetApproverIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.approverIds
    }
}
func (m *PrivilegedRoleSettings) GetElevationDuration()(*string) {
    if m == nil {
        return nil
    } else {
        return m.elevationDuration
    }
}
func (m *PrivilegedRoleSettings) GetIsMfaOnElevationConfigurable()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isMfaOnElevationConfigurable
    }
}
func (m *PrivilegedRoleSettings) GetLastGlobalAdmin()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.lastGlobalAdmin
    }
}
func (m *PrivilegedRoleSettings) GetMaxElavationDuration()(*string) {
    if m == nil {
        return nil
    } else {
        return m.maxElavationDuration
    }
}
func (m *PrivilegedRoleSettings) GetMfaOnElevation()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.mfaOnElevation
    }
}
func (m *PrivilegedRoleSettings) GetMinElevationDuration()(*string) {
    if m == nil {
        return nil
    } else {
        return m.minElevationDuration
    }
}
func (m *PrivilegedRoleSettings) GetNotificationToUserOnElevation()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.notificationToUserOnElevation
    }
}
func (m *PrivilegedRoleSettings) GetTicketingInfoOnElevation()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.ticketingInfoOnElevation
    }
}
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
func (m *PrivilegedRoleSettings) SetApprovalOnElevation(value *bool)() {
    m.approvalOnElevation = value
}
func (m *PrivilegedRoleSettings) SetApproverIds(value []string)() {
    m.approverIds = value
}
func (m *PrivilegedRoleSettings) SetElevationDuration(value *string)() {
    m.elevationDuration = value
}
func (m *PrivilegedRoleSettings) SetIsMfaOnElevationConfigurable(value *bool)() {
    m.isMfaOnElevationConfigurable = value
}
func (m *PrivilegedRoleSettings) SetLastGlobalAdmin(value *bool)() {
    m.lastGlobalAdmin = value
}
func (m *PrivilegedRoleSettings) SetMaxElavationDuration(value *string)() {
    m.maxElavationDuration = value
}
func (m *PrivilegedRoleSettings) SetMfaOnElevation(value *bool)() {
    m.mfaOnElevation = value
}
func (m *PrivilegedRoleSettings) SetMinElevationDuration(value *string)() {
    m.minElevationDuration = value
}
func (m *PrivilegedRoleSettings) SetNotificationToUserOnElevation(value *bool)() {
    m.notificationToUserOnElevation = value
}
func (m *PrivilegedRoleSettings) SetTicketingInfoOnElevation(value *bool)() {
    m.ticketingInfoOnElevation = value
}

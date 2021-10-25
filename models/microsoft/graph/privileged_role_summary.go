package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type PrivilegedRoleSummary struct {
    Entity
    elevatedCount *int32;
    managedCount *int32;
    mfaEnabled *bool;
    status *RoleSummaryStatus;
    usersCount *int32;
}
func NewPrivilegedRoleSummary()(*PrivilegedRoleSummary) {
    m := &PrivilegedRoleSummary{
        Entity: *NewEntity(),
    }
    return m
}
func (m *PrivilegedRoleSummary) GetElevatedCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.elevatedCount
    }
}
func (m *PrivilegedRoleSummary) GetManagedCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.managedCount
    }
}
func (m *PrivilegedRoleSummary) GetMfaEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.mfaEnabled
    }
}
func (m *PrivilegedRoleSummary) GetStatus()(*RoleSummaryStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
func (m *PrivilegedRoleSummary) GetUsersCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.usersCount
    }
}
func (m *PrivilegedRoleSummary) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["elevatedCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetElevatedCount(val)
        return nil
    }
    res["managedCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetManagedCount(val)
        return nil
    }
    res["mfaEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetMfaEnabled(val)
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseRoleSummaryStatus)
        if err != nil {
            return err
        }
        cast := val.(RoleSummaryStatus)
        m.SetStatus(&cast)
        return nil
    }
    res["usersCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetUsersCount(val)
        return nil
    }
    return res
}
func (m *PrivilegedRoleSummary) IsNil()(bool) {
    return m == nil
}
func (m *PrivilegedRoleSummary) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("elevatedCount", m.GetElevatedCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("managedCount", m.GetManagedCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("mfaEnabled", m.GetMfaEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := m.GetStatus().String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("usersCount", m.GetUsersCount())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *PrivilegedRoleSummary) SetElevatedCount(value *int32)() {
    m.elevatedCount = value
}
func (m *PrivilegedRoleSummary) SetManagedCount(value *int32)() {
    m.managedCount = value
}
func (m *PrivilegedRoleSummary) SetMfaEnabled(value *bool)() {
    m.mfaEnabled = value
}
func (m *PrivilegedRoleSummary) SetStatus(value *RoleSummaryStatus)() {
    m.status = value
}
func (m *PrivilegedRoleSummary) SetUsersCount(value *int32)() {
    m.usersCount = value
}

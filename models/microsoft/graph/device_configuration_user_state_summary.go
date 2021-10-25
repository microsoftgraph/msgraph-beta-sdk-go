package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type DeviceConfigurationUserStateSummary struct {
    Entity
    compliantUserCount *int32;
    conflictUserCount *int32;
    errorUserCount *int32;
    nonCompliantUserCount *int32;
    notApplicableUserCount *int32;
    remediatedUserCount *int32;
    unknownUserCount *int32;
}
func NewDeviceConfigurationUserStateSummary()(*DeviceConfigurationUserStateSummary) {
    m := &DeviceConfigurationUserStateSummary{
        Entity: *NewEntity(),
    }
    return m
}
func (m *DeviceConfigurationUserStateSummary) GetCompliantUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.compliantUserCount
    }
}
func (m *DeviceConfigurationUserStateSummary) GetConflictUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.conflictUserCount
    }
}
func (m *DeviceConfigurationUserStateSummary) GetErrorUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.errorUserCount
    }
}
func (m *DeviceConfigurationUserStateSummary) GetNonCompliantUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.nonCompliantUserCount
    }
}
func (m *DeviceConfigurationUserStateSummary) GetNotApplicableUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.notApplicableUserCount
    }
}
func (m *DeviceConfigurationUserStateSummary) GetRemediatedUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.remediatedUserCount
    }
}
func (m *DeviceConfigurationUserStateSummary) GetUnknownUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.unknownUserCount
    }
}
func (m *DeviceConfigurationUserStateSummary) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["compliantUserCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetCompliantUserCount(val)
        return nil
    }
    res["conflictUserCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetConflictUserCount(val)
        return nil
    }
    res["errorUserCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetErrorUserCount(val)
        return nil
    }
    res["nonCompliantUserCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetNonCompliantUserCount(val)
        return nil
    }
    res["notApplicableUserCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetNotApplicableUserCount(val)
        return nil
    }
    res["remediatedUserCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetRemediatedUserCount(val)
        return nil
    }
    res["unknownUserCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetUnknownUserCount(val)
        return nil
    }
    return res
}
func (m *DeviceConfigurationUserStateSummary) IsNil()(bool) {
    return m == nil
}
func (m *DeviceConfigurationUserStateSummary) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("compliantUserCount", m.GetCompliantUserCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("conflictUserCount", m.GetConflictUserCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("errorUserCount", m.GetErrorUserCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("nonCompliantUserCount", m.GetNonCompliantUserCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("notApplicableUserCount", m.GetNotApplicableUserCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("remediatedUserCount", m.GetRemediatedUserCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("unknownUserCount", m.GetUnknownUserCount())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *DeviceConfigurationUserStateSummary) SetCompliantUserCount(value *int32)() {
    m.compliantUserCount = value
}
func (m *DeviceConfigurationUserStateSummary) SetConflictUserCount(value *int32)() {
    m.conflictUserCount = value
}
func (m *DeviceConfigurationUserStateSummary) SetErrorUserCount(value *int32)() {
    m.errorUserCount = value
}
func (m *DeviceConfigurationUserStateSummary) SetNonCompliantUserCount(value *int32)() {
    m.nonCompliantUserCount = value
}
func (m *DeviceConfigurationUserStateSummary) SetNotApplicableUserCount(value *int32)() {
    m.notApplicableUserCount = value
}
func (m *DeviceConfigurationUserStateSummary) SetRemediatedUserCount(value *int32)() {
    m.remediatedUserCount = value
}
func (m *DeviceConfigurationUserStateSummary) SetUnknownUserCount(value *int32)() {
    m.unknownUserCount = value
}

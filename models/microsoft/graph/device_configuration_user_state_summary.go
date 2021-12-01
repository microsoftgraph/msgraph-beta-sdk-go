package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceConfigurationUserStateSummary 
type DeviceConfigurationUserStateSummary struct {
    Entity
    // Number of compliant users
    compliantUserCount *int32;
    // Number of conflict users
    conflictUserCount *int32;
    // Number of error users
    errorUserCount *int32;
    // Number of NonCompliant users
    nonCompliantUserCount *int32;
    // Number of not applicable users
    notApplicableUserCount *int32;
    // Number of remediated users
    remediatedUserCount *int32;
    // Number of unknown users
    unknownUserCount *int32;
}
// NewDeviceConfigurationUserStateSummary instantiates a new deviceConfigurationUserStateSummary and sets the default values.
func NewDeviceConfigurationUserStateSummary()(*DeviceConfigurationUserStateSummary) {
    m := &DeviceConfigurationUserStateSummary{
        Entity: *NewEntity(),
    }
    return m
}
// GetCompliantUserCount gets the compliantUserCount property value. Number of compliant users
func (m *DeviceConfigurationUserStateSummary) GetCompliantUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.compliantUserCount
    }
}
// GetConflictUserCount gets the conflictUserCount property value. Number of conflict users
func (m *DeviceConfigurationUserStateSummary) GetConflictUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.conflictUserCount
    }
}
// GetErrorUserCount gets the errorUserCount property value. Number of error users
func (m *DeviceConfigurationUserStateSummary) GetErrorUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.errorUserCount
    }
}
// GetNonCompliantUserCount gets the nonCompliantUserCount property value. Number of NonCompliant users
func (m *DeviceConfigurationUserStateSummary) GetNonCompliantUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.nonCompliantUserCount
    }
}
// GetNotApplicableUserCount gets the notApplicableUserCount property value. Number of not applicable users
func (m *DeviceConfigurationUserStateSummary) GetNotApplicableUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.notApplicableUserCount
    }
}
// GetRemediatedUserCount gets the remediatedUserCount property value. Number of remediated users
func (m *DeviceConfigurationUserStateSummary) GetRemediatedUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.remediatedUserCount
    }
}
// GetUnknownUserCount gets the unknownUserCount property value. Number of unknown users
func (m *DeviceConfigurationUserStateSummary) GetUnknownUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.unknownUserCount
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceConfigurationUserStateSummary) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["compliantUserCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompliantUserCount(val)
        }
        return nil
    }
    res["conflictUserCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConflictUserCount(val)
        }
        return nil
    }
    res["errorUserCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorUserCount(val)
        }
        return nil
    }
    res["nonCompliantUserCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNonCompliantUserCount(val)
        }
        return nil
    }
    res["notApplicableUserCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotApplicableUserCount(val)
        }
        return nil
    }
    res["remediatedUserCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemediatedUserCount(val)
        }
        return nil
    }
    res["unknownUserCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUnknownUserCount(val)
        }
        return nil
    }
    return res
}
func (m *DeviceConfigurationUserStateSummary) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetCompliantUserCount sets the compliantUserCount property value. Number of compliant users
func (m *DeviceConfigurationUserStateSummary) SetCompliantUserCount(value *int32)() {
    if m != nil {
        m.compliantUserCount = value
    }
}
// SetConflictUserCount sets the conflictUserCount property value. Number of conflict users
func (m *DeviceConfigurationUserStateSummary) SetConflictUserCount(value *int32)() {
    if m != nil {
        m.conflictUserCount = value
    }
}
// SetErrorUserCount sets the errorUserCount property value. Number of error users
func (m *DeviceConfigurationUserStateSummary) SetErrorUserCount(value *int32)() {
    if m != nil {
        m.errorUserCount = value
    }
}
// SetNonCompliantUserCount sets the nonCompliantUserCount property value. Number of NonCompliant users
func (m *DeviceConfigurationUserStateSummary) SetNonCompliantUserCount(value *int32)() {
    if m != nil {
        m.nonCompliantUserCount = value
    }
}
// SetNotApplicableUserCount sets the notApplicableUserCount property value. Number of not applicable users
func (m *DeviceConfigurationUserStateSummary) SetNotApplicableUserCount(value *int32)() {
    if m != nil {
        m.notApplicableUserCount = value
    }
}
// SetRemediatedUserCount sets the remediatedUserCount property value. Number of remediated users
func (m *DeviceConfigurationUserStateSummary) SetRemediatedUserCount(value *int32)() {
    if m != nil {
        m.remediatedUserCount = value
    }
}
// SetUnknownUserCount sets the unknownUserCount property value. Number of unknown users
func (m *DeviceConfigurationUserStateSummary) SetUnknownUserCount(value *int32)() {
    if m != nil {
        m.unknownUserCount = value
    }
}

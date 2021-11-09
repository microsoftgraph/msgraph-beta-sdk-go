package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new deviceConfigurationUserStateSummary and sets the default values.
func NewDeviceConfigurationUserStateSummary()(*DeviceConfigurationUserStateSummary) {
    m := &DeviceConfigurationUserStateSummary{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the compliantUserCount property value. Number of compliant users
func (m *DeviceConfigurationUserStateSummary) GetCompliantUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.compliantUserCount
    }
}
// Gets the conflictUserCount property value. Number of conflict users
func (m *DeviceConfigurationUserStateSummary) GetConflictUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.conflictUserCount
    }
}
// Gets the errorUserCount property value. Number of error users
func (m *DeviceConfigurationUserStateSummary) GetErrorUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.errorUserCount
    }
}
// Gets the nonCompliantUserCount property value. Number of NonCompliant users
func (m *DeviceConfigurationUserStateSummary) GetNonCompliantUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.nonCompliantUserCount
    }
}
// Gets the notApplicableUserCount property value. Number of not applicable users
func (m *DeviceConfigurationUserStateSummary) GetNotApplicableUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.notApplicableUserCount
    }
}
// Gets the remediatedUserCount property value. Number of remediated users
func (m *DeviceConfigurationUserStateSummary) GetRemediatedUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.remediatedUserCount
    }
}
// Gets the unknownUserCount property value. Number of unknown users
func (m *DeviceConfigurationUserStateSummary) GetUnknownUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.unknownUserCount
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the compliantUserCount property value. Number of compliant users
// Parameters:
//  - value : Value to set for the compliantUserCount property.
func (m *DeviceConfigurationUserStateSummary) SetCompliantUserCount(value *int32)() {
    m.compliantUserCount = value
}
// Sets the conflictUserCount property value. Number of conflict users
// Parameters:
//  - value : Value to set for the conflictUserCount property.
func (m *DeviceConfigurationUserStateSummary) SetConflictUserCount(value *int32)() {
    m.conflictUserCount = value
}
// Sets the errorUserCount property value. Number of error users
// Parameters:
//  - value : Value to set for the errorUserCount property.
func (m *DeviceConfigurationUserStateSummary) SetErrorUserCount(value *int32)() {
    m.errorUserCount = value
}
// Sets the nonCompliantUserCount property value. Number of NonCompliant users
// Parameters:
//  - value : Value to set for the nonCompliantUserCount property.
func (m *DeviceConfigurationUserStateSummary) SetNonCompliantUserCount(value *int32)() {
    m.nonCompliantUserCount = value
}
// Sets the notApplicableUserCount property value. Number of not applicable users
// Parameters:
//  - value : Value to set for the notApplicableUserCount property.
func (m *DeviceConfigurationUserStateSummary) SetNotApplicableUserCount(value *int32)() {
    m.notApplicableUserCount = value
}
// Sets the remediatedUserCount property value. Number of remediated users
// Parameters:
//  - value : Value to set for the remediatedUserCount property.
func (m *DeviceConfigurationUserStateSummary) SetRemediatedUserCount(value *int32)() {
    m.remediatedUserCount = value
}
// Sets the unknownUserCount property value. Number of unknown users
// Parameters:
//  - value : Value to set for the unknownUserCount property.
func (m *DeviceConfigurationUserStateSummary) SetUnknownUserCount(value *int32)() {
    m.unknownUserCount = value
}

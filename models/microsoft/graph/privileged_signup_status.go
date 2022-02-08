package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PrivilegedSignupStatus 
type PrivilegedSignupStatus struct {
    Entity
    // 
    isRegistered *bool;
    // 
    status *SetupStatus;
}
// NewPrivilegedSignupStatus instantiates a new privilegedSignupStatus and sets the default values.
func NewPrivilegedSignupStatus()(*PrivilegedSignupStatus) {
    m := &PrivilegedSignupStatus{
        Entity: *NewEntity(),
    }
    return m
}
// GetIsRegistered gets the isRegistered property value. 
func (m *PrivilegedSignupStatus) GetIsRegistered()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isRegistered
    }
}
// GetStatus gets the status property value. 
func (m *PrivilegedSignupStatus) GetStatus()(*SetupStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PrivilegedSignupStatus) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["isRegistered"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsRegistered(val)
        }
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseSetupStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*SetupStatus))
        }
        return nil
    }
    return res
}
func (m *PrivilegedSignupStatus) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *PrivilegedSignupStatus) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("isRegistered", m.GetIsRegistered())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetIsRegistered sets the isRegistered property value. 
func (m *PrivilegedSignupStatus) SetIsRegistered(value *bool)() {
    if m != nil {
        m.isRegistered = value
    }
}
// SetStatus sets the status property value. 
func (m *PrivilegedSignupStatus) SetStatus(value *SetupStatus)() {
    if m != nil {
        m.status = value
    }
}

package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type PrivilegedSignupStatus struct {
    Entity
    // 
    isRegistered *bool;
    // 
    status *SetupStatus;
}
// Instantiates a new privilegedSignupStatus and sets the default values.
func NewPrivilegedSignupStatus()(*PrivilegedSignupStatus) {
    m := &PrivilegedSignupStatus{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the isRegistered property value. 
func (m *PrivilegedSignupStatus) GetIsRegistered()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isRegistered
    }
}
// Gets the status property value. 
func (m *PrivilegedSignupStatus) GetStatus()(*SetupStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// The deserialization information for the current model
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
            cast := val.(SetupStatus)
            m.SetStatus(&cast)
        }
        return nil
    }
    return res
}
func (m *PrivilegedSignupStatus) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
        cast := m.GetStatus().String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the isRegistered property value. 
// Parameters:
//  - value : Value to set for the isRegistered property.
func (m *PrivilegedSignupStatus) SetIsRegistered(value *bool)() {
    m.isRegistered = value
}
// Sets the status property value. 
// Parameters:
//  - value : Value to set for the status property.
func (m *PrivilegedSignupStatus) SetStatus(value *SetupStatus)() {
    m.status = value
}

package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ExternalIdentitiesPolicy 
type ExternalIdentitiesPolicy struct {
    PolicyBase
    // 
    allowDeletedIdentitiesDataRemoval *bool;
    // 
    allowExternalIdentitiesToLeave *bool;
}
// NewExternalIdentitiesPolicy instantiates a new externalIdentitiesPolicy and sets the default values.
func NewExternalIdentitiesPolicy()(*ExternalIdentitiesPolicy) {
    m := &ExternalIdentitiesPolicy{
        PolicyBase: *NewPolicyBase(),
    }
    return m
}
// GetAllowDeletedIdentitiesDataRemoval gets the allowDeletedIdentitiesDataRemoval property value. 
func (m *ExternalIdentitiesPolicy) GetAllowDeletedIdentitiesDataRemoval()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowDeletedIdentitiesDataRemoval
    }
}
// GetAllowExternalIdentitiesToLeave gets the allowExternalIdentitiesToLeave property value. 
func (m *ExternalIdentitiesPolicy) GetAllowExternalIdentitiesToLeave()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowExternalIdentitiesToLeave
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ExternalIdentitiesPolicy) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.PolicyBase.GetFieldDeserializers()
    res["allowDeletedIdentitiesDataRemoval"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowDeletedIdentitiesDataRemoval(val)
        }
        return nil
    }
    res["allowExternalIdentitiesToLeave"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowExternalIdentitiesToLeave(val)
        }
        return nil
    }
    return res
}
func (m *ExternalIdentitiesPolicy) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ExternalIdentitiesPolicy) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.PolicyBase.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("allowDeletedIdentitiesDataRemoval", m.GetAllowDeletedIdentitiesDataRemoval())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("allowExternalIdentitiesToLeave", m.GetAllowExternalIdentitiesToLeave())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAllowDeletedIdentitiesDataRemoval sets the allowDeletedIdentitiesDataRemoval property value. 
func (m *ExternalIdentitiesPolicy) SetAllowDeletedIdentitiesDataRemoval(value *bool)() {
    if m != nil {
        m.allowDeletedIdentitiesDataRemoval = value
    }
}
// SetAllowExternalIdentitiesToLeave sets the allowExternalIdentitiesToLeave property value. 
func (m *ExternalIdentitiesPolicy) SetAllowExternalIdentitiesToLeave(value *bool)() {
    if m != nil {
        m.allowExternalIdentitiesToLeave = value
    }
}

package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ServicePrincipalCreationPolicy 
type ServicePrincipalCreationPolicy struct {
    PolicyBase
    // 
    excludes []ServicePrincipalCreationConditionSet;
    // 
    includes []ServicePrincipalCreationConditionSet;
    // 
    isBuiltIn *bool;
}
// NewServicePrincipalCreationPolicy instantiates a new servicePrincipalCreationPolicy and sets the default values.
func NewServicePrincipalCreationPolicy()(*ServicePrincipalCreationPolicy) {
    m := &ServicePrincipalCreationPolicy{
        PolicyBase: *NewPolicyBase(),
    }
    return m
}
// GetExcludes gets the excludes property value. 
func (m *ServicePrincipalCreationPolicy) GetExcludes()([]ServicePrincipalCreationConditionSet) {
    if m == nil {
        return nil
    } else {
        return m.excludes
    }
}
// GetIncludes gets the includes property value. 
func (m *ServicePrincipalCreationPolicy) GetIncludes()([]ServicePrincipalCreationConditionSet) {
    if m == nil {
        return nil
    } else {
        return m.includes
    }
}
// GetIsBuiltIn gets the isBuiltIn property value. 
func (m *ServicePrincipalCreationPolicy) GetIsBuiltIn()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isBuiltIn
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ServicePrincipalCreationPolicy) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.PolicyBase.GetFieldDeserializers()
    res["excludes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewServicePrincipalCreationConditionSet() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ServicePrincipalCreationConditionSet, len(val))
            for i, v := range val {
                res[i] = *(v.(*ServicePrincipalCreationConditionSet))
            }
            m.SetExcludes(res)
        }
        return nil
    }
    res["includes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewServicePrincipalCreationConditionSet() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ServicePrincipalCreationConditionSet, len(val))
            for i, v := range val {
                res[i] = *(v.(*ServicePrincipalCreationConditionSet))
            }
            m.SetIncludes(res)
        }
        return nil
    }
    res["isBuiltIn"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsBuiltIn(val)
        }
        return nil
    }
    return res
}
func (m *ServicePrincipalCreationPolicy) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ServicePrincipalCreationPolicy) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.PolicyBase.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetExcludes() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetExcludes()))
        for i, v := range m.GetExcludes() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("excludes", cast)
        if err != nil {
            return err
        }
    }
    if m.GetIncludes() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetIncludes()))
        for i, v := range m.GetIncludes() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("includes", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isBuiltIn", m.GetIsBuiltIn())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetExcludes sets the excludes property value. 
func (m *ServicePrincipalCreationPolicy) SetExcludes(value []ServicePrincipalCreationConditionSet)() {
    if m != nil {
        m.excludes = value
    }
}
// SetIncludes sets the includes property value. 
func (m *ServicePrincipalCreationPolicy) SetIncludes(value []ServicePrincipalCreationConditionSet)() {
    if m != nil {
        m.includes = value
    }
}
// SetIsBuiltIn sets the isBuiltIn property value. 
func (m *ServicePrincipalCreationPolicy) SetIsBuiltIn(value *bool)() {
    if m != nil {
        m.isBuiltIn = value
    }
}

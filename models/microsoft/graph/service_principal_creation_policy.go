package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ServicePrincipalCreationPolicy 
type ServicePrincipalCreationPolicy struct {
    PolicyBase
    // 
    excludes []ServicePrincipalCreationConditionSetable;
    // 
    includes []ServicePrincipalCreationConditionSetable;
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
// CreateServicePrincipalCreationPolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateServicePrincipalCreationPolicyFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewServicePrincipalCreationPolicy(), nil
}
// GetExcludes gets the excludes property value. 
func (m *ServicePrincipalCreationPolicy) GetExcludes()([]ServicePrincipalCreationConditionSetable) {
    if m == nil {
        return nil
    } else {
        return m.excludes
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ServicePrincipalCreationPolicy) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.PolicyBase.GetFieldDeserializers()
    res["excludes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateServicePrincipalCreationConditionSetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ServicePrincipalCreationConditionSetable, len(val))
            for i, v := range val {
                res[i] = v.(ServicePrincipalCreationConditionSetable)
            }
            m.SetExcludes(res)
        }
        return nil
    }
    res["includes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateServicePrincipalCreationConditionSetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ServicePrincipalCreationConditionSetable, len(val))
            for i, v := range val {
                res[i] = v.(ServicePrincipalCreationConditionSetable)
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
// GetIncludes gets the includes property value. 
func (m *ServicePrincipalCreationPolicy) GetIncludes()([]ServicePrincipalCreationConditionSetable) {
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
// Serialize serializes information the current object
func (m *ServicePrincipalCreationPolicy) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.PolicyBase.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetExcludes() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetExcludes()))
        for i, v := range m.GetExcludes() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("excludes", cast)
        if err != nil {
            return err
        }
    }
    if m.GetIncludes() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetIncludes()))
        for i, v := range m.GetIncludes() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
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
func (m *ServicePrincipalCreationPolicy) SetExcludes(value []ServicePrincipalCreationConditionSetable)() {
    if m != nil {
        m.excludes = value
    }
}
// SetIncludes sets the includes property value. 
func (m *ServicePrincipalCreationPolicy) SetIncludes(value []ServicePrincipalCreationConditionSetable)() {
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

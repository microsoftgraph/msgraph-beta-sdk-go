package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TenantRelationshipAccessPolicyBase 
type TenantRelationshipAccessPolicyBase struct {
    PolicyBase
    // 
    definition []string;
}
// NewTenantRelationshipAccessPolicyBase instantiates a new tenantRelationshipAccessPolicyBase and sets the default values.
func NewTenantRelationshipAccessPolicyBase()(*TenantRelationshipAccessPolicyBase) {
    m := &TenantRelationshipAccessPolicyBase{
        PolicyBase: *NewPolicyBase(),
    }
    return m
}
// GetDefinition gets the definition property value. 
func (m *TenantRelationshipAccessPolicyBase) GetDefinition()([]string) {
    if m == nil {
        return nil
    } else {
        return m.definition
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TenantRelationshipAccessPolicyBase) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.PolicyBase.GetFieldDeserializers()
    res["definition"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetDefinition(res)
        }
        return nil
    }
    return res
}
func (m *TenantRelationshipAccessPolicyBase) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *TenantRelationshipAccessPolicyBase) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.PolicyBase.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetDefinition() != nil {
        err = writer.WriteCollectionOfStringValues("definition", m.GetDefinition())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDefinition sets the definition property value. 
func (m *TenantRelationshipAccessPolicyBase) SetDefinition(value []string)() {
    if m != nil {
        m.definition = value
    }
}

package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type TrustFramework struct {
    additionalData map[string]interface{};
    keySets []TrustFrameworkKeySet;
    policies []TrustFrameworkPolicy;
}
func NewTrustFramework()(*TrustFramework) {
    m := &TrustFramework{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *TrustFramework) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *TrustFramework) GetKeySets()([]TrustFrameworkKeySet) {
    if m == nil {
        return nil
    } else {
        return m.keySets
    }
}
func (m *TrustFramework) GetPolicies()([]TrustFrameworkPolicy) {
    if m == nil {
        return nil
    } else {
        return m.policies
    }
}
func (m *TrustFramework) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["keySets"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTrustFrameworkKeySet() })
        if err != nil {
            return err
        }
        res := make([]TrustFrameworkKeySet, len(val))
        for i, v := range val {
            res[i] = *(v.(*TrustFrameworkKeySet))
        }
        m.SetKeySets(res)
        return nil
    }
    res["policies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTrustFrameworkPolicy() })
        if err != nil {
            return err
        }
        res := make([]TrustFrameworkPolicy, len(val))
        for i, v := range val {
            res[i] = *(v.(*TrustFrameworkPolicy))
        }
        m.SetPolicies(res)
        return nil
    }
    return res
}
func (m *TrustFramework) IsNil()(bool) {
    return m == nil
}
func (m *TrustFramework) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetKeySets()))
        for i, v := range m.GetKeySets() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("keySets", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetPolicies()))
        for i, v := range m.GetPolicies() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("policies", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *TrustFramework) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *TrustFramework) SetKeySets(value []TrustFrameworkKeySet)() {
    m.keySets = value
}
func (m *TrustFramework) SetPolicies(value []TrustFrameworkPolicy)() {
    m.policies = value
}

package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type TrustFramework struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    keySets []TrustFrameworkKeySet;
    // 
    policies []TrustFrameworkPolicy;
}
// Instantiates a new TrustFramework and sets the default values.
func NewTrustFramework()(*TrustFramework) {
    m := &TrustFramework{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TrustFramework) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the keySets property value. 
func (m *TrustFramework) GetKeySets()([]TrustFrameworkKeySet) {
    if m == nil {
        return nil
    } else {
        return m.keySets
    }
}
// Gets the policies property value. 
func (m *TrustFramework) GetPolicies()([]TrustFrameworkPolicy) {
    if m == nil {
        return nil
    } else {
        return m.policies
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *TrustFramework) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the keySets property value. 
// Parameters:
//  - value : Value to set for the keySets property.
func (m *TrustFramework) SetKeySets(value []TrustFrameworkKeySet)() {
    m.keySets = value
}
// Sets the policies property value. 
// Parameters:
//  - value : Value to set for the policies property.
func (m *TrustFramework) SetPolicies(value []TrustFrameworkPolicy)() {
    m.policies = value
}

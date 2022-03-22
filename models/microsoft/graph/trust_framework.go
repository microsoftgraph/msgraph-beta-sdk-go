package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TrustFramework 
type TrustFramework struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    keySets []TrustFrameworkKeySetable;
    // 
    policies []TrustFrameworkPolicyable;
}
// NewTrustFramework instantiates a new TrustFramework and sets the default values.
func NewTrustFramework()(*TrustFramework) {
    m := &TrustFramework{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateTrustFrameworkFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTrustFrameworkFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewTrustFramework(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TrustFramework) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TrustFramework) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["keySets"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTrustFrameworkKeySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TrustFrameworkKeySetable, len(val))
            for i, v := range val {
                res[i] = v.(TrustFrameworkKeySetable)
            }
            m.SetKeySets(res)
        }
        return nil
    }
    res["policies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTrustFrameworkPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TrustFrameworkPolicyable, len(val))
            for i, v := range val {
                res[i] = v.(TrustFrameworkPolicyable)
            }
            m.SetPolicies(res)
        }
        return nil
    }
    return res
}
// GetKeySets gets the keySets property value. 
func (m *TrustFramework) GetKeySets()([]TrustFrameworkKeySetable) {
    if m == nil {
        return nil
    } else {
        return m.keySets
    }
}
// GetPolicies gets the policies property value. 
func (m *TrustFramework) GetPolicies()([]TrustFrameworkPolicyable) {
    if m == nil {
        return nil
    } else {
        return m.policies
    }
}
// Serialize serializes information the current object
func (m *TrustFramework) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetKeySets() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetKeySets()))
        for i, v := range m.GetKeySets() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("keySets", cast)
        if err != nil {
            return err
        }
    }
    if m.GetPolicies() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetPolicies()))
        for i, v := range m.GetPolicies() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TrustFramework) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetKeySets sets the keySets property value. 
func (m *TrustFramework) SetKeySets(value []TrustFrameworkKeySetable)() {
    if m != nil {
        m.keySets = value
    }
}
// SetPolicies sets the policies property value. 
func (m *TrustFramework) SetPolicies(value []TrustFrameworkPolicyable)() {
    if m != nil {
        m.policies = value
    }
}

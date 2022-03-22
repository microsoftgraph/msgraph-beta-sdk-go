package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TrustFrameworkPolicy 
type TrustFrameworkPolicy struct {
    Entity
}
// NewTrustFrameworkPolicy instantiates a new trustFrameworkPolicy and sets the default values.
func NewTrustFrameworkPolicy()(*TrustFrameworkPolicy) {
    m := &TrustFrameworkPolicy{
        Entity: *NewEntity(),
    }
    return m
}
// CreateTrustFrameworkPolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTrustFrameworkPolicyFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewTrustFrameworkPolicy(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TrustFrameworkPolicy) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *TrustFrameworkPolicy) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}

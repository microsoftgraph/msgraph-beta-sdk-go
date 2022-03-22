package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TeamworkTagIdentity 
type TeamworkTagIdentity struct {
    Identity
}
// NewTeamworkTagIdentity instantiates a new teamworkTagIdentity and sets the default values.
func NewTeamworkTagIdentity()(*TeamworkTagIdentity) {
    m := &TeamworkTagIdentity{
        Identity: *NewIdentity(),
    }
    return m
}
// CreateTeamworkTagIdentityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamworkTagIdentityFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewTeamworkTagIdentity(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamworkTagIdentity) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Identity.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *TeamworkTagIdentity) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Identity.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}

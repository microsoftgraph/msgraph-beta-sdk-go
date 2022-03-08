package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PlannerDelta provides operations to manage the compliance singleton.
type PlannerDelta struct {
    Entity
}
// NewPlannerDelta instantiates a new plannerDelta and sets the default values.
func NewPlannerDelta()(*PlannerDelta) {
    m := &PlannerDelta{
        Entity: *NewEntity(),
    }
    return m
}
// CreatePlannerDeltaFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePlannerDeltaFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewPlannerDelta(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PlannerDelta) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    return res
}
func (m *PlannerDelta) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *PlannerDelta) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}

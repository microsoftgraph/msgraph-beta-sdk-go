package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UpdatableAsset 
type UpdatableAsset struct {
    Entity
}
// NewUpdatableAsset instantiates a new updatableAsset and sets the default values.
func NewUpdatableAsset()(*UpdatableAsset) {
    m := &UpdatableAsset{
        Entity: *NewEntity(),
    }
    return m
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UpdatableAsset) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    return res
}
func (m *UpdatableAsset) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *UpdatableAsset) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}

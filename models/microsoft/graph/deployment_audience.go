package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type DeploymentAudience struct {
    Entity
    // Specifies the assets to exclude from the audience.
    exclusions []UpdatableAsset;
    // Specifies the assets to include in the audience.
    members []UpdatableAsset;
}
// Instantiates a new deploymentAudience and sets the default values.
func NewDeploymentAudience()(*DeploymentAudience) {
    m := &DeploymentAudience{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the exclusions property value. Specifies the assets to exclude from the audience.
func (m *DeploymentAudience) GetExclusions()([]UpdatableAsset) {
    if m == nil {
        return nil
    } else {
        return m.exclusions
    }
}
// Gets the members property value. Specifies the assets to include in the audience.
func (m *DeploymentAudience) GetMembers()([]UpdatableAsset) {
    if m == nil {
        return nil
    } else {
        return m.members
    }
}
// The deserialization information for the current model
func (m *DeploymentAudience) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["exclusions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUpdatableAsset() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UpdatableAsset, len(val))
            for i, v := range val {
                res[i] = *(v.(*UpdatableAsset))
            }
            m.SetExclusions(res)
        }
        return nil
    }
    res["members"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUpdatableAsset() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UpdatableAsset, len(val))
            for i, v := range val {
                res[i] = *(v.(*UpdatableAsset))
            }
            m.SetMembers(res)
        }
        return nil
    }
    return res
}
func (m *DeploymentAudience) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *DeploymentAudience) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetExclusions()))
        for i, v := range m.GetExclusions() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("exclusions", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMembers()))
        for i, v := range m.GetMembers() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("members", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the exclusions property value. Specifies the assets to exclude from the audience.
// Parameters:
//  - value : Value to set for the exclusions property.
func (m *DeploymentAudience) SetExclusions(value []UpdatableAsset)() {
    m.exclusions = value
}
// Sets the members property value. Specifies the assets to include in the audience.
// Parameters:
//  - value : Value to set for the members property.
func (m *DeploymentAudience) SetMembers(value []UpdatableAsset)() {
    m.members = value
}

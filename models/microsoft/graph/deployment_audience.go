package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type DeploymentAudience struct {
    Entity
    exclusions []UpdatableAsset;
    members []UpdatableAsset;
}
func NewDeploymentAudience()(*DeploymentAudience) {
    m := &DeploymentAudience{
        Entity: *NewEntity(),
    }
    return m
}
func (m *DeploymentAudience) GetExclusions()([]UpdatableAsset) {
    if m == nil {
        return nil
    } else {
        return m.exclusions
    }
}
func (m *DeploymentAudience) GetMembers()([]UpdatableAsset) {
    if m == nil {
        return nil
    } else {
        return m.members
    }
}
func (m *DeploymentAudience) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["exclusions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUpdatableAsset() })
        if err != nil {
            return err
        }
        res := make([]UpdatableAsset, len(val))
        for i, v := range val {
            res[i] = *(v.(*UpdatableAsset))
        }
        m.SetExclusions(res)
        return nil
    }
    res["members"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUpdatableAsset() })
        if err != nil {
            return err
        }
        res := make([]UpdatableAsset, len(val))
        for i, v := range val {
            res[i] = *(v.(*UpdatableAsset))
        }
        m.SetMembers(res)
        return nil
    }
    return res
}
func (m *DeploymentAudience) IsNil()(bool) {
    return m == nil
}
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
func (m *DeploymentAudience) SetExclusions(value []UpdatableAsset)() {
    m.exclusions = value
}
func (m *DeploymentAudience) SetMembers(value []UpdatableAsset)() {
    m.members = value
}

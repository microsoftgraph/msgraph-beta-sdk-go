package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type SynchronizationSchema struct {
    Entity
    directories []DirectoryDefinition;
    synchronizationRules []SynchronizationRule;
    version *string;
}
func NewSynchronizationSchema()(*SynchronizationSchema) {
    m := &SynchronizationSchema{
        Entity: *NewEntity(),
    }
    return m
}
func (m *SynchronizationSchema) GetDirectories()([]DirectoryDefinition) {
    if m == nil {
        return nil
    } else {
        return m.directories
    }
}
func (m *SynchronizationSchema) GetSynchronizationRules()([]SynchronizationRule) {
    if m == nil {
        return nil
    } else {
        return m.synchronizationRules
    }
}
func (m *SynchronizationSchema) GetVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.version
    }
}
func (m *SynchronizationSchema) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["directories"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDirectoryDefinition() })
        if err != nil {
            return err
        }
        res := make([]DirectoryDefinition, len(val))
        for i, v := range val {
            res[i] = *(v.(*DirectoryDefinition))
        }
        m.SetDirectories(res)
        return nil
    }
    res["synchronizationRules"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSynchronizationRule() })
        if err != nil {
            return err
        }
        res := make([]SynchronizationRule, len(val))
        for i, v := range val {
            res[i] = *(v.(*SynchronizationRule))
        }
        m.SetSynchronizationRules(res)
        return nil
    }
    res["version"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetVersion(val)
        return nil
    }
    return res
}
func (m *SynchronizationSchema) IsNil()(bool) {
    return m == nil
}
func (m *SynchronizationSchema) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDirectories()))
        for i, v := range m.GetDirectories() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("directories", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSynchronizationRules()))
        for i, v := range m.GetSynchronizationRules() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("synchronizationRules", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("version", m.GetVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *SynchronizationSchema) SetDirectories(value []DirectoryDefinition)() {
    m.directories = value
}
func (m *SynchronizationSchema) SetSynchronizationRules(value []SynchronizationRule)() {
    m.synchronizationRules = value
}
func (m *SynchronizationSchema) SetVersion(value *string)() {
    m.version = value
}

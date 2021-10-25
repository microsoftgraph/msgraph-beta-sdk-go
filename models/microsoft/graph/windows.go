package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type Windows struct {
    Entity
    updates *Updates;
}
func NewWindows()(*Windows) {
    m := &Windows{
        Entity: *NewEntity(),
    }
    return m
}
func (m *Windows) GetUpdates()(*Updates) {
    if m == nil {
        return nil
    } else {
        return m.updates
    }
}
func (m *Windows) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["updates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUpdates() })
        if err != nil {
            return err
        }
        m.SetUpdates(val.(*Updates))
        return nil
    }
    return res
}
func (m *Windows) IsNil()(bool) {
    return m == nil
}
func (m *Windows) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("updates", m.GetUpdates())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *Windows) SetUpdates(value *Updates)() {
    m.updates = value
}

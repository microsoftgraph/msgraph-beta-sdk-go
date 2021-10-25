package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ProfileCardProperty struct {
    Entity
    annotations []ProfileCardAnnotation;
    directoryPropertyName *string;
}
func NewProfileCardProperty()(*ProfileCardProperty) {
    m := &ProfileCardProperty{
        Entity: *NewEntity(),
    }
    return m
}
func (m *ProfileCardProperty) GetAnnotations()([]ProfileCardAnnotation) {
    if m == nil {
        return nil
    } else {
        return m.annotations
    }
}
func (m *ProfileCardProperty) GetDirectoryPropertyName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.directoryPropertyName
    }
}
func (m *ProfileCardProperty) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["annotations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewProfileCardAnnotation() })
        if err != nil {
            return err
        }
        res := make([]ProfileCardAnnotation, len(val))
        for i, v := range val {
            res[i] = *(v.(*ProfileCardAnnotation))
        }
        m.SetAnnotations(res)
        return nil
    }
    res["directoryPropertyName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDirectoryPropertyName(val)
        return nil
    }
    return res
}
func (m *ProfileCardProperty) IsNil()(bool) {
    return m == nil
}
func (m *ProfileCardProperty) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAnnotations()))
        for i, v := range m.GetAnnotations() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("annotations", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("directoryPropertyName", m.GetDirectoryPropertyName())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *ProfileCardProperty) SetAnnotations(value []ProfileCardAnnotation)() {
    m.annotations = value
}
func (m *ProfileCardProperty) SetDirectoryPropertyName(value *string)() {
    m.directoryPropertyName = value
}

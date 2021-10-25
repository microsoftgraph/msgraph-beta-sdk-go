package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type Filter struct {
    additionalData map[string]interface{};
    categoryFilterGroups []FilterGroup;
    groups []FilterGroup;
    inputFilterGroups []FilterGroup;
}
func NewFilter()(*Filter) {
    m := &Filter{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *Filter) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *Filter) GetCategoryFilterGroups()([]FilterGroup) {
    if m == nil {
        return nil
    } else {
        return m.categoryFilterGroups
    }
}
func (m *Filter) GetGroups()([]FilterGroup) {
    if m == nil {
        return nil
    } else {
        return m.groups
    }
}
func (m *Filter) GetInputFilterGroups()([]FilterGroup) {
    if m == nil {
        return nil
    } else {
        return m.inputFilterGroups
    }
}
func (m *Filter) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["categoryFilterGroups"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewFilterGroup() })
        if err != nil {
            return err
        }
        res := make([]FilterGroup, len(val))
        for i, v := range val {
            res[i] = *(v.(*FilterGroup))
        }
        m.SetCategoryFilterGroups(res)
        return nil
    }
    res["groups"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewFilterGroup() })
        if err != nil {
            return err
        }
        res := make([]FilterGroup, len(val))
        for i, v := range val {
            res[i] = *(v.(*FilterGroup))
        }
        m.SetGroups(res)
        return nil
    }
    res["inputFilterGroups"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewFilterGroup() })
        if err != nil {
            return err
        }
        res := make([]FilterGroup, len(val))
        for i, v := range val {
            res[i] = *(v.(*FilterGroup))
        }
        m.SetInputFilterGroups(res)
        return nil
    }
    return res
}
func (m *Filter) IsNil()(bool) {
    return m == nil
}
func (m *Filter) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetCategoryFilterGroups()))
        for i, v := range m.GetCategoryFilterGroups() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("categoryFilterGroups", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetGroups()))
        for i, v := range m.GetGroups() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("groups", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetInputFilterGroups()))
        for i, v := range m.GetInputFilterGroups() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("inputFilterGroups", cast)
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
func (m *Filter) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *Filter) SetCategoryFilterGroups(value []FilterGroup)() {
    m.categoryFilterGroups = value
}
func (m *Filter) SetGroups(value []FilterGroup)() {
    m.groups = value
}
func (m *Filter) SetInputFilterGroups(value []FilterGroup)() {
    m.inputFilterGroups = value
}

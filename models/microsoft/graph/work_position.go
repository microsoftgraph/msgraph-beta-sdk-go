package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type WorkPosition struct {
    ItemFacet
    categories []string;
    colleagues []RelatedPerson;
    detail *PositionDetail;
    isCurrent *bool;
    manager *RelatedPerson;
}
func NewWorkPosition()(*WorkPosition) {
    m := &WorkPosition{
        ItemFacet: *NewItemFacet(),
    }
    return m
}
func (m *WorkPosition) GetCategories()([]string) {
    if m == nil {
        return nil
    } else {
        return m.categories
    }
}
func (m *WorkPosition) GetColleagues()([]RelatedPerson) {
    if m == nil {
        return nil
    } else {
        return m.colleagues
    }
}
func (m *WorkPosition) GetDetail()(*PositionDetail) {
    if m == nil {
        return nil
    } else {
        return m.detail
    }
}
func (m *WorkPosition) GetIsCurrent()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isCurrent
    }
}
func (m *WorkPosition) GetManager()(*RelatedPerson) {
    if m == nil {
        return nil
    } else {
        return m.manager
    }
}
func (m *WorkPosition) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.ItemFacet.GetFieldDeserializers()
    res["categories"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetCategories(res)
        return nil
    }
    res["colleagues"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRelatedPerson() })
        if err != nil {
            return err
        }
        res := make([]RelatedPerson, len(val))
        for i, v := range val {
            res[i] = *(v.(*RelatedPerson))
        }
        m.SetColleagues(res)
        return nil
    }
    res["detail"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPositionDetail() })
        if err != nil {
            return err
        }
        m.SetDetail(val.(*PositionDetail))
        return nil
    }
    res["isCurrent"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsCurrent(val)
        return nil
    }
    res["manager"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRelatedPerson() })
        if err != nil {
            return err
        }
        m.SetManager(val.(*RelatedPerson))
        return nil
    }
    return res
}
func (m *WorkPosition) IsNil()(bool) {
    return m == nil
}
func (m *WorkPosition) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.ItemFacet.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteCollectionOfStringValues("categories", m.GetCategories())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetColleagues()))
        for i, v := range m.GetColleagues() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("colleagues", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("detail", m.GetDetail())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isCurrent", m.GetIsCurrent())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("manager", m.GetManager())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *WorkPosition) SetCategories(value []string)() {
    m.categories = value
}
func (m *WorkPosition) SetColleagues(value []RelatedPerson)() {
    m.colleagues = value
}
func (m *WorkPosition) SetDetail(value *PositionDetail)() {
    m.detail = value
}
func (m *WorkPosition) SetIsCurrent(value *bool)() {
    m.isCurrent = value
}
func (m *WorkPosition) SetManager(value *RelatedPerson)() {
    m.manager = value
}

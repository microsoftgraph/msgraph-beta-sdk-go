package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type WorkPosition struct {
    ItemFacet
    // Categories that the user has associated with this position.
    categories []string;
    // Colleagues that are associated with this position.
    colleagues []RelatedPerson;
    // 
    detail *PositionDetail;
    // Denotes whether or not the position is current.
    isCurrent *bool;
    // Contains detail of the user's manager in this position.
    manager *RelatedPerson;
}
// Instantiates a new workPosition and sets the default values.
func NewWorkPosition()(*WorkPosition) {
    m := &WorkPosition{
        ItemFacet: *NewItemFacet(),
    }
    return m
}
// Gets the categories property value. Categories that the user has associated with this position.
func (m *WorkPosition) GetCategories()([]string) {
    if m == nil {
        return nil
    } else {
        return m.categories
    }
}
// Gets the colleagues property value. Colleagues that are associated with this position.
func (m *WorkPosition) GetColleagues()([]RelatedPerson) {
    if m == nil {
        return nil
    } else {
        return m.colleagues
    }
}
// Gets the detail property value. 
func (m *WorkPosition) GetDetail()(*PositionDetail) {
    if m == nil {
        return nil
    } else {
        return m.detail
    }
}
// Gets the isCurrent property value. Denotes whether or not the position is current.
func (m *WorkPosition) GetIsCurrent()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isCurrent
    }
}
// Gets the manager property value. Contains detail of the user's manager in this position.
func (m *WorkPosition) GetManager()(*RelatedPerson) {
    if m == nil {
        return nil
    } else {
        return m.manager
    }
}
// The deserialization information for the current model
func (m *WorkPosition) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.ItemFacet.GetFieldDeserializers()
    res["categories"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetCategories(res)
        }
        return nil
    }
    res["colleagues"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRelatedPerson() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RelatedPerson, len(val))
            for i, v := range val {
                res[i] = *(v.(*RelatedPerson))
            }
            m.SetColleagues(res)
        }
        return nil
    }
    res["detail"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPositionDetail() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetail(val.(*PositionDetail))
        }
        return nil
    }
    res["isCurrent"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsCurrent(val)
        }
        return nil
    }
    res["manager"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRelatedPerson() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManager(val.(*RelatedPerson))
        }
        return nil
    }
    return res
}
func (m *WorkPosition) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the categories property value. Categories that the user has associated with this position.
// Parameters:
//  - value : Value to set for the categories property.
func (m *WorkPosition) SetCategories(value []string)() {
    m.categories = value
}
// Sets the colleagues property value. Colleagues that are associated with this position.
// Parameters:
//  - value : Value to set for the colleagues property.
func (m *WorkPosition) SetColleagues(value []RelatedPerson)() {
    m.colleagues = value
}
// Sets the detail property value. 
// Parameters:
//  - value : Value to set for the detail property.
func (m *WorkPosition) SetDetail(value *PositionDetail)() {
    m.detail = value
}
// Sets the isCurrent property value. Denotes whether or not the position is current.
// Parameters:
//  - value : Value to set for the isCurrent property.
func (m *WorkPosition) SetIsCurrent(value *bool)() {
    m.isCurrent = value
}
// Sets the manager property value. Contains detail of the user's manager in this position.
// Parameters:
//  - value : Value to set for the manager property.
func (m *WorkPosition) SetManager(value *RelatedPerson)() {
    m.manager = value
}

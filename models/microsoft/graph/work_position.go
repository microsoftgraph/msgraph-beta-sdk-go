package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// WorkPosition provides operations to manage the compliance singleton.
type WorkPosition struct {
    ItemFacet
    // Categories that the user has associated with this position.
    categories []string;
    // Colleagues that are associated with this position.
    colleagues []RelatedPersonable;
    // 
    detail PositionDetailable;
    // Denotes whether or not the position is current.
    isCurrent *bool;
    // Contains detail of the user's manager in this position.
    manager RelatedPersonable;
}
// NewWorkPosition instantiates a new workPosition and sets the default values.
func NewWorkPosition()(*WorkPosition) {
    m := &WorkPosition{
        ItemFacet: *NewItemFacet(),
    }
    return m
}
// CreateWorkPositionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWorkPositionFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewWorkPosition(), nil
}
// GetCategories gets the categories property value. Categories that the user has associated with this position.
func (m *WorkPosition) GetCategories()([]string) {
    if m == nil {
        return nil
    } else {
        return m.categories
    }
}
// GetColleagues gets the colleagues property value. Colleagues that are associated with this position.
func (m *WorkPosition) GetColleagues()([]RelatedPersonable) {
    if m == nil {
        return nil
    } else {
        return m.colleagues
    }
}
// GetDetail gets the detail property value. 
func (m *WorkPosition) GetDetail()(PositionDetailable) {
    if m == nil {
        return nil
    } else {
        return m.detail
    }
}
// GetFieldDeserializers the deserialization information for the current model
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
        val, err := n.GetCollectionOfObjectValues(CreateRelatedPersonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RelatedPersonable, len(val))
            for i, v := range val {
                res[i] = v.(RelatedPersonable)
            }
            m.SetColleagues(res)
        }
        return nil
    }
    res["detail"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreatePositionDetailFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetail(val.(PositionDetailable))
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
        val, err := n.GetObjectValue(CreateRelatedPersonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManager(val.(RelatedPersonable))
        }
        return nil
    }
    return res
}
// GetIsCurrent gets the isCurrent property value. Denotes whether or not the position is current.
func (m *WorkPosition) GetIsCurrent()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isCurrent
    }
}
// GetManager gets the manager property value. Contains detail of the user's manager in this position.
func (m *WorkPosition) GetManager()(RelatedPersonable) {
    if m == nil {
        return nil
    } else {
        return m.manager
    }
}
func (m *WorkPosition) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *WorkPosition) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.ItemFacet.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetCategories() != nil {
        err = writer.WriteCollectionOfStringValues("categories", m.GetCategories())
        if err != nil {
            return err
        }
    }
    if m.GetColleagues() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetColleagues()))
        for i, v := range m.GetColleagues() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
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
// SetCategories sets the categories property value. Categories that the user has associated with this position.
func (m *WorkPosition) SetCategories(value []string)() {
    if m != nil {
        m.categories = value
    }
}
// SetColleagues sets the colleagues property value. Colleagues that are associated with this position.
func (m *WorkPosition) SetColleagues(value []RelatedPersonable)() {
    if m != nil {
        m.colleagues = value
    }
}
// SetDetail sets the detail property value. 
func (m *WorkPosition) SetDetail(value PositionDetailable)() {
    if m != nil {
        m.detail = value
    }
}
// SetIsCurrent sets the isCurrent property value. Denotes whether or not the position is current.
func (m *WorkPosition) SetIsCurrent(value *bool)() {
    if m != nil {
        m.isCurrent = value
    }
}
// SetManager sets the manager property value. Contains detail of the user's manager in this position.
func (m *WorkPosition) SetManager(value RelatedPersonable)() {
    if m != nil {
        m.manager = value
    }
}

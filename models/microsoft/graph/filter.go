package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Filter provides operations to call the instantiate method.
type Filter struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // *Experimental* Filter group set used to decide whether given object belongs and should be processed as part of this object mapping. An object is considered in scope if ANY of the groups in the collection is evaluated to true.
    categoryFilterGroups []FilterGroupable;
    // Filter group set used to decide whether given object is in scope for provisioning. This is the filter which should be used in most cases. If an object used to satisfy this filter at a given moment, and then the object or the filter was changed so that filter is not satisfied any longer, such object will get de-provisioned'. An object is considered in scope if ANY of the groups in the collection is evaluated to true.
    groups []FilterGroupable;
    // *Experimental* Filter group set used to filter out objects at the early stage of reading them from the directory. If an object doesn't satisfy this filter it will not be processed further. Important to understand is that if an object used to satisfy this filter at a given moment, and then the object or the filter was changed so that filter is no longer satisfied, such object will NOT get de-provisioned. An object is considered in scope if ANY of the groups in the collection is evaluated to true.
    inputFilterGroups []FilterGroupable;
}
// NewFilter instantiates a new filter and sets the default values.
func NewFilter()(*Filter) {
    m := &Filter{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateFilterFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateFilterFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewFilter(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Filter) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCategoryFilterGroups gets the categoryFilterGroups property value. *Experimental* Filter group set used to decide whether given object belongs and should be processed as part of this object mapping. An object is considered in scope if ANY of the groups in the collection is evaluated to true.
func (m *Filter) GetCategoryFilterGroups()([]FilterGroupable) {
    if m == nil {
        return nil
    } else {
        return m.categoryFilterGroups
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Filter) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["categoryFilterGroups"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateFilterGroupFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]FilterGroupable, len(val))
            for i, v := range val {
                res[i] = v.(FilterGroupable)
            }
            m.SetCategoryFilterGroups(res)
        }
        return nil
    }
    res["groups"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateFilterGroupFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]FilterGroupable, len(val))
            for i, v := range val {
                res[i] = v.(FilterGroupable)
            }
            m.SetGroups(res)
        }
        return nil
    }
    res["inputFilterGroups"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateFilterGroupFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]FilterGroupable, len(val))
            for i, v := range val {
                res[i] = v.(FilterGroupable)
            }
            m.SetInputFilterGroups(res)
        }
        return nil
    }
    return res
}
// GetGroups gets the groups property value. Filter group set used to decide whether given object is in scope for provisioning. This is the filter which should be used in most cases. If an object used to satisfy this filter at a given moment, and then the object or the filter was changed so that filter is not satisfied any longer, such object will get de-provisioned'. An object is considered in scope if ANY of the groups in the collection is evaluated to true.
func (m *Filter) GetGroups()([]FilterGroupable) {
    if m == nil {
        return nil
    } else {
        return m.groups
    }
}
// GetInputFilterGroups gets the inputFilterGroups property value. *Experimental* Filter group set used to filter out objects at the early stage of reading them from the directory. If an object doesn't satisfy this filter it will not be processed further. Important to understand is that if an object used to satisfy this filter at a given moment, and then the object or the filter was changed so that filter is no longer satisfied, such object will NOT get de-provisioned. An object is considered in scope if ANY of the groups in the collection is evaluated to true.
func (m *Filter) GetInputFilterGroups()([]FilterGroupable) {
    if m == nil {
        return nil
    } else {
        return m.inputFilterGroups
    }
}
func (m *Filter) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *Filter) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetCategoryFilterGroups() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetCategoryFilterGroups()))
        for i, v := range m.GetCategoryFilterGroups() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("categoryFilterGroups", cast)
        if err != nil {
            return err
        }
    }
    if m.GetGroups() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetGroups()))
        for i, v := range m.GetGroups() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("groups", cast)
        if err != nil {
            return err
        }
    }
    if m.GetInputFilterGroups() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetInputFilterGroups()))
        for i, v := range m.GetInputFilterGroups() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Filter) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetCategoryFilterGroups sets the categoryFilterGroups property value. *Experimental* Filter group set used to decide whether given object belongs and should be processed as part of this object mapping. An object is considered in scope if ANY of the groups in the collection is evaluated to true.
func (m *Filter) SetCategoryFilterGroups(value []FilterGroupable)() {
    if m != nil {
        m.categoryFilterGroups = value
    }
}
// SetGroups sets the groups property value. Filter group set used to decide whether given object is in scope for provisioning. This is the filter which should be used in most cases. If an object used to satisfy this filter at a given moment, and then the object or the filter was changed so that filter is not satisfied any longer, such object will get de-provisioned'. An object is considered in scope if ANY of the groups in the collection is evaluated to true.
func (m *Filter) SetGroups(value []FilterGroupable)() {
    if m != nil {
        m.groups = value
    }
}
// SetInputFilterGroups sets the inputFilterGroups property value. *Experimental* Filter group set used to filter out objects at the early stage of reading them from the directory. If an object doesn't satisfy this filter it will not be processed further. Important to understand is that if an object used to satisfy this filter at a given moment, and then the object or the filter was changed so that filter is no longer satisfied, such object will NOT get de-provisioned. An object is considered in scope if ANY of the groups in the collection is evaluated to true.
func (m *Filter) SetInputFilterGroups(value []FilterGroupable)() {
    if m != nil {
        m.inputFilterGroups = value
    }
}

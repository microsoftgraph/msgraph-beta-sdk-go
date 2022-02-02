package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ExactDataMatchStoreColumn 
type ExactDataMatchStoreColumn struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    ignoredDelimiters []string;
    // 
    isCaseInsensitive *bool;
    // 
    isSearchable *bool;
    // 
    name *string;
}
// NewExactDataMatchStoreColumn instantiates a new exactDataMatchStoreColumn and sets the default values.
func NewExactDataMatchStoreColumn()(*ExactDataMatchStoreColumn) {
    m := &ExactDataMatchStoreColumn{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ExactDataMatchStoreColumn) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetIgnoredDelimiters gets the ignoredDelimiters property value. 
func (m *ExactDataMatchStoreColumn) GetIgnoredDelimiters()([]string) {
    if m == nil {
        return nil
    } else {
        return m.ignoredDelimiters
    }
}
// GetIsCaseInsensitive gets the isCaseInsensitive property value. 
func (m *ExactDataMatchStoreColumn) GetIsCaseInsensitive()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isCaseInsensitive
    }
}
// GetIsSearchable gets the isSearchable property value. 
func (m *ExactDataMatchStoreColumn) GetIsSearchable()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isSearchable
    }
}
// GetName gets the name property value. 
func (m *ExactDataMatchStoreColumn) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ExactDataMatchStoreColumn) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["ignoredDelimiters"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetIgnoredDelimiters(res)
        }
        return nil
    }
    res["isCaseInsensitive"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsCaseInsensitive(val)
        }
        return nil
    }
    res["isSearchable"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsSearchable(val)
        }
        return nil
    }
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    return res
}
func (m *ExactDataMatchStoreColumn) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ExactDataMatchStoreColumn) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetIgnoredDelimiters() != nil {
        err := writer.WriteCollectionOfStringValues("ignoredDelimiters", m.GetIgnoredDelimiters())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isCaseInsensitive", m.GetIsCaseInsensitive())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isSearchable", m.GetIsSearchable())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
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
func (m *ExactDataMatchStoreColumn) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetIgnoredDelimiters sets the ignoredDelimiters property value. 
func (m *ExactDataMatchStoreColumn) SetIgnoredDelimiters(value []string)() {
    if m != nil {
        m.ignoredDelimiters = value
    }
}
// SetIsCaseInsensitive sets the isCaseInsensitive property value. 
func (m *ExactDataMatchStoreColumn) SetIsCaseInsensitive(value *bool)() {
    if m != nil {
        m.isCaseInsensitive = value
    }
}
// SetIsSearchable sets the isSearchable property value. 
func (m *ExactDataMatchStoreColumn) SetIsSearchable(value *bool)() {
    if m != nil {
        m.isSearchable = value
    }
}
// SetName sets the name property value. 
func (m *ExactDataMatchStoreColumn) SetName(value *string)() {
    if m != nil {
        m.name = value
    }
}

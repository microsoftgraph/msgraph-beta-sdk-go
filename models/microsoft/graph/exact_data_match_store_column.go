package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new exactDataMatchStoreColumn and sets the default values.
func NewExactDataMatchStoreColumn()(*ExactDataMatchStoreColumn) {
    m := &ExactDataMatchStoreColumn{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ExactDataMatchStoreColumn) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the ignoredDelimiters property value. 
func (m *ExactDataMatchStoreColumn) GetIgnoredDelimiters()([]string) {
    if m == nil {
        return nil
    } else {
        return m.ignoredDelimiters
    }
}
// Gets the isCaseInsensitive property value. 
func (m *ExactDataMatchStoreColumn) GetIsCaseInsensitive()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isCaseInsensitive
    }
}
// Gets the isSearchable property value. 
func (m *ExactDataMatchStoreColumn) GetIsSearchable()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isSearchable
    }
}
// Gets the name property value. 
func (m *ExactDataMatchStoreColumn) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ExactDataMatchStoreColumn) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *ExactDataMatchStoreColumn) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the ignoredDelimiters property value. 
// Parameters:
//  - value : Value to set for the ignoredDelimiters property.
func (m *ExactDataMatchStoreColumn) SetIgnoredDelimiters(value []string)() {
    m.ignoredDelimiters = value
}
// Sets the isCaseInsensitive property value. 
// Parameters:
//  - value : Value to set for the isCaseInsensitive property.
func (m *ExactDataMatchStoreColumn) SetIsCaseInsensitive(value *bool)() {
    m.isCaseInsensitive = value
}
// Sets the isSearchable property value. 
// Parameters:
//  - value : Value to set for the isSearchable property.
func (m *ExactDataMatchStoreColumn) SetIsSearchable(value *bool)() {
    m.isSearchable = value
}
// Sets the name property value. 
// Parameters:
//  - value : Value to set for the name property.
func (m *ExactDataMatchStoreColumn) SetName(value *string)() {
    m.name = value
}

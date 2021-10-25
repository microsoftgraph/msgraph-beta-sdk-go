package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ExactDataMatchStoreColumn struct {
    additionalData map[string]interface{};
    ignoredDelimiters []string;
    isCaseInsensitive *bool;
    isSearchable *bool;
    name *string;
}
func NewExactDataMatchStoreColumn()(*ExactDataMatchStoreColumn) {
    m := &ExactDataMatchStoreColumn{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *ExactDataMatchStoreColumn) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *ExactDataMatchStoreColumn) GetIgnoredDelimiters()([]string) {
    if m == nil {
        return nil
    } else {
        return m.ignoredDelimiters
    }
}
func (m *ExactDataMatchStoreColumn) GetIsCaseInsensitive()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isCaseInsensitive
    }
}
func (m *ExactDataMatchStoreColumn) GetIsSearchable()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isSearchable
    }
}
func (m *ExactDataMatchStoreColumn) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
func (m *ExactDataMatchStoreColumn) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["ignoredDelimiters"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetIgnoredDelimiters(res)
        return nil
    }
    res["isCaseInsensitive"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsCaseInsensitive(val)
        return nil
    }
    res["isSearchable"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsSearchable(val)
        return nil
    }
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetName(val)
        return nil
    }
    return res
}
func (m *ExactDataMatchStoreColumn) IsNil()(bool) {
    return m == nil
}
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
func (m *ExactDataMatchStoreColumn) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *ExactDataMatchStoreColumn) SetIgnoredDelimiters(value []string)() {
    m.ignoredDelimiters = value
}
func (m *ExactDataMatchStoreColumn) SetIsCaseInsensitive(value *bool)() {
    m.isCaseInsensitive = value
}
func (m *ExactDataMatchStoreColumn) SetIsSearchable(value *bool)() {
    m.isSearchable = value
}
func (m *ExactDataMatchStoreColumn) SetName(value *string)() {
    m.name = value
}

package lookup

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type LookupRequestBody struct {
    additionalData map[string]interface{};
    key *string;
    resultColumnNames []string;
    values []string;
}
func NewLookupRequestBody()(*LookupRequestBody) {
    m := &LookupRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *LookupRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *LookupRequestBody) GetKey()(*string) {
    if m == nil {
        return nil
    } else {
        return m.key
    }
}
func (m *LookupRequestBody) GetResultColumnNames()([]string) {
    if m == nil {
        return nil
    } else {
        return m.resultColumnNames
    }
}
func (m *LookupRequestBody) GetValues()([]string) {
    if m == nil {
        return nil
    } else {
        return m.values
    }
}
func (m *LookupRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["key"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetKey(val)
        return nil
    }
    res["resultColumnNames"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetResultColumnNames(res)
        return nil
    }
    res["values"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetValues(res)
        return nil
    }
    return res
}
func (m *LookupRequestBody) IsNil()(bool) {
    return m == nil
}
func (m *LookupRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("key", m.GetKey())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("resultColumnNames", m.GetResultColumnNames())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("values", m.GetValues())
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
func (m *LookupRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *LookupRequestBody) SetKey(value *string)() {
    m.key = value
}
func (m *LookupRequestBody) SetResultColumnNames(value []string)() {
    m.resultColumnNames = value
}
func (m *LookupRequestBody) SetValues(value []string)() {
    m.values = value
}

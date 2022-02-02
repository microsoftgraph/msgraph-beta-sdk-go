package lookup

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// LookupRequestBody 
type LookupRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    key *string;
    // 
    resultColumnNames []string;
    // 
    values []string;
}
// NewLookupRequestBody instantiates a new lookupRequestBody and sets the default values.
func NewLookupRequestBody()(*LookupRequestBody) {
    m := &LookupRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *LookupRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetKey gets the key property value. 
func (m *LookupRequestBody) GetKey()(*string) {
    if m == nil {
        return nil
    } else {
        return m.key
    }
}
// GetResultColumnNames gets the resultColumnNames property value. 
func (m *LookupRequestBody) GetResultColumnNames()([]string) {
    if m == nil {
        return nil
    } else {
        return m.resultColumnNames
    }
}
// GetValues gets the values property value. 
func (m *LookupRequestBody) GetValues()([]string) {
    if m == nil {
        return nil
    } else {
        return m.values
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *LookupRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["key"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKey(val)
        }
        return nil
    }
    res["resultColumnNames"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetResultColumnNames(res)
        }
        return nil
    }
    res["values"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetValues(res)
        }
        return nil
    }
    return res
}
func (m *LookupRequestBody) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *LookupRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("key", m.GetKey())
        if err != nil {
            return err
        }
    }
    if m.GetResultColumnNames() != nil {
        err := writer.WriteCollectionOfStringValues("resultColumnNames", m.GetResultColumnNames())
        if err != nil {
            return err
        }
    }
    if m.GetValues() != nil {
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *LookupRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetKey sets the key property value. 
func (m *LookupRequestBody) SetKey(value *string)() {
    if m != nil {
        m.key = value
    }
}
// SetResultColumnNames sets the resultColumnNames property value. 
func (m *LookupRequestBody) SetResultColumnNames(value []string)() {
    if m != nil {
        m.resultColumnNames = value
    }
}
// SetValues sets the values property value. 
func (m *LookupRequestBody) SetValues(value []string)() {
    if m != nil {
        m.values = value
    }
}

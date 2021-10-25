package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type SearchBucket struct {
    additionalData map[string]interface{};
    aggregationFilterToken *string;
    count *int32;
    key *string;
}
func NewSearchBucket()(*SearchBucket) {
    m := &SearchBucket{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *SearchBucket) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *SearchBucket) GetAggregationFilterToken()(*string) {
    if m == nil {
        return nil
    } else {
        return m.aggregationFilterToken
    }
}
func (m *SearchBucket) GetCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.count
    }
}
func (m *SearchBucket) GetKey()(*string) {
    if m == nil {
        return nil
    } else {
        return m.key
    }
}
func (m *SearchBucket) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["aggregationFilterToken"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAggregationFilterToken(val)
        return nil
    }
    res["count"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetCount(val)
        return nil
    }
    res["key"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetKey(val)
        return nil
    }
    return res
}
func (m *SearchBucket) IsNil()(bool) {
    return m == nil
}
func (m *SearchBucket) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("aggregationFilterToken", m.GetAggregationFilterToken())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("count", m.GetCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("key", m.GetKey())
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
func (m *SearchBucket) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *SearchBucket) SetAggregationFilterToken(value *string)() {
    m.aggregationFilterToken = value
}
func (m *SearchBucket) SetCount(value *int32)() {
    m.count = value
}
func (m *SearchBucket) SetKey(value *string)() {
    m.key = value
}

package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type AggregationOption struct {
    additionalData map[string]interface{};
    bucketDefinition *BucketAggregationDefinition;
    field *string;
    size *int32;
}
func NewAggregationOption()(*AggregationOption) {
    m := &AggregationOption{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *AggregationOption) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *AggregationOption) GetBucketDefinition()(*BucketAggregationDefinition) {
    if m == nil {
        return nil
    } else {
        return m.bucketDefinition
    }
}
func (m *AggregationOption) GetField()(*string) {
    if m == nil {
        return nil
    } else {
        return m.field
    }
}
func (m *AggregationOption) GetSize()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.size
    }
}
func (m *AggregationOption) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["bucketDefinition"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewBucketAggregationDefinition() })
        if err != nil {
            return err
        }
        m.SetBucketDefinition(val.(*BucketAggregationDefinition))
        return nil
    }
    res["field"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetField(val)
        return nil
    }
    res["size"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetSize(val)
        return nil
    }
    return res
}
func (m *AggregationOption) IsNil()(bool) {
    return m == nil
}
func (m *AggregationOption) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("bucketDefinition", m.GetBucketDefinition())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("field", m.GetField())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("size", m.GetSize())
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
func (m *AggregationOption) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *AggregationOption) SetBucketDefinition(value *BucketAggregationDefinition)() {
    m.bucketDefinition = value
}
func (m *AggregationOption) SetField(value *string)() {
    m.field = value
}
func (m *AggregationOption) SetSize(value *int32)() {
    m.size = value
}

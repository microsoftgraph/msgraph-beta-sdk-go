package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type BucketAggregationRange struct {
    additionalData map[string]interface{};
    from *string;
    to *string;
}
func NewBucketAggregationRange()(*BucketAggregationRange) {
    m := &BucketAggregationRange{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *BucketAggregationRange) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *BucketAggregationRange) GetFrom()(*string) {
    if m == nil {
        return nil
    } else {
        return m.from
    }
}
func (m *BucketAggregationRange) GetTo()(*string) {
    if m == nil {
        return nil
    } else {
        return m.to
    }
}
func (m *BucketAggregationRange) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["from"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetFrom(val)
        return nil
    }
    res["to"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTo(val)
        return nil
    }
    return res
}
func (m *BucketAggregationRange) IsNil()(bool) {
    return m == nil
}
func (m *BucketAggregationRange) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("from", m.GetFrom())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("to", m.GetTo())
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
func (m *BucketAggregationRange) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *BucketAggregationRange) SetFrom(value *string)() {
    m.from = value
}
func (m *BucketAggregationRange) SetTo(value *string)() {
    m.to = value
}

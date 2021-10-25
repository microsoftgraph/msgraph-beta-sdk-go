package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type BucketAggregationDefinition struct {
    additionalData map[string]interface{};
    isDescending *bool;
    minimumCount *int32;
    prefixFilter *string;
    ranges []BucketAggregationRange;
    sortBy *BucketAggregationSortProperty;
}
func NewBucketAggregationDefinition()(*BucketAggregationDefinition) {
    m := &BucketAggregationDefinition{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *BucketAggregationDefinition) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *BucketAggregationDefinition) GetIsDescending()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDescending
    }
}
func (m *BucketAggregationDefinition) GetMinimumCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.minimumCount
    }
}
func (m *BucketAggregationDefinition) GetPrefixFilter()(*string) {
    if m == nil {
        return nil
    } else {
        return m.prefixFilter
    }
}
func (m *BucketAggregationDefinition) GetRanges()([]BucketAggregationRange) {
    if m == nil {
        return nil
    } else {
        return m.ranges
    }
}
func (m *BucketAggregationDefinition) GetSortBy()(*BucketAggregationSortProperty) {
    if m == nil {
        return nil
    } else {
        return m.sortBy
    }
}
func (m *BucketAggregationDefinition) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["isDescending"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsDescending(val)
        return nil
    }
    res["minimumCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetMinimumCount(val)
        return nil
    }
    res["prefixFilter"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPrefixFilter(val)
        return nil
    }
    res["ranges"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewBucketAggregationRange() })
        if err != nil {
            return err
        }
        res := make([]BucketAggregationRange, len(val))
        for i, v := range val {
            res[i] = *(v.(*BucketAggregationRange))
        }
        m.SetRanges(res)
        return nil
    }
    res["sortBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseBucketAggregationSortProperty)
        if err != nil {
            return err
        }
        cast := val.(BucketAggregationSortProperty)
        m.SetSortBy(&cast)
        return nil
    }
    return res
}
func (m *BucketAggregationDefinition) IsNil()(bool) {
    return m == nil
}
func (m *BucketAggregationDefinition) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("isDescending", m.GetIsDescending())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("minimumCount", m.GetMinimumCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("prefixFilter", m.GetPrefixFilter())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetRanges()))
        for i, v := range m.GetRanges() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("ranges", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSortBy() != nil {
        cast := m.GetSortBy().String()
        err := writer.WriteStringValue("sortBy", &cast)
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
func (m *BucketAggregationDefinition) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *BucketAggregationDefinition) SetIsDescending(value *bool)() {
    m.isDescending = value
}
func (m *BucketAggregationDefinition) SetMinimumCount(value *int32)() {
    m.minimumCount = value
}
func (m *BucketAggregationDefinition) SetPrefixFilter(value *string)() {
    m.prefixFilter = value
}
func (m *BucketAggregationDefinition) SetRanges(value []BucketAggregationRange)() {
    m.ranges = value
}
func (m *BucketAggregationDefinition) SetSortBy(value *BucketAggregationSortProperty)() {
    m.sortBy = value
}

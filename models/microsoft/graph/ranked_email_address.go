package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type RankedEmailAddress struct {
    additionalData map[string]interface{};
    address *string;
    rank *float64;
}
func NewRankedEmailAddress()(*RankedEmailAddress) {
    m := &RankedEmailAddress{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *RankedEmailAddress) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *RankedEmailAddress) GetAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.address
    }
}
func (m *RankedEmailAddress) GetRank()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.rank
    }
}
func (m *RankedEmailAddress) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["address"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAddress(val)
        return nil
    }
    res["rank"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetRank(val)
        return nil
    }
    return res
}
func (m *RankedEmailAddress) IsNil()(bool) {
    return m == nil
}
func (m *RankedEmailAddress) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("address", m.GetAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("rank", m.GetRank())
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
func (m *RankedEmailAddress) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *RankedEmailAddress) SetAddress(value *string)() {
    m.address = value
}
func (m *RankedEmailAddress) SetRank(value *float64)() {
    m.rank = value
}

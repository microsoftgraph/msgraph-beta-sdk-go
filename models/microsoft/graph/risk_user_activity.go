package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type RiskUserActivity struct {
    additionalData map[string]interface{};
    detail *RiskDetail;
    eventTypes []RiskEventType;
    riskEventTypes []string;
}
func NewRiskUserActivity()(*RiskUserActivity) {
    m := &RiskUserActivity{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *RiskUserActivity) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *RiskUserActivity) GetDetail()(*RiskDetail) {
    if m == nil {
        return nil
    } else {
        return m.detail
    }
}
func (m *RiskUserActivity) GetEventTypes()([]RiskEventType) {
    if m == nil {
        return nil
    } else {
        return m.eventTypes
    }
}
func (m *RiskUserActivity) GetRiskEventTypes()([]string) {
    if m == nil {
        return nil
    } else {
        return m.riskEventTypes
    }
}
func (m *RiskUserActivity) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["detail"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseRiskDetail)
        if err != nil {
            return err
        }
        cast := val.(RiskDetail)
        m.SetDetail(&cast)
        return nil
    }
    res["eventTypes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseRiskEventType)
        if err != nil {
            return err
        }
        res := make([]RiskEventType, len(val))
        for i, v := range val {
            res[i] = *(v.(*RiskEventType))
        }
        m.SetEventTypes(res)
        return nil
    }
    res["riskEventTypes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetRiskEventTypes(res)
        return nil
    }
    return res
}
func (m *RiskUserActivity) IsNil()(bool) {
    return m == nil
}
func (m *RiskUserActivity) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetDetail() != nil {
        cast := m.GetDetail().String()
        err := writer.WriteStringValue("detail", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("eventTypes", SerializeRiskEventType(m.GetEventTypes()))
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("riskEventTypes", m.GetRiskEventTypes())
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
func (m *RiskUserActivity) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *RiskUserActivity) SetDetail(value *RiskDetail)() {
    m.detail = value
}
func (m *RiskUserActivity) SetEventTypes(value []RiskEventType)() {
    m.eventTypes = value
}
func (m *RiskUserActivity) SetRiskEventTypes(value []string)() {
    m.riskEventTypes = value
}

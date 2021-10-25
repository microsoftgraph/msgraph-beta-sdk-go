package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type AlterationResponse struct {
    additionalData map[string]interface{};
    originalQueryString *string;
    queryAlteration *SearchAlteration;
    queryAlterationType *SearchAlterationType;
}
func NewAlterationResponse()(*AlterationResponse) {
    m := &AlterationResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *AlterationResponse) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *AlterationResponse) GetOriginalQueryString()(*string) {
    if m == nil {
        return nil
    } else {
        return m.originalQueryString
    }
}
func (m *AlterationResponse) GetQueryAlteration()(*SearchAlteration) {
    if m == nil {
        return nil
    } else {
        return m.queryAlteration
    }
}
func (m *AlterationResponse) GetQueryAlterationType()(*SearchAlterationType) {
    if m == nil {
        return nil
    } else {
        return m.queryAlterationType
    }
}
func (m *AlterationResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["originalQueryString"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOriginalQueryString(val)
        return nil
    }
    res["queryAlteration"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSearchAlteration() })
        if err != nil {
            return err
        }
        m.SetQueryAlteration(val.(*SearchAlteration))
        return nil
    }
    res["queryAlterationType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseSearchAlterationType)
        if err != nil {
            return err
        }
        cast := val.(SearchAlterationType)
        m.SetQueryAlterationType(&cast)
        return nil
    }
    return res
}
func (m *AlterationResponse) IsNil()(bool) {
    return m == nil
}
func (m *AlterationResponse) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("originalQueryString", m.GetOriginalQueryString())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("queryAlteration", m.GetQueryAlteration())
        if err != nil {
            return err
        }
    }
    if m.GetQueryAlterationType() != nil {
        cast := m.GetQueryAlterationType().String()
        err := writer.WriteStringValue("queryAlterationType", &cast)
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
func (m *AlterationResponse) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *AlterationResponse) SetOriginalQueryString(value *string)() {
    m.originalQueryString = value
}
func (m *AlterationResponse) SetQueryAlteration(value *SearchAlteration)() {
    m.queryAlteration = value
}
func (m *AlterationResponse) SetQueryAlterationType(value *SearchAlterationType)() {
    m.queryAlterationType = value
}

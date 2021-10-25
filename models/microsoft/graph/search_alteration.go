package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type SearchAlteration struct {
    additionalData map[string]interface{};
    alteredHighlightedQueryString *string;
    alteredQueryString *string;
    alteredQueryTokens []AlteredQueryToken;
}
func NewSearchAlteration()(*SearchAlteration) {
    m := &SearchAlteration{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *SearchAlteration) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *SearchAlteration) GetAlteredHighlightedQueryString()(*string) {
    if m == nil {
        return nil
    } else {
        return m.alteredHighlightedQueryString
    }
}
func (m *SearchAlteration) GetAlteredQueryString()(*string) {
    if m == nil {
        return nil
    } else {
        return m.alteredQueryString
    }
}
func (m *SearchAlteration) GetAlteredQueryTokens()([]AlteredQueryToken) {
    if m == nil {
        return nil
    } else {
        return m.alteredQueryTokens
    }
}
func (m *SearchAlteration) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["alteredHighlightedQueryString"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAlteredHighlightedQueryString(val)
        return nil
    }
    res["alteredQueryString"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAlteredQueryString(val)
        return nil
    }
    res["alteredQueryTokens"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAlteredQueryToken() })
        if err != nil {
            return err
        }
        res := make([]AlteredQueryToken, len(val))
        for i, v := range val {
            res[i] = *(v.(*AlteredQueryToken))
        }
        m.SetAlteredQueryTokens(res)
        return nil
    }
    return res
}
func (m *SearchAlteration) IsNil()(bool) {
    return m == nil
}
func (m *SearchAlteration) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("alteredHighlightedQueryString", m.GetAlteredHighlightedQueryString())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("alteredQueryString", m.GetAlteredQueryString())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAlteredQueryTokens()))
        for i, v := range m.GetAlteredQueryTokens() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("alteredQueryTokens", cast)
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
func (m *SearchAlteration) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *SearchAlteration) SetAlteredHighlightedQueryString(value *string)() {
    m.alteredHighlightedQueryString = value
}
func (m *SearchAlteration) SetAlteredQueryString(value *string)() {
    m.alteredQueryString = value
}
func (m *SearchAlteration) SetAlteredQueryTokens(value []AlteredQueryToken)() {
    m.alteredQueryTokens = value
}

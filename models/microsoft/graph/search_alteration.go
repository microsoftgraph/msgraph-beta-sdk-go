package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type SearchAlteration struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Defines the altered highlighted query string with spelling correction. The annotation around the corrected segment is (/ue000, /ue001)
    alteredHighlightedQueryString *string;
    // Defines the altered query string with spelling correction.
    alteredQueryString *string;
    // Represents changed segments with respect to original query.
    alteredQueryTokens []AlteredQueryToken;
}
// Instantiates a new searchAlteration and sets the default values.
func NewSearchAlteration()(*SearchAlteration) {
    m := &SearchAlteration{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SearchAlteration) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the alteredHighlightedQueryString property value. Defines the altered highlighted query string with spelling correction. The annotation around the corrected segment is (/ue000, /ue001)
func (m *SearchAlteration) GetAlteredHighlightedQueryString()(*string) {
    if m == nil {
        return nil
    } else {
        return m.alteredHighlightedQueryString
    }
}
// Gets the alteredQueryString property value. Defines the altered query string with spelling correction.
func (m *SearchAlteration) GetAlteredQueryString()(*string) {
    if m == nil {
        return nil
    } else {
        return m.alteredQueryString
    }
}
// Gets the alteredQueryTokens property value. Represents changed segments with respect to original query.
func (m *SearchAlteration) GetAlteredQueryTokens()([]AlteredQueryToken) {
    if m == nil {
        return nil
    } else {
        return m.alteredQueryTokens
    }
}
// The deserialization information for the current model
func (m *SearchAlteration) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["alteredHighlightedQueryString"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAlteredHighlightedQueryString(val)
        }
        return nil
    }
    res["alteredQueryString"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAlteredQueryString(val)
        }
        return nil
    }
    res["alteredQueryTokens"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAlteredQueryToken() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AlteredQueryToken, len(val))
            for i, v := range val {
                res[i] = *(v.(*AlteredQueryToken))
            }
            m.SetAlteredQueryTokens(res)
        }
        return nil
    }
    return res
}
func (m *SearchAlteration) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *SearchAlteration) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the alteredHighlightedQueryString property value. Defines the altered highlighted query string with spelling correction. The annotation around the corrected segment is (/ue000, /ue001)
// Parameters:
//  - value : Value to set for the alteredHighlightedQueryString property.
func (m *SearchAlteration) SetAlteredHighlightedQueryString(value *string)() {
    m.alteredHighlightedQueryString = value
}
// Sets the alteredQueryString property value. Defines the altered query string with spelling correction.
// Parameters:
//  - value : Value to set for the alteredQueryString property.
func (m *SearchAlteration) SetAlteredQueryString(value *string)() {
    m.alteredQueryString = value
}
// Sets the alteredQueryTokens property value. Represents changed segments with respect to original query.
// Parameters:
//  - value : Value to set for the alteredQueryTokens property.
func (m *SearchAlteration) SetAlteredQueryTokens(value []AlteredQueryToken)() {
    m.alteredQueryTokens = value
}

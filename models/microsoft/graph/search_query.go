package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type SearchQuery struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    query_string *SearchQueryString;
    // The search query containing the search terms. Required.
    queryString *string;
}
// Instantiates a new searchQuery and sets the default values.
func NewSearchQuery()(*SearchQuery) {
    m := &SearchQuery{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SearchQuery) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the query_string property value. 
func (m *SearchQuery) GetQuery_string()(*SearchQueryString) {
    if m == nil {
        return nil
    } else {
        return m.query_string
    }
}
// Gets the queryString property value. The search query containing the search terms. Required.
func (m *SearchQuery) GetQueryString()(*string) {
    if m == nil {
        return nil
    } else {
        return m.queryString
    }
}
// The deserialization information for the current model
func (m *SearchQuery) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["query_string"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSearchQueryString() })
        if err != nil {
            return err
        }
        m.SetQuery_string(val.(*SearchQueryString))
        return nil
    }
    res["queryString"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetQueryString(val)
        return nil
    }
    return res
}
func (m *SearchQuery) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *SearchQuery) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("query_string", m.GetQuery_string())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("queryString", m.GetQueryString())
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
func (m *SearchQuery) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the query_string property value. 
// Parameters:
//  - value : Value to set for the query_string property.
func (m *SearchQuery) SetQuery_string(value *SearchQueryString)() {
    m.query_string = value
}
// Sets the queryString property value. The search query containing the search terms. Required.
// Parameters:
//  - value : Value to set for the queryString property.
func (m *SearchQuery) SetQueryString(value *string)() {
    m.queryString = value
}

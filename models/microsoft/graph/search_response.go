package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type SearchResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Provides details of query alteration response for spelling correction.
    queryAlterationResponse *AlterationResponse;
    // Represents results from a search query, and the terms used for the query.
    value []SearchResultSet;
}
// Instantiates a new SearchResponse and sets the default values.
func NewSearchResponse()(*SearchResponse) {
    m := &SearchResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SearchResponse) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the queryAlterationResponse property value. Provides details of query alteration response for spelling correction.
func (m *SearchResponse) GetQueryAlterationResponse()(*AlterationResponse) {
    if m == nil {
        return nil
    } else {
        return m.queryAlterationResponse
    }
}
// Gets the value property value. Represents results from a search query, and the terms used for the query.
func (m *SearchResponse) GetValue()([]SearchResultSet) {
    if m == nil {
        return nil
    } else {
        return m.value
    }
}
// The deserialization information for the current model
func (m *SearchResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["queryAlterationResponse"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAlterationResponse() })
        if err != nil {
            return err
        }
        m.SetQueryAlterationResponse(val.(*AlterationResponse))
        return nil
    }
    res["value"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSearchResultSet() })
        if err != nil {
            return err
        }
        res := make([]SearchResultSet, len(val))
        for i, v := range val {
            res[i] = *(v.(*SearchResultSet))
        }
        m.SetValue(res)
        return nil
    }
    return res
}
func (m *SearchResponse) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *SearchResponse) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("queryAlterationResponse", m.GetQueryAlterationResponse())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetValue()))
        for i, v := range m.GetValue() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("value", cast)
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
func (m *SearchResponse) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the queryAlterationResponse property value. Provides details of query alteration response for spelling correction.
// Parameters:
//  - value : Value to set for the queryAlterationResponse property.
func (m *SearchResponse) SetQueryAlterationResponse(value *AlterationResponse)() {
    m.queryAlterationResponse = value
}
// Sets the value property value. Represents results from a search query, and the terms used for the query.
// Parameters:
//  - value : Value to set for the value property.
func (m *SearchResponse) SetValue(value []SearchResultSet)() {
    m.value = value
}

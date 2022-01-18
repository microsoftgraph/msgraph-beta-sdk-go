package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SearchQuery 
type SearchQuery struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    query_string *SearchQueryString;
    // The search query containing the search terms. Required.
    queryString *string;
    // 
    queryTemplate *string;
}
// NewSearchQuery instantiates a new searchQuery and sets the default values.
func NewSearchQuery()(*SearchQuery) {
    m := &SearchQuery{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SearchQuery) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetQuery_string gets the query_string property value. 
func (m *SearchQuery) GetQuery_string()(*SearchQueryString) {
    if m == nil {
        return nil
    } else {
        return m.query_string
    }
}
// GetQueryString gets the queryString property value. The search query containing the search terms. Required.
func (m *SearchQuery) GetQueryString()(*string) {
    if m == nil {
        return nil
    } else {
        return m.queryString
    }
}
// GetQueryTemplate gets the queryTemplate property value. 
func (m *SearchQuery) GetQueryTemplate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.queryTemplate
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SearchQuery) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["query_string"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSearchQueryString() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQuery_string(val.(*SearchQueryString))
        }
        return nil
    }
    res["queryString"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQueryString(val)
        }
        return nil
    }
    res["queryTemplate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQueryTemplate(val)
        }
        return nil
    }
    return res
}
func (m *SearchQuery) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
        err := writer.WriteStringValue("queryTemplate", m.GetQueryTemplate())
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SearchQuery) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetQuery_string sets the query_string property value. 
func (m *SearchQuery) SetQuery_string(value *SearchQueryString)() {
    if m != nil {
        m.query_string = value
    }
}
// SetQueryString sets the queryString property value. The search query containing the search terms. Required.
func (m *SearchQuery) SetQueryString(value *string)() {
    if m != nil {
        m.queryString = value
    }
}
// SetQueryTemplate sets the queryTemplate property value. 
func (m *SearchQuery) SetQueryTemplate(value *string)() {
    if m != nil {
        m.queryTemplate = value
    }
}

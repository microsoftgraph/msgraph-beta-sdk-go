package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type SearchResultSet struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // A collection of search results.
    hitsContainers []SearchHitsContainer;
    // A dictionary of resultTemplateIds and associated values, which include the name and JSON schema of the result templates.
    resultTemplates *ResultTemplateDictionary;
    // Contains the search terms sent in the initial search query.
    searchTerms []string;
}
// Instantiates a new searchResultSet and sets the default values.
func NewSearchResultSet()(*SearchResultSet) {
    m := &SearchResultSet{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SearchResultSet) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the hitsContainers property value. A collection of search results.
func (m *SearchResultSet) GetHitsContainers()([]SearchHitsContainer) {
    if m == nil {
        return nil
    } else {
        return m.hitsContainers
    }
}
// Gets the resultTemplates property value. A dictionary of resultTemplateIds and associated values, which include the name and JSON schema of the result templates.
func (m *SearchResultSet) GetResultTemplates()(*ResultTemplateDictionary) {
    if m == nil {
        return nil
    } else {
        return m.resultTemplates
    }
}
// Gets the searchTerms property value. Contains the search terms sent in the initial search query.
func (m *SearchResultSet) GetSearchTerms()([]string) {
    if m == nil {
        return nil
    } else {
        return m.searchTerms
    }
}
// The deserialization information for the current model
func (m *SearchResultSet) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["hitsContainers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSearchHitsContainer() })
        if err != nil {
            return err
        }
        res := make([]SearchHitsContainer, len(val))
        for i, v := range val {
            res[i] = *(v.(*SearchHitsContainer))
        }
        m.SetHitsContainers(res)
        return nil
    }
    res["resultTemplates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewResultTemplateDictionary() })
        if err != nil {
            return err
        }
        m.SetResultTemplates(val.(*ResultTemplateDictionary))
        return nil
    }
    res["searchTerms"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetSearchTerms(res)
        return nil
    }
    return res
}
func (m *SearchResultSet) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *SearchResultSet) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetHitsContainers()))
        for i, v := range m.GetHitsContainers() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("hitsContainers", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("resultTemplates", m.GetResultTemplates())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("searchTerms", m.GetSearchTerms())
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
func (m *SearchResultSet) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the hitsContainers property value. A collection of search results.
// Parameters:
//  - value : Value to set for the hitsContainers property.
func (m *SearchResultSet) SetHitsContainers(value []SearchHitsContainer)() {
    m.hitsContainers = value
}
// Sets the resultTemplates property value. A dictionary of resultTemplateIds and associated values, which include the name and JSON schema of the result templates.
// Parameters:
//  - value : Value to set for the resultTemplates property.
func (m *SearchResultSet) SetResultTemplates(value *ResultTemplateDictionary)() {
    m.resultTemplates = value
}
// Sets the searchTerms property value. Contains the search terms sent in the initial search query.
// Parameters:
//  - value : Value to set for the searchTerms property.
func (m *SearchResultSet) SetSearchTerms(value []string)() {
    m.searchTerms = value
}

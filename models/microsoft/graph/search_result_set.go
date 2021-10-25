package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type SearchResultSet struct {
    additionalData map[string]interface{};
    hitsContainers []SearchHitsContainer;
    resultTemplates *ResultTemplateDictionary;
    searchTerms []string;
}
func NewSearchResultSet()(*SearchResultSet) {
    m := &SearchResultSet{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *SearchResultSet) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *SearchResultSet) GetHitsContainers()([]SearchHitsContainer) {
    if m == nil {
        return nil
    } else {
        return m.hitsContainers
    }
}
func (m *SearchResultSet) GetResultTemplates()(*ResultTemplateDictionary) {
    if m == nil {
        return nil
    } else {
        return m.resultTemplates
    }
}
func (m *SearchResultSet) GetSearchTerms()([]string) {
    if m == nil {
        return nil
    } else {
        return m.searchTerms
    }
}
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
func (m *SearchResultSet) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *SearchResultSet) SetHitsContainers(value []SearchHitsContainer)() {
    m.hitsContainers = value
}
func (m *SearchResultSet) SetResultTemplates(value *ResultTemplateDictionary)() {
    m.resultTemplates = value
}
func (m *SearchResultSet) SetSearchTerms(value []string)() {
    m.searchTerms = value
}

package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type SearchRequest struct {
    additionalData map[string]interface{};
    aggregationFilters []string;
    aggregations []AggregationOption;
    contentSources []string;
    enableTopResults *bool;
    entityTypes []EntityType;
    fields []string;
    from *int32;
    query *SearchQuery;
    resultTemplateOptions *ResultTemplateOption;
    size *int32;
    sortProperties []SortProperty;
    stored_fields []string;
}
func NewSearchRequest()(*SearchRequest) {
    m := &SearchRequest{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *SearchRequest) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *SearchRequest) GetAggregationFilters()([]string) {
    if m == nil {
        return nil
    } else {
        return m.aggregationFilters
    }
}
func (m *SearchRequest) GetAggregations()([]AggregationOption) {
    if m == nil {
        return nil
    } else {
        return m.aggregations
    }
}
func (m *SearchRequest) GetContentSources()([]string) {
    if m == nil {
        return nil
    } else {
        return m.contentSources
    }
}
func (m *SearchRequest) GetEnableTopResults()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.enableTopResults
    }
}
func (m *SearchRequest) GetEntityTypes()([]EntityType) {
    if m == nil {
        return nil
    } else {
        return m.entityTypes
    }
}
func (m *SearchRequest) GetFields()([]string) {
    if m == nil {
        return nil
    } else {
        return m.fields
    }
}
func (m *SearchRequest) GetFrom()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.from
    }
}
func (m *SearchRequest) GetQuery()(*SearchQuery) {
    if m == nil {
        return nil
    } else {
        return m.query
    }
}
func (m *SearchRequest) GetResultTemplateOptions()(*ResultTemplateOption) {
    if m == nil {
        return nil
    } else {
        return m.resultTemplateOptions
    }
}
func (m *SearchRequest) GetSize()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.size
    }
}
func (m *SearchRequest) GetSortProperties()([]SortProperty) {
    if m == nil {
        return nil
    } else {
        return m.sortProperties
    }
}
func (m *SearchRequest) GetStored_fields()([]string) {
    if m == nil {
        return nil
    } else {
        return m.stored_fields
    }
}
func (m *SearchRequest) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["aggregationFilters"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetAggregationFilters(res)
        return nil
    }
    res["aggregations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAggregationOption() })
        if err != nil {
            return err
        }
        res := make([]AggregationOption, len(val))
        for i, v := range val {
            res[i] = *(v.(*AggregationOption))
        }
        m.SetAggregations(res)
        return nil
    }
    res["contentSources"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetContentSources(res)
        return nil
    }
    res["enableTopResults"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetEnableTopResults(val)
        return nil
    }
    res["entityTypes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseEntityType)
        if err != nil {
            return err
        }
        res := make([]EntityType, len(val))
        for i, v := range val {
            res[i] = *(v.(*EntityType))
        }
        m.SetEntityTypes(res)
        return nil
    }
    res["fields"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetFields(res)
        return nil
    }
    res["from"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetFrom(val)
        return nil
    }
    res["query"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSearchQuery() })
        if err != nil {
            return err
        }
        m.SetQuery(val.(*SearchQuery))
        return nil
    }
    res["resultTemplateOptions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewResultTemplateOption() })
        if err != nil {
            return err
        }
        m.SetResultTemplateOptions(val.(*ResultTemplateOption))
        return nil
    }
    res["size"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetSize(val)
        return nil
    }
    res["sortProperties"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSortProperty() })
        if err != nil {
            return err
        }
        res := make([]SortProperty, len(val))
        for i, v := range val {
            res[i] = *(v.(*SortProperty))
        }
        m.SetSortProperties(res)
        return nil
    }
    res["stored_fields"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetStored_fields(res)
        return nil
    }
    return res
}
func (m *SearchRequest) IsNil()(bool) {
    return m == nil
}
func (m *SearchRequest) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteCollectionOfStringValues("aggregationFilters", m.GetAggregationFilters())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAggregations()))
        for i, v := range m.GetAggregations() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("aggregations", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("contentSources", m.GetContentSources())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("enableTopResults", m.GetEnableTopResults())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("entityTypes", SerializeEntityType(m.GetEntityTypes()))
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("fields", m.GetFields())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("from", m.GetFrom())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("query", m.GetQuery())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("resultTemplateOptions", m.GetResultTemplateOptions())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("size", m.GetSize())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSortProperties()))
        for i, v := range m.GetSortProperties() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("sortProperties", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("stored_fields", m.GetStored_fields())
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
func (m *SearchRequest) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *SearchRequest) SetAggregationFilters(value []string)() {
    m.aggregationFilters = value
}
func (m *SearchRequest) SetAggregations(value []AggregationOption)() {
    m.aggregations = value
}
func (m *SearchRequest) SetContentSources(value []string)() {
    m.contentSources = value
}
func (m *SearchRequest) SetEnableTopResults(value *bool)() {
    m.enableTopResults = value
}
func (m *SearchRequest) SetEntityTypes(value []EntityType)() {
    m.entityTypes = value
}
func (m *SearchRequest) SetFields(value []string)() {
    m.fields = value
}
func (m *SearchRequest) SetFrom(value *int32)() {
    m.from = value
}
func (m *SearchRequest) SetQuery(value *SearchQuery)() {
    m.query = value
}
func (m *SearchRequest) SetResultTemplateOptions(value *ResultTemplateOption)() {
    m.resultTemplateOptions = value
}
func (m *SearchRequest) SetSize(value *int32)() {
    m.size = value
}
func (m *SearchRequest) SetSortProperties(value []SortProperty)() {
    m.sortProperties = value
}
func (m *SearchRequest) SetStored_fields(value []string)() {
    m.stored_fields = value
}

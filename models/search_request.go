package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SearchRequest 
type SearchRequest struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Contains one or more filters to obtain search results aggregated and filtered to a specific value of a field. Optional.Build this filter based on a prior search that aggregates by the same field. From the response of the prior search, identify the searchBucket that filters results to the specific value of the field, use the string in its aggregationFilterToken property, and build an aggregation filter string in the format '{field}:/'{aggregationFilterToken}/''. If multiple values for the same field need to be provided, use the strings in its aggregationFilterToken property and build an aggregation filter string in the format '{field}:or(/'{aggregationFilterToken1}/',/'{aggregationFilterToken2}/')'. For example, searching and aggregating drive items by file type returns a searchBucket for the file type docx in the response. You can conveniently use the aggregationFilterToken returned for this searchBucket in a subsequent search query and filter matches down to drive items of the docx file type. Example 1 and example 2 show the actual requests and responses.
    aggregationFilters []string
    // Specifies aggregations (also known as refiners) to be returned alongside search results. Optional.
    aggregations []AggregationOptionable
    // Contains the ordered collection of fields and limit to collapse results. Optional.
    collapseProperties []CollapsePropertyable
    // Contains the connection to be targeted. Respects the following format : /external/connections/connectionid where connectionid is the ConnectionId defined in the Connectors Administration.  Note: contentSource is only applicable when entityType=externalItem. Optional.
    contentSources []string
    // This triggers hybrid sort for messages: the first 3 messages are the most relevant. This property is only applicable to entityType=message. Optional.
    enableTopResults *bool
    // One or more types of resources expected in the response. Possible values are: list, site, listItem, message, event, drive, driveItem, person, externalItem, acronym, bookmark, chatMessage. For details about combinations of two or more entity types that are supported in the same search request, see known limitations. Required.
    entityTypes []EntityType
    // Contains the fields to be returned for each resource object specified in entityTypes, allowing customization of the fields returned by default otherwise, including additional fields such as custom managed properties from SharePoint and OneDrive, or custom fields in externalItem from content that Microsoft Graph connectors bring in. The fields property can be using the semantic labels applied to properties. For example, if a property is label as title, you can retrieve it using the following syntax : label_title.Optional.
    fields []string
    // Specifies the offset for the search results. Offset 0 returns the very first result. Optional.
    from *int32
    // The OdataType property
    odataType *string
    // The query property
    query SearchQueryable
    // Provides query alteration options formatted as a JSON blob that contains two optional flags related to spelling correction. Optional.
    queryAlterationOptions SearchAlterationOptionsable
    // Required for searches that use application permissions. Represents the geographic location for the search. For details, see Get the region value.
    region *string
    // Provides the search result templates options for rendering connectors search results.
    resultTemplateOptions ResultTemplateOptionable
    // Indicates the kind of contents to be searched when a search is performed using application permissions. Optional.
    sharePointOneDriveOptions SharePointOneDriveOptionsable
    // The size of the page to be retrieved. Optional.
    size *int32
    // Contains the ordered collection of fields and direction to sort results. There can be at most 5 sort properties in the collection. Optional.
    sortProperties []SortPropertyable
    // The stored_fields property
    stored_fields []string
    // Indicates whether to trim away the duplicate SharePoint files from search results. Default value is false. Optional.
    trimDuplicates *bool
}
// NewSearchRequest instantiates a new searchRequest and sets the default values.
func NewSearchRequest()(*SearchRequest) {
    m := &SearchRequest{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateSearchRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSearchRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSearchRequest(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SearchRequest) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAggregationFilters gets the aggregationFilters property value. Contains one or more filters to obtain search results aggregated and filtered to a specific value of a field. Optional.Build this filter based on a prior search that aggregates by the same field. From the response of the prior search, identify the searchBucket that filters results to the specific value of the field, use the string in its aggregationFilterToken property, and build an aggregation filter string in the format '{field}:/'{aggregationFilterToken}/''. If multiple values for the same field need to be provided, use the strings in its aggregationFilterToken property and build an aggregation filter string in the format '{field}:or(/'{aggregationFilterToken1}/',/'{aggregationFilterToken2}/')'. For example, searching and aggregating drive items by file type returns a searchBucket for the file type docx in the response. You can conveniently use the aggregationFilterToken returned for this searchBucket in a subsequent search query and filter matches down to drive items of the docx file type. Example 1 and example 2 show the actual requests and responses.
func (m *SearchRequest) GetAggregationFilters()([]string) {
    return m.aggregationFilters
}
// GetAggregations gets the aggregations property value. Specifies aggregations (also known as refiners) to be returned alongside search results. Optional.
func (m *SearchRequest) GetAggregations()([]AggregationOptionable) {
    return m.aggregations
}
// GetCollapseProperties gets the collapseProperties property value. Contains the ordered collection of fields and limit to collapse results. Optional.
func (m *SearchRequest) GetCollapseProperties()([]CollapsePropertyable) {
    return m.collapseProperties
}
// GetContentSources gets the contentSources property value. Contains the connection to be targeted. Respects the following format : /external/connections/connectionid where connectionid is the ConnectionId defined in the Connectors Administration.  Note: contentSource is only applicable when entityType=externalItem. Optional.
func (m *SearchRequest) GetContentSources()([]string) {
    return m.contentSources
}
// GetEnableTopResults gets the enableTopResults property value. This triggers hybrid sort for messages: the first 3 messages are the most relevant. This property is only applicable to entityType=message. Optional.
func (m *SearchRequest) GetEnableTopResults()(*bool) {
    return m.enableTopResults
}
// GetEntityTypes gets the entityTypes property value. One or more types of resources expected in the response. Possible values are: list, site, listItem, message, event, drive, driveItem, person, externalItem, acronym, bookmark, chatMessage. For details about combinations of two or more entity types that are supported in the same search request, see known limitations. Required.
func (m *SearchRequest) GetEntityTypes()([]EntityType) {
    return m.entityTypes
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SearchRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["aggregationFilters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetAggregationFilters(res)
        }
        return nil
    }
    res["aggregations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAggregationOptionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AggregationOptionable, len(val))
            for i, v := range val {
                res[i] = v.(AggregationOptionable)
            }
            m.SetAggregations(res)
        }
        return nil
    }
    res["collapseProperties"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCollapsePropertyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CollapsePropertyable, len(val))
            for i, v := range val {
                res[i] = v.(CollapsePropertyable)
            }
            m.SetCollapseProperties(res)
        }
        return nil
    }
    res["contentSources"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetContentSources(res)
        }
        return nil
    }
    res["enableTopResults"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableTopResults(val)
        }
        return nil
    }
    res["entityTypes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseEntityType)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EntityType, len(val))
            for i, v := range val {
                res[i] = *(v.(*EntityType))
            }
            m.SetEntityTypes(res)
        }
        return nil
    }
    res["fields"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetFields(res)
        }
        return nil
    }
    res["from"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFrom(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["query"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSearchQueryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQuery(val.(SearchQueryable))
        }
        return nil
    }
    res["queryAlterationOptions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSearchAlterationOptionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQueryAlterationOptions(val.(SearchAlterationOptionsable))
        }
        return nil
    }
    res["region"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRegion(val)
        }
        return nil
    }
    res["resultTemplateOptions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateResultTemplateOptionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResultTemplateOptions(val.(ResultTemplateOptionable))
        }
        return nil
    }
    res["sharePointOneDriveOptions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSharePointOneDriveOptionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSharePointOneDriveOptions(val.(SharePointOneDriveOptionsable))
        }
        return nil
    }
    res["size"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSize(val)
        }
        return nil
    }
    res["sortProperties"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSortPropertyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SortPropertyable, len(val))
            for i, v := range val {
                res[i] = v.(SortPropertyable)
            }
            m.SetSortProperties(res)
        }
        return nil
    }
    res["stored_fields"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetStored_fields(res)
        }
        return nil
    }
    res["trimDuplicates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTrimDuplicates(val)
        }
        return nil
    }
    return res
}
// GetFields gets the fields property value. Contains the fields to be returned for each resource object specified in entityTypes, allowing customization of the fields returned by default otherwise, including additional fields such as custom managed properties from SharePoint and OneDrive, or custom fields in externalItem from content that Microsoft Graph connectors bring in. The fields property can be using the semantic labels applied to properties. For example, if a property is label as title, you can retrieve it using the following syntax : label_title.Optional.
func (m *SearchRequest) GetFields()([]string) {
    return m.fields
}
// GetFrom gets the from property value. Specifies the offset for the search results. Offset 0 returns the very first result. Optional.
func (m *SearchRequest) GetFrom()(*int32) {
    return m.from
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *SearchRequest) GetOdataType()(*string) {
    return m.odataType
}
// GetQuery gets the query property value. The query property
func (m *SearchRequest) GetQuery()(SearchQueryable) {
    return m.query
}
// GetQueryAlterationOptions gets the queryAlterationOptions property value. Provides query alteration options formatted as a JSON blob that contains two optional flags related to spelling correction. Optional.
func (m *SearchRequest) GetQueryAlterationOptions()(SearchAlterationOptionsable) {
    return m.queryAlterationOptions
}
// GetRegion gets the region property value. Required for searches that use application permissions. Represents the geographic location for the search. For details, see Get the region value.
func (m *SearchRequest) GetRegion()(*string) {
    return m.region
}
// GetResultTemplateOptions gets the resultTemplateOptions property value. Provides the search result templates options for rendering connectors search results.
func (m *SearchRequest) GetResultTemplateOptions()(ResultTemplateOptionable) {
    return m.resultTemplateOptions
}
// GetSharePointOneDriveOptions gets the sharePointOneDriveOptions property value. Indicates the kind of contents to be searched when a search is performed using application permissions. Optional.
func (m *SearchRequest) GetSharePointOneDriveOptions()(SharePointOneDriveOptionsable) {
    return m.sharePointOneDriveOptions
}
// GetSize gets the size property value. The size of the page to be retrieved. Optional.
func (m *SearchRequest) GetSize()(*int32) {
    return m.size
}
// GetSortProperties gets the sortProperties property value. Contains the ordered collection of fields and direction to sort results. There can be at most 5 sort properties in the collection. Optional.
func (m *SearchRequest) GetSortProperties()([]SortPropertyable) {
    return m.sortProperties
}
// GetStored_fields gets the stored_fields property value. The stored_fields property
func (m *SearchRequest) GetStored_fields()([]string) {
    return m.stored_fields
}
// GetTrimDuplicates gets the trimDuplicates property value. Indicates whether to trim away the duplicate SharePoint files from search results. Default value is false. Optional.
func (m *SearchRequest) GetTrimDuplicates()(*bool) {
    return m.trimDuplicates
}
// Serialize serializes information the current object
func (m *SearchRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAggregationFilters() != nil {
        err := writer.WriteCollectionOfStringValues("aggregationFilters", m.GetAggregationFilters())
        if err != nil {
            return err
        }
    }
    if m.GetAggregations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAggregations()))
        for i, v := range m.GetAggregations() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("aggregations", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCollapseProperties() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCollapseProperties()))
        for i, v := range m.GetCollapseProperties() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("collapseProperties", cast)
        if err != nil {
            return err
        }
    }
    if m.GetContentSources() != nil {
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
    if m.GetEntityTypes() != nil {
        err := writer.WriteCollectionOfStringValues("entityTypes", SerializeEntityType(m.GetEntityTypes()))
        if err != nil {
            return err
        }
    }
    if m.GetFields() != nil {
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
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
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
        err := writer.WriteObjectValue("queryAlterationOptions", m.GetQueryAlterationOptions())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("region", m.GetRegion())
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
        err := writer.WriteObjectValue("sharePointOneDriveOptions", m.GetSharePointOneDriveOptions())
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
    if m.GetSortProperties() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSortProperties()))
        for i, v := range m.GetSortProperties() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("sortProperties", cast)
        if err != nil {
            return err
        }
    }
    if m.GetStored_fields() != nil {
        err := writer.WriteCollectionOfStringValues("stored_fields", m.GetStored_fields())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("trimDuplicates", m.GetTrimDuplicates())
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
func (m *SearchRequest) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAggregationFilters sets the aggregationFilters property value. Contains one or more filters to obtain search results aggregated and filtered to a specific value of a field. Optional.Build this filter based on a prior search that aggregates by the same field. From the response of the prior search, identify the searchBucket that filters results to the specific value of the field, use the string in its aggregationFilterToken property, and build an aggregation filter string in the format '{field}:/'{aggregationFilterToken}/''. If multiple values for the same field need to be provided, use the strings in its aggregationFilterToken property and build an aggregation filter string in the format '{field}:or(/'{aggregationFilterToken1}/',/'{aggregationFilterToken2}/')'. For example, searching and aggregating drive items by file type returns a searchBucket for the file type docx in the response. You can conveniently use the aggregationFilterToken returned for this searchBucket in a subsequent search query and filter matches down to drive items of the docx file type. Example 1 and example 2 show the actual requests and responses.
func (m *SearchRequest) SetAggregationFilters(value []string)() {
    m.aggregationFilters = value
}
// SetAggregations sets the aggregations property value. Specifies aggregations (also known as refiners) to be returned alongside search results. Optional.
func (m *SearchRequest) SetAggregations(value []AggregationOptionable)() {
    m.aggregations = value
}
// SetCollapseProperties sets the collapseProperties property value. Contains the ordered collection of fields and limit to collapse results. Optional.
func (m *SearchRequest) SetCollapseProperties(value []CollapsePropertyable)() {
    m.collapseProperties = value
}
// SetContentSources sets the contentSources property value. Contains the connection to be targeted. Respects the following format : /external/connections/connectionid where connectionid is the ConnectionId defined in the Connectors Administration.  Note: contentSource is only applicable when entityType=externalItem. Optional.
func (m *SearchRequest) SetContentSources(value []string)() {
    m.contentSources = value
}
// SetEnableTopResults sets the enableTopResults property value. This triggers hybrid sort for messages: the first 3 messages are the most relevant. This property is only applicable to entityType=message. Optional.
func (m *SearchRequest) SetEnableTopResults(value *bool)() {
    m.enableTopResults = value
}
// SetEntityTypes sets the entityTypes property value. One or more types of resources expected in the response. Possible values are: list, site, listItem, message, event, drive, driveItem, person, externalItem, acronym, bookmark, chatMessage. For details about combinations of two or more entity types that are supported in the same search request, see known limitations. Required.
func (m *SearchRequest) SetEntityTypes(value []EntityType)() {
    m.entityTypes = value
}
// SetFields sets the fields property value. Contains the fields to be returned for each resource object specified in entityTypes, allowing customization of the fields returned by default otherwise, including additional fields such as custom managed properties from SharePoint and OneDrive, or custom fields in externalItem from content that Microsoft Graph connectors bring in. The fields property can be using the semantic labels applied to properties. For example, if a property is label as title, you can retrieve it using the following syntax : label_title.Optional.
func (m *SearchRequest) SetFields(value []string)() {
    m.fields = value
}
// SetFrom sets the from property value. Specifies the offset for the search results. Offset 0 returns the very first result. Optional.
func (m *SearchRequest) SetFrom(value *int32)() {
    m.from = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *SearchRequest) SetOdataType(value *string)() {
    m.odataType = value
}
// SetQuery sets the query property value. The query property
func (m *SearchRequest) SetQuery(value SearchQueryable)() {
    m.query = value
}
// SetQueryAlterationOptions sets the queryAlterationOptions property value. Provides query alteration options formatted as a JSON blob that contains two optional flags related to spelling correction. Optional.
func (m *SearchRequest) SetQueryAlterationOptions(value SearchAlterationOptionsable)() {
    m.queryAlterationOptions = value
}
// SetRegion sets the region property value. Required for searches that use application permissions. Represents the geographic location for the search. For details, see Get the region value.
func (m *SearchRequest) SetRegion(value *string)() {
    m.region = value
}
// SetResultTemplateOptions sets the resultTemplateOptions property value. Provides the search result templates options for rendering connectors search results.
func (m *SearchRequest) SetResultTemplateOptions(value ResultTemplateOptionable)() {
    m.resultTemplateOptions = value
}
// SetSharePointOneDriveOptions sets the sharePointOneDriveOptions property value. Indicates the kind of contents to be searched when a search is performed using application permissions. Optional.
func (m *SearchRequest) SetSharePointOneDriveOptions(value SharePointOneDriveOptionsable)() {
    m.sharePointOneDriveOptions = value
}
// SetSize sets the size property value. The size of the page to be retrieved. Optional.
func (m *SearchRequest) SetSize(value *int32)() {
    m.size = value
}
// SetSortProperties sets the sortProperties property value. Contains the ordered collection of fields and direction to sort results. There can be at most 5 sort properties in the collection. Optional.
func (m *SearchRequest) SetSortProperties(value []SortPropertyable)() {
    m.sortProperties = value
}
// SetStored_fields sets the stored_fields property value. The stored_fields property
func (m *SearchRequest) SetStored_fields(value []string)() {
    m.stored_fields = value
}
// SetTrimDuplicates sets the trimDuplicates property value. Indicates whether to trim away the duplicate SharePoint files from search results. Default value is false. Optional.
func (m *SearchRequest) SetTrimDuplicates(value *bool)() {
    m.trimDuplicates = value
}

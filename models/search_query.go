package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SearchQuery 
type SearchQuery struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The OdataType property
    odataType *string
    // The query_string property
    query_string SearchQueryStringable
    // The search query containing the search terms. Required.
    queryString *string
    // Provides a way to decorate the query string. Supports both KQL and query variables. Optional.
    queryTemplate *string
}
// NewSearchQuery instantiates a new searchQuery and sets the default values.
func NewSearchQuery()(*SearchQuery) {
    m := &SearchQuery{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.searchQuery";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateSearchQueryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSearchQueryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSearchQuery(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SearchQuery) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SearchQuery) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["query_string"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateSearchQueryStringFromDiscriminatorValue , m.SetQuery_string)
    res["queryString"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetQueryString)
    res["queryTemplate"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetQueryTemplate)
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *SearchQuery) GetOdataType()(*string) {
    return m.odataType
}
// GetQuery_string gets the query_string property value. The query_string property
func (m *SearchQuery) GetQuery_string()(SearchQueryStringable) {
    return m.query_string
}
// GetQueryString gets the queryString property value. The search query containing the search terms. Required.
func (m *SearchQuery) GetQueryString()(*string) {
    return m.queryString
}
// GetQueryTemplate gets the queryTemplate property value. Provides a way to decorate the query string. Supports both KQL and query variables. Optional.
func (m *SearchQuery) GetQueryTemplate()(*string) {
    return m.queryTemplate
}
// Serialize serializes information the current object
func (m *SearchQuery) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
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
        err := writer.WriteObjectValue("query_string", m.GetQuery_string())
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
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *SearchQuery) SetOdataType(value *string)() {
    m.odataType = value
}
// SetQuery_string sets the query_string property value. The query_string property
func (m *SearchQuery) SetQuery_string(value SearchQueryStringable)() {
    m.query_string = value
}
// SetQueryString sets the queryString property value. The search query containing the search terms. Required.
func (m *SearchQuery) SetQueryString(value *string)() {
    m.queryString = value
}
// SetQueryTemplate sets the queryTemplate property value. Provides a way to decorate the query string. Supports both KQL and query variables. Optional.
func (m *SearchQuery) SetQueryTemplate(value *string)() {
    m.queryTemplate = value
}

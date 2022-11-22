package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SearchHit 
type SearchHit struct {
    // The _id property
    _id *string
    // The _score property
    _score *int32
    // The _source property
    _source Entityable
    // The _summary property
    _summary *string
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The name of the content source which the externalItem is part of .
    contentSource *string
    // The internal identifier for the item. The format of the identifier varies based on the entity type. For details, see hitId format.
    hitId *string
    // The OdataType property
    odataType *string
    // The rank or the order of the result.
    rank *int32
    // The resource property
    resource Entityable
    // ID of the result template for rendering the search result. This ID must map to a display layout in the resultTemplates dictionary, included in the searchresponse as well.
    resultTemplateId *string
    // A summary of the result, if a summary is available.
    summary *string
}
// NewSearchHit instantiates a new searchHit and sets the default values.
func NewSearchHit()(*SearchHit) {
    m := &SearchHit{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateSearchHitFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSearchHitFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSearchHit(), nil
}
// Get_id gets the _id property value. The _id property
func (m *SearchHit) Get_id()(*string) {
    return m._id
}
// Get_score gets the _score property value. The _score property
func (m *SearchHit) Get_score()(*int32) {
    return m._score
}
// Get_source gets the _source property value. The _source property
func (m *SearchHit) Get_source()(Entityable) {
    return m._source
}
// Get_summary gets the _summary property value. The _summary property
func (m *SearchHit) Get_summary()(*string) {
    return m._summary
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SearchHit) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetContentSource gets the contentSource property value. The name of the content source which the externalItem is part of .
func (m *SearchHit) GetContentSource()(*string) {
    return m.contentSource
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SearchHit) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["_id"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.Set_id)
    res["_score"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.Set_score)
    res["_source"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateEntityFromDiscriminatorValue , m.Set_source)
    res["_summary"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.Set_summary)
    res["contentSource"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetContentSource)
    res["hitId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetHitId)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["rank"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetRank)
    res["resource"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateEntityFromDiscriminatorValue , m.SetResource)
    res["resultTemplateId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetResultTemplateId)
    res["summary"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetSummary)
    return res
}
// GetHitId gets the hitId property value. The internal identifier for the item. The format of the identifier varies based on the entity type. For details, see hitId format.
func (m *SearchHit) GetHitId()(*string) {
    return m.hitId
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *SearchHit) GetOdataType()(*string) {
    return m.odataType
}
// GetRank gets the rank property value. The rank or the order of the result.
func (m *SearchHit) GetRank()(*int32) {
    return m.rank
}
// GetResource gets the resource property value. The resource property
func (m *SearchHit) GetResource()(Entityable) {
    return m.resource
}
// GetResultTemplateId gets the resultTemplateId property value. ID of the result template for rendering the search result. This ID must map to a display layout in the resultTemplates dictionary, included in the searchresponse as well.
func (m *SearchHit) GetResultTemplateId()(*string) {
    return m.resultTemplateId
}
// GetSummary gets the summary property value. A summary of the result, if a summary is available.
func (m *SearchHit) GetSummary()(*string) {
    return m.summary
}
// Serialize serializes information the current object
func (m *SearchHit) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("contentSource", m.GetContentSource())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("hitId", m.GetHitId())
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
        err := writer.WriteInt32Value("rank", m.GetRank())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("resource", m.GetResource())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("resultTemplateId", m.GetResultTemplateId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("summary", m.GetSummary())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("_id", m.Get_id())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("_score", m.Get_score())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("_source", m.Get_source())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("_summary", m.Get_summary())
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
// Set_id sets the _id property value. The _id property
func (m *SearchHit) Set_id(value *string)() {
    m._id = value
}
// Set_score sets the _score property value. The _score property
func (m *SearchHit) Set_score(value *int32)() {
    m._score = value
}
// Set_source sets the _source property value. The _source property
func (m *SearchHit) Set_source(value Entityable)() {
    m._source = value
}
// Set_summary sets the _summary property value. The _summary property
func (m *SearchHit) Set_summary(value *string)() {
    m._summary = value
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SearchHit) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetContentSource sets the contentSource property value. The name of the content source which the externalItem is part of .
func (m *SearchHit) SetContentSource(value *string)() {
    m.contentSource = value
}
// SetHitId sets the hitId property value. The internal identifier for the item. The format of the identifier varies based on the entity type. For details, see hitId format.
func (m *SearchHit) SetHitId(value *string)() {
    m.hitId = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *SearchHit) SetOdataType(value *string)() {
    m.odataType = value
}
// SetRank sets the rank property value. The rank or the order of the result.
func (m *SearchHit) SetRank(value *int32)() {
    m.rank = value
}
// SetResource sets the resource property value. The resource property
func (m *SearchHit) SetResource(value Entityable)() {
    m.resource = value
}
// SetResultTemplateId sets the resultTemplateId property value. ID of the result template for rendering the search result. This ID must map to a display layout in the resultTemplates dictionary, included in the searchresponse as well.
func (m *SearchHit) SetResultTemplateId(value *string)() {
    m.resultTemplateId = value
}
// SetSummary sets the summary property value. A summary of the result, if a summary is available.
func (m *SearchHit) SetSummary(value *string)() {
    m.summary = value
}

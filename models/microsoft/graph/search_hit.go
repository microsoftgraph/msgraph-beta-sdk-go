package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type SearchHit struct {
    // 
    _id *string;
    // 
    _score *int32;
    // 
    _source *Entity;
    // 
    _summary *string;
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The name of the content source which the externalItem is part of .
    contentSource *string;
    // The internal identifier for the item.
    hitId *string;
    // The rank or the order of the result.
    rank *int32;
    // 
    resource *Entity;
    // ID of the result template for rendering the search result. This ID must map to a display layout in the resultTemplates dictionary, included in the searchresponse as well.
    resultTemplateId *string;
    // A summary of the result, if a summary is available.
    summary *string;
}
// Instantiates a new searchHit and sets the default values.
func NewSearchHit()(*SearchHit) {
    m := &SearchHit{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the _id property value. 
func (m *SearchHit) Get_id()(*string) {
    if m == nil {
        return nil
    } else {
        return m._id
    }
}
// Gets the _score property value. 
func (m *SearchHit) Get_score()(*int32) {
    if m == nil {
        return nil
    } else {
        return m._score
    }
}
// Gets the _source property value. 
func (m *SearchHit) Get_source()(*Entity) {
    if m == nil {
        return nil
    } else {
        return m._source
    }
}
// Gets the _summary property value. 
func (m *SearchHit) Get_summary()(*string) {
    if m == nil {
        return nil
    } else {
        return m._summary
    }
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SearchHit) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the contentSource property value. The name of the content source which the externalItem is part of .
func (m *SearchHit) GetContentSource()(*string) {
    if m == nil {
        return nil
    } else {
        return m.contentSource
    }
}
// Gets the hitId property value. The internal identifier for the item.
func (m *SearchHit) GetHitId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.hitId
    }
}
// Gets the rank property value. The rank or the order of the result.
func (m *SearchHit) GetRank()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.rank
    }
}
// Gets the resource property value. 
func (m *SearchHit) GetResource()(*Entity) {
    if m == nil {
        return nil
    } else {
        return m.resource
    }
}
// Gets the resultTemplateId property value. ID of the result template for rendering the search result. This ID must map to a display layout in the resultTemplates dictionary, included in the searchresponse as well.
func (m *SearchHit) GetResultTemplateId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.resultTemplateId
    }
}
// Gets the summary property value. A summary of the result, if a summary is available.
func (m *SearchHit) GetSummary()(*string) {
    if m == nil {
        return nil
    } else {
        return m.summary
    }
}
// The deserialization information for the current model
func (m *SearchHit) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["_id"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.Set_id(val)
        }
        return nil
    }
    res["_score"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.Set_score(val)
        }
        return nil
    }
    res["_source"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEntity() })
        if err != nil {
            return err
        }
        if val != nil {
            m.Set_source(val.(*Entity))
        }
        return nil
    }
    res["_summary"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.Set_summary(val)
        }
        return nil
    }
    res["contentSource"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentSource(val)
        }
        return nil
    }
    res["hitId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHitId(val)
        }
        return nil
    }
    res["rank"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRank(val)
        }
        return nil
    }
    res["resource"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEntity() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResource(val.(*Entity))
        }
        return nil
    }
    res["resultTemplateId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResultTemplateId(val)
        }
        return nil
    }
    res["summary"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSummary(val)
        }
        return nil
    }
    return res
}
func (m *SearchHit) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *SearchHit) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the _id property value. 
// Parameters:
//  - value : Value to set for the _id property.
func (m *SearchHit) Set_id(value *string)() {
    m._id = value
}
// Sets the _score property value. 
// Parameters:
//  - value : Value to set for the _score property.
func (m *SearchHit) Set_score(value *int32)() {
    m._score = value
}
// Sets the _source property value. 
// Parameters:
//  - value : Value to set for the _source property.
func (m *SearchHit) Set_source(value *Entity)() {
    m._source = value
}
// Sets the _summary property value. 
// Parameters:
//  - value : Value to set for the _summary property.
func (m *SearchHit) Set_summary(value *string)() {
    m._summary = value
}
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *SearchHit) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the contentSource property value. The name of the content source which the externalItem is part of .
// Parameters:
//  - value : Value to set for the contentSource property.
func (m *SearchHit) SetContentSource(value *string)() {
    m.contentSource = value
}
// Sets the hitId property value. The internal identifier for the item.
// Parameters:
//  - value : Value to set for the hitId property.
func (m *SearchHit) SetHitId(value *string)() {
    m.hitId = value
}
// Sets the rank property value. The rank or the order of the result.
// Parameters:
//  - value : Value to set for the rank property.
func (m *SearchHit) SetRank(value *int32)() {
    m.rank = value
}
// Sets the resource property value. 
// Parameters:
//  - value : Value to set for the resource property.
func (m *SearchHit) SetResource(value *Entity)() {
    m.resource = value
}
// Sets the resultTemplateId property value. ID of the result template for rendering the search result. This ID must map to a display layout in the resultTemplates dictionary, included in the searchresponse as well.
// Parameters:
//  - value : Value to set for the resultTemplateId property.
func (m *SearchHit) SetResultTemplateId(value *string)() {
    m.resultTemplateId = value
}
// Sets the summary property value. A summary of the result, if a summary is available.
// Parameters:
//  - value : Value to set for the summary property.
func (m *SearchHit) SetSummary(value *string)() {
    m.summary = value
}

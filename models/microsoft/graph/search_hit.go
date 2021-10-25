package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type SearchHit struct {
    _id *string;
    _score *int32;
    _source *Entity;
    _summary *string;
    additionalData map[string]interface{};
    contentSource *string;
    hitId *string;
    rank *int32;
    resource *Entity;
    resultTemplateId *string;
    summary *string;
}
func NewSearchHit()(*SearchHit) {
    m := &SearchHit{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *SearchHit) Get_id()(*string) {
    if m == nil {
        return nil
    } else {
        return m._id
    }
}
func (m *SearchHit) Get_score()(*int32) {
    if m == nil {
        return nil
    } else {
        return m._score
    }
}
func (m *SearchHit) Get_source()(*Entity) {
    if m == nil {
        return nil
    } else {
        return m._source
    }
}
func (m *SearchHit) Get_summary()(*string) {
    if m == nil {
        return nil
    } else {
        return m._summary
    }
}
func (m *SearchHit) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *SearchHit) GetContentSource()(*string) {
    if m == nil {
        return nil
    } else {
        return m.contentSource
    }
}
func (m *SearchHit) GetHitId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.hitId
    }
}
func (m *SearchHit) GetRank()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.rank
    }
}
func (m *SearchHit) GetResource()(*Entity) {
    if m == nil {
        return nil
    } else {
        return m.resource
    }
}
func (m *SearchHit) GetResultTemplateId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.resultTemplateId
    }
}
func (m *SearchHit) GetSummary()(*string) {
    if m == nil {
        return nil
    } else {
        return m.summary
    }
}
func (m *SearchHit) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["_id"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.Set_id(val)
        return nil
    }
    res["_score"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.Set_score(val)
        return nil
    }
    res["_source"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEntity() })
        if err != nil {
            return err
        }
        m.Set_source(val.(*Entity))
        return nil
    }
    res["_summary"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.Set_summary(val)
        return nil
    }
    res["contentSource"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetContentSource(val)
        return nil
    }
    res["hitId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetHitId(val)
        return nil
    }
    res["rank"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetRank(val)
        return nil
    }
    res["resource"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEntity() })
        if err != nil {
            return err
        }
        m.SetResource(val.(*Entity))
        return nil
    }
    res["resultTemplateId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetResultTemplateId(val)
        return nil
    }
    res["summary"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSummary(val)
        return nil
    }
    return res
}
func (m *SearchHit) IsNil()(bool) {
    return m == nil
}
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
func (m *SearchHit) Set_id(value *string)() {
    m._id = value
}
func (m *SearchHit) Set_score(value *int32)() {
    m._score = value
}
func (m *SearchHit) Set_source(value *Entity)() {
    m._source = value
}
func (m *SearchHit) Set_summary(value *string)() {
    m._summary = value
}
func (m *SearchHit) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *SearchHit) SetContentSource(value *string)() {
    m.contentSource = value
}
func (m *SearchHit) SetHitId(value *string)() {
    m.hitId = value
}
func (m *SearchHit) SetRank(value *int32)() {
    m.rank = value
}
func (m *SearchHit) SetResource(value *Entity)() {
    m.resource = value
}
func (m *SearchHit) SetResultTemplateId(value *string)() {
    m.resultTemplateId = value
}
func (m *SearchHit) SetSummary(value *string)() {
    m.summary = value
}

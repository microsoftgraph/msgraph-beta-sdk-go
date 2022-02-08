package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// RecommendationResource 
type RecommendationResource struct {
    Entity
    // 
    addedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    additionalDetails []KeyValue;
    // 
    apiUrl *string;
    // 
    displayName *string;
    // 
    owner *string;
    // 
    portalUrl *string;
    // 
    rank *int32;
    // 
    recommendationId *string;
    // 
    resourceType *string;
    // 
    status *RecommendationStatus;
}
// NewRecommendationResource instantiates a new recommendationResource and sets the default values.
func NewRecommendationResource()(*RecommendationResource) {
    m := &RecommendationResource{
        Entity: *NewEntity(),
    }
    return m
}
// GetAddedDateTime gets the addedDateTime property value. 
func (m *RecommendationResource) GetAddedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.addedDateTime
    }
}
// GetAdditionalDetails gets the additionalDetails property value. 
func (m *RecommendationResource) GetAdditionalDetails()([]KeyValue) {
    if m == nil {
        return nil
    } else {
        return m.additionalDetails
    }
}
// GetApiUrl gets the apiUrl property value. 
func (m *RecommendationResource) GetApiUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.apiUrl
    }
}
// GetDisplayName gets the displayName property value. 
func (m *RecommendationResource) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetOwner gets the owner property value. 
func (m *RecommendationResource) GetOwner()(*string) {
    if m == nil {
        return nil
    } else {
        return m.owner
    }
}
// GetPortalUrl gets the portalUrl property value. 
func (m *RecommendationResource) GetPortalUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.portalUrl
    }
}
// GetRank gets the rank property value. 
func (m *RecommendationResource) GetRank()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.rank
    }
}
// GetRecommendationId gets the recommendationId property value. 
func (m *RecommendationResource) GetRecommendationId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.recommendationId
    }
}
// GetResourceType gets the resourceType property value. 
func (m *RecommendationResource) GetResourceType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.resourceType
    }
}
// GetStatus gets the status property value. 
func (m *RecommendationResource) GetStatus()(*RecommendationStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RecommendationResource) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["addedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAddedDateTime(val)
        }
        return nil
    }
    res["additionalDetails"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewKeyValue() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]KeyValue, len(val))
            for i, v := range val {
                res[i] = *(v.(*KeyValue))
            }
            m.SetAdditionalDetails(res)
        }
        return nil
    }
    res["apiUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApiUrl(val)
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["owner"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwner(val)
        }
        return nil
    }
    res["portalUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPortalUrl(val)
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
    res["recommendationId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecommendationId(val)
        }
        return nil
    }
    res["resourceType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceType(val)
        }
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseRecommendationStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*RecommendationStatus))
        }
        return nil
    }
    return res
}
func (m *RecommendationResource) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *RecommendationResource) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("addedDateTime", m.GetAddedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetAdditionalDetails() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAdditionalDetails()))
        for i, v := range m.GetAdditionalDetails() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("additionalDetails", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("apiUrl", m.GetApiUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("owner", m.GetOwner())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("portalUrl", m.GetPortalUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("rank", m.GetRank())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("recommendationId", m.GetRecommendationId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("resourceType", m.GetResourceType())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAddedDateTime sets the addedDateTime property value. 
func (m *RecommendationResource) SetAddedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.addedDateTime = value
    }
}
// SetAdditionalDetails sets the additionalDetails property value. 
func (m *RecommendationResource) SetAdditionalDetails(value []KeyValue)() {
    if m != nil {
        m.additionalDetails = value
    }
}
// SetApiUrl sets the apiUrl property value. 
func (m *RecommendationResource) SetApiUrl(value *string)() {
    if m != nil {
        m.apiUrl = value
    }
}
// SetDisplayName sets the displayName property value. 
func (m *RecommendationResource) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetOwner sets the owner property value. 
func (m *RecommendationResource) SetOwner(value *string)() {
    if m != nil {
        m.owner = value
    }
}
// SetPortalUrl sets the portalUrl property value. 
func (m *RecommendationResource) SetPortalUrl(value *string)() {
    if m != nil {
        m.portalUrl = value
    }
}
// SetRank sets the rank property value. 
func (m *RecommendationResource) SetRank(value *int32)() {
    if m != nil {
        m.rank = value
    }
}
// SetRecommendationId sets the recommendationId property value. 
func (m *RecommendationResource) SetRecommendationId(value *string)() {
    if m != nil {
        m.recommendationId = value
    }
}
// SetResourceType sets the resourceType property value. 
func (m *RecommendationResource) SetResourceType(value *string)() {
    if m != nil {
        m.resourceType = value
    }
}
// SetStatus sets the status property value. 
func (m *RecommendationResource) SetStatus(value *RecommendationStatus)() {
    if m != nil {
        m.status = value
    }
}

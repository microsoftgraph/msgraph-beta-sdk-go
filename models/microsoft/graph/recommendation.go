package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Recommendation 
type Recommendation struct {
    Entity
    // 
    actionSteps []ActionStep;
    // 
    benefits *string;
    // 
    category *RecommendationCategory;
    // 
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    displayName *string;
    // 
    impactedResources []RecommendationResource;
    // 
    impactStartDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    impactType *string;
    // 
    insights *string;
    // 
    lastCheckedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    lastModifiedBy *string;
    // 
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    postponeUntilDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    priority *RecommendationPriority;
    // 
    status *RecommendationStatus;
}
// NewRecommendation instantiates a new recommendation and sets the default values.
func NewRecommendation()(*Recommendation) {
    m := &Recommendation{
        Entity: *NewEntity(),
    }
    return m
}
// GetActionSteps gets the actionSteps property value. 
func (m *Recommendation) GetActionSteps()([]ActionStep) {
    if m == nil {
        return nil
    } else {
        return m.actionSteps
    }
}
// GetBenefits gets the benefits property value. 
func (m *Recommendation) GetBenefits()(*string) {
    if m == nil {
        return nil
    } else {
        return m.benefits
    }
}
// GetCategory gets the category property value. 
func (m *Recommendation) GetCategory()(*RecommendationCategory) {
    if m == nil {
        return nil
    } else {
        return m.category
    }
}
// GetCreatedDateTime gets the createdDateTime property value. 
func (m *Recommendation) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetDisplayName gets the displayName property value. 
func (m *Recommendation) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetImpactedResources gets the impactedResources property value. 
func (m *Recommendation) GetImpactedResources()([]RecommendationResource) {
    if m == nil {
        return nil
    } else {
        return m.impactedResources
    }
}
// GetImpactStartDateTime gets the impactStartDateTime property value. 
func (m *Recommendation) GetImpactStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.impactStartDateTime
    }
}
// GetImpactType gets the impactType property value. 
func (m *Recommendation) GetImpactType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.impactType
    }
}
// GetInsights gets the insights property value. 
func (m *Recommendation) GetInsights()(*string) {
    if m == nil {
        return nil
    } else {
        return m.insights
    }
}
// GetLastCheckedDateTime gets the lastCheckedDateTime property value. 
func (m *Recommendation) GetLastCheckedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastCheckedDateTime
    }
}
// GetLastModifiedBy gets the lastModifiedBy property value. 
func (m *Recommendation) GetLastModifiedBy()(*string) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedBy
    }
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. 
func (m *Recommendation) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// GetPostponeUntilDateTime gets the postponeUntilDateTime property value. 
func (m *Recommendation) GetPostponeUntilDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.postponeUntilDateTime
    }
}
// GetPriority gets the priority property value. 
func (m *Recommendation) GetPriority()(*RecommendationPriority) {
    if m == nil {
        return nil
    } else {
        return m.priority
    }
}
// GetStatus gets the status property value. 
func (m *Recommendation) GetStatus()(*RecommendationStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Recommendation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["actionSteps"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewActionStep() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ActionStep, len(val))
            for i, v := range val {
                res[i] = *(v.(*ActionStep))
            }
            m.SetActionSteps(res)
        }
        return nil
    }
    res["benefits"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBenefits(val)
        }
        return nil
    }
    res["category"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseRecommendationCategory)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory(val.(*RecommendationCategory))
        }
        return nil
    }
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
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
    res["impactedResources"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRecommendationResource() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RecommendationResource, len(val))
            for i, v := range val {
                res[i] = *(v.(*RecommendationResource))
            }
            m.SetImpactedResources(res)
        }
        return nil
    }
    res["impactStartDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetImpactStartDateTime(val)
        }
        return nil
    }
    res["impactType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetImpactType(val)
        }
        return nil
    }
    res["insights"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInsights(val)
        }
        return nil
    }
    res["lastCheckedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastCheckedDateTime(val)
        }
        return nil
    }
    res["lastModifiedBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedBy(val)
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["postponeUntilDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPostponeUntilDateTime(val)
        }
        return nil
    }
    res["priority"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseRecommendationPriority)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPriority(val.(*RecommendationPriority))
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
func (m *Recommendation) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *Recommendation) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetActionSteps() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetActionSteps()))
        for i, v := range m.GetActionSteps() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("actionSteps", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("benefits", m.GetBenefits())
        if err != nil {
            return err
        }
    }
    if m.GetCategory() != nil {
        cast := (*m.GetCategory()).String()
        err = writer.WriteStringValue("category", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
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
    if m.GetImpactedResources() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetImpactedResources()))
        for i, v := range m.GetImpactedResources() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("impactedResources", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("impactStartDateTime", m.GetImpactStartDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("impactType", m.GetImpactType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("insights", m.GetInsights())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastCheckedDateTime", m.GetLastCheckedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("lastModifiedBy", m.GetLastModifiedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("postponeUntilDateTime", m.GetPostponeUntilDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetPriority() != nil {
        cast := (*m.GetPriority()).String()
        err = writer.WriteStringValue("priority", &cast)
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
// SetActionSteps sets the actionSteps property value. 
func (m *Recommendation) SetActionSteps(value []ActionStep)() {
    if m != nil {
        m.actionSteps = value
    }
}
// SetBenefits sets the benefits property value. 
func (m *Recommendation) SetBenefits(value *string)() {
    if m != nil {
        m.benefits = value
    }
}
// SetCategory sets the category property value. 
func (m *Recommendation) SetCategory(value *RecommendationCategory)() {
    if m != nil {
        m.category = value
    }
}
// SetCreatedDateTime sets the createdDateTime property value. 
func (m *Recommendation) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetDisplayName sets the displayName property value. 
func (m *Recommendation) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetImpactedResources sets the impactedResources property value. 
func (m *Recommendation) SetImpactedResources(value []RecommendationResource)() {
    if m != nil {
        m.impactedResources = value
    }
}
// SetImpactStartDateTime sets the impactStartDateTime property value. 
func (m *Recommendation) SetImpactStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.impactStartDateTime = value
    }
}
// SetImpactType sets the impactType property value. 
func (m *Recommendation) SetImpactType(value *string)() {
    if m != nil {
        m.impactType = value
    }
}
// SetInsights sets the insights property value. 
func (m *Recommendation) SetInsights(value *string)() {
    if m != nil {
        m.insights = value
    }
}
// SetLastCheckedDateTime sets the lastCheckedDateTime property value. 
func (m *Recommendation) SetLastCheckedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastCheckedDateTime = value
    }
}
// SetLastModifiedBy sets the lastModifiedBy property value. 
func (m *Recommendation) SetLastModifiedBy(value *string)() {
    if m != nil {
        m.lastModifiedBy = value
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. 
func (m *Recommendation) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastModifiedDateTime = value
    }
}
// SetPostponeUntilDateTime sets the postponeUntilDateTime property value. 
func (m *Recommendation) SetPostponeUntilDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.postponeUntilDateTime = value
    }
}
// SetPriority sets the priority property value. 
func (m *Recommendation) SetPriority(value *RecommendationPriority)() {
    if m != nil {
        m.priority = value
    }
}
// SetStatus sets the status property value. 
func (m *Recommendation) SetStatus(value *RecommendationStatus)() {
    if m != nil {
        m.status = value
    }
}

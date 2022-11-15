package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Recommendation provides operations to manage the collection of accessReviewDecision entities.
type Recommendation struct {
    Entity
    // The actionSteps property
    actionSteps []ActionStepable
    // The benefits property
    benefits *string
    // The category property
    category *RecommendationCategory
    // The createdDateTime property
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The displayName property
    displayName *string
    // The impactedResources property
    impactedResources []RecommendationResourceable
    // The impactStartDateTime property
    impactStartDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The impactType property
    impactType *string
    // The insights property
    insights *string
    // The lastCheckedDateTime property
    lastCheckedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The lastModifiedBy property
    lastModifiedBy *string
    // The lastModifiedDateTime property
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The postponeUntilDateTime property
    postponeUntilDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The priority property
    priority *RecommendationPriority
    // The status property
    status *RecommendationStatus
}
// NewRecommendation instantiates a new recommendation and sets the default values.
func NewRecommendation()(*Recommendation) {
    m := &Recommendation{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.recommendation";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateRecommendationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRecommendationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRecommendation(), nil
}
// GetActionSteps gets the actionSteps property value. The actionSteps property
func (m *Recommendation) GetActionSteps()([]ActionStepable) {
    return m.actionSteps
}
// GetBenefits gets the benefits property value. The benefits property
func (m *Recommendation) GetBenefits()(*string) {
    return m.benefits
}
// GetCategory gets the category property value. The category property
func (m *Recommendation) GetCategory()(*RecommendationCategory) {
    return m.category
}
// GetCreatedDateTime gets the createdDateTime property value. The createdDateTime property
func (m *Recommendation) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDisplayName gets the displayName property value. The displayName property
func (m *Recommendation) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Recommendation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["actionSteps"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateActionStepFromDiscriminatorValue , m.SetActionSteps)
    res["benefits"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetBenefits)
    res["category"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseRecommendationCategory , m.SetCategory)
    res["createdDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetCreatedDateTime)
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["impactedResources"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateRecommendationResourceFromDiscriminatorValue , m.SetImpactedResources)
    res["impactStartDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetImpactStartDateTime)
    res["impactType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetImpactType)
    res["insights"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetInsights)
    res["lastCheckedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetLastCheckedDateTime)
    res["lastModifiedBy"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetLastModifiedBy)
    res["lastModifiedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetLastModifiedDateTime)
    res["postponeUntilDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetPostponeUntilDateTime)
    res["priority"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseRecommendationPriority , m.SetPriority)
    res["status"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseRecommendationStatus , m.SetStatus)
    return res
}
// GetImpactedResources gets the impactedResources property value. The impactedResources property
func (m *Recommendation) GetImpactedResources()([]RecommendationResourceable) {
    return m.impactedResources
}
// GetImpactStartDateTime gets the impactStartDateTime property value. The impactStartDateTime property
func (m *Recommendation) GetImpactStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.impactStartDateTime
}
// GetImpactType gets the impactType property value. The impactType property
func (m *Recommendation) GetImpactType()(*string) {
    return m.impactType
}
// GetInsights gets the insights property value. The insights property
func (m *Recommendation) GetInsights()(*string) {
    return m.insights
}
// GetLastCheckedDateTime gets the lastCheckedDateTime property value. The lastCheckedDateTime property
func (m *Recommendation) GetLastCheckedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastCheckedDateTime
}
// GetLastModifiedBy gets the lastModifiedBy property value. The lastModifiedBy property
func (m *Recommendation) GetLastModifiedBy()(*string) {
    return m.lastModifiedBy
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *Recommendation) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetPostponeUntilDateTime gets the postponeUntilDateTime property value. The postponeUntilDateTime property
func (m *Recommendation) GetPostponeUntilDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.postponeUntilDateTime
}
// GetPriority gets the priority property value. The priority property
func (m *Recommendation) GetPriority()(*RecommendationPriority) {
    return m.priority
}
// GetStatus gets the status property value. The status property
func (m *Recommendation) GetStatus()(*RecommendationStatus) {
    return m.status
}
// Serialize serializes information the current object
func (m *Recommendation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetActionSteps() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetActionSteps())
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
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetImpactedResources())
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
// SetActionSteps sets the actionSteps property value. The actionSteps property
func (m *Recommendation) SetActionSteps(value []ActionStepable)() {
    m.actionSteps = value
}
// SetBenefits sets the benefits property value. The benefits property
func (m *Recommendation) SetBenefits(value *string)() {
    m.benefits = value
}
// SetCategory sets the category property value. The category property
func (m *Recommendation) SetCategory(value *RecommendationCategory)() {
    m.category = value
}
// SetCreatedDateTime sets the createdDateTime property value. The createdDateTime property
func (m *Recommendation) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *Recommendation) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetImpactedResources sets the impactedResources property value. The impactedResources property
func (m *Recommendation) SetImpactedResources(value []RecommendationResourceable)() {
    m.impactedResources = value
}
// SetImpactStartDateTime sets the impactStartDateTime property value. The impactStartDateTime property
func (m *Recommendation) SetImpactStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.impactStartDateTime = value
}
// SetImpactType sets the impactType property value. The impactType property
func (m *Recommendation) SetImpactType(value *string)() {
    m.impactType = value
}
// SetInsights sets the insights property value. The insights property
func (m *Recommendation) SetInsights(value *string)() {
    m.insights = value
}
// SetLastCheckedDateTime sets the lastCheckedDateTime property value. The lastCheckedDateTime property
func (m *Recommendation) SetLastCheckedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastCheckedDateTime = value
}
// SetLastModifiedBy sets the lastModifiedBy property value. The lastModifiedBy property
func (m *Recommendation) SetLastModifiedBy(value *string)() {
    m.lastModifiedBy = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *Recommendation) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetPostponeUntilDateTime sets the postponeUntilDateTime property value. The postponeUntilDateTime property
func (m *Recommendation) SetPostponeUntilDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.postponeUntilDateTime = value
}
// SetPriority sets the priority property value. The priority property
func (m *Recommendation) SetPriority(value *RecommendationPriority)() {
    m.priority = value
}
// SetStatus sets the status property value. The status property
func (m *Recommendation) SetStatus(value *RecommendationStatus)() {
    m.status = value
}

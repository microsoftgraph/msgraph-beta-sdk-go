package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type AccessReview struct {
    Entity
    // The business flow template identifier. Required on create.  This value is case sensitive.
    businessFlowTemplateId *string;
    // The user who created this review.
    createdBy *UserIdentity;
    // The collection of decisions for this access review.
    decisions []AccessReviewDecision;
    // The description provided by the access review creator, to show to the reviewers.
    description *string;
    // The access review name. Required on create.
    displayName *string;
    // The DateTime when the review is scheduled to end. This must be at least one day later than the start date.  Required on create.
    endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The collection of access reviews instances past, present and future, if this object is a recurring access review.
    instances []AccessReview;
    // The collection of decisions for the caller, if the caller is a reviewer.
    myDecisions []AccessReviewDecision;
    // The object for which the access reviews is reviewing the access rights assignments. This can be the group for the review of memberships of users in a group, or the app for a review of assignments of users to an application. Required on create.
    reviewedEntity *Identity;
    // The collection of reviewers for an access review, if access review reviewerType is of type delegated.
    reviewers []AccessReviewReviewer;
    // The relationship type of reviewer to the target object, one of self, delegated or entityOwners. Required on create.
    reviewerType *string;
    // The settings of an accessReview, see type definition below.
    settings *AccessReviewSettings;
    // The DateTime when the review is scheduled to be start.  This could be a date in the future.  Required on create.
    startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // This read-only field specifies the status of an accessReview. The typical states include Initializing, NotStarted, Starting,InProgress, Completing, Completed, AutoReviewing, and AutoReviewed.
    status *string;
}
// Instantiates a new accessReview and sets the default values.
func NewAccessReview()(*AccessReview) {
    m := &AccessReview{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the businessFlowTemplateId property value. The business flow template identifier. Required on create.  This value is case sensitive.
func (m *AccessReview) GetBusinessFlowTemplateId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.businessFlowTemplateId
    }
}
// Gets the createdBy property value. The user who created this review.
func (m *AccessReview) GetCreatedBy()(*UserIdentity) {
    if m == nil {
        return nil
    } else {
        return m.createdBy
    }
}
// Gets the decisions property value. The collection of decisions for this access review.
func (m *AccessReview) GetDecisions()([]AccessReviewDecision) {
    if m == nil {
        return nil
    } else {
        return m.decisions
    }
}
// Gets the description property value. The description provided by the access review creator, to show to the reviewers.
func (m *AccessReview) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the displayName property value. The access review name. Required on create.
func (m *AccessReview) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the endDateTime property value. The DateTime when the review is scheduled to end. This must be at least one day later than the start date.  Required on create.
func (m *AccessReview) GetEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.endDateTime
    }
}
// Gets the instances property value. The collection of access reviews instances past, present and future, if this object is a recurring access review.
func (m *AccessReview) GetInstances()([]AccessReview) {
    if m == nil {
        return nil
    } else {
        return m.instances
    }
}
// Gets the myDecisions property value. The collection of decisions for the caller, if the caller is a reviewer.
func (m *AccessReview) GetMyDecisions()([]AccessReviewDecision) {
    if m == nil {
        return nil
    } else {
        return m.myDecisions
    }
}
// Gets the reviewedEntity property value. The object for which the access reviews is reviewing the access rights assignments. This can be the group for the review of memberships of users in a group, or the app for a review of assignments of users to an application. Required on create.
func (m *AccessReview) GetReviewedEntity()(*Identity) {
    if m == nil {
        return nil
    } else {
        return m.reviewedEntity
    }
}
// Gets the reviewers property value. The collection of reviewers for an access review, if access review reviewerType is of type delegated.
func (m *AccessReview) GetReviewers()([]AccessReviewReviewer) {
    if m == nil {
        return nil
    } else {
        return m.reviewers
    }
}
// Gets the reviewerType property value. The relationship type of reviewer to the target object, one of self, delegated or entityOwners. Required on create.
func (m *AccessReview) GetReviewerType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reviewerType
    }
}
// Gets the settings property value. The settings of an accessReview, see type definition below.
func (m *AccessReview) GetSettings()(*AccessReviewSettings) {
    if m == nil {
        return nil
    } else {
        return m.settings
    }
}
// Gets the startDateTime property value. The DateTime when the review is scheduled to be start.  This could be a date in the future.  Required on create.
func (m *AccessReview) GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.startDateTime
    }
}
// Gets the status property value. This read-only field specifies the status of an accessReview. The typical states include Initializing, NotStarted, Starting,InProgress, Completing, Completed, AutoReviewing, and AutoReviewed.
func (m *AccessReview) GetStatus()(*string) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// The deserialization information for the current model
func (m *AccessReview) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["businessFlowTemplateId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetBusinessFlowTemplateId(val)
        return nil
    }
    res["createdBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserIdentity() })
        if err != nil {
            return err
        }
        m.SetCreatedBy(val.(*UserIdentity))
        return nil
    }
    res["decisions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessReviewDecision() })
        if err != nil {
            return err
        }
        res := make([]AccessReviewDecision, len(val))
        for i, v := range val {
            res[i] = *(v.(*AccessReviewDecision))
        }
        m.SetDecisions(res)
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDescription(val)
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["endDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetEndDateTime(val)
        return nil
    }
    res["instances"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessReview() })
        if err != nil {
            return err
        }
        res := make([]AccessReview, len(val))
        for i, v := range val {
            res[i] = *(v.(*AccessReview))
        }
        m.SetInstances(res)
        return nil
    }
    res["myDecisions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessReviewDecision() })
        if err != nil {
            return err
        }
        res := make([]AccessReviewDecision, len(val))
        for i, v := range val {
            res[i] = *(v.(*AccessReviewDecision))
        }
        m.SetMyDecisions(res)
        return nil
    }
    res["reviewedEntity"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentity() })
        if err != nil {
            return err
        }
        m.SetReviewedEntity(val.(*Identity))
        return nil
    }
    res["reviewers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessReviewReviewer() })
        if err != nil {
            return err
        }
        res := make([]AccessReviewReviewer, len(val))
        for i, v := range val {
            res[i] = *(v.(*AccessReviewReviewer))
        }
        m.SetReviewers(res)
        return nil
    }
    res["reviewerType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetReviewerType(val)
        return nil
    }
    res["settings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessReviewSettings() })
        if err != nil {
            return err
        }
        m.SetSettings(val.(*AccessReviewSettings))
        return nil
    }
    res["startDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetStartDateTime(val)
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetStatus(val)
        return nil
    }
    return res
}
func (m *AccessReview) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *AccessReview) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("businessFlowTemplateId", m.GetBusinessFlowTemplateId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("createdBy", m.GetCreatedBy())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDecisions()))
        for i, v := range m.GetDecisions() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("decisions", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
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
        err = writer.WriteTimeValue("endDateTime", m.GetEndDateTime())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetInstances()))
        for i, v := range m.GetInstances() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("instances", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMyDecisions()))
        for i, v := range m.GetMyDecisions() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("myDecisions", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("reviewedEntity", m.GetReviewedEntity())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetReviewers()))
        for i, v := range m.GetReviewers() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("reviewers", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("reviewerType", m.GetReviewerType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("settings", m.GetSettings())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("startDateTime", m.GetStartDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("status", m.GetStatus())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the businessFlowTemplateId property value. The business flow template identifier. Required on create.  This value is case sensitive.
// Parameters:
//  - value : Value to set for the businessFlowTemplateId property.
func (m *AccessReview) SetBusinessFlowTemplateId(value *string)() {
    m.businessFlowTemplateId = value
}
// Sets the createdBy property value. The user who created this review.
// Parameters:
//  - value : Value to set for the createdBy property.
func (m *AccessReview) SetCreatedBy(value *UserIdentity)() {
    m.createdBy = value
}
// Sets the decisions property value. The collection of decisions for this access review.
// Parameters:
//  - value : Value to set for the decisions property.
func (m *AccessReview) SetDecisions(value []AccessReviewDecision)() {
    m.decisions = value
}
// Sets the description property value. The description provided by the access review creator, to show to the reviewers.
// Parameters:
//  - value : Value to set for the description property.
func (m *AccessReview) SetDescription(value *string)() {
    m.description = value
}
// Sets the displayName property value. The access review name. Required on create.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *AccessReview) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the endDateTime property value. The DateTime when the review is scheduled to end. This must be at least one day later than the start date.  Required on create.
// Parameters:
//  - value : Value to set for the endDateTime property.
func (m *AccessReview) SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.endDateTime = value
}
// Sets the instances property value. The collection of access reviews instances past, present and future, if this object is a recurring access review.
// Parameters:
//  - value : Value to set for the instances property.
func (m *AccessReview) SetInstances(value []AccessReview)() {
    m.instances = value
}
// Sets the myDecisions property value. The collection of decisions for the caller, if the caller is a reviewer.
// Parameters:
//  - value : Value to set for the myDecisions property.
func (m *AccessReview) SetMyDecisions(value []AccessReviewDecision)() {
    m.myDecisions = value
}
// Sets the reviewedEntity property value. The object for which the access reviews is reviewing the access rights assignments. This can be the group for the review of memberships of users in a group, or the app for a review of assignments of users to an application. Required on create.
// Parameters:
//  - value : Value to set for the reviewedEntity property.
func (m *AccessReview) SetReviewedEntity(value *Identity)() {
    m.reviewedEntity = value
}
// Sets the reviewers property value. The collection of reviewers for an access review, if access review reviewerType is of type delegated.
// Parameters:
//  - value : Value to set for the reviewers property.
func (m *AccessReview) SetReviewers(value []AccessReviewReviewer)() {
    m.reviewers = value
}
// Sets the reviewerType property value. The relationship type of reviewer to the target object, one of self, delegated or entityOwners. Required on create.
// Parameters:
//  - value : Value to set for the reviewerType property.
func (m *AccessReview) SetReviewerType(value *string)() {
    m.reviewerType = value
}
// Sets the settings property value. The settings of an accessReview, see type definition below.
// Parameters:
//  - value : Value to set for the settings property.
func (m *AccessReview) SetSettings(value *AccessReviewSettings)() {
    m.settings = value
}
// Sets the startDateTime property value. The DateTime when the review is scheduled to be start.  This could be a date in the future.  Required on create.
// Parameters:
//  - value : Value to set for the startDateTime property.
func (m *AccessReview) SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.startDateTime = value
}
// Sets the status property value. This read-only field specifies the status of an accessReview. The typical states include Initializing, NotStarted, Starting,InProgress, Completing, Completed, AutoReviewing, and AutoReviewed.
// Parameters:
//  - value : Value to set for the status property.
func (m *AccessReview) SetStatus(value *string)() {
    m.status = value
}

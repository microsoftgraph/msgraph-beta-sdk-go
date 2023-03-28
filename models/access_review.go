package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccessReview 
type AccessReview struct {
    Entity
}
// NewAccessReview instantiates a new accessReview and sets the default values.
func NewAccessReview()(*AccessReview) {
    m := &AccessReview{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAccessReviewFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAccessReviewFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAccessReview(), nil
}
// GetBusinessFlowTemplateId gets the businessFlowTemplateId property value. The business flow template identifier. Required on create.  This value is case sensitive.
func (m *AccessReview) GetBusinessFlowTemplateId()(*string) {
    val, err := m.GetBackingStore().Get("businessFlowTemplateId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCreatedBy gets the createdBy property value. The user who created this review.
func (m *AccessReview) GetCreatedBy()(UserIdentityable) {
    val, err := m.GetBackingStore().Get("createdBy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(UserIdentityable)
    }
    return nil
}
// GetDecisions gets the decisions property value. The collection of decisions for this access review.
func (m *AccessReview) GetDecisions()([]AccessReviewDecisionable) {
    val, err := m.GetBackingStore().Get("decisions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AccessReviewDecisionable)
    }
    return nil
}
// GetDescription gets the description property value. The description provided by the access review creator, to show to the reviewers.
func (m *AccessReview) GetDescription()(*string) {
    val, err := m.GetBackingStore().Get("description")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDisplayName gets the displayName property value. The access review name. Required on create.
func (m *AccessReview) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetEndDateTime gets the endDateTime property value. The DateTime when the review is scheduled to end. This must be at least one day later than the start date.  Required on create.
func (m *AccessReview) GetEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("endDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessReview) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["businessFlowTemplateId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBusinessFlowTemplateId(val)
        }
        return nil
    }
    res["createdBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedBy(val.(UserIdentityable))
        }
        return nil
    }
    res["decisions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAccessReviewDecisionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessReviewDecisionable, len(val))
            for i, v := range val {
                res[i] = v.(AccessReviewDecisionable)
            }
            m.SetDecisions(res)
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["endDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndDateTime(val)
        }
        return nil
    }
    res["instances"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAccessReviewFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessReviewable, len(val))
            for i, v := range val {
                res[i] = v.(AccessReviewable)
            }
            m.SetInstances(res)
        }
        return nil
    }
    res["myDecisions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAccessReviewDecisionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessReviewDecisionable, len(val))
            for i, v := range val {
                res[i] = v.(AccessReviewDecisionable)
            }
            m.SetMyDecisions(res)
        }
        return nil
    }
    res["reviewedEntity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReviewedEntity(val.(Identityable))
        }
        return nil
    }
    res["reviewers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAccessReviewReviewerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessReviewReviewerable, len(val))
            for i, v := range val {
                res[i] = v.(AccessReviewReviewerable)
            }
            m.SetReviewers(res)
        }
        return nil
    }
    res["reviewerType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReviewerType(val)
        }
        return nil
    }
    res["settings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAccessReviewSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettings(val.(AccessReviewSettingsable))
        }
        return nil
    }
    res["startDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartDateTime(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val)
        }
        return nil
    }
    return res
}
// GetInstances gets the instances property value. The collection of access reviews instances past, present and future, if this object is a recurring access review.
func (m *AccessReview) GetInstances()([]AccessReviewable) {
    val, err := m.GetBackingStore().Get("instances")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AccessReviewable)
    }
    return nil
}
// GetMyDecisions gets the myDecisions property value. The collection of decisions for the caller, if the caller is a reviewer.
func (m *AccessReview) GetMyDecisions()([]AccessReviewDecisionable) {
    val, err := m.GetBackingStore().Get("myDecisions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AccessReviewDecisionable)
    }
    return nil
}
// GetReviewedEntity gets the reviewedEntity property value. The object for which the access reviews is reviewing the access rights assignments. This can be the group for the review of memberships of users in a group, or the app for a review of assignments of users to an application. Required on create.
func (m *AccessReview) GetReviewedEntity()(Identityable) {
    val, err := m.GetBackingStore().Get("reviewedEntity")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(Identityable)
    }
    return nil
}
// GetReviewers gets the reviewers property value. The collection of reviewers for an access review, if access review reviewerType is of type delegated.
func (m *AccessReview) GetReviewers()([]AccessReviewReviewerable) {
    val, err := m.GetBackingStore().Get("reviewers")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AccessReviewReviewerable)
    }
    return nil
}
// GetReviewerType gets the reviewerType property value. The relationship type of reviewer to the target object, one of self, delegated or entityOwners. Required on create.
func (m *AccessReview) GetReviewerType()(*string) {
    val, err := m.GetBackingStore().Get("reviewerType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSettings gets the settings property value. The settings of an accessReview, see type definition below.
func (m *AccessReview) GetSettings()(AccessReviewSettingsable) {
    val, err := m.GetBackingStore().Get("settings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AccessReviewSettingsable)
    }
    return nil
}
// GetStartDateTime gets the startDateTime property value. The DateTime when the review is scheduled to be start.  This could be a date in the future.  Required on create.
func (m *AccessReview) GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("startDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetStatus gets the status property value. This read-only field specifies the status of an accessReview. The typical states include Initializing, NotStarted, Starting,InProgress, Completing, Completed, AutoReviewing, and AutoReviewed.
func (m *AccessReview) GetStatus()(*string) {
    val, err := m.GetBackingStore().Get("status")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AccessReview) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    if m.GetDecisions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDecisions()))
        for i, v := range m.GetDecisions() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
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
    if m.GetInstances() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetInstances()))
        for i, v := range m.GetInstances() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("instances", cast)
        if err != nil {
            return err
        }
    }
    if m.GetMyDecisions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMyDecisions()))
        for i, v := range m.GetMyDecisions() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
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
    if m.GetReviewers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetReviewers()))
        for i, v := range m.GetReviewers() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
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
// SetBusinessFlowTemplateId sets the businessFlowTemplateId property value. The business flow template identifier. Required on create.  This value is case sensitive.
func (m *AccessReview) SetBusinessFlowTemplateId(value *string)() {
    err := m.GetBackingStore().Set("businessFlowTemplateId", value)
    if err != nil {
        panic(err)
    }
}
// SetCreatedBy sets the createdBy property value. The user who created this review.
func (m *AccessReview) SetCreatedBy(value UserIdentityable)() {
    err := m.GetBackingStore().Set("createdBy", value)
    if err != nil {
        panic(err)
    }
}
// SetDecisions sets the decisions property value. The collection of decisions for this access review.
func (m *AccessReview) SetDecisions(value []AccessReviewDecisionable)() {
    err := m.GetBackingStore().Set("decisions", value)
    if err != nil {
        panic(err)
    }
}
// SetDescription sets the description property value. The description provided by the access review creator, to show to the reviewers.
func (m *AccessReview) SetDescription(value *string)() {
    err := m.GetBackingStore().Set("description", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. The access review name. Required on create.
func (m *AccessReview) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetEndDateTime sets the endDateTime property value. The DateTime when the review is scheduled to end. This must be at least one day later than the start date.  Required on create.
func (m *AccessReview) SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("endDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetInstances sets the instances property value. The collection of access reviews instances past, present and future, if this object is a recurring access review.
func (m *AccessReview) SetInstances(value []AccessReviewable)() {
    err := m.GetBackingStore().Set("instances", value)
    if err != nil {
        panic(err)
    }
}
// SetMyDecisions sets the myDecisions property value. The collection of decisions for the caller, if the caller is a reviewer.
func (m *AccessReview) SetMyDecisions(value []AccessReviewDecisionable)() {
    err := m.GetBackingStore().Set("myDecisions", value)
    if err != nil {
        panic(err)
    }
}
// SetReviewedEntity sets the reviewedEntity property value. The object for which the access reviews is reviewing the access rights assignments. This can be the group for the review of memberships of users in a group, or the app for a review of assignments of users to an application. Required on create.
func (m *AccessReview) SetReviewedEntity(value Identityable)() {
    err := m.GetBackingStore().Set("reviewedEntity", value)
    if err != nil {
        panic(err)
    }
}
// SetReviewers sets the reviewers property value. The collection of reviewers for an access review, if access review reviewerType is of type delegated.
func (m *AccessReview) SetReviewers(value []AccessReviewReviewerable)() {
    err := m.GetBackingStore().Set("reviewers", value)
    if err != nil {
        panic(err)
    }
}
// SetReviewerType sets the reviewerType property value. The relationship type of reviewer to the target object, one of self, delegated or entityOwners. Required on create.
func (m *AccessReview) SetReviewerType(value *string)() {
    err := m.GetBackingStore().Set("reviewerType", value)
    if err != nil {
        panic(err)
    }
}
// SetSettings sets the settings property value. The settings of an accessReview, see type definition below.
func (m *AccessReview) SetSettings(value AccessReviewSettingsable)() {
    err := m.GetBackingStore().Set("settings", value)
    if err != nil {
        panic(err)
    }
}
// SetStartDateTime sets the startDateTime property value. The DateTime when the review is scheduled to be start.  This could be a date in the future.  Required on create.
func (m *AccessReview) SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("startDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetStatus sets the status property value. This read-only field specifies the status of an accessReview. The typical states include Initializing, NotStarted, Starting,InProgress, Completing, Completed, AutoReviewing, and AutoReviewed.
func (m *AccessReview) SetStatus(value *string)() {
    err := m.GetBackingStore().Set("status", value)
    if err != nil {
        panic(err)
    }
}
// AccessReviewable 
type AccessReviewable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBusinessFlowTemplateId()(*string)
    GetCreatedBy()(UserIdentityable)
    GetDecisions()([]AccessReviewDecisionable)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetInstances()([]AccessReviewable)
    GetMyDecisions()([]AccessReviewDecisionable)
    GetReviewedEntity()(Identityable)
    GetReviewers()([]AccessReviewReviewerable)
    GetReviewerType()(*string)
    GetSettings()(AccessReviewSettingsable)
    GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetStatus()(*string)
    SetBusinessFlowTemplateId(value *string)()
    SetCreatedBy(value UserIdentityable)()
    SetDecisions(value []AccessReviewDecisionable)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetInstances(value []AccessReviewable)()
    SetMyDecisions(value []AccessReviewDecisionable)()
    SetReviewedEntity(value Identityable)()
    SetReviewers(value []AccessReviewReviewerable)()
    SetReviewerType(value *string)()
    SetSettings(value AccessReviewSettingsable)()
    SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetStatus(value *string)()
}

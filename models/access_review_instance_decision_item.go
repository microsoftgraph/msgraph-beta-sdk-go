package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccessReviewInstanceDecisionItem provides operations to manage the collection of accessReviewDecision entities.
type AccessReviewInstanceDecisionItem struct {
    Entity
    // The identifier of the accessReviewInstance parent. Supports $select. Read-only.
    accessReviewId *string
    // The identifier of the user who applied the decision. 00000000-0000-0000-0000-000000000000 if the assigned reviewer hasn't applied the decision or it was automatically applied. Read-only.
    appliedBy UserIdentityable
    // The timestamp when the approval decision was applied. The DatetimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.  Supports $select. Read-only.
    appliedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The result of applying the decision. Possible values: New, AppliedSuccessfully, AppliedWithUnknownFailure, AppliedSuccessfullyButObjectNotFound and ApplyNotSupported. Supports $select, $orderby, and $filter (eq only). Read-only.
    applyResult *string
    // Result of the review. Possible values: Approve, Deny, NotReviewed, or DontKnow. Supports $select, $orderby, and $filter (eq only).
    decision *string
    // Insights are recommendations to reviewers on whether to approve or deny a decision. There can be multiple insights associated with an accessReviewInstanceDecisionItem.
    insights []GovernanceInsightable
    // There is exactly one accessReviewInstance associated with each decision. The instance is the parent of the decision item, representing the recurrence of the access review the decision is made on.
    instance AccessReviewInstanceable
    // Justification left by the reviewer when they made the decision.
    justification *string
    // Every decision item in an access review represents a principal's access to a resource. This property represents details of the principal. For example, if a decision item represents access of User 'Bob' to Group 'Sales' - The principal is 'Bob' and the resource is 'Sales'. Principals can be of two types - userIdentity and servicePrincipalIdentity. Supports $select. Read-only.
    principal Identityable
    // Link to the principal object. For example: https://graph.microsoft.com/v1.0/users/a6c7aecb-cbfd-4763-87ef-e91b4bd509d9. Read-only.
    principalLink *string
    // Every decision item in an access review represents a principal's membership to a resource. This property provides the details of the membership. For example, whether the principal has direct access or indirect access to the resource. Supports $select. Read-only.
    principalResourceMembership DecisionItemPrincipalResourceMembershipable
    // A system-generated recommendation for the approval decision based off last interactive sign-in to tenant. Recommend approve if sign-in is within thirty days of start of review. Recommend deny if sign-in is greater than thirty days of start of review. Recommendation not available otherwise. Possible values: Approve, Deny, or NoInfoAvailable. Supports $select, $orderby, and $filter (eq only). Read-only.
    recommendation *string
    // Every decision item in an access review represents a principal's access to a resource. This property represents details of the resource. For example, if a decision item represents access of User 'Bob' to Group 'Sales' - The principal is Bob and the resource is 'Sales'. Resources can be of multiple types. See accessReviewInstanceDecisionItemResource. Read-only.
    resource AccessReviewInstanceDecisionItemResourceable
    // A link to the resource. For example, https://graph.microsoft.com/v1.0/servicePrincipals/c86300f3-8695-4320-9f6e-32a2555f5ff8. Supports $select. Read-only.
    resourceLink *string
    // The identifier of the reviewer. 00000000-0000-0000-0000-000000000000 if the assigned reviewer hasn't reviewed. Supports $select. Read-only.
    reviewedBy UserIdentityable
    // The timestamp when the review decision occurred. Supports $select. Read-only.
    reviewedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The target of this specific decision. Decision targets can be of different types – each one with its own specific properties. See accessReviewInstanceDecisionItemTarget. Read-only.  This property has been replaced by the principal and resource properties in v1.0.
    target AccessReviewInstanceDecisionItemTargetable
}
// NewAccessReviewInstanceDecisionItem instantiates a new accessReviewInstanceDecisionItem and sets the default values.
func NewAccessReviewInstanceDecisionItem()(*AccessReviewInstanceDecisionItem) {
    m := &AccessReviewInstanceDecisionItem{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.accessReviewInstanceDecisionItem";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAccessReviewInstanceDecisionItemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAccessReviewInstanceDecisionItemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAccessReviewInstanceDecisionItem(), nil
}
// GetAccessReviewId gets the accessReviewId property value. The identifier of the accessReviewInstance parent. Supports $select. Read-only.
func (m *AccessReviewInstanceDecisionItem) GetAccessReviewId()(*string) {
    return m.accessReviewId
}
// GetAppliedBy gets the appliedBy property value. The identifier of the user who applied the decision. 00000000-0000-0000-0000-000000000000 if the assigned reviewer hasn't applied the decision or it was automatically applied. Read-only.
func (m *AccessReviewInstanceDecisionItem) GetAppliedBy()(UserIdentityable) {
    return m.appliedBy
}
// GetAppliedDateTime gets the appliedDateTime property value. The timestamp when the approval decision was applied. The DatetimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.  Supports $select. Read-only.
func (m *AccessReviewInstanceDecisionItem) GetAppliedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.appliedDateTime
}
// GetApplyResult gets the applyResult property value. The result of applying the decision. Possible values: New, AppliedSuccessfully, AppliedWithUnknownFailure, AppliedSuccessfullyButObjectNotFound and ApplyNotSupported. Supports $select, $orderby, and $filter (eq only). Read-only.
func (m *AccessReviewInstanceDecisionItem) GetApplyResult()(*string) {
    return m.applyResult
}
// GetDecision gets the decision property value. Result of the review. Possible values: Approve, Deny, NotReviewed, or DontKnow. Supports $select, $orderby, and $filter (eq only).
func (m *AccessReviewInstanceDecisionItem) GetDecision()(*string) {
    return m.decision
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessReviewInstanceDecisionItem) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["accessReviewId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetAccessReviewId)
    res["appliedBy"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateUserIdentityFromDiscriminatorValue , m.SetAppliedBy)
    res["appliedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetAppliedDateTime)
    res["applyResult"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetApplyResult)
    res["decision"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDecision)
    res["insights"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateGovernanceInsightFromDiscriminatorValue , m.SetInsights)
    res["instance"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateAccessReviewInstanceFromDiscriminatorValue , m.SetInstance)
    res["justification"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetJustification)
    res["principal"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateIdentityFromDiscriminatorValue , m.SetPrincipal)
    res["principalLink"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetPrincipalLink)
    res["principalResourceMembership"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateDecisionItemPrincipalResourceMembershipFromDiscriminatorValue , m.SetPrincipalResourceMembership)
    res["recommendation"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetRecommendation)
    res["resource"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateAccessReviewInstanceDecisionItemResourceFromDiscriminatorValue , m.SetResource)
    res["resourceLink"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetResourceLink)
    res["reviewedBy"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateUserIdentityFromDiscriminatorValue , m.SetReviewedBy)
    res["reviewedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetReviewedDateTime)
    res["target"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateAccessReviewInstanceDecisionItemTargetFromDiscriminatorValue , m.SetTarget)
    return res
}
// GetInsights gets the insights property value. Insights are recommendations to reviewers on whether to approve or deny a decision. There can be multiple insights associated with an accessReviewInstanceDecisionItem.
func (m *AccessReviewInstanceDecisionItem) GetInsights()([]GovernanceInsightable) {
    return m.insights
}
// GetInstance gets the instance property value. There is exactly one accessReviewInstance associated with each decision. The instance is the parent of the decision item, representing the recurrence of the access review the decision is made on.
func (m *AccessReviewInstanceDecisionItem) GetInstance()(AccessReviewInstanceable) {
    return m.instance
}
// GetJustification gets the justification property value. Justification left by the reviewer when they made the decision.
func (m *AccessReviewInstanceDecisionItem) GetJustification()(*string) {
    return m.justification
}
// GetPrincipal gets the principal property value. Every decision item in an access review represents a principal's access to a resource. This property represents details of the principal. For example, if a decision item represents access of User 'Bob' to Group 'Sales' - The principal is 'Bob' and the resource is 'Sales'. Principals can be of two types - userIdentity and servicePrincipalIdentity. Supports $select. Read-only.
func (m *AccessReviewInstanceDecisionItem) GetPrincipal()(Identityable) {
    return m.principal
}
// GetPrincipalLink gets the principalLink property value. Link to the principal object. For example: https://graph.microsoft.com/v1.0/users/a6c7aecb-cbfd-4763-87ef-e91b4bd509d9. Read-only.
func (m *AccessReviewInstanceDecisionItem) GetPrincipalLink()(*string) {
    return m.principalLink
}
// GetPrincipalResourceMembership gets the principalResourceMembership property value. Every decision item in an access review represents a principal's membership to a resource. This property provides the details of the membership. For example, whether the principal has direct access or indirect access to the resource. Supports $select. Read-only.
func (m *AccessReviewInstanceDecisionItem) GetPrincipalResourceMembership()(DecisionItemPrincipalResourceMembershipable) {
    return m.principalResourceMembership
}
// GetRecommendation gets the recommendation property value. A system-generated recommendation for the approval decision based off last interactive sign-in to tenant. Recommend approve if sign-in is within thirty days of start of review. Recommend deny if sign-in is greater than thirty days of start of review. Recommendation not available otherwise. Possible values: Approve, Deny, or NoInfoAvailable. Supports $select, $orderby, and $filter (eq only). Read-only.
func (m *AccessReviewInstanceDecisionItem) GetRecommendation()(*string) {
    return m.recommendation
}
// GetResource gets the resource property value. Every decision item in an access review represents a principal's access to a resource. This property represents details of the resource. For example, if a decision item represents access of User 'Bob' to Group 'Sales' - The principal is Bob and the resource is 'Sales'. Resources can be of multiple types. See accessReviewInstanceDecisionItemResource. Read-only.
func (m *AccessReviewInstanceDecisionItem) GetResource()(AccessReviewInstanceDecisionItemResourceable) {
    return m.resource
}
// GetResourceLink gets the resourceLink property value. A link to the resource. For example, https://graph.microsoft.com/v1.0/servicePrincipals/c86300f3-8695-4320-9f6e-32a2555f5ff8. Supports $select. Read-only.
func (m *AccessReviewInstanceDecisionItem) GetResourceLink()(*string) {
    return m.resourceLink
}
// GetReviewedBy gets the reviewedBy property value. The identifier of the reviewer. 00000000-0000-0000-0000-000000000000 if the assigned reviewer hasn't reviewed. Supports $select. Read-only.
func (m *AccessReviewInstanceDecisionItem) GetReviewedBy()(UserIdentityable) {
    return m.reviewedBy
}
// GetReviewedDateTime gets the reviewedDateTime property value. The timestamp when the review decision occurred. Supports $select. Read-only.
func (m *AccessReviewInstanceDecisionItem) GetReviewedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.reviewedDateTime
}
// GetTarget gets the target property value. The target of this specific decision. Decision targets can be of different types – each one with its own specific properties. See accessReviewInstanceDecisionItemTarget. Read-only.  This property has been replaced by the principal and resource properties in v1.0.
func (m *AccessReviewInstanceDecisionItem) GetTarget()(AccessReviewInstanceDecisionItemTargetable) {
    return m.target
}
// Serialize serializes information the current object
func (m *AccessReviewInstanceDecisionItem) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("accessReviewId", m.GetAccessReviewId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("appliedBy", m.GetAppliedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("appliedDateTime", m.GetAppliedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("applyResult", m.GetApplyResult())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("decision", m.GetDecision())
        if err != nil {
            return err
        }
    }
    if m.GetInsights() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetInsights())
        err = writer.WriteCollectionOfObjectValues("insights", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("instance", m.GetInstance())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("justification", m.GetJustification())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("principal", m.GetPrincipal())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("principalLink", m.GetPrincipalLink())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("principalResourceMembership", m.GetPrincipalResourceMembership())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("recommendation", m.GetRecommendation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("resource", m.GetResource())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("resourceLink", m.GetResourceLink())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("reviewedBy", m.GetReviewedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("reviewedDateTime", m.GetReviewedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("target", m.GetTarget())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccessReviewId sets the accessReviewId property value. The identifier of the accessReviewInstance parent. Supports $select. Read-only.
func (m *AccessReviewInstanceDecisionItem) SetAccessReviewId(value *string)() {
    m.accessReviewId = value
}
// SetAppliedBy sets the appliedBy property value. The identifier of the user who applied the decision. 00000000-0000-0000-0000-000000000000 if the assigned reviewer hasn't applied the decision or it was automatically applied. Read-only.
func (m *AccessReviewInstanceDecisionItem) SetAppliedBy(value UserIdentityable)() {
    m.appliedBy = value
}
// SetAppliedDateTime sets the appliedDateTime property value. The timestamp when the approval decision was applied. The DatetimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.  Supports $select. Read-only.
func (m *AccessReviewInstanceDecisionItem) SetAppliedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.appliedDateTime = value
}
// SetApplyResult sets the applyResult property value. The result of applying the decision. Possible values: New, AppliedSuccessfully, AppliedWithUnknownFailure, AppliedSuccessfullyButObjectNotFound and ApplyNotSupported. Supports $select, $orderby, and $filter (eq only). Read-only.
func (m *AccessReviewInstanceDecisionItem) SetApplyResult(value *string)() {
    m.applyResult = value
}
// SetDecision sets the decision property value. Result of the review. Possible values: Approve, Deny, NotReviewed, or DontKnow. Supports $select, $orderby, and $filter (eq only).
func (m *AccessReviewInstanceDecisionItem) SetDecision(value *string)() {
    m.decision = value
}
// SetInsights sets the insights property value. Insights are recommendations to reviewers on whether to approve or deny a decision. There can be multiple insights associated with an accessReviewInstanceDecisionItem.
func (m *AccessReviewInstanceDecisionItem) SetInsights(value []GovernanceInsightable)() {
    m.insights = value
}
// SetInstance sets the instance property value. There is exactly one accessReviewInstance associated with each decision. The instance is the parent of the decision item, representing the recurrence of the access review the decision is made on.
func (m *AccessReviewInstanceDecisionItem) SetInstance(value AccessReviewInstanceable)() {
    m.instance = value
}
// SetJustification sets the justification property value. Justification left by the reviewer when they made the decision.
func (m *AccessReviewInstanceDecisionItem) SetJustification(value *string)() {
    m.justification = value
}
// SetPrincipal sets the principal property value. Every decision item in an access review represents a principal's access to a resource. This property represents details of the principal. For example, if a decision item represents access of User 'Bob' to Group 'Sales' - The principal is 'Bob' and the resource is 'Sales'. Principals can be of two types - userIdentity and servicePrincipalIdentity. Supports $select. Read-only.
func (m *AccessReviewInstanceDecisionItem) SetPrincipal(value Identityable)() {
    m.principal = value
}
// SetPrincipalLink sets the principalLink property value. Link to the principal object. For example: https://graph.microsoft.com/v1.0/users/a6c7aecb-cbfd-4763-87ef-e91b4bd509d9. Read-only.
func (m *AccessReviewInstanceDecisionItem) SetPrincipalLink(value *string)() {
    m.principalLink = value
}
// SetPrincipalResourceMembership sets the principalResourceMembership property value. Every decision item in an access review represents a principal's membership to a resource. This property provides the details of the membership. For example, whether the principal has direct access or indirect access to the resource. Supports $select. Read-only.
func (m *AccessReviewInstanceDecisionItem) SetPrincipalResourceMembership(value DecisionItemPrincipalResourceMembershipable)() {
    m.principalResourceMembership = value
}
// SetRecommendation sets the recommendation property value. A system-generated recommendation for the approval decision based off last interactive sign-in to tenant. Recommend approve if sign-in is within thirty days of start of review. Recommend deny if sign-in is greater than thirty days of start of review. Recommendation not available otherwise. Possible values: Approve, Deny, or NoInfoAvailable. Supports $select, $orderby, and $filter (eq only). Read-only.
func (m *AccessReviewInstanceDecisionItem) SetRecommendation(value *string)() {
    m.recommendation = value
}
// SetResource sets the resource property value. Every decision item in an access review represents a principal's access to a resource. This property represents details of the resource. For example, if a decision item represents access of User 'Bob' to Group 'Sales' - The principal is Bob and the resource is 'Sales'. Resources can be of multiple types. See accessReviewInstanceDecisionItemResource. Read-only.
func (m *AccessReviewInstanceDecisionItem) SetResource(value AccessReviewInstanceDecisionItemResourceable)() {
    m.resource = value
}
// SetResourceLink sets the resourceLink property value. A link to the resource. For example, https://graph.microsoft.com/v1.0/servicePrincipals/c86300f3-8695-4320-9f6e-32a2555f5ff8. Supports $select. Read-only.
func (m *AccessReviewInstanceDecisionItem) SetResourceLink(value *string)() {
    m.resourceLink = value
}
// SetReviewedBy sets the reviewedBy property value. The identifier of the reviewer. 00000000-0000-0000-0000-000000000000 if the assigned reviewer hasn't reviewed. Supports $select. Read-only.
func (m *AccessReviewInstanceDecisionItem) SetReviewedBy(value UserIdentityable)() {
    m.reviewedBy = value
}
// SetReviewedDateTime sets the reviewedDateTime property value. The timestamp when the review decision occurred. Supports $select. Read-only.
func (m *AccessReviewInstanceDecisionItem) SetReviewedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.reviewedDateTime = value
}
// SetTarget sets the target property value. The target of this specific decision. Decision targets can be of different types – each one with its own specific properties. See accessReviewInstanceDecisionItemTarget. Read-only.  This property has been replaced by the principal and resource properties in v1.0.
func (m *AccessReviewInstanceDecisionItem) SetTarget(value AccessReviewInstanceDecisionItemTargetable)() {
    m.target = value
}

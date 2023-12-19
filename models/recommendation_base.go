package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RecommendationBase 
type RecommendationBase struct {
    Entity
}
// NewRecommendationBase instantiates a new recommendationBase and sets the default values.
func NewRecommendationBase()(*RecommendationBase) {
    m := &RecommendationBase{
        Entity: *NewEntity(),
    }
    return m
}
// CreateRecommendationBaseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRecommendationBaseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.recommendation":
                        return NewRecommendation(), nil
                }
            }
        }
    }
    return NewRecommendationBase(), nil
}
// GetActionSteps gets the actionSteps property value. List of actions to take to complete a recommendation.
func (m *RecommendationBase) GetActionSteps()([]ActionStepable) {
    val, err := m.GetBackingStore().Get("actionSteps")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ActionStepable)
    }
    return nil
}
// GetBenefits gets the benefits property value. An explanation of why completing the recommendation will benefit you. Corresponds to the Value section of a recommendation shown in the Microsoft Entra admin center.
func (m *RecommendationBase) GetBenefits()(*string) {
    val, err := m.GetBackingStore().Get("benefits")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCategory gets the category property value. The category property
func (m *RecommendationBase) GetCategory()(*RecommendationCategory) {
    val, err := m.GetBackingStore().Get("category")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*RecommendationCategory)
    }
    return nil
}
// GetCreatedDateTime gets the createdDateTime property value. The date and time when the recommendation was detected as applicable to your directory.
func (m *RecommendationBase) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("createdDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetCurrentScore gets the currentScore property value. The number of points the tenant has attained. Only applies to recommendations with category set to identitySecureScore.
func (m *RecommendationBase) GetCurrentScore()(*float64) {
    val, err := m.GetBackingStore().Get("currentScore")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float64)
    }
    return nil
}
// GetDisplayName gets the displayName property value. The title of the recommendation.
func (m *RecommendationBase) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFeatureAreas gets the featureAreas property value. The directory feature that the recommendation is related to.
func (m *RecommendationBase) GetFeatureAreas()([]RecommendationBase_featureAreas) {
    val, err := m.GetBackingStore().Get("featureAreas")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]RecommendationBase_featureAreas)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RecommendationBase) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["actionSteps"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateActionStepFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ActionStepable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ActionStepable)
                }
            }
            m.SetActionSteps(res)
        }
        return nil
    }
    res["benefits"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBenefits(val)
        }
        return nil
    }
    res["category"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRecommendationCategory)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory(val.(*RecommendationCategory))
        }
        return nil
    }
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["currentScore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCurrentScore(val)
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
    res["featureAreas"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseRecommendationBase_featureAreas)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RecommendationBase_featureAreas, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*RecommendationBase_featureAreas))
                }
            }
            m.SetFeatureAreas(res)
        }
        return nil
    }
    res["impactedResources"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateImpactedResourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ImpactedResourceable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ImpactedResourceable)
                }
            }
            m.SetImpactedResources(res)
        }
        return nil
    }
    res["impactStartDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetImpactStartDateTime(val)
        }
        return nil
    }
    res["impactType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetImpactType(val)
        }
        return nil
    }
    res["insights"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInsights(val)
        }
        return nil
    }
    res["lastCheckedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastCheckedDateTime(val)
        }
        return nil
    }
    res["lastModifiedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedBy(val)
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["maxScore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxScore(val)
        }
        return nil
    }
    res["postponeUntilDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPostponeUntilDateTime(val)
        }
        return nil
    }
    res["priority"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRecommendationPriority)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPriority(val.(*RecommendationPriority))
        }
        return nil
    }
    res["recommendationType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRecommendationBase_recommendationType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecommendationType(val.(*RecommendationBase_recommendationType))
        }
        return nil
    }
    res["releaseType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReleaseType(val)
        }
        return nil
    }
    res["remediationImpact"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemediationImpact(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
// GetImpactedResources gets the impactedResources property value. The list of directory objects associated with the recommendation.
func (m *RecommendationBase) GetImpactedResources()([]ImpactedResourceable) {
    val, err := m.GetBackingStore().Get("impactedResources")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ImpactedResourceable)
    }
    return nil
}
// GetImpactStartDateTime gets the impactStartDateTime property value. The future date and time when a recommendation should be completed.
func (m *RecommendationBase) GetImpactStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("impactStartDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetImpactType gets the impactType property value. Indicates the scope of impact of a recommendation. Tenant level indicates that the recommendation impacts the whole tenant. Other possible values include users, applications.
func (m *RecommendationBase) GetImpactType()(*string) {
    val, err := m.GetBackingStore().Get("impactType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetInsights gets the insights property value. Describes why a recommendation uniquely applies to your directory. Corresponds to the Description section of a recommendation shown in the Microsoft Entra admin center.
func (m *RecommendationBase) GetInsights()(*string) {
    val, err := m.GetBackingStore().Get("insights")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetLastCheckedDateTime gets the lastCheckedDateTime property value. The most recent date and time a recommendation was deemed applicable to your directory.
func (m *RecommendationBase) GetLastCheckedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastCheckedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetLastModifiedBy gets the lastModifiedBy property value. Name of the user who last updated the status of the recommendation.
func (m *RecommendationBase) GetLastModifiedBy()(*string) {
    val, err := m.GetBackingStore().Get("lastModifiedBy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The date and time the status of a recommendation was last updated.
func (m *RecommendationBase) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastModifiedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetMaxScore gets the maxScore property value. The maximum number of points attainable. Only applies to recommendations with category set to identitySecureScore.
func (m *RecommendationBase) GetMaxScore()(*float64) {
    val, err := m.GetBackingStore().Get("maxScore")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float64)
    }
    return nil
}
// GetPostponeUntilDateTime gets the postponeUntilDateTime property value. The future date and time when the status of a postponed recommendation will be active again.
func (m *RecommendationBase) GetPostponeUntilDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("postponeUntilDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetPriority gets the priority property value. The priority property
func (m *RecommendationBase) GetPriority()(*RecommendationPriority) {
    val, err := m.GetBackingStore().Get("priority")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*RecommendationPriority)
    }
    return nil
}
// GetRecommendationType gets the recommendationType property value. Friendly shortname to identify the recommendation. The possible values are: adfsAppsMigration, enableDesktopSSO, enablePHS, enableProvisioning, switchFromPerUserMFA, tenantMFA, thirdPartyApps, turnOffPerUserMFA, useAuthenticatorApp, useMyApps, staleApps, staleAppCreds, applicationCredentialExpiry, servicePrincipalKeyExpiry, adminMFAV2, blockLegacyAuthentication, integratedApps, mfaRegistrationV2, pwagePolicyNew, passwordHashSync, oneAdmin, roleOverlap, selfServicePasswordReset, signinRiskPolicy, userRiskPolicy, verifyAppPublisher, privateLinkForAAD, appRoleAssignmentsGroups, appRoleAssignmentsUsers, managedIdentity, overprivilegedApps, unknownFutureValue, longLivedCredentials, aadConnectDeprecated, adalToMsalMigration, ownerlessApps, inactiveGuests. Also, please note that you must use the Prefer: include-unknown-enum-members request header to get the following value(s) in this evolvable enum: longLivedCredentials, aadConnectDeprecated, adalToMsalMigration, ownerlessApps, inactiveGuests.
func (m *RecommendationBase) GetRecommendationType()(*RecommendationBase_recommendationType) {
    val, err := m.GetBackingStore().Get("recommendationType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*RecommendationBase_recommendationType)
    }
    return nil
}
// GetReleaseType gets the releaseType property value. The current release type of the recommendation. The possible values are: preview, generallyAvailable, unknownFutureValue.
func (m *RecommendationBase) GetReleaseType()(*string) {
    val, err := m.GetBackingStore().Get("releaseType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRemediationImpact gets the remediationImpact property value. Description of the impact on users of the remediation. Only applies to recommendations with category set to identitySecureScore.
func (m *RecommendationBase) GetRemediationImpact()(*string) {
    val, err := m.GetBackingStore().Get("remediationImpact")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetStatus gets the status property value. The status property
func (m *RecommendationBase) GetStatus()(*RecommendationStatus) {
    val, err := m.GetBackingStore().Get("status")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*RecommendationStatus)
    }
    return nil
}
// Serialize serializes information the current object
func (m *RecommendationBase) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetActionSteps() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetActionSteps()))
        for i, v := range m.GetActionSteps() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
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
        err = writer.WriteFloat64Value("currentScore", m.GetCurrentScore())
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
    if m.GetFeatureAreas() != nil {
        err = writer.WriteCollectionOfStringValues("featureAreas", SerializeRecommendationBase_featureAreas(m.GetFeatureAreas()))
        if err != nil {
            return err
        }
    }
    if m.GetImpactedResources() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetImpactedResources()))
        for i, v := range m.GetImpactedResources() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
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
        err = writer.WriteFloat64Value("maxScore", m.GetMaxScore())
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
    if m.GetRecommendationType() != nil {
        cast := (*m.GetRecommendationType()).String()
        err = writer.WriteStringValue("recommendationType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("releaseType", m.GetReleaseType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("remediationImpact", m.GetRemediationImpact())
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
// SetActionSteps sets the actionSteps property value. List of actions to take to complete a recommendation.
func (m *RecommendationBase) SetActionSteps(value []ActionStepable)() {
    err := m.GetBackingStore().Set("actionSteps", value)
    if err != nil {
        panic(err)
    }
}
// SetBenefits sets the benefits property value. An explanation of why completing the recommendation will benefit you. Corresponds to the Value section of a recommendation shown in the Microsoft Entra admin center.
func (m *RecommendationBase) SetBenefits(value *string)() {
    err := m.GetBackingStore().Set("benefits", value)
    if err != nil {
        panic(err)
    }
}
// SetCategory sets the category property value. The category property
func (m *RecommendationBase) SetCategory(value *RecommendationCategory)() {
    err := m.GetBackingStore().Set("category", value)
    if err != nil {
        panic(err)
    }
}
// SetCreatedDateTime sets the createdDateTime property value. The date and time when the recommendation was detected as applicable to your directory.
func (m *RecommendationBase) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("createdDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetCurrentScore sets the currentScore property value. The number of points the tenant has attained. Only applies to recommendations with category set to identitySecureScore.
func (m *RecommendationBase) SetCurrentScore(value *float64)() {
    err := m.GetBackingStore().Set("currentScore", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. The title of the recommendation.
func (m *RecommendationBase) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetFeatureAreas sets the featureAreas property value. The directory feature that the recommendation is related to.
func (m *RecommendationBase) SetFeatureAreas(value []RecommendationBase_featureAreas)() {
    err := m.GetBackingStore().Set("featureAreas", value)
    if err != nil {
        panic(err)
    }
}
// SetImpactedResources sets the impactedResources property value. The list of directory objects associated with the recommendation.
func (m *RecommendationBase) SetImpactedResources(value []ImpactedResourceable)() {
    err := m.GetBackingStore().Set("impactedResources", value)
    if err != nil {
        panic(err)
    }
}
// SetImpactStartDateTime sets the impactStartDateTime property value. The future date and time when a recommendation should be completed.
func (m *RecommendationBase) SetImpactStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("impactStartDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetImpactType sets the impactType property value. Indicates the scope of impact of a recommendation. Tenant level indicates that the recommendation impacts the whole tenant. Other possible values include users, applications.
func (m *RecommendationBase) SetImpactType(value *string)() {
    err := m.GetBackingStore().Set("impactType", value)
    if err != nil {
        panic(err)
    }
}
// SetInsights sets the insights property value. Describes why a recommendation uniquely applies to your directory. Corresponds to the Description section of a recommendation shown in the Microsoft Entra admin center.
func (m *RecommendationBase) SetInsights(value *string)() {
    err := m.GetBackingStore().Set("insights", value)
    if err != nil {
        panic(err)
    }
}
// SetLastCheckedDateTime sets the lastCheckedDateTime property value. The most recent date and time a recommendation was deemed applicable to your directory.
func (m *RecommendationBase) SetLastCheckedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastCheckedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetLastModifiedBy sets the lastModifiedBy property value. Name of the user who last updated the status of the recommendation.
func (m *RecommendationBase) SetLastModifiedBy(value *string)() {
    err := m.GetBackingStore().Set("lastModifiedBy", value)
    if err != nil {
        panic(err)
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The date and time the status of a recommendation was last updated.
func (m *RecommendationBase) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastModifiedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetMaxScore sets the maxScore property value. The maximum number of points attainable. Only applies to recommendations with category set to identitySecureScore.
func (m *RecommendationBase) SetMaxScore(value *float64)() {
    err := m.GetBackingStore().Set("maxScore", value)
    if err != nil {
        panic(err)
    }
}
// SetPostponeUntilDateTime sets the postponeUntilDateTime property value. The future date and time when the status of a postponed recommendation will be active again.
func (m *RecommendationBase) SetPostponeUntilDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("postponeUntilDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetPriority sets the priority property value. The priority property
func (m *RecommendationBase) SetPriority(value *RecommendationPriority)() {
    err := m.GetBackingStore().Set("priority", value)
    if err != nil {
        panic(err)
    }
}
// SetRecommendationType sets the recommendationType property value. Friendly shortname to identify the recommendation. The possible values are: adfsAppsMigration, enableDesktopSSO, enablePHS, enableProvisioning, switchFromPerUserMFA, tenantMFA, thirdPartyApps, turnOffPerUserMFA, useAuthenticatorApp, useMyApps, staleApps, staleAppCreds, applicationCredentialExpiry, servicePrincipalKeyExpiry, adminMFAV2, blockLegacyAuthentication, integratedApps, mfaRegistrationV2, pwagePolicyNew, passwordHashSync, oneAdmin, roleOverlap, selfServicePasswordReset, signinRiskPolicy, userRiskPolicy, verifyAppPublisher, privateLinkForAAD, appRoleAssignmentsGroups, appRoleAssignmentsUsers, managedIdentity, overprivilegedApps, unknownFutureValue, longLivedCredentials, aadConnectDeprecated, adalToMsalMigration, ownerlessApps, inactiveGuests. Also, please note that you must use the Prefer: include-unknown-enum-members request header to get the following value(s) in this evolvable enum: longLivedCredentials, aadConnectDeprecated, adalToMsalMigration, ownerlessApps, inactiveGuests.
func (m *RecommendationBase) SetRecommendationType(value *RecommendationBase_recommendationType)() {
    err := m.GetBackingStore().Set("recommendationType", value)
    if err != nil {
        panic(err)
    }
}
// SetReleaseType sets the releaseType property value. The current release type of the recommendation. The possible values are: preview, generallyAvailable, unknownFutureValue.
func (m *RecommendationBase) SetReleaseType(value *string)() {
    err := m.GetBackingStore().Set("releaseType", value)
    if err != nil {
        panic(err)
    }
}
// SetRemediationImpact sets the remediationImpact property value. Description of the impact on users of the remediation. Only applies to recommendations with category set to identitySecureScore.
func (m *RecommendationBase) SetRemediationImpact(value *string)() {
    err := m.GetBackingStore().Set("remediationImpact", value)
    if err != nil {
        panic(err)
    }
}
// SetStatus sets the status property value. The status property
func (m *RecommendationBase) SetStatus(value *RecommendationStatus)() {
    err := m.GetBackingStore().Set("status", value)
    if err != nil {
        panic(err)
    }
}
// RecommendationBaseable 
type RecommendationBaseable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActionSteps()([]ActionStepable)
    GetBenefits()(*string)
    GetCategory()(*RecommendationCategory)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCurrentScore()(*float64)
    GetDisplayName()(*string)
    GetFeatureAreas()([]RecommendationBase_featureAreas)
    GetImpactedResources()([]ImpactedResourceable)
    GetImpactStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetImpactType()(*string)
    GetInsights()(*string)
    GetLastCheckedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLastModifiedBy()(*string)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetMaxScore()(*float64)
    GetPostponeUntilDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetPriority()(*RecommendationPriority)
    GetRecommendationType()(*RecommendationBase_recommendationType)
    GetReleaseType()(*string)
    GetRemediationImpact()(*string)
    GetStatus()(*RecommendationStatus)
    SetActionSteps(value []ActionStepable)()
    SetBenefits(value *string)()
    SetCategory(value *RecommendationCategory)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCurrentScore(value *float64)()
    SetDisplayName(value *string)()
    SetFeatureAreas(value []RecommendationBase_featureAreas)()
    SetImpactedResources(value []ImpactedResourceable)()
    SetImpactStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetImpactType(value *string)()
    SetInsights(value *string)()
    SetLastCheckedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLastModifiedBy(value *string)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetMaxScore(value *float64)()
    SetPostponeUntilDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetPriority(value *RecommendationPriority)()
    SetRecommendationType(value *RecommendationBase_recommendationType)()
    SetReleaseType(value *string)()
    SetRemediationImpact(value *string)()
    SetStatus(value *RecommendationStatus)()
}

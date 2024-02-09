package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ImpactedResource struct {
    Entity
}
// NewImpactedResource instantiates a new ImpactedResource and sets the default values.
func NewImpactedResource()(*ImpactedResource) {
    m := &ImpactedResource{
        Entity: *NewEntity(),
    }
    return m
}
// CreateImpactedResourceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateImpactedResourceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewImpactedResource(), nil
}
// GetAddedDateTime gets the addedDateTime property value. The date and time when the impactedResource object was initially associated with the recommendation.
// returns a *Time when successful
func (m *ImpactedResource) GetAddedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("addedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetAdditionalDetails gets the additionalDetails property value. Additional information unique to the impactedResource to help contextualize the recommendation.
// returns a []KeyValueable when successful
func (m *ImpactedResource) GetAdditionalDetails()([]KeyValueable) {
    val, err := m.GetBackingStore().Get("additionalDetails")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]KeyValueable)
    }
    return nil
}
// GetApiUrl gets the apiUrl property value. The URL link to the corresponding Microsoft Entra resource.
// returns a *string when successful
func (m *ImpactedResource) GetApiUrl()(*string) {
    val, err := m.GetBackingStore().Get("apiUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDisplayName gets the displayName property value. Friendly name of the Microsoft Entra resource.
// returns a *string when successful
func (m *ImpactedResource) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ImpactedResource) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["addedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAddedDateTime(val)
        }
        return nil
    }
    res["additionalDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateKeyValueFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]KeyValueable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(KeyValueable)
                }
            }
            m.SetAdditionalDetails(res)
        }
        return nil
    }
    res["apiUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApiUrl(val)
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
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["owner"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwner(val)
        }
        return nil
    }
    res["portalUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPortalUrl(val)
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
    res["rank"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRank(val)
        }
        return nil
    }
    res["recommendationId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecommendationId(val)
        }
        return nil
    }
    res["resourceType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceType(val)
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
    res["subjectId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubjectId(val)
        }
        return nil
    }
    return res
}
// GetLastModifiedBy gets the lastModifiedBy property value. Name of the user or service that last updated the status.
// returns a *string when successful
func (m *ImpactedResource) GetLastModifiedBy()(*string) {
    val, err := m.GetBackingStore().Get("lastModifiedBy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The date and time when the status was last updated.
// returns a *string when successful
func (m *ImpactedResource) GetLastModifiedDateTime()(*string) {
    val, err := m.GetBackingStore().Get("lastModifiedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOwner gets the owner property value. The user responsible for maintaining the resource.
// returns a *string when successful
func (m *ImpactedResource) GetOwner()(*string) {
    val, err := m.GetBackingStore().Get("owner")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPortalUrl gets the portalUrl property value. The URL link to the corresponding Microsoft Entra admin center page of the resource.
// returns a *string when successful
func (m *ImpactedResource) GetPortalUrl()(*string) {
    val, err := m.GetBackingStore().Get("portalUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPostponeUntilDateTime gets the postponeUntilDateTime property value. The future date and time when the status of a postponed impactedResource will be active again.
// returns a *Time when successful
func (m *ImpactedResource) GetPostponeUntilDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("postponeUntilDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetRank gets the rank property value. Indicates the importance of the resource. A resource with a rank equal to 1 is of the highest importance.
// returns a *int32 when successful
func (m *ImpactedResource) GetRank()(*int32) {
    val, err := m.GetBackingStore().Get("rank")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetRecommendationId gets the recommendationId property value. The unique identifier of the recommendation that the resource is associated with.
// returns a *string when successful
func (m *ImpactedResource) GetRecommendationId()(*string) {
    val, err := m.GetBackingStore().Get("recommendationId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetResourceType gets the resourceType property value. Indicates the type of Microsoft Entra resource. Examples include user, application.
// returns a *string when successful
func (m *ImpactedResource) GetResourceType()(*string) {
    val, err := m.GetBackingStore().Get("resourceType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetStatus gets the status property value. The status property
// returns a *RecommendationStatus when successful
func (m *ImpactedResource) GetStatus()(*RecommendationStatus) {
    val, err := m.GetBackingStore().Get("status")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*RecommendationStatus)
    }
    return nil
}
// GetSubjectId gets the subjectId property value. The related unique identifier, depending on the resourceType. For example, this property is set to the applicationId if the resourceType is an application.
// returns a *string when successful
func (m *ImpactedResource) GetSubjectId()(*string) {
    val, err := m.GetBackingStore().Get("subjectId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ImpactedResource) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAdditionalDetails()))
        for i, v := range m.GetAdditionalDetails() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
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
        err = writer.WriteStringValue("lastModifiedBy", m.GetLastModifiedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
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
        err = writer.WriteTimeValue("postponeUntilDateTime", m.GetPostponeUntilDateTime())
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
    {
        err = writer.WriteStringValue("subjectId", m.GetSubjectId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAddedDateTime sets the addedDateTime property value. The date and time when the impactedResource object was initially associated with the recommendation.
func (m *ImpactedResource) SetAddedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("addedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetAdditionalDetails sets the additionalDetails property value. Additional information unique to the impactedResource to help contextualize the recommendation.
func (m *ImpactedResource) SetAdditionalDetails(value []KeyValueable)() {
    err := m.GetBackingStore().Set("additionalDetails", value)
    if err != nil {
        panic(err)
    }
}
// SetApiUrl sets the apiUrl property value. The URL link to the corresponding Microsoft Entra resource.
func (m *ImpactedResource) SetApiUrl(value *string)() {
    err := m.GetBackingStore().Set("apiUrl", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. Friendly name of the Microsoft Entra resource.
func (m *ImpactedResource) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetLastModifiedBy sets the lastModifiedBy property value. Name of the user or service that last updated the status.
func (m *ImpactedResource) SetLastModifiedBy(value *string)() {
    err := m.GetBackingStore().Set("lastModifiedBy", value)
    if err != nil {
        panic(err)
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The date and time when the status was last updated.
func (m *ImpactedResource) SetLastModifiedDateTime(value *string)() {
    err := m.GetBackingStore().Set("lastModifiedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetOwner sets the owner property value. The user responsible for maintaining the resource.
func (m *ImpactedResource) SetOwner(value *string)() {
    err := m.GetBackingStore().Set("owner", value)
    if err != nil {
        panic(err)
    }
}
// SetPortalUrl sets the portalUrl property value. The URL link to the corresponding Microsoft Entra admin center page of the resource.
func (m *ImpactedResource) SetPortalUrl(value *string)() {
    err := m.GetBackingStore().Set("portalUrl", value)
    if err != nil {
        panic(err)
    }
}
// SetPostponeUntilDateTime sets the postponeUntilDateTime property value. The future date and time when the status of a postponed impactedResource will be active again.
func (m *ImpactedResource) SetPostponeUntilDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("postponeUntilDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetRank sets the rank property value. Indicates the importance of the resource. A resource with a rank equal to 1 is of the highest importance.
func (m *ImpactedResource) SetRank(value *int32)() {
    err := m.GetBackingStore().Set("rank", value)
    if err != nil {
        panic(err)
    }
}
// SetRecommendationId sets the recommendationId property value. The unique identifier of the recommendation that the resource is associated with.
func (m *ImpactedResource) SetRecommendationId(value *string)() {
    err := m.GetBackingStore().Set("recommendationId", value)
    if err != nil {
        panic(err)
    }
}
// SetResourceType sets the resourceType property value. Indicates the type of Microsoft Entra resource. Examples include user, application.
func (m *ImpactedResource) SetResourceType(value *string)() {
    err := m.GetBackingStore().Set("resourceType", value)
    if err != nil {
        panic(err)
    }
}
// SetStatus sets the status property value. The status property
func (m *ImpactedResource) SetStatus(value *RecommendationStatus)() {
    err := m.GetBackingStore().Set("status", value)
    if err != nil {
        panic(err)
    }
}
// SetSubjectId sets the subjectId property value. The related unique identifier, depending on the resourceType. For example, this property is set to the applicationId if the resourceType is an application.
func (m *ImpactedResource) SetSubjectId(value *string)() {
    err := m.GetBackingStore().Set("subjectId", value)
    if err != nil {
        panic(err)
    }
}
type ImpactedResourceable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAddedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetAdditionalDetails()([]KeyValueable)
    GetApiUrl()(*string)
    GetDisplayName()(*string)
    GetLastModifiedBy()(*string)
    GetLastModifiedDateTime()(*string)
    GetOwner()(*string)
    GetPortalUrl()(*string)
    GetPostponeUntilDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetRank()(*int32)
    GetRecommendationId()(*string)
    GetResourceType()(*string)
    GetStatus()(*RecommendationStatus)
    GetSubjectId()(*string)
    SetAddedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetAdditionalDetails(value []KeyValueable)()
    SetApiUrl(value *string)()
    SetDisplayName(value *string)()
    SetLastModifiedBy(value *string)()
    SetLastModifiedDateTime(value *string)()
    SetOwner(value *string)()
    SetPortalUrl(value *string)()
    SetPostponeUntilDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetRank(value *int32)()
    SetRecommendationId(value *string)()
    SetResourceType(value *string)()
    SetStatus(value *RecommendationStatus)()
    SetSubjectId(value *string)()
}

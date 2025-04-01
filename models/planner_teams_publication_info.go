package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type PlannerTeamsPublicationInfo struct {
    PlannerTaskCreation
}
// NewPlannerTeamsPublicationInfo instantiates a new PlannerTeamsPublicationInfo and sets the default values.
func NewPlannerTeamsPublicationInfo()(*PlannerTeamsPublicationInfo) {
    m := &PlannerTeamsPublicationInfo{
        PlannerTaskCreation: *NewPlannerTaskCreation(),
    }
    odataTypeValue := "#microsoft.graph.plannerTeamsPublicationInfo"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreatePlannerTeamsPublicationInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePlannerTeamsPublicationInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPlannerTeamsPublicationInfo(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PlannerTeamsPublicationInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.PlannerTaskCreation.GetFieldDeserializers()
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
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["publicationId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublicationId(val)
        }
        return nil
    }
    res["publicationName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublicationName(val)
        }
        return nil
    }
    res["publishedToPlanId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublishedToPlanId(val)
        }
        return nil
    }
    res["publishingTeamId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublishingTeamId(val)
        }
        return nil
    }
    res["publishingTeamName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublishingTeamName(val)
        }
        return nil
    }
    return res
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The date and time when this task was last modified by the publication process. Read-only.
// returns a *Time when successful
func (m *PlannerTeamsPublicationInfo) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastModifiedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *PlannerTeamsPublicationInfo) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPublicationId gets the publicationId property value. The identifier of the publication. Read-only.
// returns a *string when successful
func (m *PlannerTeamsPublicationInfo) GetPublicationId()(*string) {
    val, err := m.GetBackingStore().Get("publicationId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPublicationName gets the publicationName property value. The name of the published task list. Read-only.
// returns a *string when successful
func (m *PlannerTeamsPublicationInfo) GetPublicationName()(*string) {
    val, err := m.GetBackingStore().Get("publicationName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPublishedToPlanId gets the publishedToPlanId property value. The identifier of the plannerPlan this task was originally placed in. Read-only.
// returns a *string when successful
func (m *PlannerTeamsPublicationInfo) GetPublishedToPlanId()(*string) {
    val, err := m.GetBackingStore().Get("publishedToPlanId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPublishingTeamId gets the publishingTeamId property value. The identifier of the team that initiated the publication process. Read-only.
// returns a *string when successful
func (m *PlannerTeamsPublicationInfo) GetPublishingTeamId()(*string) {
    val, err := m.GetBackingStore().Get("publishingTeamId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPublishingTeamName gets the publishingTeamName property value. The display name of the team that initiated the publication process. This display name is for reference only, and might not represent the most up-to-date name of the team. Read-only.
// returns a *string when successful
func (m *PlannerTeamsPublicationInfo) GetPublishingTeamName()(*string) {
    val, err := m.GetBackingStore().Get("publishingTeamName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *PlannerTeamsPublicationInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.PlannerTaskCreation.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("publicationId", m.GetPublicationId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("publicationName", m.GetPublicationName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("publishedToPlanId", m.GetPublishedToPlanId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("publishingTeamId", m.GetPublishingTeamId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("publishingTeamName", m.GetPublishingTeamName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The date and time when this task was last modified by the publication process. Read-only.
func (m *PlannerTeamsPublicationInfo) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastModifiedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *PlannerTeamsPublicationInfo) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetPublicationId sets the publicationId property value. The identifier of the publication. Read-only.
func (m *PlannerTeamsPublicationInfo) SetPublicationId(value *string)() {
    err := m.GetBackingStore().Set("publicationId", value)
    if err != nil {
        panic(err)
    }
}
// SetPublicationName sets the publicationName property value. The name of the published task list. Read-only.
func (m *PlannerTeamsPublicationInfo) SetPublicationName(value *string)() {
    err := m.GetBackingStore().Set("publicationName", value)
    if err != nil {
        panic(err)
    }
}
// SetPublishedToPlanId sets the publishedToPlanId property value. The identifier of the plannerPlan this task was originally placed in. Read-only.
func (m *PlannerTeamsPublicationInfo) SetPublishedToPlanId(value *string)() {
    err := m.GetBackingStore().Set("publishedToPlanId", value)
    if err != nil {
        panic(err)
    }
}
// SetPublishingTeamId sets the publishingTeamId property value. The identifier of the team that initiated the publication process. Read-only.
func (m *PlannerTeamsPublicationInfo) SetPublishingTeamId(value *string)() {
    err := m.GetBackingStore().Set("publishingTeamId", value)
    if err != nil {
        panic(err)
    }
}
// SetPublishingTeamName sets the publishingTeamName property value. The display name of the team that initiated the publication process. This display name is for reference only, and might not represent the most up-to-date name of the team. Read-only.
func (m *PlannerTeamsPublicationInfo) SetPublishingTeamName(value *string)() {
    err := m.GetBackingStore().Set("publishingTeamName", value)
    if err != nil {
        panic(err)
    }
}
type PlannerTeamsPublicationInfoable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    PlannerTaskCreationable
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetOdataType()(*string)
    GetPublicationId()(*string)
    GetPublicationName()(*string)
    GetPublishedToPlanId()(*string)
    GetPublishingTeamId()(*string)
    GetPublishingTeamName()(*string)
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetOdataType(value *string)()
    SetPublicationId(value *string)()
    SetPublicationName(value *string)()
    SetPublishedToPlanId(value *string)()
    SetPublishingTeamId(value *string)()
    SetPublishingTeamName(value *string)()
}

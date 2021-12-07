package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PlannerTeamsPublicationInfo 
type PlannerTeamsPublicationInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The date and time when this task was last modified by the publication process. Read-only.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The identifier of the publication. Read-only.
    publicationId *string;
    // The identifier of the plannerPlan this task was originally placed in. Read-only.
    publishedToPlanId *string;
    // The identifier of the team that initiated the publication process. Read-only.
    publishingTeamId *string;
    // The display name of the team that initiated the publication process. This display name is for reference only, and might not represent the most up-to-date name of the team. Read-only.
    publishingTeamName *string;
}
// NewPlannerTeamsPublicationInfo instantiates a new plannerTeamsPublicationInfo and sets the default values.
func NewPlannerTeamsPublicationInfo()(*PlannerTeamsPublicationInfo) {
    m := &PlannerTeamsPublicationInfo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PlannerTeamsPublicationInfo) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The date and time when this task was last modified by the publication process. Read-only.
func (m *PlannerTeamsPublicationInfo) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// GetPublicationId gets the publicationId property value. The identifier of the publication. Read-only.
func (m *PlannerTeamsPublicationInfo) GetPublicationId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.publicationId
    }
}
// GetPublishedToPlanId gets the publishedToPlanId property value. The identifier of the plannerPlan this task was originally placed in. Read-only.
func (m *PlannerTeamsPublicationInfo) GetPublishedToPlanId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.publishedToPlanId
    }
}
// GetPublishingTeamId gets the publishingTeamId property value. The identifier of the team that initiated the publication process. Read-only.
func (m *PlannerTeamsPublicationInfo) GetPublishingTeamId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.publishingTeamId
    }
}
// GetPublishingTeamName gets the publishingTeamName property value. The display name of the team that initiated the publication process. This display name is for reference only, and might not represent the most up-to-date name of the team. Read-only.
func (m *PlannerTeamsPublicationInfo) GetPublishingTeamName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.publishingTeamName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PlannerTeamsPublicationInfo) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
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
    res["publicationId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublicationId(val)
        }
        return nil
    }
    res["publishedToPlanId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublishedToPlanId(val)
        }
        return nil
    }
    res["publishingTeamId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublishingTeamId(val)
        }
        return nil
    }
    res["publishingTeamName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
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
func (m *PlannerTeamsPublicationInfo) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *PlannerTeamsPublicationInfo) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("publicationId", m.GetPublicationId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("publishedToPlanId", m.GetPublishedToPlanId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("publishingTeamId", m.GetPublishingTeamId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("publishingTeamName", m.GetPublishingTeamName())
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PlannerTeamsPublicationInfo) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The date and time when this task was last modified by the publication process. Read-only.
func (m *PlannerTeamsPublicationInfo) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastModifiedDateTime = value
    }
}
// SetPublicationId sets the publicationId property value. The identifier of the publication. Read-only.
func (m *PlannerTeamsPublicationInfo) SetPublicationId(value *string)() {
    if m != nil {
        m.publicationId = value
    }
}
// SetPublishedToPlanId sets the publishedToPlanId property value. The identifier of the plannerPlan this task was originally placed in. Read-only.
func (m *PlannerTeamsPublicationInfo) SetPublishedToPlanId(value *string)() {
    if m != nil {
        m.publishedToPlanId = value
    }
}
// SetPublishingTeamId sets the publishingTeamId property value. The identifier of the team that initiated the publication process. Read-only.
func (m *PlannerTeamsPublicationInfo) SetPublishingTeamId(value *string)() {
    if m != nil {
        m.publishingTeamId = value
    }
}
// SetPublishingTeamName sets the publishingTeamName property value. The display name of the team that initiated the publication process. This display name is for reference only, and might not represent the most up-to-date name of the team. Read-only.
func (m *PlannerTeamsPublicationInfo) SetPublishingTeamName(value *string)() {
    if m != nil {
        m.publishingTeamName = value
    }
}

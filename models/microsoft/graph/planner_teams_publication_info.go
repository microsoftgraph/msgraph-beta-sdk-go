package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new plannerTeamsPublicationInfo and sets the default values.
func NewPlannerTeamsPublicationInfo()(*PlannerTeamsPublicationInfo) {
    m := &PlannerTeamsPublicationInfo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PlannerTeamsPublicationInfo) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the lastModifiedDateTime property value. The date and time when this task was last modified by the publication process. Read-only.
func (m *PlannerTeamsPublicationInfo) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// Gets the publicationId property value. The identifier of the publication. Read-only.
func (m *PlannerTeamsPublicationInfo) GetPublicationId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.publicationId
    }
}
// Gets the publishedToPlanId property value. The identifier of the plannerPlan this task was originally placed in. Read-only.
func (m *PlannerTeamsPublicationInfo) GetPublishedToPlanId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.publishedToPlanId
    }
}
// Gets the publishingTeamId property value. The identifier of the team that initiated the publication process. Read-only.
func (m *PlannerTeamsPublicationInfo) GetPublishingTeamId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.publishingTeamId
    }
}
// Gets the publishingTeamName property value. The display name of the team that initiated the publication process. This display name is for reference only, and might not represent the most up-to-date name of the team. Read-only.
func (m *PlannerTeamsPublicationInfo) GetPublishingTeamName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.publishingTeamName
    }
}
// The deserialization information for the current model
func (m *PlannerTeamsPublicationInfo) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["lastModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastModifiedDateTime(val)
        return nil
    }
    res["publicationId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPublicationId(val)
        return nil
    }
    res["publishedToPlanId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPublishedToPlanId(val)
        return nil
    }
    res["publishingTeamId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPublishingTeamId(val)
        return nil
    }
    res["publishingTeamName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPublishingTeamName(val)
        return nil
    }
    return res
}
func (m *PlannerTeamsPublicationInfo) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *PlannerTeamsPublicationInfo) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the lastModifiedDateTime property value. The date and time when this task was last modified by the publication process. Read-only.
// Parameters:
//  - value : Value to set for the lastModifiedDateTime property.
func (m *PlannerTeamsPublicationInfo) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// Sets the publicationId property value. The identifier of the publication. Read-only.
// Parameters:
//  - value : Value to set for the publicationId property.
func (m *PlannerTeamsPublicationInfo) SetPublicationId(value *string)() {
    m.publicationId = value
}
// Sets the publishedToPlanId property value. The identifier of the plannerPlan this task was originally placed in. Read-only.
// Parameters:
//  - value : Value to set for the publishedToPlanId property.
func (m *PlannerTeamsPublicationInfo) SetPublishedToPlanId(value *string)() {
    m.publishedToPlanId = value
}
// Sets the publishingTeamId property value. The identifier of the team that initiated the publication process. Read-only.
// Parameters:
//  - value : Value to set for the publishingTeamId property.
func (m *PlannerTeamsPublicationInfo) SetPublishingTeamId(value *string)() {
    m.publishingTeamId = value
}
// Sets the publishingTeamName property value. The display name of the team that initiated the publication process. This display name is for reference only, and might not represent the most up-to-date name of the team. Read-only.
// Parameters:
//  - value : Value to set for the publishingTeamName property.
func (m *PlannerTeamsPublicationInfo) SetPublishingTeamName(value *string)() {
    m.publishingTeamName = value
}

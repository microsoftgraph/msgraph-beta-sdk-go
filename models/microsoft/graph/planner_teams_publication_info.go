package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type PlannerTeamsPublicationInfo struct {
    additionalData map[string]interface{};
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    publicationId *string;
    publishedToPlanId *string;
    publishingTeamId *string;
    publishingTeamName *string;
}
func NewPlannerTeamsPublicationInfo()(*PlannerTeamsPublicationInfo) {
    m := &PlannerTeamsPublicationInfo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *PlannerTeamsPublicationInfo) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *PlannerTeamsPublicationInfo) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
func (m *PlannerTeamsPublicationInfo) GetPublicationId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.publicationId
    }
}
func (m *PlannerTeamsPublicationInfo) GetPublishedToPlanId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.publishedToPlanId
    }
}
func (m *PlannerTeamsPublicationInfo) GetPublishingTeamId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.publishingTeamId
    }
}
func (m *PlannerTeamsPublicationInfo) GetPublishingTeamName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.publishingTeamName
    }
}
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
func (m *PlannerTeamsPublicationInfo) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *PlannerTeamsPublicationInfo) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
func (m *PlannerTeamsPublicationInfo) SetPublicationId(value *string)() {
    m.publicationId = value
}
func (m *PlannerTeamsPublicationInfo) SetPublishedToPlanId(value *string)() {
    m.publishedToPlanId = value
}
func (m *PlannerTeamsPublicationInfo) SetPublishingTeamId(value *string)() {
    m.publishingTeamId = value
}
func (m *PlannerTeamsPublicationInfo) SetPublishingTeamName(value *string)() {
    m.publishingTeamName = value
}

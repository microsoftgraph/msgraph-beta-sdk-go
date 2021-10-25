package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type PlannerTaskCreation struct {
    additionalData map[string]interface{};
    teamsPublicationInfo *PlannerTeamsPublicationInfo;
}
func NewPlannerTaskCreation()(*PlannerTaskCreation) {
    m := &PlannerTaskCreation{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *PlannerTaskCreation) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *PlannerTaskCreation) GetTeamsPublicationInfo()(*PlannerTeamsPublicationInfo) {
    if m == nil {
        return nil
    } else {
        return m.teamsPublicationInfo
    }
}
func (m *PlannerTaskCreation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["teamsPublicationInfo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPlannerTeamsPublicationInfo() })
        if err != nil {
            return err
        }
        m.SetTeamsPublicationInfo(val.(*PlannerTeamsPublicationInfo))
        return nil
    }
    return res
}
func (m *PlannerTaskCreation) IsNil()(bool) {
    return m == nil
}
func (m *PlannerTaskCreation) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("teamsPublicationInfo", m.GetTeamsPublicationInfo())
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
func (m *PlannerTaskCreation) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *PlannerTaskCreation) SetTeamsPublicationInfo(value *PlannerTeamsPublicationInfo)() {
    m.teamsPublicationInfo = value
}

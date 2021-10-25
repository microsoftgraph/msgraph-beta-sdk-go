package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type OnPremisesAgent struct {
    Entity
    agentGroups []OnPremisesAgentGroup;
    externalIp *string;
    machineName *string;
    status *AgentStatus;
    supportedPublishingTypes []OnPremisesPublishingType;
}
func NewOnPremisesAgent()(*OnPremisesAgent) {
    m := &OnPremisesAgent{
        Entity: *NewEntity(),
    }
    return m
}
func (m *OnPremisesAgent) GetAgentGroups()([]OnPremisesAgentGroup) {
    if m == nil {
        return nil
    } else {
        return m.agentGroups
    }
}
func (m *OnPremisesAgent) GetExternalIp()(*string) {
    if m == nil {
        return nil
    } else {
        return m.externalIp
    }
}
func (m *OnPremisesAgent) GetMachineName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.machineName
    }
}
func (m *OnPremisesAgent) GetStatus()(*AgentStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
func (m *OnPremisesAgent) GetSupportedPublishingTypes()([]OnPremisesPublishingType) {
    if m == nil {
        return nil
    } else {
        return m.supportedPublishingTypes
    }
}
func (m *OnPremisesAgent) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["agentGroups"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOnPremisesAgentGroup() })
        if err != nil {
            return err
        }
        res := make([]OnPremisesAgentGroup, len(val))
        for i, v := range val {
            res[i] = *(v.(*OnPremisesAgentGroup))
        }
        m.SetAgentGroups(res)
        return nil
    }
    res["externalIp"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetExternalIp(val)
        return nil
    }
    res["machineName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetMachineName(val)
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAgentStatus)
        if err != nil {
            return err
        }
        cast := val.(AgentStatus)
        m.SetStatus(&cast)
        return nil
    }
    res["supportedPublishingTypes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseOnPremisesPublishingType)
        if err != nil {
            return err
        }
        res := make([]OnPremisesPublishingType, len(val))
        for i, v := range val {
            res[i] = *(v.(*OnPremisesPublishingType))
        }
        m.SetSupportedPublishingTypes(res)
        return nil
    }
    return res
}
func (m *OnPremisesAgent) IsNil()(bool) {
    return m == nil
}
func (m *OnPremisesAgent) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAgentGroups()))
        for i, v := range m.GetAgentGroups() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("agentGroups", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("externalIp", m.GetExternalIp())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("machineName", m.GetMachineName())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := m.GetStatus().String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("supportedPublishingTypes", SerializeOnPremisesPublishingType(m.GetSupportedPublishingTypes()))
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *OnPremisesAgent) SetAgentGroups(value []OnPremisesAgentGroup)() {
    m.agentGroups = value
}
func (m *OnPremisesAgent) SetExternalIp(value *string)() {
    m.externalIp = value
}
func (m *OnPremisesAgent) SetMachineName(value *string)() {
    m.machineName = value
}
func (m *OnPremisesAgent) SetStatus(value *AgentStatus)() {
    m.status = value
}
func (m *OnPremisesAgent) SetSupportedPublishingTypes(value []OnPremisesPublishingType)() {
    m.supportedPublishingTypes = value
}

package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type OnPremisesAgent struct {
    Entity
    // List of onPremisesAgentGroups that an onPremisesAgent is assigned to. Read-only. Nullable.
    agentGroups []OnPremisesAgentGroup;
    // The external IP address as detected by the service for the agent machine. Read-only
    externalIp *string;
    // The name of the machine that the aggent is running on. Read-only
    machineName *string;
    // Possible values are: active, inactive.
    status *AgentStatus;
    // 
    supportedPublishingTypes []OnPremisesPublishingType;
}
// Instantiates a new onPremisesAgent and sets the default values.
func NewOnPremisesAgent()(*OnPremisesAgent) {
    m := &OnPremisesAgent{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the agentGroups property value. List of onPremisesAgentGroups that an onPremisesAgent is assigned to. Read-only. Nullable.
func (m *OnPremisesAgent) GetAgentGroups()([]OnPremisesAgentGroup) {
    if m == nil {
        return nil
    } else {
        return m.agentGroups
    }
}
// Gets the externalIp property value. The external IP address as detected by the service for the agent machine. Read-only
func (m *OnPremisesAgent) GetExternalIp()(*string) {
    if m == nil {
        return nil
    } else {
        return m.externalIp
    }
}
// Gets the machineName property value. The name of the machine that the aggent is running on. Read-only
func (m *OnPremisesAgent) GetMachineName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.machineName
    }
}
// Gets the status property value. Possible values are: active, inactive.
func (m *OnPremisesAgent) GetStatus()(*AgentStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// Gets the supportedPublishingTypes property value. 
func (m *OnPremisesAgent) GetSupportedPublishingTypes()([]OnPremisesPublishingType) {
    if m == nil {
        return nil
    } else {
        return m.supportedPublishingTypes
    }
}
// The deserialization information for the current model
func (m *OnPremisesAgent) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["agentGroups"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOnPremisesAgentGroup() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OnPremisesAgentGroup, len(val))
            for i, v := range val {
                res[i] = *(v.(*OnPremisesAgentGroup))
            }
            m.SetAgentGroups(res)
        }
        return nil
    }
    res["externalIp"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalIp(val)
        }
        return nil
    }
    res["machineName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMachineName(val)
        }
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAgentStatus)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(AgentStatus)
            m.SetStatus(&cast)
        }
        return nil
    }
    res["supportedPublishingTypes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseOnPremisesPublishingType)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OnPremisesPublishingType, len(val))
            for i, v := range val {
                res[i] = *(v.(*OnPremisesPublishingType))
            }
            m.SetSupportedPublishingTypes(res)
        }
        return nil
    }
    return res
}
func (m *OnPremisesAgent) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the agentGroups property value. List of onPremisesAgentGroups that an onPremisesAgent is assigned to. Read-only. Nullable.
// Parameters:
//  - value : Value to set for the agentGroups property.
func (m *OnPremisesAgent) SetAgentGroups(value []OnPremisesAgentGroup)() {
    m.agentGroups = value
}
// Sets the externalIp property value. The external IP address as detected by the service for the agent machine. Read-only
// Parameters:
//  - value : Value to set for the externalIp property.
func (m *OnPremisesAgent) SetExternalIp(value *string)() {
    m.externalIp = value
}
// Sets the machineName property value. The name of the machine that the aggent is running on. Read-only
// Parameters:
//  - value : Value to set for the machineName property.
func (m *OnPremisesAgent) SetMachineName(value *string)() {
    m.machineName = value
}
// Sets the status property value. Possible values are: active, inactive.
// Parameters:
//  - value : Value to set for the status property.
func (m *OnPremisesAgent) SetStatus(value *AgentStatus)() {
    m.status = value
}
// Sets the supportedPublishingTypes property value. 
// Parameters:
//  - value : Value to set for the supportedPublishingTypes property.
func (m *OnPremisesAgent) SetSupportedPublishingTypes(value []OnPremisesPublishingType)() {
    m.supportedPublishingTypes = value
}

package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// OnPremisesAgent 
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
// NewOnPremisesAgent instantiates a new onPremisesAgent and sets the default values.
func NewOnPremisesAgent()(*OnPremisesAgent) {
    m := &OnPremisesAgent{
        Entity: *NewEntity(),
    }
    return m
}
// GetAgentGroups gets the agentGroups property value. List of onPremisesAgentGroups that an onPremisesAgent is assigned to. Read-only. Nullable.
func (m *OnPremisesAgent) GetAgentGroups()([]OnPremisesAgentGroup) {
    if m == nil {
        return nil
    } else {
        return m.agentGroups
    }
}
// GetExternalIp gets the externalIp property value. The external IP address as detected by the service for the agent machine. Read-only
func (m *OnPremisesAgent) GetExternalIp()(*string) {
    if m == nil {
        return nil
    } else {
        return m.externalIp
    }
}
// GetMachineName gets the machineName property value. The name of the machine that the aggent is running on. Read-only
func (m *OnPremisesAgent) GetMachineName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.machineName
    }
}
// GetStatus gets the status property value. Possible values are: active, inactive.
func (m *OnPremisesAgent) GetStatus()(*AgentStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// GetSupportedPublishingTypes gets the supportedPublishingTypes property value. 
func (m *OnPremisesAgent) GetSupportedPublishingTypes()([]OnPremisesPublishingType) {
    if m == nil {
        return nil
    } else {
        return m.supportedPublishingTypes
    }
}
// GetFieldDeserializers the deserialization information for the current model
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
            m.SetStatus(val.(*AgentStatus))
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
// Serialize serializes information the current object
func (m *OnPremisesAgent) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAgentGroups() != nil {
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
        cast := (*m.GetStatus()).String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetSupportedPublishingTypes() != nil {
        err = writer.WriteCollectionOfStringValues("supportedPublishingTypes", SerializeOnPremisesPublishingType(m.GetSupportedPublishingTypes()))
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAgentGroups sets the agentGroups property value. List of onPremisesAgentGroups that an onPremisesAgent is assigned to. Read-only. Nullable.
func (m *OnPremisesAgent) SetAgentGroups(value []OnPremisesAgentGroup)() {
    if m != nil {
        m.agentGroups = value
    }
}
// SetExternalIp sets the externalIp property value. The external IP address as detected by the service for the agent machine. Read-only
func (m *OnPremisesAgent) SetExternalIp(value *string)() {
    if m != nil {
        m.externalIp = value
    }
}
// SetMachineName sets the machineName property value. The name of the machine that the aggent is running on. Read-only
func (m *OnPremisesAgent) SetMachineName(value *string)() {
    if m != nil {
        m.machineName = value
    }
}
// SetStatus sets the status property value. Possible values are: active, inactive.
func (m *OnPremisesAgent) SetStatus(value *AgentStatus)() {
    if m != nil {
        m.status = value
    }
}
// SetSupportedPublishingTypes sets the supportedPublishingTypes property value. 
func (m *OnPremisesAgent) SetSupportedPublishingTypes(value []OnPremisesPublishingType)() {
    if m != nil {
        m.supportedPublishingTypes = value
    }
}

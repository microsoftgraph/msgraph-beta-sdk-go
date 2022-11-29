package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OnPremisesAgent provides operations to manage the collection of accessReview entities.
type OnPremisesAgent struct {
    Entity
    // List of onPremisesAgentGroups that an onPremisesAgent is assigned to. Read-only. Nullable.
    agentGroups []OnPremisesAgentGroupable
    // The external IP address as detected by the service for the agent machine. Read-only
    externalIp *string
    // The name of the machine that the aggent is running on. Read-only
    machineName *string
    // The status property
    status *AgentStatus
    // The supportedPublishingTypes property
    supportedPublishingTypes []OnPremisesPublishingType
}
// NewOnPremisesAgent instantiates a new onPremisesAgent and sets the default values.
func NewOnPremisesAgent()(*OnPremisesAgent) {
    m := &OnPremisesAgent{
        Entity: *NewEntity(),
    }
    return m
}
// CreateOnPremisesAgentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOnPremisesAgentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOnPremisesAgent(), nil
}
// GetAgentGroups gets the agentGroups property value. List of onPremisesAgentGroups that an onPremisesAgent is assigned to. Read-only. Nullable.
func (m *OnPremisesAgent) GetAgentGroups()([]OnPremisesAgentGroupable) {
    return m.agentGroups
}
// GetExternalIp gets the externalIp property value. The external IP address as detected by the service for the agent machine. Read-only
func (m *OnPremisesAgent) GetExternalIp()(*string) {
    return m.externalIp
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OnPremisesAgent) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["agentGroups"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateOnPremisesAgentGroupFromDiscriminatorValue , m.SetAgentGroups)
    res["externalIp"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetExternalIp)
    res["machineName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetMachineName)
    res["status"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseAgentStatus , m.SetStatus)
    res["supportedPublishingTypes"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfEnumValues(ParseOnPremisesPublishingType , m.SetSupportedPublishingTypes)
    return res
}
// GetMachineName gets the machineName property value. The name of the machine that the aggent is running on. Read-only
func (m *OnPremisesAgent) GetMachineName()(*string) {
    return m.machineName
}
// GetStatus gets the status property value. The status property
func (m *OnPremisesAgent) GetStatus()(*AgentStatus) {
    return m.status
}
// GetSupportedPublishingTypes gets the supportedPublishingTypes property value. The supportedPublishingTypes property
func (m *OnPremisesAgent) GetSupportedPublishingTypes()([]OnPremisesPublishingType) {
    return m.supportedPublishingTypes
}
// Serialize serializes information the current object
func (m *OnPremisesAgent) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAgentGroups() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetAgentGroups())
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
func (m *OnPremisesAgent) SetAgentGroups(value []OnPremisesAgentGroupable)() {
    m.agentGroups = value
}
// SetExternalIp sets the externalIp property value. The external IP address as detected by the service for the agent machine. Read-only
func (m *OnPremisesAgent) SetExternalIp(value *string)() {
    m.externalIp = value
}
// SetMachineName sets the machineName property value. The name of the machine that the aggent is running on. Read-only
func (m *OnPremisesAgent) SetMachineName(value *string)() {
    m.machineName = value
}
// SetStatus sets the status property value. The status property
func (m *OnPremisesAgent) SetStatus(value *AgentStatus)() {
    m.status = value
}
// SetSupportedPublishingTypes sets the supportedPublishingTypes property value. The supportedPublishingTypes property
func (m *OnPremisesAgent) SetSupportedPublishingTypes(value []OnPremisesPublishingType)() {
    m.supportedPublishingTypes = value
}

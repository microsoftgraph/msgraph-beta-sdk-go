package industrydata

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

type IndustryDataRoot struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
}
// NewIndustryDataRoot instantiates a new IndustryDataRoot and sets the default values.
func NewIndustryDataRoot()(*IndustryDataRoot) {
    m := &IndustryDataRoot{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreateIndustryDataRootFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateIndustryDataRootFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIndustryDataRoot(), nil
}
// GetDataConnectors gets the dataConnectors property value. Set of connectors for importing data from source systems.
// returns a []IndustryDataConnectorable when successful
func (m *IndustryDataRoot) GetDataConnectors()([]IndustryDataConnectorable) {
    val, err := m.GetBackingStore().Get("dataConnectors")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]IndustryDataConnectorable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *IndustryDataRoot) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["dataConnectors"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateIndustryDataConnectorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IndustryDataConnectorable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(IndustryDataConnectorable)
                }
            }
            m.SetDataConnectors(res)
        }
        return nil
    }
    res["inboundFlows"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateInboundFlowFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]InboundFlowable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(InboundFlowable)
                }
            }
            m.SetInboundFlows(res)
        }
        return nil
    }
    res["operations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateLongRunningOperationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.LongRunningOperationable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.LongRunningOperationable)
                }
            }
            m.SetOperations(res)
        }
        return nil
    }
    res["outboundProvisioningFlowSets"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateOutboundProvisioningFlowSetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OutboundProvisioningFlowSetable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(OutboundProvisioningFlowSetable)
                }
            }
            m.SetOutboundProvisioningFlowSets(res)
        }
        return nil
    }
    res["referenceDefinitions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateReferenceDefinitionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ReferenceDefinitionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ReferenceDefinitionable)
                }
            }
            m.SetReferenceDefinitions(res)
        }
        return nil
    }
    res["roleGroups"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRoleGroupFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RoleGroupable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(RoleGroupable)
                }
            }
            m.SetRoleGroups(res)
        }
        return nil
    }
    res["runs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateIndustryDataRunFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IndustryDataRunable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(IndustryDataRunable)
                }
            }
            m.SetRuns(res)
        }
        return nil
    }
    res["sourceSystems"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSourceSystemDefinitionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SourceSystemDefinitionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(SourceSystemDefinitionable)
                }
            }
            m.SetSourceSystems(res)
        }
        return nil
    }
    res["years"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateYearTimePeriodDefinitionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]YearTimePeriodDefinitionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(YearTimePeriodDefinitionable)
                }
            }
            m.SetYears(res)
        }
        return nil
    }
    return res
}
// GetInboundFlows gets the inboundFlows property value. Set of data import flow activities to bring data into the canonical store via a connector.
// returns a []InboundFlowable when successful
func (m *IndustryDataRoot) GetInboundFlows()([]InboundFlowable) {
    val, err := m.GetBackingStore().Get("inboundFlows")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]InboundFlowable)
    }
    return nil
}
// GetOperations gets the operations property value. Set of ephemeral operations that the system runs currently. Read-only.
// returns a []LongRunningOperationable when successful
func (m *IndustryDataRoot) GetOperations()([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.LongRunningOperationable) {
    val, err := m.GetBackingStore().Get("operations")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.LongRunningOperationable)
    }
    return nil
}
// GetOutboundProvisioningFlowSets gets the outboundProvisioningFlowSets property value. The outboundProvisioningFlowSets property
// returns a []OutboundProvisioningFlowSetable when successful
func (m *IndustryDataRoot) GetOutboundProvisioningFlowSets()([]OutboundProvisioningFlowSetable) {
    val, err := m.GetBackingStore().Get("outboundProvisioningFlowSets")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]OutboundProvisioningFlowSetable)
    }
    return nil
}
// GetReferenceDefinitions gets the referenceDefinitions property value. Set of user modifiable system picker types.
// returns a []ReferenceDefinitionable when successful
func (m *IndustryDataRoot) GetReferenceDefinitions()([]ReferenceDefinitionable) {
    val, err := m.GetBackingStore().Get("referenceDefinitions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ReferenceDefinitionable)
    }
    return nil
}
// GetRoleGroups gets the roleGroups property value. Set of groups of individual roles that makes role-based admin simpler.
// returns a []RoleGroupable when successful
func (m *IndustryDataRoot) GetRoleGroups()([]RoleGroupable) {
    val, err := m.GetBackingStore().Get("roleGroups")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]RoleGroupable)
    }
    return nil
}
// GetRuns gets the runs property value. Set of ephemeral runs which present the point-in-time that diagnostic state of activities performed by the system. Read-only.
// returns a []IndustryDataRunable when successful
func (m *IndustryDataRoot) GetRuns()([]IndustryDataRunable) {
    val, err := m.GetBackingStore().Get("runs")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]IndustryDataRunable)
    }
    return nil
}
// GetSourceSystems gets the sourceSystems property value. Set of source definitions that represents real-world external systems.
// returns a []SourceSystemDefinitionable when successful
func (m *IndustryDataRoot) GetSourceSystems()([]SourceSystemDefinitionable) {
    val, err := m.GetBackingStore().Get("sourceSystems")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]SourceSystemDefinitionable)
    }
    return nil
}
// GetYears gets the years property value. Set of years represented in the system.
// returns a []YearTimePeriodDefinitionable when successful
func (m *IndustryDataRoot) GetYears()([]YearTimePeriodDefinitionable) {
    val, err := m.GetBackingStore().Get("years")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]YearTimePeriodDefinitionable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *IndustryDataRoot) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetDataConnectors() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDataConnectors()))
        for i, v := range m.GetDataConnectors() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("dataConnectors", cast)
        if err != nil {
            return err
        }
    }
    if m.GetInboundFlows() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetInboundFlows()))
        for i, v := range m.GetInboundFlows() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("inboundFlows", cast)
        if err != nil {
            return err
        }
    }
    if m.GetOperations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetOperations()))
        for i, v := range m.GetOperations() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("operations", cast)
        if err != nil {
            return err
        }
    }
    if m.GetOutboundProvisioningFlowSets() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetOutboundProvisioningFlowSets()))
        for i, v := range m.GetOutboundProvisioningFlowSets() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("outboundProvisioningFlowSets", cast)
        if err != nil {
            return err
        }
    }
    if m.GetReferenceDefinitions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetReferenceDefinitions()))
        for i, v := range m.GetReferenceDefinitions() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("referenceDefinitions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRoleGroups() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRoleGroups()))
        for i, v := range m.GetRoleGroups() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("roleGroups", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRuns() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRuns()))
        for i, v := range m.GetRuns() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("runs", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSourceSystems() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSourceSystems()))
        for i, v := range m.GetSourceSystems() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("sourceSystems", cast)
        if err != nil {
            return err
        }
    }
    if m.GetYears() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetYears()))
        for i, v := range m.GetYears() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("years", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDataConnectors sets the dataConnectors property value. Set of connectors for importing data from source systems.
func (m *IndustryDataRoot) SetDataConnectors(value []IndustryDataConnectorable)() {
    err := m.GetBackingStore().Set("dataConnectors", value)
    if err != nil {
        panic(err)
    }
}
// SetInboundFlows sets the inboundFlows property value. Set of data import flow activities to bring data into the canonical store via a connector.
func (m *IndustryDataRoot) SetInboundFlows(value []InboundFlowable)() {
    err := m.GetBackingStore().Set("inboundFlows", value)
    if err != nil {
        panic(err)
    }
}
// SetOperations sets the operations property value. Set of ephemeral operations that the system runs currently. Read-only.
func (m *IndustryDataRoot) SetOperations(value []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.LongRunningOperationable)() {
    err := m.GetBackingStore().Set("operations", value)
    if err != nil {
        panic(err)
    }
}
// SetOutboundProvisioningFlowSets sets the outboundProvisioningFlowSets property value. The outboundProvisioningFlowSets property
func (m *IndustryDataRoot) SetOutboundProvisioningFlowSets(value []OutboundProvisioningFlowSetable)() {
    err := m.GetBackingStore().Set("outboundProvisioningFlowSets", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceDefinitions sets the referenceDefinitions property value. Set of user modifiable system picker types.
func (m *IndustryDataRoot) SetReferenceDefinitions(value []ReferenceDefinitionable)() {
    err := m.GetBackingStore().Set("referenceDefinitions", value)
    if err != nil {
        panic(err)
    }
}
// SetRoleGroups sets the roleGroups property value. Set of groups of individual roles that makes role-based admin simpler.
func (m *IndustryDataRoot) SetRoleGroups(value []RoleGroupable)() {
    err := m.GetBackingStore().Set("roleGroups", value)
    if err != nil {
        panic(err)
    }
}
// SetRuns sets the runs property value. Set of ephemeral runs which present the point-in-time that diagnostic state of activities performed by the system. Read-only.
func (m *IndustryDataRoot) SetRuns(value []IndustryDataRunable)() {
    err := m.GetBackingStore().Set("runs", value)
    if err != nil {
        panic(err)
    }
}
// SetSourceSystems sets the sourceSystems property value. Set of source definitions that represents real-world external systems.
func (m *IndustryDataRoot) SetSourceSystems(value []SourceSystemDefinitionable)() {
    err := m.GetBackingStore().Set("sourceSystems", value)
    if err != nil {
        panic(err)
    }
}
// SetYears sets the years property value. Set of years represented in the system.
func (m *IndustryDataRoot) SetYears(value []YearTimePeriodDefinitionable)() {
    err := m.GetBackingStore().Set("years", value)
    if err != nil {
        panic(err)
    }
}
type IndustryDataRootable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDataConnectors()([]IndustryDataConnectorable)
    GetInboundFlows()([]InboundFlowable)
    GetOperations()([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.LongRunningOperationable)
    GetOutboundProvisioningFlowSets()([]OutboundProvisioningFlowSetable)
    GetReferenceDefinitions()([]ReferenceDefinitionable)
    GetRoleGroups()([]RoleGroupable)
    GetRuns()([]IndustryDataRunable)
    GetSourceSystems()([]SourceSystemDefinitionable)
    GetYears()([]YearTimePeriodDefinitionable)
    SetDataConnectors(value []IndustryDataConnectorable)()
    SetInboundFlows(value []InboundFlowable)()
    SetOperations(value []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.LongRunningOperationable)()
    SetOutboundProvisioningFlowSets(value []OutboundProvisioningFlowSetable)()
    SetReferenceDefinitions(value []ReferenceDefinitionable)()
    SetRoleGroups(value []RoleGroupable)()
    SetRuns(value []IndustryDataRunable)()
    SetSourceSystems(value []SourceSystemDefinitionable)()
    SetYears(value []YearTimePeriodDefinitionable)()
}

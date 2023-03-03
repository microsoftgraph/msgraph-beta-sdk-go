package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// PlannerTaskRoleBasedRule 
type PlannerTaskRoleBasedRule struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewPlannerTaskRoleBasedRule instantiates a new plannerTaskRoleBasedRule and sets the default values.
func NewPlannerTaskRoleBasedRule()(*PlannerTaskRoleBasedRule) {
    m := &PlannerTaskRoleBasedRule{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePlannerTaskRoleBasedRuleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePlannerTaskRoleBasedRuleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPlannerTaskRoleBasedRule(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PlannerTaskRoleBasedRule) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetBackingStore gets the backingStore property value. Stores model information.
func (m *PlannerTaskRoleBasedRule) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDefaultRule gets the defaultRule property value. Default rule that applies when a property or action-specific rule is not provided. Possible values are: Allow, Block
func (m *PlannerTaskRoleBasedRule) GetDefaultRule()(*string) {
    val, err := m.GetBackingStore().Get("defaultRule")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PlannerTaskRoleBasedRule) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["defaultRule"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultRule(val)
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
    res["propertyRule"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePlannerTaskPropertyRuleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPropertyRule(val.(PlannerTaskPropertyRuleable))
        }
        return nil
    }
    res["role"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePlannerTaskConfigurationRoleBaseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRole(val.(PlannerTaskConfigurationRoleBaseable))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *PlannerTaskRoleBasedRule) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPropertyRule gets the propertyRule property value. Rules for specific properties and actions.
func (m *PlannerTaskRoleBasedRule) GetPropertyRule()(PlannerTaskPropertyRuleable) {
    val, err := m.GetBackingStore().Get("propertyRule")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(PlannerTaskPropertyRuleable)
    }
    return nil
}
// GetRole gets the role property value. The role these rules apply to.
func (m *PlannerTaskRoleBasedRule) GetRole()(PlannerTaskConfigurationRoleBaseable) {
    val, err := m.GetBackingStore().Get("role")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(PlannerTaskConfigurationRoleBaseable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *PlannerTaskRoleBasedRule) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("defaultRule", m.GetDefaultRule())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("propertyRule", m.GetPropertyRule())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("role", m.GetRole())
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
func (m *PlannerTaskRoleBasedRule) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *PlannerTaskRoleBasedRule) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDefaultRule sets the defaultRule property value. Default rule that applies when a property or action-specific rule is not provided. Possible values are: Allow, Block
func (m *PlannerTaskRoleBasedRule) SetDefaultRule(value *string)() {
    err := m.GetBackingStore().Set("defaultRule", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *PlannerTaskRoleBasedRule) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetPropertyRule sets the propertyRule property value. Rules for specific properties and actions.
func (m *PlannerTaskRoleBasedRule) SetPropertyRule(value PlannerTaskPropertyRuleable)() {
    err := m.GetBackingStore().Set("propertyRule", value)
    if err != nil {
        panic(err)
    }
}
// SetRole sets the role property value. The role these rules apply to.
func (m *PlannerTaskRoleBasedRule) SetRole(value PlannerTaskConfigurationRoleBaseable)() {
    err := m.GetBackingStore().Set("role", value)
    if err != nil {
        panic(err)
    }
}
// PlannerTaskRoleBasedRuleable 
type PlannerTaskRoleBasedRuleable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetDefaultRule()(*string)
    GetOdataType()(*string)
    GetPropertyRule()(PlannerTaskPropertyRuleable)
    GetRole()(PlannerTaskConfigurationRoleBaseable)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetDefaultRule(value *string)()
    SetOdataType(value *string)()
    SetPropertyRule(value PlannerTaskPropertyRuleable)()
    SetRole(value PlannerTaskConfigurationRoleBaseable)()
}

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PlannerPlanPropertyRule 
type PlannerPlanPropertyRule struct {
    PlannerPropertyRule
}
// NewPlannerPlanPropertyRule instantiates a new PlannerPlanPropertyRule and sets the default values.
func NewPlannerPlanPropertyRule()(*PlannerPlanPropertyRule) {
    m := &PlannerPlanPropertyRule{
        PlannerPropertyRule: *NewPlannerPropertyRule(),
    }
    odataTypeValue := "#microsoft.graph.plannerPlanPropertyRule"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreatePlannerPlanPropertyRuleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePlannerPlanPropertyRuleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPlannerPlanPropertyRule(), nil
}
// GetBuckets gets the buckets property value. The buckets property
func (m *PlannerPlanPropertyRule) GetBuckets()([]string) {
    val, err := m.GetBackingStore().Get("buckets")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetCategoryDescriptions gets the categoryDescriptions property value. The categoryDescriptions property
func (m *PlannerPlanPropertyRule) GetCategoryDescriptions()(PlannerFieldRulesable) {
    val, err := m.GetBackingStore().Get("categoryDescriptions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(PlannerFieldRulesable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PlannerPlanPropertyRule) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.PlannerPropertyRule.GetFieldDeserializers()
    res["buckets"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetBuckets(res)
        }
        return nil
    }
    res["categoryDescriptions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePlannerFieldRulesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategoryDescriptions(val.(PlannerFieldRulesable))
        }
        return nil
    }
    res["tasks"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetTasks(res)
        }
        return nil
    }
    res["title"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePlannerFieldRulesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTitle(val.(PlannerFieldRulesable))
        }
        return nil
    }
    return res
}
// GetTasks gets the tasks property value. The tasks property
func (m *PlannerPlanPropertyRule) GetTasks()([]string) {
    val, err := m.GetBackingStore().Get("tasks")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetTitle gets the title property value. The title property
func (m *PlannerPlanPropertyRule) GetTitle()(PlannerFieldRulesable) {
    val, err := m.GetBackingStore().Get("title")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(PlannerFieldRulesable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *PlannerPlanPropertyRule) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.PlannerPropertyRule.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetBuckets() != nil {
        err = writer.WriteCollectionOfStringValues("buckets", m.GetBuckets())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("categoryDescriptions", m.GetCategoryDescriptions())
        if err != nil {
            return err
        }
    }
    if m.GetTasks() != nil {
        err = writer.WriteCollectionOfStringValues("tasks", m.GetTasks())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("title", m.GetTitle())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBuckets sets the buckets property value. The buckets property
func (m *PlannerPlanPropertyRule) SetBuckets(value []string)() {
    err := m.GetBackingStore().Set("buckets", value)
    if err != nil {
        panic(err)
    }
}
// SetCategoryDescriptions sets the categoryDescriptions property value. The categoryDescriptions property
func (m *PlannerPlanPropertyRule) SetCategoryDescriptions(value PlannerFieldRulesable)() {
    err := m.GetBackingStore().Set("categoryDescriptions", value)
    if err != nil {
        panic(err)
    }
}
// SetTasks sets the tasks property value. The tasks property
func (m *PlannerPlanPropertyRule) SetTasks(value []string)() {
    err := m.GetBackingStore().Set("tasks", value)
    if err != nil {
        panic(err)
    }
}
// SetTitle sets the title property value. The title property
func (m *PlannerPlanPropertyRule) SetTitle(value PlannerFieldRulesable)() {
    err := m.GetBackingStore().Set("title", value)
    if err != nil {
        panic(err)
    }
}
// PlannerPlanPropertyRuleable 
type PlannerPlanPropertyRuleable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    PlannerPropertyRuleable
    GetBuckets()([]string)
    GetCategoryDescriptions()(PlannerFieldRulesable)
    GetTasks()([]string)
    GetTitle()(PlannerFieldRulesable)
    SetBuckets(value []string)()
    SetCategoryDescriptions(value PlannerFieldRulesable)()
    SetTasks(value []string)()
    SetTitle(value PlannerFieldRulesable)()
}

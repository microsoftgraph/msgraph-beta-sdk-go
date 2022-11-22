package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PlannerPlanPropertyRule 
type PlannerPlanPropertyRule struct {
    PlannerPropertyRule
    // The buckets property
    buckets []string
    // The categoryDescriptions property
    categoryDescriptions PlannerFieldRulesable
    // The tasks property
    tasks []string
    // The title property
    title PlannerFieldRulesable
}
// NewPlannerPlanPropertyRule instantiates a new PlannerPlanPropertyRule and sets the default values.
func NewPlannerPlanPropertyRule()(*PlannerPlanPropertyRule) {
    m := &PlannerPlanPropertyRule{
        PlannerPropertyRule: *NewPlannerPropertyRule(),
    }
    odataTypeValue := "#microsoft.graph.plannerPlanPropertyRule";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreatePlannerPlanPropertyRuleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePlannerPlanPropertyRuleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPlannerPlanPropertyRule(), nil
}
// GetBuckets gets the buckets property value. The buckets property
func (m *PlannerPlanPropertyRule) GetBuckets()([]string) {
    return m.buckets
}
// GetCategoryDescriptions gets the categoryDescriptions property value. The categoryDescriptions property
func (m *PlannerPlanPropertyRule) GetCategoryDescriptions()(PlannerFieldRulesable) {
    return m.categoryDescriptions
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PlannerPlanPropertyRule) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.PlannerPropertyRule.GetFieldDeserializers()
    res["buckets"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetBuckets)
    res["categoryDescriptions"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreatePlannerFieldRulesFromDiscriminatorValue , m.SetCategoryDescriptions)
    res["tasks"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetTasks)
    res["title"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreatePlannerFieldRulesFromDiscriminatorValue , m.SetTitle)
    return res
}
// GetTasks gets the tasks property value. The tasks property
func (m *PlannerPlanPropertyRule) GetTasks()([]string) {
    return m.tasks
}
// GetTitle gets the title property value. The title property
func (m *PlannerPlanPropertyRule) GetTitle()(PlannerFieldRulesable) {
    return m.title
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
    m.buckets = value
}
// SetCategoryDescriptions sets the categoryDescriptions property value. The categoryDescriptions property
func (m *PlannerPlanPropertyRule) SetCategoryDescriptions(value PlannerFieldRulesable)() {
    m.categoryDescriptions = value
}
// SetTasks sets the tasks property value. The tasks property
func (m *PlannerPlanPropertyRule) SetTasks(value []string)() {
    m.tasks = value
}
// SetTitle sets the title property value. The title property
func (m *PlannerPlanPropertyRule) SetTitle(value PlannerFieldRulesable)() {
    m.title = value
}

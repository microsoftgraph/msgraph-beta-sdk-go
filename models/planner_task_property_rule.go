package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PlannerTaskPropertyRule 
type PlannerTaskPropertyRule struct {
    PlannerPropertyRule
    // The appliedCategories property
    appliedCategories PlannerFieldRulesable
    // The assignments property
    assignments PlannerFieldRulesable
    // The checkLists property
    checkLists PlannerFieldRulesable
    // The delete property
    delete []string
    // The dueDate property
    dueDate []string
    // The move property
    move []string
    // The notes property
    notes []string
    // The order property
    order []string
    // The percentComplete property
    percentComplete []string
    // The previewType property
    previewType []string
    // The priority property
    priority []string
    // The references property
    references PlannerFieldRulesable
    // The startDate property
    startDate []string
    // The title property
    title []string
}
// NewPlannerTaskPropertyRule instantiates a new PlannerTaskPropertyRule and sets the default values.
func NewPlannerTaskPropertyRule()(*PlannerTaskPropertyRule) {
    m := &PlannerTaskPropertyRule{
        PlannerPropertyRule: *NewPlannerPropertyRule(),
    }
    odataTypeValue := "#microsoft.graph.plannerTaskPropertyRule";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreatePlannerTaskPropertyRuleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePlannerTaskPropertyRuleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPlannerTaskPropertyRule(), nil
}
// GetAppliedCategories gets the appliedCategories property value. The appliedCategories property
func (m *PlannerTaskPropertyRule) GetAppliedCategories()(PlannerFieldRulesable) {
    return m.appliedCategories
}
// GetAssignments gets the assignments property value. The assignments property
func (m *PlannerTaskPropertyRule) GetAssignments()(PlannerFieldRulesable) {
    return m.assignments
}
// GetCheckLists gets the checkLists property value. The checkLists property
func (m *PlannerTaskPropertyRule) GetCheckLists()(PlannerFieldRulesable) {
    return m.checkLists
}
// GetDelete gets the delete property value. The delete property
func (m *PlannerTaskPropertyRule) GetDelete()([]string) {
    return m.delete
}
// GetDueDate gets the dueDate property value. The dueDate property
func (m *PlannerTaskPropertyRule) GetDueDate()([]string) {
    return m.dueDate
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PlannerTaskPropertyRule) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.PlannerPropertyRule.GetFieldDeserializers()
    res["appliedCategories"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreatePlannerFieldRulesFromDiscriminatorValue , m.SetAppliedCategories)
    res["assignments"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreatePlannerFieldRulesFromDiscriminatorValue , m.SetAssignments)
    res["checkLists"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreatePlannerFieldRulesFromDiscriminatorValue , m.SetCheckLists)
    res["delete"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetDelete)
    res["dueDate"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetDueDate)
    res["move"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetMove)
    res["notes"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetNotes)
    res["order"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetOrder)
    res["percentComplete"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetPercentComplete)
    res["previewType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetPreviewType)
    res["priority"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetPriority)
    res["references"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreatePlannerFieldRulesFromDiscriminatorValue , m.SetReferences)
    res["startDate"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetStartDate)
    res["title"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetTitle)
    return res
}
// GetMove gets the move property value. The move property
func (m *PlannerTaskPropertyRule) GetMove()([]string) {
    return m.move
}
// GetNotes gets the notes property value. The notes property
func (m *PlannerTaskPropertyRule) GetNotes()([]string) {
    return m.notes
}
// GetOrder gets the order property value. The order property
func (m *PlannerTaskPropertyRule) GetOrder()([]string) {
    return m.order
}
// GetPercentComplete gets the percentComplete property value. The percentComplete property
func (m *PlannerTaskPropertyRule) GetPercentComplete()([]string) {
    return m.percentComplete
}
// GetPreviewType gets the previewType property value. The previewType property
func (m *PlannerTaskPropertyRule) GetPreviewType()([]string) {
    return m.previewType
}
// GetPriority gets the priority property value. The priority property
func (m *PlannerTaskPropertyRule) GetPriority()([]string) {
    return m.priority
}
// GetReferences gets the references property value. The references property
func (m *PlannerTaskPropertyRule) GetReferences()(PlannerFieldRulesable) {
    return m.references
}
// GetStartDate gets the startDate property value. The startDate property
func (m *PlannerTaskPropertyRule) GetStartDate()([]string) {
    return m.startDate
}
// GetTitle gets the title property value. The title property
func (m *PlannerTaskPropertyRule) GetTitle()([]string) {
    return m.title
}
// Serialize serializes information the current object
func (m *PlannerTaskPropertyRule) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.PlannerPropertyRule.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("appliedCategories", m.GetAppliedCategories())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("assignments", m.GetAssignments())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("checkLists", m.GetCheckLists())
        if err != nil {
            return err
        }
    }
    if m.GetDelete() != nil {
        err = writer.WriteCollectionOfStringValues("delete", m.GetDelete())
        if err != nil {
            return err
        }
    }
    if m.GetDueDate() != nil {
        err = writer.WriteCollectionOfStringValues("dueDate", m.GetDueDate())
        if err != nil {
            return err
        }
    }
    if m.GetMove() != nil {
        err = writer.WriteCollectionOfStringValues("move", m.GetMove())
        if err != nil {
            return err
        }
    }
    if m.GetNotes() != nil {
        err = writer.WriteCollectionOfStringValues("notes", m.GetNotes())
        if err != nil {
            return err
        }
    }
    if m.GetOrder() != nil {
        err = writer.WriteCollectionOfStringValues("order", m.GetOrder())
        if err != nil {
            return err
        }
    }
    if m.GetPercentComplete() != nil {
        err = writer.WriteCollectionOfStringValues("percentComplete", m.GetPercentComplete())
        if err != nil {
            return err
        }
    }
    if m.GetPreviewType() != nil {
        err = writer.WriteCollectionOfStringValues("previewType", m.GetPreviewType())
        if err != nil {
            return err
        }
    }
    if m.GetPriority() != nil {
        err = writer.WriteCollectionOfStringValues("priority", m.GetPriority())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("references", m.GetReferences())
        if err != nil {
            return err
        }
    }
    if m.GetStartDate() != nil {
        err = writer.WriteCollectionOfStringValues("startDate", m.GetStartDate())
        if err != nil {
            return err
        }
    }
    if m.GetTitle() != nil {
        err = writer.WriteCollectionOfStringValues("title", m.GetTitle())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAppliedCategories sets the appliedCategories property value. The appliedCategories property
func (m *PlannerTaskPropertyRule) SetAppliedCategories(value PlannerFieldRulesable)() {
    m.appliedCategories = value
}
// SetAssignments sets the assignments property value. The assignments property
func (m *PlannerTaskPropertyRule) SetAssignments(value PlannerFieldRulesable)() {
    m.assignments = value
}
// SetCheckLists sets the checkLists property value. The checkLists property
func (m *PlannerTaskPropertyRule) SetCheckLists(value PlannerFieldRulesable)() {
    m.checkLists = value
}
// SetDelete sets the delete property value. The delete property
func (m *PlannerTaskPropertyRule) SetDelete(value []string)() {
    m.delete = value
}
// SetDueDate sets the dueDate property value. The dueDate property
func (m *PlannerTaskPropertyRule) SetDueDate(value []string)() {
    m.dueDate = value
}
// SetMove sets the move property value. The move property
func (m *PlannerTaskPropertyRule) SetMove(value []string)() {
    m.move = value
}
// SetNotes sets the notes property value. The notes property
func (m *PlannerTaskPropertyRule) SetNotes(value []string)() {
    m.notes = value
}
// SetOrder sets the order property value. The order property
func (m *PlannerTaskPropertyRule) SetOrder(value []string)() {
    m.order = value
}
// SetPercentComplete sets the percentComplete property value. The percentComplete property
func (m *PlannerTaskPropertyRule) SetPercentComplete(value []string)() {
    m.percentComplete = value
}
// SetPreviewType sets the previewType property value. The previewType property
func (m *PlannerTaskPropertyRule) SetPreviewType(value []string)() {
    m.previewType = value
}
// SetPriority sets the priority property value. The priority property
func (m *PlannerTaskPropertyRule) SetPriority(value []string)() {
    m.priority = value
}
// SetReferences sets the references property value. The references property
func (m *PlannerTaskPropertyRule) SetReferences(value PlannerFieldRulesable)() {
    m.references = value
}
// SetStartDate sets the startDate property value. The startDate property
func (m *PlannerTaskPropertyRule) SetStartDate(value []string)() {
    m.startDate = value
}
// SetTitle sets the title property value. The title property
func (m *PlannerTaskPropertyRule) SetTitle(value []string)() {
    m.title = value
}

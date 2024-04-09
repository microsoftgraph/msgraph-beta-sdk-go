package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type PlannerTaskPropertyRule struct {
    PlannerPropertyRule
}
// NewPlannerTaskPropertyRule instantiates a new PlannerTaskPropertyRule and sets the default values.
func NewPlannerTaskPropertyRule()(*PlannerTaskPropertyRule) {
    m := &PlannerTaskPropertyRule{
        PlannerPropertyRule: *NewPlannerPropertyRule(),
    }
    odataTypeValue := "#microsoft.graph.plannerTaskPropertyRule"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreatePlannerTaskPropertyRuleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePlannerTaskPropertyRuleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPlannerTaskPropertyRule(), nil
}
// GetAppliedCategories gets the appliedCategories property value. Rules and restrictions for applied categories. This value does not currently support overrides. Accepted values for the default rule and individual overrides are allow, block.
// returns a PlannerFieldRulesable when successful
func (m *PlannerTaskPropertyRule) GetAppliedCategories()(PlannerFieldRulesable) {
    val, err := m.GetBackingStore().Get("appliedCategories")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(PlannerFieldRulesable)
    }
    return nil
}
// GetApprovalAttachment gets the approvalAttachment property value. The approvalAttachment property
// returns a PlannerFieldRulesable when successful
func (m *PlannerTaskPropertyRule) GetApprovalAttachment()(PlannerFieldRulesable) {
    val, err := m.GetBackingStore().Get("approvalAttachment")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(PlannerFieldRulesable)
    }
    return nil
}
// GetAssignments gets the assignments property value. Rules and restrictions for assignments. Allowed overrides are userCreated and applicationCreated. Accepted values for the default rule and individual overrides are allow, add, addSelf, addOther, remove, removeSelf, removeOther, block.
// returns a PlannerFieldRulesable when successful
func (m *PlannerTaskPropertyRule) GetAssignments()(PlannerFieldRulesable) {
    val, err := m.GetBackingStore().Get("assignments")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(PlannerFieldRulesable)
    }
    return nil
}
// GetCheckLists gets the checkLists property value. Rules and restrictions for checklist. Allowed overrides are userCreated and applicationCreated. Accepted values for the default rule and individual overrides are allow, add, remove, update, check, reorder, block.
// returns a PlannerFieldRulesable when successful
func (m *PlannerTaskPropertyRule) GetCheckLists()(PlannerFieldRulesable) {
    val, err := m.GetBackingStore().Get("checkLists")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(PlannerFieldRulesable)
    }
    return nil
}
// GetCompletionRequirements gets the completionRequirements property value. Rules and restrictions for completion requirements of the task. Accepted values are allow, add, remove, edit, and block.
// returns a []string when successful
func (m *PlannerTaskPropertyRule) GetCompletionRequirements()([]string) {
    val, err := m.GetBackingStore().Get("completionRequirements")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetDelete gets the delete property value. Rules and restrictions for deleting the task. Accepted values are allow and block.
// returns a []string when successful
func (m *PlannerTaskPropertyRule) GetDelete()([]string) {
    val, err := m.GetBackingStore().Get("delete")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetDueDate gets the dueDate property value. Rules and restrictions for changing the due date of the task. Accepted values are allow and block.
// returns a []string when successful
func (m *PlannerTaskPropertyRule) GetDueDate()([]string) {
    val, err := m.GetBackingStore().Get("dueDate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PlannerTaskPropertyRule) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.PlannerPropertyRule.GetFieldDeserializers()
    res["appliedCategories"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePlannerFieldRulesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppliedCategories(val.(PlannerFieldRulesable))
        }
        return nil
    }
    res["approvalAttachment"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePlannerFieldRulesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApprovalAttachment(val.(PlannerFieldRulesable))
        }
        return nil
    }
    res["assignments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePlannerFieldRulesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignments(val.(PlannerFieldRulesable))
        }
        return nil
    }
    res["checkLists"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePlannerFieldRulesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCheckLists(val.(PlannerFieldRulesable))
        }
        return nil
    }
    res["completionRequirements"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetCompletionRequirements(res)
        }
        return nil
    }
    res["delete"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetDelete(res)
        }
        return nil
    }
    res["dueDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetDueDate(res)
        }
        return nil
    }
    res["forms"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePlannerFieldRulesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetForms(val.(PlannerFieldRulesable))
        }
        return nil
    }
    res["move"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetMove(res)
        }
        return nil
    }
    res["notes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetNotes(res)
        }
        return nil
    }
    res["order"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetOrder(res)
        }
        return nil
    }
    res["percentComplete"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetPercentComplete(res)
        }
        return nil
    }
    res["previewType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetPreviewType(res)
        }
        return nil
    }
    res["priority"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetPriority(res)
        }
        return nil
    }
    res["references"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePlannerFieldRulesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReferences(val.(PlannerFieldRulesable))
        }
        return nil
    }
    res["startDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetStartDate(res)
        }
        return nil
    }
    res["title"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetTitle(res)
        }
        return nil
    }
    return res
}
// GetForms gets the forms property value. Rules and restrictions for forms. Allowed overrides are userCreated and applicationCreated. The following are the accepted values for the default rule and individual overrides: allow, add, addResponse, remove, update, block.
// returns a PlannerFieldRulesable when successful
func (m *PlannerTaskPropertyRule) GetForms()(PlannerFieldRulesable) {
    val, err := m.GetBackingStore().Get("forms")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(PlannerFieldRulesable)
    }
    return nil
}
// GetMove gets the move property value. Rules and restrictions for moving the task between buckets or plans. Accepted values are allow, moveBetweenPlans, moveBetweenBuckets, and block.
// returns a []string when successful
func (m *PlannerTaskPropertyRule) GetMove()([]string) {
    val, err := m.GetBackingStore().Get("move")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetNotes gets the notes property value. Rules and restrictions for changing the notes of the task. Accepted values are allow and block.
// returns a []string when successful
func (m *PlannerTaskPropertyRule) GetNotes()([]string) {
    val, err := m.GetBackingStore().Get("notes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetOrder gets the order property value. Rules and restrictions for changing the order of the task. Accepted values are allow and block.
// returns a []string when successful
func (m *PlannerTaskPropertyRule) GetOrder()([]string) {
    val, err := m.GetBackingStore().Get("order")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetPercentComplete gets the percentComplete property value. Rules and restrictions for changing the completion percentage of the task. Accepted values are allow, setToComplete, overrideRequirements, setToNotStarted, setToInProgress, and block.
// returns a []string when successful
func (m *PlannerTaskPropertyRule) GetPercentComplete()([]string) {
    val, err := m.GetBackingStore().Get("percentComplete")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetPreviewType gets the previewType property value. Rules and restrictions for changing the preview type of the task. Accepted values are allow and block.
// returns a []string when successful
func (m *PlannerTaskPropertyRule) GetPreviewType()([]string) {
    val, err := m.GetBackingStore().Get("previewType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetPriority gets the priority property value. Rules and restrictions for changing the priority of the task. Accepted values are allow and block.
// returns a []string when successful
func (m *PlannerTaskPropertyRule) GetPriority()([]string) {
    val, err := m.GetBackingStore().Get("priority")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetReferences gets the references property value. Rules and restrictions for references. Allowed overrides are userCreated and applicationCreated. Accepted values for the default rule and individual overrides are allow, add, remove, block.
// returns a PlannerFieldRulesable when successful
func (m *PlannerTaskPropertyRule) GetReferences()(PlannerFieldRulesable) {
    val, err := m.GetBackingStore().Get("references")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(PlannerFieldRulesable)
    }
    return nil
}
// GetStartDate gets the startDate property value. Rules and restrictions for changing the start date of the task. Accepted values are allow and block.
// returns a []string when successful
func (m *PlannerTaskPropertyRule) GetStartDate()([]string) {
    val, err := m.GetBackingStore().Get("startDate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetTitle gets the title property value. Rules and restrictions for changing the title of the task. Accepted values are allow and block.
// returns a []string when successful
func (m *PlannerTaskPropertyRule) GetTitle()([]string) {
    val, err := m.GetBackingStore().Get("title")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
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
        err = writer.WriteObjectValue("approvalAttachment", m.GetApprovalAttachment())
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
    if m.GetCompletionRequirements() != nil {
        err = writer.WriteCollectionOfStringValues("completionRequirements", m.GetCompletionRequirements())
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
    {
        err = writer.WriteObjectValue("forms", m.GetForms())
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
// SetAppliedCategories sets the appliedCategories property value. Rules and restrictions for applied categories. This value does not currently support overrides. Accepted values for the default rule and individual overrides are allow, block.
func (m *PlannerTaskPropertyRule) SetAppliedCategories(value PlannerFieldRulesable)() {
    err := m.GetBackingStore().Set("appliedCategories", value)
    if err != nil {
        panic(err)
    }
}
// SetApprovalAttachment sets the approvalAttachment property value. The approvalAttachment property
func (m *PlannerTaskPropertyRule) SetApprovalAttachment(value PlannerFieldRulesable)() {
    err := m.GetBackingStore().Set("approvalAttachment", value)
    if err != nil {
        panic(err)
    }
}
// SetAssignments sets the assignments property value. Rules and restrictions for assignments. Allowed overrides are userCreated and applicationCreated. Accepted values for the default rule and individual overrides are allow, add, addSelf, addOther, remove, removeSelf, removeOther, block.
func (m *PlannerTaskPropertyRule) SetAssignments(value PlannerFieldRulesable)() {
    err := m.GetBackingStore().Set("assignments", value)
    if err != nil {
        panic(err)
    }
}
// SetCheckLists sets the checkLists property value. Rules and restrictions for checklist. Allowed overrides are userCreated and applicationCreated. Accepted values for the default rule and individual overrides are allow, add, remove, update, check, reorder, block.
func (m *PlannerTaskPropertyRule) SetCheckLists(value PlannerFieldRulesable)() {
    err := m.GetBackingStore().Set("checkLists", value)
    if err != nil {
        panic(err)
    }
}
// SetCompletionRequirements sets the completionRequirements property value. Rules and restrictions for completion requirements of the task. Accepted values are allow, add, remove, edit, and block.
func (m *PlannerTaskPropertyRule) SetCompletionRequirements(value []string)() {
    err := m.GetBackingStore().Set("completionRequirements", value)
    if err != nil {
        panic(err)
    }
}
// SetDelete sets the delete property value. Rules and restrictions for deleting the task. Accepted values are allow and block.
func (m *PlannerTaskPropertyRule) SetDelete(value []string)() {
    err := m.GetBackingStore().Set("delete", value)
    if err != nil {
        panic(err)
    }
}
// SetDueDate sets the dueDate property value. Rules and restrictions for changing the due date of the task. Accepted values are allow and block.
func (m *PlannerTaskPropertyRule) SetDueDate(value []string)() {
    err := m.GetBackingStore().Set("dueDate", value)
    if err != nil {
        panic(err)
    }
}
// SetForms sets the forms property value. Rules and restrictions for forms. Allowed overrides are userCreated and applicationCreated. The following are the accepted values for the default rule and individual overrides: allow, add, addResponse, remove, update, block.
func (m *PlannerTaskPropertyRule) SetForms(value PlannerFieldRulesable)() {
    err := m.GetBackingStore().Set("forms", value)
    if err != nil {
        panic(err)
    }
}
// SetMove sets the move property value. Rules and restrictions for moving the task between buckets or plans. Accepted values are allow, moveBetweenPlans, moveBetweenBuckets, and block.
func (m *PlannerTaskPropertyRule) SetMove(value []string)() {
    err := m.GetBackingStore().Set("move", value)
    if err != nil {
        panic(err)
    }
}
// SetNotes sets the notes property value. Rules and restrictions for changing the notes of the task. Accepted values are allow and block.
func (m *PlannerTaskPropertyRule) SetNotes(value []string)() {
    err := m.GetBackingStore().Set("notes", value)
    if err != nil {
        panic(err)
    }
}
// SetOrder sets the order property value. Rules and restrictions for changing the order of the task. Accepted values are allow and block.
func (m *PlannerTaskPropertyRule) SetOrder(value []string)() {
    err := m.GetBackingStore().Set("order", value)
    if err != nil {
        panic(err)
    }
}
// SetPercentComplete sets the percentComplete property value. Rules and restrictions for changing the completion percentage of the task. Accepted values are allow, setToComplete, overrideRequirements, setToNotStarted, setToInProgress, and block.
func (m *PlannerTaskPropertyRule) SetPercentComplete(value []string)() {
    err := m.GetBackingStore().Set("percentComplete", value)
    if err != nil {
        panic(err)
    }
}
// SetPreviewType sets the previewType property value. Rules and restrictions for changing the preview type of the task. Accepted values are allow and block.
func (m *PlannerTaskPropertyRule) SetPreviewType(value []string)() {
    err := m.GetBackingStore().Set("previewType", value)
    if err != nil {
        panic(err)
    }
}
// SetPriority sets the priority property value. Rules and restrictions for changing the priority of the task. Accepted values are allow and block.
func (m *PlannerTaskPropertyRule) SetPriority(value []string)() {
    err := m.GetBackingStore().Set("priority", value)
    if err != nil {
        panic(err)
    }
}
// SetReferences sets the references property value. Rules and restrictions for references. Allowed overrides are userCreated and applicationCreated. Accepted values for the default rule and individual overrides are allow, add, remove, block.
func (m *PlannerTaskPropertyRule) SetReferences(value PlannerFieldRulesable)() {
    err := m.GetBackingStore().Set("references", value)
    if err != nil {
        panic(err)
    }
}
// SetStartDate sets the startDate property value. Rules and restrictions for changing the start date of the task. Accepted values are allow and block.
func (m *PlannerTaskPropertyRule) SetStartDate(value []string)() {
    err := m.GetBackingStore().Set("startDate", value)
    if err != nil {
        panic(err)
    }
}
// SetTitle sets the title property value. Rules and restrictions for changing the title of the task. Accepted values are allow and block.
func (m *PlannerTaskPropertyRule) SetTitle(value []string)() {
    err := m.GetBackingStore().Set("title", value)
    if err != nil {
        panic(err)
    }
}
type PlannerTaskPropertyRuleable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    PlannerPropertyRuleable
    GetAppliedCategories()(PlannerFieldRulesable)
    GetApprovalAttachment()(PlannerFieldRulesable)
    GetAssignments()(PlannerFieldRulesable)
    GetCheckLists()(PlannerFieldRulesable)
    GetCompletionRequirements()([]string)
    GetDelete()([]string)
    GetDueDate()([]string)
    GetForms()(PlannerFieldRulesable)
    GetMove()([]string)
    GetNotes()([]string)
    GetOrder()([]string)
    GetPercentComplete()([]string)
    GetPreviewType()([]string)
    GetPriority()([]string)
    GetReferences()(PlannerFieldRulesable)
    GetStartDate()([]string)
    GetTitle()([]string)
    SetAppliedCategories(value PlannerFieldRulesable)()
    SetApprovalAttachment(value PlannerFieldRulesable)()
    SetAssignments(value PlannerFieldRulesable)()
    SetCheckLists(value PlannerFieldRulesable)()
    SetCompletionRequirements(value []string)()
    SetDelete(value []string)()
    SetDueDate(value []string)()
    SetForms(value PlannerFieldRulesable)()
    SetMove(value []string)()
    SetNotes(value []string)()
    SetOrder(value []string)()
    SetPercentComplete(value []string)()
    SetPreviewType(value []string)()
    SetPriority(value []string)()
    SetReferences(value PlannerFieldRulesable)()
    SetStartDate(value []string)()
    SetTitle(value []string)()
}

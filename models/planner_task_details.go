package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PlannerTaskDetails 
type PlannerTaskDetails struct {
    PlannerDelta
}
// NewPlannerTaskDetails instantiates a new plannerTaskDetails and sets the default values.
func NewPlannerTaskDetails()(*PlannerTaskDetails) {
    m := &PlannerTaskDetails{
        PlannerDelta: *NewPlannerDelta(),
    }
    return m
}
// CreatePlannerTaskDetailsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePlannerTaskDetailsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPlannerTaskDetails(), nil
}
// GetApprovalAttachment gets the approvalAttachment property value. The approvalAttachment property
func (m *PlannerTaskDetails) GetApprovalAttachment()(PlannerBaseApprovalAttachmentable) {
    val, err := m.GetBackingStore().Get("approvalAttachment")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(PlannerBaseApprovalAttachmentable)
    }
    return nil
}
// GetChecklist gets the checklist property value. The collection of checklist items on the task.
func (m *PlannerTaskDetails) GetChecklist()(PlannerChecklistItemsable) {
    val, err := m.GetBackingStore().Get("checklist")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(PlannerChecklistItemsable)
    }
    return nil
}
// GetCompletionRequirements gets the completionRequirements property value. Contains detailed information about requirements on the task.
func (m *PlannerTaskDetails) GetCompletionRequirements()(PlannerTaskCompletionRequirementDetailsable) {
    val, err := m.GetBackingStore().Get("completionRequirements")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(PlannerTaskCompletionRequirementDetailsable)
    }
    return nil
}
// GetDescription gets the description property value. Description of the task.
func (m *PlannerTaskDetails) GetDescription()(*string) {
    val, err := m.GetBackingStore().Get("description")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PlannerTaskDetails) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.PlannerDelta.GetFieldDeserializers()
    res["approvalAttachment"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePlannerBaseApprovalAttachmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApprovalAttachment(val.(PlannerBaseApprovalAttachmentable))
        }
        return nil
    }
    res["checklist"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePlannerChecklistItemsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChecklist(val.(PlannerChecklistItemsable))
        }
        return nil
    }
    res["completionRequirements"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePlannerTaskCompletionRequirementDetailsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompletionRequirements(val.(PlannerTaskCompletionRequirementDetailsable))
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["forms"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePlannerFormsDictionaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetForms(val.(PlannerFormsDictionaryable))
        }
        return nil
    }
    res["notes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemBodyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotes(val.(ItemBodyable))
        }
        return nil
    }
    res["previewType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePlannerPreviewType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreviewType(val.(*PlannerPreviewType))
        }
        return nil
    }
    res["references"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePlannerExternalReferencesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReferences(val.(PlannerExternalReferencesable))
        }
        return nil
    }
    return res
}
// GetForms gets the forms property value. The forms property
func (m *PlannerTaskDetails) GetForms()(PlannerFormsDictionaryable) {
    val, err := m.GetBackingStore().Get("forms")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(PlannerFormsDictionaryable)
    }
    return nil
}
// GetNotes gets the notes property value. Rich text description of the task. To be used by HTML-aware clients. For backwards compatibility, a plain-text version of the HTML description will be synced to the 'description' field. If this field hasn't previously been set but 'description' has been, the existing description is synchronized to 'notes' with minimal whitespace-preserving HTML markup. Setting both 'description' and 'notes' is an error and will result in an exception.
func (m *PlannerTaskDetails) GetNotes()(ItemBodyable) {
    val, err := m.GetBackingStore().Get("notes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ItemBodyable)
    }
    return nil
}
// GetPreviewType gets the previewType property value. This sets the type of preview that shows up on the task. Possible values are: automatic, noPreview, checklist, description, reference. When set to automatic the displayed preview is chosen by the app viewing the task.
func (m *PlannerTaskDetails) GetPreviewType()(*PlannerPreviewType) {
    val, err := m.GetBackingStore().Get("previewType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*PlannerPreviewType)
    }
    return nil
}
// GetReferences gets the references property value. The collection of references on the task.
func (m *PlannerTaskDetails) GetReferences()(PlannerExternalReferencesable) {
    val, err := m.GetBackingStore().Get("references")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(PlannerExternalReferencesable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *PlannerTaskDetails) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.PlannerDelta.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("approvalAttachment", m.GetApprovalAttachment())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("checklist", m.GetChecklist())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("completionRequirements", m.GetCompletionRequirements())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
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
    {
        err = writer.WriteObjectValue("notes", m.GetNotes())
        if err != nil {
            return err
        }
    }
    if m.GetPreviewType() != nil {
        cast := (*m.GetPreviewType()).String()
        err = writer.WriteStringValue("previewType", &cast)
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
    return nil
}
// SetApprovalAttachment sets the approvalAttachment property value. The approvalAttachment property
func (m *PlannerTaskDetails) SetApprovalAttachment(value PlannerBaseApprovalAttachmentable)() {
    err := m.GetBackingStore().Set("approvalAttachment", value)
    if err != nil {
        panic(err)
    }
}
// SetChecklist sets the checklist property value. The collection of checklist items on the task.
func (m *PlannerTaskDetails) SetChecklist(value PlannerChecklistItemsable)() {
    err := m.GetBackingStore().Set("checklist", value)
    if err != nil {
        panic(err)
    }
}
// SetCompletionRequirements sets the completionRequirements property value. Contains detailed information about requirements on the task.
func (m *PlannerTaskDetails) SetCompletionRequirements(value PlannerTaskCompletionRequirementDetailsable)() {
    err := m.GetBackingStore().Set("completionRequirements", value)
    if err != nil {
        panic(err)
    }
}
// SetDescription sets the description property value. Description of the task.
func (m *PlannerTaskDetails) SetDescription(value *string)() {
    err := m.GetBackingStore().Set("description", value)
    if err != nil {
        panic(err)
    }
}
// SetForms sets the forms property value. The forms property
func (m *PlannerTaskDetails) SetForms(value PlannerFormsDictionaryable)() {
    err := m.GetBackingStore().Set("forms", value)
    if err != nil {
        panic(err)
    }
}
// SetNotes sets the notes property value. Rich text description of the task. To be used by HTML-aware clients. For backwards compatibility, a plain-text version of the HTML description will be synced to the 'description' field. If this field hasn't previously been set but 'description' has been, the existing description is synchronized to 'notes' with minimal whitespace-preserving HTML markup. Setting both 'description' and 'notes' is an error and will result in an exception.
func (m *PlannerTaskDetails) SetNotes(value ItemBodyable)() {
    err := m.GetBackingStore().Set("notes", value)
    if err != nil {
        panic(err)
    }
}
// SetPreviewType sets the previewType property value. This sets the type of preview that shows up on the task. Possible values are: automatic, noPreview, checklist, description, reference. When set to automatic the displayed preview is chosen by the app viewing the task.
func (m *PlannerTaskDetails) SetPreviewType(value *PlannerPreviewType)() {
    err := m.GetBackingStore().Set("previewType", value)
    if err != nil {
        panic(err)
    }
}
// SetReferences sets the references property value. The collection of references on the task.
func (m *PlannerTaskDetails) SetReferences(value PlannerExternalReferencesable)() {
    err := m.GetBackingStore().Set("references", value)
    if err != nil {
        panic(err)
    }
}
// PlannerTaskDetailsable 
type PlannerTaskDetailsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    PlannerDeltaable
    GetApprovalAttachment()(PlannerBaseApprovalAttachmentable)
    GetChecklist()(PlannerChecklistItemsable)
    GetCompletionRequirements()(PlannerTaskCompletionRequirementDetailsable)
    GetDescription()(*string)
    GetForms()(PlannerFormsDictionaryable)
    GetNotes()(ItemBodyable)
    GetPreviewType()(*PlannerPreviewType)
    GetReferences()(PlannerExternalReferencesable)
    SetApprovalAttachment(value PlannerBaseApprovalAttachmentable)()
    SetChecklist(value PlannerChecklistItemsable)()
    SetCompletionRequirements(value PlannerTaskCompletionRequirementDetailsable)()
    SetDescription(value *string)()
    SetForms(value PlannerFormsDictionaryable)()
    SetNotes(value ItemBodyable)()
    SetPreviewType(value *PlannerPreviewType)()
    SetReferences(value PlannerExternalReferencesable)()
}

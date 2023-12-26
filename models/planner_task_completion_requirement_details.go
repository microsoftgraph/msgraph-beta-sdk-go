package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// PlannerTaskCompletionRequirementDetails 
type PlannerTaskCompletionRequirementDetails struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewPlannerTaskCompletionRequirementDetails instantiates a new plannerTaskCompletionRequirementDetails and sets the default values.
func NewPlannerTaskCompletionRequirementDetails()(*PlannerTaskCompletionRequirementDetails) {
    m := &PlannerTaskCompletionRequirementDetails{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePlannerTaskCompletionRequirementDetailsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePlannerTaskCompletionRequirementDetailsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPlannerTaskCompletionRequirementDetails(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PlannerTaskCompletionRequirementDetails) GetAdditionalData()(map[string]any) {
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
// GetApprovalRequirement gets the approvalRequirement property value. The approvalRequirement property
func (m *PlannerTaskCompletionRequirementDetails) GetApprovalRequirement()(PlannerApprovalRequirementable) {
    val, err := m.GetBackingStore().Get("approvalRequirement")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(PlannerApprovalRequirementable)
    }
    return nil
}
// GetBackingStore gets the BackingStore property value. Stores model information.
func (m *PlannerTaskCompletionRequirementDetails) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetChecklistRequirement gets the checklistRequirement property value. Information about the requirements for completing the checklist.
func (m *PlannerTaskCompletionRequirementDetails) GetChecklistRequirement()(PlannerChecklistRequirementable) {
    val, err := m.GetBackingStore().Get("checklistRequirement")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(PlannerChecklistRequirementable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PlannerTaskCompletionRequirementDetails) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["approvalRequirement"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePlannerApprovalRequirementFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApprovalRequirement(val.(PlannerApprovalRequirementable))
        }
        return nil
    }
    res["checklistRequirement"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePlannerChecklistRequirementFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChecklistRequirement(val.(PlannerChecklistRequirementable))
        }
        return nil
    }
    res["formsRequirement"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePlannerFormsRequirementFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFormsRequirement(val.(PlannerFormsRequirementable))
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
    return res
}
// GetFormsRequirement gets the formsRequirement property value. The formsRequirement property
func (m *PlannerTaskCompletionRequirementDetails) GetFormsRequirement()(PlannerFormsRequirementable) {
    val, err := m.GetBackingStore().Get("formsRequirement")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(PlannerFormsRequirementable)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *PlannerTaskCompletionRequirementDetails) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *PlannerTaskCompletionRequirementDetails) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("approvalRequirement", m.GetApprovalRequirement())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("checklistRequirement", m.GetChecklistRequirement())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("formsRequirement", m.GetFormsRequirement())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PlannerTaskCompletionRequirementDetails) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetApprovalRequirement sets the approvalRequirement property value. The approvalRequirement property
func (m *PlannerTaskCompletionRequirementDetails) SetApprovalRequirement(value PlannerApprovalRequirementable)() {
    err := m.GetBackingStore().Set("approvalRequirement", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *PlannerTaskCompletionRequirementDetails) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetChecklistRequirement sets the checklistRequirement property value. Information about the requirements for completing the checklist.
func (m *PlannerTaskCompletionRequirementDetails) SetChecklistRequirement(value PlannerChecklistRequirementable)() {
    err := m.GetBackingStore().Set("checklistRequirement", value)
    if err != nil {
        panic(err)
    }
}
// SetFormsRequirement sets the formsRequirement property value. The formsRequirement property
func (m *PlannerTaskCompletionRequirementDetails) SetFormsRequirement(value PlannerFormsRequirementable)() {
    err := m.GetBackingStore().Set("formsRequirement", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *PlannerTaskCompletionRequirementDetails) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// PlannerTaskCompletionRequirementDetailsable 
type PlannerTaskCompletionRequirementDetailsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApprovalRequirement()(PlannerApprovalRequirementable)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetChecklistRequirement()(PlannerChecklistRequirementable)
    GetFormsRequirement()(PlannerFormsRequirementable)
    GetOdataType()(*string)
    SetApprovalRequirement(value PlannerApprovalRequirementable)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetChecklistRequirement(value PlannerChecklistRequirementable)()
    SetFormsRequirement(value PlannerFormsRequirementable)()
    SetOdataType(value *string)()
}

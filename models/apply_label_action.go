package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ApplyLabelAction 
type ApplyLabelAction struct {
    InformationProtectionAction
    // The collection of specific actions that should be taken by the consuming application to label the document. See  informationProtectionAction for the full list.
    actions []InformationProtectionActionable
    // The actionSource property
    actionSource *ActionSource
    // Object that describes the details of the label to apply.
    label LabelDetailsable
    // If the label was the result of an automatic classification, supply the list of sensitive info type GUIDs that resulted in the returned label.
    responsibleSensitiveTypeIds []string
}
// NewApplyLabelAction instantiates a new ApplyLabelAction and sets the default values.
func NewApplyLabelAction()(*ApplyLabelAction) {
    m := &ApplyLabelAction{
        InformationProtectionAction: *NewInformationProtectionAction(),
    }
    odataTypeValue := "#microsoft.graph.applyLabelAction";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateApplyLabelActionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateApplyLabelActionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewApplyLabelAction(), nil
}
// GetActions gets the actions property value. The collection of specific actions that should be taken by the consuming application to label the document. See  informationProtectionAction for the full list.
func (m *ApplyLabelAction) GetActions()([]InformationProtectionActionable) {
    return m.actions
}
// GetActionSource gets the actionSource property value. The actionSource property
func (m *ApplyLabelAction) GetActionSource()(*ActionSource) {
    return m.actionSource
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ApplyLabelAction) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.InformationProtectionAction.GetFieldDeserializers()
    res["actions"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateInformationProtectionActionFromDiscriminatorValue , m.SetActions)
    res["actionSource"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseActionSource , m.SetActionSource)
    res["label"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateLabelDetailsFromDiscriminatorValue , m.SetLabel)
    res["responsibleSensitiveTypeIds"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetResponsibleSensitiveTypeIds)
    return res
}
// GetLabel gets the label property value. Object that describes the details of the label to apply.
func (m *ApplyLabelAction) GetLabel()(LabelDetailsable) {
    return m.label
}
// GetResponsibleSensitiveTypeIds gets the responsibleSensitiveTypeIds property value. If the label was the result of an automatic classification, supply the list of sensitive info type GUIDs that resulted in the returned label.
func (m *ApplyLabelAction) GetResponsibleSensitiveTypeIds()([]string) {
    return m.responsibleSensitiveTypeIds
}
// Serialize serializes information the current object
func (m *ApplyLabelAction) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.InformationProtectionAction.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetActions() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetActions())
        err = writer.WriteCollectionOfObjectValues("actions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetActionSource() != nil {
        cast := (*m.GetActionSource()).String()
        err = writer.WriteStringValue("actionSource", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("label", m.GetLabel())
        if err != nil {
            return err
        }
    }
    if m.GetResponsibleSensitiveTypeIds() != nil {
        err = writer.WriteCollectionOfStringValues("responsibleSensitiveTypeIds", m.GetResponsibleSensitiveTypeIds())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActions sets the actions property value. The collection of specific actions that should be taken by the consuming application to label the document. See  informationProtectionAction for the full list.
func (m *ApplyLabelAction) SetActions(value []InformationProtectionActionable)() {
    m.actions = value
}
// SetActionSource sets the actionSource property value. The actionSource property
func (m *ApplyLabelAction) SetActionSource(value *ActionSource)() {
    m.actionSource = value
}
// SetLabel sets the label property value. Object that describes the details of the label to apply.
func (m *ApplyLabelAction) SetLabel(value LabelDetailsable)() {
    m.label = value
}
// SetResponsibleSensitiveTypeIds sets the responsibleSensitiveTypeIds property value. If the label was the result of an automatic classification, supply the list of sensitive info type GUIDs that resulted in the returned label.
func (m *ApplyLabelAction) SetResponsibleSensitiveTypeIds(value []string)() {
    m.responsibleSensitiveTypeIds = value
}

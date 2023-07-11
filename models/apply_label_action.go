package models

import (
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ApplyLabelAction 
type ApplyLabelAction struct {
    InformationProtectionAction
    // The OdataType property
    OdataType *string
}
// NewApplyLabelAction instantiates a new applyLabelAction and sets the default values.
func NewApplyLabelAction()(*ApplyLabelAction) {
    m := &ApplyLabelAction{
        InformationProtectionAction: *NewInformationProtectionAction(),
    }
    odataTypeValue := "#microsoft.graph.applyLabelAction"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateApplyLabelActionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateApplyLabelActionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewApplyLabelAction(), nil
}
// GetActions gets the actions property value. The collection of specific actions that should be taken by the consuming application to label the document. See  informationProtectionAction for the full list.
func (m *ApplyLabelAction) GetActions()([]InformationProtectionActionable) {
    val, err := m.GetBackingStore().Get("actions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]InformationProtectionActionable)
    }
    return nil
}
// GetActionSource gets the actionSource property value. The actionSource property
func (m *ApplyLabelAction) GetActionSource()(*ActionSource) {
    val, err := m.GetBackingStore().Get("actionSource")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ActionSource)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ApplyLabelAction) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.InformationProtectionAction.GetFieldDeserializers()
    res["actions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateInformationProtectionActionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]InformationProtectionActionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(InformationProtectionActionable)
                }
            }
            m.SetActions(res)
        }
        return nil
    }
    res["actionSource"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseActionSource)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActionSource(val.(*ActionSource))
        }
        return nil
    }
    res["label"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateLabelDetailsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLabel(val.(LabelDetailsable))
        }
        return nil
    }
    res["responsibleSensitiveTypeIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("uuid")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID))
                }
            }
            m.SetResponsibleSensitiveTypeIds(res)
        }
        return nil
    }
    return res
}
// GetLabel gets the label property value. Object that describes the details of the label to apply.
func (m *ApplyLabelAction) GetLabel()(LabelDetailsable) {
    val, err := m.GetBackingStore().Get("label")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(LabelDetailsable)
    }
    return nil
}
// GetResponsibleSensitiveTypeIds gets the responsibleSensitiveTypeIds property value. If the label was the result of an automatic classification, supply the list of sensitive info type GUIDs that resulted in the returned label.
func (m *ApplyLabelAction) GetResponsibleSensitiveTypeIds()([]i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    val, err := m.GetBackingStore().Get("responsibleSensitiveTypeIds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ApplyLabelAction) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.InformationProtectionAction.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetActions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetActions()))
        for i, v := range m.GetActions() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
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
        err = writer.WriteCollectionOfUUIDValues("responsibleSensitiveTypeIds", m.GetResponsibleSensitiveTypeIds())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActions sets the actions property value. The collection of specific actions that should be taken by the consuming application to label the document. See  informationProtectionAction for the full list.
func (m *ApplyLabelAction) SetActions(value []InformationProtectionActionable)() {
    err := m.GetBackingStore().Set("actions", value)
    if err != nil {
        panic(err)
    }
}
// SetActionSource sets the actionSource property value. The actionSource property
func (m *ApplyLabelAction) SetActionSource(value *ActionSource)() {
    err := m.GetBackingStore().Set("actionSource", value)
    if err != nil {
        panic(err)
    }
}
// SetLabel sets the label property value. Object that describes the details of the label to apply.
func (m *ApplyLabelAction) SetLabel(value LabelDetailsable)() {
    err := m.GetBackingStore().Set("label", value)
    if err != nil {
        panic(err)
    }
}
// SetResponsibleSensitiveTypeIds sets the responsibleSensitiveTypeIds property value. If the label was the result of an automatic classification, supply the list of sensitive info type GUIDs that resulted in the returned label.
func (m *ApplyLabelAction) SetResponsibleSensitiveTypeIds(value []i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    err := m.GetBackingStore().Set("responsibleSensitiveTypeIds", value)
    if err != nil {
        panic(err)
    }
}
// ApplyLabelActionable 
type ApplyLabelActionable interface {
    InformationProtectionActionable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActions()([]InformationProtectionActionable)
    GetActionSource()(*ActionSource)
    GetLabel()(LabelDetailsable)
    GetResponsibleSensitiveTypeIds()([]i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    SetActions(value []InformationProtectionActionable)()
    SetActionSource(value *ActionSource)()
    SetLabel(value LabelDetailsable)()
    SetResponsibleSensitiveTypeIds(value []i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)()
}

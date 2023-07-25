package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RecommendLabelAction 
type RecommendLabelAction struct {
    InformationProtectionAction
}
// NewRecommendLabelAction instantiates a new recommendLabelAction and sets the default values.
func NewRecommendLabelAction()(*RecommendLabelAction) {
    m := &RecommendLabelAction{
        InformationProtectionAction: *NewInformationProtectionAction(),
    }
    odataTypeValue := "#microsoft.graph.security.recommendLabelAction"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateRecommendLabelActionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRecommendLabelActionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRecommendLabelAction(), nil
}
// GetActions gets the actions property value. Actions to take if the label is accepted by the user.
func (m *RecommendLabelAction) GetActions()([]InformationProtectionActionable) {
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
func (m *RecommendLabelAction) GetActionSource()(*ActionSource) {
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
func (m *RecommendLabelAction) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["responsibleSensitiveTypeIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetResponsibleSensitiveTypeIds(res)
        }
        return nil
    }
    res["sensitivityLabelId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSensitivityLabelId(val)
        }
        return nil
    }
    return res
}
// GetResponsibleSensitiveTypeIds gets the responsibleSensitiveTypeIds property value. The sensitive information type GUIDs that caused the recommendation to be given.
func (m *RecommendLabelAction) GetResponsibleSensitiveTypeIds()([]string) {
    val, err := m.GetBackingStore().Get("responsibleSensitiveTypeIds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetSensitivityLabelId gets the sensitivityLabelId property value. The sensitivityLabelId property
func (m *RecommendLabelAction) GetSensitivityLabelId()(*string) {
    val, err := m.GetBackingStore().Get("sensitivityLabelId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *RecommendLabelAction) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    if m.GetResponsibleSensitiveTypeIds() != nil {
        err = writer.WriteCollectionOfStringValues("responsibleSensitiveTypeIds", m.GetResponsibleSensitiveTypeIds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("sensitivityLabelId", m.GetSensitivityLabelId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActions sets the actions property value. Actions to take if the label is accepted by the user.
func (m *RecommendLabelAction) SetActions(value []InformationProtectionActionable)() {
    err := m.GetBackingStore().Set("actions", value)
    if err != nil {
        panic(err)
    }
}
// SetActionSource sets the actionSource property value. The actionSource property
func (m *RecommendLabelAction) SetActionSource(value *ActionSource)() {
    err := m.GetBackingStore().Set("actionSource", value)
    if err != nil {
        panic(err)
    }
}
// SetResponsibleSensitiveTypeIds sets the responsibleSensitiveTypeIds property value. The sensitive information type GUIDs that caused the recommendation to be given.
func (m *RecommendLabelAction) SetResponsibleSensitiveTypeIds(value []string)() {
    err := m.GetBackingStore().Set("responsibleSensitiveTypeIds", value)
    if err != nil {
        panic(err)
    }
}
// SetSensitivityLabelId sets the sensitivityLabelId property value. The sensitivityLabelId property
func (m *RecommendLabelAction) SetSensitivityLabelId(value *string)() {
    err := m.GetBackingStore().Set("sensitivityLabelId", value)
    if err != nil {
        panic(err)
    }
}
// RecommendLabelActionable 
type RecommendLabelActionable interface {
    InformationProtectionActionable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActions()([]InformationProtectionActionable)
    GetActionSource()(*ActionSource)
    GetResponsibleSensitiveTypeIds()([]string)
    GetSensitivityLabelId()(*string)
    SetActions(value []InformationProtectionActionable)()
    SetActionSource(value *ActionSource)()
    SetResponsibleSensitiveTypeIds(value []string)()
    SetSensitivityLabelId(value *string)()
}

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RecommendLabelAction 
type RecommendLabelAction struct {
    InformationProtectionAction
    // Actions to take if the label is accepted by the user.
    actions []InformationProtectionActionable
    // Possible values are: manual, automatic, recommended, default.
    actionSource *ActionSource
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The label that is being recommended.
    label LabelDetailsable
    // The sensitive information type GUIDs that caused the recommendation to be given.
    responsibleSensitiveTypeIds []string
}
// NewRecommendLabelAction instantiates a new RecommendLabelAction and sets the default values.
func NewRecommendLabelAction()(*RecommendLabelAction) {
    m := &RecommendLabelAction{
        InformationProtectionAction: *NewInformationProtectionAction(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateRecommendLabelActionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRecommendLabelActionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRecommendLabelAction(), nil
}
// GetActions gets the actions property value. Actions to take if the label is accepted by the user.
func (m *RecommendLabelAction) GetActions()([]InformationProtectionActionable) {
    if m == nil {
        return nil
    } else {
        return m.actions
    }
}
// GetActionSource gets the actionSource property value. Possible values are: manual, automatic, recommended, default.
func (m *RecommendLabelAction) GetActionSource()(*ActionSource) {
    if m == nil {
        return nil
    } else {
        return m.actionSource
    }
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RecommendLabelAction) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
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
                res[i] = v.(InformationProtectionActionable)
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
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetResponsibleSensitiveTypeIds(res)
        }
        return nil
    }
    return res
}
// GetLabel gets the label property value. The label that is being recommended.
func (m *RecommendLabelAction) GetLabel()(LabelDetailsable) {
    if m == nil {
        return nil
    } else {
        return m.label
    }
}
// GetResponsibleSensitiveTypeIds gets the responsibleSensitiveTypeIds property value. The sensitive information type GUIDs that caused the recommendation to be given.
func (m *RecommendLabelAction) GetResponsibleSensitiveTypeIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.responsibleSensitiveTypeIds
    }
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
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
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
        err = writer.WriteCollectionOfStringValues("responsibleSensitiveTypeIds", m.GetResponsibleSensitiveTypeIds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActions sets the actions property value. Actions to take if the label is accepted by the user.
func (m *RecommendLabelAction) SetActions(value []InformationProtectionActionable)() {
    if m != nil {
        m.actions = value
    }
}
// SetActionSource sets the actionSource property value. Possible values are: manual, automatic, recommended, default.
func (m *RecommendLabelAction) SetActionSource(value *ActionSource)() {
    if m != nil {
        m.actionSource = value
    }
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RecommendLabelAction) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetLabel sets the label property value. The label that is being recommended.
func (m *RecommendLabelAction) SetLabel(value LabelDetailsable)() {
    if m != nil {
        m.label = value
    }
}
// SetResponsibleSensitiveTypeIds sets the responsibleSensitiveTypeIds property value. The sensitive information type GUIDs that caused the recommendation to be given.
func (m *RecommendLabelAction) SetResponsibleSensitiveTypeIds(value []string)() {
    if m != nil {
        m.responsibleSensitiveTypeIds = value
    }
}

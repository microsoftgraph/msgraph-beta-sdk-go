package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// PlannerTaskCreation 
type PlannerTaskCreation struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewPlannerTaskCreation instantiates a new plannerTaskCreation and sets the default values.
func NewPlannerTaskCreation()(*PlannerTaskCreation) {
    m := &PlannerTaskCreation{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePlannerTaskCreationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePlannerTaskCreationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.plannerExternalTaskSource":
                        return NewPlannerExternalTaskSource(), nil
                    case "#microsoft.graph.plannerTeamsPublicationInfo":
                        return NewPlannerTeamsPublicationInfo(), nil
                }
            }
        }
    }
    return NewPlannerTaskCreation(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PlannerTaskCreation) GetAdditionalData()(map[string]any) {
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
// GetBackingStore gets the BackingStore property value. Stores model information.
func (m *PlannerTaskCreation) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetCreationSourceKind gets the creationSourceKind property value. Specifies what kind of creation source the task is created with. The possible values are: external, publication and unknownFutureValue.
func (m *PlannerTaskCreation) GetCreationSourceKind()(*PlannerTaskCreation_creationSourceKind) {
    val, err := m.GetBackingStore().Get("creationSourceKind")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*PlannerTaskCreation_creationSourceKind)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PlannerTaskCreation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["creationSourceKind"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePlannerTaskCreation_creationSourceKind)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreationSourceKind(val.(*PlannerTaskCreation_creationSourceKind))
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
    res["teamsPublicationInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePlannerTeamsPublicationInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeamsPublicationInfo(val.(PlannerTeamsPublicationInfoable))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *PlannerTaskCreation) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTeamsPublicationInfo gets the teamsPublicationInfo property value. Information about the publication process that created this task. This field is deprecated and clients should move to using the new inheritance model.
func (m *PlannerTaskCreation) GetTeamsPublicationInfo()(PlannerTeamsPublicationInfoable) {
    val, err := m.GetBackingStore().Get("teamsPublicationInfo")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(PlannerTeamsPublicationInfoable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *PlannerTaskCreation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetCreationSourceKind() != nil {
        cast := (*m.GetCreationSourceKind()).String()
        err := writer.WriteStringValue("creationSourceKind", &cast)
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
        err := writer.WriteObjectValue("teamsPublicationInfo", m.GetTeamsPublicationInfo())
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
func (m *PlannerTaskCreation) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *PlannerTaskCreation) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetCreationSourceKind sets the creationSourceKind property value. Specifies what kind of creation source the task is created with. The possible values are: external, publication and unknownFutureValue.
func (m *PlannerTaskCreation) SetCreationSourceKind(value *PlannerTaskCreation_creationSourceKind)() {
    err := m.GetBackingStore().Set("creationSourceKind", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *PlannerTaskCreation) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetTeamsPublicationInfo sets the teamsPublicationInfo property value. Information about the publication process that created this task. This field is deprecated and clients should move to using the new inheritance model.
func (m *PlannerTaskCreation) SetTeamsPublicationInfo(value PlannerTeamsPublicationInfoable)() {
    err := m.GetBackingStore().Set("teamsPublicationInfo", value)
    if err != nil {
        panic(err)
    }
}
// PlannerTaskCreationable 
type PlannerTaskCreationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetCreationSourceKind()(*PlannerTaskCreation_creationSourceKind)
    GetOdataType()(*string)
    GetTeamsPublicationInfo()(PlannerTeamsPublicationInfoable)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetCreationSourceKind(value *PlannerTaskCreation_creationSourceKind)()
    SetOdataType(value *string)()
    SetTeamsPublicationInfo(value PlannerTeamsPublicationInfoable)()
}

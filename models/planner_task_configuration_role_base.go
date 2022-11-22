package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PlannerTaskConfigurationRoleBase 
type PlannerTaskConfigurationRoleBase struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The OdataType property
    odataType *string
    // The roleKind property
    roleKind *PlannerUserRoleKind
}
// NewPlannerTaskConfigurationRoleBase instantiates a new plannerTaskConfigurationRoleBase and sets the default values.
func NewPlannerTaskConfigurationRoleBase()(*PlannerTaskConfigurationRoleBase) {
    m := &PlannerTaskConfigurationRoleBase{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreatePlannerTaskConfigurationRoleBaseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePlannerTaskConfigurationRoleBaseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.plannerRelationshipBasedUserType":
                        return NewPlannerRelationshipBasedUserType(), nil
                }
            }
        }
    }
    return NewPlannerTaskConfigurationRoleBase(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PlannerTaskConfigurationRoleBase) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PlannerTaskConfigurationRoleBase) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["roleKind"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParsePlannerUserRoleKind , m.SetRoleKind)
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *PlannerTaskConfigurationRoleBase) GetOdataType()(*string) {
    return m.odataType
}
// GetRoleKind gets the roleKind property value. The roleKind property
func (m *PlannerTaskConfigurationRoleBase) GetRoleKind()(*PlannerUserRoleKind) {
    return m.roleKind
}
// Serialize serializes information the current object
func (m *PlannerTaskConfigurationRoleBase) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    if m.GetRoleKind() != nil {
        cast := (*m.GetRoleKind()).String()
        err := writer.WriteStringValue("roleKind", &cast)
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PlannerTaskConfigurationRoleBase) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *PlannerTaskConfigurationRoleBase) SetOdataType(value *string)() {
    m.odataType = value
}
// SetRoleKind sets the roleKind property value. The roleKind property
func (m *PlannerTaskConfigurationRoleBase) SetRoleKind(value *PlannerUserRoleKind)() {
    m.roleKind = value
}

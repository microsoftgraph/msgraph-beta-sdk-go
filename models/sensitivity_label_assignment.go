package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SensitivityLabelAssignment 
type SensitivityLabelAssignment struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The assignmentMethod property
    assignmentMethod *SensitivityLabelAssignmentMethod
    // The OdataType property
    odataType *string
    // The unique identifier for the sensitivity label assigned to the file.
    sensitivityLabelId *string
    // The unique identifier for the tenant that hosts the file when this label is applied.
    tenantId *string
}
// NewSensitivityLabelAssignment instantiates a new sensitivityLabelAssignment and sets the default values.
func NewSensitivityLabelAssignment()(*SensitivityLabelAssignment) {
    m := &SensitivityLabelAssignment{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.sensitivityLabelAssignment";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateSensitivityLabelAssignmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSensitivityLabelAssignmentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSensitivityLabelAssignment(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SensitivityLabelAssignment) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAssignmentMethod gets the assignmentMethod property value. The assignmentMethod property
func (m *SensitivityLabelAssignment) GetAssignmentMethod()(*SensitivityLabelAssignmentMethod) {
    return m.assignmentMethod
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SensitivityLabelAssignment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["assignmentMethod"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseSensitivityLabelAssignmentMethod , m.SetAssignmentMethod)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["sensitivityLabelId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetSensitivityLabelId)
    res["tenantId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetTenantId)
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *SensitivityLabelAssignment) GetOdataType()(*string) {
    return m.odataType
}
// GetSensitivityLabelId gets the sensitivityLabelId property value. The unique identifier for the sensitivity label assigned to the file.
func (m *SensitivityLabelAssignment) GetSensitivityLabelId()(*string) {
    return m.sensitivityLabelId
}
// GetTenantId gets the tenantId property value. The unique identifier for the tenant that hosts the file when this label is applied.
func (m *SensitivityLabelAssignment) GetTenantId()(*string) {
    return m.tenantId
}
// Serialize serializes information the current object
func (m *SensitivityLabelAssignment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAssignmentMethod() != nil {
        cast := (*m.GetAssignmentMethod()).String()
        err := writer.WriteStringValue("assignmentMethod", &cast)
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
        err := writer.WriteStringValue("sensitivityLabelId", m.GetSensitivityLabelId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("tenantId", m.GetTenantId())
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
func (m *SensitivityLabelAssignment) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAssignmentMethod sets the assignmentMethod property value. The assignmentMethod property
func (m *SensitivityLabelAssignment) SetAssignmentMethod(value *SensitivityLabelAssignmentMethod)() {
    m.assignmentMethod = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *SensitivityLabelAssignment) SetOdataType(value *string)() {
    m.odataType = value
}
// SetSensitivityLabelId sets the sensitivityLabelId property value. The unique identifier for the sensitivity label assigned to the file.
func (m *SensitivityLabelAssignment) SetSensitivityLabelId(value *string)() {
    m.sensitivityLabelId = value
}
// SetTenantId sets the tenantId property value. The unique identifier for the tenant that hosts the file when this label is applied.
func (m *SensitivityLabelAssignment) SetTenantId(value *string)() {
    m.tenantId = value
}

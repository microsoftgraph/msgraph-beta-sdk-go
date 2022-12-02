package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DriveRootAssignSensitivityLabelPostRequestBody provides operations to call the assignSensitivityLabel method.
type DriveRootAssignSensitivityLabelPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The assignmentMethod property
    assignmentMethod *SensitivityLabelAssignmentMethod
    // The justificationText property
    justificationText *string
    // The sensitivityLabelId property
    sensitivityLabelId *string
}
// NewDriveRootAssignSensitivityLabelPostRequestBody instantiates a new DriveRootAssignSensitivityLabelPostRequestBody and sets the default values.
func NewDriveRootAssignSensitivityLabelPostRequestBody()(*DriveRootAssignSensitivityLabelPostRequestBody) {
    m := &DriveRootAssignSensitivityLabelPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDriveRootAssignSensitivityLabelPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDriveRootAssignSensitivityLabelPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDriveRootAssignSensitivityLabelPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DriveRootAssignSensitivityLabelPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAssignmentMethod gets the assignmentMethod property value. The assignmentMethod property
func (m *DriveRootAssignSensitivityLabelPostRequestBody) GetAssignmentMethod()(*SensitivityLabelAssignmentMethod) {
    return m.assignmentMethod
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DriveRootAssignSensitivityLabelPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["assignmentMethod"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSensitivityLabelAssignmentMethod)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignmentMethod(val.(*SensitivityLabelAssignmentMethod))
        }
        return nil
    }
    res["justificationText"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJustificationText(val)
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
// GetJustificationText gets the justificationText property value. The justificationText property
func (m *DriveRootAssignSensitivityLabelPostRequestBody) GetJustificationText()(*string) {
    return m.justificationText
}
// GetSensitivityLabelId gets the sensitivityLabelId property value. The sensitivityLabelId property
func (m *DriveRootAssignSensitivityLabelPostRequestBody) GetSensitivityLabelId()(*string) {
    return m.sensitivityLabelId
}
// Serialize serializes information the current object
func (m *DriveRootAssignSensitivityLabelPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAssignmentMethod() != nil {
        cast := (*m.GetAssignmentMethod()).String()
        err := writer.WriteStringValue("assignmentMethod", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("justificationText", m.GetJustificationText())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DriveRootAssignSensitivityLabelPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAssignmentMethod sets the assignmentMethod property value. The assignmentMethod property
func (m *DriveRootAssignSensitivityLabelPostRequestBody) SetAssignmentMethod(value *SensitivityLabelAssignmentMethod)() {
    m.assignmentMethod = value
}
// SetJustificationText sets the justificationText property value. The justificationText property
func (m *DriveRootAssignSensitivityLabelPostRequestBody) SetJustificationText(value *string)() {
    m.justificationText = value
}
// SetSensitivityLabelId sets the sensitivityLabelId property value. The sensitivityLabelId property
func (m *DriveRootAssignSensitivityLabelPostRequestBody) SetSensitivityLabelId(value *string)() {
    m.sensitivityLabelId = value
}

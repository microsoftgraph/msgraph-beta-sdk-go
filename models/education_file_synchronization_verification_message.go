package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EducationFileSynchronizationVerificationMessage 
type EducationFileSynchronizationVerificationMessage struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Detailed information about the message type.
    description *string
    // The fileName property
    fileName *string
    // The OdataType property
    odataType *string
    // Type of the message. Possible values are: error, warning, information.
    type_escaped *string
}
// NewEducationFileSynchronizationVerificationMessage instantiates a new educationFileSynchronizationVerificationMessage and sets the default values.
func NewEducationFileSynchronizationVerificationMessage()(*EducationFileSynchronizationVerificationMessage) {
    m := &EducationFileSynchronizationVerificationMessage{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateEducationFileSynchronizationVerificationMessageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEducationFileSynchronizationVerificationMessageFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEducationFileSynchronizationVerificationMessage(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EducationFileSynchronizationVerificationMessage) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDescription gets the description property value. Detailed information about the message type.
func (m *EducationFileSynchronizationVerificationMessage) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EducationFileSynchronizationVerificationMessage) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["fileName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFileName(val)
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
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val)
        }
        return nil
    }
    return res
}
// GetFileName gets the fileName property value. The fileName property
func (m *EducationFileSynchronizationVerificationMessage) GetFileName()(*string) {
    return m.fileName
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *EducationFileSynchronizationVerificationMessage) GetOdataType()(*string) {
    return m.odataType
}
// GetType gets the type property value. Type of the message. Possible values are: error, warning, information.
func (m *EducationFileSynchronizationVerificationMessage) GetType()(*string) {
    return m.type_escaped
}
// Serialize serializes information the current object
func (m *EducationFileSynchronizationVerificationMessage) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("fileName", m.GetFileName())
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
        err := writer.WriteStringValue("type", m.GetType())
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
func (m *EducationFileSynchronizationVerificationMessage) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDescription sets the description property value. Detailed information about the message type.
func (m *EducationFileSynchronizationVerificationMessage) SetDescription(value *string)() {
    m.description = value
}
// SetFileName sets the fileName property value. The fileName property
func (m *EducationFileSynchronizationVerificationMessage) SetFileName(value *string)() {
    m.fileName = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *EducationFileSynchronizationVerificationMessage) SetOdataType(value *string)() {
    m.odataType = value
}
// SetType sets the type property value. Type of the message. Possible values are: error, warning, information.
func (m *EducationFileSynchronizationVerificationMessage) SetType(value *string)() {
    m.type_escaped = value
}

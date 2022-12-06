package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CloudPcStatusDetails 
type CloudPcStatusDetails struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Any additional information about the Cloud PC status.
    additionalInformation []KeyValuePairable
    // The code associated with the Cloud PC status.
    code *string
    // The status message.
    message *string
    // The OdataType property
    odataType *string
}
// NewCloudPcStatusDetails instantiates a new cloudPcStatusDetails and sets the default values.
func NewCloudPcStatusDetails()(*CloudPcStatusDetails) {
    m := &CloudPcStatusDetails{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateCloudPcStatusDetailsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCloudPcStatusDetailsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudPcStatusDetails(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CloudPcStatusDetails) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAdditionalInformation gets the additionalInformation property value. Any additional information about the Cloud PC status.
func (m *CloudPcStatusDetails) GetAdditionalInformation()([]KeyValuePairable) {
    return m.additionalInformation
}
// GetCode gets the code property value. The code associated with the Cloud PC status.
func (m *CloudPcStatusDetails) GetCode()(*string) {
    return m.code
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CloudPcStatusDetails) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["additionalInformation"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateKeyValuePairFromDiscriminatorValue , m.SetAdditionalInformation)
    res["code"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetCode)
    res["message"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetMessage)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    return res
}
// GetMessage gets the message property value. The status message.
func (m *CloudPcStatusDetails) GetMessage()(*string) {
    return m.message
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *CloudPcStatusDetails) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *CloudPcStatusDetails) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAdditionalInformation() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetAdditionalInformation())
        err := writer.WriteCollectionOfObjectValues("additionalInformation", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("code", m.GetCode())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("message", m.GetMessage())
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CloudPcStatusDetails) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAdditionalInformation sets the additionalInformation property value. Any additional information about the Cloud PC status.
func (m *CloudPcStatusDetails) SetAdditionalInformation(value []KeyValuePairable)() {
    m.additionalInformation = value
}
// SetCode sets the code property value. The code associated with the Cloud PC status.
func (m *CloudPcStatusDetails) SetCode(value *string)() {
    m.code = value
}
// SetMessage sets the message property value. The status message.
func (m *CloudPcStatusDetails) SetMessage(value *string)() {
    m.message = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *CloudPcStatusDetails) SetOdataType(value *string)() {
    m.odataType = value
}

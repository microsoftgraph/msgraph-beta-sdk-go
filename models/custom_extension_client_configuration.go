package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CustomExtensionClientConfiguration 
type CustomExtensionClientConfiguration struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The OdataType property
    odataType *string
    // The max duration in milliseconds that Azure AD will wait for a response from the logic app before it shuts down the connection. The valid range is between 200 and 2000 milliseconds. Default duration is 1000.
    timeoutInMilliseconds *int32
}
// NewCustomExtensionClientConfiguration instantiates a new customExtensionClientConfiguration and sets the default values.
func NewCustomExtensionClientConfiguration()(*CustomExtensionClientConfiguration) {
    m := &CustomExtensionClientConfiguration{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateCustomExtensionClientConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCustomExtensionClientConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCustomExtensionClientConfiguration(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CustomExtensionClientConfiguration) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CustomExtensionClientConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["timeoutInMilliseconds"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetTimeoutInMilliseconds)
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *CustomExtensionClientConfiguration) GetOdataType()(*string) {
    return m.odataType
}
// GetTimeoutInMilliseconds gets the timeoutInMilliseconds property value. The max duration in milliseconds that Azure AD will wait for a response from the logic app before it shuts down the connection. The valid range is between 200 and 2000 milliseconds. Default duration is 1000.
func (m *CustomExtensionClientConfiguration) GetTimeoutInMilliseconds()(*int32) {
    return m.timeoutInMilliseconds
}
// Serialize serializes information the current object
func (m *CustomExtensionClientConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("timeoutInMilliseconds", m.GetTimeoutInMilliseconds())
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
func (m *CustomExtensionClientConfiguration) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *CustomExtensionClientConfiguration) SetOdataType(value *string)() {
    m.odataType = value
}
// SetTimeoutInMilliseconds sets the timeoutInMilliseconds property value. The max duration in milliseconds that Azure AD will wait for a response from the logic app before it shuts down the connection. The valid range is between 200 and 2000 milliseconds. Default duration is 1000.
func (m *CustomExtensionClientConfiguration) SetTimeoutInMilliseconds(value *int32)() {
    m.timeoutInMilliseconds = value
}

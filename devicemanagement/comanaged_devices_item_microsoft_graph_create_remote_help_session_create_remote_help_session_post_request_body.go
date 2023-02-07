package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ComanagedDevicesItemMicrosoftGraphCreateRemoteHelpSessionCreateRemoteHelpSessionPostRequestBody 
type ComanagedDevicesItemMicrosoftGraphCreateRemoteHelpSessionCreateRemoteHelpSessionPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The sessionType property
    sessionType *string
}
// NewComanagedDevicesItemMicrosoftGraphCreateRemoteHelpSessionCreateRemoteHelpSessionPostRequestBody instantiates a new ComanagedDevicesItemMicrosoftGraphCreateRemoteHelpSessionCreateRemoteHelpSessionPostRequestBody and sets the default values.
func NewComanagedDevicesItemMicrosoftGraphCreateRemoteHelpSessionCreateRemoteHelpSessionPostRequestBody()(*ComanagedDevicesItemMicrosoftGraphCreateRemoteHelpSessionCreateRemoteHelpSessionPostRequestBody) {
    m := &ComanagedDevicesItemMicrosoftGraphCreateRemoteHelpSessionCreateRemoteHelpSessionPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateComanagedDevicesItemMicrosoftGraphCreateRemoteHelpSessionCreateRemoteHelpSessionPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateComanagedDevicesItemMicrosoftGraphCreateRemoteHelpSessionCreateRemoteHelpSessionPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewComanagedDevicesItemMicrosoftGraphCreateRemoteHelpSessionCreateRemoteHelpSessionPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ComanagedDevicesItemMicrosoftGraphCreateRemoteHelpSessionCreateRemoteHelpSessionPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ComanagedDevicesItemMicrosoftGraphCreateRemoteHelpSessionCreateRemoteHelpSessionPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["sessionType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSessionType(val)
        }
        return nil
    }
    return res
}
// GetSessionType gets the sessionType property value. The sessionType property
func (m *ComanagedDevicesItemMicrosoftGraphCreateRemoteHelpSessionCreateRemoteHelpSessionPostRequestBody) GetSessionType()(*string) {
    return m.sessionType
}
// Serialize serializes information the current object
func (m *ComanagedDevicesItemMicrosoftGraphCreateRemoteHelpSessionCreateRemoteHelpSessionPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("sessionType", m.GetSessionType())
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
func (m *ComanagedDevicesItemMicrosoftGraphCreateRemoteHelpSessionCreateRemoteHelpSessionPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetSessionType sets the sessionType property value. The sessionType property
func (m *ComanagedDevicesItemMicrosoftGraphCreateRemoteHelpSessionCreateRemoteHelpSessionPostRequestBody) SetSessionType(value *string)() {
    m.sessionType = value
}

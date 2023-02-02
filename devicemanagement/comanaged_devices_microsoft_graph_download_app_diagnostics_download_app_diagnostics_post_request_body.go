package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// ComanagedDevicesMicrosoftGraphDownloadAppDiagnosticsDownloadAppDiagnosticsPostRequestBody 
type ComanagedDevicesMicrosoftGraphDownloadAppDiagnosticsDownloadAppDiagnosticsPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The request property
    request ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PowerliftDownloadRequestable
}
// NewComanagedDevicesMicrosoftGraphDownloadAppDiagnosticsDownloadAppDiagnosticsPostRequestBody instantiates a new ComanagedDevicesMicrosoftGraphDownloadAppDiagnosticsDownloadAppDiagnosticsPostRequestBody and sets the default values.
func NewComanagedDevicesMicrosoftGraphDownloadAppDiagnosticsDownloadAppDiagnosticsPostRequestBody()(*ComanagedDevicesMicrosoftGraphDownloadAppDiagnosticsDownloadAppDiagnosticsPostRequestBody) {
    m := &ComanagedDevicesMicrosoftGraphDownloadAppDiagnosticsDownloadAppDiagnosticsPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateComanagedDevicesMicrosoftGraphDownloadAppDiagnosticsDownloadAppDiagnosticsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateComanagedDevicesMicrosoftGraphDownloadAppDiagnosticsDownloadAppDiagnosticsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewComanagedDevicesMicrosoftGraphDownloadAppDiagnosticsDownloadAppDiagnosticsPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ComanagedDevicesMicrosoftGraphDownloadAppDiagnosticsDownloadAppDiagnosticsPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ComanagedDevicesMicrosoftGraphDownloadAppDiagnosticsDownloadAppDiagnosticsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["request"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePowerliftDownloadRequestFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequest(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PowerliftDownloadRequestable))
        }
        return nil
    }
    return res
}
// GetRequest gets the request property value. The request property
func (m *ComanagedDevicesMicrosoftGraphDownloadAppDiagnosticsDownloadAppDiagnosticsPostRequestBody) GetRequest()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PowerliftDownloadRequestable) {
    return m.request
}
// Serialize serializes information the current object
func (m *ComanagedDevicesMicrosoftGraphDownloadAppDiagnosticsDownloadAppDiagnosticsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("request", m.GetRequest())
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
func (m *ComanagedDevicesMicrosoftGraphDownloadAppDiagnosticsDownloadAppDiagnosticsPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetRequest sets the request property value. The request property
func (m *ComanagedDevicesMicrosoftGraphDownloadAppDiagnosticsDownloadAppDiagnosticsPostRequestBody) SetRequest(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PowerliftDownloadRequestable)() {
    m.request = value
}

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MeManagedDevicesDownloadAppDiagnosticsPostRequestBody provides operations to call the downloadAppDiagnostics method.
type MeManagedDevicesDownloadAppDiagnosticsPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The request property
    request PowerliftDownloadRequestable
}
// NewMeManagedDevicesDownloadAppDiagnosticsPostRequestBody instantiates a new MeManagedDevicesDownloadAppDiagnosticsPostRequestBody and sets the default values.
func NewMeManagedDevicesDownloadAppDiagnosticsPostRequestBody()(*MeManagedDevicesDownloadAppDiagnosticsPostRequestBody) {
    m := &MeManagedDevicesDownloadAppDiagnosticsPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateMeManagedDevicesDownloadAppDiagnosticsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMeManagedDevicesDownloadAppDiagnosticsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMeManagedDevicesDownloadAppDiagnosticsPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MeManagedDevicesDownloadAppDiagnosticsPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MeManagedDevicesDownloadAppDiagnosticsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["request"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePowerliftDownloadRequestFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequest(val.(PowerliftDownloadRequestable))
        }
        return nil
    }
    return res
}
// GetRequest gets the request property value. The request property
func (m *MeManagedDevicesDownloadAppDiagnosticsPostRequestBody) GetRequest()(PowerliftDownloadRequestable) {
    return m.request
}
// Serialize serializes information the current object
func (m *MeManagedDevicesDownloadAppDiagnosticsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *MeManagedDevicesDownloadAppDiagnosticsPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetRequest sets the request property value. The request property
func (m *MeManagedDevicesDownloadAppDiagnosticsPostRequestBody) SetRequest(value PowerliftDownloadRequestable)() {
    m.request = value
}

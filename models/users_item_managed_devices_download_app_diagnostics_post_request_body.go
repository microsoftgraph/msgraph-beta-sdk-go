package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UsersItemManagedDevicesDownloadAppDiagnosticsPostRequestBody provides operations to call the downloadAppDiagnostics method.
type UsersItemManagedDevicesDownloadAppDiagnosticsPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The request property
    request PowerliftDownloadRequestable
}
// NewUsersItemManagedDevicesDownloadAppDiagnosticsPostRequestBody instantiates a new UsersItemManagedDevicesDownloadAppDiagnosticsPostRequestBody and sets the default values.
func NewUsersItemManagedDevicesDownloadAppDiagnosticsPostRequestBody()(*UsersItemManagedDevicesDownloadAppDiagnosticsPostRequestBody) {
    m := &UsersItemManagedDevicesDownloadAppDiagnosticsPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateUsersItemManagedDevicesDownloadAppDiagnosticsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUsersItemManagedDevicesDownloadAppDiagnosticsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUsersItemManagedDevicesDownloadAppDiagnosticsPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UsersItemManagedDevicesDownloadAppDiagnosticsPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UsersItemManagedDevicesDownloadAppDiagnosticsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *UsersItemManagedDevicesDownloadAppDiagnosticsPostRequestBody) GetRequest()(PowerliftDownloadRequestable) {
    return m.request
}
// Serialize serializes information the current object
func (m *UsersItemManagedDevicesDownloadAppDiagnosticsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *UsersItemManagedDevicesDownloadAppDiagnosticsPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetRequest sets the request property value. The request property
func (m *UsersItemManagedDevicesDownloadAppDiagnosticsPostRequestBody) SetRequest(value PowerliftDownloadRequestable)() {
    m.request = value
}

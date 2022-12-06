package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PreAuthorizedApplication 
type PreAuthorizedApplication struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The unique identifier for the client application.
    appId *string
    // The OdataType property
    odataType *string
    // The unique identifier for the scopes the client application is granted.
    permissionIds []string
}
// NewPreAuthorizedApplication instantiates a new preAuthorizedApplication and sets the default values.
func NewPreAuthorizedApplication()(*PreAuthorizedApplication) {
    m := &PreAuthorizedApplication{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreatePreAuthorizedApplicationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePreAuthorizedApplicationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPreAuthorizedApplication(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PreAuthorizedApplication) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAppId gets the appId property value. The unique identifier for the client application.
func (m *PreAuthorizedApplication) GetAppId()(*string) {
    return m.appId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PreAuthorizedApplication) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["appId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetAppId)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["permissionIds"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetPermissionIds)
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *PreAuthorizedApplication) GetOdataType()(*string) {
    return m.odataType
}
// GetPermissionIds gets the permissionIds property value. The unique identifier for the scopes the client application is granted.
func (m *PreAuthorizedApplication) GetPermissionIds()([]string) {
    return m.permissionIds
}
// Serialize serializes information the current object
func (m *PreAuthorizedApplication) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("appId", m.GetAppId())
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
    if m.GetPermissionIds() != nil {
        err := writer.WriteCollectionOfStringValues("permissionIds", m.GetPermissionIds())
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
func (m *PreAuthorizedApplication) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAppId sets the appId property value. The unique identifier for the client application.
func (m *PreAuthorizedApplication) SetAppId(value *string)() {
    m.appId = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *PreAuthorizedApplication) SetOdataType(value *string)() {
    m.odataType = value
}
// SetPermissionIds sets the permissionIds property value. The unique identifier for the scopes the client application is granted.
func (m *PreAuthorizedApplication) SetPermissionIds(value []string)() {
    m.permissionIds = value
}

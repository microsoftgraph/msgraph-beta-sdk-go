package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MeManagedDevicesMoveDevicesToOUPostRequestBody provides operations to call the moveDevicesToOU method.
type MeManagedDevicesMoveDevicesToOUPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The deviceIds property
    deviceIds []string
    // The organizationalUnitPath property
    organizationalUnitPath *string
}
// NewMeManagedDevicesMoveDevicesToOUPostRequestBody instantiates a new MeManagedDevicesMoveDevicesToOUPostRequestBody and sets the default values.
func NewMeManagedDevicesMoveDevicesToOUPostRequestBody()(*MeManagedDevicesMoveDevicesToOUPostRequestBody) {
    m := &MeManagedDevicesMoveDevicesToOUPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateMeManagedDevicesMoveDevicesToOUPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMeManagedDevicesMoveDevicesToOUPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMeManagedDevicesMoveDevicesToOUPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MeManagedDevicesMoveDevicesToOUPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDeviceIds gets the deviceIds property value. The deviceIds property
func (m *MeManagedDevicesMoveDevicesToOUPostRequestBody) GetDeviceIds()([]string) {
    return m.deviceIds
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MeManagedDevicesMoveDevicesToOUPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["deviceIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetDeviceIds(res)
        }
        return nil
    }
    res["organizationalUnitPath"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganizationalUnitPath(val)
        }
        return nil
    }
    return res
}
// GetOrganizationalUnitPath gets the organizationalUnitPath property value. The organizationalUnitPath property
func (m *MeManagedDevicesMoveDevicesToOUPostRequestBody) GetOrganizationalUnitPath()(*string) {
    return m.organizationalUnitPath
}
// Serialize serializes information the current object
func (m *MeManagedDevicesMoveDevicesToOUPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetDeviceIds() != nil {
        err := writer.WriteCollectionOfStringValues("deviceIds", m.GetDeviceIds())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("organizationalUnitPath", m.GetOrganizationalUnitPath())
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
func (m *MeManagedDevicesMoveDevicesToOUPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDeviceIds sets the deviceIds property value. The deviceIds property
func (m *MeManagedDevicesMoveDevicesToOUPostRequestBody) SetDeviceIds(value []string)() {
    m.deviceIds = value
}
// SetOrganizationalUnitPath sets the organizationalUnitPath property value. The organizationalUnitPath property
func (m *MeManagedDevicesMoveDevicesToOUPostRequestBody) SetOrganizationalUnitPath(value *string)() {
    m.organizationalUnitPath = value
}

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceAppManagementMdmWindowsInformationProtectionPoliciesHasPayloadLinksPostRequestBody provides operations to call the hasPayloadLinks method.
type DeviceAppManagementMdmWindowsInformationProtectionPoliciesHasPayloadLinksPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The payloadIds property
    payloadIds []string
}
// NewDeviceAppManagementMdmWindowsInformationProtectionPoliciesHasPayloadLinksPostRequestBody instantiates a new DeviceAppManagementMdmWindowsInformationProtectionPoliciesHasPayloadLinksPostRequestBody and sets the default values.
func NewDeviceAppManagementMdmWindowsInformationProtectionPoliciesHasPayloadLinksPostRequestBody()(*DeviceAppManagementMdmWindowsInformationProtectionPoliciesHasPayloadLinksPostRequestBody) {
    m := &DeviceAppManagementMdmWindowsInformationProtectionPoliciesHasPayloadLinksPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceAppManagementMdmWindowsInformationProtectionPoliciesHasPayloadLinksPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceAppManagementMdmWindowsInformationProtectionPoliciesHasPayloadLinksPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceAppManagementMdmWindowsInformationProtectionPoliciesHasPayloadLinksPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceAppManagementMdmWindowsInformationProtectionPoliciesHasPayloadLinksPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceAppManagementMdmWindowsInformationProtectionPoliciesHasPayloadLinksPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["payloadIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetPayloadIds(res)
        }
        return nil
    }
    return res
}
// GetPayloadIds gets the payloadIds property value. The payloadIds property
func (m *DeviceAppManagementMdmWindowsInformationProtectionPoliciesHasPayloadLinksPostRequestBody) GetPayloadIds()([]string) {
    return m.payloadIds
}
// Serialize serializes information the current object
func (m *DeviceAppManagementMdmWindowsInformationProtectionPoliciesHasPayloadLinksPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetPayloadIds() != nil {
        err := writer.WriteCollectionOfStringValues("payloadIds", m.GetPayloadIds())
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
func (m *DeviceAppManagementMdmWindowsInformationProtectionPoliciesHasPayloadLinksPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetPayloadIds sets the payloadIds property value. The payloadIds property
func (m *DeviceAppManagementMdmWindowsInformationProtectionPoliciesHasPayloadLinksPostRequestBody) SetPayloadIds(value []string)() {
    m.payloadIds = value
}

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CloudPcBulkRemoteActionResult 
type CloudPcBulkRemoteActionResult struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // A list of all the Intune managed device IDs that completed the bulk action with a failure.
    failedDeviceIds []string
    // A list of all the Intune managed device IDs that were not found when the bulk action was attempted.
    notFoundDeviceIds []string
    // A list of all the Intune managed device IDs that were identified as unsupported for the bulk action.
    notSupportedDeviceIds []string
    // The OdataType property
    odataType *string
    // A list of all the Intune managed device IDs that completed the bulk action successfully.
    successfulDeviceIds []string
}
// NewCloudPcBulkRemoteActionResult instantiates a new cloudPcBulkRemoteActionResult and sets the default values.
func NewCloudPcBulkRemoteActionResult()(*CloudPcBulkRemoteActionResult) {
    m := &CloudPcBulkRemoteActionResult{
    }
    m.SetAdditionalData(make(map[string]any));
    return m
}
// CreateCloudPcBulkRemoteActionResultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCloudPcBulkRemoteActionResultFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudPcBulkRemoteActionResult(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CloudPcBulkRemoteActionResult) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFailedDeviceIds gets the failedDeviceIds property value. A list of all the Intune managed device IDs that completed the bulk action with a failure.
func (m *CloudPcBulkRemoteActionResult) GetFailedDeviceIds()([]string) {
    return m.failedDeviceIds
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CloudPcBulkRemoteActionResult) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["failedDeviceIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetFailedDeviceIds(res)
        }
        return nil
    }
    res["notFoundDeviceIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetNotFoundDeviceIds(res)
        }
        return nil
    }
    res["notSupportedDeviceIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetNotSupportedDeviceIds(res)
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
    res["successfulDeviceIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetSuccessfulDeviceIds(res)
        }
        return nil
    }
    return res
}
// GetNotFoundDeviceIds gets the notFoundDeviceIds property value. A list of all the Intune managed device IDs that were not found when the bulk action was attempted.
func (m *CloudPcBulkRemoteActionResult) GetNotFoundDeviceIds()([]string) {
    return m.notFoundDeviceIds
}
// GetNotSupportedDeviceIds gets the notSupportedDeviceIds property value. A list of all the Intune managed device IDs that were identified as unsupported for the bulk action.
func (m *CloudPcBulkRemoteActionResult) GetNotSupportedDeviceIds()([]string) {
    return m.notSupportedDeviceIds
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *CloudPcBulkRemoteActionResult) GetOdataType()(*string) {
    return m.odataType
}
// GetSuccessfulDeviceIds gets the successfulDeviceIds property value. A list of all the Intune managed device IDs that completed the bulk action successfully.
func (m *CloudPcBulkRemoteActionResult) GetSuccessfulDeviceIds()([]string) {
    return m.successfulDeviceIds
}
// Serialize serializes information the current object
func (m *CloudPcBulkRemoteActionResult) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetFailedDeviceIds() != nil {
        err := writer.WriteCollectionOfStringValues("failedDeviceIds", m.GetFailedDeviceIds())
        if err != nil {
            return err
        }
    }
    if m.GetNotFoundDeviceIds() != nil {
        err := writer.WriteCollectionOfStringValues("notFoundDeviceIds", m.GetNotFoundDeviceIds())
        if err != nil {
            return err
        }
    }
    if m.GetNotSupportedDeviceIds() != nil {
        err := writer.WriteCollectionOfStringValues("notSupportedDeviceIds", m.GetNotSupportedDeviceIds())
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
    if m.GetSuccessfulDeviceIds() != nil {
        err := writer.WriteCollectionOfStringValues("successfulDeviceIds", m.GetSuccessfulDeviceIds())
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
func (m *CloudPcBulkRemoteActionResult) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetFailedDeviceIds sets the failedDeviceIds property value. A list of all the Intune managed device IDs that completed the bulk action with a failure.
func (m *CloudPcBulkRemoteActionResult) SetFailedDeviceIds(value []string)() {
    m.failedDeviceIds = value
}
// SetNotFoundDeviceIds sets the notFoundDeviceIds property value. A list of all the Intune managed device IDs that were not found when the bulk action was attempted.
func (m *CloudPcBulkRemoteActionResult) SetNotFoundDeviceIds(value []string)() {
    m.notFoundDeviceIds = value
}
// SetNotSupportedDeviceIds sets the notSupportedDeviceIds property value. A list of all the Intune managed device IDs that were identified as unsupported for the bulk action.
func (m *CloudPcBulkRemoteActionResult) SetNotSupportedDeviceIds(value []string)() {
    m.notSupportedDeviceIds = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *CloudPcBulkRemoteActionResult) SetOdataType(value *string)() {
    m.odataType = value
}
// SetSuccessfulDeviceIds sets the successfulDeviceIds property value. A list of all the Intune managed device IDs that completed the bulk action successfully.
func (m *CloudPcBulkRemoteActionResult) SetSuccessfulDeviceIds(value []string)() {
    m.successfulDeviceIds = value
}

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BulkManagedDeviceActionResult 
type BulkManagedDeviceActionResult struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Failed devices
    failedDeviceIds []string
    // Not found devices
    notFoundDeviceIds []string
    // Not supported devices
    notSupportedDeviceIds []string
    // The OdataType property
    odataType *string
    // Successful devices
    successfulDeviceIds []string
}
// NewBulkManagedDeviceActionResult instantiates a new bulkManagedDeviceActionResult and sets the default values.
func NewBulkManagedDeviceActionResult()(*BulkManagedDeviceActionResult) {
    m := &BulkManagedDeviceActionResult{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.bulkManagedDeviceActionResult";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateBulkManagedDeviceActionResultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateBulkManagedDeviceActionResultFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBulkManagedDeviceActionResult(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *BulkManagedDeviceActionResult) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFailedDeviceIds gets the failedDeviceIds property value. Failed devices
func (m *BulkManagedDeviceActionResult) GetFailedDeviceIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.failedDeviceIds
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *BulkManagedDeviceActionResult) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
// GetNotFoundDeviceIds gets the notFoundDeviceIds property value. Not found devices
func (m *BulkManagedDeviceActionResult) GetNotFoundDeviceIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.notFoundDeviceIds
    }
}
// GetNotSupportedDeviceIds gets the notSupportedDeviceIds property value. Not supported devices
func (m *BulkManagedDeviceActionResult) GetNotSupportedDeviceIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.notSupportedDeviceIds
    }
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *BulkManagedDeviceActionResult) GetOdataType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.odataType
    }
}
// GetSuccessfulDeviceIds gets the successfulDeviceIds property value. Successful devices
func (m *BulkManagedDeviceActionResult) GetSuccessfulDeviceIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.successfulDeviceIds
    }
}
// Serialize serializes information the current object
func (m *BulkManagedDeviceActionResult) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *BulkManagedDeviceActionResult) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetFailedDeviceIds sets the failedDeviceIds property value. Failed devices
func (m *BulkManagedDeviceActionResult) SetFailedDeviceIds(value []string)() {
    if m != nil {
        m.failedDeviceIds = value
    }
}
// SetNotFoundDeviceIds sets the notFoundDeviceIds property value. Not found devices
func (m *BulkManagedDeviceActionResult) SetNotFoundDeviceIds(value []string)() {
    if m != nil {
        m.notFoundDeviceIds = value
    }
}
// SetNotSupportedDeviceIds sets the notSupportedDeviceIds property value. Not supported devices
func (m *BulkManagedDeviceActionResult) SetNotSupportedDeviceIds(value []string)() {
    if m != nil {
        m.notSupportedDeviceIds = value
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *BulkManagedDeviceActionResult) SetOdataType(value *string)() {
    if m != nil {
        m.odataType = value
    }
}
// SetSuccessfulDeviceIds sets the successfulDeviceIds property value. Successful devices
func (m *BulkManagedDeviceActionResult) SetSuccessfulDeviceIds(value []string)() {
    if m != nil {
        m.successfulDeviceIds = value
    }
}

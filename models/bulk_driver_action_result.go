package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BulkDriverActionResult a complex type to represent the result of bulk driver action.
type BulkDriverActionResult struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // List of driver Ids where the action is failed.
    failedDriverIds []string
    // List of driver Ids that are not found.
    notFoundDriverIds []string
    // The OdataType property
    odataType *string
    // List of driver Ids where the action is successful.
    successfulDriverIds []string
}
// NewBulkDriverActionResult instantiates a new bulkDriverActionResult and sets the default values.
func NewBulkDriverActionResult()(*BulkDriverActionResult) {
    m := &BulkDriverActionResult{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateBulkDriverActionResultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateBulkDriverActionResultFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBulkDriverActionResult(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *BulkDriverActionResult) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFailedDriverIds gets the failedDriverIds property value. List of driver Ids where the action is failed.
func (m *BulkDriverActionResult) GetFailedDriverIds()([]string) {
    return m.failedDriverIds
}
// GetFieldDeserializers the deserialization information for the current model
func (m *BulkDriverActionResult) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["failedDriverIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetFailedDriverIds(res)
        }
        return nil
    }
    res["notFoundDriverIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetNotFoundDriverIds(res)
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
    res["successfulDriverIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetSuccessfulDriverIds(res)
        }
        return nil
    }
    return res
}
// GetNotFoundDriverIds gets the notFoundDriverIds property value. List of driver Ids that are not found.
func (m *BulkDriverActionResult) GetNotFoundDriverIds()([]string) {
    return m.notFoundDriverIds
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *BulkDriverActionResult) GetOdataType()(*string) {
    return m.odataType
}
// GetSuccessfulDriverIds gets the successfulDriverIds property value. List of driver Ids where the action is successful.
func (m *BulkDriverActionResult) GetSuccessfulDriverIds()([]string) {
    return m.successfulDriverIds
}
// Serialize serializes information the current object
func (m *BulkDriverActionResult) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetFailedDriverIds() != nil {
        err := writer.WriteCollectionOfStringValues("failedDriverIds", m.GetFailedDriverIds())
        if err != nil {
            return err
        }
    }
    if m.GetNotFoundDriverIds() != nil {
        err := writer.WriteCollectionOfStringValues("notFoundDriverIds", m.GetNotFoundDriverIds())
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
    if m.GetSuccessfulDriverIds() != nil {
        err := writer.WriteCollectionOfStringValues("successfulDriverIds", m.GetSuccessfulDriverIds())
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
func (m *BulkDriverActionResult) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetFailedDriverIds sets the failedDriverIds property value. List of driver Ids where the action is failed.
func (m *BulkDriverActionResult) SetFailedDriverIds(value []string)() {
    m.failedDriverIds = value
}
// SetNotFoundDriverIds sets the notFoundDriverIds property value. List of driver Ids that are not found.
func (m *BulkDriverActionResult) SetNotFoundDriverIds(value []string)() {
    m.notFoundDriverIds = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *BulkDriverActionResult) SetOdataType(value *string)() {
    m.odataType = value
}
// SetSuccessfulDriverIds sets the successfulDriverIds property value. List of driver Ids where the action is successful.
func (m *BulkDriverActionResult) SetSuccessfulDriverIds(value []string)() {
    m.successfulDriverIds = value
}

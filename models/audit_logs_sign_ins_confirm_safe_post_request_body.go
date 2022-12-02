package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AuditLogsSignInsConfirmSafePostRequestBody provides operations to call the confirmSafe method.
type AuditLogsSignInsConfirmSafePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The requestIds property
    requestIds []string
}
// NewAuditLogsSignInsConfirmSafePostRequestBody instantiates a new AuditLogsSignInsConfirmSafePostRequestBody and sets the default values.
func NewAuditLogsSignInsConfirmSafePostRequestBody()(*AuditLogsSignInsConfirmSafePostRequestBody) {
    m := &AuditLogsSignInsConfirmSafePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAuditLogsSignInsConfirmSafePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAuditLogsSignInsConfirmSafePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuditLogsSignInsConfirmSafePostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AuditLogsSignInsConfirmSafePostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AuditLogsSignInsConfirmSafePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["requestIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetRequestIds(res)
        }
        return nil
    }
    return res
}
// GetRequestIds gets the requestIds property value. The requestIds property
func (m *AuditLogsSignInsConfirmSafePostRequestBody) GetRequestIds()([]string) {
    return m.requestIds
}
// Serialize serializes information the current object
func (m *AuditLogsSignInsConfirmSafePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetRequestIds() != nil {
        err := writer.WriteCollectionOfStringValues("requestIds", m.GetRequestIds())
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
func (m *AuditLogsSignInsConfirmSafePostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetRequestIds sets the requestIds property value. The requestIds property
func (m *AuditLogsSignInsConfirmSafePostRequestBody) SetRequestIds(value []string)() {
    m.requestIds = value
}

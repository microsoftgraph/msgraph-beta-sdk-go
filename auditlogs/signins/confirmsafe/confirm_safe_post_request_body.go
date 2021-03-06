package confirmsafe

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ConfirmSafePostRequestBody provides operations to call the confirmSafe method.
type ConfirmSafePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The requestIds property
    requestIds []string
}
// NewConfirmSafePostRequestBody instantiates a new confirmSafePostRequestBody and sets the default values.
func NewConfirmSafePostRequestBody()(*ConfirmSafePostRequestBody) {
    m := &ConfirmSafePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateConfirmSafePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateConfirmSafePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewConfirmSafePostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ConfirmSafePostRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ConfirmSafePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *ConfirmSafePostRequestBody) GetRequestIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.requestIds
    }
}
// Serialize serializes information the current object
func (m *ConfirmSafePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ConfirmSafePostRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetRequestIds sets the requestIds property value. The requestIds property
func (m *ConfirmSafePostRequestBody) SetRequestIds(value []string)() {
    if m != nil {
        m.requestIds = value
    }
}

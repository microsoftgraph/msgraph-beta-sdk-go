package getmobileappcountwithstatus

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GetMobileAppCountWithStatusResponse provides operations to call the getMobileAppCount method.
type GetMobileAppCountWithStatusResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The value property
    value *int64
}
// NewGetMobileAppCountWithStatusResponse instantiates a new getMobileAppCountWithStatusResponse and sets the default values.
func NewGetMobileAppCountWithStatusResponse()(*GetMobileAppCountWithStatusResponse) {
    m := &GetMobileAppCountWithStatusResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateGetMobileAppCountWithStatusResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGetMobileAppCountWithStatusResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGetMobileAppCountWithStatusResponse(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GetMobileAppCountWithStatusResponse) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GetMobileAppCountWithStatusResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["value"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt64Value(m.SetValue)
    return res
}
// GetValue gets the value property value. The value property
func (m *GetMobileAppCountWithStatusResponse) GetValue()(*int64) {
    return m.value
}
// Serialize serializes information the current object
func (m *GetMobileAppCountWithStatusResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt64Value("value", m.GetValue())
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
func (m *GetMobileAppCountWithStatusResponse) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetValue sets the value property value. The value property
func (m *GetMobileAppCountWithStatusResponse) SetValue(value *int64)() {
    m.value = value
}

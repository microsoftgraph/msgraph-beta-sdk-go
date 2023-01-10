package directory

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RecommendationsItemDismissPostRequestBody 
type RecommendationsItemDismissPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The dismissReason property
    dismissReason *string
}
// NewRecommendationsItemDismissPostRequestBody instantiates a new RecommendationsItemDismissPostRequestBody and sets the default values.
func NewRecommendationsItemDismissPostRequestBody()(*RecommendationsItemDismissPostRequestBody) {
    m := &RecommendationsItemDismissPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateRecommendationsItemDismissPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRecommendationsItemDismissPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRecommendationsItemDismissPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RecommendationsItemDismissPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDismissReason gets the dismissReason property value. The dismissReason property
func (m *RecommendationsItemDismissPostRequestBody) GetDismissReason()(*string) {
    return m.dismissReason
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RecommendationsItemDismissPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["dismissReason"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDismissReason(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *RecommendationsItemDismissPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("dismissReason", m.GetDismissReason())
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
func (m *RecommendationsItemDismissPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDismissReason sets the dismissReason property value. The dismissReason property
func (m *RecommendationsItemDismissPostRequestBody) SetDismissReason(value *string)() {
    m.dismissReason = value
}

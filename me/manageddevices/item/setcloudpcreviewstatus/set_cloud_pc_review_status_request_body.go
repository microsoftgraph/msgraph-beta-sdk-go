package setcloudpcreviewstatus

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// SetCloudPcReviewStatusRequestBody provides operations to call the setCloudPcReviewStatus method.
type SetCloudPcReviewStatusRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The reviewStatus property
    reviewStatus ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcReviewStatusable
}
// NewSetCloudPcReviewStatusRequestBody instantiates a new setCloudPcReviewStatusRequestBody and sets the default values.
func NewSetCloudPcReviewStatusRequestBody()(*SetCloudPcReviewStatusRequestBody) {
    m := &SetCloudPcReviewStatusRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateSetCloudPcReviewStatusRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSetCloudPcReviewStatusRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSetCloudPcReviewStatusRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SetCloudPcReviewStatusRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SetCloudPcReviewStatusRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["reviewStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCloudPcReviewStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReviewStatus(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcReviewStatusable))
        }
        return nil
    }
    return res
}
// GetReviewStatus gets the reviewStatus property value. The reviewStatus property
func (m *SetCloudPcReviewStatusRequestBody) GetReviewStatus()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcReviewStatusable) {
    if m == nil {
        return nil
    } else {
        return m.reviewStatus
    }
}
// Serialize serializes information the current object
func (m *SetCloudPcReviewStatusRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("reviewStatus", m.GetReviewStatus())
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
func (m *SetCloudPcReviewStatusRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetReviewStatus sets the reviewStatus property value. The reviewStatus property
func (m *SetCloudPcReviewStatusRequestBody) SetReviewStatus(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcReviewStatusable)() {
    if m != nil {
        m.reviewStatus = value
    }
}

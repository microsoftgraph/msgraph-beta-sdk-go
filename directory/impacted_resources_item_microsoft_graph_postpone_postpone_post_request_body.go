package directory

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ImpactedResourcesItemMicrosoftGraphPostponePostponePostRequestBody 
type ImpactedResourcesItemMicrosoftGraphPostponePostponePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The postponeUntilDateTime property
    postponeUntilDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewImpactedResourcesItemMicrosoftGraphPostponePostponePostRequestBody instantiates a new ImpactedResourcesItemMicrosoftGraphPostponePostponePostRequestBody and sets the default values.
func NewImpactedResourcesItemMicrosoftGraphPostponePostponePostRequestBody()(*ImpactedResourcesItemMicrosoftGraphPostponePostponePostRequestBody) {
    m := &ImpactedResourcesItemMicrosoftGraphPostponePostponePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateImpactedResourcesItemMicrosoftGraphPostponePostponePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateImpactedResourcesItemMicrosoftGraphPostponePostponePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewImpactedResourcesItemMicrosoftGraphPostponePostponePostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ImpactedResourcesItemMicrosoftGraphPostponePostponePostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ImpactedResourcesItemMicrosoftGraphPostponePostponePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["postponeUntilDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPostponeUntilDateTime(val)
        }
        return nil
    }
    return res
}
// GetPostponeUntilDateTime gets the postponeUntilDateTime property value. The postponeUntilDateTime property
func (m *ImpactedResourcesItemMicrosoftGraphPostponePostponePostRequestBody) GetPostponeUntilDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.postponeUntilDateTime
}
// Serialize serializes information the current object
func (m *ImpactedResourcesItemMicrosoftGraphPostponePostponePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteTimeValue("postponeUntilDateTime", m.GetPostponeUntilDateTime())
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
func (m *ImpactedResourcesItemMicrosoftGraphPostponePostponePostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetPostponeUntilDateTime sets the postponeUntilDateTime property value. The postponeUntilDateTime property
func (m *ImpactedResourcesItemMicrosoftGraphPostponePostponePostRequestBody) SetPostponeUntilDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.postponeUntilDateTime = value
}

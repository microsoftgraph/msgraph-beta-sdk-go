package security

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// RetentionEventStatus 
type RetentionEventStatus struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The error if the status is not successful.
    error ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PublicErrorable
    // The OdataType property
    odataType *string
    // The status of the distribution. The possible values are: pending, error, success, notAvaliable.
    status *EventStatusType
}
// NewRetentionEventStatus instantiates a new retentionEventStatus and sets the default values.
func NewRetentionEventStatus()(*RetentionEventStatus) {
    m := &RetentionEventStatus{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.security.retentionEventStatus";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateRetentionEventStatusFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRetentionEventStatusFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRetentionEventStatus(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RetentionEventStatus) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetError gets the error property value. The error if the status is not successful.
func (m *RetentionEventStatus) GetError()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PublicErrorable) {
    return m.error
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RetentionEventStatus) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["error"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePublicErrorFromDiscriminatorValue , m.SetError)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["status"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseEventStatusType , m.SetStatus)
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *RetentionEventStatus) GetOdataType()(*string) {
    return m.odataType
}
// GetStatus gets the status property value. The status of the distribution. The possible values are: pending, error, success, notAvaliable.
func (m *RetentionEventStatus) GetStatus()(*EventStatusType) {
    return m.status
}
// Serialize serializes information the current object
func (m *RetentionEventStatus) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("error", m.GetError())
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
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err := writer.WriteStringValue("status", &cast)
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
func (m *RetentionEventStatus) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetError sets the error property value. The error if the status is not successful.
func (m *RetentionEventStatus) SetError(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PublicErrorable)() {
    m.error = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *RetentionEventStatus) SetOdataType(value *string)() {
    m.odataType = value
}
// SetStatus sets the status property value. The status of the distribution. The possible values are: pending, error, success, notAvaliable.
func (m *RetentionEventStatus) SetStatus(value *EventStatusType)() {
    m.status = value
}

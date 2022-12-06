package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// HasPayloadLinkResultItem a class containing the result of HasPayloadLinks action.
type HasPayloadLinkResultItem struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Exception information indicates if check for this item was successful or not.Empty string for no error.
    error *string
    // Indicate whether a payload has any link or not.
    hasLink *bool
    // The OdataType property
    odataType *string
    // Key of the Payload, In the format of Guid.
    payloadId *string
    // The reason where the link comes from.
    sources []DeviceAndAppManagementAssignmentSource
}
// NewHasPayloadLinkResultItem instantiates a new hasPayloadLinkResultItem and sets the default values.
func NewHasPayloadLinkResultItem()(*HasPayloadLinkResultItem) {
    m := &HasPayloadLinkResultItem{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateHasPayloadLinkResultItemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateHasPayloadLinkResultItemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewHasPayloadLinkResultItem(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *HasPayloadLinkResultItem) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetError gets the error property value. Exception information indicates if check for this item was successful or not.Empty string for no error.
func (m *HasPayloadLinkResultItem) GetError()(*string) {
    return m.error
}
// GetFieldDeserializers the deserialization information for the current model
func (m *HasPayloadLinkResultItem) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["error"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetError)
    res["hasLink"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetHasLink)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["payloadId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetPayloadId)
    res["sources"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfEnumValues(ParseDeviceAndAppManagementAssignmentSource , m.SetSources)
    return res
}
// GetHasLink gets the hasLink property value. Indicate whether a payload has any link or not.
func (m *HasPayloadLinkResultItem) GetHasLink()(*bool) {
    return m.hasLink
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *HasPayloadLinkResultItem) GetOdataType()(*string) {
    return m.odataType
}
// GetPayloadId gets the payloadId property value. Key of the Payload, In the format of Guid.
func (m *HasPayloadLinkResultItem) GetPayloadId()(*string) {
    return m.payloadId
}
// GetSources gets the sources property value. The reason where the link comes from.
func (m *HasPayloadLinkResultItem) GetSources()([]DeviceAndAppManagementAssignmentSource) {
    return m.sources
}
// Serialize serializes information the current object
func (m *HasPayloadLinkResultItem) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("error", m.GetError())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("hasLink", m.GetHasLink())
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
    {
        err := writer.WriteStringValue("payloadId", m.GetPayloadId())
        if err != nil {
            return err
        }
    }
    if m.GetSources() != nil {
        err := writer.WriteCollectionOfStringValues("sources", SerializeDeviceAndAppManagementAssignmentSource(m.GetSources()))
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
func (m *HasPayloadLinkResultItem) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetError sets the error property value. Exception information indicates if check for this item was successful or not.Empty string for no error.
func (m *HasPayloadLinkResultItem) SetError(value *string)() {
    m.error = value
}
// SetHasLink sets the hasLink property value. Indicate whether a payload has any link or not.
func (m *HasPayloadLinkResultItem) SetHasLink(value *bool)() {
    m.hasLink = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *HasPayloadLinkResultItem) SetOdataType(value *string)() {
    m.odataType = value
}
// SetPayloadId sets the payloadId property value. Key of the Payload, In the format of Guid.
func (m *HasPayloadLinkResultItem) SetPayloadId(value *string)() {
    m.payloadId = value
}
// SetSources sets the sources property value. The reason where the link comes from.
func (m *HasPayloadLinkResultItem) SetSources(value []DeviceAndAppManagementAssignmentSource)() {
    m.sources = value
}

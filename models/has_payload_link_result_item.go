package models

import (
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
    // Key of the Payload, In the format of Guid.
    payloadId *string
    // The reason where the link comes from.
    sources []string
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
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetError gets the error property value. Exception information indicates if check for this item was successful or not.Empty string for no error.
func (m *HasPayloadLinkResultItem) GetError()(*string) {
    if m == nil {
        return nil
    } else {
        return m.error
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *HasPayloadLinkResultItem) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["error"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetError(val)
        }
        return nil
    }
    res["hasLink"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHasLink(val)
        }
        return nil
    }
    res["payloadId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPayloadId(val)
        }
        return nil
    }
    res["sources"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetSources(res)
        }
        return nil
    }
    return res
}
// GetHasLink gets the hasLink property value. Indicate whether a payload has any link or not.
func (m *HasPayloadLinkResultItem) GetHasLink()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hasLink
    }
}
// GetPayloadId gets the payloadId property value. Key of the Payload, In the format of Guid.
func (m *HasPayloadLinkResultItem) GetPayloadId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.payloadId
    }
}
// GetSources gets the sources property value. The reason where the link comes from.
func (m *HasPayloadLinkResultItem) GetSources()([]string) {
    if m == nil {
        return nil
    } else {
        return m.sources
    }
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
        err := writer.WriteStringValue("payloadId", m.GetPayloadId())
        if err != nil {
            return err
        }
    }
    if m.GetSources() != nil {
        err := writer.WriteCollectionOfStringValues("sources", m.GetSources())
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
    if m != nil {
        m.additionalData = value
    }
}
// SetError sets the error property value. Exception information indicates if check for this item was successful or not.Empty string for no error.
func (m *HasPayloadLinkResultItem) SetError(value *string)() {
    if m != nil {
        m.error = value
    }
}
// SetHasLink sets the hasLink property value. Indicate whether a payload has any link or not.
func (m *HasPayloadLinkResultItem) SetHasLink(value *bool)() {
    if m != nil {
        m.hasLink = value
    }
}
// SetPayloadId sets the payloadId property value. Key of the Payload, In the format of Guid.
func (m *HasPayloadLinkResultItem) SetPayloadId(value *string)() {
    if m != nil {
        m.payloadId = value
    }
}
// SetSources sets the sources property value. The reason where the link comes from.
func (m *HasPayloadLinkResultItem) SetSources(value []string)() {
    if m != nil {
        m.sources = value
    }
}

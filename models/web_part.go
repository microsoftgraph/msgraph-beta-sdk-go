package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WebPart 
type WebPart struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The required properties for the webPart (varies by webPart)
    data SitePageDataable
    // The OdataType property
    odataType *string
    // A unique identifier specifying the webPart type. Read-only.
    type_escaped *string
}
// NewWebPart instantiates a new webPart and sets the default values.
func NewWebPart()(*WebPart) {
    m := &WebPart{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateWebPartFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWebPartFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWebPart(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WebPart) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetData gets the data property value. The required properties for the webPart (varies by webPart)
func (m *WebPart) GetData()(SitePageDataable) {
    return m.data
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WebPart) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["data"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSitePageDataFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetData(val.(SitePageDataable))
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
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *WebPart) GetOdataType()(*string) {
    return m.odataType
}
// GetType gets the type property value. A unique identifier specifying the webPart type. Read-only.
func (m *WebPart) GetType()(*string) {
    return m.type_escaped
}
// Serialize serializes information the current object
func (m *WebPart) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("data", m.GetData())
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
        err := writer.WriteStringValue("type", m.GetType())
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
func (m *WebPart) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetData sets the data property value. The required properties for the webPart (varies by webPart)
func (m *WebPart) SetData(value SitePageDataable)() {
    m.data = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *WebPart) SetOdataType(value *string)() {
    m.odataType = value
}
// SetType sets the type property value. A unique identifier specifying the webPart type. Read-only.
func (m *WebPart) SetType(value *string)() {
    m.type_escaped = value
}

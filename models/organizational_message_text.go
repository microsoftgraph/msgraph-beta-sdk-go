package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OrganizationalMessageText contains the text that will be displayed to users for a particular variant
type OrganizationalMessageText struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Indicates the text that will be displayed on the button of the message. This field applies to the softLanding surface
    buttonText *string
    // Indicates the url that the user will be directed to when the message is clicked
    clickUrl *string
    // Indicates the message that will be displayed
    message *string
    // The OdataType property
    odataType *string
    // Indicates the title that will be displayed
    title *string
}
// NewOrganizationalMessageText instantiates a new organizationalMessageText and sets the default values.
func NewOrganizationalMessageText()(*OrganizationalMessageText) {
    m := &OrganizationalMessageText{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.organizationalMessageText";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateOrganizationalMessageTextFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOrganizationalMessageTextFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOrganizationalMessageText(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OrganizationalMessageText) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetButtonText gets the buttonText property value. Indicates the text that will be displayed on the button of the message. This field applies to the softLanding surface
func (m *OrganizationalMessageText) GetButtonText()(*string) {
    return m.buttonText
}
// GetClickUrl gets the clickUrl property value. Indicates the url that the user will be directed to when the message is clicked
func (m *OrganizationalMessageText) GetClickUrl()(*string) {
    return m.clickUrl
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OrganizationalMessageText) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["buttonText"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetButtonText(val)
        }
        return nil
    }
    res["clickUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClickUrl(val)
        }
        return nil
    }
    res["message"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMessage(val)
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
    res["title"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTitle(val)
        }
        return nil
    }
    return res
}
// GetMessage gets the message property value. Indicates the message that will be displayed
func (m *OrganizationalMessageText) GetMessage()(*string) {
    return m.message
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *OrganizationalMessageText) GetOdataType()(*string) {
    return m.odataType
}
// GetTitle gets the title property value. Indicates the title that will be displayed
func (m *OrganizationalMessageText) GetTitle()(*string) {
    return m.title
}
// Serialize serializes information the current object
func (m *OrganizationalMessageText) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("buttonText", m.GetButtonText())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("clickUrl", m.GetClickUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("message", m.GetMessage())
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
        err := writer.WriteStringValue("title", m.GetTitle())
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
func (m *OrganizationalMessageText) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetButtonText sets the buttonText property value. Indicates the text that will be displayed on the button of the message. This field applies to the softLanding surface
func (m *OrganizationalMessageText) SetButtonText(value *string)() {
    m.buttonText = value
}
// SetClickUrl sets the clickUrl property value. Indicates the url that the user will be directed to when the message is clicked
func (m *OrganizationalMessageText) SetClickUrl(value *string)() {
    m.clickUrl = value
}
// SetMessage sets the message property value. Indicates the message that will be displayed
func (m *OrganizationalMessageText) SetMessage(value *string)() {
    m.message = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *OrganizationalMessageText) SetOdataType(value *string)() {
    m.odataType = value
}
// SetTitle sets the title property value. Indicates the title that will be displayed
func (m *OrganizationalMessageText) SetTitle(value *string)() {
    m.title = value
}

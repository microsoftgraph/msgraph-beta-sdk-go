package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OrganizationalMessageLogo contains the logo's binary content or a url to the logo's downloadable location. Either both logo and contentType contain valid values or logoCdnUrl contains a valid url
type OrganizationalMessageLogo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The content type of the logo that is contained in the logo array. This is null when logoCdnUrl is used to send the logo. Possible values are: png.
    contentType *OrganizationalMessageLogoType
    // The binary contents of the logo. This is null when logoCdnUrl is used to send the logo
    logo []byte
    // The url at which the logo resides. This is null when logo and contentType are used to send the logo
    logoCdnUrl *string
    // The OdataType property
    odataType *string
}
// NewOrganizationalMessageLogo instantiates a new organizationalMessageLogo and sets the default values.
func NewOrganizationalMessageLogo()(*OrganizationalMessageLogo) {
    m := &OrganizationalMessageLogo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.organizationalMessageLogo";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateOrganizationalMessageLogoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOrganizationalMessageLogoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOrganizationalMessageLogo(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OrganizationalMessageLogo) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetContentType gets the contentType property value. The content type of the logo that is contained in the logo array. This is null when logoCdnUrl is used to send the logo. Possible values are: png.
func (m *OrganizationalMessageLogo) GetContentType()(*OrganizationalMessageLogoType) {
    return m.contentType
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OrganizationalMessageLogo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["contentType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseOrganizationalMessageLogoType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentType(val.(*OrganizationalMessageLogoType))
        }
        return nil
    }
    res["logo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLogo(val)
        }
        return nil
    }
    res["logoCdnUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLogoCdnUrl(val)
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
    return res
}
// GetLogo gets the logo property value. The binary contents of the logo. This is null when logoCdnUrl is used to send the logo
func (m *OrganizationalMessageLogo) GetLogo()([]byte) {
    return m.logo
}
// GetLogoCdnUrl gets the logoCdnUrl property value. The url at which the logo resides. This is null when logo and contentType are used to send the logo
func (m *OrganizationalMessageLogo) GetLogoCdnUrl()(*string) {
    return m.logoCdnUrl
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *OrganizationalMessageLogo) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *OrganizationalMessageLogo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetContentType() != nil {
        cast := (*m.GetContentType()).String()
        err := writer.WriteStringValue("contentType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteByteArrayValue("logo", m.GetLogo())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("logoCdnUrl", m.GetLogoCdnUrl())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OrganizationalMessageLogo) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetContentType sets the contentType property value. The content type of the logo that is contained in the logo array. This is null when logoCdnUrl is used to send the logo. Possible values are: png.
func (m *OrganizationalMessageLogo) SetContentType(value *OrganizationalMessageLogoType)() {
    m.contentType = value
}
// SetLogo sets the logo property value. The binary contents of the logo. This is null when logoCdnUrl is used to send the logo
func (m *OrganizationalMessageLogo) SetLogo(value []byte)() {
    m.logo = value
}
// SetLogoCdnUrl sets the logoCdnUrl property value. The url at which the logo resides. This is null when logo and contentType are used to send the logo
func (m *OrganizationalMessageLogo) SetLogoCdnUrl(value *string)() {
    m.logoCdnUrl = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *OrganizationalMessageLogo) SetOdataType(value *string)() {
    m.odataType = value
}

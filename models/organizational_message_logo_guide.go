package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OrganizationalMessageLogoGuide example logo and its size requirements
type OrganizationalMessageLogoGuide struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The asset name is the key for this specific logo. This is used to compute the required token when accessing the logoCdnUrl to fetch the logo
    assetName *string
    // The required size dimensions of the logo
    dimensions OrganizationalMessageLogoDimensionsable
    // The url at which the logo resides
    logoCdnUrl *string
    // The OdataType property
    odataType *string
}
// NewOrganizationalMessageLogoGuide instantiates a new organizationalMessageLogoGuide and sets the default values.
func NewOrganizationalMessageLogoGuide()(*OrganizationalMessageLogoGuide) {
    m := &OrganizationalMessageLogoGuide{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.organizationalMessageLogoGuide";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateOrganizationalMessageLogoGuideFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOrganizationalMessageLogoGuideFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOrganizationalMessageLogoGuide(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OrganizationalMessageLogoGuide) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAssetName gets the assetName property value. The asset name is the key for this specific logo. This is used to compute the required token when accessing the logoCdnUrl to fetch the logo
func (m *OrganizationalMessageLogoGuide) GetAssetName()(*string) {
    return m.assetName
}
// GetDimensions gets the dimensions property value. The required size dimensions of the logo
func (m *OrganizationalMessageLogoGuide) GetDimensions()(OrganizationalMessageLogoDimensionsable) {
    return m.dimensions
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OrganizationalMessageLogoGuide) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["assetName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssetName(val)
        }
        return nil
    }
    res["dimensions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateOrganizationalMessageLogoDimensionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDimensions(val.(OrganizationalMessageLogoDimensionsable))
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
// GetLogoCdnUrl gets the logoCdnUrl property value. The url at which the logo resides
func (m *OrganizationalMessageLogoGuide) GetLogoCdnUrl()(*string) {
    return m.logoCdnUrl
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *OrganizationalMessageLogoGuide) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *OrganizationalMessageLogoGuide) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("assetName", m.GetAssetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("dimensions", m.GetDimensions())
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
func (m *OrganizationalMessageLogoGuide) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAssetName sets the assetName property value. The asset name is the key for this specific logo. This is used to compute the required token when accessing the logoCdnUrl to fetch the logo
func (m *OrganizationalMessageLogoGuide) SetAssetName(value *string)() {
    m.assetName = value
}
// SetDimensions sets the dimensions property value. The required size dimensions of the logo
func (m *OrganizationalMessageLogoGuide) SetDimensions(value OrganizationalMessageLogoDimensionsable)() {
    m.dimensions = value
}
// SetLogoCdnUrl sets the logoCdnUrl property value. The url at which the logo resides
func (m *OrganizationalMessageLogoGuide) SetLogoCdnUrl(value *string)() {
    m.logoCdnUrl = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *OrganizationalMessageLogoGuide) SetOdataType(value *string)() {
    m.odataType = value
}

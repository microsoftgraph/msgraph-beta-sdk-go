package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OrganizationalMessageContent contains the entire content of the message that will be displayed to the clients
type OrganizationalMessageContent struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The ID of the guided content that this content is using
    guidedContentId *string
    // The logo that will be displayed to the clients. This will contain ether the binary contents of the logo or a url to the logo's location
    logoInfo OrganizationalMessageLogoable
    // The OdataType property
    odataType *string
    // Contains the different types of text content that can be displayed to customers along with their localized values
    placementDetails []OrganizationalMessagePlacementDetailable
}
// NewOrganizationalMessageContent instantiates a new organizationalMessageContent and sets the default values.
func NewOrganizationalMessageContent()(*OrganizationalMessageContent) {
    m := &OrganizationalMessageContent{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.organizationalMessageContent";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateOrganizationalMessageContentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOrganizationalMessageContentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOrganizationalMessageContent(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OrganizationalMessageContent) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OrganizationalMessageContent) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["guidedContentId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGuidedContentId(val)
        }
        return nil
    }
    res["logoInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateOrganizationalMessageLogoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLogoInfo(val.(OrganizationalMessageLogoable))
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
    res["placementDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateOrganizationalMessagePlacementDetailFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OrganizationalMessagePlacementDetailable, len(val))
            for i, v := range val {
                res[i] = v.(OrganizationalMessagePlacementDetailable)
            }
            m.SetPlacementDetails(res)
        }
        return nil
    }
    return res
}
// GetGuidedContentId gets the guidedContentId property value. The ID of the guided content that this content is using
func (m *OrganizationalMessageContent) GetGuidedContentId()(*string) {
    return m.guidedContentId
}
// GetLogoInfo gets the logoInfo property value. The logo that will be displayed to the clients. This will contain ether the binary contents of the logo or a url to the logo's location
func (m *OrganizationalMessageContent) GetLogoInfo()(OrganizationalMessageLogoable) {
    return m.logoInfo
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *OrganizationalMessageContent) GetOdataType()(*string) {
    return m.odataType
}
// GetPlacementDetails gets the placementDetails property value. Contains the different types of text content that can be displayed to customers along with their localized values
func (m *OrganizationalMessageContent) GetPlacementDetails()([]OrganizationalMessagePlacementDetailable) {
    return m.placementDetails
}
// Serialize serializes information the current object
func (m *OrganizationalMessageContent) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("guidedContentId", m.GetGuidedContentId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("logoInfo", m.GetLogoInfo())
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
    if m.GetPlacementDetails() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPlacementDetails()))
        for i, v := range m.GetPlacementDetails() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("placementDetails", cast)
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
func (m *OrganizationalMessageContent) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetGuidedContentId sets the guidedContentId property value. The ID of the guided content that this content is using
func (m *OrganizationalMessageContent) SetGuidedContentId(value *string)() {
    m.guidedContentId = value
}
// SetLogoInfo sets the logoInfo property value. The logo that will be displayed to the clients. This will contain ether the binary contents of the logo or a url to the logo's location
func (m *OrganizationalMessageContent) SetLogoInfo(value OrganizationalMessageLogoable)() {
    m.logoInfo = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *OrganizationalMessageContent) SetOdataType(value *string)() {
    m.odataType = value
}
// SetPlacementDetails sets the placementDetails property value. Contains the different types of text content that can be displayed to customers along with their localized values
func (m *OrganizationalMessageContent) SetPlacementDetails(value []OrganizationalMessagePlacementDetailable)() {
    m.placementDetails = value
}

package me

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// CloudPCsItemMicrosoftGraphReprovisionReprovisionPostRequestBody 
type CloudPCsItemMicrosoftGraphReprovisionReprovisionPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The osVersion property
    osVersion *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcOperatingSystem
    // The userAccountType property
    userAccountType *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcUserAccountType
}
// NewCloudPCsItemMicrosoftGraphReprovisionReprovisionPostRequestBody instantiates a new CloudPCsItemMicrosoftGraphReprovisionReprovisionPostRequestBody and sets the default values.
func NewCloudPCsItemMicrosoftGraphReprovisionReprovisionPostRequestBody()(*CloudPCsItemMicrosoftGraphReprovisionReprovisionPostRequestBody) {
    m := &CloudPCsItemMicrosoftGraphReprovisionReprovisionPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCloudPCsItemMicrosoftGraphReprovisionReprovisionPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCloudPCsItemMicrosoftGraphReprovisionReprovisionPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudPCsItemMicrosoftGraphReprovisionReprovisionPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CloudPCsItemMicrosoftGraphReprovisionReprovisionPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CloudPCsItemMicrosoftGraphReprovisionReprovisionPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["osVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ParseCloudPcOperatingSystem)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsVersion(val.(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcOperatingSystem))
        }
        return nil
    }
    res["userAccountType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ParseCloudPcUserAccountType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserAccountType(val.(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcUserAccountType))
        }
        return nil
    }
    return res
}
// GetOsVersion gets the osVersion property value. The osVersion property
func (m *CloudPCsItemMicrosoftGraphReprovisionReprovisionPostRequestBody) GetOsVersion()(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcOperatingSystem) {
    return m.osVersion
}
// GetUserAccountType gets the userAccountType property value. The userAccountType property
func (m *CloudPCsItemMicrosoftGraphReprovisionReprovisionPostRequestBody) GetUserAccountType()(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcUserAccountType) {
    return m.userAccountType
}
// Serialize serializes information the current object
func (m *CloudPCsItemMicrosoftGraphReprovisionReprovisionPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetOsVersion() != nil {
        cast := (*m.GetOsVersion()).String()
        err := writer.WriteStringValue("osVersion", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserAccountType() != nil {
        cast := (*m.GetUserAccountType()).String()
        err := writer.WriteStringValue("userAccountType", &cast)
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
func (m *CloudPCsItemMicrosoftGraphReprovisionReprovisionPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetOsVersion sets the osVersion property value. The osVersion property
func (m *CloudPCsItemMicrosoftGraphReprovisionReprovisionPostRequestBody) SetOsVersion(value *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcOperatingSystem)() {
    m.osVersion = value
}
// SetUserAccountType sets the userAccountType property value. The userAccountType property
func (m *CloudPCsItemMicrosoftGraphReprovisionReprovisionPostRequestBody) SetUserAccountType(value *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcUserAccountType)() {
    m.userAccountType = value
}

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UsersItemCloudPCsItemReprovisionPostRequestBody provides operations to call the reprovision method.
type UsersItemCloudPCsItemReprovisionPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The osVersion property
    osVersion *CloudPcOperatingSystem
    // The userAccountType property
    userAccountType *CloudPcUserAccountType
}
// NewUsersItemCloudPCsItemReprovisionPostRequestBody instantiates a new UsersItemCloudPCsItemReprovisionPostRequestBody and sets the default values.
func NewUsersItemCloudPCsItemReprovisionPostRequestBody()(*UsersItemCloudPCsItemReprovisionPostRequestBody) {
    m := &UsersItemCloudPCsItemReprovisionPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateUsersItemCloudPCsItemReprovisionPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUsersItemCloudPCsItemReprovisionPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUsersItemCloudPCsItemReprovisionPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UsersItemCloudPCsItemReprovisionPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UsersItemCloudPCsItemReprovisionPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["osVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcOperatingSystem)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsVersion(val.(*CloudPcOperatingSystem))
        }
        return nil
    }
    res["userAccountType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcUserAccountType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserAccountType(val.(*CloudPcUserAccountType))
        }
        return nil
    }
    return res
}
// GetOsVersion gets the osVersion property value. The osVersion property
func (m *UsersItemCloudPCsItemReprovisionPostRequestBody) GetOsVersion()(*CloudPcOperatingSystem) {
    return m.osVersion
}
// GetUserAccountType gets the userAccountType property value. The userAccountType property
func (m *UsersItemCloudPCsItemReprovisionPostRequestBody) GetUserAccountType()(*CloudPcUserAccountType) {
    return m.userAccountType
}
// Serialize serializes information the current object
func (m *UsersItemCloudPCsItemReprovisionPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *UsersItemCloudPCsItemReprovisionPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetOsVersion sets the osVersion property value. The osVersion property
func (m *UsersItemCloudPCsItemReprovisionPostRequestBody) SetOsVersion(value *CloudPcOperatingSystem)() {
    m.osVersion = value
}
// SetUserAccountType sets the userAccountType property value. The userAccountType property
func (m *UsersItemCloudPCsItemReprovisionPostRequestBody) SetUserAccountType(value *CloudPcUserAccountType)() {
    m.userAccountType = value
}

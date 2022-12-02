package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamworkDevicesItemUpdateSoftwarePostRequestBody provides operations to call the updateSoftware method.
type TeamworkDevicesItemUpdateSoftwarePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The softwareType property
    softwareType *TeamworkSoftwareType
    // The softwareVersion property
    softwareVersion *string
}
// NewTeamworkDevicesItemUpdateSoftwarePostRequestBody instantiates a new TeamworkDevicesItemUpdateSoftwarePostRequestBody and sets the default values.
func NewTeamworkDevicesItemUpdateSoftwarePostRequestBody()(*TeamworkDevicesItemUpdateSoftwarePostRequestBody) {
    m := &TeamworkDevicesItemUpdateSoftwarePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateTeamworkDevicesItemUpdateSoftwarePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamworkDevicesItemUpdateSoftwarePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamworkDevicesItemUpdateSoftwarePostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamworkDevicesItemUpdateSoftwarePostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamworkDevicesItemUpdateSoftwarePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["softwareType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseTeamworkSoftwareType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSoftwareType(val.(*TeamworkSoftwareType))
        }
        return nil
    }
    res["softwareVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSoftwareVersion(val)
        }
        return nil
    }
    return res
}
// GetSoftwareType gets the softwareType property value. The softwareType property
func (m *TeamworkDevicesItemUpdateSoftwarePostRequestBody) GetSoftwareType()(*TeamworkSoftwareType) {
    return m.softwareType
}
// GetSoftwareVersion gets the softwareVersion property value. The softwareVersion property
func (m *TeamworkDevicesItemUpdateSoftwarePostRequestBody) GetSoftwareVersion()(*string) {
    return m.softwareVersion
}
// Serialize serializes information the current object
func (m *TeamworkDevicesItemUpdateSoftwarePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetSoftwareType() != nil {
        cast := (*m.GetSoftwareType()).String()
        err := writer.WriteStringValue("softwareType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("softwareVersion", m.GetSoftwareVersion())
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
func (m *TeamworkDevicesItemUpdateSoftwarePostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetSoftwareType sets the softwareType property value. The softwareType property
func (m *TeamworkDevicesItemUpdateSoftwarePostRequestBody) SetSoftwareType(value *TeamworkSoftwareType)() {
    m.softwareType = value
}
// SetSoftwareVersion sets the softwareVersion property value. The softwareVersion property
func (m *TeamworkDevicesItemUpdateSoftwarePostRequestBody) SetSoftwareVersion(value *string)() {
    m.softwareVersion = value
}

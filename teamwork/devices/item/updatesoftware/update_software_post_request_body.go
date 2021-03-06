package updatesoftware

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// UpdateSoftwarePostRequestBody provides operations to call the updateSoftware method.
type UpdateSoftwarePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The softwareType property
    softwareType *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TeamworkSoftwareType
    // The softwareVersion property
    softwareVersion *string
}
// NewUpdateSoftwarePostRequestBody instantiates a new updateSoftwarePostRequestBody and sets the default values.
func NewUpdateSoftwarePostRequestBody()(*UpdateSoftwarePostRequestBody) {
    m := &UpdateSoftwarePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateUpdateSoftwarePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUpdateSoftwarePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUpdateSoftwarePostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UpdateSoftwarePostRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UpdateSoftwarePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["softwareType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ParseTeamworkSoftwareType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSoftwareType(val.(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TeamworkSoftwareType))
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
func (m *UpdateSoftwarePostRequestBody) GetSoftwareType()(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TeamworkSoftwareType) {
    if m == nil {
        return nil
    } else {
        return m.softwareType
    }
}
// GetSoftwareVersion gets the softwareVersion property value. The softwareVersion property
func (m *UpdateSoftwarePostRequestBody) GetSoftwareVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.softwareVersion
    }
}
// Serialize serializes information the current object
func (m *UpdateSoftwarePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *UpdateSoftwarePostRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetSoftwareType sets the softwareType property value. The softwareType property
func (m *UpdateSoftwarePostRequestBody) SetSoftwareType(value *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TeamworkSoftwareType)() {
    if m != nil {
        m.softwareType = value
    }
}
// SetSoftwareVersion sets the softwareVersion property value. The softwareVersion property
func (m *UpdateSoftwarePostRequestBody) SetSoftwareVersion(value *string)() {
    if m != nil {
        m.softwareVersion = value
    }
}

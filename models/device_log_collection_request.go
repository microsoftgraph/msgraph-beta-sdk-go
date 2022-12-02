package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceLogCollectionRequest windows Log Collection request entity.
type DeviceLogCollectionRequest struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The unique identifier
    id *string
    // The OdataType property
    odataType *string
    // Enum for the template type used for collecting logs
    templateType *DeviceLogCollectionTemplateType
}
// NewDeviceLogCollectionRequest instantiates a new deviceLogCollectionRequest and sets the default values.
func NewDeviceLogCollectionRequest()(*DeviceLogCollectionRequest) {
    m := &DeviceLogCollectionRequest{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceLogCollectionRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceLogCollectionRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceLogCollectionRequest(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceLogCollectionRequest) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceLogCollectionRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
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
    res["templateType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceLogCollectionTemplateType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTemplateType(val.(*DeviceLogCollectionTemplateType))
        }
        return nil
    }
    return res
}
// GetId gets the id property value. The unique identifier
func (m *DeviceLogCollectionRequest) GetId()(*string) {
    return m.id
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *DeviceLogCollectionRequest) GetOdataType()(*string) {
    return m.odataType
}
// GetTemplateType gets the templateType property value. Enum for the template type used for collecting logs
func (m *DeviceLogCollectionRequest) GetTemplateType()(*DeviceLogCollectionTemplateType) {
    return m.templateType
}
// Serialize serializes information the current object
func (m *DeviceLogCollectionRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("id", m.GetId())
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
    if m.GetTemplateType() != nil {
        cast := (*m.GetTemplateType()).String()
        err := writer.WriteStringValue("templateType", &cast)
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
func (m *DeviceLogCollectionRequest) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetId sets the id property value. The unique identifier
func (m *DeviceLogCollectionRequest) SetId(value *string)() {
    m.id = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DeviceLogCollectionRequest) SetOdataType(value *string)() {
    m.odataType = value
}
// SetTemplateType sets the templateType property value. Enum for the template type used for collecting logs
func (m *DeviceLogCollectionRequest) SetTemplateType(value *DeviceLogCollectionTemplateType)() {
    m.templateType = value
}

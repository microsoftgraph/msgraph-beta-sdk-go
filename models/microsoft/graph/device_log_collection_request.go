package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type DeviceLogCollectionRequest struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The unique identifier
    id *string;
    // The template type that is sent with the collection request. Possible values are: predefined.
    templateType *DeviceLogCollectionTemplateType;
}
// Instantiates a new deviceLogCollectionRequest and sets the default values.
func NewDeviceLogCollectionRequest()(*DeviceLogCollectionRequest) {
    m := &DeviceLogCollectionRequest{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceLogCollectionRequest) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the id property value. The unique identifier
func (m *DeviceLogCollectionRequest) GetId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.id
    }
}
// Gets the templateType property value. The template type that is sent with the collection request. Possible values are: predefined.
func (m *DeviceLogCollectionRequest) GetTemplateType()(*DeviceLogCollectionTemplateType) {
    if m == nil {
        return nil
    } else {
        return m.templateType
    }
}
// The deserialization information for the current model
func (m *DeviceLogCollectionRequest) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["id"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["templateType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceLogCollectionTemplateType)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(DeviceLogCollectionTemplateType)
            m.SetTemplateType(&cast)
        }
        return nil
    }
    return res
}
func (m *DeviceLogCollectionRequest) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *DeviceLogCollectionRequest) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    if m.GetTemplateType() != nil {
        cast := m.GetTemplateType().String()
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *DeviceLogCollectionRequest) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the id property value. The unique identifier
// Parameters:
//  - value : Value to set for the id property.
func (m *DeviceLogCollectionRequest) SetId(value *string)() {
    m.id = value
}
// Sets the templateType property value. The template type that is sent with the collection request. Possible values are: predefined.
// Parameters:
//  - value : Value to set for the templateType property.
func (m *DeviceLogCollectionRequest) SetTemplateType(value *DeviceLogCollectionTemplateType)() {
    m.templateType = value
}

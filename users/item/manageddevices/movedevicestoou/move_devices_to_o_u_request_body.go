package movedevicestoou

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type MoveDevicesToOURequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    deviceIds []string;
    // 
    organizationalUnitPath *string;
}
// Instantiates a new moveDevicesToOURequestBody and sets the default values.
func NewMoveDevicesToOURequestBody()(*MoveDevicesToOURequestBody) {
    m := &MoveDevicesToOURequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MoveDevicesToOURequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the deviceIds property value. 
func (m *MoveDevicesToOURequestBody) GetDeviceIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.deviceIds
    }
}
// Gets the organizationalUnitPath property value. 
func (m *MoveDevicesToOURequestBody) GetOrganizationalUnitPath()(*string) {
    if m == nil {
        return nil
    } else {
        return m.organizationalUnitPath
    }
}
// The deserialization information for the current model
func (m *MoveDevicesToOURequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["deviceIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = *(v.(*string))
        }
        m.SetDeviceIds(res)
        return nil
    }
    res["organizationalUnitPath"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOrganizationalUnitPath(val)
        return nil
    }
    return res
}
func (m *MoveDevicesToOURequestBody) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *MoveDevicesToOURequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteCollectionOfStringValues("deviceIds", m.GetDeviceIds())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("organizationalUnitPath", m.GetOrganizationalUnitPath())
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
func (m *MoveDevicesToOURequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the deviceIds property value. 
// Parameters:
//  - value : Value to set for the deviceIds property.
func (m *MoveDevicesToOURequestBody) SetDeviceIds(value []string)() {
    m.deviceIds = value
}
// Sets the organizationalUnitPath property value. 
// Parameters:
//  - value : Value to set for the organizationalUnitPath property.
func (m *MoveDevicesToOURequestBody) SetOrganizationalUnitPath(value *string)() {
    m.organizationalUnitPath = value
}

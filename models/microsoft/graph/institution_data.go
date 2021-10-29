package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type InstitutionData struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Short description of the institution the user studied at.
    description *string;
    // Name of the institution the user studied at.
    displayName *string;
    // Address or location of the institute.
    location *PhysicalAddress;
    // Link to the institution or department homepage.
    webUrl *string;
}
// Instantiates a new institutionData and sets the default values.
func NewInstitutionData()(*InstitutionData) {
    m := &InstitutionData{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *InstitutionData) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the description property value. Short description of the institution the user studied at.
func (m *InstitutionData) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the displayName property value. Name of the institution the user studied at.
func (m *InstitutionData) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the location property value. Address or location of the institute.
func (m *InstitutionData) GetLocation()(*PhysicalAddress) {
    if m == nil {
        return nil
    } else {
        return m.location
    }
}
// Gets the webUrl property value. Link to the institution or department homepage.
func (m *InstitutionData) GetWebUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.webUrl
    }
}
// The deserialization information for the current model
func (m *InstitutionData) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDescription(val)
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["location"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPhysicalAddress() })
        if err != nil {
            return err
        }
        m.SetLocation(val.(*PhysicalAddress))
        return nil
    }
    res["webUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetWebUrl(val)
        return nil
    }
    return res
}
func (m *InstitutionData) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *InstitutionData) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("location", m.GetLocation())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("webUrl", m.GetWebUrl())
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
func (m *InstitutionData) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the description property value. Short description of the institution the user studied at.
// Parameters:
//  - value : Value to set for the description property.
func (m *InstitutionData) SetDescription(value *string)() {
    m.description = value
}
// Sets the displayName property value. Name of the institution the user studied at.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *InstitutionData) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the location property value. Address or location of the institute.
// Parameters:
//  - value : Value to set for the location property.
func (m *InstitutionData) SetLocation(value *PhysicalAddress)() {
    m.location = value
}
// Sets the webUrl property value. Link to the institution or department homepage.
// Parameters:
//  - value : Value to set for the webUrl property.
func (m *InstitutionData) SetWebUrl(value *string)() {
    m.webUrl = value
}

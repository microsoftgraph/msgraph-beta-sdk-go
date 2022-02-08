package updatesoftware

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// UpdateSoftwareRequestBody 
type UpdateSoftwareRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    softwareType *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TeamworkSoftwareType;
    // 
    softwareVersion *string;
}
// NewUpdateSoftwareRequestBody instantiates a new updateSoftwareRequestBody and sets the default values.
func NewUpdateSoftwareRequestBody()(*UpdateSoftwareRequestBody) {
    m := &UpdateSoftwareRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UpdateSoftwareRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetSoftwareType gets the softwareType property value. 
func (m *UpdateSoftwareRequestBody) GetSoftwareType()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TeamworkSoftwareType) {
    if m == nil {
        return nil
    } else {
        return m.softwareType
    }
}
// GetSoftwareVersion gets the softwareVersion property value. 
func (m *UpdateSoftwareRequestBody) GetSoftwareVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.softwareVersion
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UpdateSoftwareRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["softwareType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ParseTeamworkSoftwareType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSoftwareType(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TeamworkSoftwareType))
        }
        return nil
    }
    res["softwareVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
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
func (m *UpdateSoftwareRequestBody) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *UpdateSoftwareRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
func (m *UpdateSoftwareRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetSoftwareType sets the softwareType property value. 
func (m *UpdateSoftwareRequestBody) SetSoftwareType(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TeamworkSoftwareType)() {
    if m != nil {
        m.softwareType = value
    }
}
// SetSoftwareVersion sets the softwareVersion property value. 
func (m *UpdateSoftwareRequestBody) SetSoftwareVersion(value *string)() {
    if m != nil {
        m.softwareVersion = value
    }
}

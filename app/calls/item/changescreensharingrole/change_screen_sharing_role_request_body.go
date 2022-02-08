package changescreensharingrole

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// ChangeScreenSharingRoleRequestBody 
type ChangeScreenSharingRoleRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    role *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ScreenSharingRole;
}
// NewChangeScreenSharingRoleRequestBody instantiates a new changeScreenSharingRoleRequestBody and sets the default values.
func NewChangeScreenSharingRoleRequestBody()(*ChangeScreenSharingRoleRequestBody) {
    m := &ChangeScreenSharingRoleRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ChangeScreenSharingRoleRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetRole gets the role property value. 
func (m *ChangeScreenSharingRoleRequestBody) GetRole()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ScreenSharingRole) {
    if m == nil {
        return nil
    } else {
        return m.role
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ChangeScreenSharingRoleRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["role"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ParseScreenSharingRole)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRole(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ScreenSharingRole))
        }
        return nil
    }
    return res
}
func (m *ChangeScreenSharingRoleRequestBody) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ChangeScreenSharingRoleRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetRole() != nil {
        cast := (*m.GetRole()).String()
        err := writer.WriteStringValue("role", &cast)
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
func (m *ChangeScreenSharingRoleRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetRole sets the role property value. 
func (m *ChangeScreenSharingRoleRequestBody) SetRole(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ScreenSharingRole)() {
    if m != nil {
        m.role = value
    }
}

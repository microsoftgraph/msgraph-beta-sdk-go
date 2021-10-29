package assignedaccessmultimodeprofiles

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// 
type AssignedAccessMultiModeProfilesRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    assignedAccessMultiModeProfiles []i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WindowsAssignedAccessProfile;
}
// Instantiates a new assignedAccessMultiModeProfilesRequestBody and sets the default values.
func NewAssignedAccessMultiModeProfilesRequestBody()(*AssignedAccessMultiModeProfilesRequestBody) {
    m := &AssignedAccessMultiModeProfilesRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AssignedAccessMultiModeProfilesRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the assignedAccessMultiModeProfiles property value. 
func (m *AssignedAccessMultiModeProfilesRequestBody) GetAssignedAccessMultiModeProfiles()([]i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WindowsAssignedAccessProfile) {
    if m == nil {
        return nil
    } else {
        return m.assignedAccessMultiModeProfiles
    }
}
// The deserialization information for the current model
func (m *AssignedAccessMultiModeProfilesRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["assignedAccessMultiModeProfiles"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewWindowsAssignedAccessProfile() })
        if err != nil {
            return err
        }
        res := make([]i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WindowsAssignedAccessProfile, len(val))
        for i, v := range val {
            res[i] = *(v.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WindowsAssignedAccessProfile))
        }
        m.SetAssignedAccessMultiModeProfiles(res)
        return nil
    }
    return res
}
func (m *AssignedAccessMultiModeProfilesRequestBody) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *AssignedAccessMultiModeProfilesRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAssignedAccessMultiModeProfiles()))
        for i, v := range m.GetAssignedAccessMultiModeProfiles() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("assignedAccessMultiModeProfiles", cast)
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
func (m *AssignedAccessMultiModeProfilesRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the assignedAccessMultiModeProfiles property value. 
// Parameters:
//  - value : Value to set for the assignedAccessMultiModeProfiles property.
func (m *AssignedAccessMultiModeProfilesRequestBody) SetAssignedAccessMultiModeProfiles(value []i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WindowsAssignedAccessProfile)() {
    m.assignedAccessMultiModeProfiles = value
}

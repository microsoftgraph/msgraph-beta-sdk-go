package importdeviceidentitylist

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// 
type ImportDeviceIdentityListRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    importedDeviceIdentities []i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ImportedDeviceIdentity;
    // 
    overwriteImportedDeviceIdentities *bool;
}
// Instantiates a new importDeviceIdentityListRequestBody and sets the default values.
func NewImportDeviceIdentityListRequestBody()(*ImportDeviceIdentityListRequestBody) {
    m := &ImportDeviceIdentityListRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ImportDeviceIdentityListRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the importedDeviceIdentities property value. 
func (m *ImportDeviceIdentityListRequestBody) GetImportedDeviceIdentities()([]i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ImportedDeviceIdentity) {
    if m == nil {
        return nil
    } else {
        return m.importedDeviceIdentities
    }
}
// Gets the overwriteImportedDeviceIdentities property value. 
func (m *ImportDeviceIdentityListRequestBody) GetOverwriteImportedDeviceIdentities()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.overwriteImportedDeviceIdentities
    }
}
// The deserialization information for the current model
func (m *ImportDeviceIdentityListRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["importedDeviceIdentities"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewImportedDeviceIdentity() })
        if err != nil {
            return err
        }
        res := make([]i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ImportedDeviceIdentity, len(val))
        for i, v := range val {
            res[i] = *(v.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ImportedDeviceIdentity))
        }
        m.SetImportedDeviceIdentities(res)
        return nil
    }
    res["overwriteImportedDeviceIdentities"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetOverwriteImportedDeviceIdentities(val)
        return nil
    }
    return res
}
func (m *ImportDeviceIdentityListRequestBody) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ImportDeviceIdentityListRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetImportedDeviceIdentities()))
        for i, v := range m.GetImportedDeviceIdentities() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("importedDeviceIdentities", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("overwriteImportedDeviceIdentities", m.GetOverwriteImportedDeviceIdentities())
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
func (m *ImportDeviceIdentityListRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the importedDeviceIdentities property value. 
// Parameters:
//  - value : Value to set for the importedDeviceIdentities property.
func (m *ImportDeviceIdentityListRequestBody) SetImportedDeviceIdentities(value []i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ImportedDeviceIdentity)() {
    m.importedDeviceIdentities = value
}
// Sets the overwriteImportedDeviceIdentities property value. 
// Parameters:
//  - value : Value to set for the overwriteImportedDeviceIdentities property.
func (m *ImportDeviceIdentityListRequestBody) SetOverwriteImportedDeviceIdentities(value *bool)() {
    m.overwriteImportedDeviceIdentities = value
}

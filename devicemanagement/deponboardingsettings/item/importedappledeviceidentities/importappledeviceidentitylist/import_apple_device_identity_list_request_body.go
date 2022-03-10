package importappledeviceidentitylist

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// ImportAppleDeviceIdentityListRequestBody provides operations to call the importAppleDeviceIdentityList method.
type ImportAppleDeviceIdentityListRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    importedAppleDeviceIdentities []i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ImportedAppleDeviceIdentityable;
    // 
    overwriteImportedDeviceIdentities *bool;
}
// NewImportAppleDeviceIdentityListRequestBody instantiates a new importAppleDeviceIdentityListRequestBody and sets the default values.
func NewImportAppleDeviceIdentityListRequestBody()(*ImportAppleDeviceIdentityListRequestBody) {
    m := &ImportAppleDeviceIdentityListRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateImportAppleDeviceIdentityListRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateImportAppleDeviceIdentityListRequestBodyFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewImportAppleDeviceIdentityListRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ImportAppleDeviceIdentityListRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ImportAppleDeviceIdentityListRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["importedAppleDeviceIdentities"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateImportedAppleDeviceIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ImportedAppleDeviceIdentityable, len(val))
            for i, v := range val {
                res[i] = v.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ImportedAppleDeviceIdentityable)
            }
            m.SetImportedAppleDeviceIdentities(res)
        }
        return nil
    }
    res["overwriteImportedDeviceIdentities"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOverwriteImportedDeviceIdentities(val)
        }
        return nil
    }
    return res
}
// GetImportedAppleDeviceIdentities gets the importedAppleDeviceIdentities property value. 
func (m *ImportAppleDeviceIdentityListRequestBody) GetImportedAppleDeviceIdentities()([]i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ImportedAppleDeviceIdentityable) {
    if m == nil {
        return nil
    } else {
        return m.importedAppleDeviceIdentities
    }
}
// GetOverwriteImportedDeviceIdentities gets the overwriteImportedDeviceIdentities property value. 
func (m *ImportAppleDeviceIdentityListRequestBody) GetOverwriteImportedDeviceIdentities()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.overwriteImportedDeviceIdentities
    }
}
func (m *ImportAppleDeviceIdentityListRequestBody) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ImportAppleDeviceIdentityListRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetImportedAppleDeviceIdentities() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetImportedAppleDeviceIdentities()))
        for i, v := range m.GetImportedAppleDeviceIdentities() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("importedAppleDeviceIdentities", cast)
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ImportAppleDeviceIdentityListRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetImportedAppleDeviceIdentities sets the importedAppleDeviceIdentities property value. 
func (m *ImportAppleDeviceIdentityListRequestBody) SetImportedAppleDeviceIdentities(value []i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ImportedAppleDeviceIdentityable)() {
    if m != nil {
        m.importedAppleDeviceIdentities = value
    }
}
// SetOverwriteImportedDeviceIdentities sets the overwriteImportedDeviceIdentities property value. 
func (m *ImportAppleDeviceIdentityListRequestBody) SetOverwriteImportedDeviceIdentities(value *bool)() {
    if m != nil {
        m.overwriteImportedDeviceIdentities = value
    }
}

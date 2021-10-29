package windowsprivacyaccesscontrols

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// 
type WindowsPrivacyAccessControlsRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    windowsPrivacyAccessControls []i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WindowsPrivacyDataAccessControlItem;
}
// Instantiates a new windowsPrivacyAccessControlsRequestBody and sets the default values.
func NewWindowsPrivacyAccessControlsRequestBody()(*WindowsPrivacyAccessControlsRequestBody) {
    m := &WindowsPrivacyAccessControlsRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WindowsPrivacyAccessControlsRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the windowsPrivacyAccessControls property value. 
func (m *WindowsPrivacyAccessControlsRequestBody) GetWindowsPrivacyAccessControls()([]i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WindowsPrivacyDataAccessControlItem) {
    if m == nil {
        return nil
    } else {
        return m.windowsPrivacyAccessControls
    }
}
// The deserialization information for the current model
func (m *WindowsPrivacyAccessControlsRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["windowsPrivacyAccessControls"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewWindowsPrivacyDataAccessControlItem() })
        if err != nil {
            return err
        }
        res := make([]i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WindowsPrivacyDataAccessControlItem, len(val))
        for i, v := range val {
            res[i] = *(v.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WindowsPrivacyDataAccessControlItem))
        }
        m.SetWindowsPrivacyAccessControls(res)
        return nil
    }
    return res
}
func (m *WindowsPrivacyAccessControlsRequestBody) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *WindowsPrivacyAccessControlsRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetWindowsPrivacyAccessControls()))
        for i, v := range m.GetWindowsPrivacyAccessControls() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("windowsPrivacyAccessControls", cast)
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
func (m *WindowsPrivacyAccessControlsRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the windowsPrivacyAccessControls property value. 
// Parameters:
//  - value : Value to set for the windowsPrivacyAccessControls property.
func (m *WindowsPrivacyAccessControlsRequestBody) SetWindowsPrivacyAccessControls(value []i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WindowsPrivacyDataAccessControlItem)() {
    m.windowsPrivacyAccessControls = value
}

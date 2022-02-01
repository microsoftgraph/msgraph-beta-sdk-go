package windowsprivacyaccesscontrols

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// WindowsPrivacyAccessControlsRequestBody 
type WindowsPrivacyAccessControlsRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    windowsPrivacyAccessControls []i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WindowsPrivacyDataAccessControlItem;
}
// NewWindowsPrivacyAccessControlsRequestBody instantiates a new windowsPrivacyAccessControlsRequestBody and sets the default values.
func NewWindowsPrivacyAccessControlsRequestBody()(*WindowsPrivacyAccessControlsRequestBody) {
    m := &WindowsPrivacyAccessControlsRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WindowsPrivacyAccessControlsRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetWindowsPrivacyAccessControls gets the windowsPrivacyAccessControls property value. 
func (m *WindowsPrivacyAccessControlsRequestBody) GetWindowsPrivacyAccessControls()([]i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WindowsPrivacyDataAccessControlItem) {
    if m == nil {
        return nil
    } else {
        return m.windowsPrivacyAccessControls
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsPrivacyAccessControlsRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["windowsPrivacyAccessControls"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewWindowsPrivacyDataAccessControlItem() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WindowsPrivacyDataAccessControlItem, len(val))
            for i, v := range val {
                res[i] = *(v.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WindowsPrivacyDataAccessControlItem))
            }
            m.SetWindowsPrivacyAccessControls(res)
        }
        return nil
    }
    return res
}
func (m *WindowsPrivacyAccessControlsRequestBody) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *WindowsPrivacyAccessControlsRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetWindowsPrivacyAccessControls() != nil {
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WindowsPrivacyAccessControlsRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetWindowsPrivacyAccessControls sets the windowsPrivacyAccessControls property value. 
func (m *WindowsPrivacyAccessControlsRequestBody) SetWindowsPrivacyAccessControls(value []i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WindowsPrivacyDataAccessControlItem)() {
    if m != nil {
        m.windowsPrivacyAccessControls = value
    }
}

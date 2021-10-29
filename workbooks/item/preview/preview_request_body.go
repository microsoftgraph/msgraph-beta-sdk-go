package preview

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type PreviewRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    allowEdit *bool;
    // 
    chromeless *bool;
    // 
    page *string;
    // 
    viewer *string;
    // 
    zoom *float64;
}
// Instantiates a new previewRequestBody and sets the default values.
func NewPreviewRequestBody()(*PreviewRequestBody) {
    m := &PreviewRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PreviewRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the allowEdit property value. 
func (m *PreviewRequestBody) GetAllowEdit()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowEdit
    }
}
// Gets the chromeless property value. 
func (m *PreviewRequestBody) GetChromeless()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.chromeless
    }
}
// Gets the page property value. 
func (m *PreviewRequestBody) GetPage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.page
    }
}
// Gets the viewer property value. 
func (m *PreviewRequestBody) GetViewer()(*string) {
    if m == nil {
        return nil
    } else {
        return m.viewer
    }
}
// Gets the zoom property value. 
func (m *PreviewRequestBody) GetZoom()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.zoom
    }
}
// The deserialization information for the current model
func (m *PreviewRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["allowEdit"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAllowEdit(val)
        return nil
    }
    res["chromeless"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetChromeless(val)
        return nil
    }
    res["page"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPage(val)
        return nil
    }
    res["viewer"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetViewer(val)
        return nil
    }
    res["zoom"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetZoom(val)
        return nil
    }
    return res
}
func (m *PreviewRequestBody) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *PreviewRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("allowEdit", m.GetAllowEdit())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("chromeless", m.GetChromeless())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("page", m.GetPage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("viewer", m.GetViewer())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("zoom", m.GetZoom())
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
func (m *PreviewRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the allowEdit property value. 
// Parameters:
//  - value : Value to set for the allowEdit property.
func (m *PreviewRequestBody) SetAllowEdit(value *bool)() {
    m.allowEdit = value
}
// Sets the chromeless property value. 
// Parameters:
//  - value : Value to set for the chromeless property.
func (m *PreviewRequestBody) SetChromeless(value *bool)() {
    m.chromeless = value
}
// Sets the page property value. 
// Parameters:
//  - value : Value to set for the page property.
func (m *PreviewRequestBody) SetPage(value *string)() {
    m.page = value
}
// Sets the viewer property value. 
// Parameters:
//  - value : Value to set for the viewer property.
func (m *PreviewRequestBody) SetViewer(value *string)() {
    m.viewer = value
}
// Sets the zoom property value. 
// Parameters:
//  - value : Value to set for the zoom property.
func (m *PreviewRequestBody) SetZoom(value *float64)() {
    m.zoom = value
}

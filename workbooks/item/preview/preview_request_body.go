package preview

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PreviewRequestBody 
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
// NewPreviewRequestBody instantiates a new previewRequestBody and sets the default values.
func NewPreviewRequestBody()(*PreviewRequestBody) {
    m := &PreviewRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PreviewRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAllowEdit gets the allowEdit property value. 
func (m *PreviewRequestBody) GetAllowEdit()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowEdit
    }
}
// GetChromeless gets the chromeless property value. 
func (m *PreviewRequestBody) GetChromeless()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.chromeless
    }
}
// GetPage gets the page property value. 
func (m *PreviewRequestBody) GetPage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.page
    }
}
// GetViewer gets the viewer property value. 
func (m *PreviewRequestBody) GetViewer()(*string) {
    if m == nil {
        return nil
    } else {
        return m.viewer
    }
}
// GetZoom gets the zoom property value. 
func (m *PreviewRequestBody) GetZoom()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.zoom
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PreviewRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["allowEdit"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowEdit(val)
        }
        return nil
    }
    res["chromeless"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChromeless(val)
        }
        return nil
    }
    res["page"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPage(val)
        }
        return nil
    }
    res["viewer"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetViewer(val)
        }
        return nil
    }
    res["zoom"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetZoom(val)
        }
        return nil
    }
    return res
}
func (m *PreviewRequestBody) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PreviewRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAllowEdit sets the allowEdit property value. 
func (m *PreviewRequestBody) SetAllowEdit(value *bool)() {
    if m != nil {
        m.allowEdit = value
    }
}
// SetChromeless sets the chromeless property value. 
func (m *PreviewRequestBody) SetChromeless(value *bool)() {
    if m != nil {
        m.chromeless = value
    }
}
// SetPage sets the page property value. 
func (m *PreviewRequestBody) SetPage(value *string)() {
    if m != nil {
        m.page = value
    }
}
// SetViewer sets the viewer property value. 
func (m *PreviewRequestBody) SetViewer(value *string)() {
    if m != nil {
        m.viewer = value
    }
}
// SetZoom sets the zoom property value. 
func (m *PreviewRequestBody) SetZoom(value *float64)() {
    if m != nil {
        m.zoom = value
    }
}

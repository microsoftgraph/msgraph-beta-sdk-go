package preview

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PreviewRequestBody provides operations to call the preview method.
type PreviewRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The allowEdit property
    allowEdit *bool
    // The chromeless property
    chromeless *bool
    // The page property
    page *string
    // The viewer property
    viewer *string
    // The zoom property
    zoom *float64
}
// NewPreviewRequestBody instantiates a new previewRequestBody and sets the default values.
func NewPreviewRequestBody()(*PreviewRequestBody) {
    m := &PreviewRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreatePreviewRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePreviewRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPreviewRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PreviewRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAllowEdit gets the allowEdit property value. The allowEdit property
func (m *PreviewRequestBody) GetAllowEdit()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowEdit
    }
}
// GetChromeless gets the chromeless property value. The chromeless property
func (m *PreviewRequestBody) GetChromeless()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.chromeless
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PreviewRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["allowEdit"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowEdit(val)
        }
        return nil
    }
    res["chromeless"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChromeless(val)
        }
        return nil
    }
    res["page"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPage(val)
        }
        return nil
    }
    res["viewer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetViewer(val)
        }
        return nil
    }
    res["zoom"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
// GetPage gets the page property value. The page property
func (m *PreviewRequestBody) GetPage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.page
    }
}
// GetViewer gets the viewer property value. The viewer property
func (m *PreviewRequestBody) GetViewer()(*string) {
    if m == nil {
        return nil
    } else {
        return m.viewer
    }
}
// GetZoom gets the zoom property value. The zoom property
func (m *PreviewRequestBody) GetZoom()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.zoom
    }
}
// Serialize serializes information the current object
func (m *PreviewRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
// SetAllowEdit sets the allowEdit property value. The allowEdit property
func (m *PreviewRequestBody) SetAllowEdit(value *bool)() {
    if m != nil {
        m.allowEdit = value
    }
}
// SetChromeless sets the chromeless property value. The chromeless property
func (m *PreviewRequestBody) SetChromeless(value *bool)() {
    if m != nil {
        m.chromeless = value
    }
}
// SetPage sets the page property value. The page property
func (m *PreviewRequestBody) SetPage(value *string)() {
    if m != nil {
        m.page = value
    }
}
// SetViewer sets the viewer property value. The viewer property
func (m *PreviewRequestBody) SetViewer(value *string)() {
    if m != nil {
        m.viewer = value
    }
}
// SetZoom sets the zoom property value. The zoom property
func (m *PreviewRequestBody) SetZoom(value *float64)() {
    if m != nil {
        m.zoom = value
    }
}

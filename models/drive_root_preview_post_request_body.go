package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DriveRootPreviewPostRequestBody provides operations to call the preview method.
type DriveRootPreviewPostRequestBody struct {
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
// NewDriveRootPreviewPostRequestBody instantiates a new DriveRootPreviewPostRequestBody and sets the default values.
func NewDriveRootPreviewPostRequestBody()(*DriveRootPreviewPostRequestBody) {
    m := &DriveRootPreviewPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDriveRootPreviewPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDriveRootPreviewPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDriveRootPreviewPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DriveRootPreviewPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAllowEdit gets the allowEdit property value. The allowEdit property
func (m *DriveRootPreviewPostRequestBody) GetAllowEdit()(*bool) {
    return m.allowEdit
}
// GetChromeless gets the chromeless property value. The chromeless property
func (m *DriveRootPreviewPostRequestBody) GetChromeless()(*bool) {
    return m.chromeless
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DriveRootPreviewPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *DriveRootPreviewPostRequestBody) GetPage()(*string) {
    return m.page
}
// GetViewer gets the viewer property value. The viewer property
func (m *DriveRootPreviewPostRequestBody) GetViewer()(*string) {
    return m.viewer
}
// GetZoom gets the zoom property value. The zoom property
func (m *DriveRootPreviewPostRequestBody) GetZoom()(*float64) {
    return m.zoom
}
// Serialize serializes information the current object
func (m *DriveRootPreviewPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *DriveRootPreviewPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAllowEdit sets the allowEdit property value. The allowEdit property
func (m *DriveRootPreviewPostRequestBody) SetAllowEdit(value *bool)() {
    m.allowEdit = value
}
// SetChromeless sets the chromeless property value. The chromeless property
func (m *DriveRootPreviewPostRequestBody) SetChromeless(value *bool)() {
    m.chromeless = value
}
// SetPage sets the page property value. The page property
func (m *DriveRootPreviewPostRequestBody) SetPage(value *string)() {
    m.page = value
}
// SetViewer sets the viewer property value. The viewer property
func (m *DriveRootPreviewPostRequestBody) SetViewer(value *string)() {
    m.viewer = value
}
// SetZoom sets the zoom property value. The zoom property
func (m *DriveRootPreviewPostRequestBody) SetZoom(value *float64)() {
    m.zoom = value
}

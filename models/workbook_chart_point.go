package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WorkbookChartPoint provides operations to manage the collection of accessReview entities.
type WorkbookChartPoint struct {
    Entity
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Encapsulates the format properties chart point. Read-only.
    format WorkbookChartPointFormatable
    // Returns the value of a chart point. Read-only.
    value Jsonable
}
// NewWorkbookChartPoint instantiates a new workbookChartPoint and sets the default values.
func NewWorkbookChartPoint()(*WorkbookChartPoint) {
    m := &WorkbookChartPoint{
        Entity: *NewEntity(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateWorkbookChartPointFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWorkbookChartPointFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWorkbookChartPoint(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WorkbookChartPoint) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkbookChartPoint) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["format"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkbookChartPointFormatFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFormat(val.(WorkbookChartPointFormatable))
        }
        return nil
    }
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val.(Jsonable))
        }
        return nil
    }
    return res
}
// GetFormat gets the format property value. Encapsulates the format properties chart point. Read-only.
func (m *WorkbookChartPoint) GetFormat()(WorkbookChartPointFormatable) {
    if m == nil {
        return nil
    } else {
        return m.format
    }
}
// GetValue gets the value property value. Returns the value of a chart point. Read-only.
func (m *WorkbookChartPoint) GetValue()(Jsonable) {
    if m == nil {
        return nil
    } else {
        return m.value
    }
}
// Serialize serializes information the current object
func (m *WorkbookChartPoint) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("format", m.GetFormat())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("value", m.GetValue())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WorkbookChartPoint) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetFormat sets the format property value. Encapsulates the format properties chart point. Read-only.
func (m *WorkbookChartPoint) SetFormat(value WorkbookChartPointFormatable)() {
    if m != nil {
        m.format = value
    }
}
// SetValue sets the value property value. Returns the value of a chart point. Read-only.
func (m *WorkbookChartPoint) SetValue(value Jsonable)() {
    if m != nil {
        m.value = value
    }
}

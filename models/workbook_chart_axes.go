package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WorkbookChartAxes provides operations to manage the collection of accessReview entities.
type WorkbookChartAxes struct {
    Entity
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Represents the category axis in a chart. Read-only.
    categoryAxis WorkbookChartAxisable
    // Represents the series axis of a 3-dimensional chart. Read-only.
    seriesAxis WorkbookChartAxisable
    // Represents the value axis in an axis. Read-only.
    valueAxis WorkbookChartAxisable
}
// NewWorkbookChartAxes instantiates a new workbookChartAxes and sets the default values.
func NewWorkbookChartAxes()(*WorkbookChartAxes) {
    m := &WorkbookChartAxes{
        Entity: *NewEntity(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateWorkbookChartAxesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWorkbookChartAxesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWorkbookChartAxes(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WorkbookChartAxes) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCategoryAxis gets the categoryAxis property value. Represents the category axis in a chart. Read-only.
func (m *WorkbookChartAxes) GetCategoryAxis()(WorkbookChartAxisable) {
    if m == nil {
        return nil
    } else {
        return m.categoryAxis
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkbookChartAxes) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["categoryAxis"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkbookChartAxisFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategoryAxis(val.(WorkbookChartAxisable))
        }
        return nil
    }
    res["seriesAxis"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkbookChartAxisFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSeriesAxis(val.(WorkbookChartAxisable))
        }
        return nil
    }
    res["valueAxis"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkbookChartAxisFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValueAxis(val.(WorkbookChartAxisable))
        }
        return nil
    }
    return res
}
// GetSeriesAxis gets the seriesAxis property value. Represents the series axis of a 3-dimensional chart. Read-only.
func (m *WorkbookChartAxes) GetSeriesAxis()(WorkbookChartAxisable) {
    if m == nil {
        return nil
    } else {
        return m.seriesAxis
    }
}
// GetValueAxis gets the valueAxis property value. Represents the value axis in an axis. Read-only.
func (m *WorkbookChartAxes) GetValueAxis()(WorkbookChartAxisable) {
    if m == nil {
        return nil
    } else {
        return m.valueAxis
    }
}
// Serialize serializes information the current object
func (m *WorkbookChartAxes) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("categoryAxis", m.GetCategoryAxis())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("seriesAxis", m.GetSeriesAxis())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("valueAxis", m.GetValueAxis())
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
func (m *WorkbookChartAxes) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetCategoryAxis sets the categoryAxis property value. Represents the category axis in a chart. Read-only.
func (m *WorkbookChartAxes) SetCategoryAxis(value WorkbookChartAxisable)() {
    if m != nil {
        m.categoryAxis = value
    }
}
// SetSeriesAxis sets the seriesAxis property value. Represents the series axis of a 3-dimensional chart. Read-only.
func (m *WorkbookChartAxes) SetSeriesAxis(value WorkbookChartAxisable)() {
    if m != nil {
        m.seriesAxis = value
    }
}
// SetValueAxis sets the valueAxis property value. Represents the value axis in an axis. Read-only.
func (m *WorkbookChartAxes) SetValueAxis(value WorkbookChartAxisable)() {
    if m != nil {
        m.valueAxis = value
    }
}

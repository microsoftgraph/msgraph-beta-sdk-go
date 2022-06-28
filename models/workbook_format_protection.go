package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WorkbookFormatProtection 
type WorkbookFormatProtection struct {
    Entity
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Indicates if Excel hides the formula for the cells in the range. A null value indicates that the entire range doesn't have uniform formula hidden setting.
    formulaHidden *bool
    // Indicates if Excel locks the cells in the object. A null value indicates that the entire range doesn't have uniform lock setting.
    locked *bool
}
// NewWorkbookFormatProtection instantiates a new WorkbookFormatProtection and sets the default values.
func NewWorkbookFormatProtection()(*WorkbookFormatProtection) {
    m := &WorkbookFormatProtection{
        Entity: *NewEntity(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateWorkbookFormatProtectionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWorkbookFormatProtectionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWorkbookFormatProtection(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WorkbookFormatProtection) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkbookFormatProtection) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["formulaHidden"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFormulaHidden(val)
        }
        return nil
    }
    res["locked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocked(val)
        }
        return nil
    }
    return res
}
// GetFormulaHidden gets the formulaHidden property value. Indicates if Excel hides the formula for the cells in the range. A null value indicates that the entire range doesn't have uniform formula hidden setting.
func (m *WorkbookFormatProtection) GetFormulaHidden()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.formulaHidden
    }
}
// GetLocked gets the locked property value. Indicates if Excel locks the cells in the object. A null value indicates that the entire range doesn't have uniform lock setting.
func (m *WorkbookFormatProtection) GetLocked()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.locked
    }
}
// Serialize serializes information the current object
func (m *WorkbookFormatProtection) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("formulaHidden", m.GetFormulaHidden())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("locked", m.GetLocked())
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
func (m *WorkbookFormatProtection) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetFormulaHidden sets the formulaHidden property value. Indicates if Excel hides the formula for the cells in the range. A null value indicates that the entire range doesn't have uniform formula hidden setting.
func (m *WorkbookFormatProtection) SetFormulaHidden(value *bool)() {
    if m != nil {
        m.formulaHidden = value
    }
}
// SetLocked sets the locked property value. Indicates if Excel locks the cells in the object. A null value indicates that the entire range doesn't have uniform lock setting.
func (m *WorkbookFormatProtection) SetLocked(value *bool)() {
    if m != nil {
        m.locked = value
    }
}

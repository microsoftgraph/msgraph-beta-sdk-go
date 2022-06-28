package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Win32LobAppRegistryDetection 
type Win32LobAppRegistryDetection struct {
    Win32LobAppDetection
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // A value indicating whether this registry path is for checking 32-bit app on 64-bit system
    check32BitOn64System *bool
    // The registry data detection type. Possible values are: notConfigured, exists, doesNotExist, string, integer, version.
    detectionType *Win32LobAppRegistryDetectionType
    // The registry detection value
    detectionValue *string
    // The registry key path to detect Win32 Line of Business (LoB) app
    keyPath *string
    // The operator for registry data detection. Possible values are: notConfigured, equal, notEqual, greaterThan, greaterThanOrEqual, lessThan, lessThanOrEqual.
    operator *Win32LobAppDetectionOperator
    // The registry value name
    valueName *string
}
// NewWin32LobAppRegistryDetection instantiates a new Win32LobAppRegistryDetection and sets the default values.
func NewWin32LobAppRegistryDetection()(*Win32LobAppRegistryDetection) {
    m := &Win32LobAppRegistryDetection{
        Win32LobAppDetection: *NewWin32LobAppDetection(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateWin32LobAppRegistryDetectionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWin32LobAppRegistryDetectionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWin32LobAppRegistryDetection(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Win32LobAppRegistryDetection) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCheck32BitOn64System gets the check32BitOn64System property value. A value indicating whether this registry path is for checking 32-bit app on 64-bit system
func (m *Win32LobAppRegistryDetection) GetCheck32BitOn64System()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.check32BitOn64System
    }
}
// GetDetectionType gets the detectionType property value. The registry data detection type. Possible values are: notConfigured, exists, doesNotExist, string, integer, version.
func (m *Win32LobAppRegistryDetection) GetDetectionType()(*Win32LobAppRegistryDetectionType) {
    if m == nil {
        return nil
    } else {
        return m.detectionType
    }
}
// GetDetectionValue gets the detectionValue property value. The registry detection value
func (m *Win32LobAppRegistryDetection) GetDetectionValue()(*string) {
    if m == nil {
        return nil
    } else {
        return m.detectionValue
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Win32LobAppRegistryDetection) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Win32LobAppDetection.GetFieldDeserializers()
    res["check32BitOn64System"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCheck32BitOn64System(val)
        }
        return nil
    }
    res["detectionType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWin32LobAppRegistryDetectionType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetectionType(val.(*Win32LobAppRegistryDetectionType))
        }
        return nil
    }
    res["detectionValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetectionValue(val)
        }
        return nil
    }
    res["keyPath"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeyPath(val)
        }
        return nil
    }
    res["operator"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWin32LobAppDetectionOperator)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOperator(val.(*Win32LobAppDetectionOperator))
        }
        return nil
    }
    res["valueName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValueName(val)
        }
        return nil
    }
    return res
}
// GetKeyPath gets the keyPath property value. The registry key path to detect Win32 Line of Business (LoB) app
func (m *Win32LobAppRegistryDetection) GetKeyPath()(*string) {
    if m == nil {
        return nil
    } else {
        return m.keyPath
    }
}
// GetOperator gets the operator property value. The operator for registry data detection. Possible values are: notConfigured, equal, notEqual, greaterThan, greaterThanOrEqual, lessThan, lessThanOrEqual.
func (m *Win32LobAppRegistryDetection) GetOperator()(*Win32LobAppDetectionOperator) {
    if m == nil {
        return nil
    } else {
        return m.operator
    }
}
// GetValueName gets the valueName property value. The registry value name
func (m *Win32LobAppRegistryDetection) GetValueName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.valueName
    }
}
// Serialize serializes information the current object
func (m *Win32LobAppRegistryDetection) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Win32LobAppDetection.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("check32BitOn64System", m.GetCheck32BitOn64System())
        if err != nil {
            return err
        }
    }
    if m.GetDetectionType() != nil {
        cast := (*m.GetDetectionType()).String()
        err = writer.WriteStringValue("detectionType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("detectionValue", m.GetDetectionValue())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("keyPath", m.GetKeyPath())
        if err != nil {
            return err
        }
    }
    if m.GetOperator() != nil {
        cast := (*m.GetOperator()).String()
        err = writer.WriteStringValue("operator", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("valueName", m.GetValueName())
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
func (m *Win32LobAppRegistryDetection) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetCheck32BitOn64System sets the check32BitOn64System property value. A value indicating whether this registry path is for checking 32-bit app on 64-bit system
func (m *Win32LobAppRegistryDetection) SetCheck32BitOn64System(value *bool)() {
    if m != nil {
        m.check32BitOn64System = value
    }
}
// SetDetectionType sets the detectionType property value. The registry data detection type. Possible values are: notConfigured, exists, doesNotExist, string, integer, version.
func (m *Win32LobAppRegistryDetection) SetDetectionType(value *Win32LobAppRegistryDetectionType)() {
    if m != nil {
        m.detectionType = value
    }
}
// SetDetectionValue sets the detectionValue property value. The registry detection value
func (m *Win32LobAppRegistryDetection) SetDetectionValue(value *string)() {
    if m != nil {
        m.detectionValue = value
    }
}
// SetKeyPath sets the keyPath property value. The registry key path to detect Win32 Line of Business (LoB) app
func (m *Win32LobAppRegistryDetection) SetKeyPath(value *string)() {
    if m != nil {
        m.keyPath = value
    }
}
// SetOperator sets the operator property value. The operator for registry data detection. Possible values are: notConfigured, equal, notEqual, greaterThan, greaterThanOrEqual, lessThan, lessThanOrEqual.
func (m *Win32LobAppRegistryDetection) SetOperator(value *Win32LobAppDetectionOperator)() {
    if m != nil {
        m.operator = value
    }
}
// SetValueName sets the valueName property value. The registry value name
func (m *Win32LobAppRegistryDetection) SetValueName(value *string)() {
    if m != nil {
        m.valueName = value
    }
}

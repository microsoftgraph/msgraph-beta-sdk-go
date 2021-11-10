package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type DeviceProtectionOverview struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Clean device count.
    cleanDeviceCount *int32;
    // Critical failures device count.
    criticalFailuresDeviceCount *int32;
    // Device with inactive threat agent count
    inactiveThreatAgentDeviceCount *int32;
    // Pending full scan device count.
    pendingFullScanDeviceCount *int32;
    // Pending manual steps device count.
    pendingManualStepsDeviceCount *int32;
    // Pending offline scan device count.
    pendingOfflineScanDeviceCount *int32;
    // Pending quick scan device count. Valid values -2147483648 to 2147483647
    pendingQuickScanDeviceCount *int32;
    // Pending restart device count.
    pendingRestartDeviceCount *int32;
    // Device with old signature count.
    pendingSignatureUpdateDeviceCount *int32;
    // Total device count.
    totalReportedDeviceCount *int32;
    // Device with threat agent state as unknown count.
    unknownStateThreatAgentDeviceCount *int32;
}
// Instantiates a new deviceProtectionOverview and sets the default values.
func NewDeviceProtectionOverview()(*DeviceProtectionOverview) {
    m := &DeviceProtectionOverview{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceProtectionOverview) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the cleanDeviceCount property value. Clean device count.
func (m *DeviceProtectionOverview) GetCleanDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.cleanDeviceCount
    }
}
// Gets the criticalFailuresDeviceCount property value. Critical failures device count.
func (m *DeviceProtectionOverview) GetCriticalFailuresDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.criticalFailuresDeviceCount
    }
}
// Gets the inactiveThreatAgentDeviceCount property value. Device with inactive threat agent count
func (m *DeviceProtectionOverview) GetInactiveThreatAgentDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.inactiveThreatAgentDeviceCount
    }
}
// Gets the pendingFullScanDeviceCount property value. Pending full scan device count.
func (m *DeviceProtectionOverview) GetPendingFullScanDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.pendingFullScanDeviceCount
    }
}
// Gets the pendingManualStepsDeviceCount property value. Pending manual steps device count.
func (m *DeviceProtectionOverview) GetPendingManualStepsDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.pendingManualStepsDeviceCount
    }
}
// Gets the pendingOfflineScanDeviceCount property value. Pending offline scan device count.
func (m *DeviceProtectionOverview) GetPendingOfflineScanDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.pendingOfflineScanDeviceCount
    }
}
// Gets the pendingQuickScanDeviceCount property value. Pending quick scan device count. Valid values -2147483648 to 2147483647
func (m *DeviceProtectionOverview) GetPendingQuickScanDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.pendingQuickScanDeviceCount
    }
}
// Gets the pendingRestartDeviceCount property value. Pending restart device count.
func (m *DeviceProtectionOverview) GetPendingRestartDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.pendingRestartDeviceCount
    }
}
// Gets the pendingSignatureUpdateDeviceCount property value. Device with old signature count.
func (m *DeviceProtectionOverview) GetPendingSignatureUpdateDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.pendingSignatureUpdateDeviceCount
    }
}
// Gets the totalReportedDeviceCount property value. Total device count.
func (m *DeviceProtectionOverview) GetTotalReportedDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.totalReportedDeviceCount
    }
}
// Gets the unknownStateThreatAgentDeviceCount property value. Device with threat agent state as unknown count.
func (m *DeviceProtectionOverview) GetUnknownStateThreatAgentDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.unknownStateThreatAgentDeviceCount
    }
}
// The deserialization information for the current model
func (m *DeviceProtectionOverview) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["cleanDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCleanDeviceCount(val)
        }
        return nil
    }
    res["criticalFailuresDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCriticalFailuresDeviceCount(val)
        }
        return nil
    }
    res["inactiveThreatAgentDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInactiveThreatAgentDeviceCount(val)
        }
        return nil
    }
    res["pendingFullScanDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPendingFullScanDeviceCount(val)
        }
        return nil
    }
    res["pendingManualStepsDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPendingManualStepsDeviceCount(val)
        }
        return nil
    }
    res["pendingOfflineScanDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPendingOfflineScanDeviceCount(val)
        }
        return nil
    }
    res["pendingQuickScanDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPendingQuickScanDeviceCount(val)
        }
        return nil
    }
    res["pendingRestartDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPendingRestartDeviceCount(val)
        }
        return nil
    }
    res["pendingSignatureUpdateDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPendingSignatureUpdateDeviceCount(val)
        }
        return nil
    }
    res["totalReportedDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalReportedDeviceCount(val)
        }
        return nil
    }
    res["unknownStateThreatAgentDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUnknownStateThreatAgentDeviceCount(val)
        }
        return nil
    }
    return res
}
func (m *DeviceProtectionOverview) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *DeviceProtectionOverview) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("cleanDeviceCount", m.GetCleanDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("criticalFailuresDeviceCount", m.GetCriticalFailuresDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("inactiveThreatAgentDeviceCount", m.GetInactiveThreatAgentDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("pendingFullScanDeviceCount", m.GetPendingFullScanDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("pendingManualStepsDeviceCount", m.GetPendingManualStepsDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("pendingOfflineScanDeviceCount", m.GetPendingOfflineScanDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("pendingQuickScanDeviceCount", m.GetPendingQuickScanDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("pendingRestartDeviceCount", m.GetPendingRestartDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("pendingSignatureUpdateDeviceCount", m.GetPendingSignatureUpdateDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("totalReportedDeviceCount", m.GetTotalReportedDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("unknownStateThreatAgentDeviceCount", m.GetUnknownStateThreatAgentDeviceCount())
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
func (m *DeviceProtectionOverview) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the cleanDeviceCount property value. Clean device count.
// Parameters:
//  - value : Value to set for the cleanDeviceCount property.
func (m *DeviceProtectionOverview) SetCleanDeviceCount(value *int32)() {
    m.cleanDeviceCount = value
}
// Sets the criticalFailuresDeviceCount property value. Critical failures device count.
// Parameters:
//  - value : Value to set for the criticalFailuresDeviceCount property.
func (m *DeviceProtectionOverview) SetCriticalFailuresDeviceCount(value *int32)() {
    m.criticalFailuresDeviceCount = value
}
// Sets the inactiveThreatAgentDeviceCount property value. Device with inactive threat agent count
// Parameters:
//  - value : Value to set for the inactiveThreatAgentDeviceCount property.
func (m *DeviceProtectionOverview) SetInactiveThreatAgentDeviceCount(value *int32)() {
    m.inactiveThreatAgentDeviceCount = value
}
// Sets the pendingFullScanDeviceCount property value. Pending full scan device count.
// Parameters:
//  - value : Value to set for the pendingFullScanDeviceCount property.
func (m *DeviceProtectionOverview) SetPendingFullScanDeviceCount(value *int32)() {
    m.pendingFullScanDeviceCount = value
}
// Sets the pendingManualStepsDeviceCount property value. Pending manual steps device count.
// Parameters:
//  - value : Value to set for the pendingManualStepsDeviceCount property.
func (m *DeviceProtectionOverview) SetPendingManualStepsDeviceCount(value *int32)() {
    m.pendingManualStepsDeviceCount = value
}
// Sets the pendingOfflineScanDeviceCount property value. Pending offline scan device count.
// Parameters:
//  - value : Value to set for the pendingOfflineScanDeviceCount property.
func (m *DeviceProtectionOverview) SetPendingOfflineScanDeviceCount(value *int32)() {
    m.pendingOfflineScanDeviceCount = value
}
// Sets the pendingQuickScanDeviceCount property value. Pending quick scan device count. Valid values -2147483648 to 2147483647
// Parameters:
//  - value : Value to set for the pendingQuickScanDeviceCount property.
func (m *DeviceProtectionOverview) SetPendingQuickScanDeviceCount(value *int32)() {
    m.pendingQuickScanDeviceCount = value
}
// Sets the pendingRestartDeviceCount property value. Pending restart device count.
// Parameters:
//  - value : Value to set for the pendingRestartDeviceCount property.
func (m *DeviceProtectionOverview) SetPendingRestartDeviceCount(value *int32)() {
    m.pendingRestartDeviceCount = value
}
// Sets the pendingSignatureUpdateDeviceCount property value. Device with old signature count.
// Parameters:
//  - value : Value to set for the pendingSignatureUpdateDeviceCount property.
func (m *DeviceProtectionOverview) SetPendingSignatureUpdateDeviceCount(value *int32)() {
    m.pendingSignatureUpdateDeviceCount = value
}
// Sets the totalReportedDeviceCount property value. Total device count.
// Parameters:
//  - value : Value to set for the totalReportedDeviceCount property.
func (m *DeviceProtectionOverview) SetTotalReportedDeviceCount(value *int32)() {
    m.totalReportedDeviceCount = value
}
// Sets the unknownStateThreatAgentDeviceCount property value. Device with threat agent state as unknown count.
// Parameters:
//  - value : Value to set for the unknownStateThreatAgentDeviceCount property.
func (m *DeviceProtectionOverview) SetUnknownStateThreatAgentDeviceCount(value *int32)() {
    m.unknownStateThreatAgentDeviceCount = value
}

package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceProtectionOverview 
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
// NewDeviceProtectionOverview instantiates a new deviceProtectionOverview and sets the default values.
func NewDeviceProtectionOverview()(*DeviceProtectionOverview) {
    m := &DeviceProtectionOverview{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceProtectionOverview) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCleanDeviceCount gets the cleanDeviceCount property value. Clean device count.
func (m *DeviceProtectionOverview) GetCleanDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.cleanDeviceCount
    }
}
// GetCriticalFailuresDeviceCount gets the criticalFailuresDeviceCount property value. Critical failures device count.
func (m *DeviceProtectionOverview) GetCriticalFailuresDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.criticalFailuresDeviceCount
    }
}
// GetInactiveThreatAgentDeviceCount gets the inactiveThreatAgentDeviceCount property value. Device with inactive threat agent count
func (m *DeviceProtectionOverview) GetInactiveThreatAgentDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.inactiveThreatAgentDeviceCount
    }
}
// GetPendingFullScanDeviceCount gets the pendingFullScanDeviceCount property value. Pending full scan device count.
func (m *DeviceProtectionOverview) GetPendingFullScanDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.pendingFullScanDeviceCount
    }
}
// GetPendingManualStepsDeviceCount gets the pendingManualStepsDeviceCount property value. Pending manual steps device count.
func (m *DeviceProtectionOverview) GetPendingManualStepsDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.pendingManualStepsDeviceCount
    }
}
// GetPendingOfflineScanDeviceCount gets the pendingOfflineScanDeviceCount property value. Pending offline scan device count.
func (m *DeviceProtectionOverview) GetPendingOfflineScanDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.pendingOfflineScanDeviceCount
    }
}
// GetPendingQuickScanDeviceCount gets the pendingQuickScanDeviceCount property value. Pending quick scan device count. Valid values -2147483648 to 2147483647
func (m *DeviceProtectionOverview) GetPendingQuickScanDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.pendingQuickScanDeviceCount
    }
}
// GetPendingRestartDeviceCount gets the pendingRestartDeviceCount property value. Pending restart device count.
func (m *DeviceProtectionOverview) GetPendingRestartDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.pendingRestartDeviceCount
    }
}
// GetPendingSignatureUpdateDeviceCount gets the pendingSignatureUpdateDeviceCount property value. Device with old signature count.
func (m *DeviceProtectionOverview) GetPendingSignatureUpdateDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.pendingSignatureUpdateDeviceCount
    }
}
// GetTotalReportedDeviceCount gets the totalReportedDeviceCount property value. Total device count.
func (m *DeviceProtectionOverview) GetTotalReportedDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.totalReportedDeviceCount
    }
}
// GetUnknownStateThreatAgentDeviceCount gets the unknownStateThreatAgentDeviceCount property value. Device with threat agent state as unknown count.
func (m *DeviceProtectionOverview) GetUnknownStateThreatAgentDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.unknownStateThreatAgentDeviceCount
    }
}
// GetFieldDeserializers the deserialization information for the current model
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
// Serialize serializes information the current object
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceProtectionOverview) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetCleanDeviceCount sets the cleanDeviceCount property value. Clean device count.
func (m *DeviceProtectionOverview) SetCleanDeviceCount(value *int32)() {
    if m != nil {
        m.cleanDeviceCount = value
    }
}
// SetCriticalFailuresDeviceCount sets the criticalFailuresDeviceCount property value. Critical failures device count.
func (m *DeviceProtectionOverview) SetCriticalFailuresDeviceCount(value *int32)() {
    if m != nil {
        m.criticalFailuresDeviceCount = value
    }
}
// SetInactiveThreatAgentDeviceCount sets the inactiveThreatAgentDeviceCount property value. Device with inactive threat agent count
func (m *DeviceProtectionOverview) SetInactiveThreatAgentDeviceCount(value *int32)() {
    if m != nil {
        m.inactiveThreatAgentDeviceCount = value
    }
}
// SetPendingFullScanDeviceCount sets the pendingFullScanDeviceCount property value. Pending full scan device count.
func (m *DeviceProtectionOverview) SetPendingFullScanDeviceCount(value *int32)() {
    if m != nil {
        m.pendingFullScanDeviceCount = value
    }
}
// SetPendingManualStepsDeviceCount sets the pendingManualStepsDeviceCount property value. Pending manual steps device count.
func (m *DeviceProtectionOverview) SetPendingManualStepsDeviceCount(value *int32)() {
    if m != nil {
        m.pendingManualStepsDeviceCount = value
    }
}
// SetPendingOfflineScanDeviceCount sets the pendingOfflineScanDeviceCount property value. Pending offline scan device count.
func (m *DeviceProtectionOverview) SetPendingOfflineScanDeviceCount(value *int32)() {
    if m != nil {
        m.pendingOfflineScanDeviceCount = value
    }
}
// SetPendingQuickScanDeviceCount sets the pendingQuickScanDeviceCount property value. Pending quick scan device count. Valid values -2147483648 to 2147483647
func (m *DeviceProtectionOverview) SetPendingQuickScanDeviceCount(value *int32)() {
    if m != nil {
        m.pendingQuickScanDeviceCount = value
    }
}
// SetPendingRestartDeviceCount sets the pendingRestartDeviceCount property value. Pending restart device count.
func (m *DeviceProtectionOverview) SetPendingRestartDeviceCount(value *int32)() {
    if m != nil {
        m.pendingRestartDeviceCount = value
    }
}
// SetPendingSignatureUpdateDeviceCount sets the pendingSignatureUpdateDeviceCount property value. Device with old signature count.
func (m *DeviceProtectionOverview) SetPendingSignatureUpdateDeviceCount(value *int32)() {
    if m != nil {
        m.pendingSignatureUpdateDeviceCount = value
    }
}
// SetTotalReportedDeviceCount sets the totalReportedDeviceCount property value. Total device count.
func (m *DeviceProtectionOverview) SetTotalReportedDeviceCount(value *int32)() {
    if m != nil {
        m.totalReportedDeviceCount = value
    }
}
// SetUnknownStateThreatAgentDeviceCount sets the unknownStateThreatAgentDeviceCount property value. Device with threat agent state as unknown count.
func (m *DeviceProtectionOverview) SetUnknownStateThreatAgentDeviceCount(value *int32)() {
    if m != nil {
        m.unknownStateThreatAgentDeviceCount = value
    }
}

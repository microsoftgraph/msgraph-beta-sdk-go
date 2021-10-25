package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type DeviceProtectionOverview struct {
    additionalData map[string]interface{};
    cleanDeviceCount *int32;
    criticalFailuresDeviceCount *int32;
    inactiveThreatAgentDeviceCount *int32;
    pendingFullScanDeviceCount *int32;
    pendingManualStepsDeviceCount *int32;
    pendingOfflineScanDeviceCount *int32;
    pendingQuickScanDeviceCount *int32;
    pendingRestartDeviceCount *int32;
    pendingSignatureUpdateDeviceCount *int32;
    totalReportedDeviceCount *int32;
    unknownStateThreatAgentDeviceCount *int32;
}
func NewDeviceProtectionOverview()(*DeviceProtectionOverview) {
    m := &DeviceProtectionOverview{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *DeviceProtectionOverview) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *DeviceProtectionOverview) GetCleanDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.cleanDeviceCount
    }
}
func (m *DeviceProtectionOverview) GetCriticalFailuresDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.criticalFailuresDeviceCount
    }
}
func (m *DeviceProtectionOverview) GetInactiveThreatAgentDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.inactiveThreatAgentDeviceCount
    }
}
func (m *DeviceProtectionOverview) GetPendingFullScanDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.pendingFullScanDeviceCount
    }
}
func (m *DeviceProtectionOverview) GetPendingManualStepsDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.pendingManualStepsDeviceCount
    }
}
func (m *DeviceProtectionOverview) GetPendingOfflineScanDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.pendingOfflineScanDeviceCount
    }
}
func (m *DeviceProtectionOverview) GetPendingQuickScanDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.pendingQuickScanDeviceCount
    }
}
func (m *DeviceProtectionOverview) GetPendingRestartDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.pendingRestartDeviceCount
    }
}
func (m *DeviceProtectionOverview) GetPendingSignatureUpdateDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.pendingSignatureUpdateDeviceCount
    }
}
func (m *DeviceProtectionOverview) GetTotalReportedDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.totalReportedDeviceCount
    }
}
func (m *DeviceProtectionOverview) GetUnknownStateThreatAgentDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.unknownStateThreatAgentDeviceCount
    }
}
func (m *DeviceProtectionOverview) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["cleanDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetCleanDeviceCount(val)
        return nil
    }
    res["criticalFailuresDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetCriticalFailuresDeviceCount(val)
        return nil
    }
    res["inactiveThreatAgentDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetInactiveThreatAgentDeviceCount(val)
        return nil
    }
    res["pendingFullScanDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetPendingFullScanDeviceCount(val)
        return nil
    }
    res["pendingManualStepsDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetPendingManualStepsDeviceCount(val)
        return nil
    }
    res["pendingOfflineScanDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetPendingOfflineScanDeviceCount(val)
        return nil
    }
    res["pendingQuickScanDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetPendingQuickScanDeviceCount(val)
        return nil
    }
    res["pendingRestartDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetPendingRestartDeviceCount(val)
        return nil
    }
    res["pendingSignatureUpdateDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetPendingSignatureUpdateDeviceCount(val)
        return nil
    }
    res["totalReportedDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetTotalReportedDeviceCount(val)
        return nil
    }
    res["unknownStateThreatAgentDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetUnknownStateThreatAgentDeviceCount(val)
        return nil
    }
    return res
}
func (m *DeviceProtectionOverview) IsNil()(bool) {
    return m == nil
}
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
func (m *DeviceProtectionOverview) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *DeviceProtectionOverview) SetCleanDeviceCount(value *int32)() {
    m.cleanDeviceCount = value
}
func (m *DeviceProtectionOverview) SetCriticalFailuresDeviceCount(value *int32)() {
    m.criticalFailuresDeviceCount = value
}
func (m *DeviceProtectionOverview) SetInactiveThreatAgentDeviceCount(value *int32)() {
    m.inactiveThreatAgentDeviceCount = value
}
func (m *DeviceProtectionOverview) SetPendingFullScanDeviceCount(value *int32)() {
    m.pendingFullScanDeviceCount = value
}
func (m *DeviceProtectionOverview) SetPendingManualStepsDeviceCount(value *int32)() {
    m.pendingManualStepsDeviceCount = value
}
func (m *DeviceProtectionOverview) SetPendingOfflineScanDeviceCount(value *int32)() {
    m.pendingOfflineScanDeviceCount = value
}
func (m *DeviceProtectionOverview) SetPendingQuickScanDeviceCount(value *int32)() {
    m.pendingQuickScanDeviceCount = value
}
func (m *DeviceProtectionOverview) SetPendingRestartDeviceCount(value *int32)() {
    m.pendingRestartDeviceCount = value
}
func (m *DeviceProtectionOverview) SetPendingSignatureUpdateDeviceCount(value *int32)() {
    m.pendingSignatureUpdateDeviceCount = value
}
func (m *DeviceProtectionOverview) SetTotalReportedDeviceCount(value *int32)() {
    m.totalReportedDeviceCount = value
}
func (m *DeviceProtectionOverview) SetUnknownStateThreatAgentDeviceCount(value *int32)() {
    m.unknownStateThreatAgentDeviceCount = value
}

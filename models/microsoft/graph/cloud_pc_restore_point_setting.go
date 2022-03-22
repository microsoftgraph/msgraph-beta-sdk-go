package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// CloudPcRestorePointSetting 
type CloudPcRestorePointSetting struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The time interval in hours to take snapshots (restore points) of a Cloud PC automatically. Possible values are 4, 6, 12, 16, and 24. The default frequency is 12 hours.
    frequencyInHours *int32;
    // If true, the user has the ability to use snapshots to restore Cloud PCs. If false, non-admin users cannot use snapshots to restore the Cloud PC.
    userRestoreEnabled *bool;
}
// NewCloudPcRestorePointSetting instantiates a new cloudPcRestorePointSetting and sets the default values.
func NewCloudPcRestorePointSetting()(*CloudPcRestorePointSetting) {
    m := &CloudPcRestorePointSetting{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateCloudPcRestorePointSettingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCloudPcRestorePointSettingFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewCloudPcRestorePointSetting(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CloudPcRestorePointSetting) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CloudPcRestorePointSetting) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["frequencyInHours"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFrequencyInHours(val)
        }
        return nil
    }
    res["userRestoreEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserRestoreEnabled(val)
        }
        return nil
    }
    return res
}
// GetFrequencyInHours gets the frequencyInHours property value. The time interval in hours to take snapshots (restore points) of a Cloud PC automatically. Possible values are 4, 6, 12, 16, and 24. The default frequency is 12 hours.
func (m *CloudPcRestorePointSetting) GetFrequencyInHours()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.frequencyInHours
    }
}
// GetUserRestoreEnabled gets the userRestoreEnabled property value. If true, the user has the ability to use snapshots to restore Cloud PCs. If false, non-admin users cannot use snapshots to restore the Cloud PC.
func (m *CloudPcRestorePointSetting) GetUserRestoreEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.userRestoreEnabled
    }
}
// Serialize serializes information the current object
func (m *CloudPcRestorePointSetting) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("frequencyInHours", m.GetFrequencyInHours())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("userRestoreEnabled", m.GetUserRestoreEnabled())
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
func (m *CloudPcRestorePointSetting) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetFrequencyInHours sets the frequencyInHours property value. The time interval in hours to take snapshots (restore points) of a Cloud PC automatically. Possible values are 4, 6, 12, 16, and 24. The default frequency is 12 hours.
func (m *CloudPcRestorePointSetting) SetFrequencyInHours(value *int32)() {
    if m != nil {
        m.frequencyInHours = value
    }
}
// SetUserRestoreEnabled sets the userRestoreEnabled property value. If true, the user has the ability to use snapshots to restore Cloud PCs. If false, non-admin users cannot use snapshots to restore the Cloud PC.
func (m *CloudPcRestorePointSetting) SetUserRestoreEnabled(value *bool)() {
    if m != nil {
        m.userRestoreEnabled = value
    }
}

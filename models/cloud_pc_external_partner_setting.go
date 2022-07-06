package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CloudPcExternalPartnerSetting 
type CloudPcExternalPartnerSetting struct {
    Entity
    // The enableConnection property
    enableConnection *bool
    // The lastSyncDateTime property
    lastSyncDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The partnerId property
    partnerId *string
    // The status property
    status *CloudPcExternalPartnerStatus
    // The statusDetails property
    statusDetails *string
}
// NewCloudPcExternalPartnerSetting instantiates a new CloudPcExternalPartnerSetting and sets the default values.
func NewCloudPcExternalPartnerSetting()(*CloudPcExternalPartnerSetting) {
    m := &CloudPcExternalPartnerSetting{
        Entity: *NewEntity(),
    }
    return m
}
// CreateCloudPcExternalPartnerSettingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCloudPcExternalPartnerSettingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudPcExternalPartnerSetting(), nil
}
// GetEnableConnection gets the enableConnection property value. The enableConnection property
func (m *CloudPcExternalPartnerSetting) GetEnableConnection()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.enableConnection
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CloudPcExternalPartnerSetting) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["enableConnection"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableConnection(val)
        }
        return nil
    }
    res["lastSyncDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastSyncDateTime(val)
        }
        return nil
    }
    res["partnerId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPartnerId(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcExternalPartnerStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*CloudPcExternalPartnerStatus))
        }
        return nil
    }
    res["statusDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatusDetails(val)
        }
        return nil
    }
    return res
}
// GetLastSyncDateTime gets the lastSyncDateTime property value. The lastSyncDateTime property
func (m *CloudPcExternalPartnerSetting) GetLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastSyncDateTime
    }
}
// GetPartnerId gets the partnerId property value. The partnerId property
func (m *CloudPcExternalPartnerSetting) GetPartnerId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.partnerId
    }
}
// GetStatus gets the status property value. The status property
func (m *CloudPcExternalPartnerSetting) GetStatus()(*CloudPcExternalPartnerStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// GetStatusDetails gets the statusDetails property value. The statusDetails property
func (m *CloudPcExternalPartnerSetting) GetStatusDetails()(*string) {
    if m == nil {
        return nil
    } else {
        return m.statusDetails
    }
}
// Serialize serializes information the current object
func (m *CloudPcExternalPartnerSetting) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("enableConnection", m.GetEnableConnection())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastSyncDateTime", m.GetLastSyncDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("partnerId", m.GetPartnerId())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("statusDetails", m.GetStatusDetails())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetEnableConnection sets the enableConnection property value. The enableConnection property
func (m *CloudPcExternalPartnerSetting) SetEnableConnection(value *bool)() {
    if m != nil {
        m.enableConnection = value
    }
}
// SetLastSyncDateTime sets the lastSyncDateTime property value. The lastSyncDateTime property
func (m *CloudPcExternalPartnerSetting) SetLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastSyncDateTime = value
    }
}
// SetPartnerId sets the partnerId property value. The partnerId property
func (m *CloudPcExternalPartnerSetting) SetPartnerId(value *string)() {
    if m != nil {
        m.partnerId = value
    }
}
// SetStatus sets the status property value. The status property
func (m *CloudPcExternalPartnerSetting) SetStatus(value *CloudPcExternalPartnerStatus)() {
    if m != nil {
        m.status = value
    }
}
// SetStatusDetails sets the statusDetails property value. The statusDetails property
func (m *CloudPcExternalPartnerSetting) SetStatusDetails(value *string)() {
    if m != nil {
        m.statusDetails = value
    }
}

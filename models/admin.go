package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Admin 
type Admin struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // A container for Microsoft Edge resources. Read-only.
    edge Edgeable
    // The OdataType property
    odataType *string
    // A container for administrative resources to manage reports.
    reportSettings AdminReportSettingsable
    // A container for service communications resources. Read-only.
    serviceAnnouncement ServiceAnnouncementable
}
// NewAdmin instantiates a new Admin and sets the default values.
func NewAdmin()(*Admin) {
    m := &Admin{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAdminFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAdminFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAdmin(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Admin) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetEdge gets the edge property value. A container for Microsoft Edge resources. Read-only.
func (m *Admin) GetEdge()(Edgeable) {
    return m.edge
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Admin) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["edge"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEdgeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEdge(val.(Edgeable))
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["reportSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAdminReportSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReportSettings(val.(AdminReportSettingsable))
        }
        return nil
    }
    res["serviceAnnouncement"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateServiceAnnouncementFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServiceAnnouncement(val.(ServiceAnnouncementable))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *Admin) GetOdataType()(*string) {
    return m.odataType
}
// GetReportSettings gets the reportSettings property value. A container for administrative resources to manage reports.
func (m *Admin) GetReportSettings()(AdminReportSettingsable) {
    return m.reportSettings
}
// GetServiceAnnouncement gets the serviceAnnouncement property value. A container for service communications resources. Read-only.
func (m *Admin) GetServiceAnnouncement()(ServiceAnnouncementable) {
    return m.serviceAnnouncement
}
// Serialize serializes information the current object
func (m *Admin) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("edge", m.GetEdge())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("reportSettings", m.GetReportSettings())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("serviceAnnouncement", m.GetServiceAnnouncement())
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
func (m *Admin) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetEdge sets the edge property value. A container for Microsoft Edge resources. Read-only.
func (m *Admin) SetEdge(value Edgeable)() {
    m.edge = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *Admin) SetOdataType(value *string)() {
    m.odataType = value
}
// SetReportSettings sets the reportSettings property value. A container for administrative resources to manage reports.
func (m *Admin) SetReportSettings(value AdminReportSettingsable)() {
    m.reportSettings = value
}
// SetServiceAnnouncement sets the serviceAnnouncement property value. A container for service communications resources. Read-only.
func (m *Admin) SetServiceAnnouncement(value ServiceAnnouncementable)() {
    m.serviceAnnouncement = value
}

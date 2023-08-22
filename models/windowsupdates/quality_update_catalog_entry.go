package windowsupdates

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// QualityUpdateCatalogEntry 
type QualityUpdateCatalogEntry struct {
    SoftwareUpdateCatalogEntry
}
// NewQualityUpdateCatalogEntry instantiates a new qualityUpdateCatalogEntry and sets the default values.
func NewQualityUpdateCatalogEntry()(*QualityUpdateCatalogEntry) {
    m := &QualityUpdateCatalogEntry{
        SoftwareUpdateCatalogEntry: *NewSoftwareUpdateCatalogEntry(),
    }
    odataTypeValue := "#microsoft.graph.windowsUpdates.qualityUpdateCatalogEntry"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateQualityUpdateCatalogEntryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateQualityUpdateCatalogEntryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewQualityUpdateCatalogEntry(), nil
}
// GetCatalogName gets the catalogName property value. The catalogName property
func (m *QualityUpdateCatalogEntry) GetCatalogName()(*string) {
    val, err := m.GetBackingStore().Get("catalogName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCveSeverityInformation gets the cveSeverityInformation property value. The cveSeverityInformation property
func (m *QualityUpdateCatalogEntry) GetCveSeverityInformation()(QualityUpdateCveSeverityInformationable) {
    val, err := m.GetBackingStore().Get("cveSeverityInformation")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(QualityUpdateCveSeverityInformationable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *QualityUpdateCatalogEntry) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.SoftwareUpdateCatalogEntry.GetFieldDeserializers()
    res["catalogName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCatalogName(val)
        }
        return nil
    }
    res["cveSeverityInformation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateQualityUpdateCveSeverityInformationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCveSeverityInformation(val.(QualityUpdateCveSeverityInformationable))
        }
        return nil
    }
    res["isExpeditable"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsExpeditable(val)
        }
        return nil
    }
    res["productRevisions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateProductRevisionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ProductRevisionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ProductRevisionable)
                }
            }
            m.SetProductRevisions(res)
        }
        return nil
    }
    res["qualityUpdateCadence"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseQualityUpdateCadence)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQualityUpdateCadence(val.(*QualityUpdateCadence))
        }
        return nil
    }
    res["qualityUpdateClassification"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseQualityUpdateClassification)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQualityUpdateClassification(val.(*QualityUpdateClassification))
        }
        return nil
    }
    res["shortName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShortName(val)
        }
        return nil
    }
    return res
}
// GetIsExpeditable gets the isExpeditable property value. Indicates whether the content can be deployed as an expedited quality update. Read-only.
func (m *QualityUpdateCatalogEntry) GetIsExpeditable()(*bool) {
    val, err := m.GetBackingStore().Get("isExpeditable")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetProductRevisions gets the productRevisions property value. The productRevisions property
func (m *QualityUpdateCatalogEntry) GetProductRevisions()([]ProductRevisionable) {
    val, err := m.GetBackingStore().Get("productRevisions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ProductRevisionable)
    }
    return nil
}
// GetQualityUpdateCadence gets the qualityUpdateCadence property value. The qualityUpdateCadence property
func (m *QualityUpdateCatalogEntry) GetQualityUpdateCadence()(*QualityUpdateCadence) {
    val, err := m.GetBackingStore().Get("qualityUpdateCadence")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*QualityUpdateCadence)
    }
    return nil
}
// GetQualityUpdateClassification gets the qualityUpdateClassification property value. The qualityUpdateClassification property
func (m *QualityUpdateCatalogEntry) GetQualityUpdateClassification()(*QualityUpdateClassification) {
    val, err := m.GetBackingStore().Get("qualityUpdateClassification")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*QualityUpdateClassification)
    }
    return nil
}
// GetShortName gets the shortName property value. The shortName property
func (m *QualityUpdateCatalogEntry) GetShortName()(*string) {
    val, err := m.GetBackingStore().Get("shortName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *QualityUpdateCatalogEntry) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.SoftwareUpdateCatalogEntry.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("catalogName", m.GetCatalogName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("cveSeverityInformation", m.GetCveSeverityInformation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isExpeditable", m.GetIsExpeditable())
        if err != nil {
            return err
        }
    }
    if m.GetProductRevisions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetProductRevisions()))
        for i, v := range m.GetProductRevisions() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("productRevisions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetQualityUpdateCadence() != nil {
        cast := (*m.GetQualityUpdateCadence()).String()
        err = writer.WriteStringValue("qualityUpdateCadence", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetQualityUpdateClassification() != nil {
        cast := (*m.GetQualityUpdateClassification()).String()
        err = writer.WriteStringValue("qualityUpdateClassification", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("shortName", m.GetShortName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCatalogName sets the catalogName property value. The catalogName property
func (m *QualityUpdateCatalogEntry) SetCatalogName(value *string)() {
    err := m.GetBackingStore().Set("catalogName", value)
    if err != nil {
        panic(err)
    }
}
// SetCveSeverityInformation sets the cveSeverityInformation property value. The cveSeverityInformation property
func (m *QualityUpdateCatalogEntry) SetCveSeverityInformation(value QualityUpdateCveSeverityInformationable)() {
    err := m.GetBackingStore().Set("cveSeverityInformation", value)
    if err != nil {
        panic(err)
    }
}
// SetIsExpeditable sets the isExpeditable property value. Indicates whether the content can be deployed as an expedited quality update. Read-only.
func (m *QualityUpdateCatalogEntry) SetIsExpeditable(value *bool)() {
    err := m.GetBackingStore().Set("isExpeditable", value)
    if err != nil {
        panic(err)
    }
}
// SetProductRevisions sets the productRevisions property value. The productRevisions property
func (m *QualityUpdateCatalogEntry) SetProductRevisions(value []ProductRevisionable)() {
    err := m.GetBackingStore().Set("productRevisions", value)
    if err != nil {
        panic(err)
    }
}
// SetQualityUpdateCadence sets the qualityUpdateCadence property value. The qualityUpdateCadence property
func (m *QualityUpdateCatalogEntry) SetQualityUpdateCadence(value *QualityUpdateCadence)() {
    err := m.GetBackingStore().Set("qualityUpdateCadence", value)
    if err != nil {
        panic(err)
    }
}
// SetQualityUpdateClassification sets the qualityUpdateClassification property value. The qualityUpdateClassification property
func (m *QualityUpdateCatalogEntry) SetQualityUpdateClassification(value *QualityUpdateClassification)() {
    err := m.GetBackingStore().Set("qualityUpdateClassification", value)
    if err != nil {
        panic(err)
    }
}
// SetShortName sets the shortName property value. The shortName property
func (m *QualityUpdateCatalogEntry) SetShortName(value *string)() {
    err := m.GetBackingStore().Set("shortName", value)
    if err != nil {
        panic(err)
    }
}
// QualityUpdateCatalogEntryable 
type QualityUpdateCatalogEntryable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    SoftwareUpdateCatalogEntryable
    GetCatalogName()(*string)
    GetCveSeverityInformation()(QualityUpdateCveSeverityInformationable)
    GetIsExpeditable()(*bool)
    GetProductRevisions()([]ProductRevisionable)
    GetQualityUpdateCadence()(*QualityUpdateCadence)
    GetQualityUpdateClassification()(*QualityUpdateClassification)
    GetShortName()(*string)
    SetCatalogName(value *string)()
    SetCveSeverityInformation(value QualityUpdateCveSeverityInformationable)()
    SetIsExpeditable(value *bool)()
    SetProductRevisions(value []ProductRevisionable)()
    SetQualityUpdateCadence(value *QualityUpdateCadence)()
    SetQualityUpdateClassification(value *QualityUpdateClassification)()
    SetShortName(value *string)()
}

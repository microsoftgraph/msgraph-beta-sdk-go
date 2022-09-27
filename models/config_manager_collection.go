package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ConfigManagerCollection a ConfigManager defined collection of devices or users.
type ConfigManagerCollection struct {
    Entity
    // The collection identifier in SCCM.
    collectionIdentifier *string
    // The created date.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The DisplayName.
    displayName *string
    // The Hierarchy Identifier.
    hierarchyIdentifier *string
    // The HierarchyName.
    hierarchyName *string
    // The last modified date.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewConfigManagerCollection instantiates a new configManagerCollection and sets the default values.
func NewConfigManagerCollection()(*ConfigManagerCollection) {
    m := &ConfigManagerCollection{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.configManagerCollection";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateConfigManagerCollectionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateConfigManagerCollectionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewConfigManagerCollection(), nil
}
// GetCollectionIdentifier gets the collectionIdentifier property value. The collection identifier in SCCM.
func (m *ConfigManagerCollection) GetCollectionIdentifier()(*string) {
    return m.collectionIdentifier
}
// GetCreatedDateTime gets the createdDateTime property value. The created date.
func (m *ConfigManagerCollection) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDisplayName gets the displayName property value. The DisplayName.
func (m *ConfigManagerCollection) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ConfigManagerCollection) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["collectionIdentifier"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetCollectionIdentifier)
    res["createdDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetCreatedDateTime)
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["hierarchyIdentifier"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetHierarchyIdentifier)
    res["hierarchyName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetHierarchyName)
    res["lastModifiedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetLastModifiedDateTime)
    return res
}
// GetHierarchyIdentifier gets the hierarchyIdentifier property value. The Hierarchy Identifier.
func (m *ConfigManagerCollection) GetHierarchyIdentifier()(*string) {
    return m.hierarchyIdentifier
}
// GetHierarchyName gets the hierarchyName property value. The HierarchyName.
func (m *ConfigManagerCollection) GetHierarchyName()(*string) {
    return m.hierarchyName
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The last modified date.
func (m *ConfigManagerCollection) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// Serialize serializes information the current object
func (m *ConfigManagerCollection) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("collectionIdentifier", m.GetCollectionIdentifier())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("hierarchyIdentifier", m.GetHierarchyIdentifier())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("hierarchyName", m.GetHierarchyName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCollectionIdentifier sets the collectionIdentifier property value. The collection identifier in SCCM.
func (m *ConfigManagerCollection) SetCollectionIdentifier(value *string)() {
    m.collectionIdentifier = value
}
// SetCreatedDateTime sets the createdDateTime property value. The created date.
func (m *ConfigManagerCollection) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDisplayName sets the displayName property value. The DisplayName.
func (m *ConfigManagerCollection) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetHierarchyIdentifier sets the hierarchyIdentifier property value. The Hierarchy Identifier.
func (m *ConfigManagerCollection) SetHierarchyIdentifier(value *string)() {
    m.hierarchyIdentifier = value
}
// SetHierarchyName sets the hierarchyName property value. The HierarchyName.
func (m *ConfigManagerCollection) SetHierarchyName(value *string)() {
    m.hierarchyName = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The last modified date.
func (m *ConfigManagerCollection) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}

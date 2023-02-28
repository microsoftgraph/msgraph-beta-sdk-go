package externalconnectors

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// ExternalConnection 
type ExternalConnection struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
}
// NewExternalConnection instantiates a new ExternalConnection and sets the default values.
func NewExternalConnection()(*ExternalConnection) {
    m := &ExternalConnection{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreateExternalConnectionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateExternalConnectionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewExternalConnection(), nil
}
// GetActivitySettings gets the activitySettings property value. Collects configurable settings related to activities involving connector content.
func (m *ExternalConnection) GetActivitySettings()(ActivitySettingsable) {
    val, err := m.GetBackingStore().Get("activitySettings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ActivitySettingsable)
    }
    return nil
}
// GetComplianceSettings gets the complianceSettings property value. The settings required for the connection to participate in eDiscovery, such as the display templates for eDiscovery results.
func (m *ExternalConnection) GetComplianceSettings()(ComplianceSettingsable) {
    val, err := m.GetBackingStore().Get("complianceSettings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ComplianceSettingsable)
    }
    return nil
}
// GetConfiguration gets the configuration property value. Specifies additional application IDs that are allowed to manage the connection and to index content in the connection. Optional.
func (m *ExternalConnection) GetConfiguration()(Configurationable) {
    val, err := m.GetBackingStore().Get("configuration")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(Configurationable)
    }
    return nil
}
// GetConnectorId gets the connectorId property value. The Teams App ID. Optional.
func (m *ExternalConnection) GetConnectorId()(*string) {
    val, err := m.GetBackingStore().Get("connectorId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDescription gets the description property value. Description of the connection displayed in the Microsoft 365 admin center. Optional.
func (m *ExternalConnection) GetDescription()(*string) {
    val, err := m.GetBackingStore().Get("description")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetEnabledContentExperiences gets the enabledContentExperiences property value. The list of content experiences the connection will participate in. Possible values are search and compliance.
func (m *ExternalConnection) GetEnabledContentExperiences()(*ContentExperienceType) {
    val, err := m.GetBackingStore().Get("enabledContentExperiences")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ContentExperienceType)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ExternalConnection) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["activitySettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateActivitySettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActivitySettings(val.(ActivitySettingsable))
        }
        return nil
    }
    res["complianceSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateComplianceSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComplianceSettings(val.(ComplianceSettingsable))
        }
        return nil
    }
    res["configuration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfiguration(val.(Configurationable))
        }
        return nil
    }
    res["connectorId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConnectorId(val)
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["enabledContentExperiences"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseContentExperienceType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnabledContentExperiences(val.(*ContentExperienceType))
        }
        return nil
    }
    res["groups"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateExternalGroupFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ExternalGroupable, len(val))
            for i, v := range val {
                res[i] = v.(ExternalGroupable)
            }
            m.SetGroups(res)
        }
        return nil
    }
    res["ingestedItemsCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIngestedItemsCount(val)
        }
        return nil
    }
    res["items"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateExternalItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ExternalItemable, len(val))
            for i, v := range val {
                res[i] = v.(ExternalItemable)
            }
            m.SetItems(res)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["operations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateConnectionOperationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ConnectionOperationable, len(val))
            for i, v := range val {
                res[i] = v.(ConnectionOperationable)
            }
            m.SetOperations(res)
        }
        return nil
    }
    res["quota"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateConnectionQuotaFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQuota(val.(ConnectionQuotaable))
        }
        return nil
    }
    res["schema"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSchemaFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSchema(val.(Schemaable))
        }
        return nil
    }
    res["searchSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSearchSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSearchSettings(val.(SearchSettingsable))
        }
        return nil
    }
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseConnectionState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*ConnectionState))
        }
        return nil
    }
    return res
}
// GetGroups gets the groups property value. The groups property
func (m *ExternalConnection) GetGroups()([]ExternalGroupable) {
    val, err := m.GetBackingStore().Get("groups")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ExternalGroupable)
    }
    return nil
}
// GetIngestedItemsCount gets the ingestedItemsCount property value. The number of items ingested into a connection. This value is refreshed every 15 minutes. If the connection state is draft, then ingestedItemsCount will be null.
func (m *ExternalConnection) GetIngestedItemsCount()(*int64) {
    val, err := m.GetBackingStore().Get("ingestedItemsCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetItems gets the items property value. The items property
func (m *ExternalConnection) GetItems()([]ExternalItemable) {
    val, err := m.GetBackingStore().Get("items")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ExternalItemable)
    }
    return nil
}
// GetName gets the name property value. The display name of the connection to be displayed in the Microsoft 365 admin center. Maximum length of 128 characters. Required.
func (m *ExternalConnection) GetName()(*string) {
    val, err := m.GetBackingStore().Get("name")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOperations gets the operations property value. The operations property
func (m *ExternalConnection) GetOperations()([]ConnectionOperationable) {
    val, err := m.GetBackingStore().Get("operations")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ConnectionOperationable)
    }
    return nil
}
// GetQuota gets the quota property value. The quota property
func (m *ExternalConnection) GetQuota()(ConnectionQuotaable) {
    val, err := m.GetBackingStore().Get("quota")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ConnectionQuotaable)
    }
    return nil
}
// GetSchema gets the schema property value. The schema property
func (m *ExternalConnection) GetSchema()(Schemaable) {
    val, err := m.GetBackingStore().Get("schema")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(Schemaable)
    }
    return nil
}
// GetSearchSettings gets the searchSettings property value. The settings configuring the search experience for content in this connection, such as the display templates for search results.
func (m *ExternalConnection) GetSearchSettings()(SearchSettingsable) {
    val, err := m.GetBackingStore().Get("searchSettings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(SearchSettingsable)
    }
    return nil
}
// GetState gets the state property value. Indicates the current state of the connection. Possible values are draft, ready, obsolete, and limitExceeded. Required.
func (m *ExternalConnection) GetState()(*ConnectionState) {
    val, err := m.GetBackingStore().Get("state")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ConnectionState)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ExternalConnection) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("activitySettings", m.GetActivitySettings())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("complianceSettings", m.GetComplianceSettings())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("configuration", m.GetConfiguration())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("connectorId", m.GetConnectorId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    if m.GetEnabledContentExperiences() != nil {
        cast := (*m.GetEnabledContentExperiences()).String()
        err = writer.WriteStringValue("enabledContentExperiences", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetGroups() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetGroups()))
        for i, v := range m.GetGroups() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("groups", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("ingestedItemsCount", m.GetIngestedItemsCount())
        if err != nil {
            return err
        }
    }
    if m.GetItems() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetItems()))
        for i, v := range m.GetItems() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("items", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    if m.GetOperations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetOperations()))
        for i, v := range m.GetOperations() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("operations", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("quota", m.GetQuota())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("schema", m.GetSchema())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("searchSettings", m.GetSearchSettings())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActivitySettings sets the activitySettings property value. Collects configurable settings related to activities involving connector content.
func (m *ExternalConnection) SetActivitySettings(value ActivitySettingsable)() {
    err := m.GetBackingStore().Set("activitySettings", value)
    if err != nil {
        panic(err)
    }
}
// SetComplianceSettings sets the complianceSettings property value. The settings required for the connection to participate in eDiscovery, such as the display templates for eDiscovery results.
func (m *ExternalConnection) SetComplianceSettings(value ComplianceSettingsable)() {
    err := m.GetBackingStore().Set("complianceSettings", value)
    if err != nil {
        panic(err)
    }
}
// SetConfiguration sets the configuration property value. Specifies additional application IDs that are allowed to manage the connection and to index content in the connection. Optional.
func (m *ExternalConnection) SetConfiguration(value Configurationable)() {
    err := m.GetBackingStore().Set("configuration", value)
    if err != nil {
        panic(err)
    }
}
// SetConnectorId sets the connectorId property value. The Teams App ID. Optional.
func (m *ExternalConnection) SetConnectorId(value *string)() {
    err := m.GetBackingStore().Set("connectorId", value)
    if err != nil {
        panic(err)
    }
}
// SetDescription sets the description property value. Description of the connection displayed in the Microsoft 365 admin center. Optional.
func (m *ExternalConnection) SetDescription(value *string)() {
    err := m.GetBackingStore().Set("description", value)
    if err != nil {
        panic(err)
    }
}
// SetEnabledContentExperiences sets the enabledContentExperiences property value. The list of content experiences the connection will participate in. Possible values are search and compliance.
func (m *ExternalConnection) SetEnabledContentExperiences(value *ContentExperienceType)() {
    err := m.GetBackingStore().Set("enabledContentExperiences", value)
    if err != nil {
        panic(err)
    }
}
// SetGroups sets the groups property value. The groups property
func (m *ExternalConnection) SetGroups(value []ExternalGroupable)() {
    err := m.GetBackingStore().Set("groups", value)
    if err != nil {
        panic(err)
    }
}
// SetIngestedItemsCount sets the ingestedItemsCount property value. The number of items ingested into a connection. This value is refreshed every 15 minutes. If the connection state is draft, then ingestedItemsCount will be null.
func (m *ExternalConnection) SetIngestedItemsCount(value *int64)() {
    err := m.GetBackingStore().Set("ingestedItemsCount", value)
    if err != nil {
        panic(err)
    }
}
// SetItems sets the items property value. The items property
func (m *ExternalConnection) SetItems(value []ExternalItemable)() {
    err := m.GetBackingStore().Set("items", value)
    if err != nil {
        panic(err)
    }
}
// SetName sets the name property value. The display name of the connection to be displayed in the Microsoft 365 admin center. Maximum length of 128 characters. Required.
func (m *ExternalConnection) SetName(value *string)() {
    err := m.GetBackingStore().Set("name", value)
    if err != nil {
        panic(err)
    }
}
// SetOperations sets the operations property value. The operations property
func (m *ExternalConnection) SetOperations(value []ConnectionOperationable)() {
    err := m.GetBackingStore().Set("operations", value)
    if err != nil {
        panic(err)
    }
}
// SetQuota sets the quota property value. The quota property
func (m *ExternalConnection) SetQuota(value ConnectionQuotaable)() {
    err := m.GetBackingStore().Set("quota", value)
    if err != nil {
        panic(err)
    }
}
// SetSchema sets the schema property value. The schema property
func (m *ExternalConnection) SetSchema(value Schemaable)() {
    err := m.GetBackingStore().Set("schema", value)
    if err != nil {
        panic(err)
    }
}
// SetSearchSettings sets the searchSettings property value. The settings configuring the search experience for content in this connection, such as the display templates for search results.
func (m *ExternalConnection) SetSearchSettings(value SearchSettingsable)() {
    err := m.GetBackingStore().Set("searchSettings", value)
    if err != nil {
        panic(err)
    }
}
// SetState sets the state property value. Indicates the current state of the connection. Possible values are draft, ready, obsolete, and limitExceeded. Required.
func (m *ExternalConnection) SetState(value *ConnectionState)() {
    err := m.GetBackingStore().Set("state", value)
    if err != nil {
        panic(err)
    }
}
// ExternalConnectionable 
type ExternalConnectionable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActivitySettings()(ActivitySettingsable)
    GetComplianceSettings()(ComplianceSettingsable)
    GetConfiguration()(Configurationable)
    GetConnectorId()(*string)
    GetDescription()(*string)
    GetEnabledContentExperiences()(*ContentExperienceType)
    GetGroups()([]ExternalGroupable)
    GetIngestedItemsCount()(*int64)
    GetItems()([]ExternalItemable)
    GetName()(*string)
    GetOperations()([]ConnectionOperationable)
    GetQuota()(ConnectionQuotaable)
    GetSchema()(Schemaable)
    GetSearchSettings()(SearchSettingsable)
    GetState()(*ConnectionState)
    SetActivitySettings(value ActivitySettingsable)()
    SetComplianceSettings(value ComplianceSettingsable)()
    SetConfiguration(value Configurationable)()
    SetConnectorId(value *string)()
    SetDescription(value *string)()
    SetEnabledContentExperiences(value *ContentExperienceType)()
    SetGroups(value []ExternalGroupable)()
    SetIngestedItemsCount(value *int64)()
    SetItems(value []ExternalItemable)()
    SetName(value *string)()
    SetOperations(value []ConnectionOperationable)()
    SetQuota(value ConnectionQuotaable)()
    SetSchema(value Schemaable)()
    SetSearchSettings(value SearchSettingsable)()
    SetState(value *ConnectionState)()
}

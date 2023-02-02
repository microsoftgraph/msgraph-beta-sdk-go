package externalconnectors

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// ExternalConnection 
type ExternalConnection struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
    // Collects configurable settings related to activities involving connector content.
    activitySettings ActivitySettingsable
    // The settings required for the connection to participate in eDiscovery, such as the display templates for eDiscovery results.
    complianceSettings ComplianceSettingsable
    // Specifies additional application IDs that are allowed to manage the connection and to index content in the connection. Optional.
    configuration Configurationable
    // The Teams App ID. Optional.
    connectorId *string
    // Description of the connection displayed in the Microsoft 365 admin center. Optional.
    description *string
    // The list of content experiences the connection will participate in. Possible values are search and compliance.
    enabledContentExperiences *ContentExperienceType
    // The groups property
    groups []ExternalGroupable
    // The number of items ingested into a connection. This value is refreshed every 15 minutes. If the connection state is draft, then ingestedItemsCount will be null.
    ingestedItemsCount *int64
    // The items property
    items []ExternalItemable
    // The display name of the connection to be displayed in the Microsoft 365 admin center. Maximum length of 128 characters. Required.
    name *string
    // The operations property
    operations []ConnectionOperationable
    // The quota property
    quota ConnectionQuotaable
    // The schema property
    schema Schemaable
    // The settings configuring the search experience for content in this connection, such as the display templates for search results.
    searchSettings SearchSettingsable
    // Indicates the current state of the connection. Possible values are draft, ready, obsolete, and limitExceeded. Required.
    state *ConnectionState
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
    return m.activitySettings
}
// GetComplianceSettings gets the complianceSettings property value. The settings required for the connection to participate in eDiscovery, such as the display templates for eDiscovery results.
func (m *ExternalConnection) GetComplianceSettings()(ComplianceSettingsable) {
    return m.complianceSettings
}
// GetConfiguration gets the configuration property value. Specifies additional application IDs that are allowed to manage the connection and to index content in the connection. Optional.
func (m *ExternalConnection) GetConfiguration()(Configurationable) {
    return m.configuration
}
// GetConnectorId gets the connectorId property value. The Teams App ID. Optional.
func (m *ExternalConnection) GetConnectorId()(*string) {
    return m.connectorId
}
// GetDescription gets the description property value. Description of the connection displayed in the Microsoft 365 admin center. Optional.
func (m *ExternalConnection) GetDescription()(*string) {
    return m.description
}
// GetEnabledContentExperiences gets the enabledContentExperiences property value. The list of content experiences the connection will participate in. Possible values are search and compliance.
func (m *ExternalConnection) GetEnabledContentExperiences()(*ContentExperienceType) {
    return m.enabledContentExperiences
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
    return m.groups
}
// GetIngestedItemsCount gets the ingestedItemsCount property value. The number of items ingested into a connection. This value is refreshed every 15 minutes. If the connection state is draft, then ingestedItemsCount will be null.
func (m *ExternalConnection) GetIngestedItemsCount()(*int64) {
    return m.ingestedItemsCount
}
// GetItems gets the items property value. The items property
func (m *ExternalConnection) GetItems()([]ExternalItemable) {
    return m.items
}
// GetName gets the name property value. The display name of the connection to be displayed in the Microsoft 365 admin center. Maximum length of 128 characters. Required.
func (m *ExternalConnection) GetName()(*string) {
    return m.name
}
// GetOperations gets the operations property value. The operations property
func (m *ExternalConnection) GetOperations()([]ConnectionOperationable) {
    return m.operations
}
// GetQuota gets the quota property value. The quota property
func (m *ExternalConnection) GetQuota()(ConnectionQuotaable) {
    return m.quota
}
// GetSchema gets the schema property value. The schema property
func (m *ExternalConnection) GetSchema()(Schemaable) {
    return m.schema
}
// GetSearchSettings gets the searchSettings property value. The settings configuring the search experience for content in this connection, such as the display templates for search results.
func (m *ExternalConnection) GetSearchSettings()(SearchSettingsable) {
    return m.searchSettings
}
// GetState gets the state property value. Indicates the current state of the connection. Possible values are draft, ready, obsolete, and limitExceeded. Required.
func (m *ExternalConnection) GetState()(*ConnectionState) {
    return m.state
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
    m.activitySettings = value
}
// SetComplianceSettings sets the complianceSettings property value. The settings required for the connection to participate in eDiscovery, such as the display templates for eDiscovery results.
func (m *ExternalConnection) SetComplianceSettings(value ComplianceSettingsable)() {
    m.complianceSettings = value
}
// SetConfiguration sets the configuration property value. Specifies additional application IDs that are allowed to manage the connection and to index content in the connection. Optional.
func (m *ExternalConnection) SetConfiguration(value Configurationable)() {
    m.configuration = value
}
// SetConnectorId sets the connectorId property value. The Teams App ID. Optional.
func (m *ExternalConnection) SetConnectorId(value *string)() {
    m.connectorId = value
}
// SetDescription sets the description property value. Description of the connection displayed in the Microsoft 365 admin center. Optional.
func (m *ExternalConnection) SetDescription(value *string)() {
    m.description = value
}
// SetEnabledContentExperiences sets the enabledContentExperiences property value. The list of content experiences the connection will participate in. Possible values are search and compliance.
func (m *ExternalConnection) SetEnabledContentExperiences(value *ContentExperienceType)() {
    m.enabledContentExperiences = value
}
// SetGroups sets the groups property value. The groups property
func (m *ExternalConnection) SetGroups(value []ExternalGroupable)() {
    m.groups = value
}
// SetIngestedItemsCount sets the ingestedItemsCount property value. The number of items ingested into a connection. This value is refreshed every 15 minutes. If the connection state is draft, then ingestedItemsCount will be null.
func (m *ExternalConnection) SetIngestedItemsCount(value *int64)() {
    m.ingestedItemsCount = value
}
// SetItems sets the items property value. The items property
func (m *ExternalConnection) SetItems(value []ExternalItemable)() {
    m.items = value
}
// SetName sets the name property value. The display name of the connection to be displayed in the Microsoft 365 admin center. Maximum length of 128 characters. Required.
func (m *ExternalConnection) SetName(value *string)() {
    m.name = value
}
// SetOperations sets the operations property value. The operations property
func (m *ExternalConnection) SetOperations(value []ConnectionOperationable)() {
    m.operations = value
}
// SetQuota sets the quota property value. The quota property
func (m *ExternalConnection) SetQuota(value ConnectionQuotaable)() {
    m.quota = value
}
// SetSchema sets the schema property value. The schema property
func (m *ExternalConnection) SetSchema(value Schemaable)() {
    m.schema = value
}
// SetSearchSettings sets the searchSettings property value. The settings configuring the search experience for content in this connection, such as the display templates for search results.
func (m *ExternalConnection) SetSearchSettings(value SearchSettingsable)() {
    m.searchSettings = value
}
// SetState sets the state property value. Indicates the current state of the connection. Possible values are draft, ready, obsolete, and limitExceeded. Required.
func (m *ExternalConnection) SetState(value *ConnectionState)() {
    m.state = value
}

package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i3af76fce9a0d8c03f22ff90ccd64c93d01bbef0102a1c4e80376e26d2e22a367 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/externalconnectors"
)

// ExternalConnection provides operations to manage the collection of externalConnection entities.
type ExternalConnection struct {
    Entity
    // Specifies additional application IDs that are allowed to manage the connection and to index content in the connection. Optional.
    configuration Configurationable;
    // The Teams App ID. Optional.
    connectorId *string;
    // Description of the connection displayed in the Microsoft 365 admin center. Optional.
    description *string;
    // Read-only. Nullable.
    groups []ExternalGroupable;
    // 
    ingestedItemsCount *int64;
    // Read-only. Nullable.
    items []ExternalItemable;
    // The display name of the connection to be displayed in the Microsoft 365 admin center. Maximum length of 128 characters. Required.
    name *string;
    // Read-only. Nullable.
    operations []ConnectionOperationable;
    // 
    quota ConnectionQuotaable;
    // Read-only. Nullable.
    schema Schemaable;
    // The settings configuring the search experience for content in this connection, such as the display templates for search results.
    searchSettings SearchSettingsable;
    // Indicates the current state of the connection. Possible values are: draft, ready, obsolete, limitExceeded, unknownFutureValue.
    state *i3af76fce9a0d8c03f22ff90ccd64c93d01bbef0102a1c4e80376e26d2e22a367.ConnectionState;
}
// NewExternalConnection instantiates a new externalConnection and sets the default values.
func NewExternalConnection()(*ExternalConnection) {
    m := &ExternalConnection{
        Entity: *NewEntity(),
    }
    return m
}
// CreateExternalConnectionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateExternalConnectionFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewExternalConnection(), nil
}
// GetConfiguration gets the configuration property value. Specifies additional application IDs that are allowed to manage the connection and to index content in the connection. Optional.
func (m *ExternalConnection) GetConfiguration()(Configurationable) {
    if m == nil {
        return nil
    } else {
        return m.configuration
    }
}
// GetConnectorId gets the connectorId property value. The Teams App ID. Optional.
func (m *ExternalConnection) GetConnectorId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.connectorId
    }
}
// GetDescription gets the description property value. Description of the connection displayed in the Microsoft 365 admin center. Optional.
func (m *ExternalConnection) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ExternalConnection) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["configuration"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfiguration(val.(Configurationable))
        }
        return nil
    }
    res["connectorId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConnectorId(val)
        }
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["groups"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
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
    res["ingestedItemsCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIngestedItemsCount(val)
        }
        return nil
    }
    res["items"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
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
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["operations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
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
    res["quota"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateConnectionQuotaFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQuota(val.(ConnectionQuotaable))
        }
        return nil
    }
    res["schema"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateSchemaFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSchema(val.(Schemaable))
        }
        return nil
    }
    res["searchSettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateSearchSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSearchSettings(val.(SearchSettingsable))
        }
        return nil
    }
    res["state"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i3af76fce9a0d8c03f22ff90ccd64c93d01bbef0102a1c4e80376e26d2e22a367.ParseConnectionState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*i3af76fce9a0d8c03f22ff90ccd64c93d01bbef0102a1c4e80376e26d2e22a367.ConnectionState))
        }
        return nil
    }
    return res
}
// GetGroups gets the groups property value. Read-only. Nullable.
func (m *ExternalConnection) GetGroups()([]ExternalGroupable) {
    if m == nil {
        return nil
    } else {
        return m.groups
    }
}
// GetIngestedItemsCount gets the ingestedItemsCount property value. 
func (m *ExternalConnection) GetIngestedItemsCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.ingestedItemsCount
    }
}
// GetItems gets the items property value. Read-only. Nullable.
func (m *ExternalConnection) GetItems()([]ExternalItemable) {
    if m == nil {
        return nil
    } else {
        return m.items
    }
}
// GetName gets the name property value. The display name of the connection to be displayed in the Microsoft 365 admin center. Maximum length of 128 characters. Required.
func (m *ExternalConnection) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// GetOperations gets the operations property value. Read-only. Nullable.
func (m *ExternalConnection) GetOperations()([]ConnectionOperationable) {
    if m == nil {
        return nil
    } else {
        return m.operations
    }
}
// GetQuota gets the quota property value. 
func (m *ExternalConnection) GetQuota()(ConnectionQuotaable) {
    if m == nil {
        return nil
    } else {
        return m.quota
    }
}
// GetSchema gets the schema property value. Read-only. Nullable.
func (m *ExternalConnection) GetSchema()(Schemaable) {
    if m == nil {
        return nil
    } else {
        return m.schema
    }
}
// GetSearchSettings gets the searchSettings property value. The settings configuring the search experience for content in this connection, such as the display templates for search results.
func (m *ExternalConnection) GetSearchSettings()(SearchSettingsable) {
    if m == nil {
        return nil
    } else {
        return m.searchSettings
    }
}
// GetState gets the state property value. Indicates the current state of the connection. Possible values are: draft, ready, obsolete, limitExceeded, unknownFutureValue.
func (m *ExternalConnection) GetState()(*i3af76fce9a0d8c03f22ff90ccd64c93d01bbef0102a1c4e80376e26d2e22a367.ConnectionState) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
func (m *ExternalConnection) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ExternalConnection) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
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
    if m.GetGroups() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetGroups()))
        for i, v := range m.GetGroups() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetItems()))
        for i, v := range m.GetItems() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetOperations()))
        for i, v := range m.GetOperations() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
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
    if m.GetState() != nil {
        cast := (*m.GetState()).String()
        err = writer.WriteStringValue("state", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetConfiguration sets the configuration property value. Specifies additional application IDs that are allowed to manage the connection and to index content in the connection. Optional.
func (m *ExternalConnection) SetConfiguration(value Configurationable)() {
    if m != nil {
        m.configuration = value
    }
}
// SetConnectorId sets the connectorId property value. The Teams App ID. Optional.
func (m *ExternalConnection) SetConnectorId(value *string)() {
    if m != nil {
        m.connectorId = value
    }
}
// SetDescription sets the description property value. Description of the connection displayed in the Microsoft 365 admin center. Optional.
func (m *ExternalConnection) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetGroups sets the groups property value. Read-only. Nullable.
func (m *ExternalConnection) SetGroups(value []ExternalGroupable)() {
    if m != nil {
        m.groups = value
    }
}
// SetIngestedItemsCount sets the ingestedItemsCount property value. 
func (m *ExternalConnection) SetIngestedItemsCount(value *int64)() {
    if m != nil {
        m.ingestedItemsCount = value
    }
}
// SetItems sets the items property value. Read-only. Nullable.
func (m *ExternalConnection) SetItems(value []ExternalItemable)() {
    if m != nil {
        m.items = value
    }
}
// SetName sets the name property value. The display name of the connection to be displayed in the Microsoft 365 admin center. Maximum length of 128 characters. Required.
func (m *ExternalConnection) SetName(value *string)() {
    if m != nil {
        m.name = value
    }
}
// SetOperations sets the operations property value. Read-only. Nullable.
func (m *ExternalConnection) SetOperations(value []ConnectionOperationable)() {
    if m != nil {
        m.operations = value
    }
}
// SetQuota sets the quota property value. 
func (m *ExternalConnection) SetQuota(value ConnectionQuotaable)() {
    if m != nil {
        m.quota = value
    }
}
// SetSchema sets the schema property value. Read-only. Nullable.
func (m *ExternalConnection) SetSchema(value Schemaable)() {
    if m != nil {
        m.schema = value
    }
}
// SetSearchSettings sets the searchSettings property value. The settings configuring the search experience for content in this connection, such as the display templates for search results.
func (m *ExternalConnection) SetSearchSettings(value SearchSettingsable)() {
    if m != nil {
        m.searchSettings = value
    }
}
// SetState sets the state property value. Indicates the current state of the connection. Possible values are: draft, ready, obsolete, limitExceeded, unknownFutureValue.
func (m *ExternalConnection) SetState(value *i3af76fce9a0d8c03f22ff90ccd64c93d01bbef0102a1c4e80376e26d2e22a367.ConnectionState)() {
    if m != nil {
        m.state = value
    }
}

package managedtenants

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// TenantGroup provides operations to manage the tenantRelationship singleton.
type TenantGroup struct {
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Entity
    // A flag indicating whether all managed tenant are included in the tenant group. Required. Read-only.
    allTenantsIncluded *bool;
    // The display name for the tenant group. Optional. Read-only.
    displayName *string;
    // The collection of management action associated with the tenant group. Optional. Read-only.
    managementActions []ManagementActionInfoable;
    // The collection of management intents associated with the tenant group. Optional. Read-only.
    managementIntents []ManagementIntentInfoable;
    // The collection of managed tenant identifiers include in the tenant group. Optional. Read-only.
    tenantIds []string;
}
// NewTenantGroup instantiates a new tenantGroup and sets the default values.
func NewTenantGroup()(*TenantGroup) {
    m := &TenantGroup{
        Entity: *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEntity(),
    }
    return m
}
// CreateTenantGroupFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTenantGroupFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewTenantGroup(), nil
}
// GetAllTenantsIncluded gets the allTenantsIncluded property value. A flag indicating whether all managed tenant are included in the tenant group. Required. Read-only.
func (m *TenantGroup) GetAllTenantsIncluded()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allTenantsIncluded
    }
}
// GetDisplayName gets the displayName property value. The display name for the tenant group. Optional. Read-only.
func (m *TenantGroup) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TenantGroup) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["allTenantsIncluded"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllTenantsIncluded(val)
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["managementActions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagementActionInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagementActionInfoable, len(val))
            for i, v := range val {
                res[i] = v.(ManagementActionInfoable)
            }
            m.SetManagementActions(res)
        }
        return nil
    }
    res["managementIntents"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagementIntentInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagementIntentInfoable, len(val))
            for i, v := range val {
                res[i] = v.(ManagementIntentInfoable)
            }
            m.SetManagementIntents(res)
        }
        return nil
    }
    res["tenantIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetTenantIds(res)
        }
        return nil
    }
    return res
}
// GetManagementActions gets the managementActions property value. The collection of management action associated with the tenant group. Optional. Read-only.
func (m *TenantGroup) GetManagementActions()([]ManagementActionInfoable) {
    if m == nil {
        return nil
    } else {
        return m.managementActions
    }
}
// GetManagementIntents gets the managementIntents property value. The collection of management intents associated with the tenant group. Optional. Read-only.
func (m *TenantGroup) GetManagementIntents()([]ManagementIntentInfoable) {
    if m == nil {
        return nil
    } else {
        return m.managementIntents
    }
}
// GetTenantIds gets the tenantIds property value. The collection of managed tenant identifiers include in the tenant group. Optional. Read-only.
func (m *TenantGroup) GetTenantIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.tenantIds
    }
}
func (m *TenantGroup) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *TenantGroup) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("allTenantsIncluded", m.GetAllTenantsIncluded())
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
    if m.GetManagementActions() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetManagementActions()))
        for i, v := range m.GetManagementActions() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("managementActions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetManagementIntents() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetManagementIntents()))
        for i, v := range m.GetManagementIntents() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("managementIntents", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTenantIds() != nil {
        err = writer.WriteCollectionOfStringValues("tenantIds", m.GetTenantIds())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAllTenantsIncluded sets the allTenantsIncluded property value. A flag indicating whether all managed tenant are included in the tenant group. Required. Read-only.
func (m *TenantGroup) SetAllTenantsIncluded(value *bool)() {
    if m != nil {
        m.allTenantsIncluded = value
    }
}
// SetDisplayName sets the displayName property value. The display name for the tenant group. Optional. Read-only.
func (m *TenantGroup) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetManagementActions sets the managementActions property value. The collection of management action associated with the tenant group. Optional. Read-only.
func (m *TenantGroup) SetManagementActions(value []ManagementActionInfoable)() {
    if m != nil {
        m.managementActions = value
    }
}
// SetManagementIntents sets the managementIntents property value. The collection of management intents associated with the tenant group. Optional. Read-only.
func (m *TenantGroup) SetManagementIntents(value []ManagementIntentInfoable)() {
    if m != nil {
        m.managementIntents = value
    }
}
// SetTenantIds sets the tenantIds property value. The collection of managed tenant identifiers include in the tenant group. Optional. Read-only.
func (m *TenantGroup) SetTenantIds(value []string)() {
    if m != nil {
        m.tenantIds = value
    }
}

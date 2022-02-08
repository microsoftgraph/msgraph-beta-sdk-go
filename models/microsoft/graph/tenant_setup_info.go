package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TenantSetupInfo 
type TenantSetupInfo struct {
    Entity
    // 
    defaultRolesSettings *PrivilegedRoleSettings;
    // 
    firstTimeSetup *bool;
    // 
    relevantRolesSettings []string;
    // 
    setupStatus *SetupStatus;
    // 
    skipSetup *bool;
    // 
    userRolesActions *string;
}
// NewTenantSetupInfo instantiates a new tenantSetupInfo and sets the default values.
func NewTenantSetupInfo()(*TenantSetupInfo) {
    m := &TenantSetupInfo{
        Entity: *NewEntity(),
    }
    return m
}
// GetDefaultRolesSettings gets the defaultRolesSettings property value. 
func (m *TenantSetupInfo) GetDefaultRolesSettings()(*PrivilegedRoleSettings) {
    if m == nil {
        return nil
    } else {
        return m.defaultRolesSettings
    }
}
// GetFirstTimeSetup gets the firstTimeSetup property value. 
func (m *TenantSetupInfo) GetFirstTimeSetup()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.firstTimeSetup
    }
}
// GetRelevantRolesSettings gets the relevantRolesSettings property value. 
func (m *TenantSetupInfo) GetRelevantRolesSettings()([]string) {
    if m == nil {
        return nil
    } else {
        return m.relevantRolesSettings
    }
}
// GetSetupStatus gets the setupStatus property value. 
func (m *TenantSetupInfo) GetSetupStatus()(*SetupStatus) {
    if m == nil {
        return nil
    } else {
        return m.setupStatus
    }
}
// GetSkipSetup gets the skipSetup property value. 
func (m *TenantSetupInfo) GetSkipSetup()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.skipSetup
    }
}
// GetUserRolesActions gets the userRolesActions property value. 
func (m *TenantSetupInfo) GetUserRolesActions()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userRolesActions
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TenantSetupInfo) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["defaultRolesSettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPrivilegedRoleSettings() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultRolesSettings(val.(*PrivilegedRoleSettings))
        }
        return nil
    }
    res["firstTimeSetup"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirstTimeSetup(val)
        }
        return nil
    }
    res["relevantRolesSettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetRelevantRolesSettings(res)
        }
        return nil
    }
    res["setupStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseSetupStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSetupStatus(val.(*SetupStatus))
        }
        return nil
    }
    res["skipSetup"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSkipSetup(val)
        }
        return nil
    }
    res["userRolesActions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserRolesActions(val)
        }
        return nil
    }
    return res
}
func (m *TenantSetupInfo) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *TenantSetupInfo) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("defaultRolesSettings", m.GetDefaultRolesSettings())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("firstTimeSetup", m.GetFirstTimeSetup())
        if err != nil {
            return err
        }
    }
    if m.GetRelevantRolesSettings() != nil {
        err = writer.WriteCollectionOfStringValues("relevantRolesSettings", m.GetRelevantRolesSettings())
        if err != nil {
            return err
        }
    }
    if m.GetSetupStatus() != nil {
        cast := (*m.GetSetupStatus()).String()
        err = writer.WriteStringValue("setupStatus", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("skipSetup", m.GetSkipSetup())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userRolesActions", m.GetUserRolesActions())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDefaultRolesSettings sets the defaultRolesSettings property value. 
func (m *TenantSetupInfo) SetDefaultRolesSettings(value *PrivilegedRoleSettings)() {
    if m != nil {
        m.defaultRolesSettings = value
    }
}
// SetFirstTimeSetup sets the firstTimeSetup property value. 
func (m *TenantSetupInfo) SetFirstTimeSetup(value *bool)() {
    if m != nil {
        m.firstTimeSetup = value
    }
}
// SetRelevantRolesSettings sets the relevantRolesSettings property value. 
func (m *TenantSetupInfo) SetRelevantRolesSettings(value []string)() {
    if m != nil {
        m.relevantRolesSettings = value
    }
}
// SetSetupStatus sets the setupStatus property value. 
func (m *TenantSetupInfo) SetSetupStatus(value *SetupStatus)() {
    if m != nil {
        m.setupStatus = value
    }
}
// SetSkipSetup sets the skipSetup property value. 
func (m *TenantSetupInfo) SetSkipSetup(value *bool)() {
    if m != nil {
        m.skipSetup = value
    }
}
// SetUserRolesActions sets the userRolesActions property value. 
func (m *TenantSetupInfo) SetUserRolesActions(value *string)() {
    if m != nil {
        m.userRolesActions = value
    }
}

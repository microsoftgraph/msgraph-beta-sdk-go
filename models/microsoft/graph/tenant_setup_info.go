package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type TenantSetupInfo struct {
    Entity
    defaultRolesSettings *PrivilegedRoleSettings;
    firstTimeSetup *bool;
    relevantRolesSettings []string;
    setupStatus *SetupStatus;
    skipSetup *bool;
    userRolesActions *string;
}
func NewTenantSetupInfo()(*TenantSetupInfo) {
    m := &TenantSetupInfo{
        Entity: *NewEntity(),
    }
    return m
}
func (m *TenantSetupInfo) GetDefaultRolesSettings()(*PrivilegedRoleSettings) {
    if m == nil {
        return nil
    } else {
        return m.defaultRolesSettings
    }
}
func (m *TenantSetupInfo) GetFirstTimeSetup()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.firstTimeSetup
    }
}
func (m *TenantSetupInfo) GetRelevantRolesSettings()([]string) {
    if m == nil {
        return nil
    } else {
        return m.relevantRolesSettings
    }
}
func (m *TenantSetupInfo) GetSetupStatus()(*SetupStatus) {
    if m == nil {
        return nil
    } else {
        return m.setupStatus
    }
}
func (m *TenantSetupInfo) GetSkipSetup()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.skipSetup
    }
}
func (m *TenantSetupInfo) GetUserRolesActions()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userRolesActions
    }
}
func (m *TenantSetupInfo) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["defaultRolesSettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPrivilegedRoleSettings() })
        if err != nil {
            return err
        }
        m.SetDefaultRolesSettings(val.(*PrivilegedRoleSettings))
        return nil
    }
    res["firstTimeSetup"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetFirstTimeSetup(val)
        return nil
    }
    res["relevantRolesSettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetRelevantRolesSettings(res)
        return nil
    }
    res["setupStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseSetupStatus)
        if err != nil {
            return err
        }
        cast := val.(SetupStatus)
        m.SetSetupStatus(&cast)
        return nil
    }
    res["skipSetup"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetSkipSetup(val)
        return nil
    }
    res["userRolesActions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUserRolesActions(val)
        return nil
    }
    return res
}
func (m *TenantSetupInfo) IsNil()(bool) {
    return m == nil
}
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
    {
        err = writer.WriteCollectionOfStringValues("relevantRolesSettings", m.GetRelevantRolesSettings())
        if err != nil {
            return err
        }
    }
    if m.GetSetupStatus() != nil {
        cast := m.GetSetupStatus().String()
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
func (m *TenantSetupInfo) SetDefaultRolesSettings(value *PrivilegedRoleSettings)() {
    m.defaultRolesSettings = value
}
func (m *TenantSetupInfo) SetFirstTimeSetup(value *bool)() {
    m.firstTimeSetup = value
}
func (m *TenantSetupInfo) SetRelevantRolesSettings(value []string)() {
    m.relevantRolesSettings = value
}
func (m *TenantSetupInfo) SetSetupStatus(value *SetupStatus)() {
    m.setupStatus = value
}
func (m *TenantSetupInfo) SetSkipSetup(value *bool)() {
    m.skipSetup = value
}
func (m *TenantSetupInfo) SetUserRolesActions(value *string)() {
    m.userRolesActions = value
}

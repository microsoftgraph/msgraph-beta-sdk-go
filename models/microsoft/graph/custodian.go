package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type Custodian struct {
    DataSourceContainer
    acknowledgedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    applyHoldToSources *bool;
    email *string;
    siteSources []SiteSource;
    unifiedGroupSources []UnifiedGroupSource;
    userSources []UserSource;
}
func NewCustodian()(*Custodian) {
    m := &Custodian{
        DataSourceContainer: *NewDataSourceContainer(),
    }
    return m
}
func (m *Custodian) GetAcknowledgedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.acknowledgedDateTime
    }
}
func (m *Custodian) GetApplyHoldToSources()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.applyHoldToSources
    }
}
func (m *Custodian) GetEmail()(*string) {
    if m == nil {
        return nil
    } else {
        return m.email
    }
}
func (m *Custodian) GetSiteSources()([]SiteSource) {
    if m == nil {
        return nil
    } else {
        return m.siteSources
    }
}
func (m *Custodian) GetUnifiedGroupSources()([]UnifiedGroupSource) {
    if m == nil {
        return nil
    } else {
        return m.unifiedGroupSources
    }
}
func (m *Custodian) GetUserSources()([]UserSource) {
    if m == nil {
        return nil
    } else {
        return m.userSources
    }
}
func (m *Custodian) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.DataSourceContainer.GetFieldDeserializers()
    res["acknowledgedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetAcknowledgedDateTime(val)
        return nil
    }
    res["applyHoldToSources"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetApplyHoldToSources(val)
        return nil
    }
    res["email"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetEmail(val)
        return nil
    }
    res["siteSources"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSiteSource() })
        if err != nil {
            return err
        }
        res := make([]SiteSource, len(val))
        for i, v := range val {
            res[i] = *(v.(*SiteSource))
        }
        m.SetSiteSources(res)
        return nil
    }
    res["unifiedGroupSources"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUnifiedGroupSource() })
        if err != nil {
            return err
        }
        res := make([]UnifiedGroupSource, len(val))
        for i, v := range val {
            res[i] = *(v.(*UnifiedGroupSource))
        }
        m.SetUnifiedGroupSources(res)
        return nil
    }
    res["userSources"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserSource() })
        if err != nil {
            return err
        }
        res := make([]UserSource, len(val))
        for i, v := range val {
            res[i] = *(v.(*UserSource))
        }
        m.SetUserSources(res)
        return nil
    }
    return res
}
func (m *Custodian) IsNil()(bool) {
    return m == nil
}
func (m *Custodian) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.DataSourceContainer.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("acknowledgedDateTime", m.GetAcknowledgedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("applyHoldToSources", m.GetApplyHoldToSources())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("email", m.GetEmail())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSiteSources()))
        for i, v := range m.GetSiteSources() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("siteSources", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetUnifiedGroupSources()))
        for i, v := range m.GetUnifiedGroupSources() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("unifiedGroupSources", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetUserSources()))
        for i, v := range m.GetUserSources() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("userSources", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *Custodian) SetAcknowledgedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.acknowledgedDateTime = value
}
func (m *Custodian) SetApplyHoldToSources(value *bool)() {
    m.applyHoldToSources = value
}
func (m *Custodian) SetEmail(value *string)() {
    m.email = value
}
func (m *Custodian) SetSiteSources(value []SiteSource)() {
    m.siteSources = value
}
func (m *Custodian) SetUnifiedGroupSources(value []UnifiedGroupSource)() {
    m.unifiedGroupSources = value
}
func (m *Custodian) SetUserSources(value []UserSource)() {
    m.userSources = value
}

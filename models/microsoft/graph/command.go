package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type Command struct {
    Entity
    appServiceName *string;
    error *string;
    packageFamilyName *string;
    payload *PayloadRequest;
    permissionTicket *string;
    postBackUri *string;
    responsepayload *PayloadResponse;
    status *string;
    type_escaped *string;
}
func NewCommand()(*Command) {
    m := &Command{
        Entity: *NewEntity(),
    }
    return m
}
func (m *Command) GetAppServiceName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appServiceName
    }
}
func (m *Command) GetError()(*string) {
    if m == nil {
        return nil
    } else {
        return m.error
    }
}
func (m *Command) GetPackageFamilyName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.packageFamilyName
    }
}
func (m *Command) GetPayload()(*PayloadRequest) {
    if m == nil {
        return nil
    } else {
        return m.payload
    }
}
func (m *Command) GetPermissionTicket()(*string) {
    if m == nil {
        return nil
    } else {
        return m.permissionTicket
    }
}
func (m *Command) GetPostBackUri()(*string) {
    if m == nil {
        return nil
    } else {
        return m.postBackUri
    }
}
func (m *Command) GetResponsepayload()(*PayloadResponse) {
    if m == nil {
        return nil
    } else {
        return m.responsepayload
    }
}
func (m *Command) GetStatus()(*string) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
func (m *Command) GetType_escaped()(*string) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
func (m *Command) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["appServiceName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAppServiceName(val)
        return nil
    }
    res["error"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetError(val)
        return nil
    }
    res["packageFamilyName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPackageFamilyName(val)
        return nil
    }
    res["payload"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPayloadRequest() })
        if err != nil {
            return err
        }
        m.SetPayload(val.(*PayloadRequest))
        return nil
    }
    res["permissionTicket"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPermissionTicket(val)
        return nil
    }
    res["postBackUri"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPostBackUri(val)
        return nil
    }
    res["responsepayload"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPayloadResponse() })
        if err != nil {
            return err
        }
        m.SetResponsepayload(val.(*PayloadResponse))
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetStatus(val)
        return nil
    }
    res["type_escaped"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetType_escaped(val)
        return nil
    }
    return res
}
func (m *Command) IsNil()(bool) {
    return m == nil
}
func (m *Command) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("appServiceName", m.GetAppServiceName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("error", m.GetError())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("packageFamilyName", m.GetPackageFamilyName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("payload", m.GetPayload())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("permissionTicket", m.GetPermissionTicket())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("postBackUri", m.GetPostBackUri())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("responsepayload", m.GetResponsepayload())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("status", m.GetStatus())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("type_escaped", m.GetType_escaped())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *Command) SetAppServiceName(value *string)() {
    m.appServiceName = value
}
func (m *Command) SetError(value *string)() {
    m.error = value
}
func (m *Command) SetPackageFamilyName(value *string)() {
    m.packageFamilyName = value
}
func (m *Command) SetPayload(value *PayloadRequest)() {
    m.payload = value
}
func (m *Command) SetPermissionTicket(value *string)() {
    m.permissionTicket = value
}
func (m *Command) SetPostBackUri(value *string)() {
    m.postBackUri = value
}
func (m *Command) SetResponsepayload(value *PayloadResponse)() {
    m.responsepayload = value
}
func (m *Command) SetStatus(value *string)() {
    m.status = value
}
func (m *Command) SetType_escaped(value *string)() {
    m.type_escaped = value
}

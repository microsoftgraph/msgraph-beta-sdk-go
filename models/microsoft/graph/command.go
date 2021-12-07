package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Command 
type Command struct {
    Entity
    // 
    appServiceName *string;
    // 
    error *string;
    // 
    packageFamilyName *string;
    // 
    payload *PayloadRequest;
    // 
    permissionTicket *string;
    // 
    postBackUri *string;
    // 
    responsepayload *PayloadResponse;
    // 
    status *string;
    // 
    type_escaped *string;
}
// NewCommand instantiates a new command and sets the default values.
func NewCommand()(*Command) {
    m := &Command{
        Entity: *NewEntity(),
    }
    return m
}
// GetAppServiceName gets the appServiceName property value. 
func (m *Command) GetAppServiceName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appServiceName
    }
}
// GetError gets the error property value. 
func (m *Command) GetError()(*string) {
    if m == nil {
        return nil
    } else {
        return m.error
    }
}
// GetPackageFamilyName gets the packageFamilyName property value. 
func (m *Command) GetPackageFamilyName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.packageFamilyName
    }
}
// GetPayload gets the payload property value. 
func (m *Command) GetPayload()(*PayloadRequest) {
    if m == nil {
        return nil
    } else {
        return m.payload
    }
}
// GetPermissionTicket gets the permissionTicket property value. 
func (m *Command) GetPermissionTicket()(*string) {
    if m == nil {
        return nil
    } else {
        return m.permissionTicket
    }
}
// GetPostBackUri gets the postBackUri property value. 
func (m *Command) GetPostBackUri()(*string) {
    if m == nil {
        return nil
    } else {
        return m.postBackUri
    }
}
// GetResponsepayload gets the responsepayload property value. 
func (m *Command) GetResponsepayload()(*PayloadResponse) {
    if m == nil {
        return nil
    } else {
        return m.responsepayload
    }
}
// GetStatus gets the status property value. 
func (m *Command) GetStatus()(*string) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// GetType gets the type property value. 
func (m *Command) GetType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Command) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["appServiceName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppServiceName(val)
        }
        return nil
    }
    res["error"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetError(val)
        }
        return nil
    }
    res["packageFamilyName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPackageFamilyName(val)
        }
        return nil
    }
    res["payload"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPayloadRequest() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPayload(val.(*PayloadRequest))
        }
        return nil
    }
    res["permissionTicket"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPermissionTicket(val)
        }
        return nil
    }
    res["postBackUri"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPostBackUri(val)
        }
        return nil
    }
    res["responsepayload"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPayloadResponse() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResponsepayload(val.(*PayloadResponse))
        }
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val)
        }
        return nil
    }
    res["type"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val)
        }
        return nil
    }
    return res
}
func (m *Command) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
        err = writer.WriteStringValue("type", m.GetType())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAppServiceName sets the appServiceName property value. 
func (m *Command) SetAppServiceName(value *string)() {
    if m != nil {
        m.appServiceName = value
    }
}
// SetError sets the error property value. 
func (m *Command) SetError(value *string)() {
    if m != nil {
        m.error = value
    }
}
// SetPackageFamilyName sets the packageFamilyName property value. 
func (m *Command) SetPackageFamilyName(value *string)() {
    if m != nil {
        m.packageFamilyName = value
    }
}
// SetPayload sets the payload property value. 
func (m *Command) SetPayload(value *PayloadRequest)() {
    if m != nil {
        m.payload = value
    }
}
// SetPermissionTicket sets the permissionTicket property value. 
func (m *Command) SetPermissionTicket(value *string)() {
    if m != nil {
        m.permissionTicket = value
    }
}
// SetPostBackUri sets the postBackUri property value. 
func (m *Command) SetPostBackUri(value *string)() {
    if m != nil {
        m.postBackUri = value
    }
}
// SetResponsepayload sets the responsepayload property value. 
func (m *Command) SetResponsepayload(value *PayloadResponse)() {
    if m != nil {
        m.responsepayload = value
    }
}
// SetStatus sets the status property value. 
func (m *Command) SetStatus(value *string)() {
    if m != nil {
        m.status = value
    }
}
// SetType sets the type property value. 
func (m *Command) SetType(value *string)() {
    if m != nil {
        m.type_escaped = value
    }
}

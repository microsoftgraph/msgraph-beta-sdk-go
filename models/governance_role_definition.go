package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GovernanceRoleDefinition 
type GovernanceRoleDefinition struct {
    Entity
    // The display name of the role definition.
    displayName *string;
    // The external id of the role definition.
    externalId *string;
    // Read-only. The associated resource for the role definition.
    resource GovernanceResourceable;
    // Required. The id of the resource associated with the role definition.
    resourceId *string;
    // The associated role setting for the role definition.
    roleSetting GovernanceRoleSettingable;
    // The templateId property
    templateId *string;
}
// NewGovernanceRoleDefinition instantiates a new governanceRoleDefinition and sets the default values.
func NewGovernanceRoleDefinition()(*GovernanceRoleDefinition) {
    m := &GovernanceRoleDefinition{
        Entity: *NewEntity(),
    }
    return m
}
// CreateGovernanceRoleDefinitionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGovernanceRoleDefinitionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGovernanceRoleDefinition(), nil
}
// GetDisplayName gets the displayName property value. The display name of the role definition.
func (m *GovernanceRoleDefinition) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetExternalId gets the externalId property value. The external id of the role definition.
func (m *GovernanceRoleDefinition) GetExternalId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.externalId
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GovernanceRoleDefinition) GetFieldDeserializers()(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["displayName"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["externalId"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalId(val)
        }
        return nil
    }
    res["resource"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGovernanceResourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResource(val.(GovernanceResourceable))
        }
        return nil
    }
    res["resourceId"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceId(val)
        }
        return nil
    }
    res["roleSetting"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGovernanceRoleSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoleSetting(val.(GovernanceRoleSettingable))
        }
        return nil
    }
    res["templateId"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTemplateId(val)
        }
        return nil
    }
    return res
}
// GetResource gets the resource property value. Read-only. The associated resource for the role definition.
func (m *GovernanceRoleDefinition) GetResource()(GovernanceResourceable) {
    if m == nil {
        return nil
    } else {
        return m.resource
    }
}
// GetResourceId gets the resourceId property value. Required. The id of the resource associated with the role definition.
func (m *GovernanceRoleDefinition) GetResourceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.resourceId
    }
}
// GetRoleSetting gets the roleSetting property value. The associated role setting for the role definition.
func (m *GovernanceRoleDefinition) GetRoleSetting()(GovernanceRoleSettingable) {
    if m == nil {
        return nil
    } else {
        return m.roleSetting
    }
}
// GetTemplateId gets the templateId property value. The templateId property
func (m *GovernanceRoleDefinition) GetTemplateId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.templateId
    }
}
// Serialize serializes information the current object
func (m *GovernanceRoleDefinition) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("externalId", m.GetExternalId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("resource", m.GetResource())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("resourceId", m.GetResourceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("roleSetting", m.GetRoleSetting())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("templateId", m.GetTemplateId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDisplayName sets the displayName property value. The display name of the role definition.
func (m *GovernanceRoleDefinition) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetExternalId sets the externalId property value. The external id of the role definition.
func (m *GovernanceRoleDefinition) SetExternalId(value *string)() {
    if m != nil {
        m.externalId = value
    }
}
// SetResource sets the resource property value. Read-only. The associated resource for the role definition.
func (m *GovernanceRoleDefinition) SetResource(value GovernanceResourceable)() {
    if m != nil {
        m.resource = value
    }
}
// SetResourceId sets the resourceId property value. Required. The id of the resource associated with the role definition.
func (m *GovernanceRoleDefinition) SetResourceId(value *string)() {
    if m != nil {
        m.resourceId = value
    }
}
// SetRoleSetting sets the roleSetting property value. The associated role setting for the role definition.
func (m *GovernanceRoleDefinition) SetRoleSetting(value GovernanceRoleSettingable)() {
    if m != nil {
        m.roleSetting = value
    }
}
// SetTemplateId sets the templateId property value. The templateId property
func (m *GovernanceRoleDefinition) SetTemplateId(value *string)() {
    if m != nil {
        m.templateId = value
    }
}

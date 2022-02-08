package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// CloudPcOrganizationSettings 
type CloudPcOrganizationSettings struct {
    Entity
    // The version of the operating system (OS) to provision on Cloud PCs. The possible values are: windows10, windows11, unknownFutureValue.
    osVersion *CloudPcOperatingSystem;
    // The account type of the user on provisioned Cloud PCs. The possible values are: standardUser, administrator, unknownFutureValue.
    userAccountType *CloudPcUserAccountType;
}
// NewCloudPcOrganizationSettings instantiates a new cloudPcOrganizationSettings and sets the default values.
func NewCloudPcOrganizationSettings()(*CloudPcOrganizationSettings) {
    m := &CloudPcOrganizationSettings{
        Entity: *NewEntity(),
    }
    return m
}
// GetOsVersion gets the osVersion property value. The version of the operating system (OS) to provision on Cloud PCs. The possible values are: windows10, windows11, unknownFutureValue.
func (m *CloudPcOrganizationSettings) GetOsVersion()(*CloudPcOperatingSystem) {
    if m == nil {
        return nil
    } else {
        return m.osVersion
    }
}
// GetUserAccountType gets the userAccountType property value. The account type of the user on provisioned Cloud PCs. The possible values are: standardUser, administrator, unknownFutureValue.
func (m *CloudPcOrganizationSettings) GetUserAccountType()(*CloudPcUserAccountType) {
    if m == nil {
        return nil
    } else {
        return m.userAccountType
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CloudPcOrganizationSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["osVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcOperatingSystem)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsVersion(val.(*CloudPcOperatingSystem))
        }
        return nil
    }
    res["userAccountType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcUserAccountType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserAccountType(val.(*CloudPcUserAccountType))
        }
        return nil
    }
    return res
}
func (m *CloudPcOrganizationSettings) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *CloudPcOrganizationSettings) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetOsVersion() != nil {
        cast := (*m.GetOsVersion()).String()
        err = writer.WriteStringValue("osVersion", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserAccountType() != nil {
        cast := (*m.GetUserAccountType()).String()
        err = writer.WriteStringValue("userAccountType", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetOsVersion sets the osVersion property value. The version of the operating system (OS) to provision on Cloud PCs. The possible values are: windows10, windows11, unknownFutureValue.
func (m *CloudPcOrganizationSettings) SetOsVersion(value *CloudPcOperatingSystem)() {
    if m != nil {
        m.osVersion = value
    }
}
// SetUserAccountType sets the userAccountType property value. The account type of the user on provisioned Cloud PCs. The possible values are: standardUser, administrator, unknownFutureValue.
func (m *CloudPcOrganizationSettings) SetUserAccountType(value *CloudPcUserAccountType)() {
    if m != nil {
        m.userAccountType = value
    }
}

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TenantInformation 
type TenantInformation struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Primary domain name of an Azure AD tenant.
    defaultDomainName *string
    // Display name of an Azure AD tenant.
    displayName *string
    // Name shown to users that sign in to an Azure AD tenant.
    federationBrandName *string
    // The OdataType property
    odataType *string
    // Unique identifier of an Azure AD tenant.
    tenantId *string
}
// NewTenantInformation instantiates a new tenantInformation and sets the default values.
func NewTenantInformation()(*TenantInformation) {
    m := &TenantInformation{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.tenantInformation";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateTenantInformationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTenantInformationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTenantInformation(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TenantInformation) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDefaultDomainName gets the defaultDomainName property value. Primary domain name of an Azure AD tenant.
func (m *TenantInformation) GetDefaultDomainName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.defaultDomainName
    }
}
// GetDisplayName gets the displayName property value. Display name of an Azure AD tenant.
func (m *TenantInformation) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFederationBrandName gets the federationBrandName property value. Name shown to users that sign in to an Azure AD tenant.
func (m *TenantInformation) GetFederationBrandName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.federationBrandName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TenantInformation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["defaultDomainName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultDomainName(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["federationBrandName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFederationBrandName(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["tenantId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantId(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *TenantInformation) GetOdataType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.odataType
    }
}
// GetTenantId gets the tenantId property value. Unique identifier of an Azure AD tenant.
func (m *TenantInformation) GetTenantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tenantId
    }
}
// Serialize serializes information the current object
func (m *TenantInformation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("defaultDomainName", m.GetDefaultDomainName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("federationBrandName", m.GetFederationBrandName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("tenantId", m.GetTenantId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TenantInformation) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDefaultDomainName sets the defaultDomainName property value. Primary domain name of an Azure AD tenant.
func (m *TenantInformation) SetDefaultDomainName(value *string)() {
    if m != nil {
        m.defaultDomainName = value
    }
}
// SetDisplayName sets the displayName property value. Display name of an Azure AD tenant.
func (m *TenantInformation) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetFederationBrandName sets the federationBrandName property value. Name shown to users that sign in to an Azure AD tenant.
func (m *TenantInformation) SetFederationBrandName(value *string)() {
    if m != nil {
        m.federationBrandName = value
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *TenantInformation) SetOdataType(value *string)() {
    if m != nil {
        m.odataType = value
    }
}
// SetTenantId sets the tenantId property value. Unique identifier of an Azure AD tenant.
func (m *TenantInformation) SetTenantId(value *string)() {
    if m != nil {
        m.tenantId = value
    }
}

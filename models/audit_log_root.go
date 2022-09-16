package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AuditLogRoot 
type AuditLogRoot struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The directoryAudits property
    directoryAudits []DirectoryAuditable
    // The directoryProvisioning property
    directoryProvisioning []ProvisioningObjectSummaryable
    // The OdataType property
    odataType *string
    // The provisioning property
    provisioning []ProvisioningObjectSummaryable
    // The signIns property
    signIns []SignInable
}
// NewAuditLogRoot instantiates a new AuditLogRoot and sets the default values.
func NewAuditLogRoot()(*AuditLogRoot) {
    m := &AuditLogRoot{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.auditLogRoot";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAuditLogRootFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAuditLogRootFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuditLogRoot(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AuditLogRoot) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDirectoryAudits gets the directoryAudits property value. The directoryAudits property
func (m *AuditLogRoot) GetDirectoryAudits()([]DirectoryAuditable) {
    return m.directoryAudits
}
// GetDirectoryProvisioning gets the directoryProvisioning property value. The directoryProvisioning property
func (m *AuditLogRoot) GetDirectoryProvisioning()([]ProvisioningObjectSummaryable) {
    return m.directoryProvisioning
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AuditLogRoot) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["directoryAudits"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDirectoryAuditFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DirectoryAuditable, len(val))
            for i, v := range val {
                res[i] = v.(DirectoryAuditable)
            }
            m.SetDirectoryAudits(res)
        }
        return nil
    }
    res["directoryProvisioning"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateProvisioningObjectSummaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ProvisioningObjectSummaryable, len(val))
            for i, v := range val {
                res[i] = v.(ProvisioningObjectSummaryable)
            }
            m.SetDirectoryProvisioning(res)
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
    res["provisioning"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateProvisioningObjectSummaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ProvisioningObjectSummaryable, len(val))
            for i, v := range val {
                res[i] = v.(ProvisioningObjectSummaryable)
            }
            m.SetProvisioning(res)
        }
        return nil
    }
    res["signIns"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSignInFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SignInable, len(val))
            for i, v := range val {
                res[i] = v.(SignInable)
            }
            m.SetSignIns(res)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AuditLogRoot) GetOdataType()(*string) {
    return m.odataType
}
// GetProvisioning gets the provisioning property value. The provisioning property
func (m *AuditLogRoot) GetProvisioning()([]ProvisioningObjectSummaryable) {
    return m.provisioning
}
// GetSignIns gets the signIns property value. The signIns property
func (m *AuditLogRoot) GetSignIns()([]SignInable) {
    return m.signIns
}
// Serialize serializes information the current object
func (m *AuditLogRoot) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetDirectoryAudits() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDirectoryAudits()))
        for i, v := range m.GetDirectoryAudits() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("directoryAudits", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDirectoryProvisioning() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDirectoryProvisioning()))
        for i, v := range m.GetDirectoryProvisioning() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("directoryProvisioning", cast)
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
    if m.GetProvisioning() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetProvisioning()))
        for i, v := range m.GetProvisioning() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("provisioning", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSignIns() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSignIns()))
        for i, v := range m.GetSignIns() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("signIns", cast)
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
func (m *AuditLogRoot) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDirectoryAudits sets the directoryAudits property value. The directoryAudits property
func (m *AuditLogRoot) SetDirectoryAudits(value []DirectoryAuditable)() {
    m.directoryAudits = value
}
// SetDirectoryProvisioning sets the directoryProvisioning property value. The directoryProvisioning property
func (m *AuditLogRoot) SetDirectoryProvisioning(value []ProvisioningObjectSummaryable)() {
    m.directoryProvisioning = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AuditLogRoot) SetOdataType(value *string)() {
    m.odataType = value
}
// SetProvisioning sets the provisioning property value. The provisioning property
func (m *AuditLogRoot) SetProvisioning(value []ProvisioningObjectSummaryable)() {
    m.provisioning = value
}
// SetSignIns sets the signIns property value. The signIns property
func (m *AuditLogRoot) SetSignIns(value []SignInable)() {
    m.signIns = value
}

package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
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
    res["directoryAudits"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDirectoryAuditFromDiscriminatorValue , m.SetDirectoryAudits)
    res["directoryProvisioning"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateProvisioningObjectSummaryFromDiscriminatorValue , m.SetDirectoryProvisioning)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["provisioning"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateProvisioningObjectSummaryFromDiscriminatorValue , m.SetProvisioning)
    res["signIns"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateSignInFromDiscriminatorValue , m.SetSignIns)
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
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetDirectoryAudits())
        err := writer.WriteCollectionOfObjectValues("directoryAudits", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDirectoryProvisioning() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetDirectoryProvisioning())
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
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetProvisioning())
        err := writer.WriteCollectionOfObjectValues("provisioning", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSignIns() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetSignIns())
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

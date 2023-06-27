package managedtenants

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// ManagementTemplateStepTenantSummary 
type ManagementTemplateStepTenantSummary struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
}
// NewManagementTemplateStepTenantSummary instantiates a new ManagementTemplateStepTenantSummary and sets the default values.
func NewManagementTemplateStepTenantSummary()(*ManagementTemplateStepTenantSummary) {
    m := &ManagementTemplateStepTenantSummary{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreateManagementTemplateStepTenantSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagementTemplateStepTenantSummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewManagementTemplateStepTenantSummary(), nil
}
// GetAssignedTenantsCount gets the assignedTenantsCount property value. The assignedTenantsCount property
func (m *ManagementTemplateStepTenantSummary) GetAssignedTenantsCount()(*int32) {
    val, err := m.GetBackingStore().Get("assignedTenantsCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetCompliantTenantsCount gets the compliantTenantsCount property value. The compliantTenantsCount property
func (m *ManagementTemplateStepTenantSummary) GetCompliantTenantsCount()(*int32) {
    val, err := m.GetBackingStore().Get("compliantTenantsCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetCreatedByUserId gets the createdByUserId property value. The createdByUserId property
func (m *ManagementTemplateStepTenantSummary) GetCreatedByUserId()(*string) {
    val, err := m.GetBackingStore().Get("createdByUserId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCreatedDateTime gets the createdDateTime property value. The createdDateTime property
func (m *ManagementTemplateStepTenantSummary) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("createdDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetDismissedTenantsCount gets the dismissedTenantsCount property value. The dismissedTenantsCount property
func (m *ManagementTemplateStepTenantSummary) GetDismissedTenantsCount()(*int32) {
    val, err := m.GetBackingStore().Get("dismissedTenantsCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagementTemplateStepTenantSummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignedTenantsCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignedTenantsCount(val)
        }
        return nil
    }
    res["compliantTenantsCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompliantTenantsCount(val)
        }
        return nil
    }
    res["createdByUserId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedByUserId(val)
        }
        return nil
    }
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["dismissedTenantsCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDismissedTenantsCount(val)
        }
        return nil
    }
    res["ineligibleTenantsCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIneligibleTenantsCount(val)
        }
        return nil
    }
    res["lastActionByUserId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastActionByUserId(val)
        }
        return nil
    }
    res["lastActionDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastActionDateTime(val)
        }
        return nil
    }
    res["managementTemplateCollectionDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagementTemplateCollectionDisplayName(val)
        }
        return nil
    }
    res["managementTemplateCollectionId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagementTemplateCollectionId(val)
        }
        return nil
    }
    res["managementTemplateDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagementTemplateDisplayName(val)
        }
        return nil
    }
    res["managementTemplateId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagementTemplateId(val)
        }
        return nil
    }
    res["managementTemplateStepDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagementTemplateStepDisplayName(val)
        }
        return nil
    }
    res["managementTemplateStepId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagementTemplateStepId(val)
        }
        return nil
    }
    res["notCompliantTenantsCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotCompliantTenantsCount(val)
        }
        return nil
    }
    return res
}
// GetIneligibleTenantsCount gets the ineligibleTenantsCount property value. The ineligibleTenantsCount property
func (m *ManagementTemplateStepTenantSummary) GetIneligibleTenantsCount()(*int32) {
    val, err := m.GetBackingStore().Get("ineligibleTenantsCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetLastActionByUserId gets the lastActionByUserId property value. The lastActionByUserId property
func (m *ManagementTemplateStepTenantSummary) GetLastActionByUserId()(*string) {
    val, err := m.GetBackingStore().Get("lastActionByUserId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetLastActionDateTime gets the lastActionDateTime property value. The lastActionDateTime property
func (m *ManagementTemplateStepTenantSummary) GetLastActionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastActionDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetManagementTemplateCollectionDisplayName gets the managementTemplateCollectionDisplayName property value. The managementTemplateCollectionDisplayName property
func (m *ManagementTemplateStepTenantSummary) GetManagementTemplateCollectionDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("managementTemplateCollectionDisplayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetManagementTemplateCollectionId gets the managementTemplateCollectionId property value. The managementTemplateCollectionId property
func (m *ManagementTemplateStepTenantSummary) GetManagementTemplateCollectionId()(*string) {
    val, err := m.GetBackingStore().Get("managementTemplateCollectionId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetManagementTemplateDisplayName gets the managementTemplateDisplayName property value. The managementTemplateDisplayName property
func (m *ManagementTemplateStepTenantSummary) GetManagementTemplateDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("managementTemplateDisplayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetManagementTemplateId gets the managementTemplateId property value. The managementTemplateId property
func (m *ManagementTemplateStepTenantSummary) GetManagementTemplateId()(*string) {
    val, err := m.GetBackingStore().Get("managementTemplateId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetManagementTemplateStepDisplayName gets the managementTemplateStepDisplayName property value. The managementTemplateStepDisplayName property
func (m *ManagementTemplateStepTenantSummary) GetManagementTemplateStepDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("managementTemplateStepDisplayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetManagementTemplateStepId gets the managementTemplateStepId property value. The managementTemplateStepId property
func (m *ManagementTemplateStepTenantSummary) GetManagementTemplateStepId()(*string) {
    val, err := m.GetBackingStore().Get("managementTemplateStepId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetNotCompliantTenantsCount gets the notCompliantTenantsCount property value. The notCompliantTenantsCount property
func (m *ManagementTemplateStepTenantSummary) GetNotCompliantTenantsCount()(*int32) {
    val, err := m.GetBackingStore().Get("notCompliantTenantsCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ManagementTemplateStepTenantSummary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("assignedTenantsCount", m.GetAssignedTenantsCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("compliantTenantsCount", m.GetCompliantTenantsCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("createdByUserId", m.GetCreatedByUserId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("dismissedTenantsCount", m.GetDismissedTenantsCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("ineligibleTenantsCount", m.GetIneligibleTenantsCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("lastActionByUserId", m.GetLastActionByUserId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastActionDateTime", m.GetLastActionDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("managementTemplateCollectionDisplayName", m.GetManagementTemplateCollectionDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("managementTemplateCollectionId", m.GetManagementTemplateCollectionId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("managementTemplateDisplayName", m.GetManagementTemplateDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("managementTemplateId", m.GetManagementTemplateId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("managementTemplateStepDisplayName", m.GetManagementTemplateStepDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("managementTemplateStepId", m.GetManagementTemplateStepId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("notCompliantTenantsCount", m.GetNotCompliantTenantsCount())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAssignedTenantsCount sets the assignedTenantsCount property value. The assignedTenantsCount property
func (m *ManagementTemplateStepTenantSummary) SetAssignedTenantsCount(value *int32)() {
    err := m.GetBackingStore().Set("assignedTenantsCount", value)
    if err != nil {
        panic(err)
    }
}
// SetCompliantTenantsCount sets the compliantTenantsCount property value. The compliantTenantsCount property
func (m *ManagementTemplateStepTenantSummary) SetCompliantTenantsCount(value *int32)() {
    err := m.GetBackingStore().Set("compliantTenantsCount", value)
    if err != nil {
        panic(err)
    }
}
// SetCreatedByUserId sets the createdByUserId property value. The createdByUserId property
func (m *ManagementTemplateStepTenantSummary) SetCreatedByUserId(value *string)() {
    err := m.GetBackingStore().Set("createdByUserId", value)
    if err != nil {
        panic(err)
    }
}
// SetCreatedDateTime sets the createdDateTime property value. The createdDateTime property
func (m *ManagementTemplateStepTenantSummary) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("createdDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetDismissedTenantsCount sets the dismissedTenantsCount property value. The dismissedTenantsCount property
func (m *ManagementTemplateStepTenantSummary) SetDismissedTenantsCount(value *int32)() {
    err := m.GetBackingStore().Set("dismissedTenantsCount", value)
    if err != nil {
        panic(err)
    }
}
// SetIneligibleTenantsCount sets the ineligibleTenantsCount property value. The ineligibleTenantsCount property
func (m *ManagementTemplateStepTenantSummary) SetIneligibleTenantsCount(value *int32)() {
    err := m.GetBackingStore().Set("ineligibleTenantsCount", value)
    if err != nil {
        panic(err)
    }
}
// SetLastActionByUserId sets the lastActionByUserId property value. The lastActionByUserId property
func (m *ManagementTemplateStepTenantSummary) SetLastActionByUserId(value *string)() {
    err := m.GetBackingStore().Set("lastActionByUserId", value)
    if err != nil {
        panic(err)
    }
}
// SetLastActionDateTime sets the lastActionDateTime property value. The lastActionDateTime property
func (m *ManagementTemplateStepTenantSummary) SetLastActionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastActionDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetManagementTemplateCollectionDisplayName sets the managementTemplateCollectionDisplayName property value. The managementTemplateCollectionDisplayName property
func (m *ManagementTemplateStepTenantSummary) SetManagementTemplateCollectionDisplayName(value *string)() {
    err := m.GetBackingStore().Set("managementTemplateCollectionDisplayName", value)
    if err != nil {
        panic(err)
    }
}
// SetManagementTemplateCollectionId sets the managementTemplateCollectionId property value. The managementTemplateCollectionId property
func (m *ManagementTemplateStepTenantSummary) SetManagementTemplateCollectionId(value *string)() {
    err := m.GetBackingStore().Set("managementTemplateCollectionId", value)
    if err != nil {
        panic(err)
    }
}
// SetManagementTemplateDisplayName sets the managementTemplateDisplayName property value. The managementTemplateDisplayName property
func (m *ManagementTemplateStepTenantSummary) SetManagementTemplateDisplayName(value *string)() {
    err := m.GetBackingStore().Set("managementTemplateDisplayName", value)
    if err != nil {
        panic(err)
    }
}
// SetManagementTemplateId sets the managementTemplateId property value. The managementTemplateId property
func (m *ManagementTemplateStepTenantSummary) SetManagementTemplateId(value *string)() {
    err := m.GetBackingStore().Set("managementTemplateId", value)
    if err != nil {
        panic(err)
    }
}
// SetManagementTemplateStepDisplayName sets the managementTemplateStepDisplayName property value. The managementTemplateStepDisplayName property
func (m *ManagementTemplateStepTenantSummary) SetManagementTemplateStepDisplayName(value *string)() {
    err := m.GetBackingStore().Set("managementTemplateStepDisplayName", value)
    if err != nil {
        panic(err)
    }
}
// SetManagementTemplateStepId sets the managementTemplateStepId property value. The managementTemplateStepId property
func (m *ManagementTemplateStepTenantSummary) SetManagementTemplateStepId(value *string)() {
    err := m.GetBackingStore().Set("managementTemplateStepId", value)
    if err != nil {
        panic(err)
    }
}
// SetNotCompliantTenantsCount sets the notCompliantTenantsCount property value. The notCompliantTenantsCount property
func (m *ManagementTemplateStepTenantSummary) SetNotCompliantTenantsCount(value *int32)() {
    err := m.GetBackingStore().Set("notCompliantTenantsCount", value)
    if err != nil {
        panic(err)
    }
}
// ManagementTemplateStepTenantSummaryable 
type ManagementTemplateStepTenantSummaryable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAssignedTenantsCount()(*int32)
    GetCompliantTenantsCount()(*int32)
    GetCreatedByUserId()(*string)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDismissedTenantsCount()(*int32)
    GetIneligibleTenantsCount()(*int32)
    GetLastActionByUserId()(*string)
    GetLastActionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetManagementTemplateCollectionDisplayName()(*string)
    GetManagementTemplateCollectionId()(*string)
    GetManagementTemplateDisplayName()(*string)
    GetManagementTemplateId()(*string)
    GetManagementTemplateStepDisplayName()(*string)
    GetManagementTemplateStepId()(*string)
    GetNotCompliantTenantsCount()(*int32)
    SetAssignedTenantsCount(value *int32)()
    SetCompliantTenantsCount(value *int32)()
    SetCreatedByUserId(value *string)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDismissedTenantsCount(value *int32)()
    SetIneligibleTenantsCount(value *int32)()
    SetLastActionByUserId(value *string)()
    SetLastActionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetManagementTemplateCollectionDisplayName(value *string)()
    SetManagementTemplateCollectionId(value *string)()
    SetManagementTemplateDisplayName(value *string)()
    SetManagementTemplateId(value *string)()
    SetManagementTemplateStepDisplayName(value *string)()
    SetManagementTemplateStepId(value *string)()
    SetNotCompliantTenantsCount(value *int32)()
}

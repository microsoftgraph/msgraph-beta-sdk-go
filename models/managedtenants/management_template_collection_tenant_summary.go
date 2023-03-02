package managedtenants

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// ManagementTemplateCollectionTenantSummary 
type ManagementTemplateCollectionTenantSummary struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
}
// NewManagementTemplateCollectionTenantSummary instantiates a new managementTemplateCollectionTenantSummary and sets the default values.
func NewManagementTemplateCollectionTenantSummary()(*ManagementTemplateCollectionTenantSummary) {
    m := &ManagementTemplateCollectionTenantSummary{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreateManagementTemplateCollectionTenantSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagementTemplateCollectionTenantSummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewManagementTemplateCollectionTenantSummary(), nil
}
// GetCompleteStepsCount gets the completeStepsCount property value. The completeStepsCount property
func (m *ManagementTemplateCollectionTenantSummary) GetCompleteStepsCount()(*int32) {
    val, err := m.GetBackingStore().Get("completeStepsCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetCompleteUsersCount gets the completeUsersCount property value. The completeUsersCount property
func (m *ManagementTemplateCollectionTenantSummary) GetCompleteUsersCount()(*int32) {
    val, err := m.GetBackingStore().Get("completeUsersCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetCreatedByUserId gets the createdByUserId property value. The createdByUserId property
func (m *ManagementTemplateCollectionTenantSummary) GetCreatedByUserId()(*string) {
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
func (m *ManagementTemplateCollectionTenantSummary) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("createdDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetDismissedStepsCount gets the dismissedStepsCount property value. The dismissedStepsCount property
func (m *ManagementTemplateCollectionTenantSummary) GetDismissedStepsCount()(*int32) {
    val, err := m.GetBackingStore().Get("dismissedStepsCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetExcludedUsersCount gets the excludedUsersCount property value. The excludedUsersCount property
func (m *ManagementTemplateCollectionTenantSummary) GetExcludedUsersCount()(*int32) {
    val, err := m.GetBackingStore().Get("excludedUsersCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetExcludedUsersDistinctCount gets the excludedUsersDistinctCount property value. The excludedUsersDistinctCount property
func (m *ManagementTemplateCollectionTenantSummary) GetExcludedUsersDistinctCount()(*int32) {
    val, err := m.GetBackingStore().Get("excludedUsersDistinctCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagementTemplateCollectionTenantSummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["completeStepsCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompleteStepsCount(val)
        }
        return nil
    }
    res["completeUsersCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompleteUsersCount(val)
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
    res["dismissedStepsCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDismissedStepsCount(val)
        }
        return nil
    }
    res["excludedUsersCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExcludedUsersCount(val)
        }
        return nil
    }
    res["excludedUsersDistinctCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExcludedUsersDistinctCount(val)
        }
        return nil
    }
    res["incompleteStepsCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIncompleteStepsCount(val)
        }
        return nil
    }
    res["incompleteUsersCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIncompleteUsersCount(val)
        }
        return nil
    }
    res["ineligibleStepsCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIneligibleStepsCount(val)
        }
        return nil
    }
    res["isComplete"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsComplete(val)
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
// GetIncompleteStepsCount gets the incompleteStepsCount property value. The incompleteStepsCount property
func (m *ManagementTemplateCollectionTenantSummary) GetIncompleteStepsCount()(*int32) {
    val, err := m.GetBackingStore().Get("incompleteStepsCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetIncompleteUsersCount gets the incompleteUsersCount property value. The incompleteUsersCount property
func (m *ManagementTemplateCollectionTenantSummary) GetIncompleteUsersCount()(*int32) {
    val, err := m.GetBackingStore().Get("incompleteUsersCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetIneligibleStepsCount gets the ineligibleStepsCount property value. The ineligibleStepsCount property
func (m *ManagementTemplateCollectionTenantSummary) GetIneligibleStepsCount()(*int32) {
    val, err := m.GetBackingStore().Get("ineligibleStepsCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetIsComplete gets the isComplete property value. The isComplete property
func (m *ManagementTemplateCollectionTenantSummary) GetIsComplete()(*bool) {
    val, err := m.GetBackingStore().Get("isComplete")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetLastActionByUserId gets the lastActionByUserId property value. The lastActionByUserId property
func (m *ManagementTemplateCollectionTenantSummary) GetLastActionByUserId()(*string) {
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
func (m *ManagementTemplateCollectionTenantSummary) GetLastActionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
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
func (m *ManagementTemplateCollectionTenantSummary) GetManagementTemplateCollectionDisplayName()(*string) {
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
func (m *ManagementTemplateCollectionTenantSummary) GetManagementTemplateCollectionId()(*string) {
    val, err := m.GetBackingStore().Get("managementTemplateCollectionId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTenantId gets the tenantId property value. The tenantId property
func (m *ManagementTemplateCollectionTenantSummary) GetTenantId()(*string) {
    val, err := m.GetBackingStore().Get("tenantId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ManagementTemplateCollectionTenantSummary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("completeStepsCount", m.GetCompleteStepsCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("completeUsersCount", m.GetCompleteUsersCount())
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
        err = writer.WriteInt32Value("dismissedStepsCount", m.GetDismissedStepsCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("excludedUsersCount", m.GetExcludedUsersCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("excludedUsersDistinctCount", m.GetExcludedUsersDistinctCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("incompleteStepsCount", m.GetIncompleteStepsCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("incompleteUsersCount", m.GetIncompleteUsersCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("ineligibleStepsCount", m.GetIneligibleStepsCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isComplete", m.GetIsComplete())
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
        err = writer.WriteStringValue("tenantId", m.GetTenantId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCompleteStepsCount sets the completeStepsCount property value. The completeStepsCount property
func (m *ManagementTemplateCollectionTenantSummary) SetCompleteStepsCount(value *int32)() {
    err := m.GetBackingStore().Set("completeStepsCount", value)
    if err != nil {
        panic(err)
    }
}
// SetCompleteUsersCount sets the completeUsersCount property value. The completeUsersCount property
func (m *ManagementTemplateCollectionTenantSummary) SetCompleteUsersCount(value *int32)() {
    err := m.GetBackingStore().Set("completeUsersCount", value)
    if err != nil {
        panic(err)
    }
}
// SetCreatedByUserId sets the createdByUserId property value. The createdByUserId property
func (m *ManagementTemplateCollectionTenantSummary) SetCreatedByUserId(value *string)() {
    err := m.GetBackingStore().Set("createdByUserId", value)
    if err != nil {
        panic(err)
    }
}
// SetCreatedDateTime sets the createdDateTime property value. The createdDateTime property
func (m *ManagementTemplateCollectionTenantSummary) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("createdDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetDismissedStepsCount sets the dismissedStepsCount property value. The dismissedStepsCount property
func (m *ManagementTemplateCollectionTenantSummary) SetDismissedStepsCount(value *int32)() {
    err := m.GetBackingStore().Set("dismissedStepsCount", value)
    if err != nil {
        panic(err)
    }
}
// SetExcludedUsersCount sets the excludedUsersCount property value. The excludedUsersCount property
func (m *ManagementTemplateCollectionTenantSummary) SetExcludedUsersCount(value *int32)() {
    err := m.GetBackingStore().Set("excludedUsersCount", value)
    if err != nil {
        panic(err)
    }
}
// SetExcludedUsersDistinctCount sets the excludedUsersDistinctCount property value. The excludedUsersDistinctCount property
func (m *ManagementTemplateCollectionTenantSummary) SetExcludedUsersDistinctCount(value *int32)() {
    err := m.GetBackingStore().Set("excludedUsersDistinctCount", value)
    if err != nil {
        panic(err)
    }
}
// SetIncompleteStepsCount sets the incompleteStepsCount property value. The incompleteStepsCount property
func (m *ManagementTemplateCollectionTenantSummary) SetIncompleteStepsCount(value *int32)() {
    err := m.GetBackingStore().Set("incompleteStepsCount", value)
    if err != nil {
        panic(err)
    }
}
// SetIncompleteUsersCount sets the incompleteUsersCount property value. The incompleteUsersCount property
func (m *ManagementTemplateCollectionTenantSummary) SetIncompleteUsersCount(value *int32)() {
    err := m.GetBackingStore().Set("incompleteUsersCount", value)
    if err != nil {
        panic(err)
    }
}
// SetIneligibleStepsCount sets the ineligibleStepsCount property value. The ineligibleStepsCount property
func (m *ManagementTemplateCollectionTenantSummary) SetIneligibleStepsCount(value *int32)() {
    err := m.GetBackingStore().Set("ineligibleStepsCount", value)
    if err != nil {
        panic(err)
    }
}
// SetIsComplete sets the isComplete property value. The isComplete property
func (m *ManagementTemplateCollectionTenantSummary) SetIsComplete(value *bool)() {
    err := m.GetBackingStore().Set("isComplete", value)
    if err != nil {
        panic(err)
    }
}
// SetLastActionByUserId sets the lastActionByUserId property value. The lastActionByUserId property
func (m *ManagementTemplateCollectionTenantSummary) SetLastActionByUserId(value *string)() {
    err := m.GetBackingStore().Set("lastActionByUserId", value)
    if err != nil {
        panic(err)
    }
}
// SetLastActionDateTime sets the lastActionDateTime property value. The lastActionDateTime property
func (m *ManagementTemplateCollectionTenantSummary) SetLastActionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastActionDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetManagementTemplateCollectionDisplayName sets the managementTemplateCollectionDisplayName property value. The managementTemplateCollectionDisplayName property
func (m *ManagementTemplateCollectionTenantSummary) SetManagementTemplateCollectionDisplayName(value *string)() {
    err := m.GetBackingStore().Set("managementTemplateCollectionDisplayName", value)
    if err != nil {
        panic(err)
    }
}
// SetManagementTemplateCollectionId sets the managementTemplateCollectionId property value. The managementTemplateCollectionId property
func (m *ManagementTemplateCollectionTenantSummary) SetManagementTemplateCollectionId(value *string)() {
    err := m.GetBackingStore().Set("managementTemplateCollectionId", value)
    if err != nil {
        panic(err)
    }
}
// SetTenantId sets the tenantId property value. The tenantId property
func (m *ManagementTemplateCollectionTenantSummary) SetTenantId(value *string)() {
    err := m.GetBackingStore().Set("tenantId", value)
    if err != nil {
        panic(err)
    }
}
// ManagementTemplateCollectionTenantSummaryable 
type ManagementTemplateCollectionTenantSummaryable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCompleteStepsCount()(*int32)
    GetCompleteUsersCount()(*int32)
    GetCreatedByUserId()(*string)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDismissedStepsCount()(*int32)
    GetExcludedUsersCount()(*int32)
    GetExcludedUsersDistinctCount()(*int32)
    GetIncompleteStepsCount()(*int32)
    GetIncompleteUsersCount()(*int32)
    GetIneligibleStepsCount()(*int32)
    GetIsComplete()(*bool)
    GetLastActionByUserId()(*string)
    GetLastActionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetManagementTemplateCollectionDisplayName()(*string)
    GetManagementTemplateCollectionId()(*string)
    GetTenantId()(*string)
    SetCompleteStepsCount(value *int32)()
    SetCompleteUsersCount(value *int32)()
    SetCreatedByUserId(value *string)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDismissedStepsCount(value *int32)()
    SetExcludedUsersCount(value *int32)()
    SetExcludedUsersDistinctCount(value *int32)()
    SetIncompleteStepsCount(value *int32)()
    SetIncompleteUsersCount(value *int32)()
    SetIneligibleStepsCount(value *int32)()
    SetIsComplete(value *bool)()
    SetLastActionByUserId(value *string)()
    SetLastActionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetManagementTemplateCollectionDisplayName(value *string)()
    SetManagementTemplateCollectionId(value *string)()
    SetTenantId(value *string)()
}

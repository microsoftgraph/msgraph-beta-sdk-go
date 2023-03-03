package managedtenants

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// ManagementTemplateStep 
type ManagementTemplateStep struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
}
// NewManagementTemplateStep instantiates a new managementTemplateStep and sets the default values.
func NewManagementTemplateStep()(*ManagementTemplateStep) {
    m := &ManagementTemplateStep{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreateManagementTemplateStepFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagementTemplateStepFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewManagementTemplateStep(), nil
}
// GetAcceptedVersion gets the acceptedVersion property value. The acceptedVersion property
func (m *ManagementTemplateStep) GetAcceptedVersion()(ManagementTemplateStepVersionable) {
    val, err := m.GetBackingStore().Get("acceptedVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ManagementTemplateStepVersionable)
    }
    return nil
}
// GetCategory gets the category property value. The category property
func (m *ManagementTemplateStep) GetCategory()(*ManagementCategory) {
    val, err := m.GetBackingStore().Get("category")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ManagementCategory)
    }
    return nil
}
// GetCreatedByUserId gets the createdByUserId property value. The createdByUserId property
func (m *ManagementTemplateStep) GetCreatedByUserId()(*string) {
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
func (m *ManagementTemplateStep) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("createdDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetDescription gets the description property value. The description property
func (m *ManagementTemplateStep) GetDescription()(*string) {
    val, err := m.GetBackingStore().Get("description")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDisplayName gets the displayName property value. The displayName property
func (m *ManagementTemplateStep) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagementTemplateStep) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["acceptedVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateManagementTemplateStepVersionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAcceptedVersion(val.(ManagementTemplateStepVersionable))
        }
        return nil
    }
    res["category"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagementCategory)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory(val.(*ManagementCategory))
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
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
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
    res["managementTemplate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateManagementTemplateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagementTemplate(val.(ManagementTemplateable))
        }
        return nil
    }
    res["portalLink"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateActionUrlFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPortalLink(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ActionUrlable))
        }
        return nil
    }
    res["priority"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPriority(val)
        }
        return nil
    }
    res["versions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagementTemplateStepVersionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagementTemplateStepVersionable, len(val))
            for i, v := range val {
                res[i] = v.(ManagementTemplateStepVersionable)
            }
            m.SetVersions(res)
        }
        return nil
    }
    return res
}
// GetLastActionByUserId gets the lastActionByUserId property value. The lastActionByUserId property
func (m *ManagementTemplateStep) GetLastActionByUserId()(*string) {
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
func (m *ManagementTemplateStep) GetLastActionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastActionDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetManagementTemplate gets the managementTemplate property value. The managementTemplate property
func (m *ManagementTemplateStep) GetManagementTemplate()(ManagementTemplateable) {
    val, err := m.GetBackingStore().Get("managementTemplate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ManagementTemplateable)
    }
    return nil
}
// GetPortalLink gets the portalLink property value. The portalLink property
func (m *ManagementTemplateStep) GetPortalLink()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ActionUrlable) {
    val, err := m.GetBackingStore().Get("portalLink")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ActionUrlable)
    }
    return nil
}
// GetPriority gets the priority property value. The priority property
func (m *ManagementTemplateStep) GetPriority()(*int32) {
    val, err := m.GetBackingStore().Get("priority")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetVersions gets the versions property value. The versions property
func (m *ManagementTemplateStep) GetVersions()([]ManagementTemplateStepVersionable) {
    val, err := m.GetBackingStore().Get("versions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ManagementTemplateStepVersionable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ManagementTemplateStep) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("acceptedVersion", m.GetAcceptedVersion())
        if err != nil {
            return err
        }
    }
    if m.GetCategory() != nil {
        cast := (*m.GetCategory()).String()
        err = writer.WriteStringValue("category", &cast)
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
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
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
        err = writer.WriteObjectValue("managementTemplate", m.GetManagementTemplate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("portalLink", m.GetPortalLink())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("priority", m.GetPriority())
        if err != nil {
            return err
        }
    }
    if m.GetVersions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetVersions()))
        for i, v := range m.GetVersions() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("versions", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAcceptedVersion sets the acceptedVersion property value. The acceptedVersion property
func (m *ManagementTemplateStep) SetAcceptedVersion(value ManagementTemplateStepVersionable)() {
    err := m.GetBackingStore().Set("acceptedVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetCategory sets the category property value. The category property
func (m *ManagementTemplateStep) SetCategory(value *ManagementCategory)() {
    err := m.GetBackingStore().Set("category", value)
    if err != nil {
        panic(err)
    }
}
// SetCreatedByUserId sets the createdByUserId property value. The createdByUserId property
func (m *ManagementTemplateStep) SetCreatedByUserId(value *string)() {
    err := m.GetBackingStore().Set("createdByUserId", value)
    if err != nil {
        panic(err)
    }
}
// SetCreatedDateTime sets the createdDateTime property value. The createdDateTime property
func (m *ManagementTemplateStep) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("createdDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetDescription sets the description property value. The description property
func (m *ManagementTemplateStep) SetDescription(value *string)() {
    err := m.GetBackingStore().Set("description", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *ManagementTemplateStep) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetLastActionByUserId sets the lastActionByUserId property value. The lastActionByUserId property
func (m *ManagementTemplateStep) SetLastActionByUserId(value *string)() {
    err := m.GetBackingStore().Set("lastActionByUserId", value)
    if err != nil {
        panic(err)
    }
}
// SetLastActionDateTime sets the lastActionDateTime property value. The lastActionDateTime property
func (m *ManagementTemplateStep) SetLastActionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastActionDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetManagementTemplate sets the managementTemplate property value. The managementTemplate property
func (m *ManagementTemplateStep) SetManagementTemplate(value ManagementTemplateable)() {
    err := m.GetBackingStore().Set("managementTemplate", value)
    if err != nil {
        panic(err)
    }
}
// SetPortalLink sets the portalLink property value. The portalLink property
func (m *ManagementTemplateStep) SetPortalLink(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ActionUrlable)() {
    err := m.GetBackingStore().Set("portalLink", value)
    if err != nil {
        panic(err)
    }
}
// SetPriority sets the priority property value. The priority property
func (m *ManagementTemplateStep) SetPriority(value *int32)() {
    err := m.GetBackingStore().Set("priority", value)
    if err != nil {
        panic(err)
    }
}
// SetVersions sets the versions property value. The versions property
func (m *ManagementTemplateStep) SetVersions(value []ManagementTemplateStepVersionable)() {
    err := m.GetBackingStore().Set("versions", value)
    if err != nil {
        panic(err)
    }
}
// ManagementTemplateStepable 
type ManagementTemplateStepable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAcceptedVersion()(ManagementTemplateStepVersionable)
    GetCategory()(*ManagementCategory)
    GetCreatedByUserId()(*string)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetLastActionByUserId()(*string)
    GetLastActionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetManagementTemplate()(ManagementTemplateable)
    GetPortalLink()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ActionUrlable)
    GetPriority()(*int32)
    GetVersions()([]ManagementTemplateStepVersionable)
    SetAcceptedVersion(value ManagementTemplateStepVersionable)()
    SetCategory(value *ManagementCategory)()
    SetCreatedByUserId(value *string)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetLastActionByUserId(value *string)()
    SetLastActionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetManagementTemplate(value ManagementTemplateable)()
    SetPortalLink(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ActionUrlable)()
    SetPriority(value *int32)()
    SetVersions(value []ManagementTemplateStepVersionable)()
}

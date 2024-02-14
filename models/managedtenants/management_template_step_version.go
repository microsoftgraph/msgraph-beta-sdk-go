package managedtenants

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

type ManagementTemplateStepVersion struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
}
// NewManagementTemplateStepVersion instantiates a new ManagementTemplateStepVersion and sets the default values.
func NewManagementTemplateStepVersion()(*ManagementTemplateStepVersion) {
    m := &ManagementTemplateStepVersion{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreateManagementTemplateStepVersionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateManagementTemplateStepVersionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewManagementTemplateStepVersion(), nil
}
// GetAcceptedFor gets the acceptedFor property value. The acceptedFor property
// returns a ManagementTemplateStepable when successful
func (m *ManagementTemplateStepVersion) GetAcceptedFor()(ManagementTemplateStepable) {
    val, err := m.GetBackingStore().Get("acceptedFor")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ManagementTemplateStepable)
    }
    return nil
}
// GetContentMarkdown gets the contentMarkdown property value. The contentMarkdown property
// returns a *string when successful
func (m *ManagementTemplateStepVersion) GetContentMarkdown()(*string) {
    val, err := m.GetBackingStore().Get("contentMarkdown")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCreatedByUserId gets the createdByUserId property value. The createdByUserId property
// returns a *string when successful
func (m *ManagementTemplateStepVersion) GetCreatedByUserId()(*string) {
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
// returns a *Time when successful
func (m *ManagementTemplateStepVersion) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("createdDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetDeployments gets the deployments property value. The deployments property
// returns a []ManagementTemplateStepDeploymentable when successful
func (m *ManagementTemplateStepVersion) GetDeployments()([]ManagementTemplateStepDeploymentable) {
    val, err := m.GetBackingStore().Get("deployments")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ManagementTemplateStepDeploymentable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ManagementTemplateStepVersion) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["acceptedFor"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateManagementTemplateStepFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAcceptedFor(val.(ManagementTemplateStepable))
        }
        return nil
    }
    res["contentMarkdown"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentMarkdown(val)
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
    res["deployments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagementTemplateStepDeploymentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagementTemplateStepDeploymentable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ManagementTemplateStepDeploymentable)
                }
            }
            m.SetDeployments(res)
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
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["templateStep"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateManagementTemplateStepFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTemplateStep(val.(ManagementTemplateStepable))
        }
        return nil
    }
    res["version"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersion(val)
        }
        return nil
    }
    res["versionInformation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersionInformation(val)
        }
        return nil
    }
    return res
}
// GetLastActionByUserId gets the lastActionByUserId property value. The lastActionByUserId property
// returns a *string when successful
func (m *ManagementTemplateStepVersion) GetLastActionByUserId()(*string) {
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
// returns a *Time when successful
func (m *ManagementTemplateStepVersion) GetLastActionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastActionDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetName gets the name property value. The name property
// returns a *string when successful
func (m *ManagementTemplateStepVersion) GetName()(*string) {
    val, err := m.GetBackingStore().Get("name")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTemplateStep gets the templateStep property value. The templateStep property
// returns a ManagementTemplateStepable when successful
func (m *ManagementTemplateStepVersion) GetTemplateStep()(ManagementTemplateStepable) {
    val, err := m.GetBackingStore().Get("templateStep")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ManagementTemplateStepable)
    }
    return nil
}
// GetVersion gets the version property value. The version property
// returns a *int32 when successful
func (m *ManagementTemplateStepVersion) GetVersion()(*int32) {
    val, err := m.GetBackingStore().Get("version")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetVersionInformation gets the versionInformation property value. The versionInformation property
// returns a *string when successful
func (m *ManagementTemplateStepVersion) GetVersionInformation()(*string) {
    val, err := m.GetBackingStore().Get("versionInformation")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ManagementTemplateStepVersion) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("acceptedFor", m.GetAcceptedFor())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("contentMarkdown", m.GetContentMarkdown())
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
    if m.GetDeployments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeployments()))
        for i, v := range m.GetDeployments() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("deployments", cast)
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
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("templateStep", m.GetTemplateStep())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("version", m.GetVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("versionInformation", m.GetVersionInformation())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAcceptedFor sets the acceptedFor property value. The acceptedFor property
func (m *ManagementTemplateStepVersion) SetAcceptedFor(value ManagementTemplateStepable)() {
    err := m.GetBackingStore().Set("acceptedFor", value)
    if err != nil {
        panic(err)
    }
}
// SetContentMarkdown sets the contentMarkdown property value. The contentMarkdown property
func (m *ManagementTemplateStepVersion) SetContentMarkdown(value *string)() {
    err := m.GetBackingStore().Set("contentMarkdown", value)
    if err != nil {
        panic(err)
    }
}
// SetCreatedByUserId sets the createdByUserId property value. The createdByUserId property
func (m *ManagementTemplateStepVersion) SetCreatedByUserId(value *string)() {
    err := m.GetBackingStore().Set("createdByUserId", value)
    if err != nil {
        panic(err)
    }
}
// SetCreatedDateTime sets the createdDateTime property value. The createdDateTime property
func (m *ManagementTemplateStepVersion) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("createdDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetDeployments sets the deployments property value. The deployments property
func (m *ManagementTemplateStepVersion) SetDeployments(value []ManagementTemplateStepDeploymentable)() {
    err := m.GetBackingStore().Set("deployments", value)
    if err != nil {
        panic(err)
    }
}
// SetLastActionByUserId sets the lastActionByUserId property value. The lastActionByUserId property
func (m *ManagementTemplateStepVersion) SetLastActionByUserId(value *string)() {
    err := m.GetBackingStore().Set("lastActionByUserId", value)
    if err != nil {
        panic(err)
    }
}
// SetLastActionDateTime sets the lastActionDateTime property value. The lastActionDateTime property
func (m *ManagementTemplateStepVersion) SetLastActionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastActionDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetName sets the name property value. The name property
func (m *ManagementTemplateStepVersion) SetName(value *string)() {
    err := m.GetBackingStore().Set("name", value)
    if err != nil {
        panic(err)
    }
}
// SetTemplateStep sets the templateStep property value. The templateStep property
func (m *ManagementTemplateStepVersion) SetTemplateStep(value ManagementTemplateStepable)() {
    err := m.GetBackingStore().Set("templateStep", value)
    if err != nil {
        panic(err)
    }
}
// SetVersion sets the version property value. The version property
func (m *ManagementTemplateStepVersion) SetVersion(value *int32)() {
    err := m.GetBackingStore().Set("version", value)
    if err != nil {
        panic(err)
    }
}
// SetVersionInformation sets the versionInformation property value. The versionInformation property
func (m *ManagementTemplateStepVersion) SetVersionInformation(value *string)() {
    err := m.GetBackingStore().Set("versionInformation", value)
    if err != nil {
        panic(err)
    }
}
type ManagementTemplateStepVersionable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAcceptedFor()(ManagementTemplateStepable)
    GetContentMarkdown()(*string)
    GetCreatedByUserId()(*string)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDeployments()([]ManagementTemplateStepDeploymentable)
    GetLastActionByUserId()(*string)
    GetLastActionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetName()(*string)
    GetTemplateStep()(ManagementTemplateStepable)
    GetVersion()(*int32)
    GetVersionInformation()(*string)
    SetAcceptedFor(value ManagementTemplateStepable)()
    SetContentMarkdown(value *string)()
    SetCreatedByUserId(value *string)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDeployments(value []ManagementTemplateStepDeploymentable)()
    SetLastActionByUserId(value *string)()
    SetLastActionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetName(value *string)()
    SetTemplateStep(value ManagementTemplateStepable)()
    SetVersion(value *int32)()
    SetVersionInformation(value *string)()
}

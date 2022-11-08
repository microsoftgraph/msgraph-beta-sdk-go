package identitygovernance

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// CustomTaskExtension 
type CustomTaskExtension struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CustomCalloutExtension
    // The callback configuration for a custom extension.
    callbackConfiguration ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CustomExtensionCallbackConfigurationable
    // The unique identifier of the Azure AD user that created the custom task extension.Supports $filter(eq, ne) and $expand.
    createdBy ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable
    // When the custom task extension was created.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The unique identifier of the Azure AD user that modified the custom task extension last.Supports $filter(eq, ne) and $expand.
    lastModifiedBy ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable
    // When the custom extension was last modified.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewCustomTaskExtension instantiates a new CustomTaskExtension and sets the default values.
func NewCustomTaskExtension()(*CustomTaskExtension) {
    m := &CustomTaskExtension{
        CustomCalloutExtension: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewCustomCalloutExtension(),
    }
    odataTypeValue := "#microsoft.graph.identityGovernance.customTaskExtension";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateCustomTaskExtensionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCustomTaskExtensionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCustomTaskExtension(), nil
}
// GetCallbackConfiguration gets the callbackConfiguration property value. The callback configuration for a custom extension.
func (m *CustomTaskExtension) GetCallbackConfiguration()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CustomExtensionCallbackConfigurationable) {
    return m.callbackConfiguration
}
// GetCreatedBy gets the createdBy property value. The unique identifier of the Azure AD user that created the custom task extension.Supports $filter(eq, ne) and $expand.
func (m *CustomTaskExtension) GetCreatedBy()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable) {
    return m.createdBy
}
// GetCreatedDateTime gets the createdDateTime property value. When the custom task extension was created.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
func (m *CustomTaskExtension) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CustomTaskExtension) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.CustomCalloutExtension.GetFieldDeserializers()
    res["callbackConfiguration"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCustomExtensionCallbackConfigurationFromDiscriminatorValue , m.SetCallbackConfiguration)
    res["createdBy"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserFromDiscriminatorValue , m.SetCreatedBy)
    res["createdDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetCreatedDateTime)
    res["lastModifiedBy"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserFromDiscriminatorValue , m.SetLastModifiedBy)
    res["lastModifiedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetLastModifiedDateTime)
    return res
}
// GetLastModifiedBy gets the lastModifiedBy property value. The unique identifier of the Azure AD user that modified the custom task extension last.Supports $filter(eq, ne) and $expand.
func (m *CustomTaskExtension) GetLastModifiedBy()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable) {
    return m.lastModifiedBy
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. When the custom extension was last modified.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
func (m *CustomTaskExtension) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// Serialize serializes information the current object
func (m *CustomTaskExtension) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.CustomCalloutExtension.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("callbackConfiguration", m.GetCallbackConfiguration())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("createdBy", m.GetCreatedBy())
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
        err = writer.WriteObjectValue("lastModifiedBy", m.GetLastModifiedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCallbackConfiguration sets the callbackConfiguration property value. The callback configuration for a custom extension.
func (m *CustomTaskExtension) SetCallbackConfiguration(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CustomExtensionCallbackConfigurationable)() {
    m.callbackConfiguration = value
}
// SetCreatedBy sets the createdBy property value. The unique identifier of the Azure AD user that created the custom task extension.Supports $filter(eq, ne) and $expand.
func (m *CustomTaskExtension) SetCreatedBy(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable)() {
    m.createdBy = value
}
// SetCreatedDateTime sets the createdDateTime property value. When the custom task extension was created.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
func (m *CustomTaskExtension) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetLastModifiedBy sets the lastModifiedBy property value. The unique identifier of the Azure AD user that modified the custom task extension last.Supports $filter(eq, ne) and $expand.
func (m *CustomTaskExtension) SetLastModifiedBy(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable)() {
    m.lastModifiedBy = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. When the custom extension was last modified.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
func (m *CustomTaskExtension) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}

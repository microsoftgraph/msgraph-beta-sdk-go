package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MobileAppRelationshipState describes the installation status details of the child app in the context of UPN and device id.
type MobileAppRelationshipState struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The corresponding device id.
    deviceId *string
    // The error code for install or uninstall failures of target app.
    errorCode *int32
    // A list of possible states for application status on an individual device. When devices contact the Intune service and find targeted application enforcement intent, the status of the enforcement is recorded and becomes accessible in the Graph API. Since the application status is identified during device interaction with the Intune service, status records do not immediately appear upon application group assignment; it is created only after the assignment is evaluated in the service and devices start receiving the policy during check-ins.
    installState *ResultantAppState
    // Enum indicating additional details regarding why an application has a particular install state.
    installStateDetail *ResultantAppStateDetail
    // The OdataType property
    odataType *string
    // The collection of source mobile app's ids.
    sourceIds []string
    // The related target app's display name.
    targetDisplayName *string
    // The related target app's id.
    targetId *string
    // The last sync time of the target app.
    targetLastSyncDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewMobileAppRelationshipState instantiates a new mobileAppRelationshipState and sets the default values.
func NewMobileAppRelationshipState()(*MobileAppRelationshipState) {
    m := &MobileAppRelationshipState{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateMobileAppRelationshipStateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMobileAppRelationshipStateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMobileAppRelationshipState(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MobileAppRelationshipState) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDeviceId gets the deviceId property value. The corresponding device id.
func (m *MobileAppRelationshipState) GetDeviceId()(*string) {
    return m.deviceId
}
// GetErrorCode gets the errorCode property value. The error code for install or uninstall failures of target app.
func (m *MobileAppRelationshipState) GetErrorCode()(*int32) {
    return m.errorCode
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MobileAppRelationshipState) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["deviceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceId(val)
        }
        return nil
    }
    res["errorCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorCode(val)
        }
        return nil
    }
    res["installState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseResultantAppState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInstallState(val.(*ResultantAppState))
        }
        return nil
    }
    res["installStateDetail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseResultantAppStateDetail)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInstallStateDetail(val.(*ResultantAppStateDetail))
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
    res["sourceIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetSourceIds(res)
        }
        return nil
    }
    res["targetDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetDisplayName(val)
        }
        return nil
    }
    res["targetId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetId(val)
        }
        return nil
    }
    res["targetLastSyncDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetLastSyncDateTime(val)
        }
        return nil
    }
    return res
}
// GetInstallState gets the installState property value. A list of possible states for application status on an individual device. When devices contact the Intune service and find targeted application enforcement intent, the status of the enforcement is recorded and becomes accessible in the Graph API. Since the application status is identified during device interaction with the Intune service, status records do not immediately appear upon application group assignment; it is created only after the assignment is evaluated in the service and devices start receiving the policy during check-ins.
func (m *MobileAppRelationshipState) GetInstallState()(*ResultantAppState) {
    return m.installState
}
// GetInstallStateDetail gets the installStateDetail property value. Enum indicating additional details regarding why an application has a particular install state.
func (m *MobileAppRelationshipState) GetInstallStateDetail()(*ResultantAppStateDetail) {
    return m.installStateDetail
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *MobileAppRelationshipState) GetOdataType()(*string) {
    return m.odataType
}
// GetSourceIds gets the sourceIds property value. The collection of source mobile app's ids.
func (m *MobileAppRelationshipState) GetSourceIds()([]string) {
    return m.sourceIds
}
// GetTargetDisplayName gets the targetDisplayName property value. The related target app's display name.
func (m *MobileAppRelationshipState) GetTargetDisplayName()(*string) {
    return m.targetDisplayName
}
// GetTargetId gets the targetId property value. The related target app's id.
func (m *MobileAppRelationshipState) GetTargetId()(*string) {
    return m.targetId
}
// GetTargetLastSyncDateTime gets the targetLastSyncDateTime property value. The last sync time of the target app.
func (m *MobileAppRelationshipState) GetTargetLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.targetLastSyncDateTime
}
// Serialize serializes information the current object
func (m *MobileAppRelationshipState) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("deviceId", m.GetDeviceId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("errorCode", m.GetErrorCode())
        if err != nil {
            return err
        }
    }
    if m.GetInstallState() != nil {
        cast := (*m.GetInstallState()).String()
        err := writer.WriteStringValue("installState", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetInstallStateDetail() != nil {
        cast := (*m.GetInstallStateDetail()).String()
        err := writer.WriteStringValue("installStateDetail", &cast)
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
    if m.GetSourceIds() != nil {
        err := writer.WriteCollectionOfStringValues("sourceIds", m.GetSourceIds())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("targetDisplayName", m.GetTargetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("targetId", m.GetTargetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("targetLastSyncDateTime", m.GetTargetLastSyncDateTime())
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
func (m *MobileAppRelationshipState) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDeviceId sets the deviceId property value. The corresponding device id.
func (m *MobileAppRelationshipState) SetDeviceId(value *string)() {
    m.deviceId = value
}
// SetErrorCode sets the errorCode property value. The error code for install or uninstall failures of target app.
func (m *MobileAppRelationshipState) SetErrorCode(value *int32)() {
    m.errorCode = value
}
// SetInstallState sets the installState property value. A list of possible states for application status on an individual device. When devices contact the Intune service and find targeted application enforcement intent, the status of the enforcement is recorded and becomes accessible in the Graph API. Since the application status is identified during device interaction with the Intune service, status records do not immediately appear upon application group assignment; it is created only after the assignment is evaluated in the service and devices start receiving the policy during check-ins.
func (m *MobileAppRelationshipState) SetInstallState(value *ResultantAppState)() {
    m.installState = value
}
// SetInstallStateDetail sets the installStateDetail property value. Enum indicating additional details regarding why an application has a particular install state.
func (m *MobileAppRelationshipState) SetInstallStateDetail(value *ResultantAppStateDetail)() {
    m.installStateDetail = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *MobileAppRelationshipState) SetOdataType(value *string)() {
    m.odataType = value
}
// SetSourceIds sets the sourceIds property value. The collection of source mobile app's ids.
func (m *MobileAppRelationshipState) SetSourceIds(value []string)() {
    m.sourceIds = value
}
// SetTargetDisplayName sets the targetDisplayName property value. The related target app's display name.
func (m *MobileAppRelationshipState) SetTargetDisplayName(value *string)() {
    m.targetDisplayName = value
}
// SetTargetId sets the targetId property value. The related target app's id.
func (m *MobileAppRelationshipState) SetTargetId(value *string)() {
    m.targetId = value
}
// SetTargetLastSyncDateTime sets the targetLastSyncDateTime property value. The last sync time of the target app.
func (m *MobileAppRelationshipState) SetTargetLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.targetLastSyncDateTime = value
}

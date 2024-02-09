package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// MacOSPrivacyAccessControlItem represents per-process privacy preferences.
type MacOSPrivacyAccessControlItem struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewMacOSPrivacyAccessControlItem instantiates a new MacOSPrivacyAccessControlItem and sets the default values.
func NewMacOSPrivacyAccessControlItem()(*MacOSPrivacyAccessControlItem) {
    m := &MacOSPrivacyAccessControlItem{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateMacOSPrivacyAccessControlItemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMacOSPrivacyAccessControlItemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMacOSPrivacyAccessControlItem(), nil
}
// GetAccessibility gets the accessibility property value. Possible values of a property
// returns a *Enablement when successful
func (m *MacOSPrivacyAccessControlItem) GetAccessibility()(*Enablement) {
    val, err := m.GetBackingStore().Get("accessibility")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*Enablement)
    }
    return nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *MacOSPrivacyAccessControlItem) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetAddressBook gets the addressBook property value. Possible values of a property
// returns a *Enablement when successful
func (m *MacOSPrivacyAccessControlItem) GetAddressBook()(*Enablement) {
    val, err := m.GetBackingStore().Get("addressBook")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*Enablement)
    }
    return nil
}
// GetAppleEventsAllowedReceivers gets the appleEventsAllowedReceivers property value. Allow or deny the app or process to send a restricted Apple event to another app or process. You will need to know the identifier, identifier type, and code requirement of the receiving app or process. This collection can contain a maximum of 500 elements.
// returns a []MacOSAppleEventReceiverable when successful
func (m *MacOSPrivacyAccessControlItem) GetAppleEventsAllowedReceivers()([]MacOSAppleEventReceiverable) {
    val, err := m.GetBackingStore().Get("appleEventsAllowedReceivers")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]MacOSAppleEventReceiverable)
    }
    return nil
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *MacOSPrivacyAccessControlItem) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetBlockCamera gets the blockCamera property value. Block access to camera app.
// returns a *bool when successful
func (m *MacOSPrivacyAccessControlItem) GetBlockCamera()(*bool) {
    val, err := m.GetBackingStore().Get("blockCamera")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetBlockListenEvent gets the blockListenEvent property value. Block the app or process from listening to events from input devices such as mouse, keyboard, and trackpad.Requires macOS 10.15 or later.
// returns a *bool when successful
func (m *MacOSPrivacyAccessControlItem) GetBlockListenEvent()(*bool) {
    val, err := m.GetBackingStore().Get("blockListenEvent")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetBlockMicrophone gets the blockMicrophone property value. Block access to microphone.
// returns a *bool when successful
func (m *MacOSPrivacyAccessControlItem) GetBlockMicrophone()(*bool) {
    val, err := m.GetBackingStore().Get("blockMicrophone")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetBlockScreenCapture gets the blockScreenCapture property value. Block app from capturing contents of system display. Requires macOS 10.15 or later.
// returns a *bool when successful
func (m *MacOSPrivacyAccessControlItem) GetBlockScreenCapture()(*bool) {
    val, err := m.GetBackingStore().Get("blockScreenCapture")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetCalendar gets the calendar property value. Possible values of a property
// returns a *Enablement when successful
func (m *MacOSPrivacyAccessControlItem) GetCalendar()(*Enablement) {
    val, err := m.GetBackingStore().Get("calendar")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*Enablement)
    }
    return nil
}
// GetCodeRequirement gets the codeRequirement property value. Enter the code requirement, which can be obtained with the command 'codesign –display -r –' in the Terminal app. Include everything after '=>'.
// returns a *string when successful
func (m *MacOSPrivacyAccessControlItem) GetCodeRequirement()(*string) {
    val, err := m.GetBackingStore().Get("codeRequirement")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDisplayName gets the displayName property value. The display name of the app, process, or executable.
// returns a *string when successful
func (m *MacOSPrivacyAccessControlItem) GetDisplayName()(*string) {
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
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *MacOSPrivacyAccessControlItem) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["accessibility"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnablement)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessibility(val.(*Enablement))
        }
        return nil
    }
    res["addressBook"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnablement)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAddressBook(val.(*Enablement))
        }
        return nil
    }
    res["appleEventsAllowedReceivers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMacOSAppleEventReceiverFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MacOSAppleEventReceiverable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(MacOSAppleEventReceiverable)
                }
            }
            m.SetAppleEventsAllowedReceivers(res)
        }
        return nil
    }
    res["blockCamera"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBlockCamera(val)
        }
        return nil
    }
    res["blockListenEvent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBlockListenEvent(val)
        }
        return nil
    }
    res["blockMicrophone"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBlockMicrophone(val)
        }
        return nil
    }
    res["blockScreenCapture"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBlockScreenCapture(val)
        }
        return nil
    }
    res["calendar"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnablement)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCalendar(val.(*Enablement))
        }
        return nil
    }
    res["codeRequirement"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCodeRequirement(val)
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
    res["fileProviderPresence"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnablement)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFileProviderPresence(val.(*Enablement))
        }
        return nil
    }
    res["identifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentifier(val)
        }
        return nil
    }
    res["identifierType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMacOSProcessIdentifierType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentifierType(val.(*MacOSProcessIdentifierType))
        }
        return nil
    }
    res["mediaLibrary"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnablement)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMediaLibrary(val.(*Enablement))
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
    res["photos"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnablement)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPhotos(val.(*Enablement))
        }
        return nil
    }
    res["postEvent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnablement)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPostEvent(val.(*Enablement))
        }
        return nil
    }
    res["reminders"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnablement)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReminders(val.(*Enablement))
        }
        return nil
    }
    res["speechRecognition"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnablement)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSpeechRecognition(val.(*Enablement))
        }
        return nil
    }
    res["staticCodeValidation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStaticCodeValidation(val)
        }
        return nil
    }
    res["systemPolicyAllFiles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnablement)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSystemPolicyAllFiles(val.(*Enablement))
        }
        return nil
    }
    res["systemPolicyDesktopFolder"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnablement)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSystemPolicyDesktopFolder(val.(*Enablement))
        }
        return nil
    }
    res["systemPolicyDocumentsFolder"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnablement)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSystemPolicyDocumentsFolder(val.(*Enablement))
        }
        return nil
    }
    res["systemPolicyDownloadsFolder"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnablement)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSystemPolicyDownloadsFolder(val.(*Enablement))
        }
        return nil
    }
    res["systemPolicyNetworkVolumes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnablement)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSystemPolicyNetworkVolumes(val.(*Enablement))
        }
        return nil
    }
    res["systemPolicyRemovableVolumes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnablement)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSystemPolicyRemovableVolumes(val.(*Enablement))
        }
        return nil
    }
    res["systemPolicySystemAdminFiles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnablement)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSystemPolicySystemAdminFiles(val.(*Enablement))
        }
        return nil
    }
    return res
}
// GetFileProviderPresence gets the fileProviderPresence property value. Possible values of a property
// returns a *Enablement when successful
func (m *MacOSPrivacyAccessControlItem) GetFileProviderPresence()(*Enablement) {
    val, err := m.GetBackingStore().Get("fileProviderPresence")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*Enablement)
    }
    return nil
}
// GetIdentifier gets the identifier property value. The bundle ID or path of the app, process, or executable.
// returns a *string when successful
func (m *MacOSPrivacyAccessControlItem) GetIdentifier()(*string) {
    val, err := m.GetBackingStore().Get("identifier")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetIdentifierType gets the identifierType property value. Process identifier types for MacOS Privacy Preferences
// returns a *MacOSProcessIdentifierType when successful
func (m *MacOSPrivacyAccessControlItem) GetIdentifierType()(*MacOSProcessIdentifierType) {
    val, err := m.GetBackingStore().Get("identifierType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*MacOSProcessIdentifierType)
    }
    return nil
}
// GetMediaLibrary gets the mediaLibrary property value. Possible values of a property
// returns a *Enablement when successful
func (m *MacOSPrivacyAccessControlItem) GetMediaLibrary()(*Enablement) {
    val, err := m.GetBackingStore().Get("mediaLibrary")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*Enablement)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *MacOSPrivacyAccessControlItem) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPhotos gets the photos property value. Possible values of a property
// returns a *Enablement when successful
func (m *MacOSPrivacyAccessControlItem) GetPhotos()(*Enablement) {
    val, err := m.GetBackingStore().Get("photos")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*Enablement)
    }
    return nil
}
// GetPostEvent gets the postEvent property value. Possible values of a property
// returns a *Enablement when successful
func (m *MacOSPrivacyAccessControlItem) GetPostEvent()(*Enablement) {
    val, err := m.GetBackingStore().Get("postEvent")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*Enablement)
    }
    return nil
}
// GetReminders gets the reminders property value. Possible values of a property
// returns a *Enablement when successful
func (m *MacOSPrivacyAccessControlItem) GetReminders()(*Enablement) {
    val, err := m.GetBackingStore().Get("reminders")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*Enablement)
    }
    return nil
}
// GetSpeechRecognition gets the speechRecognition property value. Possible values of a property
// returns a *Enablement when successful
func (m *MacOSPrivacyAccessControlItem) GetSpeechRecognition()(*Enablement) {
    val, err := m.GetBackingStore().Get("speechRecognition")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*Enablement)
    }
    return nil
}
// GetStaticCodeValidation gets the staticCodeValidation property value. Statically validates the code requirement. Use this setting if the process invalidates its dynamic code signature.
// returns a *bool when successful
func (m *MacOSPrivacyAccessControlItem) GetStaticCodeValidation()(*bool) {
    val, err := m.GetBackingStore().Get("staticCodeValidation")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetSystemPolicyAllFiles gets the systemPolicyAllFiles property value. Possible values of a property
// returns a *Enablement when successful
func (m *MacOSPrivacyAccessControlItem) GetSystemPolicyAllFiles()(*Enablement) {
    val, err := m.GetBackingStore().Get("systemPolicyAllFiles")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*Enablement)
    }
    return nil
}
// GetSystemPolicyDesktopFolder gets the systemPolicyDesktopFolder property value. Possible values of a property
// returns a *Enablement when successful
func (m *MacOSPrivacyAccessControlItem) GetSystemPolicyDesktopFolder()(*Enablement) {
    val, err := m.GetBackingStore().Get("systemPolicyDesktopFolder")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*Enablement)
    }
    return nil
}
// GetSystemPolicyDocumentsFolder gets the systemPolicyDocumentsFolder property value. Possible values of a property
// returns a *Enablement when successful
func (m *MacOSPrivacyAccessControlItem) GetSystemPolicyDocumentsFolder()(*Enablement) {
    val, err := m.GetBackingStore().Get("systemPolicyDocumentsFolder")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*Enablement)
    }
    return nil
}
// GetSystemPolicyDownloadsFolder gets the systemPolicyDownloadsFolder property value. Possible values of a property
// returns a *Enablement when successful
func (m *MacOSPrivacyAccessControlItem) GetSystemPolicyDownloadsFolder()(*Enablement) {
    val, err := m.GetBackingStore().Get("systemPolicyDownloadsFolder")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*Enablement)
    }
    return nil
}
// GetSystemPolicyNetworkVolumes gets the systemPolicyNetworkVolumes property value. Possible values of a property
// returns a *Enablement when successful
func (m *MacOSPrivacyAccessControlItem) GetSystemPolicyNetworkVolumes()(*Enablement) {
    val, err := m.GetBackingStore().Get("systemPolicyNetworkVolumes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*Enablement)
    }
    return nil
}
// GetSystemPolicyRemovableVolumes gets the systemPolicyRemovableVolumes property value. Possible values of a property
// returns a *Enablement when successful
func (m *MacOSPrivacyAccessControlItem) GetSystemPolicyRemovableVolumes()(*Enablement) {
    val, err := m.GetBackingStore().Get("systemPolicyRemovableVolumes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*Enablement)
    }
    return nil
}
// GetSystemPolicySystemAdminFiles gets the systemPolicySystemAdminFiles property value. Possible values of a property
// returns a *Enablement when successful
func (m *MacOSPrivacyAccessControlItem) GetSystemPolicySystemAdminFiles()(*Enablement) {
    val, err := m.GetBackingStore().Get("systemPolicySystemAdminFiles")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*Enablement)
    }
    return nil
}
// Serialize serializes information the current object
func (m *MacOSPrivacyAccessControlItem) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAccessibility() != nil {
        cast := (*m.GetAccessibility()).String()
        err := writer.WriteStringValue("accessibility", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetAddressBook() != nil {
        cast := (*m.GetAddressBook()).String()
        err := writer.WriteStringValue("addressBook", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetAppleEventsAllowedReceivers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAppleEventsAllowedReceivers()))
        for i, v := range m.GetAppleEventsAllowedReceivers() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("appleEventsAllowedReceivers", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("blockCamera", m.GetBlockCamera())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("blockListenEvent", m.GetBlockListenEvent())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("blockMicrophone", m.GetBlockMicrophone())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("blockScreenCapture", m.GetBlockScreenCapture())
        if err != nil {
            return err
        }
    }
    if m.GetCalendar() != nil {
        cast := (*m.GetCalendar()).String()
        err := writer.WriteStringValue("calendar", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("codeRequirement", m.GetCodeRequirement())
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
    if m.GetFileProviderPresence() != nil {
        cast := (*m.GetFileProviderPresence()).String()
        err := writer.WriteStringValue("fileProviderPresence", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("identifier", m.GetIdentifier())
        if err != nil {
            return err
        }
    }
    if m.GetIdentifierType() != nil {
        cast := (*m.GetIdentifierType()).String()
        err := writer.WriteStringValue("identifierType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetMediaLibrary() != nil {
        cast := (*m.GetMediaLibrary()).String()
        err := writer.WriteStringValue("mediaLibrary", &cast)
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
    if m.GetPhotos() != nil {
        cast := (*m.GetPhotos()).String()
        err := writer.WriteStringValue("photos", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetPostEvent() != nil {
        cast := (*m.GetPostEvent()).String()
        err := writer.WriteStringValue("postEvent", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetReminders() != nil {
        cast := (*m.GetReminders()).String()
        err := writer.WriteStringValue("reminders", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetSpeechRecognition() != nil {
        cast := (*m.GetSpeechRecognition()).String()
        err := writer.WriteStringValue("speechRecognition", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("staticCodeValidation", m.GetStaticCodeValidation())
        if err != nil {
            return err
        }
    }
    if m.GetSystemPolicyAllFiles() != nil {
        cast := (*m.GetSystemPolicyAllFiles()).String()
        err := writer.WriteStringValue("systemPolicyAllFiles", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetSystemPolicyDesktopFolder() != nil {
        cast := (*m.GetSystemPolicyDesktopFolder()).String()
        err := writer.WriteStringValue("systemPolicyDesktopFolder", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetSystemPolicyDocumentsFolder() != nil {
        cast := (*m.GetSystemPolicyDocumentsFolder()).String()
        err := writer.WriteStringValue("systemPolicyDocumentsFolder", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetSystemPolicyDownloadsFolder() != nil {
        cast := (*m.GetSystemPolicyDownloadsFolder()).String()
        err := writer.WriteStringValue("systemPolicyDownloadsFolder", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetSystemPolicyNetworkVolumes() != nil {
        cast := (*m.GetSystemPolicyNetworkVolumes()).String()
        err := writer.WriteStringValue("systemPolicyNetworkVolumes", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetSystemPolicyRemovableVolumes() != nil {
        cast := (*m.GetSystemPolicyRemovableVolumes()).String()
        err := writer.WriteStringValue("systemPolicyRemovableVolumes", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetSystemPolicySystemAdminFiles() != nil {
        cast := (*m.GetSystemPolicySystemAdminFiles()).String()
        err := writer.WriteStringValue("systemPolicySystemAdminFiles", &cast)
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
// SetAccessibility sets the accessibility property value. Possible values of a property
func (m *MacOSPrivacyAccessControlItem) SetAccessibility(value *Enablement)() {
    err := m.GetBackingStore().Set("accessibility", value)
    if err != nil {
        panic(err)
    }
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MacOSPrivacyAccessControlItem) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetAddressBook sets the addressBook property value. Possible values of a property
func (m *MacOSPrivacyAccessControlItem) SetAddressBook(value *Enablement)() {
    err := m.GetBackingStore().Set("addressBook", value)
    if err != nil {
        panic(err)
    }
}
// SetAppleEventsAllowedReceivers sets the appleEventsAllowedReceivers property value. Allow or deny the app or process to send a restricted Apple event to another app or process. You will need to know the identifier, identifier type, and code requirement of the receiving app or process. This collection can contain a maximum of 500 elements.
func (m *MacOSPrivacyAccessControlItem) SetAppleEventsAllowedReceivers(value []MacOSAppleEventReceiverable)() {
    err := m.GetBackingStore().Set("appleEventsAllowedReceivers", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *MacOSPrivacyAccessControlItem) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetBlockCamera sets the blockCamera property value. Block access to camera app.
func (m *MacOSPrivacyAccessControlItem) SetBlockCamera(value *bool)() {
    err := m.GetBackingStore().Set("blockCamera", value)
    if err != nil {
        panic(err)
    }
}
// SetBlockListenEvent sets the blockListenEvent property value. Block the app or process from listening to events from input devices such as mouse, keyboard, and trackpad.Requires macOS 10.15 or later.
func (m *MacOSPrivacyAccessControlItem) SetBlockListenEvent(value *bool)() {
    err := m.GetBackingStore().Set("blockListenEvent", value)
    if err != nil {
        panic(err)
    }
}
// SetBlockMicrophone sets the blockMicrophone property value. Block access to microphone.
func (m *MacOSPrivacyAccessControlItem) SetBlockMicrophone(value *bool)() {
    err := m.GetBackingStore().Set("blockMicrophone", value)
    if err != nil {
        panic(err)
    }
}
// SetBlockScreenCapture sets the blockScreenCapture property value. Block app from capturing contents of system display. Requires macOS 10.15 or later.
func (m *MacOSPrivacyAccessControlItem) SetBlockScreenCapture(value *bool)() {
    err := m.GetBackingStore().Set("blockScreenCapture", value)
    if err != nil {
        panic(err)
    }
}
// SetCalendar sets the calendar property value. Possible values of a property
func (m *MacOSPrivacyAccessControlItem) SetCalendar(value *Enablement)() {
    err := m.GetBackingStore().Set("calendar", value)
    if err != nil {
        panic(err)
    }
}
// SetCodeRequirement sets the codeRequirement property value. Enter the code requirement, which can be obtained with the command 'codesign –display -r –' in the Terminal app. Include everything after '=>'.
func (m *MacOSPrivacyAccessControlItem) SetCodeRequirement(value *string)() {
    err := m.GetBackingStore().Set("codeRequirement", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. The display name of the app, process, or executable.
func (m *MacOSPrivacyAccessControlItem) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetFileProviderPresence sets the fileProviderPresence property value. Possible values of a property
func (m *MacOSPrivacyAccessControlItem) SetFileProviderPresence(value *Enablement)() {
    err := m.GetBackingStore().Set("fileProviderPresence", value)
    if err != nil {
        panic(err)
    }
}
// SetIdentifier sets the identifier property value. The bundle ID or path of the app, process, or executable.
func (m *MacOSPrivacyAccessControlItem) SetIdentifier(value *string)() {
    err := m.GetBackingStore().Set("identifier", value)
    if err != nil {
        panic(err)
    }
}
// SetIdentifierType sets the identifierType property value. Process identifier types for MacOS Privacy Preferences
func (m *MacOSPrivacyAccessControlItem) SetIdentifierType(value *MacOSProcessIdentifierType)() {
    err := m.GetBackingStore().Set("identifierType", value)
    if err != nil {
        panic(err)
    }
}
// SetMediaLibrary sets the mediaLibrary property value. Possible values of a property
func (m *MacOSPrivacyAccessControlItem) SetMediaLibrary(value *Enablement)() {
    err := m.GetBackingStore().Set("mediaLibrary", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *MacOSPrivacyAccessControlItem) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetPhotos sets the photos property value. Possible values of a property
func (m *MacOSPrivacyAccessControlItem) SetPhotos(value *Enablement)() {
    err := m.GetBackingStore().Set("photos", value)
    if err != nil {
        panic(err)
    }
}
// SetPostEvent sets the postEvent property value. Possible values of a property
func (m *MacOSPrivacyAccessControlItem) SetPostEvent(value *Enablement)() {
    err := m.GetBackingStore().Set("postEvent", value)
    if err != nil {
        panic(err)
    }
}
// SetReminders sets the reminders property value. Possible values of a property
func (m *MacOSPrivacyAccessControlItem) SetReminders(value *Enablement)() {
    err := m.GetBackingStore().Set("reminders", value)
    if err != nil {
        panic(err)
    }
}
// SetSpeechRecognition sets the speechRecognition property value. Possible values of a property
func (m *MacOSPrivacyAccessControlItem) SetSpeechRecognition(value *Enablement)() {
    err := m.GetBackingStore().Set("speechRecognition", value)
    if err != nil {
        panic(err)
    }
}
// SetStaticCodeValidation sets the staticCodeValidation property value. Statically validates the code requirement. Use this setting if the process invalidates its dynamic code signature.
func (m *MacOSPrivacyAccessControlItem) SetStaticCodeValidation(value *bool)() {
    err := m.GetBackingStore().Set("staticCodeValidation", value)
    if err != nil {
        panic(err)
    }
}
// SetSystemPolicyAllFiles sets the systemPolicyAllFiles property value. Possible values of a property
func (m *MacOSPrivacyAccessControlItem) SetSystemPolicyAllFiles(value *Enablement)() {
    err := m.GetBackingStore().Set("systemPolicyAllFiles", value)
    if err != nil {
        panic(err)
    }
}
// SetSystemPolicyDesktopFolder sets the systemPolicyDesktopFolder property value. Possible values of a property
func (m *MacOSPrivacyAccessControlItem) SetSystemPolicyDesktopFolder(value *Enablement)() {
    err := m.GetBackingStore().Set("systemPolicyDesktopFolder", value)
    if err != nil {
        panic(err)
    }
}
// SetSystemPolicyDocumentsFolder sets the systemPolicyDocumentsFolder property value. Possible values of a property
func (m *MacOSPrivacyAccessControlItem) SetSystemPolicyDocumentsFolder(value *Enablement)() {
    err := m.GetBackingStore().Set("systemPolicyDocumentsFolder", value)
    if err != nil {
        panic(err)
    }
}
// SetSystemPolicyDownloadsFolder sets the systemPolicyDownloadsFolder property value. Possible values of a property
func (m *MacOSPrivacyAccessControlItem) SetSystemPolicyDownloadsFolder(value *Enablement)() {
    err := m.GetBackingStore().Set("systemPolicyDownloadsFolder", value)
    if err != nil {
        panic(err)
    }
}
// SetSystemPolicyNetworkVolumes sets the systemPolicyNetworkVolumes property value. Possible values of a property
func (m *MacOSPrivacyAccessControlItem) SetSystemPolicyNetworkVolumes(value *Enablement)() {
    err := m.GetBackingStore().Set("systemPolicyNetworkVolumes", value)
    if err != nil {
        panic(err)
    }
}
// SetSystemPolicyRemovableVolumes sets the systemPolicyRemovableVolumes property value. Possible values of a property
func (m *MacOSPrivacyAccessControlItem) SetSystemPolicyRemovableVolumes(value *Enablement)() {
    err := m.GetBackingStore().Set("systemPolicyRemovableVolumes", value)
    if err != nil {
        panic(err)
    }
}
// SetSystemPolicySystemAdminFiles sets the systemPolicySystemAdminFiles property value. Possible values of a property
func (m *MacOSPrivacyAccessControlItem) SetSystemPolicySystemAdminFiles(value *Enablement)() {
    err := m.GetBackingStore().Set("systemPolicySystemAdminFiles", value)
    if err != nil {
        panic(err)
    }
}
type MacOSPrivacyAccessControlItemable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccessibility()(*Enablement)
    GetAddressBook()(*Enablement)
    GetAppleEventsAllowedReceivers()([]MacOSAppleEventReceiverable)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetBlockCamera()(*bool)
    GetBlockListenEvent()(*bool)
    GetBlockMicrophone()(*bool)
    GetBlockScreenCapture()(*bool)
    GetCalendar()(*Enablement)
    GetCodeRequirement()(*string)
    GetDisplayName()(*string)
    GetFileProviderPresence()(*Enablement)
    GetIdentifier()(*string)
    GetIdentifierType()(*MacOSProcessIdentifierType)
    GetMediaLibrary()(*Enablement)
    GetOdataType()(*string)
    GetPhotos()(*Enablement)
    GetPostEvent()(*Enablement)
    GetReminders()(*Enablement)
    GetSpeechRecognition()(*Enablement)
    GetStaticCodeValidation()(*bool)
    GetSystemPolicyAllFiles()(*Enablement)
    GetSystemPolicyDesktopFolder()(*Enablement)
    GetSystemPolicyDocumentsFolder()(*Enablement)
    GetSystemPolicyDownloadsFolder()(*Enablement)
    GetSystemPolicyNetworkVolumes()(*Enablement)
    GetSystemPolicyRemovableVolumes()(*Enablement)
    GetSystemPolicySystemAdminFiles()(*Enablement)
    SetAccessibility(value *Enablement)()
    SetAddressBook(value *Enablement)()
    SetAppleEventsAllowedReceivers(value []MacOSAppleEventReceiverable)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetBlockCamera(value *bool)()
    SetBlockListenEvent(value *bool)()
    SetBlockMicrophone(value *bool)()
    SetBlockScreenCapture(value *bool)()
    SetCalendar(value *Enablement)()
    SetCodeRequirement(value *string)()
    SetDisplayName(value *string)()
    SetFileProviderPresence(value *Enablement)()
    SetIdentifier(value *string)()
    SetIdentifierType(value *MacOSProcessIdentifierType)()
    SetMediaLibrary(value *Enablement)()
    SetOdataType(value *string)()
    SetPhotos(value *Enablement)()
    SetPostEvent(value *Enablement)()
    SetReminders(value *Enablement)()
    SetSpeechRecognition(value *Enablement)()
    SetStaticCodeValidation(value *bool)()
    SetSystemPolicyAllFiles(value *Enablement)()
    SetSystemPolicyDesktopFolder(value *Enablement)()
    SetSystemPolicyDocumentsFolder(value *Enablement)()
    SetSystemPolicyDownloadsFolder(value *Enablement)()
    SetSystemPolicyNetworkVolumes(value *Enablement)()
    SetSystemPolicyRemovableVolumes(value *Enablement)()
    SetSystemPolicySystemAdminFiles(value *Enablement)()
}

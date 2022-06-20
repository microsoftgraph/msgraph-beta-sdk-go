package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MacOSPrivacyAccessControlItem represents per-process privacy preferences.
type MacOSPrivacyAccessControlItem struct {
    // Allow the app or process to control the Mac via the Accessibility subsystem. Possible values are: notConfigured, enabled, disabled.
    accessibility *Enablement
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Allow or block access to contact information managed by Contacts. Possible values are: notConfigured, enabled, disabled.
    addressBook *Enablement
    // Allow or deny the app or process to send a restricted Apple event to another app or process. You will need to know the identifier, identifier type, and code requirement of the receiving app or process. This collection can contain a maximum of 500 elements.
    appleEventsAllowedReceivers []MacOSAppleEventReceiverable
    // Block access to camera app.
    blockCamera *bool
    // Block the app or process from listening to events from input devices such as mouse, keyboard, and trackpad.Requires macOS 10.15 or later.
    blockListenEvent *bool
    // Block access to microphone.
    blockMicrophone *bool
    // Block app from capturing contents of system display. Requires macOS 10.15 or later.
    blockScreenCapture *bool
    // Allow or block access to event information managed by Calendar. Possible values are: notConfigured, enabled, disabled.
    calendar *Enablement
    // Enter the code requirement, which can be obtained with the command 'codesign –display -r –' in the Terminal app. Include everything after '=>'.
    codeRequirement *string
    // The display name of the app, process, or executable.
    displayName *string
    // Allow the app or process to access files managed by another app’s file provider extension. Requires macOS 10.15 or later. . Possible values are: notConfigured, enabled, disabled.
    fileProviderPresence *Enablement
    // The bundle ID or path of the app, process, or executable.
    identifier *string
    // A bundle ID is used to identify an app. A path is used to identify a process or executable. Possible values are: bundleID, path.
    identifierType *MacOSProcessIdentifierType
    // Allow or block access to music and the media library. Possible values are: notConfigured, enabled, disabled.
    mediaLibrary *Enablement
    // Allow or block access to images managed by Photos. Possible values are: notConfigured, enabled, disabled.
    photos *Enablement
    // Control access to CoreGraphics APIs, which are used to send CGEvents to the system event stream. Possible values are: notConfigured, enabled, disabled.
    postEvent *Enablement
    // Allow or block access to information managed by Reminders. Possible values are: notConfigured, enabled, disabled.
    reminders *Enablement
    // Allow or block access to system speech recognition facility. Possible values are: notConfigured, enabled, disabled.
    speechRecognition *Enablement
    // Statically validates the code requirement. Use this setting if the process invalidates its dynamic code signature.
    staticCodeValidation *bool
    // Control access to all protected files on a device. Files might be in locations such as emails, messages, apps, and administrative settings. Apply this setting with caution. Possible values are: notConfigured, enabled, disabled.
    systemPolicyAllFiles *Enablement
    // Allow or block access to Desktop folder. Possible values are: notConfigured, enabled, disabled.
    systemPolicyDesktopFolder *Enablement
    // Allow or block access to Documents folder. Possible values are: notConfigured, enabled, disabled.
    systemPolicyDocumentsFolder *Enablement
    // Allow or block access to Downloads folder. Possible values are: notConfigured, enabled, disabled.
    systemPolicyDownloadsFolder *Enablement
    // Allow or block access to network volumes. Requires macOS 10.15 or later. Possible values are: notConfigured, enabled, disabled.
    systemPolicyNetworkVolumes *Enablement
    // Control access to removable  volumes on the device, such as an external hard drive. Requires macOS 10.15 or later. Possible values are: notConfigured, enabled, disabled.
    systemPolicyRemovableVolumes *Enablement
    // Allow app or process to access files used in system administration. Possible values are: notConfigured, enabled, disabled.
    systemPolicySystemAdminFiles *Enablement
}
// NewMacOSPrivacyAccessControlItem instantiates a new macOSPrivacyAccessControlItem and sets the default values.
func NewMacOSPrivacyAccessControlItem()(*MacOSPrivacyAccessControlItem) {
    m := &MacOSPrivacyAccessControlItem{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateMacOSPrivacyAccessControlItemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMacOSPrivacyAccessControlItemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMacOSPrivacyAccessControlItem(), nil
}
// GetAccessibility gets the accessibility property value. Allow the app or process to control the Mac via the Accessibility subsystem. Possible values are: notConfigured, enabled, disabled.
func (m *MacOSPrivacyAccessControlItem) GetAccessibility()(*Enablement) {
    if m == nil {
        return nil
    } else {
        return m.accessibility
    }
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MacOSPrivacyAccessControlItem) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAddressBook gets the addressBook property value. Allow or block access to contact information managed by Contacts. Possible values are: notConfigured, enabled, disabled.
func (m *MacOSPrivacyAccessControlItem) GetAddressBook()(*Enablement) {
    if m == nil {
        return nil
    } else {
        return m.addressBook
    }
}
// GetAppleEventsAllowedReceivers gets the appleEventsAllowedReceivers property value. Allow or deny the app or process to send a restricted Apple event to another app or process. You will need to know the identifier, identifier type, and code requirement of the receiving app or process. This collection can contain a maximum of 500 elements.
func (m *MacOSPrivacyAccessControlItem) GetAppleEventsAllowedReceivers()([]MacOSAppleEventReceiverable) {
    if m == nil {
        return nil
    } else {
        return m.appleEventsAllowedReceivers
    }
}
// GetBlockCamera gets the blockCamera property value. Block access to camera app.
func (m *MacOSPrivacyAccessControlItem) GetBlockCamera()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.blockCamera
    }
}
// GetBlockListenEvent gets the blockListenEvent property value. Block the app or process from listening to events from input devices such as mouse, keyboard, and trackpad.Requires macOS 10.15 or later.
func (m *MacOSPrivacyAccessControlItem) GetBlockListenEvent()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.blockListenEvent
    }
}
// GetBlockMicrophone gets the blockMicrophone property value. Block access to microphone.
func (m *MacOSPrivacyAccessControlItem) GetBlockMicrophone()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.blockMicrophone
    }
}
// GetBlockScreenCapture gets the blockScreenCapture property value. Block app from capturing contents of system display. Requires macOS 10.15 or later.
func (m *MacOSPrivacyAccessControlItem) GetBlockScreenCapture()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.blockScreenCapture
    }
}
// GetCalendar gets the calendar property value. Allow or block access to event information managed by Calendar. Possible values are: notConfigured, enabled, disabled.
func (m *MacOSPrivacyAccessControlItem) GetCalendar()(*Enablement) {
    if m == nil {
        return nil
    } else {
        return m.calendar
    }
}
// GetCodeRequirement gets the codeRequirement property value. Enter the code requirement, which can be obtained with the command 'codesign –display -r –' in the Terminal app. Include everything after '=>'.
func (m *MacOSPrivacyAccessControlItem) GetCodeRequirement()(*string) {
    if m == nil {
        return nil
    } else {
        return m.codeRequirement
    }
}
// GetDisplayName gets the displayName property value. The display name of the app, process, or executable.
func (m *MacOSPrivacyAccessControlItem) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFieldDeserializers the deserialization information for the current model
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
                res[i] = v.(MacOSAppleEventReceiverable)
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
// GetFileProviderPresence gets the fileProviderPresence property value. Allow the app or process to access files managed by another app’s file provider extension. Requires macOS 10.15 or later. . Possible values are: notConfigured, enabled, disabled.
func (m *MacOSPrivacyAccessControlItem) GetFileProviderPresence()(*Enablement) {
    if m == nil {
        return nil
    } else {
        return m.fileProviderPresence
    }
}
// GetIdentifier gets the identifier property value. The bundle ID or path of the app, process, or executable.
func (m *MacOSPrivacyAccessControlItem) GetIdentifier()(*string) {
    if m == nil {
        return nil
    } else {
        return m.identifier
    }
}
// GetIdentifierType gets the identifierType property value. A bundle ID is used to identify an app. A path is used to identify a process or executable. Possible values are: bundleID, path.
func (m *MacOSPrivacyAccessControlItem) GetIdentifierType()(*MacOSProcessIdentifierType) {
    if m == nil {
        return nil
    } else {
        return m.identifierType
    }
}
// GetMediaLibrary gets the mediaLibrary property value. Allow or block access to music and the media library. Possible values are: notConfigured, enabled, disabled.
func (m *MacOSPrivacyAccessControlItem) GetMediaLibrary()(*Enablement) {
    if m == nil {
        return nil
    } else {
        return m.mediaLibrary
    }
}
// GetPhotos gets the photos property value. Allow or block access to images managed by Photos. Possible values are: notConfigured, enabled, disabled.
func (m *MacOSPrivacyAccessControlItem) GetPhotos()(*Enablement) {
    if m == nil {
        return nil
    } else {
        return m.photos
    }
}
// GetPostEvent gets the postEvent property value. Control access to CoreGraphics APIs, which are used to send CGEvents to the system event stream. Possible values are: notConfigured, enabled, disabled.
func (m *MacOSPrivacyAccessControlItem) GetPostEvent()(*Enablement) {
    if m == nil {
        return nil
    } else {
        return m.postEvent
    }
}
// GetReminders gets the reminders property value. Allow or block access to information managed by Reminders. Possible values are: notConfigured, enabled, disabled.
func (m *MacOSPrivacyAccessControlItem) GetReminders()(*Enablement) {
    if m == nil {
        return nil
    } else {
        return m.reminders
    }
}
// GetSpeechRecognition gets the speechRecognition property value. Allow or block access to system speech recognition facility. Possible values are: notConfigured, enabled, disabled.
func (m *MacOSPrivacyAccessControlItem) GetSpeechRecognition()(*Enablement) {
    if m == nil {
        return nil
    } else {
        return m.speechRecognition
    }
}
// GetStaticCodeValidation gets the staticCodeValidation property value. Statically validates the code requirement. Use this setting if the process invalidates its dynamic code signature.
func (m *MacOSPrivacyAccessControlItem) GetStaticCodeValidation()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.staticCodeValidation
    }
}
// GetSystemPolicyAllFiles gets the systemPolicyAllFiles property value. Control access to all protected files on a device. Files might be in locations such as emails, messages, apps, and administrative settings. Apply this setting with caution. Possible values are: notConfigured, enabled, disabled.
func (m *MacOSPrivacyAccessControlItem) GetSystemPolicyAllFiles()(*Enablement) {
    if m == nil {
        return nil
    } else {
        return m.systemPolicyAllFiles
    }
}
// GetSystemPolicyDesktopFolder gets the systemPolicyDesktopFolder property value. Allow or block access to Desktop folder. Possible values are: notConfigured, enabled, disabled.
func (m *MacOSPrivacyAccessControlItem) GetSystemPolicyDesktopFolder()(*Enablement) {
    if m == nil {
        return nil
    } else {
        return m.systemPolicyDesktopFolder
    }
}
// GetSystemPolicyDocumentsFolder gets the systemPolicyDocumentsFolder property value. Allow or block access to Documents folder. Possible values are: notConfigured, enabled, disabled.
func (m *MacOSPrivacyAccessControlItem) GetSystemPolicyDocumentsFolder()(*Enablement) {
    if m == nil {
        return nil
    } else {
        return m.systemPolicyDocumentsFolder
    }
}
// GetSystemPolicyDownloadsFolder gets the systemPolicyDownloadsFolder property value. Allow or block access to Downloads folder. Possible values are: notConfigured, enabled, disabled.
func (m *MacOSPrivacyAccessControlItem) GetSystemPolicyDownloadsFolder()(*Enablement) {
    if m == nil {
        return nil
    } else {
        return m.systemPolicyDownloadsFolder
    }
}
// GetSystemPolicyNetworkVolumes gets the systemPolicyNetworkVolumes property value. Allow or block access to network volumes. Requires macOS 10.15 or later. Possible values are: notConfigured, enabled, disabled.
func (m *MacOSPrivacyAccessControlItem) GetSystemPolicyNetworkVolumes()(*Enablement) {
    if m == nil {
        return nil
    } else {
        return m.systemPolicyNetworkVolumes
    }
}
// GetSystemPolicyRemovableVolumes gets the systemPolicyRemovableVolumes property value. Control access to removable  volumes on the device, such as an external hard drive. Requires macOS 10.15 or later. Possible values are: notConfigured, enabled, disabled.
func (m *MacOSPrivacyAccessControlItem) GetSystemPolicyRemovableVolumes()(*Enablement) {
    if m == nil {
        return nil
    } else {
        return m.systemPolicyRemovableVolumes
    }
}
// GetSystemPolicySystemAdminFiles gets the systemPolicySystemAdminFiles property value. Allow app or process to access files used in system administration. Possible values are: notConfigured, enabled, disabled.
func (m *MacOSPrivacyAccessControlItem) GetSystemPolicySystemAdminFiles()(*Enablement) {
    if m == nil {
        return nil
    } else {
        return m.systemPolicySystemAdminFiles
    }
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
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
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
// SetAccessibility sets the accessibility property value. Allow the app or process to control the Mac via the Accessibility subsystem. Possible values are: notConfigured, enabled, disabled.
func (m *MacOSPrivacyAccessControlItem) SetAccessibility(value *Enablement)() {
    if m != nil {
        m.accessibility = value
    }
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MacOSPrivacyAccessControlItem) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAddressBook sets the addressBook property value. Allow or block access to contact information managed by Contacts. Possible values are: notConfigured, enabled, disabled.
func (m *MacOSPrivacyAccessControlItem) SetAddressBook(value *Enablement)() {
    if m != nil {
        m.addressBook = value
    }
}
// SetAppleEventsAllowedReceivers sets the appleEventsAllowedReceivers property value. Allow or deny the app or process to send a restricted Apple event to another app or process. You will need to know the identifier, identifier type, and code requirement of the receiving app or process. This collection can contain a maximum of 500 elements.
func (m *MacOSPrivacyAccessControlItem) SetAppleEventsAllowedReceivers(value []MacOSAppleEventReceiverable)() {
    if m != nil {
        m.appleEventsAllowedReceivers = value
    }
}
// SetBlockCamera sets the blockCamera property value. Block access to camera app.
func (m *MacOSPrivacyAccessControlItem) SetBlockCamera(value *bool)() {
    if m != nil {
        m.blockCamera = value
    }
}
// SetBlockListenEvent sets the blockListenEvent property value. Block the app or process from listening to events from input devices such as mouse, keyboard, and trackpad.Requires macOS 10.15 or later.
func (m *MacOSPrivacyAccessControlItem) SetBlockListenEvent(value *bool)() {
    if m != nil {
        m.blockListenEvent = value
    }
}
// SetBlockMicrophone sets the blockMicrophone property value. Block access to microphone.
func (m *MacOSPrivacyAccessControlItem) SetBlockMicrophone(value *bool)() {
    if m != nil {
        m.blockMicrophone = value
    }
}
// SetBlockScreenCapture sets the blockScreenCapture property value. Block app from capturing contents of system display. Requires macOS 10.15 or later.
func (m *MacOSPrivacyAccessControlItem) SetBlockScreenCapture(value *bool)() {
    if m != nil {
        m.blockScreenCapture = value
    }
}
// SetCalendar sets the calendar property value. Allow or block access to event information managed by Calendar. Possible values are: notConfigured, enabled, disabled.
func (m *MacOSPrivacyAccessControlItem) SetCalendar(value *Enablement)() {
    if m != nil {
        m.calendar = value
    }
}
// SetCodeRequirement sets the codeRequirement property value. Enter the code requirement, which can be obtained with the command 'codesign –display -r –' in the Terminal app. Include everything after '=>'.
func (m *MacOSPrivacyAccessControlItem) SetCodeRequirement(value *string)() {
    if m != nil {
        m.codeRequirement = value
    }
}
// SetDisplayName sets the displayName property value. The display name of the app, process, or executable.
func (m *MacOSPrivacyAccessControlItem) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetFileProviderPresence sets the fileProviderPresence property value. Allow the app or process to access files managed by another app’s file provider extension. Requires macOS 10.15 or later. . Possible values are: notConfigured, enabled, disabled.
func (m *MacOSPrivacyAccessControlItem) SetFileProviderPresence(value *Enablement)() {
    if m != nil {
        m.fileProviderPresence = value
    }
}
// SetIdentifier sets the identifier property value. The bundle ID or path of the app, process, or executable.
func (m *MacOSPrivacyAccessControlItem) SetIdentifier(value *string)() {
    if m != nil {
        m.identifier = value
    }
}
// SetIdentifierType sets the identifierType property value. A bundle ID is used to identify an app. A path is used to identify a process or executable. Possible values are: bundleID, path.
func (m *MacOSPrivacyAccessControlItem) SetIdentifierType(value *MacOSProcessIdentifierType)() {
    if m != nil {
        m.identifierType = value
    }
}
// SetMediaLibrary sets the mediaLibrary property value. Allow or block access to music and the media library. Possible values are: notConfigured, enabled, disabled.
func (m *MacOSPrivacyAccessControlItem) SetMediaLibrary(value *Enablement)() {
    if m != nil {
        m.mediaLibrary = value
    }
}
// SetPhotos sets the photos property value. Allow or block access to images managed by Photos. Possible values are: notConfigured, enabled, disabled.
func (m *MacOSPrivacyAccessControlItem) SetPhotos(value *Enablement)() {
    if m != nil {
        m.photos = value
    }
}
// SetPostEvent sets the postEvent property value. Control access to CoreGraphics APIs, which are used to send CGEvents to the system event stream. Possible values are: notConfigured, enabled, disabled.
func (m *MacOSPrivacyAccessControlItem) SetPostEvent(value *Enablement)() {
    if m != nil {
        m.postEvent = value
    }
}
// SetReminders sets the reminders property value. Allow or block access to information managed by Reminders. Possible values are: notConfigured, enabled, disabled.
func (m *MacOSPrivacyAccessControlItem) SetReminders(value *Enablement)() {
    if m != nil {
        m.reminders = value
    }
}
// SetSpeechRecognition sets the speechRecognition property value. Allow or block access to system speech recognition facility. Possible values are: notConfigured, enabled, disabled.
func (m *MacOSPrivacyAccessControlItem) SetSpeechRecognition(value *Enablement)() {
    if m != nil {
        m.speechRecognition = value
    }
}
// SetStaticCodeValidation sets the staticCodeValidation property value. Statically validates the code requirement. Use this setting if the process invalidates its dynamic code signature.
func (m *MacOSPrivacyAccessControlItem) SetStaticCodeValidation(value *bool)() {
    if m != nil {
        m.staticCodeValidation = value
    }
}
// SetSystemPolicyAllFiles sets the systemPolicyAllFiles property value. Control access to all protected files on a device. Files might be in locations such as emails, messages, apps, and administrative settings. Apply this setting with caution. Possible values are: notConfigured, enabled, disabled.
func (m *MacOSPrivacyAccessControlItem) SetSystemPolicyAllFiles(value *Enablement)() {
    if m != nil {
        m.systemPolicyAllFiles = value
    }
}
// SetSystemPolicyDesktopFolder sets the systemPolicyDesktopFolder property value. Allow or block access to Desktop folder. Possible values are: notConfigured, enabled, disabled.
func (m *MacOSPrivacyAccessControlItem) SetSystemPolicyDesktopFolder(value *Enablement)() {
    if m != nil {
        m.systemPolicyDesktopFolder = value
    }
}
// SetSystemPolicyDocumentsFolder sets the systemPolicyDocumentsFolder property value. Allow or block access to Documents folder. Possible values are: notConfigured, enabled, disabled.
func (m *MacOSPrivacyAccessControlItem) SetSystemPolicyDocumentsFolder(value *Enablement)() {
    if m != nil {
        m.systemPolicyDocumentsFolder = value
    }
}
// SetSystemPolicyDownloadsFolder sets the systemPolicyDownloadsFolder property value. Allow or block access to Downloads folder. Possible values are: notConfigured, enabled, disabled.
func (m *MacOSPrivacyAccessControlItem) SetSystemPolicyDownloadsFolder(value *Enablement)() {
    if m != nil {
        m.systemPolicyDownloadsFolder = value
    }
}
// SetSystemPolicyNetworkVolumes sets the systemPolicyNetworkVolumes property value. Allow or block access to network volumes. Requires macOS 10.15 or later. Possible values are: notConfigured, enabled, disabled.
func (m *MacOSPrivacyAccessControlItem) SetSystemPolicyNetworkVolumes(value *Enablement)() {
    if m != nil {
        m.systemPolicyNetworkVolumes = value
    }
}
// SetSystemPolicyRemovableVolumes sets the systemPolicyRemovableVolumes property value. Control access to removable  volumes on the device, such as an external hard drive. Requires macOS 10.15 or later. Possible values are: notConfigured, enabled, disabled.
func (m *MacOSPrivacyAccessControlItem) SetSystemPolicyRemovableVolumes(value *Enablement)() {
    if m != nil {
        m.systemPolicyRemovableVolumes = value
    }
}
// SetSystemPolicySystemAdminFiles sets the systemPolicySystemAdminFiles property value. Allow app or process to access files used in system administration. Possible values are: notConfigured, enabled, disabled.
func (m *MacOSPrivacyAccessControlItem) SetSystemPolicySystemAdminFiles(value *Enablement)() {
    if m != nil {
        m.systemPolicySystemAdminFiles = value
    }
}

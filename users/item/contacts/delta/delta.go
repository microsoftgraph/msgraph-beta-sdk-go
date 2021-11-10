package delta

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// 
type Delta struct {
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OutlookItem
    // The name of the contact's assistant.
    assistantName *string;
    // The contact's birthday. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    birthday *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The names of the contact's children.
    children []string;
    // The name of the contact's company.
    companyName *string;
    // The contact's department.
    department *string;
    // The contact's display name. You can specify the display name in a create or update operation. Note that later updates to other properties may cause an automatically generated value to overwrite the displayName value you have specified. To preserve a pre-existing value, always include it as displayName in an update operation.
    displayName *string;
    // The contact's email addresses.
    emailAddresses []i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TypedEmailAddress;
    // The collection of open extensions defined for the contact. Read-only. Nullable.
    extensions []i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Extension;
    // The name the contact is filed under.
    fileAs *string;
    // The flag value that indicates the status, start date, due date, or completion date for the contact.
    flag *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.FollowupFlag;
    // The contact's gender.
    gender *string;
    // The contact's generation.
    generation *string;
    // The contact's given name.
    givenName *string;
    // 
    imAddresses []string;
    // 
    initials *string;
    // 
    isFavorite *bool;
    // 
    jobTitle *string;
    // 
    manager *string;
    // 
    middleName *string;
    // The collection of multi-value extended properties defined for the contact. Read-only. Nullable.
    multiValueExtendedProperties []i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MultiValueLegacyExtendedProperty;
    // 
    nickName *string;
    // 
    officeLocation *string;
    // 
    parentFolderId *string;
    // 
    personalNotes *string;
    // 
    phones []i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Phone;
    // Optional contact picture. You can get or set a photo for a contact.
    photo *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ProfilePhoto;
    // 
    postalAddresses []i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PhysicalAddress;
    // 
    profession *string;
    // The collection of single-value extended properties defined for the contact. Read-only. Nullable.
    singleValueExtendedProperties []i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SingleValueLegacyExtendedProperty;
    // 
    spouseName *string;
    // 
    surname *string;
    // 
    title *string;
    // 
    websites []i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Website;
    // 
    weddingAnniversary *string;
    // 
    yomiCompanyName *string;
    // 
    yomiGivenName *string;
    // 
    yomiSurname *string;
}
// Instantiates a new delta and sets the default values.
func NewDelta()(*Delta) {
    m := &Delta{
        OutlookItem: *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewOutlookItem(),
    }
    return m
}
// Gets the assistantName property value. The name of the contact's assistant.
func (m *Delta) GetAssistantName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.assistantName
    }
}
// Gets the birthday property value. The contact's birthday. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *Delta) GetBirthday()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.birthday
    }
}
// Gets the children property value. The names of the contact's children.
func (m *Delta) GetChildren()([]string) {
    if m == nil {
        return nil
    } else {
        return m.children
    }
}
// Gets the companyName property value. The name of the contact's company.
func (m *Delta) GetCompanyName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.companyName
    }
}
// Gets the department property value. The contact's department.
func (m *Delta) GetDepartment()(*string) {
    if m == nil {
        return nil
    } else {
        return m.department
    }
}
// Gets the displayName property value. The contact's display name. You can specify the display name in a create or update operation. Note that later updates to other properties may cause an automatically generated value to overwrite the displayName value you have specified. To preserve a pre-existing value, always include it as displayName in an update operation.
func (m *Delta) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the emailAddresses property value. The contact's email addresses.
func (m *Delta) GetEmailAddresses()([]i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TypedEmailAddress) {
    if m == nil {
        return nil
    } else {
        return m.emailAddresses
    }
}
// Gets the extensions property value. The collection of open extensions defined for the contact. Read-only. Nullable.
func (m *Delta) GetExtensions()([]i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Extension) {
    if m == nil {
        return nil
    } else {
        return m.extensions
    }
}
// Gets the fileAs property value. The name the contact is filed under.
func (m *Delta) GetFileAs()(*string) {
    if m == nil {
        return nil
    } else {
        return m.fileAs
    }
}
// Gets the flag property value. The flag value that indicates the status, start date, due date, or completion date for the contact.
func (m *Delta) GetFlag()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.FollowupFlag) {
    if m == nil {
        return nil
    } else {
        return m.flag
    }
}
// Gets the gender property value. The contact's gender.
func (m *Delta) GetGender()(*string) {
    if m == nil {
        return nil
    } else {
        return m.gender
    }
}
// Gets the generation property value. The contact's generation.
func (m *Delta) GetGeneration()(*string) {
    if m == nil {
        return nil
    } else {
        return m.generation
    }
}
// Gets the givenName property value. The contact's given name.
func (m *Delta) GetGivenName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.givenName
    }
}
// Gets the imAddresses property value. 
func (m *Delta) GetImAddresses()([]string) {
    if m == nil {
        return nil
    } else {
        return m.imAddresses
    }
}
// Gets the initials property value. 
func (m *Delta) GetInitials()(*string) {
    if m == nil {
        return nil
    } else {
        return m.initials
    }
}
// Gets the isFavorite property value. 
func (m *Delta) GetIsFavorite()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isFavorite
    }
}
// Gets the jobTitle property value. 
func (m *Delta) GetJobTitle()(*string) {
    if m == nil {
        return nil
    } else {
        return m.jobTitle
    }
}
// Gets the manager property value. 
func (m *Delta) GetManager()(*string) {
    if m == nil {
        return nil
    } else {
        return m.manager
    }
}
// Gets the middleName property value. 
func (m *Delta) GetMiddleName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.middleName
    }
}
// Gets the multiValueExtendedProperties property value. The collection of multi-value extended properties defined for the contact. Read-only. Nullable.
func (m *Delta) GetMultiValueExtendedProperties()([]i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MultiValueLegacyExtendedProperty) {
    if m == nil {
        return nil
    } else {
        return m.multiValueExtendedProperties
    }
}
// Gets the nickName property value. 
func (m *Delta) GetNickName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.nickName
    }
}
// Gets the officeLocation property value. 
func (m *Delta) GetOfficeLocation()(*string) {
    if m == nil {
        return nil
    } else {
        return m.officeLocation
    }
}
// Gets the parentFolderId property value. 
func (m *Delta) GetParentFolderId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.parentFolderId
    }
}
// Gets the personalNotes property value. 
func (m *Delta) GetPersonalNotes()(*string) {
    if m == nil {
        return nil
    } else {
        return m.personalNotes
    }
}
// Gets the phones property value. 
func (m *Delta) GetPhones()([]i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Phone) {
    if m == nil {
        return nil
    } else {
        return m.phones
    }
}
// Gets the photo property value. Optional contact picture. You can get or set a photo for a contact.
func (m *Delta) GetPhoto()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ProfilePhoto) {
    if m == nil {
        return nil
    } else {
        return m.photo
    }
}
// Gets the postalAddresses property value. 
func (m *Delta) GetPostalAddresses()([]i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PhysicalAddress) {
    if m == nil {
        return nil
    } else {
        return m.postalAddresses
    }
}
// Gets the profession property value. 
func (m *Delta) GetProfession()(*string) {
    if m == nil {
        return nil
    } else {
        return m.profession
    }
}
// Gets the singleValueExtendedProperties property value. The collection of single-value extended properties defined for the contact. Read-only. Nullable.
func (m *Delta) GetSingleValueExtendedProperties()([]i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SingleValueLegacyExtendedProperty) {
    if m == nil {
        return nil
    } else {
        return m.singleValueExtendedProperties
    }
}
// Gets the spouseName property value. 
func (m *Delta) GetSpouseName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.spouseName
    }
}
// Gets the surname property value. 
func (m *Delta) GetSurname()(*string) {
    if m == nil {
        return nil
    } else {
        return m.surname
    }
}
// Gets the title property value. 
func (m *Delta) GetTitle()(*string) {
    if m == nil {
        return nil
    } else {
        return m.title
    }
}
// Gets the websites property value. 
func (m *Delta) GetWebsites()([]i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Website) {
    if m == nil {
        return nil
    } else {
        return m.websites
    }
}
// Gets the weddingAnniversary property value. 
func (m *Delta) GetWeddingAnniversary()(*string) {
    if m == nil {
        return nil
    } else {
        return m.weddingAnniversary
    }
}
// Gets the yomiCompanyName property value. 
func (m *Delta) GetYomiCompanyName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.yomiCompanyName
    }
}
// Gets the yomiGivenName property value. 
func (m *Delta) GetYomiGivenName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.yomiGivenName
    }
}
// Gets the yomiSurname property value. 
func (m *Delta) GetYomiSurname()(*string) {
    if m == nil {
        return nil
    } else {
        return m.yomiSurname
    }
}
// The deserialization information for the current model
func (m *Delta) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.OutlookItem.GetFieldDeserializers()
    res["assistantName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssistantName(val)
        }
        return nil
    }
    res["birthday"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBirthday(val)
        }
        return nil
    }
    res["children"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetChildren(res)
        }
        return nil
    }
    res["companyName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompanyName(val)
        }
        return nil
    }
    res["department"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDepartment(val)
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["emailAddresses"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewTypedEmailAddress() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TypedEmailAddress, len(val))
            for i, v := range val {
                res[i] = *(v.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TypedEmailAddress))
            }
            m.SetEmailAddresses(res)
        }
        return nil
    }
    res["extensions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewExtension() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Extension, len(val))
            for i, v := range val {
                res[i] = *(v.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Extension))
            }
            m.SetExtensions(res)
        }
        return nil
    }
    res["fileAs"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFileAs(val)
        }
        return nil
    }
    res["flag"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewFollowupFlag() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFlag(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.FollowupFlag))
        }
        return nil
    }
    res["gender"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGender(val)
        }
        return nil
    }
    res["generation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGeneration(val)
        }
        return nil
    }
    res["givenName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGivenName(val)
        }
        return nil
    }
    res["imAddresses"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetImAddresses(res)
        }
        return nil
    }
    res["initials"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInitials(val)
        }
        return nil
    }
    res["isFavorite"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsFavorite(val)
        }
        return nil
    }
    res["jobTitle"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJobTitle(val)
        }
        return nil
    }
    res["manager"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManager(val)
        }
        return nil
    }
    res["middleName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMiddleName(val)
        }
        return nil
    }
    res["multiValueExtendedProperties"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewMultiValueLegacyExtendedProperty() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MultiValueLegacyExtendedProperty, len(val))
            for i, v := range val {
                res[i] = *(v.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MultiValueLegacyExtendedProperty))
            }
            m.SetMultiValueExtendedProperties(res)
        }
        return nil
    }
    res["nickName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNickName(val)
        }
        return nil
    }
    res["officeLocation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOfficeLocation(val)
        }
        return nil
    }
    res["parentFolderId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParentFolderId(val)
        }
        return nil
    }
    res["personalNotes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPersonalNotes(val)
        }
        return nil
    }
    res["phones"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewPhone() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Phone, len(val))
            for i, v := range val {
                res[i] = *(v.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Phone))
            }
            m.SetPhones(res)
        }
        return nil
    }
    res["photo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewProfilePhoto() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPhoto(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ProfilePhoto))
        }
        return nil
    }
    res["postalAddresses"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewPhysicalAddress() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PhysicalAddress, len(val))
            for i, v := range val {
                res[i] = *(v.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PhysicalAddress))
            }
            m.SetPostalAddresses(res)
        }
        return nil
    }
    res["profession"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProfession(val)
        }
        return nil
    }
    res["singleValueExtendedProperties"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewSingleValueLegacyExtendedProperty() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SingleValueLegacyExtendedProperty, len(val))
            for i, v := range val {
                res[i] = *(v.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SingleValueLegacyExtendedProperty))
            }
            m.SetSingleValueExtendedProperties(res)
        }
        return nil
    }
    res["spouseName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSpouseName(val)
        }
        return nil
    }
    res["surname"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSurname(val)
        }
        return nil
    }
    res["title"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTitle(val)
        }
        return nil
    }
    res["websites"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewWebsite() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Website, len(val))
            for i, v := range val {
                res[i] = *(v.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Website))
            }
            m.SetWebsites(res)
        }
        return nil
    }
    res["weddingAnniversary"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWeddingAnniversary(val)
        }
        return nil
    }
    res["yomiCompanyName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetYomiCompanyName(val)
        }
        return nil
    }
    res["yomiGivenName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetYomiGivenName(val)
        }
        return nil
    }
    res["yomiSurname"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetYomiSurname(val)
        }
        return nil
    }
    return res
}
func (m *Delta) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *Delta) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.OutlookItem.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("assistantName", m.GetAssistantName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("birthday", m.GetBirthday())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("children", m.GetChildren())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("companyName", m.GetCompanyName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("department", m.GetDepartment())
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetEmailAddresses()))
        for i, v := range m.GetEmailAddresses() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("emailAddresses", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetExtensions()))
        for i, v := range m.GetExtensions() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("extensions", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("fileAs", m.GetFileAs())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("flag", m.GetFlag())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("gender", m.GetGender())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("generation", m.GetGeneration())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("givenName", m.GetGivenName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("imAddresses", m.GetImAddresses())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("initials", m.GetInitials())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isFavorite", m.GetIsFavorite())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("jobTitle", m.GetJobTitle())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("manager", m.GetManager())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("middleName", m.GetMiddleName())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMultiValueExtendedProperties()))
        for i, v := range m.GetMultiValueExtendedProperties() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("multiValueExtendedProperties", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("nickName", m.GetNickName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("officeLocation", m.GetOfficeLocation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("parentFolderId", m.GetParentFolderId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("personalNotes", m.GetPersonalNotes())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetPhones()))
        for i, v := range m.GetPhones() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("phones", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("photo", m.GetPhoto())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetPostalAddresses()))
        for i, v := range m.GetPostalAddresses() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("postalAddresses", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("profession", m.GetProfession())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSingleValueExtendedProperties()))
        for i, v := range m.GetSingleValueExtendedProperties() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("singleValueExtendedProperties", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("spouseName", m.GetSpouseName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("surname", m.GetSurname())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("title", m.GetTitle())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetWebsites()))
        for i, v := range m.GetWebsites() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("websites", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("weddingAnniversary", m.GetWeddingAnniversary())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("yomiCompanyName", m.GetYomiCompanyName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("yomiGivenName", m.GetYomiGivenName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("yomiSurname", m.GetYomiSurname())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the assistantName property value. The name of the contact's assistant.
// Parameters:
//  - value : Value to set for the assistantName property.
func (m *Delta) SetAssistantName(value *string)() {
    m.assistantName = value
}
// Sets the birthday property value. The contact's birthday. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
// Parameters:
//  - value : Value to set for the birthday property.
func (m *Delta) SetBirthday(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.birthday = value
}
// Sets the children property value. The names of the contact's children.
// Parameters:
//  - value : Value to set for the children property.
func (m *Delta) SetChildren(value []string)() {
    m.children = value
}
// Sets the companyName property value. The name of the contact's company.
// Parameters:
//  - value : Value to set for the companyName property.
func (m *Delta) SetCompanyName(value *string)() {
    m.companyName = value
}
// Sets the department property value. The contact's department.
// Parameters:
//  - value : Value to set for the department property.
func (m *Delta) SetDepartment(value *string)() {
    m.department = value
}
// Sets the displayName property value. The contact's display name. You can specify the display name in a create or update operation. Note that later updates to other properties may cause an automatically generated value to overwrite the displayName value you have specified. To preserve a pre-existing value, always include it as displayName in an update operation.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *Delta) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the emailAddresses property value. The contact's email addresses.
// Parameters:
//  - value : Value to set for the emailAddresses property.
func (m *Delta) SetEmailAddresses(value []i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TypedEmailAddress)() {
    m.emailAddresses = value
}
// Sets the extensions property value. The collection of open extensions defined for the contact. Read-only. Nullable.
// Parameters:
//  - value : Value to set for the extensions property.
func (m *Delta) SetExtensions(value []i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Extension)() {
    m.extensions = value
}
// Sets the fileAs property value. The name the contact is filed under.
// Parameters:
//  - value : Value to set for the fileAs property.
func (m *Delta) SetFileAs(value *string)() {
    m.fileAs = value
}
// Sets the flag property value. The flag value that indicates the status, start date, due date, or completion date for the contact.
// Parameters:
//  - value : Value to set for the flag property.
func (m *Delta) SetFlag(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.FollowupFlag)() {
    m.flag = value
}
// Sets the gender property value. The contact's gender.
// Parameters:
//  - value : Value to set for the gender property.
func (m *Delta) SetGender(value *string)() {
    m.gender = value
}
// Sets the generation property value. The contact's generation.
// Parameters:
//  - value : Value to set for the generation property.
func (m *Delta) SetGeneration(value *string)() {
    m.generation = value
}
// Sets the givenName property value. The contact's given name.
// Parameters:
//  - value : Value to set for the givenName property.
func (m *Delta) SetGivenName(value *string)() {
    m.givenName = value
}
// Sets the imAddresses property value. 
// Parameters:
//  - value : Value to set for the imAddresses property.
func (m *Delta) SetImAddresses(value []string)() {
    m.imAddresses = value
}
// Sets the initials property value. 
// Parameters:
//  - value : Value to set for the initials property.
func (m *Delta) SetInitials(value *string)() {
    m.initials = value
}
// Sets the isFavorite property value. 
// Parameters:
//  - value : Value to set for the isFavorite property.
func (m *Delta) SetIsFavorite(value *bool)() {
    m.isFavorite = value
}
// Sets the jobTitle property value. 
// Parameters:
//  - value : Value to set for the jobTitle property.
func (m *Delta) SetJobTitle(value *string)() {
    m.jobTitle = value
}
// Sets the manager property value. 
// Parameters:
//  - value : Value to set for the manager property.
func (m *Delta) SetManager(value *string)() {
    m.manager = value
}
// Sets the middleName property value. 
// Parameters:
//  - value : Value to set for the middleName property.
func (m *Delta) SetMiddleName(value *string)() {
    m.middleName = value
}
// Sets the multiValueExtendedProperties property value. The collection of multi-value extended properties defined for the contact. Read-only. Nullable.
// Parameters:
//  - value : Value to set for the multiValueExtendedProperties property.
func (m *Delta) SetMultiValueExtendedProperties(value []i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MultiValueLegacyExtendedProperty)() {
    m.multiValueExtendedProperties = value
}
// Sets the nickName property value. 
// Parameters:
//  - value : Value to set for the nickName property.
func (m *Delta) SetNickName(value *string)() {
    m.nickName = value
}
// Sets the officeLocation property value. 
// Parameters:
//  - value : Value to set for the officeLocation property.
func (m *Delta) SetOfficeLocation(value *string)() {
    m.officeLocation = value
}
// Sets the parentFolderId property value. 
// Parameters:
//  - value : Value to set for the parentFolderId property.
func (m *Delta) SetParentFolderId(value *string)() {
    m.parentFolderId = value
}
// Sets the personalNotes property value. 
// Parameters:
//  - value : Value to set for the personalNotes property.
func (m *Delta) SetPersonalNotes(value *string)() {
    m.personalNotes = value
}
// Sets the phones property value. 
// Parameters:
//  - value : Value to set for the phones property.
func (m *Delta) SetPhones(value []i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Phone)() {
    m.phones = value
}
// Sets the photo property value. Optional contact picture. You can get or set a photo for a contact.
// Parameters:
//  - value : Value to set for the photo property.
func (m *Delta) SetPhoto(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ProfilePhoto)() {
    m.photo = value
}
// Sets the postalAddresses property value. 
// Parameters:
//  - value : Value to set for the postalAddresses property.
func (m *Delta) SetPostalAddresses(value []i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PhysicalAddress)() {
    m.postalAddresses = value
}
// Sets the profession property value. 
// Parameters:
//  - value : Value to set for the profession property.
func (m *Delta) SetProfession(value *string)() {
    m.profession = value
}
// Sets the singleValueExtendedProperties property value. The collection of single-value extended properties defined for the contact. Read-only. Nullable.
// Parameters:
//  - value : Value to set for the singleValueExtendedProperties property.
func (m *Delta) SetSingleValueExtendedProperties(value []i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SingleValueLegacyExtendedProperty)() {
    m.singleValueExtendedProperties = value
}
// Sets the spouseName property value. 
// Parameters:
//  - value : Value to set for the spouseName property.
func (m *Delta) SetSpouseName(value *string)() {
    m.spouseName = value
}
// Sets the surname property value. 
// Parameters:
//  - value : Value to set for the surname property.
func (m *Delta) SetSurname(value *string)() {
    m.surname = value
}
// Sets the title property value. 
// Parameters:
//  - value : Value to set for the title property.
func (m *Delta) SetTitle(value *string)() {
    m.title = value
}
// Sets the websites property value. 
// Parameters:
//  - value : Value to set for the websites property.
func (m *Delta) SetWebsites(value []i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Website)() {
    m.websites = value
}
// Sets the weddingAnniversary property value. 
// Parameters:
//  - value : Value to set for the weddingAnniversary property.
func (m *Delta) SetWeddingAnniversary(value *string)() {
    m.weddingAnniversary = value
}
// Sets the yomiCompanyName property value. 
// Parameters:
//  - value : Value to set for the yomiCompanyName property.
func (m *Delta) SetYomiCompanyName(value *string)() {
    m.yomiCompanyName = value
}
// Sets the yomiGivenName property value. 
// Parameters:
//  - value : Value to set for the yomiGivenName property.
func (m *Delta) SetYomiGivenName(value *string)() {
    m.yomiGivenName = value
}
// Sets the yomiSurname property value. 
// Parameters:
//  - value : Value to set for the yomiSurname property.
func (m *Delta) SetYomiSurname(value *string)() {
    m.yomiSurname = value
}

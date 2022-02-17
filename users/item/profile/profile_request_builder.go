package profile

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i26a7f8ef2723df7b554c727d88e4c0fade18dfb137078fee4bff0aaadde79d6b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/profile/skills"
    i376a1c9fb041f9485ffc4b05d361c22c731bfc23b1f7d7bd1ee1ccab102bbec4 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/profile/websites"
    i3fa49a978d179839dbea3f8dd172662623ea0846568dc802484468eb4f5b89d8 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/profile/notes"
    i46891830191b8b241de7376dce57ac07cfd53ae6fd0013bd8e1e973a36e35d0a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/profile/interests"
    i54063d6549cfffd18152df5918e18636c36e95ba53e40f0a76aab9477de26aa6 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/profile/phones"
    i624b8f0e161cffd6542ffd795e358c6db683b8120f9e592078e0ed9f8fed0eca "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/profile/patents"
    i66a493024db3c54c6d1d0a7a037cc42558a6cc37204f250864539ee332b7e218 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/profile/emails"
    i6e5ffca00124afb0244dcbd530449478920e5af0605e6c2b4ffc6742b991ee48 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/profile/awards"
    i7d8235f2ef58f14d9b48524d25c0ea6e65bf8d9e92bb04a40fe31ae1f66a3908 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/profile/webaccounts"
    i8708d000e3379870d86078c6f35b29d11715cdfd03b38c7765791918f0672600 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/profile/addresses"
    i8e8af53beda11d49c2fe34af67e33e01fbfc8215ef688bcc9a6a99b22213c975 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/profile/names"
    i8f9bf67bcb3cda9c5f1048d104da873d4944ff13d895646ec920d2c76bfff28d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/profile/anniversaries"
    ia594d9f0fde4f12633b29cf2184447c01329657b9029df63e316ec608c382ce7 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/profile/projects"
    id61a36d1ca39b427b83fcc2c83fef77cca045a5aab72b1488d7f99f1260ac02f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/profile/certifications"
    id664db371311b8f8e8c63b1bcd1e9e389320ce441b0df7fd171ae3e2183e3ef2 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/profile/educationalactivities"
    id9a14f7d08a6d2f477b59c98cd195160e5cc687372332cdcac1fb9a6c27a6f32 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/profile/account"
    idb2fd763c8ceac7682affef82cbd0bac0889f63a9904631022f51d8f81d00202 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/profile/publications"
    idcb0b21c16e8bb5245cebf25791c36ae69b0d39c1f2f0e8defa8dd5c56a64600 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/profile/positions"
    ie07f4e96ea74d56fa71e6fc9c3cf4c8a7aa0b94b884c31955e02827215f4b842 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/profile/languages"
    i042c646cb6e32821c8fad98a8f86ea3810ba71e5ce5f207b7727d32a421aa35d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/profile/websites/item"
    i09dbdbf1122b782cd8c1844df30cff1030466125501cfa48f39bace4667a50ab "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/profile/phones/item"
    i0c6587bb49b5339ae5004ef05df87a4393eea15a87469e8b9347e8ed7ad9f098 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/profile/skills/item"
    i1447f59bce732e48357e480ddb4ec0ee411855640d3f9e1da6ed7fc9d4086d64 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/profile/names/item"
    i1942a76b21bedb363b344695ae63b964a7c1e7f8cbee2dd8a053356f4cdea025 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/profile/account/item"
    i1d6a22e1958f8e610c82479a0a2d9fbfbe737522bed6b1ae128aed42eec5de72 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/profile/interests/item"
    i42cff4ee1200327093e5e7ebc241155dc70d33e9a1329d8d67d8fb56f8b951a9 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/profile/notes/item"
    i42f9ec73b968bdabb69618c8eb1906cb3c1655e5c30183b4a20aae652614b827 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/profile/educationalactivities/item"
    i4701b652703be8fdd84b7e4c9738fbe17c3e8ba6dd477ae00e44ab3e65041138 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/profile/certifications/item"
    i5a184e11a95aa2344b38024a42ff6a91fc4849920be15fb51f531e12f69b1ca6 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/profile/webaccounts/item"
    i6b52d42e9fa352b8843a815622d380fd88fb3108ec3297f67307e922fe108110 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/profile/awards/item"
    i6cb42d6092f89243d60ec2eff66d3051d0f316e72ca67a0846be1031eb378549 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/profile/emails/item"
    i770fc10a67210474d14ec729151f910b743e9684eb17d1374540e9da62dc6f85 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/profile/anniversaries/item"
    i8450d9ff8966cf928b8991d8901a612103583e7b8300bc32e543a25ab8c60952 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/profile/positions/item"
    i86cafdce36cb3eb513134ac90aeb1f13b19256a2d7e03fc8eb4d029c011328e7 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/profile/addresses/item"
    i91b663b48ef705f87f8309ddf035d579ff8c2c2b8eb0689c1eff786c8be74538 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/profile/patents/item"
    ib34c8a74c1ba3d0d92a6d260bd9015795a3b2c70dbaef9ad4c485e831646e025 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/profile/projects/item"
    ib3dd68449e016ac77f213ca43faab02d165f0dd15c9fbbb56bc0094df6270197 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/profile/languages/item"
    iea021cd498fe18da6c27911f34b60740750e87aa77d2bc8965521a0e5f36c311 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/profile/publications/item"
)

// ProfileRequestBuilder builds and executes requests for operations under \users\{user-id}\profile
type ProfileRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ProfileRequestBuilderDeleteOptions options for Delete
type ProfileRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ProfileRequestBuilderGetOptions options for Get
type ProfileRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ProfileRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ProfileRequestBuilderGetQueryParameters represents properties that are descriptive of a user in a tenant.
type ProfileRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// ProfileRequestBuilderPatchOptions options for Patch
type ProfileRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Profile;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *ProfileRequestBuilder) Account()(*id9a14f7d08a6d2f477b59c98cd195160e5cc687372332cdcac1fb9a6c27a6f32.AccountRequestBuilder) {
    return id9a14f7d08a6d2f477b59c98cd195160e5cc687372332cdcac1fb9a6c27a6f32.NewAccountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccountById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.profile.account.item collection
func (m *ProfileRequestBuilder) AccountById(id string)(*i1942a76b21bedb363b344695ae63b964a7c1e7f8cbee2dd8a053356f4cdea025.UserAccountInformationRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userAccountInformation_id"] = id
    }
    return i1942a76b21bedb363b344695ae63b964a7c1e7f8cbee2dd8a053356f4cdea025.NewUserAccountInformationRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ProfileRequestBuilder) Addresses()(*i8708d000e3379870d86078c6f35b29d11715cdfd03b38c7765791918f0672600.AddressesRequestBuilder) {
    return i8708d000e3379870d86078c6f35b29d11715cdfd03b38c7765791918f0672600.NewAddressesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AddressesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.profile.addresses.item collection
func (m *ProfileRequestBuilder) AddressesById(id string)(*i86cafdce36cb3eb513134ac90aeb1f13b19256a2d7e03fc8eb4d029c011328e7.ItemAddressRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemAddress_id"] = id
    }
    return i86cafdce36cb3eb513134ac90aeb1f13b19256a2d7e03fc8eb4d029c011328e7.NewItemAddressRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ProfileRequestBuilder) Anniversaries()(*i8f9bf67bcb3cda9c5f1048d104da873d4944ff13d895646ec920d2c76bfff28d.AnniversariesRequestBuilder) {
    return i8f9bf67bcb3cda9c5f1048d104da873d4944ff13d895646ec920d2c76bfff28d.NewAnniversariesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AnniversariesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.profile.anniversaries.item collection
func (m *ProfileRequestBuilder) AnniversariesById(id string)(*i770fc10a67210474d14ec729151f910b743e9684eb17d1374540e9da62dc6f85.PersonAnnualEventRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["personAnnualEvent_id"] = id
    }
    return i770fc10a67210474d14ec729151f910b743e9684eb17d1374540e9da62dc6f85.NewPersonAnnualEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ProfileRequestBuilder) Awards()(*i6e5ffca00124afb0244dcbd530449478920e5af0605e6c2b4ffc6742b991ee48.AwardsRequestBuilder) {
    return i6e5ffca00124afb0244dcbd530449478920e5af0605e6c2b4ffc6742b991ee48.NewAwardsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AwardsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.profile.awards.item collection
func (m *ProfileRequestBuilder) AwardsById(id string)(*i6b52d42e9fa352b8843a815622d380fd88fb3108ec3297f67307e922fe108110.PersonAwardRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["personAward_id"] = id
    }
    return i6b52d42e9fa352b8843a815622d380fd88fb3108ec3297f67307e922fe108110.NewPersonAwardRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ProfileRequestBuilder) Certifications()(*id61a36d1ca39b427b83fcc2c83fef77cca045a5aab72b1488d7f99f1260ac02f.CertificationsRequestBuilder) {
    return id61a36d1ca39b427b83fcc2c83fef77cca045a5aab72b1488d7f99f1260ac02f.NewCertificationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CertificationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.profile.certifications.item collection
func (m *ProfileRequestBuilder) CertificationsById(id string)(*i4701b652703be8fdd84b7e4c9738fbe17c3e8ba6dd477ae00e44ab3e65041138.PersonCertificationRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["personCertification_id"] = id
    }
    return i4701b652703be8fdd84b7e4c9738fbe17c3e8ba6dd477ae00e44ab3e65041138.NewPersonCertificationRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewProfileRequestBuilderInternal instantiates a new ProfileRequestBuilder and sets the default values.
func NewProfileRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ProfileRequestBuilder) {
    m := &ProfileRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/profile{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewProfileRequestBuilder instantiates a new ProfileRequestBuilder and sets the default values.
func NewProfileRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ProfileRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewProfileRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation represents properties that are descriptive of a user in a tenant.
func (m *ProfileRequestBuilder) CreateDeleteRequestInformation(options *ProfileRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation represents properties that are descriptive of a user in a tenant.
func (m *ProfileRequestBuilder) CreateGetRequestInformation(options *ProfileRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
    }
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation represents properties that are descriptive of a user in a tenant.
func (m *ProfileRequestBuilder) CreatePatchRequestInformation(options *ProfileRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete represents properties that are descriptive of a user in a tenant.
func (m *ProfileRequestBuilder) Delete(options *ProfileRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *ProfileRequestBuilder) EducationalActivities()(*id664db371311b8f8e8c63b1bcd1e9e389320ce441b0df7fd171ae3e2183e3ef2.EducationalActivitiesRequestBuilder) {
    return id664db371311b8f8e8c63b1bcd1e9e389320ce441b0df7fd171ae3e2183e3ef2.NewEducationalActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EducationalActivitiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.profile.educationalActivities.item collection
func (m *ProfileRequestBuilder) EducationalActivitiesById(id string)(*i42f9ec73b968bdabb69618c8eb1906cb3c1655e5c30183b4a20aae652614b827.EducationalActivityRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationalActivity_id"] = id
    }
    return i42f9ec73b968bdabb69618c8eb1906cb3c1655e5c30183b4a20aae652614b827.NewEducationalActivityRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ProfileRequestBuilder) Emails()(*i66a493024db3c54c6d1d0a7a037cc42558a6cc37204f250864539ee332b7e218.EmailsRequestBuilder) {
    return i66a493024db3c54c6d1d0a7a037cc42558a6cc37204f250864539ee332b7e218.NewEmailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EmailsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.profile.emails.item collection
func (m *ProfileRequestBuilder) EmailsById(id string)(*i6cb42d6092f89243d60ec2eff66d3051d0f316e72ca67a0846be1031eb378549.ItemEmailRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemEmail_id"] = id
    }
    return i6cb42d6092f89243d60ec2eff66d3051d0f316e72ca67a0846be1031eb378549.NewItemEmailRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get represents properties that are descriptive of a user in a tenant.
func (m *ProfileRequestBuilder) Get(options *ProfileRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Profile, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewProfile() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Profile), nil
}
func (m *ProfileRequestBuilder) Interests()(*i46891830191b8b241de7376dce57ac07cfd53ae6fd0013bd8e1e973a36e35d0a.InterestsRequestBuilder) {
    return i46891830191b8b241de7376dce57ac07cfd53ae6fd0013bd8e1e973a36e35d0a.NewInterestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InterestsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.profile.interests.item collection
func (m *ProfileRequestBuilder) InterestsById(id string)(*i1d6a22e1958f8e610c82479a0a2d9fbfbe737522bed6b1ae128aed42eec5de72.PersonInterestRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["personInterest_id"] = id
    }
    return i1d6a22e1958f8e610c82479a0a2d9fbfbe737522bed6b1ae128aed42eec5de72.NewPersonInterestRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ProfileRequestBuilder) Languages()(*ie07f4e96ea74d56fa71e6fc9c3cf4c8a7aa0b94b884c31955e02827215f4b842.LanguagesRequestBuilder) {
    return ie07f4e96ea74d56fa71e6fc9c3cf4c8a7aa0b94b884c31955e02827215f4b842.NewLanguagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// LanguagesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.profile.languages.item collection
func (m *ProfileRequestBuilder) LanguagesById(id string)(*ib3dd68449e016ac77f213ca43faab02d165f0dd15c9fbbb56bc0094df6270197.LanguageProficiencyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["languageProficiency_id"] = id
    }
    return ib3dd68449e016ac77f213ca43faab02d165f0dd15c9fbbb56bc0094df6270197.NewLanguageProficiencyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ProfileRequestBuilder) Names()(*i8e8af53beda11d49c2fe34af67e33e01fbfc8215ef688bcc9a6a99b22213c975.NamesRequestBuilder) {
    return i8e8af53beda11d49c2fe34af67e33e01fbfc8215ef688bcc9a6a99b22213c975.NewNamesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NamesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.profile.names.item collection
func (m *ProfileRequestBuilder) NamesById(id string)(*i1447f59bce732e48357e480ddb4ec0ee411855640d3f9e1da6ed7fc9d4086d64.PersonNameRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["personName_id"] = id
    }
    return i1447f59bce732e48357e480ddb4ec0ee411855640d3f9e1da6ed7fc9d4086d64.NewPersonNameRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ProfileRequestBuilder) Notes()(*i3fa49a978d179839dbea3f8dd172662623ea0846568dc802484468eb4f5b89d8.NotesRequestBuilder) {
    return i3fa49a978d179839dbea3f8dd172662623ea0846568dc802484468eb4f5b89d8.NewNotesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NotesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.profile.notes.item collection
func (m *ProfileRequestBuilder) NotesById(id string)(*i42cff4ee1200327093e5e7ebc241155dc70d33e9a1329d8d67d8fb56f8b951a9.PersonAnnotationRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["personAnnotation_id"] = id
    }
    return i42cff4ee1200327093e5e7ebc241155dc70d33e9a1329d8d67d8fb56f8b951a9.NewPersonAnnotationRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch represents properties that are descriptive of a user in a tenant.
func (m *ProfileRequestBuilder) Patch(options *ProfileRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *ProfileRequestBuilder) Patents()(*i624b8f0e161cffd6542ffd795e358c6db683b8120f9e592078e0ed9f8fed0eca.PatentsRequestBuilder) {
    return i624b8f0e161cffd6542ffd795e358c6db683b8120f9e592078e0ed9f8fed0eca.NewPatentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PatentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.profile.patents.item collection
func (m *ProfileRequestBuilder) PatentsById(id string)(*i91b663b48ef705f87f8309ddf035d579ff8c2c2b8eb0689c1eff786c8be74538.ItemPatentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemPatent_id"] = id
    }
    return i91b663b48ef705f87f8309ddf035d579ff8c2c2b8eb0689c1eff786c8be74538.NewItemPatentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ProfileRequestBuilder) Phones()(*i54063d6549cfffd18152df5918e18636c36e95ba53e40f0a76aab9477de26aa6.PhonesRequestBuilder) {
    return i54063d6549cfffd18152df5918e18636c36e95ba53e40f0a76aab9477de26aa6.NewPhonesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PhonesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.profile.phones.item collection
func (m *ProfileRequestBuilder) PhonesById(id string)(*i09dbdbf1122b782cd8c1844df30cff1030466125501cfa48f39bace4667a50ab.ItemPhoneRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemPhone_id"] = id
    }
    return i09dbdbf1122b782cd8c1844df30cff1030466125501cfa48f39bace4667a50ab.NewItemPhoneRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ProfileRequestBuilder) Positions()(*idcb0b21c16e8bb5245cebf25791c36ae69b0d39c1f2f0e8defa8dd5c56a64600.PositionsRequestBuilder) {
    return idcb0b21c16e8bb5245cebf25791c36ae69b0d39c1f2f0e8defa8dd5c56a64600.NewPositionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PositionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.profile.positions.item collection
func (m *ProfileRequestBuilder) PositionsById(id string)(*i8450d9ff8966cf928b8991d8901a612103583e7b8300bc32e543a25ab8c60952.WorkPositionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["workPosition_id"] = id
    }
    return i8450d9ff8966cf928b8991d8901a612103583e7b8300bc32e543a25ab8c60952.NewWorkPositionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ProfileRequestBuilder) Projects()(*ia594d9f0fde4f12633b29cf2184447c01329657b9029df63e316ec608c382ce7.ProjectsRequestBuilder) {
    return ia594d9f0fde4f12633b29cf2184447c01329657b9029df63e316ec608c382ce7.NewProjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ProjectsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.profile.projects.item collection
func (m *ProfileRequestBuilder) ProjectsById(id string)(*ib34c8a74c1ba3d0d92a6d260bd9015795a3b2c70dbaef9ad4c485e831646e025.ProjectParticipationRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["projectParticipation_id"] = id
    }
    return ib34c8a74c1ba3d0d92a6d260bd9015795a3b2c70dbaef9ad4c485e831646e025.NewProjectParticipationRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ProfileRequestBuilder) Publications()(*idb2fd763c8ceac7682affef82cbd0bac0889f63a9904631022f51d8f81d00202.PublicationsRequestBuilder) {
    return idb2fd763c8ceac7682affef82cbd0bac0889f63a9904631022f51d8f81d00202.NewPublicationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PublicationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.profile.publications.item collection
func (m *ProfileRequestBuilder) PublicationsById(id string)(*iea021cd498fe18da6c27911f34b60740750e87aa77d2bc8965521a0e5f36c311.ItemPublicationRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemPublication_id"] = id
    }
    return iea021cd498fe18da6c27911f34b60740750e87aa77d2bc8965521a0e5f36c311.NewItemPublicationRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ProfileRequestBuilder) Skills()(*i26a7f8ef2723df7b554c727d88e4c0fade18dfb137078fee4bff0aaadde79d6b.SkillsRequestBuilder) {
    return i26a7f8ef2723df7b554c727d88e4c0fade18dfb137078fee4bff0aaadde79d6b.NewSkillsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SkillsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.profile.skills.item collection
func (m *ProfileRequestBuilder) SkillsById(id string)(*i0c6587bb49b5339ae5004ef05df87a4393eea15a87469e8b9347e8ed7ad9f098.SkillProficiencyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["skillProficiency_id"] = id
    }
    return i0c6587bb49b5339ae5004ef05df87a4393eea15a87469e8b9347e8ed7ad9f098.NewSkillProficiencyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ProfileRequestBuilder) WebAccounts()(*i7d8235f2ef58f14d9b48524d25c0ea6e65bf8d9e92bb04a40fe31ae1f66a3908.WebAccountsRequestBuilder) {
    return i7d8235f2ef58f14d9b48524d25c0ea6e65bf8d9e92bb04a40fe31ae1f66a3908.NewWebAccountsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WebAccountsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.profile.webAccounts.item collection
func (m *ProfileRequestBuilder) WebAccountsById(id string)(*i5a184e11a95aa2344b38024a42ff6a91fc4849920be15fb51f531e12f69b1ca6.WebAccountRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["webAccount_id"] = id
    }
    return i5a184e11a95aa2344b38024a42ff6a91fc4849920be15fb51f531e12f69b1ca6.NewWebAccountRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ProfileRequestBuilder) Websites()(*i376a1c9fb041f9485ffc4b05d361c22c731bfc23b1f7d7bd1ee1ccab102bbec4.WebsitesRequestBuilder) {
    return i376a1c9fb041f9485ffc4b05d361c22c731bfc23b1f7d7bd1ee1ccab102bbec4.NewWebsitesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WebsitesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.profile.websites.item collection
func (m *ProfileRequestBuilder) WebsitesById(id string)(*i042c646cb6e32821c8fad98a8f86ea3810ba71e5ce5f207b7727d32a421aa35d.PersonWebsiteRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["personWebsite_id"] = id
    }
    return i042c646cb6e32821c8fad98a8f86ea3810ba71e5ce5f207b7727d32a421aa35d.NewPersonWebsiteRequestBuilderInternal(urlTplParams, m.requestAdapter);
}

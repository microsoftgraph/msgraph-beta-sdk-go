package profile

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i0ca2cf05aa0ab8d575b29715dcde49cfad85816e7bd0dd6be6d8f28c5a57e0cc "github.com/microsoftgraph/msgraph-beta-sdk-go/me/profile/certifications"
    i17eb4344c5a0ad3def19ab63c562d6a940af904120cf3ba8cc2e5f247c5766f4 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/profile/publications"
    i1ee4d65d3257fb4603ee33eeff42b441404e6e3132f8c1c4158b20f867e38ba1 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/profile/languages"
    i2669dd5ed2aedac6dfd0c6508f784eed419fc338e30b07f910bb466760259fd1 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/profile/patents"
    i2d051d9e5b7dd4aa7e3df2ef2eef981e643ad151808b56ed2b5d7541b92e13a1 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/profile/emails"
    i51f3dc08dba3b9e52299dbf08214cce0f1ef2e7b1e7a2e495a3674b40711a3a9 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/profile/websites"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i56d5cd06e8002d9e89308b824b11d54f00236639eaa954149f137eff70aab691 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/profile/anniversaries"
    i5fac82f5724c79260d3690b79cac6213e87abd5d5b4732fc04b7ba6502202051 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/profile/notes"
    i66da464502bfa2f3d7040895c2dabd17e6fb785bd3beb256e7fc2924d2f8c581 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/profile/skills"
    i71cc69ae1f1790552d0d1f7b97f8fea4bcc00e0e41a5f5e400f8922a081f91d6 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/profile/projects"
    i7898d9baa7e9e2d6a7d32a761e8091d644c5366604eb46532fb8589fb563f424 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/profile/interests"
    i922a0397b3daba132d5755f51b216f39966edb1f0630f18a65831cbbf7ee6232 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/profile/names"
    ibdb9324e3601482a2267b1d9c0949c7d55a8080f5a4519f50c9653dc218292c0 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/profile/account"
    ic78d01b72199f27822bf3cf257e4574c3a9f57c17b0caaf0199ca0ab34c5051d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/profile/webaccounts"
    ic849b639cf0ed78d9c66d6ef19cccba41d1930bd2c2121ed950c74bc47bf74e5 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/profile/awards"
    idcb1d0f8bedd7256bb0cc1971e06d710738eda2e86753fc975ef06dbabd0dc98 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/profile/educationalactivities"
    idcfa0e287194c044fbfe4f5ea53c45e3e9260f37e9b15f0dcb0364f06774b230 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/profile/phones"
    ie9dcb8d043e17128686c43a95d98248fba5483cacc977feb238c5d674e1c89b4 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/profile/addresses"
    ief55708ca4039096496623eb43c04c4cab978519cc01987773642856bea25f15 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/profile/positions"
    i275a2397cc537d02a6de5d369b78e7f403c669152aed225ff04fe014e548e2bb "github.com/microsoftgraph/msgraph-beta-sdk-go/me/profile/addresses/item"
    i277bbd8b071a4d604a140752a228f28f247bd56c8d6ae617cd3a07f842cb6bf9 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/profile/awards/item"
    i2d0d6da11cb59700ac3ab9dd6b295e2ff497574052fdc8dc8cb99cfc50b629df "github.com/microsoftgraph/msgraph-beta-sdk-go/me/profile/anniversaries/item"
    i31cef5d667451b5423d5636714b09b3604e36345dba29f9ef24b76262d978aa5 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/profile/emails/item"
    i37b1fa517414e7c68dfd4a974844cbe1bf639c58cc2428815d70a432874a7634 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/profile/educationalactivities/item"
    i4833578e352a59b2b789284b4629d725d6ab0b3e30e29b40d1cb5c47837338b1 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/profile/account/item"
    i5052f9ec51b905e901e62b11f8a759a974f0cc7ab06fa1608d9d0c8d5fe5405b "github.com/microsoftgraph/msgraph-beta-sdk-go/me/profile/certifications/item"
    i59e9bd57c2f134047fa9d3ccde515ff620af19f86e8137229ea2935f1d483683 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/profile/positions/item"
    i5aa75bab3c02d7858a6bebecc29f8be43656cd44abcb0c3034d29e7461821067 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/profile/websites/item"
    i69ec9c16bbd15d4073a76419fab5dc22c3bddfcd0d392900a5ce53a55c061e2f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/profile/publications/item"
    i6a6a22b5e119da36bfd8b745f8ed59696bcbb84cbfc7c4b034075d1c1251b1f0 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/profile/interests/item"
    i6bbd51d2e2c2915bc291e644b1546ce4e0decac456826169fb0b87161ca8409e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/profile/patents/item"
    i7058e34b903fc8140673df6d1d8a644b427bfd66ffd15315f8c31382480ed4f6 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/profile/webaccounts/item"
    i89842cc8c04e65e1b1e88f5b767cc55ce09a0f30cc2e56c4f131d32c16de1159 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/profile/names/item"
    i9ff4662c3294ee098e0632304f1ef015fe0a686ef176fe16350508a3f9841068 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/profile/languages/item"
    ia79fb7e80b53cbbeaf958a755dd5660915548586799859d56f03c9b38f615a9d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/profile/phones/item"
    ib1d4f89ea65a822d70151ba9c171cd4c59a63443f8b798ab2f57a88825b43c6b "github.com/microsoftgraph/msgraph-beta-sdk-go/me/profile/projects/item"
    ib252e438900aca60fecae4f455e0915c86da07bd52b6015206b1a2debaf6ed36 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/profile/skills/item"
    ifa892d3191635eb436fe2e0ead51770cb25545dd430d2739450e3a0ab2fcff3a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/profile/notes/item"
)

// ProfileRequestBuilder builds and executes requests for operations under \me\profile
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
func (m *ProfileRequestBuilder) Account()(*ibdb9324e3601482a2267b1d9c0949c7d55a8080f5a4519f50c9653dc218292c0.AccountRequestBuilder) {
    return ibdb9324e3601482a2267b1d9c0949c7d55a8080f5a4519f50c9653dc218292c0.NewAccountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccountById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.profile.account.item collection
func (m *ProfileRequestBuilder) AccountById(id string)(*i4833578e352a59b2b789284b4629d725d6ab0b3e30e29b40d1cb5c47837338b1.UserAccountInformationRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userAccountInformation_id"] = id
    }
    return i4833578e352a59b2b789284b4629d725d6ab0b3e30e29b40d1cb5c47837338b1.NewUserAccountInformationRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ProfileRequestBuilder) Addresses()(*ie9dcb8d043e17128686c43a95d98248fba5483cacc977feb238c5d674e1c89b4.AddressesRequestBuilder) {
    return ie9dcb8d043e17128686c43a95d98248fba5483cacc977feb238c5d674e1c89b4.NewAddressesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AddressesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.profile.addresses.item collection
func (m *ProfileRequestBuilder) AddressesById(id string)(*i275a2397cc537d02a6de5d369b78e7f403c669152aed225ff04fe014e548e2bb.ItemAddressRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemAddress_id"] = id
    }
    return i275a2397cc537d02a6de5d369b78e7f403c669152aed225ff04fe014e548e2bb.NewItemAddressRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ProfileRequestBuilder) Anniversaries()(*i56d5cd06e8002d9e89308b824b11d54f00236639eaa954149f137eff70aab691.AnniversariesRequestBuilder) {
    return i56d5cd06e8002d9e89308b824b11d54f00236639eaa954149f137eff70aab691.NewAnniversariesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AnniversariesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.profile.anniversaries.item collection
func (m *ProfileRequestBuilder) AnniversariesById(id string)(*i2d0d6da11cb59700ac3ab9dd6b295e2ff497574052fdc8dc8cb99cfc50b629df.PersonAnnualEventRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["personAnnualEvent_id"] = id
    }
    return i2d0d6da11cb59700ac3ab9dd6b295e2ff497574052fdc8dc8cb99cfc50b629df.NewPersonAnnualEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ProfileRequestBuilder) Awards()(*ic849b639cf0ed78d9c66d6ef19cccba41d1930bd2c2121ed950c74bc47bf74e5.AwardsRequestBuilder) {
    return ic849b639cf0ed78d9c66d6ef19cccba41d1930bd2c2121ed950c74bc47bf74e5.NewAwardsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AwardsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.profile.awards.item collection
func (m *ProfileRequestBuilder) AwardsById(id string)(*i277bbd8b071a4d604a140752a228f28f247bd56c8d6ae617cd3a07f842cb6bf9.PersonAwardRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["personAward_id"] = id
    }
    return i277bbd8b071a4d604a140752a228f28f247bd56c8d6ae617cd3a07f842cb6bf9.NewPersonAwardRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ProfileRequestBuilder) Certifications()(*i0ca2cf05aa0ab8d575b29715dcde49cfad85816e7bd0dd6be6d8f28c5a57e0cc.CertificationsRequestBuilder) {
    return i0ca2cf05aa0ab8d575b29715dcde49cfad85816e7bd0dd6be6d8f28c5a57e0cc.NewCertificationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CertificationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.profile.certifications.item collection
func (m *ProfileRequestBuilder) CertificationsById(id string)(*i5052f9ec51b905e901e62b11f8a759a974f0cc7ab06fa1608d9d0c8d5fe5405b.PersonCertificationRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["personCertification_id"] = id
    }
    return i5052f9ec51b905e901e62b11f8a759a974f0cc7ab06fa1608d9d0c8d5fe5405b.NewPersonCertificationRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewProfileRequestBuilderInternal instantiates a new ProfileRequestBuilder and sets the default values.
func NewProfileRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ProfileRequestBuilder) {
    m := &ProfileRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/profile{?select,expand}";
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
func (m *ProfileRequestBuilder) EducationalActivities()(*idcb1d0f8bedd7256bb0cc1971e06d710738eda2e86753fc975ef06dbabd0dc98.EducationalActivitiesRequestBuilder) {
    return idcb1d0f8bedd7256bb0cc1971e06d710738eda2e86753fc975ef06dbabd0dc98.NewEducationalActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EducationalActivitiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.profile.educationalActivities.item collection
func (m *ProfileRequestBuilder) EducationalActivitiesById(id string)(*i37b1fa517414e7c68dfd4a974844cbe1bf639c58cc2428815d70a432874a7634.EducationalActivityRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationalActivity_id"] = id
    }
    return i37b1fa517414e7c68dfd4a974844cbe1bf639c58cc2428815d70a432874a7634.NewEducationalActivityRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ProfileRequestBuilder) Emails()(*i2d051d9e5b7dd4aa7e3df2ef2eef981e643ad151808b56ed2b5d7541b92e13a1.EmailsRequestBuilder) {
    return i2d051d9e5b7dd4aa7e3df2ef2eef981e643ad151808b56ed2b5d7541b92e13a1.NewEmailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EmailsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.profile.emails.item collection
func (m *ProfileRequestBuilder) EmailsById(id string)(*i31cef5d667451b5423d5636714b09b3604e36345dba29f9ef24b76262d978aa5.ItemEmailRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemEmail_id"] = id
    }
    return i31cef5d667451b5423d5636714b09b3604e36345dba29f9ef24b76262d978aa5.NewItemEmailRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *ProfileRequestBuilder) Interests()(*i7898d9baa7e9e2d6a7d32a761e8091d644c5366604eb46532fb8589fb563f424.InterestsRequestBuilder) {
    return i7898d9baa7e9e2d6a7d32a761e8091d644c5366604eb46532fb8589fb563f424.NewInterestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InterestsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.profile.interests.item collection
func (m *ProfileRequestBuilder) InterestsById(id string)(*i6a6a22b5e119da36bfd8b745f8ed59696bcbb84cbfc7c4b034075d1c1251b1f0.PersonInterestRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["personInterest_id"] = id
    }
    return i6a6a22b5e119da36bfd8b745f8ed59696bcbb84cbfc7c4b034075d1c1251b1f0.NewPersonInterestRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ProfileRequestBuilder) Languages()(*i1ee4d65d3257fb4603ee33eeff42b441404e6e3132f8c1c4158b20f867e38ba1.LanguagesRequestBuilder) {
    return i1ee4d65d3257fb4603ee33eeff42b441404e6e3132f8c1c4158b20f867e38ba1.NewLanguagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// LanguagesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.profile.languages.item collection
func (m *ProfileRequestBuilder) LanguagesById(id string)(*i9ff4662c3294ee098e0632304f1ef015fe0a686ef176fe16350508a3f9841068.LanguageProficiencyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["languageProficiency_id"] = id
    }
    return i9ff4662c3294ee098e0632304f1ef015fe0a686ef176fe16350508a3f9841068.NewLanguageProficiencyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ProfileRequestBuilder) Names()(*i922a0397b3daba132d5755f51b216f39966edb1f0630f18a65831cbbf7ee6232.NamesRequestBuilder) {
    return i922a0397b3daba132d5755f51b216f39966edb1f0630f18a65831cbbf7ee6232.NewNamesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NamesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.profile.names.item collection
func (m *ProfileRequestBuilder) NamesById(id string)(*i89842cc8c04e65e1b1e88f5b767cc55ce09a0f30cc2e56c4f131d32c16de1159.PersonNameRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["personName_id"] = id
    }
    return i89842cc8c04e65e1b1e88f5b767cc55ce09a0f30cc2e56c4f131d32c16de1159.NewPersonNameRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ProfileRequestBuilder) Notes()(*i5fac82f5724c79260d3690b79cac6213e87abd5d5b4732fc04b7ba6502202051.NotesRequestBuilder) {
    return i5fac82f5724c79260d3690b79cac6213e87abd5d5b4732fc04b7ba6502202051.NewNotesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NotesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.profile.notes.item collection
func (m *ProfileRequestBuilder) NotesById(id string)(*ifa892d3191635eb436fe2e0ead51770cb25545dd430d2739450e3a0ab2fcff3a.PersonAnnotationRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["personAnnotation_id"] = id
    }
    return ifa892d3191635eb436fe2e0ead51770cb25545dd430d2739450e3a0ab2fcff3a.NewPersonAnnotationRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *ProfileRequestBuilder) Patents()(*i2669dd5ed2aedac6dfd0c6508f784eed419fc338e30b07f910bb466760259fd1.PatentsRequestBuilder) {
    return i2669dd5ed2aedac6dfd0c6508f784eed419fc338e30b07f910bb466760259fd1.NewPatentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PatentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.profile.patents.item collection
func (m *ProfileRequestBuilder) PatentsById(id string)(*i6bbd51d2e2c2915bc291e644b1546ce4e0decac456826169fb0b87161ca8409e.ItemPatentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemPatent_id"] = id
    }
    return i6bbd51d2e2c2915bc291e644b1546ce4e0decac456826169fb0b87161ca8409e.NewItemPatentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ProfileRequestBuilder) Phones()(*idcfa0e287194c044fbfe4f5ea53c45e3e9260f37e9b15f0dcb0364f06774b230.PhonesRequestBuilder) {
    return idcfa0e287194c044fbfe4f5ea53c45e3e9260f37e9b15f0dcb0364f06774b230.NewPhonesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PhonesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.profile.phones.item collection
func (m *ProfileRequestBuilder) PhonesById(id string)(*ia79fb7e80b53cbbeaf958a755dd5660915548586799859d56f03c9b38f615a9d.ItemPhoneRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemPhone_id"] = id
    }
    return ia79fb7e80b53cbbeaf958a755dd5660915548586799859d56f03c9b38f615a9d.NewItemPhoneRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ProfileRequestBuilder) Positions()(*ief55708ca4039096496623eb43c04c4cab978519cc01987773642856bea25f15.PositionsRequestBuilder) {
    return ief55708ca4039096496623eb43c04c4cab978519cc01987773642856bea25f15.NewPositionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PositionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.profile.positions.item collection
func (m *ProfileRequestBuilder) PositionsById(id string)(*i59e9bd57c2f134047fa9d3ccde515ff620af19f86e8137229ea2935f1d483683.WorkPositionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["workPosition_id"] = id
    }
    return i59e9bd57c2f134047fa9d3ccde515ff620af19f86e8137229ea2935f1d483683.NewWorkPositionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ProfileRequestBuilder) Projects()(*i71cc69ae1f1790552d0d1f7b97f8fea4bcc00e0e41a5f5e400f8922a081f91d6.ProjectsRequestBuilder) {
    return i71cc69ae1f1790552d0d1f7b97f8fea4bcc00e0e41a5f5e400f8922a081f91d6.NewProjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ProjectsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.profile.projects.item collection
func (m *ProfileRequestBuilder) ProjectsById(id string)(*ib1d4f89ea65a822d70151ba9c171cd4c59a63443f8b798ab2f57a88825b43c6b.ProjectParticipationRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["projectParticipation_id"] = id
    }
    return ib1d4f89ea65a822d70151ba9c171cd4c59a63443f8b798ab2f57a88825b43c6b.NewProjectParticipationRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ProfileRequestBuilder) Publications()(*i17eb4344c5a0ad3def19ab63c562d6a940af904120cf3ba8cc2e5f247c5766f4.PublicationsRequestBuilder) {
    return i17eb4344c5a0ad3def19ab63c562d6a940af904120cf3ba8cc2e5f247c5766f4.NewPublicationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PublicationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.profile.publications.item collection
func (m *ProfileRequestBuilder) PublicationsById(id string)(*i69ec9c16bbd15d4073a76419fab5dc22c3bddfcd0d392900a5ce53a55c061e2f.ItemPublicationRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemPublication_id"] = id
    }
    return i69ec9c16bbd15d4073a76419fab5dc22c3bddfcd0d392900a5ce53a55c061e2f.NewItemPublicationRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ProfileRequestBuilder) Skills()(*i66da464502bfa2f3d7040895c2dabd17e6fb785bd3beb256e7fc2924d2f8c581.SkillsRequestBuilder) {
    return i66da464502bfa2f3d7040895c2dabd17e6fb785bd3beb256e7fc2924d2f8c581.NewSkillsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SkillsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.profile.skills.item collection
func (m *ProfileRequestBuilder) SkillsById(id string)(*ib252e438900aca60fecae4f455e0915c86da07bd52b6015206b1a2debaf6ed36.SkillProficiencyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["skillProficiency_id"] = id
    }
    return ib252e438900aca60fecae4f455e0915c86da07bd52b6015206b1a2debaf6ed36.NewSkillProficiencyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ProfileRequestBuilder) WebAccounts()(*ic78d01b72199f27822bf3cf257e4574c3a9f57c17b0caaf0199ca0ab34c5051d.WebAccountsRequestBuilder) {
    return ic78d01b72199f27822bf3cf257e4574c3a9f57c17b0caaf0199ca0ab34c5051d.NewWebAccountsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WebAccountsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.profile.webAccounts.item collection
func (m *ProfileRequestBuilder) WebAccountsById(id string)(*i7058e34b903fc8140673df6d1d8a644b427bfd66ffd15315f8c31382480ed4f6.WebAccountRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["webAccount_id"] = id
    }
    return i7058e34b903fc8140673df6d1d8a644b427bfd66ffd15315f8c31382480ed4f6.NewWebAccountRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ProfileRequestBuilder) Websites()(*i51f3dc08dba3b9e52299dbf08214cce0f1ef2e7b1e7a2e495a3674b40711a3a9.WebsitesRequestBuilder) {
    return i51f3dc08dba3b9e52299dbf08214cce0f1ef2e7b1e7a2e495a3674b40711a3a9.NewWebsitesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WebsitesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.profile.websites.item collection
func (m *ProfileRequestBuilder) WebsitesById(id string)(*i5aa75bab3c02d7858a6bebecc29f8be43656cd44abcb0c3034d29e7461821067.PersonWebsiteRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["personWebsite_id"] = id
    }
    return i5aa75bab3c02d7858a6bebecc29f8be43656cd44abcb0c3034d29e7461821067.NewPersonWebsiteRequestBuilderInternal(urlTplParams, m.requestAdapter);
}

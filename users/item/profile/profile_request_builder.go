package profile

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
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

// ProfileRequestBuilder provides operations to manage the profile property of the microsoft.graph.user entity.
type ProfileRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ProfileRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ProfileRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ProfileRequestBuilderGetQueryParameters retrieve the properties and relationships of a profile object for a given user. The **profile** resource exposes various rich properties that are descriptive of the user as relationships, for example, anniversaries and education activities. To get one of these navigation properties, use the corresponding GET method on that property. See the methods exposed by **profile**.
type ProfileRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ProfileRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ProfileRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ProfileRequestBuilderGetQueryParameters
}
// ProfileRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ProfileRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Account the account property
func (m *ProfileRequestBuilder) Account()(*id9a14f7d08a6d2f477b59c98cd195160e5cc687372332cdcac1fb9a6c27a6f32.AccountRequestBuilder) {
    return id9a14f7d08a6d2f477b59c98cd195160e5cc687372332cdcac1fb9a6c27a6f32.NewAccountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccountById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.profile.account.item collection
func (m *ProfileRequestBuilder) AccountById(id string)(*i1942a76b21bedb363b344695ae63b964a7c1e7f8cbee2dd8a053356f4cdea025.UserAccountInformationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userAccountInformation%2Did"] = id
    }
    return i1942a76b21bedb363b344695ae63b964a7c1e7f8cbee2dd8a053356f4cdea025.NewUserAccountInformationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Addresses the addresses property
func (m *ProfileRequestBuilder) Addresses()(*i8708d000e3379870d86078c6f35b29d11715cdfd03b38c7765791918f0672600.AddressesRequestBuilder) {
    return i8708d000e3379870d86078c6f35b29d11715cdfd03b38c7765791918f0672600.NewAddressesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AddressesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.profile.addresses.item collection
func (m *ProfileRequestBuilder) AddressesById(id string)(*i86cafdce36cb3eb513134ac90aeb1f13b19256a2d7e03fc8eb4d029c011328e7.ItemAddressItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemAddress%2Did"] = id
    }
    return i86cafdce36cb3eb513134ac90aeb1f13b19256a2d7e03fc8eb4d029c011328e7.NewItemAddressItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Anniversaries the anniversaries property
func (m *ProfileRequestBuilder) Anniversaries()(*i8f9bf67bcb3cda9c5f1048d104da873d4944ff13d895646ec920d2c76bfff28d.AnniversariesRequestBuilder) {
    return i8f9bf67bcb3cda9c5f1048d104da873d4944ff13d895646ec920d2c76bfff28d.NewAnniversariesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AnniversariesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.profile.anniversaries.item collection
func (m *ProfileRequestBuilder) AnniversariesById(id string)(*i770fc10a67210474d14ec729151f910b743e9684eb17d1374540e9da62dc6f85.PersonAnnualEventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["personAnnualEvent%2Did"] = id
    }
    return i770fc10a67210474d14ec729151f910b743e9684eb17d1374540e9da62dc6f85.NewPersonAnnualEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Awards the awards property
func (m *ProfileRequestBuilder) Awards()(*i6e5ffca00124afb0244dcbd530449478920e5af0605e6c2b4ffc6742b991ee48.AwardsRequestBuilder) {
    return i6e5ffca00124afb0244dcbd530449478920e5af0605e6c2b4ffc6742b991ee48.NewAwardsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AwardsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.profile.awards.item collection
func (m *ProfileRequestBuilder) AwardsById(id string)(*i6b52d42e9fa352b8843a815622d380fd88fb3108ec3297f67307e922fe108110.PersonAwardItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["personAward%2Did"] = id
    }
    return i6b52d42e9fa352b8843a815622d380fd88fb3108ec3297f67307e922fe108110.NewPersonAwardItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Certifications the certifications property
func (m *ProfileRequestBuilder) Certifications()(*id61a36d1ca39b427b83fcc2c83fef77cca045a5aab72b1488d7f99f1260ac02f.CertificationsRequestBuilder) {
    return id61a36d1ca39b427b83fcc2c83fef77cca045a5aab72b1488d7f99f1260ac02f.NewCertificationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CertificationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.profile.certifications.item collection
func (m *ProfileRequestBuilder) CertificationsById(id string)(*i4701b652703be8fdd84b7e4c9738fbe17c3e8ba6dd477ae00e44ab3e65041138.PersonCertificationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["personCertification%2Did"] = id
    }
    return i4701b652703be8fdd84b7e4c9738fbe17c3e8ba6dd477ae00e44ab3e65041138.NewPersonCertificationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewProfileRequestBuilderInternal instantiates a new ProfileRequestBuilder and sets the default values.
func NewProfileRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ProfileRequestBuilder) {
    m := &ProfileRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/profile{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewProfileRequestBuilder instantiates a new ProfileRequestBuilder and sets the default values.
func NewProfileRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ProfileRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewProfileRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation deletes a profile object from a user's account.
func (m *ProfileRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *ProfileRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation retrieve the properties and relationships of a profile object for a given user. The **profile** resource exposes various rich properties that are descriptive of the user as relationships, for example, anniversaries and education activities. To get one of these navigation properties, use the corresponding GET method on that property. See the methods exposed by **profile**.
func (m *ProfileRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *ProfileRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property profile in users
func (m *ProfileRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Profileable, requestConfiguration *ProfileRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete deletes a profile object from a user's account.
func (m *ProfileRequestBuilder) Delete(ctx context.Context, requestConfiguration *ProfileRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// EducationalActivities the educationalActivities property
func (m *ProfileRequestBuilder) EducationalActivities()(*id664db371311b8f8e8c63b1bcd1e9e389320ce441b0df7fd171ae3e2183e3ef2.EducationalActivitiesRequestBuilder) {
    return id664db371311b8f8e8c63b1bcd1e9e389320ce441b0df7fd171ae3e2183e3ef2.NewEducationalActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EducationalActivitiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.profile.educationalActivities.item collection
func (m *ProfileRequestBuilder) EducationalActivitiesById(id string)(*i42f9ec73b968bdabb69618c8eb1906cb3c1655e5c30183b4a20aae652614b827.EducationalActivityItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationalActivity%2Did"] = id
    }
    return i42f9ec73b968bdabb69618c8eb1906cb3c1655e5c30183b4a20aae652614b827.NewEducationalActivityItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Emails the emails property
func (m *ProfileRequestBuilder) Emails()(*i66a493024db3c54c6d1d0a7a037cc42558a6cc37204f250864539ee332b7e218.EmailsRequestBuilder) {
    return i66a493024db3c54c6d1d0a7a037cc42558a6cc37204f250864539ee332b7e218.NewEmailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EmailsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.profile.emails.item collection
func (m *ProfileRequestBuilder) EmailsById(id string)(*i6cb42d6092f89243d60ec2eff66d3051d0f316e72ca67a0846be1031eb378549.ItemEmailItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemEmail%2Did"] = id
    }
    return i6cb42d6092f89243d60ec2eff66d3051d0f316e72ca67a0846be1031eb378549.NewItemEmailItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get retrieve the properties and relationships of a profile object for a given user. The **profile** resource exposes various rich properties that are descriptive of the user as relationships, for example, anniversaries and education activities. To get one of these navigation properties, use the corresponding GET method on that property. See the methods exposed by **profile**.
func (m *ProfileRequestBuilder) Get(ctx context.Context, requestConfiguration *ProfileRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Profileable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateProfileFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Profileable), nil
}
// Interests the interests property
func (m *ProfileRequestBuilder) Interests()(*i46891830191b8b241de7376dce57ac07cfd53ae6fd0013bd8e1e973a36e35d0a.InterestsRequestBuilder) {
    return i46891830191b8b241de7376dce57ac07cfd53ae6fd0013bd8e1e973a36e35d0a.NewInterestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InterestsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.profile.interests.item collection
func (m *ProfileRequestBuilder) InterestsById(id string)(*i1d6a22e1958f8e610c82479a0a2d9fbfbe737522bed6b1ae128aed42eec5de72.PersonInterestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["personInterest%2Did"] = id
    }
    return i1d6a22e1958f8e610c82479a0a2d9fbfbe737522bed6b1ae128aed42eec5de72.NewPersonInterestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Languages the languages property
func (m *ProfileRequestBuilder) Languages()(*ie07f4e96ea74d56fa71e6fc9c3cf4c8a7aa0b94b884c31955e02827215f4b842.LanguagesRequestBuilder) {
    return ie07f4e96ea74d56fa71e6fc9c3cf4c8a7aa0b94b884c31955e02827215f4b842.NewLanguagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// LanguagesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.profile.languages.item collection
func (m *ProfileRequestBuilder) LanguagesById(id string)(*ib3dd68449e016ac77f213ca43faab02d165f0dd15c9fbbb56bc0094df6270197.LanguageProficiencyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["languageProficiency%2Did"] = id
    }
    return ib3dd68449e016ac77f213ca43faab02d165f0dd15c9fbbb56bc0094df6270197.NewLanguageProficiencyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Names the names property
func (m *ProfileRequestBuilder) Names()(*i8e8af53beda11d49c2fe34af67e33e01fbfc8215ef688bcc9a6a99b22213c975.NamesRequestBuilder) {
    return i8e8af53beda11d49c2fe34af67e33e01fbfc8215ef688bcc9a6a99b22213c975.NewNamesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NamesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.profile.names.item collection
func (m *ProfileRequestBuilder) NamesById(id string)(*i1447f59bce732e48357e480ddb4ec0ee411855640d3f9e1da6ed7fc9d4086d64.PersonNameItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["personName%2Did"] = id
    }
    return i1447f59bce732e48357e480ddb4ec0ee411855640d3f9e1da6ed7fc9d4086d64.NewPersonNameItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Notes the notes property
func (m *ProfileRequestBuilder) Notes()(*i3fa49a978d179839dbea3f8dd172662623ea0846568dc802484468eb4f5b89d8.NotesRequestBuilder) {
    return i3fa49a978d179839dbea3f8dd172662623ea0846568dc802484468eb4f5b89d8.NewNotesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NotesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.profile.notes.item collection
func (m *ProfileRequestBuilder) NotesById(id string)(*i42cff4ee1200327093e5e7ebc241155dc70d33e9a1329d8d67d8fb56f8b951a9.PersonAnnotationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["personAnnotation%2Did"] = id
    }
    return i42cff4ee1200327093e5e7ebc241155dc70d33e9a1329d8d67d8fb56f8b951a9.NewPersonAnnotationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property profile in users
func (m *ProfileRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Profileable, requestConfiguration *ProfileRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Profileable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateProfileFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Profileable), nil
}
// Patents the patents property
func (m *ProfileRequestBuilder) Patents()(*i624b8f0e161cffd6542ffd795e358c6db683b8120f9e592078e0ed9f8fed0eca.PatentsRequestBuilder) {
    return i624b8f0e161cffd6542ffd795e358c6db683b8120f9e592078e0ed9f8fed0eca.NewPatentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PatentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.profile.patents.item collection
func (m *ProfileRequestBuilder) PatentsById(id string)(*i91b663b48ef705f87f8309ddf035d579ff8c2c2b8eb0689c1eff786c8be74538.ItemPatentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemPatent%2Did"] = id
    }
    return i91b663b48ef705f87f8309ddf035d579ff8c2c2b8eb0689c1eff786c8be74538.NewItemPatentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Phones the phones property
func (m *ProfileRequestBuilder) Phones()(*i54063d6549cfffd18152df5918e18636c36e95ba53e40f0a76aab9477de26aa6.PhonesRequestBuilder) {
    return i54063d6549cfffd18152df5918e18636c36e95ba53e40f0a76aab9477de26aa6.NewPhonesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PhonesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.profile.phones.item collection
func (m *ProfileRequestBuilder) PhonesById(id string)(*i09dbdbf1122b782cd8c1844df30cff1030466125501cfa48f39bace4667a50ab.ItemPhoneItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemPhone%2Did"] = id
    }
    return i09dbdbf1122b782cd8c1844df30cff1030466125501cfa48f39bace4667a50ab.NewItemPhoneItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Positions the positions property
func (m *ProfileRequestBuilder) Positions()(*idcb0b21c16e8bb5245cebf25791c36ae69b0d39c1f2f0e8defa8dd5c56a64600.PositionsRequestBuilder) {
    return idcb0b21c16e8bb5245cebf25791c36ae69b0d39c1f2f0e8defa8dd5c56a64600.NewPositionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PositionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.profile.positions.item collection
func (m *ProfileRequestBuilder) PositionsById(id string)(*i8450d9ff8966cf928b8991d8901a612103583e7b8300bc32e543a25ab8c60952.WorkPositionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["workPosition%2Did"] = id
    }
    return i8450d9ff8966cf928b8991d8901a612103583e7b8300bc32e543a25ab8c60952.NewWorkPositionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Projects the projects property
func (m *ProfileRequestBuilder) Projects()(*ia594d9f0fde4f12633b29cf2184447c01329657b9029df63e316ec608c382ce7.ProjectsRequestBuilder) {
    return ia594d9f0fde4f12633b29cf2184447c01329657b9029df63e316ec608c382ce7.NewProjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ProjectsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.profile.projects.item collection
func (m *ProfileRequestBuilder) ProjectsById(id string)(*ib34c8a74c1ba3d0d92a6d260bd9015795a3b2c70dbaef9ad4c485e831646e025.ProjectParticipationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["projectParticipation%2Did"] = id
    }
    return ib34c8a74c1ba3d0d92a6d260bd9015795a3b2c70dbaef9ad4c485e831646e025.NewProjectParticipationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Publications the publications property
func (m *ProfileRequestBuilder) Publications()(*idb2fd763c8ceac7682affef82cbd0bac0889f63a9904631022f51d8f81d00202.PublicationsRequestBuilder) {
    return idb2fd763c8ceac7682affef82cbd0bac0889f63a9904631022f51d8f81d00202.NewPublicationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PublicationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.profile.publications.item collection
func (m *ProfileRequestBuilder) PublicationsById(id string)(*iea021cd498fe18da6c27911f34b60740750e87aa77d2bc8965521a0e5f36c311.ItemPublicationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemPublication%2Did"] = id
    }
    return iea021cd498fe18da6c27911f34b60740750e87aa77d2bc8965521a0e5f36c311.NewItemPublicationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Skills the skills property
func (m *ProfileRequestBuilder) Skills()(*i26a7f8ef2723df7b554c727d88e4c0fade18dfb137078fee4bff0aaadde79d6b.SkillsRequestBuilder) {
    return i26a7f8ef2723df7b554c727d88e4c0fade18dfb137078fee4bff0aaadde79d6b.NewSkillsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SkillsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.profile.skills.item collection
func (m *ProfileRequestBuilder) SkillsById(id string)(*i0c6587bb49b5339ae5004ef05df87a4393eea15a87469e8b9347e8ed7ad9f098.SkillProficiencyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["skillProficiency%2Did"] = id
    }
    return i0c6587bb49b5339ae5004ef05df87a4393eea15a87469e8b9347e8ed7ad9f098.NewSkillProficiencyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// WebAccounts the webAccounts property
func (m *ProfileRequestBuilder) WebAccounts()(*i7d8235f2ef58f14d9b48524d25c0ea6e65bf8d9e92bb04a40fe31ae1f66a3908.WebAccountsRequestBuilder) {
    return i7d8235f2ef58f14d9b48524d25c0ea6e65bf8d9e92bb04a40fe31ae1f66a3908.NewWebAccountsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WebAccountsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.profile.webAccounts.item collection
func (m *ProfileRequestBuilder) WebAccountsById(id string)(*i5a184e11a95aa2344b38024a42ff6a91fc4849920be15fb51f531e12f69b1ca6.WebAccountItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["webAccount%2Did"] = id
    }
    return i5a184e11a95aa2344b38024a42ff6a91fc4849920be15fb51f531e12f69b1ca6.NewWebAccountItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Websites the websites property
func (m *ProfileRequestBuilder) Websites()(*i376a1c9fb041f9485ffc4b05d361c22c731bfc23b1f7d7bd1ee1ccab102bbec4.WebsitesRequestBuilder) {
    return i376a1c9fb041f9485ffc4b05d361c22c731bfc23b1f7d7bd1ee1ccab102bbec4.NewWebsitesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WebsitesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.profile.websites.item collection
func (m *ProfileRequestBuilder) WebsitesById(id string)(*i042c646cb6e32821c8fad98a8f86ea3810ba71e5ce5f207b7727d32a421aa35d.PersonWebsiteItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["personWebsite%2Did"] = id
    }
    return i042c646cb6e32821c8fad98a8f86ea3810ba71e5ce5f207b7727d32a421aa35d.NewPersonWebsiteItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}

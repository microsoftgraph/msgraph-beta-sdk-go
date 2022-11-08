package profile

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i0ca2cf05aa0ab8d575b29715dcde49cfad85816e7bd0dd6be6d8f28c5a57e0cc "github.com/microsoftgraph/msgraph-beta-sdk-go/me/profile/certifications"
    i17eb4344c5a0ad3def19ab63c562d6a940af904120cf3ba8cc2e5f247c5766f4 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/profile/publications"
    i1ee4d65d3257fb4603ee33eeff42b441404e6e3132f8c1c4158b20f867e38ba1 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/profile/languages"
    i2669dd5ed2aedac6dfd0c6508f784eed419fc338e30b07f910bb466760259fd1 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/profile/patents"
    i2d051d9e5b7dd4aa7e3df2ef2eef981e643ad151808b56ed2b5d7541b92e13a1 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/profile/emails"
    i51f3dc08dba3b9e52299dbf08214cce0f1ef2e7b1e7a2e495a3674b40711a3a9 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/profile/websites"
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
// Account provides operations to manage the account property of the microsoft.graph.profile entity.
func (m *ProfileRequestBuilder) Account()(*ibdb9324e3601482a2267b1d9c0949c7d55a8080f5a4519f50c9653dc218292c0.AccountRequestBuilder) {
    return ibdb9324e3601482a2267b1d9c0949c7d55a8080f5a4519f50c9653dc218292c0.NewAccountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccountById provides operations to manage the account property of the microsoft.graph.profile entity.
func (m *ProfileRequestBuilder) AccountById(id string)(*i4833578e352a59b2b789284b4629d725d6ab0b3e30e29b40d1cb5c47837338b1.UserAccountInformationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userAccountInformation%2Did"] = id
    }
    return i4833578e352a59b2b789284b4629d725d6ab0b3e30e29b40d1cb5c47837338b1.NewUserAccountInformationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Addresses provides operations to manage the addresses property of the microsoft.graph.profile entity.
func (m *ProfileRequestBuilder) Addresses()(*ie9dcb8d043e17128686c43a95d98248fba5483cacc977feb238c5d674e1c89b4.AddressesRequestBuilder) {
    return ie9dcb8d043e17128686c43a95d98248fba5483cacc977feb238c5d674e1c89b4.NewAddressesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AddressesById provides operations to manage the addresses property of the microsoft.graph.profile entity.
func (m *ProfileRequestBuilder) AddressesById(id string)(*i275a2397cc537d02a6de5d369b78e7f403c669152aed225ff04fe014e548e2bb.ItemAddressItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemAddress%2Did"] = id
    }
    return i275a2397cc537d02a6de5d369b78e7f403c669152aed225ff04fe014e548e2bb.NewItemAddressItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Anniversaries provides operations to manage the anniversaries property of the microsoft.graph.profile entity.
func (m *ProfileRequestBuilder) Anniversaries()(*i56d5cd06e8002d9e89308b824b11d54f00236639eaa954149f137eff70aab691.AnniversariesRequestBuilder) {
    return i56d5cd06e8002d9e89308b824b11d54f00236639eaa954149f137eff70aab691.NewAnniversariesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AnniversariesById provides operations to manage the anniversaries property of the microsoft.graph.profile entity.
func (m *ProfileRequestBuilder) AnniversariesById(id string)(*i2d0d6da11cb59700ac3ab9dd6b295e2ff497574052fdc8dc8cb99cfc50b629df.PersonAnnualEventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["personAnnualEvent%2Did"] = id
    }
    return i2d0d6da11cb59700ac3ab9dd6b295e2ff497574052fdc8dc8cb99cfc50b629df.NewPersonAnnualEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Awards provides operations to manage the awards property of the microsoft.graph.profile entity.
func (m *ProfileRequestBuilder) Awards()(*ic849b639cf0ed78d9c66d6ef19cccba41d1930bd2c2121ed950c74bc47bf74e5.AwardsRequestBuilder) {
    return ic849b639cf0ed78d9c66d6ef19cccba41d1930bd2c2121ed950c74bc47bf74e5.NewAwardsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AwardsById provides operations to manage the awards property of the microsoft.graph.profile entity.
func (m *ProfileRequestBuilder) AwardsById(id string)(*i277bbd8b071a4d604a140752a228f28f247bd56c8d6ae617cd3a07f842cb6bf9.PersonAwardItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["personAward%2Did"] = id
    }
    return i277bbd8b071a4d604a140752a228f28f247bd56c8d6ae617cd3a07f842cb6bf9.NewPersonAwardItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Certifications provides operations to manage the certifications property of the microsoft.graph.profile entity.
func (m *ProfileRequestBuilder) Certifications()(*i0ca2cf05aa0ab8d575b29715dcde49cfad85816e7bd0dd6be6d8f28c5a57e0cc.CertificationsRequestBuilder) {
    return i0ca2cf05aa0ab8d575b29715dcde49cfad85816e7bd0dd6be6d8f28c5a57e0cc.NewCertificationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CertificationsById provides operations to manage the certifications property of the microsoft.graph.profile entity.
func (m *ProfileRequestBuilder) CertificationsById(id string)(*i5052f9ec51b905e901e62b11f8a759a974f0cc7ab06fa1608d9d0c8d5fe5405b.PersonCertificationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["personCertification%2Did"] = id
    }
    return i5052f9ec51b905e901e62b11f8a759a974f0cc7ab06fa1608d9d0c8d5fe5405b.NewPersonCertificationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewProfileRequestBuilderInternal instantiates a new ProfileRequestBuilder and sets the default values.
func NewProfileRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ProfileRequestBuilder) {
    m := &ProfileRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/profile{?%24select,%24expand}";
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
// CreatePatchRequestInformation update the navigation property profile in me
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
// EducationalActivities provides operations to manage the educationalActivities property of the microsoft.graph.profile entity.
func (m *ProfileRequestBuilder) EducationalActivities()(*idcb1d0f8bedd7256bb0cc1971e06d710738eda2e86753fc975ef06dbabd0dc98.EducationalActivitiesRequestBuilder) {
    return idcb1d0f8bedd7256bb0cc1971e06d710738eda2e86753fc975ef06dbabd0dc98.NewEducationalActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EducationalActivitiesById provides operations to manage the educationalActivities property of the microsoft.graph.profile entity.
func (m *ProfileRequestBuilder) EducationalActivitiesById(id string)(*i37b1fa517414e7c68dfd4a974844cbe1bf639c58cc2428815d70a432874a7634.EducationalActivityItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationalActivity%2Did"] = id
    }
    return i37b1fa517414e7c68dfd4a974844cbe1bf639c58cc2428815d70a432874a7634.NewEducationalActivityItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Emails provides operations to manage the emails property of the microsoft.graph.profile entity.
func (m *ProfileRequestBuilder) Emails()(*i2d051d9e5b7dd4aa7e3df2ef2eef981e643ad151808b56ed2b5d7541b92e13a1.EmailsRequestBuilder) {
    return i2d051d9e5b7dd4aa7e3df2ef2eef981e643ad151808b56ed2b5d7541b92e13a1.NewEmailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EmailsById provides operations to manage the emails property of the microsoft.graph.profile entity.
func (m *ProfileRequestBuilder) EmailsById(id string)(*i31cef5d667451b5423d5636714b09b3604e36345dba29f9ef24b76262d978aa5.ItemEmailItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemEmail%2Did"] = id
    }
    return i31cef5d667451b5423d5636714b09b3604e36345dba29f9ef24b76262d978aa5.NewItemEmailItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
// Interests provides operations to manage the interests property of the microsoft.graph.profile entity.
func (m *ProfileRequestBuilder) Interests()(*i7898d9baa7e9e2d6a7d32a761e8091d644c5366604eb46532fb8589fb563f424.InterestsRequestBuilder) {
    return i7898d9baa7e9e2d6a7d32a761e8091d644c5366604eb46532fb8589fb563f424.NewInterestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InterestsById provides operations to manage the interests property of the microsoft.graph.profile entity.
func (m *ProfileRequestBuilder) InterestsById(id string)(*i6a6a22b5e119da36bfd8b745f8ed59696bcbb84cbfc7c4b034075d1c1251b1f0.PersonInterestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["personInterest%2Did"] = id
    }
    return i6a6a22b5e119da36bfd8b745f8ed59696bcbb84cbfc7c4b034075d1c1251b1f0.NewPersonInterestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Languages provides operations to manage the languages property of the microsoft.graph.profile entity.
func (m *ProfileRequestBuilder) Languages()(*i1ee4d65d3257fb4603ee33eeff42b441404e6e3132f8c1c4158b20f867e38ba1.LanguagesRequestBuilder) {
    return i1ee4d65d3257fb4603ee33eeff42b441404e6e3132f8c1c4158b20f867e38ba1.NewLanguagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// LanguagesById provides operations to manage the languages property of the microsoft.graph.profile entity.
func (m *ProfileRequestBuilder) LanguagesById(id string)(*i9ff4662c3294ee098e0632304f1ef015fe0a686ef176fe16350508a3f9841068.LanguageProficiencyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["languageProficiency%2Did"] = id
    }
    return i9ff4662c3294ee098e0632304f1ef015fe0a686ef176fe16350508a3f9841068.NewLanguageProficiencyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Names provides operations to manage the names property of the microsoft.graph.profile entity.
func (m *ProfileRequestBuilder) Names()(*i922a0397b3daba132d5755f51b216f39966edb1f0630f18a65831cbbf7ee6232.NamesRequestBuilder) {
    return i922a0397b3daba132d5755f51b216f39966edb1f0630f18a65831cbbf7ee6232.NewNamesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NamesById provides operations to manage the names property of the microsoft.graph.profile entity.
func (m *ProfileRequestBuilder) NamesById(id string)(*i89842cc8c04e65e1b1e88f5b767cc55ce09a0f30cc2e56c4f131d32c16de1159.PersonNameItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["personName%2Did"] = id
    }
    return i89842cc8c04e65e1b1e88f5b767cc55ce09a0f30cc2e56c4f131d32c16de1159.NewPersonNameItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Notes provides operations to manage the notes property of the microsoft.graph.profile entity.
func (m *ProfileRequestBuilder) Notes()(*i5fac82f5724c79260d3690b79cac6213e87abd5d5b4732fc04b7ba6502202051.NotesRequestBuilder) {
    return i5fac82f5724c79260d3690b79cac6213e87abd5d5b4732fc04b7ba6502202051.NewNotesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NotesById provides operations to manage the notes property of the microsoft.graph.profile entity.
func (m *ProfileRequestBuilder) NotesById(id string)(*ifa892d3191635eb436fe2e0ead51770cb25545dd430d2739450e3a0ab2fcff3a.PersonAnnotationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["personAnnotation%2Did"] = id
    }
    return ifa892d3191635eb436fe2e0ead51770cb25545dd430d2739450e3a0ab2fcff3a.NewPersonAnnotationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property profile in me
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
// Patents provides operations to manage the patents property of the microsoft.graph.profile entity.
func (m *ProfileRequestBuilder) Patents()(*i2669dd5ed2aedac6dfd0c6508f784eed419fc338e30b07f910bb466760259fd1.PatentsRequestBuilder) {
    return i2669dd5ed2aedac6dfd0c6508f784eed419fc338e30b07f910bb466760259fd1.NewPatentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PatentsById provides operations to manage the patents property of the microsoft.graph.profile entity.
func (m *ProfileRequestBuilder) PatentsById(id string)(*i6bbd51d2e2c2915bc291e644b1546ce4e0decac456826169fb0b87161ca8409e.ItemPatentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemPatent%2Did"] = id
    }
    return i6bbd51d2e2c2915bc291e644b1546ce4e0decac456826169fb0b87161ca8409e.NewItemPatentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Phones provides operations to manage the phones property of the microsoft.graph.profile entity.
func (m *ProfileRequestBuilder) Phones()(*idcfa0e287194c044fbfe4f5ea53c45e3e9260f37e9b15f0dcb0364f06774b230.PhonesRequestBuilder) {
    return idcfa0e287194c044fbfe4f5ea53c45e3e9260f37e9b15f0dcb0364f06774b230.NewPhonesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PhonesById provides operations to manage the phones property of the microsoft.graph.profile entity.
func (m *ProfileRequestBuilder) PhonesById(id string)(*ia79fb7e80b53cbbeaf958a755dd5660915548586799859d56f03c9b38f615a9d.ItemPhoneItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemPhone%2Did"] = id
    }
    return ia79fb7e80b53cbbeaf958a755dd5660915548586799859d56f03c9b38f615a9d.NewItemPhoneItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Positions provides operations to manage the positions property of the microsoft.graph.profile entity.
func (m *ProfileRequestBuilder) Positions()(*ief55708ca4039096496623eb43c04c4cab978519cc01987773642856bea25f15.PositionsRequestBuilder) {
    return ief55708ca4039096496623eb43c04c4cab978519cc01987773642856bea25f15.NewPositionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PositionsById provides operations to manage the positions property of the microsoft.graph.profile entity.
func (m *ProfileRequestBuilder) PositionsById(id string)(*i59e9bd57c2f134047fa9d3ccde515ff620af19f86e8137229ea2935f1d483683.WorkPositionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["workPosition%2Did"] = id
    }
    return i59e9bd57c2f134047fa9d3ccde515ff620af19f86e8137229ea2935f1d483683.NewWorkPositionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Projects provides operations to manage the projects property of the microsoft.graph.profile entity.
func (m *ProfileRequestBuilder) Projects()(*i71cc69ae1f1790552d0d1f7b97f8fea4bcc00e0e41a5f5e400f8922a081f91d6.ProjectsRequestBuilder) {
    return i71cc69ae1f1790552d0d1f7b97f8fea4bcc00e0e41a5f5e400f8922a081f91d6.NewProjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ProjectsById provides operations to manage the projects property of the microsoft.graph.profile entity.
func (m *ProfileRequestBuilder) ProjectsById(id string)(*ib1d4f89ea65a822d70151ba9c171cd4c59a63443f8b798ab2f57a88825b43c6b.ProjectParticipationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["projectParticipation%2Did"] = id
    }
    return ib1d4f89ea65a822d70151ba9c171cd4c59a63443f8b798ab2f57a88825b43c6b.NewProjectParticipationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Publications provides operations to manage the publications property of the microsoft.graph.profile entity.
func (m *ProfileRequestBuilder) Publications()(*i17eb4344c5a0ad3def19ab63c562d6a940af904120cf3ba8cc2e5f247c5766f4.PublicationsRequestBuilder) {
    return i17eb4344c5a0ad3def19ab63c562d6a940af904120cf3ba8cc2e5f247c5766f4.NewPublicationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PublicationsById provides operations to manage the publications property of the microsoft.graph.profile entity.
func (m *ProfileRequestBuilder) PublicationsById(id string)(*i69ec9c16bbd15d4073a76419fab5dc22c3bddfcd0d392900a5ce53a55c061e2f.ItemPublicationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemPublication%2Did"] = id
    }
    return i69ec9c16bbd15d4073a76419fab5dc22c3bddfcd0d392900a5ce53a55c061e2f.NewItemPublicationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Skills provides operations to manage the skills property of the microsoft.graph.profile entity.
func (m *ProfileRequestBuilder) Skills()(*i66da464502bfa2f3d7040895c2dabd17e6fb785bd3beb256e7fc2924d2f8c581.SkillsRequestBuilder) {
    return i66da464502bfa2f3d7040895c2dabd17e6fb785bd3beb256e7fc2924d2f8c581.NewSkillsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SkillsById provides operations to manage the skills property of the microsoft.graph.profile entity.
func (m *ProfileRequestBuilder) SkillsById(id string)(*ib252e438900aca60fecae4f455e0915c86da07bd52b6015206b1a2debaf6ed36.SkillProficiencyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["skillProficiency%2Did"] = id
    }
    return ib252e438900aca60fecae4f455e0915c86da07bd52b6015206b1a2debaf6ed36.NewSkillProficiencyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// WebAccounts provides operations to manage the webAccounts property of the microsoft.graph.profile entity.
func (m *ProfileRequestBuilder) WebAccounts()(*ic78d01b72199f27822bf3cf257e4574c3a9f57c17b0caaf0199ca0ab34c5051d.WebAccountsRequestBuilder) {
    return ic78d01b72199f27822bf3cf257e4574c3a9f57c17b0caaf0199ca0ab34c5051d.NewWebAccountsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WebAccountsById provides operations to manage the webAccounts property of the microsoft.graph.profile entity.
func (m *ProfileRequestBuilder) WebAccountsById(id string)(*i7058e34b903fc8140673df6d1d8a644b427bfd66ffd15315f8c31382480ed4f6.WebAccountItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["webAccount%2Did"] = id
    }
    return i7058e34b903fc8140673df6d1d8a644b427bfd66ffd15315f8c31382480ed4f6.NewWebAccountItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Websites provides operations to manage the websites property of the microsoft.graph.profile entity.
func (m *ProfileRequestBuilder) Websites()(*i51f3dc08dba3b9e52299dbf08214cce0f1ef2e7b1e7a2e495a3674b40711a3a9.WebsitesRequestBuilder) {
    return i51f3dc08dba3b9e52299dbf08214cce0f1ef2e7b1e7a2e495a3674b40711a3a9.NewWebsitesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WebsitesById provides operations to manage the websites property of the microsoft.graph.profile entity.
func (m *ProfileRequestBuilder) WebsitesById(id string)(*i5aa75bab3c02d7858a6bebecc29f8be43656cd44abcb0c3034d29e7461821067.PersonWebsiteItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["personWebsite%2Did"] = id
    }
    return i5aa75bab3c02d7858a6bebecc29f8be43656cd44abcb0c3034d29e7461821067.NewPersonWebsiteItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}

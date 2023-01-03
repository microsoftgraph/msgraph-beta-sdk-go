package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemProfileRequestBuilder provides operations to manage the profile property of the microsoft.graph.user entity.
type ItemProfileRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemProfileRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemProfileRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemProfileRequestBuilderGetQueryParameters retrieve the properties and relationships of a profile object for a given user. The **profile** resource exposes various rich properties that are descriptive of the user as relationships, for example, anniversaries and education activities. To get one of these navigation properties, use the corresponding GET method on that property. See the methods exposed by **profile**.
type ItemProfileRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemProfileRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemProfileRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemProfileRequestBuilderGetQueryParameters
}
// ItemProfileRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemProfileRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Account provides operations to manage the account property of the microsoft.graph.profile entity.
func (m *ItemProfileRequestBuilder) Account()(*ItemProfileAccountRequestBuilder) {
    return NewItemProfileAccountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccountById provides operations to manage the account property of the microsoft.graph.profile entity.
func (m *ItemProfileRequestBuilder) AccountById(id string)(*UserAccountInformationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userAccountInformation%2Did"] = id
    }
    return NewUserAccountInformationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Addresses provides operations to manage the addresses property of the microsoft.graph.profile entity.
func (m *ItemProfileRequestBuilder) Addresses()(*ItemProfileAddressesRequestBuilder) {
    return NewItemProfileAddressesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AddressesById provides operations to manage the addresses property of the microsoft.graph.profile entity.
func (m *ItemProfileRequestBuilder) AddressesById(id string)(*ItemAddressItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemAddress%2Did"] = id
    }
    return NewItemAddressItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Anniversaries provides operations to manage the anniversaries property of the microsoft.graph.profile entity.
func (m *ItemProfileRequestBuilder) Anniversaries()(*ItemProfileAnniversariesRequestBuilder) {
    return NewItemProfileAnniversariesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AnniversariesById provides operations to manage the anniversaries property of the microsoft.graph.profile entity.
func (m *ItemProfileRequestBuilder) AnniversariesById(id string)(*PersonAnnualEventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["personAnnualEvent%2Did"] = id
    }
    return NewPersonAnnualEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Awards provides operations to manage the awards property of the microsoft.graph.profile entity.
func (m *ItemProfileRequestBuilder) Awards()(*ItemProfileAwardsRequestBuilder) {
    return NewItemProfileAwardsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AwardsById provides operations to manage the awards property of the microsoft.graph.profile entity.
func (m *ItemProfileRequestBuilder) AwardsById(id string)(*PersonAwardItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["personAward%2Did"] = id
    }
    return NewPersonAwardItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Certifications provides operations to manage the certifications property of the microsoft.graph.profile entity.
func (m *ItemProfileRequestBuilder) Certifications()(*ItemProfileCertificationsRequestBuilder) {
    return NewItemProfileCertificationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CertificationsById provides operations to manage the certifications property of the microsoft.graph.profile entity.
func (m *ItemProfileRequestBuilder) CertificationsById(id string)(*PersonCertificationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["personCertification%2Did"] = id
    }
    return NewPersonCertificationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewItemProfileRequestBuilderInternal instantiates a new ProfileRequestBuilder and sets the default values.
func NewItemProfileRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemProfileRequestBuilder) {
    m := &ItemProfileRequestBuilder{
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
// NewItemProfileRequestBuilder instantiates a new ProfileRequestBuilder and sets the default values.
func NewItemProfileRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemProfileRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemProfileRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation deletes a profile object from a user's account.
func (m *ItemProfileRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemProfileRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation retrieve the properties and relationships of a profile object for a given user. The **profile** resource exposes various rich properties that are descriptive of the user as relationships, for example, anniversaries and education activities. To get one of these navigation properties, use the corresponding GET method on that property. See the methods exposed by **profile**.
func (m *ItemProfileRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *ItemProfileRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property profile in users
func (m *ItemProfileRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Profileable, requestConfiguration *ItemProfileRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete deletes a profile object from a user's account.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/profile-delete?view=graph-rest-1.0
func (m *ItemProfileRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemProfileRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *ItemProfileRequestBuilder) EducationalActivities()(*ItemProfileEducationalActivitiesRequestBuilder) {
    return NewItemProfileEducationalActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EducationalActivitiesById provides operations to manage the educationalActivities property of the microsoft.graph.profile entity.
func (m *ItemProfileRequestBuilder) EducationalActivitiesById(id string)(*EducationalActivityItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationalActivity%2Did"] = id
    }
    return NewEducationalActivityItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Emails provides operations to manage the emails property of the microsoft.graph.profile entity.
func (m *ItemProfileRequestBuilder) Emails()(*ItemProfileEmailsRequestBuilder) {
    return NewItemProfileEmailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EmailsById provides operations to manage the emails property of the microsoft.graph.profile entity.
func (m *ItemProfileRequestBuilder) EmailsById(id string)(*ItemEmailItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemEmail%2Did"] = id
    }
    return NewItemEmailItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get retrieve the properties and relationships of a profile object for a given user. The **profile** resource exposes various rich properties that are descriptive of the user as relationships, for example, anniversaries and education activities. To get one of these navigation properties, use the corresponding GET method on that property. See the methods exposed by **profile**.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/profile-get?view=graph-rest-1.0
func (m *ItemProfileRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemProfileRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Profileable, error) {
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
func (m *ItemProfileRequestBuilder) Interests()(*ItemProfileInterestsRequestBuilder) {
    return NewItemProfileInterestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InterestsById provides operations to manage the interests property of the microsoft.graph.profile entity.
func (m *ItemProfileRequestBuilder) InterestsById(id string)(*PersonInterestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["personInterest%2Did"] = id
    }
    return NewPersonInterestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Languages provides operations to manage the languages property of the microsoft.graph.profile entity.
func (m *ItemProfileRequestBuilder) Languages()(*ItemProfileLanguagesRequestBuilder) {
    return NewItemProfileLanguagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// LanguagesById provides operations to manage the languages property of the microsoft.graph.profile entity.
func (m *ItemProfileRequestBuilder) LanguagesById(id string)(*LanguageProficiencyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["languageProficiency%2Did"] = id
    }
    return NewLanguageProficiencyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Names provides operations to manage the names property of the microsoft.graph.profile entity.
func (m *ItemProfileRequestBuilder) Names()(*ItemProfileNamesRequestBuilder) {
    return NewItemProfileNamesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NamesById provides operations to manage the names property of the microsoft.graph.profile entity.
func (m *ItemProfileRequestBuilder) NamesById(id string)(*PersonNameItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["personName%2Did"] = id
    }
    return NewPersonNameItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Notes provides operations to manage the notes property of the microsoft.graph.profile entity.
func (m *ItemProfileRequestBuilder) Notes()(*ItemProfileNotesRequestBuilder) {
    return NewItemProfileNotesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NotesById provides operations to manage the notes property of the microsoft.graph.profile entity.
func (m *ItemProfileRequestBuilder) NotesById(id string)(*PersonAnnotationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["personAnnotation%2Did"] = id
    }
    return NewPersonAnnotationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property profile in users
func (m *ItemProfileRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Profileable, requestConfiguration *ItemProfileRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Profileable, error) {
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
func (m *ItemProfileRequestBuilder) Patents()(*ItemProfilePatentsRequestBuilder) {
    return NewItemProfilePatentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PatentsById provides operations to manage the patents property of the microsoft.graph.profile entity.
func (m *ItemProfileRequestBuilder) PatentsById(id string)(*ItemPatentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemPatent%2Did"] = id
    }
    return NewItemPatentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Phones provides operations to manage the phones property of the microsoft.graph.profile entity.
func (m *ItemProfileRequestBuilder) Phones()(*ItemProfilePhonesRequestBuilder) {
    return NewItemProfilePhonesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PhonesById provides operations to manage the phones property of the microsoft.graph.profile entity.
func (m *ItemProfileRequestBuilder) PhonesById(id string)(*ItemPhoneItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemPhone%2Did"] = id
    }
    return NewItemPhoneItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Positions provides operations to manage the positions property of the microsoft.graph.profile entity.
func (m *ItemProfileRequestBuilder) Positions()(*ItemProfilePositionsRequestBuilder) {
    return NewItemProfilePositionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PositionsById provides operations to manage the positions property of the microsoft.graph.profile entity.
func (m *ItemProfileRequestBuilder) PositionsById(id string)(*WorkPositionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["workPosition%2Did"] = id
    }
    return NewWorkPositionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Projects provides operations to manage the projects property of the microsoft.graph.profile entity.
func (m *ItemProfileRequestBuilder) Projects()(*ItemProfileProjectsRequestBuilder) {
    return NewItemProfileProjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ProjectsById provides operations to manage the projects property of the microsoft.graph.profile entity.
func (m *ItemProfileRequestBuilder) ProjectsById(id string)(*ProjectParticipationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["projectParticipation%2Did"] = id
    }
    return NewProjectParticipationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Publications provides operations to manage the publications property of the microsoft.graph.profile entity.
func (m *ItemProfileRequestBuilder) Publications()(*ItemProfilePublicationsRequestBuilder) {
    return NewItemProfilePublicationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PublicationsById provides operations to manage the publications property of the microsoft.graph.profile entity.
func (m *ItemProfileRequestBuilder) PublicationsById(id string)(*ItemPublicationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemPublication%2Did"] = id
    }
    return NewItemPublicationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Skills provides operations to manage the skills property of the microsoft.graph.profile entity.
func (m *ItemProfileRequestBuilder) Skills()(*ItemProfileSkillsRequestBuilder) {
    return NewItemProfileSkillsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SkillsById provides operations to manage the skills property of the microsoft.graph.profile entity.
func (m *ItemProfileRequestBuilder) SkillsById(id string)(*SkillProficiencyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["skillProficiency%2Did"] = id
    }
    return NewSkillProficiencyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// WebAccounts provides operations to manage the webAccounts property of the microsoft.graph.profile entity.
func (m *ItemProfileRequestBuilder) WebAccounts()(*ItemProfileWebAccountsRequestBuilder) {
    return NewItemProfileWebAccountsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WebAccountsById provides operations to manage the webAccounts property of the microsoft.graph.profile entity.
func (m *ItemProfileRequestBuilder) WebAccountsById(id string)(*WebAccountItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["webAccount%2Did"] = id
    }
    return NewWebAccountItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Websites provides operations to manage the websites property of the microsoft.graph.profile entity.
func (m *ItemProfileRequestBuilder) Websites()(*ItemProfileWebsitesRequestBuilder) {
    return NewItemProfileWebsitesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WebsitesById provides operations to manage the websites property of the microsoft.graph.profile entity.
func (m *ItemProfileRequestBuilder) WebsitesById(id string)(*PersonWebsiteItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["personWebsite%2Did"] = id
    }
    return NewPersonWebsiteItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}

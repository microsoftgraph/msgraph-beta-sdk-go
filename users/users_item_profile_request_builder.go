package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UsersItemProfileRequestBuilder provides operations to manage the profile property of the microsoft.graph.user entity.
type UsersItemProfileRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// UsersItemProfileRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemProfileRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// UsersItemProfileRequestBuilderGetQueryParameters retrieve the properties and relationships of a profile object for a given user. The **profile** resource exposes various rich properties that are descriptive of the user as relationships, for example, anniversaries and education activities. To get one of these navigation properties, use the corresponding GET method on that property. See the methods exposed by **profile**.
type UsersItemProfileRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UsersItemProfileRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemProfileRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UsersItemProfileRequestBuilderGetQueryParameters
}
// UsersItemProfileRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemProfileRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Account provides operations to manage the account property of the microsoft.graph.profile entity.
func (m *UsersItemProfileRequestBuilder) Account()(*UsersItemProfileAccountRequestBuilder) {
    return NewUsersItemProfileAccountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccountById provides operations to manage the account property of the microsoft.graph.profile entity.
func (m *UsersItemProfileRequestBuilder) AccountById(id string)(*UsersItemProfileAccountUserAccountInformationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userAccountInformation%2Did"] = id
    }
    return NewUsersItemProfileAccountUserAccountInformationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Addresses provides operations to manage the addresses property of the microsoft.graph.profile entity.
func (m *UsersItemProfileRequestBuilder) Addresses()(*UsersItemProfileAddressesRequestBuilder) {
    return NewUsersItemProfileAddressesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AddressesById provides operations to manage the addresses property of the microsoft.graph.profile entity.
func (m *UsersItemProfileRequestBuilder) AddressesById(id string)(*UsersItemProfileAddressesItemAddressItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemAddress%2Did"] = id
    }
    return NewUsersItemProfileAddressesItemAddressItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Anniversaries provides operations to manage the anniversaries property of the microsoft.graph.profile entity.
func (m *UsersItemProfileRequestBuilder) Anniversaries()(*UsersItemProfileAnniversariesRequestBuilder) {
    return NewUsersItemProfileAnniversariesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AnniversariesById provides operations to manage the anniversaries property of the microsoft.graph.profile entity.
func (m *UsersItemProfileRequestBuilder) AnniversariesById(id string)(*UsersItemProfileAnniversariesPersonAnnualEventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["personAnnualEvent%2Did"] = id
    }
    return NewUsersItemProfileAnniversariesPersonAnnualEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Awards provides operations to manage the awards property of the microsoft.graph.profile entity.
func (m *UsersItemProfileRequestBuilder) Awards()(*UsersItemProfileAwardsRequestBuilder) {
    return NewUsersItemProfileAwardsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AwardsById provides operations to manage the awards property of the microsoft.graph.profile entity.
func (m *UsersItemProfileRequestBuilder) AwardsById(id string)(*UsersItemProfileAwardsPersonAwardItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["personAward%2Did"] = id
    }
    return NewUsersItemProfileAwardsPersonAwardItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Certifications provides operations to manage the certifications property of the microsoft.graph.profile entity.
func (m *UsersItemProfileRequestBuilder) Certifications()(*UsersItemProfileCertificationsRequestBuilder) {
    return NewUsersItemProfileCertificationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CertificationsById provides operations to manage the certifications property of the microsoft.graph.profile entity.
func (m *UsersItemProfileRequestBuilder) CertificationsById(id string)(*UsersItemProfileCertificationsPersonCertificationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["personCertification%2Did"] = id
    }
    return NewUsersItemProfileCertificationsPersonCertificationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewUsersItemProfileRequestBuilderInternal instantiates a new ProfileRequestBuilder and sets the default values.
func NewUsersItemProfileRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemProfileRequestBuilder) {
    m := &UsersItemProfileRequestBuilder{
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
// NewUsersItemProfileRequestBuilder instantiates a new ProfileRequestBuilder and sets the default values.
func NewUsersItemProfileRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemProfileRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUsersItemProfileRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation deletes a profile object from a user's account.
func (m *UsersItemProfileRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *UsersItemProfileRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *UsersItemProfileRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *UsersItemProfileRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *UsersItemProfileRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Profileable, requestConfiguration *UsersItemProfileRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *UsersItemProfileRequestBuilder) Delete(ctx context.Context, requestConfiguration *UsersItemProfileRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *UsersItemProfileRequestBuilder) EducationalActivities()(*UsersItemProfileEducationalActivitiesRequestBuilder) {
    return NewUsersItemProfileEducationalActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EducationalActivitiesById provides operations to manage the educationalActivities property of the microsoft.graph.profile entity.
func (m *UsersItemProfileRequestBuilder) EducationalActivitiesById(id string)(*UsersItemProfileEducationalActivitiesEducationalActivityItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationalActivity%2Did"] = id
    }
    return NewUsersItemProfileEducationalActivitiesEducationalActivityItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Emails provides operations to manage the emails property of the microsoft.graph.profile entity.
func (m *UsersItemProfileRequestBuilder) Emails()(*UsersItemProfileEmailsRequestBuilder) {
    return NewUsersItemProfileEmailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EmailsById provides operations to manage the emails property of the microsoft.graph.profile entity.
func (m *UsersItemProfileRequestBuilder) EmailsById(id string)(*UsersItemProfileEmailsItemEmailItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemEmail%2Did"] = id
    }
    return NewUsersItemProfileEmailsItemEmailItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get retrieve the properties and relationships of a profile object for a given user. The **profile** resource exposes various rich properties that are descriptive of the user as relationships, for example, anniversaries and education activities. To get one of these navigation properties, use the corresponding GET method on that property. See the methods exposed by **profile**.
func (m *UsersItemProfileRequestBuilder) Get(ctx context.Context, requestConfiguration *UsersItemProfileRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Profileable, error) {
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
func (m *UsersItemProfileRequestBuilder) Interests()(*UsersItemProfileInterestsRequestBuilder) {
    return NewUsersItemProfileInterestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InterestsById provides operations to manage the interests property of the microsoft.graph.profile entity.
func (m *UsersItemProfileRequestBuilder) InterestsById(id string)(*UsersItemProfileInterestsPersonInterestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["personInterest%2Did"] = id
    }
    return NewUsersItemProfileInterestsPersonInterestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Languages provides operations to manage the languages property of the microsoft.graph.profile entity.
func (m *UsersItemProfileRequestBuilder) Languages()(*UsersItemProfileLanguagesRequestBuilder) {
    return NewUsersItemProfileLanguagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// LanguagesById provides operations to manage the languages property of the microsoft.graph.profile entity.
func (m *UsersItemProfileRequestBuilder) LanguagesById(id string)(*UsersItemProfileLanguagesLanguageProficiencyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["languageProficiency%2Did"] = id
    }
    return NewUsersItemProfileLanguagesLanguageProficiencyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Names provides operations to manage the names property of the microsoft.graph.profile entity.
func (m *UsersItemProfileRequestBuilder) Names()(*UsersItemProfileNamesRequestBuilder) {
    return NewUsersItemProfileNamesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NamesById provides operations to manage the names property of the microsoft.graph.profile entity.
func (m *UsersItemProfileRequestBuilder) NamesById(id string)(*UsersItemProfileNamesPersonNameItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["personName%2Did"] = id
    }
    return NewUsersItemProfileNamesPersonNameItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Notes provides operations to manage the notes property of the microsoft.graph.profile entity.
func (m *UsersItemProfileRequestBuilder) Notes()(*UsersItemProfileNotesRequestBuilder) {
    return NewUsersItemProfileNotesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NotesById provides operations to manage the notes property of the microsoft.graph.profile entity.
func (m *UsersItemProfileRequestBuilder) NotesById(id string)(*UsersItemProfileNotesPersonAnnotationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["personAnnotation%2Did"] = id
    }
    return NewUsersItemProfileNotesPersonAnnotationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property profile in users
func (m *UsersItemProfileRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Profileable, requestConfiguration *UsersItemProfileRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Profileable, error) {
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
func (m *UsersItemProfileRequestBuilder) Patents()(*UsersItemProfilePatentsRequestBuilder) {
    return NewUsersItemProfilePatentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PatentsById provides operations to manage the patents property of the microsoft.graph.profile entity.
func (m *UsersItemProfileRequestBuilder) PatentsById(id string)(*UsersItemProfilePatentsItemPatentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemPatent%2Did"] = id
    }
    return NewUsersItemProfilePatentsItemPatentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Phones provides operations to manage the phones property of the microsoft.graph.profile entity.
func (m *UsersItemProfileRequestBuilder) Phones()(*UsersItemProfilePhonesRequestBuilder) {
    return NewUsersItemProfilePhonesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PhonesById provides operations to manage the phones property of the microsoft.graph.profile entity.
func (m *UsersItemProfileRequestBuilder) PhonesById(id string)(*UsersItemProfilePhonesItemPhoneItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemPhone%2Did"] = id
    }
    return NewUsersItemProfilePhonesItemPhoneItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Positions provides operations to manage the positions property of the microsoft.graph.profile entity.
func (m *UsersItemProfileRequestBuilder) Positions()(*UsersItemProfilePositionsRequestBuilder) {
    return NewUsersItemProfilePositionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PositionsById provides operations to manage the positions property of the microsoft.graph.profile entity.
func (m *UsersItemProfileRequestBuilder) PositionsById(id string)(*UsersItemProfilePositionsWorkPositionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["workPosition%2Did"] = id
    }
    return NewUsersItemProfilePositionsWorkPositionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Projects provides operations to manage the projects property of the microsoft.graph.profile entity.
func (m *UsersItemProfileRequestBuilder) Projects()(*UsersItemProfileProjectsRequestBuilder) {
    return NewUsersItemProfileProjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ProjectsById provides operations to manage the projects property of the microsoft.graph.profile entity.
func (m *UsersItemProfileRequestBuilder) ProjectsById(id string)(*UsersItemProfileProjectsProjectParticipationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["projectParticipation%2Did"] = id
    }
    return NewUsersItemProfileProjectsProjectParticipationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Publications provides operations to manage the publications property of the microsoft.graph.profile entity.
func (m *UsersItemProfileRequestBuilder) Publications()(*UsersItemProfilePublicationsRequestBuilder) {
    return NewUsersItemProfilePublicationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PublicationsById provides operations to manage the publications property of the microsoft.graph.profile entity.
func (m *UsersItemProfileRequestBuilder) PublicationsById(id string)(*UsersItemProfilePublicationsItemPublicationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemPublication%2Did"] = id
    }
    return NewUsersItemProfilePublicationsItemPublicationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Skills provides operations to manage the skills property of the microsoft.graph.profile entity.
func (m *UsersItemProfileRequestBuilder) Skills()(*UsersItemProfileSkillsRequestBuilder) {
    return NewUsersItemProfileSkillsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SkillsById provides operations to manage the skills property of the microsoft.graph.profile entity.
func (m *UsersItemProfileRequestBuilder) SkillsById(id string)(*UsersItemProfileSkillsSkillProficiencyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["skillProficiency%2Did"] = id
    }
    return NewUsersItemProfileSkillsSkillProficiencyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// WebAccounts provides operations to manage the webAccounts property of the microsoft.graph.profile entity.
func (m *UsersItemProfileRequestBuilder) WebAccounts()(*UsersItemProfileWebAccountsRequestBuilder) {
    return NewUsersItemProfileWebAccountsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WebAccountsById provides operations to manage the webAccounts property of the microsoft.graph.profile entity.
func (m *UsersItemProfileRequestBuilder) WebAccountsById(id string)(*UsersItemProfileWebAccountsWebAccountItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["webAccount%2Did"] = id
    }
    return NewUsersItemProfileWebAccountsWebAccountItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Websites provides operations to manage the websites property of the microsoft.graph.profile entity.
func (m *UsersItemProfileRequestBuilder) Websites()(*UsersItemProfileWebsitesRequestBuilder) {
    return NewUsersItemProfileWebsitesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WebsitesById provides operations to manage the websites property of the microsoft.graph.profile entity.
func (m *UsersItemProfileRequestBuilder) WebsitesById(id string)(*UsersItemProfileWebsitesPersonWebsiteItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["personWebsite%2Did"] = id
    }
    return NewUsersItemProfileWebsitesPersonWebsiteItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}

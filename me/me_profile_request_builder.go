package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MeProfileRequestBuilder provides operations to manage the profile property of the microsoft.graph.user entity.
type MeProfileRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MeProfileRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MeProfileRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// MeProfileRequestBuilderGetQueryParameters retrieve the properties and relationships of a profile object for a given user. The **profile** resource exposes various rich properties that are descriptive of the user as relationships, for example, anniversaries and education activities. To get one of these navigation properties, use the corresponding GET method on that property. See the methods exposed by **profile**.
type MeProfileRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MeProfileRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MeProfileRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MeProfileRequestBuilderGetQueryParameters
}
// MeProfileRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MeProfileRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Account provides operations to manage the account property of the microsoft.graph.profile entity.
func (m *MeProfileRequestBuilder) Account()(*MeProfileAccountRequestBuilder) {
    return NewMeProfileAccountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccountById provides operations to manage the account property of the microsoft.graph.profile entity.
func (m *MeProfileRequestBuilder) AccountById(id string)(*MeProfileAccountUserAccountInformationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userAccountInformation%2Did"] = id
    }
    return NewMeProfileAccountUserAccountInformationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Addresses provides operations to manage the addresses property of the microsoft.graph.profile entity.
func (m *MeProfileRequestBuilder) Addresses()(*MeProfileAddressesRequestBuilder) {
    return NewMeProfileAddressesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AddressesById provides operations to manage the addresses property of the microsoft.graph.profile entity.
func (m *MeProfileRequestBuilder) AddressesById(id string)(*MeProfileAddressesItemAddressItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemAddress%2Did"] = id
    }
    return NewMeProfileAddressesItemAddressItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Anniversaries provides operations to manage the anniversaries property of the microsoft.graph.profile entity.
func (m *MeProfileRequestBuilder) Anniversaries()(*MeProfileAnniversariesRequestBuilder) {
    return NewMeProfileAnniversariesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AnniversariesById provides operations to manage the anniversaries property of the microsoft.graph.profile entity.
func (m *MeProfileRequestBuilder) AnniversariesById(id string)(*MeProfileAnniversariesPersonAnnualEventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["personAnnualEvent%2Did"] = id
    }
    return NewMeProfileAnniversariesPersonAnnualEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Awards provides operations to manage the awards property of the microsoft.graph.profile entity.
func (m *MeProfileRequestBuilder) Awards()(*MeProfileAwardsRequestBuilder) {
    return NewMeProfileAwardsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AwardsById provides operations to manage the awards property of the microsoft.graph.profile entity.
func (m *MeProfileRequestBuilder) AwardsById(id string)(*MeProfileAwardsPersonAwardItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["personAward%2Did"] = id
    }
    return NewMeProfileAwardsPersonAwardItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Certifications provides operations to manage the certifications property of the microsoft.graph.profile entity.
func (m *MeProfileRequestBuilder) Certifications()(*MeProfileCertificationsRequestBuilder) {
    return NewMeProfileCertificationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CertificationsById provides operations to manage the certifications property of the microsoft.graph.profile entity.
func (m *MeProfileRequestBuilder) CertificationsById(id string)(*MeProfileCertificationsPersonCertificationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["personCertification%2Did"] = id
    }
    return NewMeProfileCertificationsPersonCertificationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewMeProfileRequestBuilderInternal instantiates a new ProfileRequestBuilder and sets the default values.
func NewMeProfileRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeProfileRequestBuilder) {
    m := &MeProfileRequestBuilder{
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
// NewMeProfileRequestBuilder instantiates a new ProfileRequestBuilder and sets the default values.
func NewMeProfileRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeProfileRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMeProfileRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation deletes a profile object from a user's account.
func (m *MeProfileRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *MeProfileRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *MeProfileRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *MeProfileRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *MeProfileRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Profileable, requestConfiguration *MeProfileRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *MeProfileRequestBuilder) Delete(ctx context.Context, requestConfiguration *MeProfileRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *MeProfileRequestBuilder) EducationalActivities()(*MeProfileEducationalActivitiesRequestBuilder) {
    return NewMeProfileEducationalActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EducationalActivitiesById provides operations to manage the educationalActivities property of the microsoft.graph.profile entity.
func (m *MeProfileRequestBuilder) EducationalActivitiesById(id string)(*MeProfileEducationalActivitiesEducationalActivityItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationalActivity%2Did"] = id
    }
    return NewMeProfileEducationalActivitiesEducationalActivityItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Emails provides operations to manage the emails property of the microsoft.graph.profile entity.
func (m *MeProfileRequestBuilder) Emails()(*MeProfileEmailsRequestBuilder) {
    return NewMeProfileEmailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EmailsById provides operations to manage the emails property of the microsoft.graph.profile entity.
func (m *MeProfileRequestBuilder) EmailsById(id string)(*MeProfileEmailsItemEmailItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemEmail%2Did"] = id
    }
    return NewMeProfileEmailsItemEmailItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get retrieve the properties and relationships of a profile object for a given user. The **profile** resource exposes various rich properties that are descriptive of the user as relationships, for example, anniversaries and education activities. To get one of these navigation properties, use the corresponding GET method on that property. See the methods exposed by **profile**.
func (m *MeProfileRequestBuilder) Get(ctx context.Context, requestConfiguration *MeProfileRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Profileable, error) {
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
func (m *MeProfileRequestBuilder) Interests()(*MeProfileInterestsRequestBuilder) {
    return NewMeProfileInterestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InterestsById provides operations to manage the interests property of the microsoft.graph.profile entity.
func (m *MeProfileRequestBuilder) InterestsById(id string)(*MeProfileInterestsPersonInterestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["personInterest%2Did"] = id
    }
    return NewMeProfileInterestsPersonInterestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Languages provides operations to manage the languages property of the microsoft.graph.profile entity.
func (m *MeProfileRequestBuilder) Languages()(*MeProfileLanguagesRequestBuilder) {
    return NewMeProfileLanguagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// LanguagesById provides operations to manage the languages property of the microsoft.graph.profile entity.
func (m *MeProfileRequestBuilder) LanguagesById(id string)(*MeProfileLanguagesLanguageProficiencyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["languageProficiency%2Did"] = id
    }
    return NewMeProfileLanguagesLanguageProficiencyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Names provides operations to manage the names property of the microsoft.graph.profile entity.
func (m *MeProfileRequestBuilder) Names()(*MeProfileNamesRequestBuilder) {
    return NewMeProfileNamesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NamesById provides operations to manage the names property of the microsoft.graph.profile entity.
func (m *MeProfileRequestBuilder) NamesById(id string)(*MeProfileNamesPersonNameItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["personName%2Did"] = id
    }
    return NewMeProfileNamesPersonNameItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Notes provides operations to manage the notes property of the microsoft.graph.profile entity.
func (m *MeProfileRequestBuilder) Notes()(*MeProfileNotesRequestBuilder) {
    return NewMeProfileNotesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NotesById provides operations to manage the notes property of the microsoft.graph.profile entity.
func (m *MeProfileRequestBuilder) NotesById(id string)(*MeProfileNotesPersonAnnotationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["personAnnotation%2Did"] = id
    }
    return NewMeProfileNotesPersonAnnotationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property profile in me
func (m *MeProfileRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Profileable, requestConfiguration *MeProfileRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Profileable, error) {
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
func (m *MeProfileRequestBuilder) Patents()(*MeProfilePatentsRequestBuilder) {
    return NewMeProfilePatentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PatentsById provides operations to manage the patents property of the microsoft.graph.profile entity.
func (m *MeProfileRequestBuilder) PatentsById(id string)(*MeProfilePatentsItemPatentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemPatent%2Did"] = id
    }
    return NewMeProfilePatentsItemPatentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Phones provides operations to manage the phones property of the microsoft.graph.profile entity.
func (m *MeProfileRequestBuilder) Phones()(*MeProfilePhonesRequestBuilder) {
    return NewMeProfilePhonesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PhonesById provides operations to manage the phones property of the microsoft.graph.profile entity.
func (m *MeProfileRequestBuilder) PhonesById(id string)(*MeProfilePhonesItemPhoneItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemPhone%2Did"] = id
    }
    return NewMeProfilePhonesItemPhoneItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Positions provides operations to manage the positions property of the microsoft.graph.profile entity.
func (m *MeProfileRequestBuilder) Positions()(*MeProfilePositionsRequestBuilder) {
    return NewMeProfilePositionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PositionsById provides operations to manage the positions property of the microsoft.graph.profile entity.
func (m *MeProfileRequestBuilder) PositionsById(id string)(*MeProfilePositionsWorkPositionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["workPosition%2Did"] = id
    }
    return NewMeProfilePositionsWorkPositionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Projects provides operations to manage the projects property of the microsoft.graph.profile entity.
func (m *MeProfileRequestBuilder) Projects()(*MeProfileProjectsRequestBuilder) {
    return NewMeProfileProjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ProjectsById provides operations to manage the projects property of the microsoft.graph.profile entity.
func (m *MeProfileRequestBuilder) ProjectsById(id string)(*MeProfileProjectsProjectParticipationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["projectParticipation%2Did"] = id
    }
    return NewMeProfileProjectsProjectParticipationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Publications provides operations to manage the publications property of the microsoft.graph.profile entity.
func (m *MeProfileRequestBuilder) Publications()(*MeProfilePublicationsRequestBuilder) {
    return NewMeProfilePublicationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PublicationsById provides operations to manage the publications property of the microsoft.graph.profile entity.
func (m *MeProfileRequestBuilder) PublicationsById(id string)(*MeProfilePublicationsItemPublicationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemPublication%2Did"] = id
    }
    return NewMeProfilePublicationsItemPublicationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Skills provides operations to manage the skills property of the microsoft.graph.profile entity.
func (m *MeProfileRequestBuilder) Skills()(*MeProfileSkillsRequestBuilder) {
    return NewMeProfileSkillsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SkillsById provides operations to manage the skills property of the microsoft.graph.profile entity.
func (m *MeProfileRequestBuilder) SkillsById(id string)(*MeProfileSkillsSkillProficiencyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["skillProficiency%2Did"] = id
    }
    return NewMeProfileSkillsSkillProficiencyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// WebAccounts provides operations to manage the webAccounts property of the microsoft.graph.profile entity.
func (m *MeProfileRequestBuilder) WebAccounts()(*MeProfileWebAccountsRequestBuilder) {
    return NewMeProfileWebAccountsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WebAccountsById provides operations to manage the webAccounts property of the microsoft.graph.profile entity.
func (m *MeProfileRequestBuilder) WebAccountsById(id string)(*MeProfileWebAccountsWebAccountItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["webAccount%2Did"] = id
    }
    return NewMeProfileWebAccountsWebAccountItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Websites provides operations to manage the websites property of the microsoft.graph.profile entity.
func (m *MeProfileRequestBuilder) Websites()(*MeProfileWebsitesRequestBuilder) {
    return NewMeProfileWebsitesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WebsitesById provides operations to manage the websites property of the microsoft.graph.profile entity.
func (m *MeProfileRequestBuilder) WebsitesById(id string)(*MeProfileWebsitesPersonWebsiteItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["personWebsite%2Did"] = id
    }
    return NewMeProfileWebsitesPersonWebsiteItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}

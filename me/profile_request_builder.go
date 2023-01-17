package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
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
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
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
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ProfileRequestBuilderGetQueryParameters
}
// ProfileRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ProfileRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Account provides operations to manage the account property of the microsoft.graph.profile entity.
func (m *ProfileRequestBuilder) Account()(*ProfileAccountRequestBuilder) {
    return NewProfileAccountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccountById provides operations to manage the account property of the microsoft.graph.profile entity.
func (m *ProfileRequestBuilder) AccountById(id string)(*ProfileAccountUserAccountInformationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userAccountInformation%2Did"] = id
    }
    return NewProfileAccountUserAccountInformationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Addresses provides operations to manage the addresses property of the microsoft.graph.profile entity.
func (m *ProfileRequestBuilder) Addresses()(*ProfileAddressesRequestBuilder) {
    return NewProfileAddressesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AddressesById provides operations to manage the addresses property of the microsoft.graph.profile entity.
func (m *ProfileRequestBuilder) AddressesById(id string)(*ProfileAddressesItemAddressItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemAddress%2Did"] = id
    }
    return NewProfileAddressesItemAddressItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Anniversaries provides operations to manage the anniversaries property of the microsoft.graph.profile entity.
func (m *ProfileRequestBuilder) Anniversaries()(*ProfileAnniversariesRequestBuilder) {
    return NewProfileAnniversariesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AnniversariesById provides operations to manage the anniversaries property of the microsoft.graph.profile entity.
func (m *ProfileRequestBuilder) AnniversariesById(id string)(*ProfileAnniversariesPersonAnnualEventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["personAnnualEvent%2Did"] = id
    }
    return NewProfileAnniversariesPersonAnnualEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Awards provides operations to manage the awards property of the microsoft.graph.profile entity.
func (m *ProfileRequestBuilder) Awards()(*ProfileAwardsRequestBuilder) {
    return NewProfileAwardsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AwardsById provides operations to manage the awards property of the microsoft.graph.profile entity.
func (m *ProfileRequestBuilder) AwardsById(id string)(*ProfileAwardsPersonAwardItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["personAward%2Did"] = id
    }
    return NewProfileAwardsPersonAwardItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Certifications provides operations to manage the certifications property of the microsoft.graph.profile entity.
func (m *ProfileRequestBuilder) Certifications()(*ProfileCertificationsRequestBuilder) {
    return NewProfileCertificationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CertificationsById provides operations to manage the certifications property of the microsoft.graph.profile entity.
func (m *ProfileRequestBuilder) CertificationsById(id string)(*ProfileCertificationsPersonCertificationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["personCertification%2Did"] = id
    }
    return NewProfileCertificationsPersonCertificationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
// Delete deletes a profile object from a user's account.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/profile-delete?view=graph-rest-1.0
func (m *ProfileRequestBuilder) Delete(ctx context.Context, requestConfiguration *ProfileRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// EducationalActivities provides operations to manage the educationalActivities property of the microsoft.graph.profile entity.
func (m *ProfileRequestBuilder) EducationalActivities()(*ProfileEducationalActivitiesRequestBuilder) {
    return NewProfileEducationalActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EducationalActivitiesById provides operations to manage the educationalActivities property of the microsoft.graph.profile entity.
func (m *ProfileRequestBuilder) EducationalActivitiesById(id string)(*ProfileEducationalActivitiesEducationalActivityItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationalActivity%2Did"] = id
    }
    return NewProfileEducationalActivitiesEducationalActivityItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Emails provides operations to manage the emails property of the microsoft.graph.profile entity.
func (m *ProfileRequestBuilder) Emails()(*ProfileEmailsRequestBuilder) {
    return NewProfileEmailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EmailsById provides operations to manage the emails property of the microsoft.graph.profile entity.
func (m *ProfileRequestBuilder) EmailsById(id string)(*ProfileEmailsItemEmailItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemEmail%2Did"] = id
    }
    return NewProfileEmailsItemEmailItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get retrieve the properties and relationships of a profile object for a given user. The **profile** resource exposes various rich properties that are descriptive of the user as relationships, for example, anniversaries and education activities. To get one of these navigation properties, use the corresponding GET method on that property. See the methods exposed by **profile**.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/profile-get?view=graph-rest-1.0
func (m *ProfileRequestBuilder) Get(ctx context.Context, requestConfiguration *ProfileRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Profileable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateProfileFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Profileable), nil
}
// Interests provides operations to manage the interests property of the microsoft.graph.profile entity.
func (m *ProfileRequestBuilder) Interests()(*ProfileInterestsRequestBuilder) {
    return NewProfileInterestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InterestsById provides operations to manage the interests property of the microsoft.graph.profile entity.
func (m *ProfileRequestBuilder) InterestsById(id string)(*ProfileInterestsPersonInterestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["personInterest%2Did"] = id
    }
    return NewProfileInterestsPersonInterestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Languages provides operations to manage the languages property of the microsoft.graph.profile entity.
func (m *ProfileRequestBuilder) Languages()(*ProfileLanguagesRequestBuilder) {
    return NewProfileLanguagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// LanguagesById provides operations to manage the languages property of the microsoft.graph.profile entity.
func (m *ProfileRequestBuilder) LanguagesById(id string)(*ProfileLanguagesLanguageProficiencyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["languageProficiency%2Did"] = id
    }
    return NewProfileLanguagesLanguageProficiencyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Names provides operations to manage the names property of the microsoft.graph.profile entity.
func (m *ProfileRequestBuilder) Names()(*ProfileNamesRequestBuilder) {
    return NewProfileNamesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NamesById provides operations to manage the names property of the microsoft.graph.profile entity.
func (m *ProfileRequestBuilder) NamesById(id string)(*ProfileNamesPersonNameItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["personName%2Did"] = id
    }
    return NewProfileNamesPersonNameItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Notes provides operations to manage the notes property of the microsoft.graph.profile entity.
func (m *ProfileRequestBuilder) Notes()(*ProfileNotesRequestBuilder) {
    return NewProfileNotesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NotesById provides operations to manage the notes property of the microsoft.graph.profile entity.
func (m *ProfileRequestBuilder) NotesById(id string)(*ProfileNotesPersonAnnotationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["personAnnotation%2Did"] = id
    }
    return NewProfileNotesPersonAnnotationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property profile in me
func (m *ProfileRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Profileable, requestConfiguration *ProfileRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Profileable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateProfileFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Profileable), nil
}
// Patents provides operations to manage the patents property of the microsoft.graph.profile entity.
func (m *ProfileRequestBuilder) Patents()(*ProfilePatentsRequestBuilder) {
    return NewProfilePatentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PatentsById provides operations to manage the patents property of the microsoft.graph.profile entity.
func (m *ProfileRequestBuilder) PatentsById(id string)(*ProfilePatentsItemPatentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemPatent%2Did"] = id
    }
    return NewProfilePatentsItemPatentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Phones provides operations to manage the phones property of the microsoft.graph.profile entity.
func (m *ProfileRequestBuilder) Phones()(*ProfilePhonesRequestBuilder) {
    return NewProfilePhonesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PhonesById provides operations to manage the phones property of the microsoft.graph.profile entity.
func (m *ProfileRequestBuilder) PhonesById(id string)(*ProfilePhonesItemPhoneItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemPhone%2Did"] = id
    }
    return NewProfilePhonesItemPhoneItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Positions provides operations to manage the positions property of the microsoft.graph.profile entity.
func (m *ProfileRequestBuilder) Positions()(*ProfilePositionsRequestBuilder) {
    return NewProfilePositionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PositionsById provides operations to manage the positions property of the microsoft.graph.profile entity.
func (m *ProfileRequestBuilder) PositionsById(id string)(*ProfilePositionsWorkPositionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["workPosition%2Did"] = id
    }
    return NewProfilePositionsWorkPositionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Projects provides operations to manage the projects property of the microsoft.graph.profile entity.
func (m *ProfileRequestBuilder) Projects()(*ProfileProjectsRequestBuilder) {
    return NewProfileProjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ProjectsById provides operations to manage the projects property of the microsoft.graph.profile entity.
func (m *ProfileRequestBuilder) ProjectsById(id string)(*ProfileProjectsProjectParticipationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["projectParticipation%2Did"] = id
    }
    return NewProfileProjectsProjectParticipationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Publications provides operations to manage the publications property of the microsoft.graph.profile entity.
func (m *ProfileRequestBuilder) Publications()(*ProfilePublicationsRequestBuilder) {
    return NewProfilePublicationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PublicationsById provides operations to manage the publications property of the microsoft.graph.profile entity.
func (m *ProfileRequestBuilder) PublicationsById(id string)(*ProfilePublicationsItemPublicationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemPublication%2Did"] = id
    }
    return NewProfilePublicationsItemPublicationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Skills provides operations to manage the skills property of the microsoft.graph.profile entity.
func (m *ProfileRequestBuilder) Skills()(*ProfileSkillsRequestBuilder) {
    return NewProfileSkillsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SkillsById provides operations to manage the skills property of the microsoft.graph.profile entity.
func (m *ProfileRequestBuilder) SkillsById(id string)(*ProfileSkillsSkillProficiencyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["skillProficiency%2Did"] = id
    }
    return NewProfileSkillsSkillProficiencyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ToDeleteRequestInformation deletes a profile object from a user's account.
func (m *ProfileRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ProfileRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation retrieve the properties and relationships of a profile object for a given user. The **profile** resource exposes various rich properties that are descriptive of the user as relationships, for example, anniversaries and education activities. To get one of these navigation properties, use the corresponding GET method on that property. See the methods exposed by **profile**.
func (m *ProfileRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ProfileRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property profile in me
func (m *ProfileRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Profileable, requestConfiguration *ProfileRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WebAccounts provides operations to manage the webAccounts property of the microsoft.graph.profile entity.
func (m *ProfileRequestBuilder) WebAccounts()(*ProfileWebAccountsRequestBuilder) {
    return NewProfileWebAccountsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WebAccountsById provides operations to manage the webAccounts property of the microsoft.graph.profile entity.
func (m *ProfileRequestBuilder) WebAccountsById(id string)(*ProfileWebAccountsWebAccountItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["webAccount%2Did"] = id
    }
    return NewProfileWebAccountsWebAccountItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Websites provides operations to manage the websites property of the microsoft.graph.profile entity.
func (m *ProfileRequestBuilder) Websites()(*ProfileWebsitesRequestBuilder) {
    return NewProfileWebsitesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WebsitesById provides operations to manage the websites property of the microsoft.graph.profile entity.
func (m *ProfileRequestBuilder) WebsitesById(id string)(*ProfileWebsitesPersonWebsiteItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["personWebsite%2Did"] = id
    }
    return NewProfileWebsitesPersonWebsiteItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}

package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemProfileRequestBuilder provides operations to manage the profile property of the microsoft.graph.user entity.
type ItemProfileRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemProfileRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemProfileRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemProfileRequestBuilderGetQueryParameters represents properties that are descriptive of a user in a tenant.
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
// returns a *ItemProfileAccountRequestBuilder when successful
func (m *ItemProfileRequestBuilder) Account()(*ItemProfileAccountRequestBuilder) {
    return NewItemProfileAccountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Addresses provides operations to manage the addresses property of the microsoft.graph.profile entity.
// returns a *ItemProfileAddressesRequestBuilder when successful
func (m *ItemProfileRequestBuilder) Addresses()(*ItemProfileAddressesRequestBuilder) {
    return NewItemProfileAddressesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Anniversaries provides operations to manage the anniversaries property of the microsoft.graph.profile entity.
// returns a *ItemProfileAnniversariesRequestBuilder when successful
func (m *ItemProfileRequestBuilder) Anniversaries()(*ItemProfileAnniversariesRequestBuilder) {
    return NewItemProfileAnniversariesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Awards provides operations to manage the awards property of the microsoft.graph.profile entity.
// returns a *ItemProfileAwardsRequestBuilder when successful
func (m *ItemProfileRequestBuilder) Awards()(*ItemProfileAwardsRequestBuilder) {
    return NewItemProfileAwardsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Certifications provides operations to manage the certifications property of the microsoft.graph.profile entity.
// returns a *ItemProfileCertificationsRequestBuilder when successful
func (m *ItemProfileRequestBuilder) Certifications()(*ItemProfileCertificationsRequestBuilder) {
    return NewItemProfileCertificationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemProfileRequestBuilderInternal instantiates a new ItemProfileRequestBuilder and sets the default values.
func NewItemProfileRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemProfileRequestBuilder) {
    m := &ItemProfileRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/profile{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemProfileRequestBuilder instantiates a new ItemProfileRequestBuilder and sets the default values.
func NewItemProfileRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemProfileRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemProfileRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property profile for users
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemProfileRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemProfileRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// EducationalActivities provides operations to manage the educationalActivities property of the microsoft.graph.profile entity.
// returns a *ItemProfileEducationalactivitiesEducationalActivitiesRequestBuilder when successful
func (m *ItemProfileRequestBuilder) EducationalActivities()(*ItemProfileEducationalactivitiesEducationalActivitiesRequestBuilder) {
    return NewItemProfileEducationalactivitiesEducationalActivitiesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Emails provides operations to manage the emails property of the microsoft.graph.profile entity.
// returns a *ItemProfileEmailsRequestBuilder when successful
func (m *ItemProfileRequestBuilder) Emails()(*ItemProfileEmailsRequestBuilder) {
    return NewItemProfileEmailsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get represents properties that are descriptive of a user in a tenant.
// returns a Profileable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemProfileRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemProfileRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Profileable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateProfileFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Profileable), nil
}
// Interests provides operations to manage the interests property of the microsoft.graph.profile entity.
// returns a *ItemProfileInterestsRequestBuilder when successful
func (m *ItemProfileRequestBuilder) Interests()(*ItemProfileInterestsRequestBuilder) {
    return NewItemProfileInterestsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Languages provides operations to manage the languages property of the microsoft.graph.profile entity.
// returns a *ItemProfileLanguagesRequestBuilder when successful
func (m *ItemProfileRequestBuilder) Languages()(*ItemProfileLanguagesRequestBuilder) {
    return NewItemProfileLanguagesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Names provides operations to manage the names property of the microsoft.graph.profile entity.
// returns a *ItemProfileNamesRequestBuilder when successful
func (m *ItemProfileRequestBuilder) Names()(*ItemProfileNamesRequestBuilder) {
    return NewItemProfileNamesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Notes provides operations to manage the notes property of the microsoft.graph.profile entity.
// returns a *ItemProfileNotesRequestBuilder when successful
func (m *ItemProfileRequestBuilder) Notes()(*ItemProfileNotesRequestBuilder) {
    return NewItemProfileNotesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property profile in users
// returns a Profileable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemProfileRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Profileable, requestConfiguration *ItemProfileRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Profileable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateProfileFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Profileable), nil
}
// Patents provides operations to manage the patents property of the microsoft.graph.profile entity.
// returns a *ItemProfilePatentsRequestBuilder when successful
func (m *ItemProfileRequestBuilder) Patents()(*ItemProfilePatentsRequestBuilder) {
    return NewItemProfilePatentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Phones provides operations to manage the phones property of the microsoft.graph.profile entity.
// returns a *ItemProfilePhonesRequestBuilder when successful
func (m *ItemProfileRequestBuilder) Phones()(*ItemProfilePhonesRequestBuilder) {
    return NewItemProfilePhonesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Positions provides operations to manage the positions property of the microsoft.graph.profile entity.
// returns a *ItemProfilePositionsRequestBuilder when successful
func (m *ItemProfileRequestBuilder) Positions()(*ItemProfilePositionsRequestBuilder) {
    return NewItemProfilePositionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Projects provides operations to manage the projects property of the microsoft.graph.profile entity.
// returns a *ItemProfileProjectsRequestBuilder when successful
func (m *ItemProfileRequestBuilder) Projects()(*ItemProfileProjectsRequestBuilder) {
    return NewItemProfileProjectsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Publications provides operations to manage the publications property of the microsoft.graph.profile entity.
// returns a *ItemProfilePublicationsRequestBuilder when successful
func (m *ItemProfileRequestBuilder) Publications()(*ItemProfilePublicationsRequestBuilder) {
    return NewItemProfilePublicationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Skills provides operations to manage the skills property of the microsoft.graph.profile entity.
// returns a *ItemProfileSkillsRequestBuilder when successful
func (m *ItemProfileRequestBuilder) Skills()(*ItemProfileSkillsRequestBuilder) {
    return NewItemProfileSkillsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property profile for users
// returns a *RequestInformation when successful
func (m *ItemProfileRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemProfileRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation represents properties that are descriptive of a user in a tenant.
// returns a *RequestInformation when successful
func (m *ItemProfileRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemProfileRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPatchRequestInformation update the navigation property profile in users
// returns a *RequestInformation when successful
func (m *ItemProfileRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Profileable, requestConfiguration *ItemProfileRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WebAccounts provides operations to manage the webAccounts property of the microsoft.graph.profile entity.
// returns a *ItemProfileWebaccountsWebAccountsRequestBuilder when successful
func (m *ItemProfileRequestBuilder) WebAccounts()(*ItemProfileWebaccountsWebAccountsRequestBuilder) {
    return NewItemProfileWebaccountsWebAccountsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Websites provides operations to manage the websites property of the microsoft.graph.profile entity.
// returns a *ItemProfileWebsitesRequestBuilder when successful
func (m *ItemProfileRequestBuilder) Websites()(*ItemProfileWebsitesRequestBuilder) {
    return NewItemProfileWebsitesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemProfileRequestBuilder when successful
func (m *ItemProfileRequestBuilder) WithUrl(rawUrl string)(*ItemProfileRequestBuilder) {
    return NewItemProfileRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

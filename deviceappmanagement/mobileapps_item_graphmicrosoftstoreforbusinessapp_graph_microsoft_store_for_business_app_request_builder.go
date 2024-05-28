package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MobileappsItemGraphmicrosoftstoreforbusinessappGraphMicrosoftStoreForBusinessAppRequestBuilder casts the previous resource to microsoftStoreForBusinessApp.
type MobileappsItemGraphmicrosoftstoreforbusinessappGraphMicrosoftStoreForBusinessAppRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MobileappsItemGraphmicrosoftstoreforbusinessappGraphMicrosoftStoreForBusinessAppRequestBuilderGetQueryParameters get the item of type microsoft.graph.mobileApp as microsoft.graph.microsoftStoreForBusinessApp
type MobileappsItemGraphmicrosoftstoreforbusinessappGraphMicrosoftStoreForBusinessAppRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MobileappsItemGraphmicrosoftstoreforbusinessappGraphMicrosoftStoreForBusinessAppRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileappsItemGraphmicrosoftstoreforbusinessappGraphMicrosoftStoreForBusinessAppRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MobileappsItemGraphmicrosoftstoreforbusinessappGraphMicrosoftStoreForBusinessAppRequestBuilderGetQueryParameters
}
// Assignments provides operations to manage the assignments property of the microsoft.graph.mobileApp entity.
// returns a *MobileappsItemGraphmicrosoftstoreforbusinessappAssignmentsRequestBuilder when successful
func (m *MobileappsItemGraphmicrosoftstoreforbusinessappGraphMicrosoftStoreForBusinessAppRequestBuilder) Assignments()(*MobileappsItemGraphmicrosoftstoreforbusinessappAssignmentsRequestBuilder) {
    return NewMobileappsItemGraphmicrosoftstoreforbusinessappAssignmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Categories provides operations to manage the categories property of the microsoft.graph.mobileApp entity.
// returns a *MobileappsItemGraphmicrosoftstoreforbusinessappCategoriesRequestBuilder when successful
func (m *MobileappsItemGraphmicrosoftstoreforbusinessappGraphMicrosoftStoreForBusinessAppRequestBuilder) Categories()(*MobileappsItemGraphmicrosoftstoreforbusinessappCategoriesRequestBuilder) {
    return NewMobileappsItemGraphmicrosoftstoreforbusinessappCategoriesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewMobileappsItemGraphmicrosoftstoreforbusinessappGraphMicrosoftStoreForBusinessAppRequestBuilderInternal instantiates a new MobileappsItemGraphmicrosoftstoreforbusinessappGraphMicrosoftStoreForBusinessAppRequestBuilder and sets the default values.
func NewMobileappsItemGraphmicrosoftstoreforbusinessappGraphMicrosoftStoreForBusinessAppRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileappsItemGraphmicrosoftstoreforbusinessappGraphMicrosoftStoreForBusinessAppRequestBuilder) {
    m := &MobileappsItemGraphmicrosoftstoreforbusinessappGraphMicrosoftStoreForBusinessAppRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/mobileApps/{mobileApp%2Did}/graph.microsoftStoreForBusinessApp{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewMobileappsItemGraphmicrosoftstoreforbusinessappGraphMicrosoftStoreForBusinessAppRequestBuilder instantiates a new MobileappsItemGraphmicrosoftstoreforbusinessappGraphMicrosoftStoreForBusinessAppRequestBuilder and sets the default values.
func NewMobileappsItemGraphmicrosoftstoreforbusinessappGraphMicrosoftStoreForBusinessAppRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileappsItemGraphmicrosoftstoreforbusinessappGraphMicrosoftStoreForBusinessAppRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMobileappsItemGraphmicrosoftstoreforbusinessappGraphMicrosoftStoreForBusinessAppRequestBuilderInternal(urlParams, requestAdapter)
}
// ContainedApps provides operations to manage the containedApps property of the microsoft.graph.microsoftStoreForBusinessApp entity.
// returns a *MobileappsItemGraphmicrosoftstoreforbusinessappContainedappsContainedAppsRequestBuilder when successful
func (m *MobileappsItemGraphmicrosoftstoreforbusinessappGraphMicrosoftStoreForBusinessAppRequestBuilder) ContainedApps()(*MobileappsItemGraphmicrosoftstoreforbusinessappContainedappsContainedAppsRequestBuilder) {
    return NewMobileappsItemGraphmicrosoftstoreforbusinessappContainedappsContainedAppsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get the item of type microsoft.graph.mobileApp as microsoft.graph.microsoftStoreForBusinessApp
// returns a MicrosoftStoreForBusinessAppable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MobileappsItemGraphmicrosoftstoreforbusinessappGraphMicrosoftStoreForBusinessAppRequestBuilder) Get(ctx context.Context, requestConfiguration *MobileappsItemGraphmicrosoftstoreforbusinessappGraphMicrosoftStoreForBusinessAppRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MicrosoftStoreForBusinessAppable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMicrosoftStoreForBusinessAppFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MicrosoftStoreForBusinessAppable), nil
}
// Relationships provides operations to manage the relationships property of the microsoft.graph.mobileApp entity.
// returns a *MobileappsItemGraphmicrosoftstoreforbusinessappRelationshipsRequestBuilder when successful
func (m *MobileappsItemGraphmicrosoftstoreforbusinessappGraphMicrosoftStoreForBusinessAppRequestBuilder) Relationships()(*MobileappsItemGraphmicrosoftstoreforbusinessappRelationshipsRequestBuilder) {
    return NewMobileappsItemGraphmicrosoftstoreforbusinessappRelationshipsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation get the item of type microsoft.graph.mobileApp as microsoft.graph.microsoftStoreForBusinessApp
// returns a *RequestInformation when successful
func (m *MobileappsItemGraphmicrosoftstoreforbusinessappGraphMicrosoftStoreForBusinessAppRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MobileappsItemGraphmicrosoftstoreforbusinessappGraphMicrosoftStoreForBusinessAppRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *MobileappsItemGraphmicrosoftstoreforbusinessappGraphMicrosoftStoreForBusinessAppRequestBuilder when successful
func (m *MobileappsItemGraphmicrosoftstoreforbusinessappGraphMicrosoftStoreForBusinessAppRequestBuilder) WithUrl(rawUrl string)(*MobileappsItemGraphmicrosoftstoreforbusinessappGraphMicrosoftStoreForBusinessAppRequestBuilder) {
    return NewMobileappsItemGraphmicrosoftstoreforbusinessappGraphMicrosoftStoreForBusinessAppRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

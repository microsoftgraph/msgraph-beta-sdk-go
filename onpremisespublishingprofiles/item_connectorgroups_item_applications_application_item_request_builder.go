package onpremisespublishingprofiles

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemConnectorgroupsItemApplicationsApplicationItemRequestBuilder provides operations to manage the applications property of the microsoft.graph.connectorGroup entity.
type ItemConnectorgroupsItemApplicationsApplicationItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemConnectorgroupsItemApplicationsApplicationItemRequestBuilderGetQueryParameters get applications from onPremisesPublishingProfiles
type ItemConnectorgroupsItemApplicationsApplicationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemConnectorgroupsItemApplicationsApplicationItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemConnectorgroupsItemApplicationsApplicationItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemConnectorgroupsItemApplicationsApplicationItemRequestBuilderGetQueryParameters
}
// NewItemConnectorgroupsItemApplicationsApplicationItemRequestBuilderInternal instantiates a new ItemConnectorgroupsItemApplicationsApplicationItemRequestBuilder and sets the default values.
func NewItemConnectorgroupsItemApplicationsApplicationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemConnectorgroupsItemApplicationsApplicationItemRequestBuilder) {
    m := &ItemConnectorgroupsItemApplicationsApplicationItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/onPremisesPublishingProfiles/{onPremisesPublishingProfile%2Did}/connectorGroups/{connectorGroup%2Did}/applications/{application%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemConnectorgroupsItemApplicationsApplicationItemRequestBuilder instantiates a new ItemConnectorgroupsItemApplicationsApplicationItemRequestBuilder and sets the default values.
func NewItemConnectorgroupsItemApplicationsApplicationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemConnectorgroupsItemApplicationsApplicationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemConnectorgroupsItemApplicationsApplicationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get applications from onPremisesPublishingProfiles
// returns a Applicationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemConnectorgroupsItemApplicationsApplicationItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemConnectorgroupsItemApplicationsApplicationItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Applicationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateApplicationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Applicationable), nil
}
// Logo provides operations to manage the media for the onPremisesPublishingProfile entity.
// returns a *ItemConnectorgroupsItemApplicationsItemLogoRequestBuilder when successful
func (m *ItemConnectorgroupsItemApplicationsApplicationItemRequestBuilder) Logo()(*ItemConnectorgroupsItemApplicationsItemLogoRequestBuilder) {
    return NewItemConnectorgroupsItemApplicationsItemLogoRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation get applications from onPremisesPublishingProfiles
// returns a *RequestInformation when successful
func (m *ItemConnectorgroupsItemApplicationsApplicationItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemConnectorgroupsItemApplicationsApplicationItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemConnectorgroupsItemApplicationsApplicationItemRequestBuilder when successful
func (m *ItemConnectorgroupsItemApplicationsApplicationItemRequestBuilder) WithUrl(rawUrl string)(*ItemConnectorgroupsItemApplicationsApplicationItemRequestBuilder) {
    return NewItemConnectorgroupsItemApplicationsApplicationItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

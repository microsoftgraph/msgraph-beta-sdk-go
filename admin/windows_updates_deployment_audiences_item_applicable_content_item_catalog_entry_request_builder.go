package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c "github.com/microsoftgraph/msgraph-beta-sdk-go/models/windowsupdates"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// WindowsUpdatesDeploymentAudiencesItemApplicableContentItemCatalogEntryRequestBuilder provides operations to manage the catalogEntry property of the microsoft.graph.windowsUpdates.applicableContent entity.
type WindowsUpdatesDeploymentAudiencesItemApplicableContentItemCatalogEntryRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WindowsUpdatesDeploymentAudiencesItemApplicableContentItemCatalogEntryRequestBuilderGetQueryParameters catalog entry for the update or content.
type WindowsUpdatesDeploymentAudiencesItemApplicableContentItemCatalogEntryRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// WindowsUpdatesDeploymentAudiencesItemApplicableContentItemCatalogEntryRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsUpdatesDeploymentAudiencesItemApplicableContentItemCatalogEntryRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *WindowsUpdatesDeploymentAudiencesItemApplicableContentItemCatalogEntryRequestBuilderGetQueryParameters
}
// NewWindowsUpdatesDeploymentAudiencesItemApplicableContentItemCatalogEntryRequestBuilderInternal instantiates a new WindowsUpdatesDeploymentAudiencesItemApplicableContentItemCatalogEntryRequestBuilder and sets the default values.
func NewWindowsUpdatesDeploymentAudiencesItemApplicableContentItemCatalogEntryRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsUpdatesDeploymentAudiencesItemApplicableContentItemCatalogEntryRequestBuilder) {
    m := &WindowsUpdatesDeploymentAudiencesItemApplicableContentItemCatalogEntryRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/admin/windows/updates/deploymentAudiences/{deploymentAudience%2Did}/applicableContent/{applicableContent%2DcatalogEntryId}/catalogEntry{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewWindowsUpdatesDeploymentAudiencesItemApplicableContentItemCatalogEntryRequestBuilder instantiates a new WindowsUpdatesDeploymentAudiencesItemApplicableContentItemCatalogEntryRequestBuilder and sets the default values.
func NewWindowsUpdatesDeploymentAudiencesItemApplicableContentItemCatalogEntryRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsUpdatesDeploymentAudiencesItemApplicableContentItemCatalogEntryRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsUpdatesDeploymentAudiencesItemApplicableContentItemCatalogEntryRequestBuilderInternal(urlParams, requestAdapter)
}
// Get catalog entry for the update or content.
// returns a CatalogEntryable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsUpdatesDeploymentAudiencesItemApplicableContentItemCatalogEntryRequestBuilder) Get(ctx context.Context, requestConfiguration *WindowsUpdatesDeploymentAudiencesItemApplicableContentItemCatalogEntryRequestBuilderGetRequestConfiguration)(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.CatalogEntryable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.CreateCatalogEntryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.CatalogEntryable), nil
}
// ToGetRequestInformation catalog entry for the update or content.
// returns a *RequestInformation when successful
func (m *WindowsUpdatesDeploymentAudiencesItemApplicableContentItemCatalogEntryRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *WindowsUpdatesDeploymentAudiencesItemApplicableContentItemCatalogEntryRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *WindowsUpdatesDeploymentAudiencesItemApplicableContentItemCatalogEntryRequestBuilder when successful
func (m *WindowsUpdatesDeploymentAudiencesItemApplicableContentItemCatalogEntryRequestBuilder) WithUrl(rawUrl string)(*WindowsUpdatesDeploymentAudiencesItemApplicableContentItemCatalogEntryRequestBuilder) {
    return NewWindowsUpdatesDeploymentAudiencesItemApplicableContentItemCatalogEntryRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MacOSSoftwareUpdateAccountSummariesItemCategorySummariesItemUpdateStateSummariesMacOSSoftwareUpdateStateSummaryItemRequestBuilder provides operations to manage the updateStateSummaries property of the microsoft.graph.macOSSoftwareUpdateCategorySummary entity.
type MacOSSoftwareUpdateAccountSummariesItemCategorySummariesItemUpdateStateSummariesMacOSSoftwareUpdateStateSummaryItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MacOSSoftwareUpdateAccountSummariesItemCategorySummariesItemUpdateStateSummariesMacOSSoftwareUpdateStateSummaryItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MacOSSoftwareUpdateAccountSummariesItemCategorySummariesItemUpdateStateSummariesMacOSSoftwareUpdateStateSummaryItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// MacOSSoftwareUpdateAccountSummariesItemCategorySummariesItemUpdateStateSummariesMacOSSoftwareUpdateStateSummaryItemRequestBuilderGetQueryParameters summary of the update states.
type MacOSSoftwareUpdateAccountSummariesItemCategorySummariesItemUpdateStateSummariesMacOSSoftwareUpdateStateSummaryItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MacOSSoftwareUpdateAccountSummariesItemCategorySummariesItemUpdateStateSummariesMacOSSoftwareUpdateStateSummaryItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MacOSSoftwareUpdateAccountSummariesItemCategorySummariesItemUpdateStateSummariesMacOSSoftwareUpdateStateSummaryItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MacOSSoftwareUpdateAccountSummariesItemCategorySummariesItemUpdateStateSummariesMacOSSoftwareUpdateStateSummaryItemRequestBuilderGetQueryParameters
}
// MacOSSoftwareUpdateAccountSummariesItemCategorySummariesItemUpdateStateSummariesMacOSSoftwareUpdateStateSummaryItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MacOSSoftwareUpdateAccountSummariesItemCategorySummariesItemUpdateStateSummariesMacOSSoftwareUpdateStateSummaryItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMacOSSoftwareUpdateAccountSummariesItemCategorySummariesItemUpdateStateSummariesMacOSSoftwareUpdateStateSummaryItemRequestBuilderInternal instantiates a new MacOSSoftwareUpdateAccountSummariesItemCategorySummariesItemUpdateStateSummariesMacOSSoftwareUpdateStateSummaryItemRequestBuilder and sets the default values.
func NewMacOSSoftwareUpdateAccountSummariesItemCategorySummariesItemUpdateStateSummariesMacOSSoftwareUpdateStateSummaryItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MacOSSoftwareUpdateAccountSummariesItemCategorySummariesItemUpdateStateSummariesMacOSSoftwareUpdateStateSummaryItemRequestBuilder) {
    m := &MacOSSoftwareUpdateAccountSummariesItemCategorySummariesItemUpdateStateSummariesMacOSSoftwareUpdateStateSummaryItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/macOSSoftwareUpdateAccountSummaries/{macOSSoftwareUpdateAccountSummary%2Did}/categorySummaries/{macOSSoftwareUpdateCategorySummary%2Did}/updateStateSummaries/{macOSSoftwareUpdateStateSummary%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewMacOSSoftwareUpdateAccountSummariesItemCategorySummariesItemUpdateStateSummariesMacOSSoftwareUpdateStateSummaryItemRequestBuilder instantiates a new MacOSSoftwareUpdateAccountSummariesItemCategorySummariesItemUpdateStateSummariesMacOSSoftwareUpdateStateSummaryItemRequestBuilder and sets the default values.
func NewMacOSSoftwareUpdateAccountSummariesItemCategorySummariesItemUpdateStateSummariesMacOSSoftwareUpdateStateSummaryItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MacOSSoftwareUpdateAccountSummariesItemCategorySummariesItemUpdateStateSummariesMacOSSoftwareUpdateStateSummaryItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMacOSSoftwareUpdateAccountSummariesItemCategorySummariesItemUpdateStateSummariesMacOSSoftwareUpdateStateSummaryItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property updateStateSummaries for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MacOSSoftwareUpdateAccountSummariesItemCategorySummariesItemUpdateStateSummariesMacOSSoftwareUpdateStateSummaryItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *MacOSSoftwareUpdateAccountSummariesItemCategorySummariesItemUpdateStateSummariesMacOSSoftwareUpdateStateSummaryItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get summary of the update states.
// returns a MacOSSoftwareUpdateStateSummaryable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MacOSSoftwareUpdateAccountSummariesItemCategorySummariesItemUpdateStateSummariesMacOSSoftwareUpdateStateSummaryItemRequestBuilder) Get(ctx context.Context, requestConfiguration *MacOSSoftwareUpdateAccountSummariesItemCategorySummariesItemUpdateStateSummariesMacOSSoftwareUpdateStateSummaryItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MacOSSoftwareUpdateStateSummaryable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMacOSSoftwareUpdateStateSummaryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MacOSSoftwareUpdateStateSummaryable), nil
}
// Patch update the navigation property updateStateSummaries in deviceManagement
// returns a MacOSSoftwareUpdateStateSummaryable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MacOSSoftwareUpdateAccountSummariesItemCategorySummariesItemUpdateStateSummariesMacOSSoftwareUpdateStateSummaryItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MacOSSoftwareUpdateStateSummaryable, requestConfiguration *MacOSSoftwareUpdateAccountSummariesItemCategorySummariesItemUpdateStateSummariesMacOSSoftwareUpdateStateSummaryItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MacOSSoftwareUpdateStateSummaryable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMacOSSoftwareUpdateStateSummaryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MacOSSoftwareUpdateStateSummaryable), nil
}
// ToDeleteRequestInformation delete navigation property updateStateSummaries for deviceManagement
// returns a *RequestInformation when successful
func (m *MacOSSoftwareUpdateAccountSummariesItemCategorySummariesItemUpdateStateSummariesMacOSSoftwareUpdateStateSummaryItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *MacOSSoftwareUpdateAccountSummariesItemCategorySummariesItemUpdateStateSummariesMacOSSoftwareUpdateStateSummaryItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation summary of the update states.
// returns a *RequestInformation when successful
func (m *MacOSSoftwareUpdateAccountSummariesItemCategorySummariesItemUpdateStateSummariesMacOSSoftwareUpdateStateSummaryItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MacOSSoftwareUpdateAccountSummariesItemCategorySummariesItemUpdateStateSummariesMacOSSoftwareUpdateStateSummaryItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property updateStateSummaries in deviceManagement
// returns a *RequestInformation when successful
func (m *MacOSSoftwareUpdateAccountSummariesItemCategorySummariesItemUpdateStateSummariesMacOSSoftwareUpdateStateSummaryItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MacOSSoftwareUpdateStateSummaryable, requestConfiguration *MacOSSoftwareUpdateAccountSummariesItemCategorySummariesItemUpdateStateSummariesMacOSSoftwareUpdateStateSummaryItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *MacOSSoftwareUpdateAccountSummariesItemCategorySummariesItemUpdateStateSummariesMacOSSoftwareUpdateStateSummaryItemRequestBuilder when successful
func (m *MacOSSoftwareUpdateAccountSummariesItemCategorySummariesItemUpdateStateSummariesMacOSSoftwareUpdateStateSummaryItemRequestBuilder) WithUrl(rawUrl string)(*MacOSSoftwareUpdateAccountSummariesItemCategorySummariesItemUpdateStateSummariesMacOSSoftwareUpdateStateSummaryItemRequestBuilder) {
    return NewMacOSSoftwareUpdateAccountSummariesItemCategorySummariesItemUpdateStateSummariesMacOSSoftwareUpdateStateSummaryItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

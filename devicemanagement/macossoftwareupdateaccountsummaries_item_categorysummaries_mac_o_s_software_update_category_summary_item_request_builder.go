package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MacossoftwareupdateaccountsummariesItemCategorysummariesMacOSSoftwareUpdateCategorySummaryItemRequestBuilder provides operations to manage the categorySummaries property of the microsoft.graph.macOSSoftwareUpdateAccountSummary entity.
type MacossoftwareupdateaccountsummariesItemCategorysummariesMacOSSoftwareUpdateCategorySummaryItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MacossoftwareupdateaccountsummariesItemCategorysummariesMacOSSoftwareUpdateCategorySummaryItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MacossoftwareupdateaccountsummariesItemCategorysummariesMacOSSoftwareUpdateCategorySummaryItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// MacossoftwareupdateaccountsummariesItemCategorysummariesMacOSSoftwareUpdateCategorySummaryItemRequestBuilderGetQueryParameters summary of the updates by category.
type MacossoftwareupdateaccountsummariesItemCategorysummariesMacOSSoftwareUpdateCategorySummaryItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MacossoftwareupdateaccountsummariesItemCategorysummariesMacOSSoftwareUpdateCategorySummaryItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MacossoftwareupdateaccountsummariesItemCategorysummariesMacOSSoftwareUpdateCategorySummaryItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MacossoftwareupdateaccountsummariesItemCategorysummariesMacOSSoftwareUpdateCategorySummaryItemRequestBuilderGetQueryParameters
}
// MacossoftwareupdateaccountsummariesItemCategorysummariesMacOSSoftwareUpdateCategorySummaryItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MacossoftwareupdateaccountsummariesItemCategorysummariesMacOSSoftwareUpdateCategorySummaryItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMacossoftwareupdateaccountsummariesItemCategorysummariesMacOSSoftwareUpdateCategorySummaryItemRequestBuilderInternal instantiates a new MacossoftwareupdateaccountsummariesItemCategorysummariesMacOSSoftwareUpdateCategorySummaryItemRequestBuilder and sets the default values.
func NewMacossoftwareupdateaccountsummariesItemCategorysummariesMacOSSoftwareUpdateCategorySummaryItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MacossoftwareupdateaccountsummariesItemCategorysummariesMacOSSoftwareUpdateCategorySummaryItemRequestBuilder) {
    m := &MacossoftwareupdateaccountsummariesItemCategorysummariesMacOSSoftwareUpdateCategorySummaryItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/macOSSoftwareUpdateAccountSummaries/{macOSSoftwareUpdateAccountSummary%2Did}/categorySummaries/{macOSSoftwareUpdateCategorySummary%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewMacossoftwareupdateaccountsummariesItemCategorysummariesMacOSSoftwareUpdateCategorySummaryItemRequestBuilder instantiates a new MacossoftwareupdateaccountsummariesItemCategorysummariesMacOSSoftwareUpdateCategorySummaryItemRequestBuilder and sets the default values.
func NewMacossoftwareupdateaccountsummariesItemCategorysummariesMacOSSoftwareUpdateCategorySummaryItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MacossoftwareupdateaccountsummariesItemCategorysummariesMacOSSoftwareUpdateCategorySummaryItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMacossoftwareupdateaccountsummariesItemCategorysummariesMacOSSoftwareUpdateCategorySummaryItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property categorySummaries for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MacossoftwareupdateaccountsummariesItemCategorysummariesMacOSSoftwareUpdateCategorySummaryItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *MacossoftwareupdateaccountsummariesItemCategorysummariesMacOSSoftwareUpdateCategorySummaryItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get summary of the updates by category.
// returns a MacOSSoftwareUpdateCategorySummaryable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MacossoftwareupdateaccountsummariesItemCategorysummariesMacOSSoftwareUpdateCategorySummaryItemRequestBuilder) Get(ctx context.Context, requestConfiguration *MacossoftwareupdateaccountsummariesItemCategorysummariesMacOSSoftwareUpdateCategorySummaryItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MacOSSoftwareUpdateCategorySummaryable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMacOSSoftwareUpdateCategorySummaryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MacOSSoftwareUpdateCategorySummaryable), nil
}
// Patch update the navigation property categorySummaries in deviceManagement
// returns a MacOSSoftwareUpdateCategorySummaryable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MacossoftwareupdateaccountsummariesItemCategorysummariesMacOSSoftwareUpdateCategorySummaryItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MacOSSoftwareUpdateCategorySummaryable, requestConfiguration *MacossoftwareupdateaccountsummariesItemCategorysummariesMacOSSoftwareUpdateCategorySummaryItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MacOSSoftwareUpdateCategorySummaryable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMacOSSoftwareUpdateCategorySummaryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MacOSSoftwareUpdateCategorySummaryable), nil
}
// ToDeleteRequestInformation delete navigation property categorySummaries for deviceManagement
// returns a *RequestInformation when successful
func (m *MacossoftwareupdateaccountsummariesItemCategorysummariesMacOSSoftwareUpdateCategorySummaryItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *MacossoftwareupdateaccountsummariesItemCategorysummariesMacOSSoftwareUpdateCategorySummaryItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation summary of the updates by category.
// returns a *RequestInformation when successful
func (m *MacossoftwareupdateaccountsummariesItemCategorysummariesMacOSSoftwareUpdateCategorySummaryItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MacossoftwareupdateaccountsummariesItemCategorysummariesMacOSSoftwareUpdateCategorySummaryItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property categorySummaries in deviceManagement
// returns a *RequestInformation when successful
func (m *MacossoftwareupdateaccountsummariesItemCategorysummariesMacOSSoftwareUpdateCategorySummaryItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MacOSSoftwareUpdateCategorySummaryable, requestConfiguration *MacossoftwareupdateaccountsummariesItemCategorysummariesMacOSSoftwareUpdateCategorySummaryItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// UpdateStateSummaries provides operations to manage the updateStateSummaries property of the microsoft.graph.macOSSoftwareUpdateCategorySummary entity.
// returns a *MacossoftwareupdateaccountsummariesItemCategorysummariesItemUpdatestatesummariesUpdateStateSummariesRequestBuilder when successful
func (m *MacossoftwareupdateaccountsummariesItemCategorysummariesMacOSSoftwareUpdateCategorySummaryItemRequestBuilder) UpdateStateSummaries()(*MacossoftwareupdateaccountsummariesItemCategorysummariesItemUpdatestatesummariesUpdateStateSummariesRequestBuilder) {
    return NewMacossoftwareupdateaccountsummariesItemCategorysummariesItemUpdatestatesummariesUpdateStateSummariesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *MacossoftwareupdateaccountsummariesItemCategorysummariesMacOSSoftwareUpdateCategorySummaryItemRequestBuilder when successful
func (m *MacossoftwareupdateaccountsummariesItemCategorysummariesMacOSSoftwareUpdateCategorySummaryItemRequestBuilder) WithUrl(rawUrl string)(*MacossoftwareupdateaccountsummariesItemCategorysummariesMacOSSoftwareUpdateCategorySummaryItemRequestBuilder) {
    return NewMacossoftwareupdateaccountsummariesItemCategorysummariesMacOSSoftwareUpdateCategorySummaryItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

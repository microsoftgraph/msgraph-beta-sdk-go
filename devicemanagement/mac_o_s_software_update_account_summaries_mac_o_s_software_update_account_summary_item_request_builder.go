package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MacOSSoftwareUpdateAccountSummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilder provides operations to manage the macOSSoftwareUpdateAccountSummaries property of the microsoft.graph.deviceManagement entity.
type MacOSSoftwareUpdateAccountSummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MacOSSoftwareUpdateAccountSummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MacOSSoftwareUpdateAccountSummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// MacOSSoftwareUpdateAccountSummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilderGetQueryParameters the MacOS software update account summaries for this account.
type MacOSSoftwareUpdateAccountSummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MacOSSoftwareUpdateAccountSummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MacOSSoftwareUpdateAccountSummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MacOSSoftwareUpdateAccountSummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilderGetQueryParameters
}
// MacOSSoftwareUpdateAccountSummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MacOSSoftwareUpdateAccountSummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CategorySummaries provides operations to manage the categorySummaries property of the microsoft.graph.macOSSoftwareUpdateAccountSummary entity.
func (m *MacOSSoftwareUpdateAccountSummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilder) CategorySummaries()(*MacOSSoftwareUpdateAccountSummariesItemCategorySummariesRequestBuilder) {
    return NewMacOSSoftwareUpdateAccountSummariesItemCategorySummariesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CategorySummariesById provides operations to manage the categorySummaries property of the microsoft.graph.macOSSoftwareUpdateAccountSummary entity.
func (m *MacOSSoftwareUpdateAccountSummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilder) CategorySummariesById(id string)(*MacOSSoftwareUpdateAccountSummariesItemCategorySummariesMacOSSoftwareUpdateCategorySummaryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["macOSSoftwareUpdateCategorySummary%2Did"] = id
    }
    return NewMacOSSoftwareUpdateAccountSummariesItemCategorySummariesMacOSSoftwareUpdateCategorySummaryItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewMacOSSoftwareUpdateAccountSummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilderInternal instantiates a new MacOSSoftwareUpdateAccountSummaryItemRequestBuilder and sets the default values.
func NewMacOSSoftwareUpdateAccountSummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MacOSSoftwareUpdateAccountSummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilder) {
    m := &MacOSSoftwareUpdateAccountSummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/macOSSoftwareUpdateAccountSummaries/{macOSSoftwareUpdateAccountSummary%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewMacOSSoftwareUpdateAccountSummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilder instantiates a new MacOSSoftwareUpdateAccountSummaryItemRequestBuilder and sets the default values.
func NewMacOSSoftwareUpdateAccountSummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MacOSSoftwareUpdateAccountSummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMacOSSoftwareUpdateAccountSummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property macOSSoftwareUpdateAccountSummaries for deviceManagement
func (m *MacOSSoftwareUpdateAccountSummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *MacOSSoftwareUpdateAccountSummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get the MacOS software update account summaries for this account.
func (m *MacOSSoftwareUpdateAccountSummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilder) Get(ctx context.Context, requestConfiguration *MacOSSoftwareUpdateAccountSummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MacOSSoftwareUpdateAccountSummaryable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMacOSSoftwareUpdateAccountSummaryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MacOSSoftwareUpdateAccountSummaryable), nil
}
// Patch update the navigation property macOSSoftwareUpdateAccountSummaries in deviceManagement
func (m *MacOSSoftwareUpdateAccountSummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MacOSSoftwareUpdateAccountSummaryable, requestConfiguration *MacOSSoftwareUpdateAccountSummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MacOSSoftwareUpdateAccountSummaryable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMacOSSoftwareUpdateAccountSummaryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MacOSSoftwareUpdateAccountSummaryable), nil
}
// ToDeleteRequestInformation delete navigation property macOSSoftwareUpdateAccountSummaries for deviceManagement
func (m *MacOSSoftwareUpdateAccountSummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *MacOSSoftwareUpdateAccountSummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation the MacOS software update account summaries for this account.
func (m *MacOSSoftwareUpdateAccountSummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MacOSSoftwareUpdateAccountSummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
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
// ToPatchRequestInformation update the navigation property macOSSoftwareUpdateAccountSummaries in deviceManagement
func (m *MacOSSoftwareUpdateAccountSummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MacOSSoftwareUpdateAccountSummaryable, requestConfiguration *MacOSSoftwareUpdateAccountSummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}

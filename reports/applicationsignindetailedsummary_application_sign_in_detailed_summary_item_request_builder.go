package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ApplicationsignindetailedsummaryApplicationSignInDetailedSummaryItemRequestBuilder provides operations to manage the applicationSignInDetailedSummary property of the microsoft.graph.reportRoot entity.
type ApplicationsignindetailedsummaryApplicationSignInDetailedSummaryItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ApplicationsignindetailedsummaryApplicationSignInDetailedSummaryItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ApplicationsignindetailedsummaryApplicationSignInDetailedSummaryItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ApplicationsignindetailedsummaryApplicationSignInDetailedSummaryItemRequestBuilderGetQueryParameters retrieve the properties and relationships of an applicationSignInDetailedSummary object.
type ApplicationsignindetailedsummaryApplicationSignInDetailedSummaryItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ApplicationsignindetailedsummaryApplicationSignInDetailedSummaryItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ApplicationsignindetailedsummaryApplicationSignInDetailedSummaryItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ApplicationsignindetailedsummaryApplicationSignInDetailedSummaryItemRequestBuilderGetQueryParameters
}
// ApplicationsignindetailedsummaryApplicationSignInDetailedSummaryItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ApplicationsignindetailedsummaryApplicationSignInDetailedSummaryItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewApplicationsignindetailedsummaryApplicationSignInDetailedSummaryItemRequestBuilderInternal instantiates a new ApplicationsignindetailedsummaryApplicationSignInDetailedSummaryItemRequestBuilder and sets the default values.
func NewApplicationsignindetailedsummaryApplicationSignInDetailedSummaryItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ApplicationsignindetailedsummaryApplicationSignInDetailedSummaryItemRequestBuilder) {
    m := &ApplicationsignindetailedsummaryApplicationSignInDetailedSummaryItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/applicationSignInDetailedSummary/{applicationSignInDetailedSummary%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewApplicationsignindetailedsummaryApplicationSignInDetailedSummaryItemRequestBuilder instantiates a new ApplicationsignindetailedsummaryApplicationSignInDetailedSummaryItemRequestBuilder and sets the default values.
func NewApplicationsignindetailedsummaryApplicationSignInDetailedSummaryItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ApplicationsignindetailedsummaryApplicationSignInDetailedSummaryItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewApplicationsignindetailedsummaryApplicationSignInDetailedSummaryItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property applicationSignInDetailedSummary for reports
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ApplicationsignindetailedsummaryApplicationSignInDetailedSummaryItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ApplicationsignindetailedsummaryApplicationSignInDetailedSummaryItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get retrieve the properties and relationships of an applicationSignInDetailedSummary object.
// returns a ApplicationSignInDetailedSummaryable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/applicationsignindetailedsummary-get?view=graph-rest-beta
func (m *ApplicationsignindetailedsummaryApplicationSignInDetailedSummaryItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ApplicationsignindetailedsummaryApplicationSignInDetailedSummaryItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ApplicationSignInDetailedSummaryable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateApplicationSignInDetailedSummaryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ApplicationSignInDetailedSummaryable), nil
}
// Patch update the navigation property applicationSignInDetailedSummary in reports
// returns a ApplicationSignInDetailedSummaryable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ApplicationsignindetailedsummaryApplicationSignInDetailedSummaryItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ApplicationSignInDetailedSummaryable, requestConfiguration *ApplicationsignindetailedsummaryApplicationSignInDetailedSummaryItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ApplicationSignInDetailedSummaryable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateApplicationSignInDetailedSummaryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ApplicationSignInDetailedSummaryable), nil
}
// ToDeleteRequestInformation delete navigation property applicationSignInDetailedSummary for reports
// returns a *RequestInformation when successful
func (m *ApplicationsignindetailedsummaryApplicationSignInDetailedSummaryItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ApplicationsignindetailedsummaryApplicationSignInDetailedSummaryItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation retrieve the properties and relationships of an applicationSignInDetailedSummary object.
// returns a *RequestInformation when successful
func (m *ApplicationsignindetailedsummaryApplicationSignInDetailedSummaryItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ApplicationsignindetailedsummaryApplicationSignInDetailedSummaryItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property applicationSignInDetailedSummary in reports
// returns a *RequestInformation when successful
func (m *ApplicationsignindetailedsummaryApplicationSignInDetailedSummaryItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ApplicationSignInDetailedSummaryable, requestConfiguration *ApplicationsignindetailedsummaryApplicationSignInDetailedSummaryItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ApplicationsignindetailedsummaryApplicationSignInDetailedSummaryItemRequestBuilder when successful
func (m *ApplicationsignindetailedsummaryApplicationSignInDetailedSummaryItemRequestBuilder) WithUrl(rawUrl string)(*ApplicationsignindetailedsummaryApplicationSignInDetailedSummaryItemRequestBuilder) {
    return NewApplicationsignindetailedsummaryApplicationSignInDetailedSummaryItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

package admin

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// WindowsUpdatesProductsItemMicrosoftgraphwindowsupdatesgetknownissuesbytimerangewithdaysinpastwithincludeallactiveMicrosoftGraphWindowsUpdatesGetKnownIssuesByTimeRangeWithDaysInPastWithIncludeAllActiveRequestBuilder provides operations to call the getKnownIssuesByTimeRange method.
type WindowsUpdatesProductsItemMicrosoftgraphwindowsupdatesgetknownissuesbytimerangewithdaysinpastwithincludeallactiveMicrosoftGraphWindowsUpdatesGetKnownIssuesByTimeRangeWithDaysInPastWithIncludeAllActiveRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WindowsUpdatesProductsItemMicrosoftgraphwindowsupdatesgetknownissuesbytimerangewithdaysinpastwithincludeallactiveMicrosoftGraphWindowsUpdatesGetKnownIssuesByTimeRangeWithDaysInPastWithIncludeAllActiveRequestBuilderGetQueryParameters get known issues related to a particular product based on a specified timeframe in the past.
type WindowsUpdatesProductsItemMicrosoftgraphwindowsupdatesgetknownissuesbytimerangewithdaysinpastwithincludeallactiveMicrosoftGraphWindowsUpdatesGetKnownIssuesByTimeRangeWithDaysInPastWithIncludeAllActiveRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Usage: includeAllActive=@includeAllActive
    IncludeAllActive *bool `uriparametername:"includeAllActive"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// WindowsUpdatesProductsItemMicrosoftgraphwindowsupdatesgetknownissuesbytimerangewithdaysinpastwithincludeallactiveMicrosoftGraphWindowsUpdatesGetKnownIssuesByTimeRangeWithDaysInPastWithIncludeAllActiveRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsUpdatesProductsItemMicrosoftgraphwindowsupdatesgetknownissuesbytimerangewithdaysinpastwithincludeallactiveMicrosoftGraphWindowsUpdatesGetKnownIssuesByTimeRangeWithDaysInPastWithIncludeAllActiveRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *WindowsUpdatesProductsItemMicrosoftgraphwindowsupdatesgetknownissuesbytimerangewithdaysinpastwithincludeallactiveMicrosoftGraphWindowsUpdatesGetKnownIssuesByTimeRangeWithDaysInPastWithIncludeAllActiveRequestBuilderGetQueryParameters
}
// NewWindowsUpdatesProductsItemMicrosoftgraphwindowsupdatesgetknownissuesbytimerangewithdaysinpastwithincludeallactiveMicrosoftGraphWindowsUpdatesGetKnownIssuesByTimeRangeWithDaysInPastWithIncludeAllActiveRequestBuilderInternal instantiates a new WindowsUpdatesProductsItemMicrosoftgraphwindowsupdatesgetknownissuesbytimerangewithdaysinpastwithincludeallactiveMicrosoftGraphWindowsUpdatesGetKnownIssuesByTimeRangeWithDaysInPastWithIncludeAllActiveRequestBuilder and sets the default values.
func NewWindowsUpdatesProductsItemMicrosoftgraphwindowsupdatesgetknownissuesbytimerangewithdaysinpastwithincludeallactiveMicrosoftGraphWindowsUpdatesGetKnownIssuesByTimeRangeWithDaysInPastWithIncludeAllActiveRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, daysInPast *int32)(*WindowsUpdatesProductsItemMicrosoftgraphwindowsupdatesgetknownissuesbytimerangewithdaysinpastwithincludeallactiveMicrosoftGraphWindowsUpdatesGetKnownIssuesByTimeRangeWithDaysInPastWithIncludeAllActiveRequestBuilder) {
    m := &WindowsUpdatesProductsItemMicrosoftgraphwindowsupdatesgetknownissuesbytimerangewithdaysinpastwithincludeallactiveMicrosoftGraphWindowsUpdatesGetKnownIssuesByTimeRangeWithDaysInPastWithIncludeAllActiveRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/admin/windows/updates/products/{product%2Did}/microsoft.graph.windowsUpdates.getKnownIssuesByTimeRange(daysInPast={daysInPast},includeAllActive=@includeAllActive){?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top,includeAllActive*}", pathParameters),
    }
    if daysInPast != nil {
        m.BaseRequestBuilder.PathParameters["daysInPast"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(*daysInPast), 10)
    }
    return m
}
// NewWindowsUpdatesProductsItemMicrosoftgraphwindowsupdatesgetknownissuesbytimerangewithdaysinpastwithincludeallactiveMicrosoftGraphWindowsUpdatesGetKnownIssuesByTimeRangeWithDaysInPastWithIncludeAllActiveRequestBuilder instantiates a new WindowsUpdatesProductsItemMicrosoftgraphwindowsupdatesgetknownissuesbytimerangewithdaysinpastwithincludeallactiveMicrosoftGraphWindowsUpdatesGetKnownIssuesByTimeRangeWithDaysInPastWithIncludeAllActiveRequestBuilder and sets the default values.
func NewWindowsUpdatesProductsItemMicrosoftgraphwindowsupdatesgetknownissuesbytimerangewithdaysinpastwithincludeallactiveMicrosoftGraphWindowsUpdatesGetKnownIssuesByTimeRangeWithDaysInPastWithIncludeAllActiveRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsUpdatesProductsItemMicrosoftgraphwindowsupdatesgetknownissuesbytimerangewithdaysinpastwithincludeallactiveMicrosoftGraphWindowsUpdatesGetKnownIssuesByTimeRangeWithDaysInPastWithIncludeAllActiveRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsUpdatesProductsItemMicrosoftgraphwindowsupdatesgetknownissuesbytimerangewithdaysinpastwithincludeallactiveMicrosoftGraphWindowsUpdatesGetKnownIssuesByTimeRangeWithDaysInPastWithIncludeAllActiveRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get get known issues related to a particular product based on a specified timeframe in the past.
// Deprecated: This method is obsolete. Use GetAsGetKnownIssuesByTimeRangeWithDaysInPastWithIncludeAllActiveGetResponse instead.
// returns a WindowsUpdatesProductsItemMicrosoftgraphwindowsupdatesgetknownissuesbytimerangewithdaysinpastwithincludeallactiveGetKnownIssuesByTimeRangeWithDaysInPastWithIncludeAllActiveResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/windowsupdates-product-getknownissuesbytimerange?view=graph-rest-beta
func (m *WindowsUpdatesProductsItemMicrosoftgraphwindowsupdatesgetknownissuesbytimerangewithdaysinpastwithincludeallactiveMicrosoftGraphWindowsUpdatesGetKnownIssuesByTimeRangeWithDaysInPastWithIncludeAllActiveRequestBuilder) Get(ctx context.Context, requestConfiguration *WindowsUpdatesProductsItemMicrosoftgraphwindowsupdatesgetknownissuesbytimerangewithdaysinpastwithincludeallactiveMicrosoftGraphWindowsUpdatesGetKnownIssuesByTimeRangeWithDaysInPastWithIncludeAllActiveRequestBuilderGetRequestConfiguration)(WindowsUpdatesProductsItemMicrosoftgraphwindowsupdatesgetknownissuesbytimerangewithdaysinpastwithincludeallactiveGetKnownIssuesByTimeRangeWithDaysInPastWithIncludeAllActiveResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateWindowsUpdatesProductsItemMicrosoftgraphwindowsupdatesgetknownissuesbytimerangewithdaysinpastwithincludeallactiveGetKnownIssuesByTimeRangeWithDaysInPastWithIncludeAllActiveResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(WindowsUpdatesProductsItemMicrosoftgraphwindowsupdatesgetknownissuesbytimerangewithdaysinpastwithincludeallactiveGetKnownIssuesByTimeRangeWithDaysInPastWithIncludeAllActiveResponseable), nil
}
// GetAsGetKnownIssuesByTimeRangeWithDaysInPastWithIncludeAllActiveGetResponse get known issues related to a particular product based on a specified timeframe in the past.
// returns a WindowsUpdatesProductsItemMicrosoftgraphwindowsupdatesgetknownissuesbytimerangewithdaysinpastwithincludeallactiveGetKnownIssuesByTimeRangeWithDaysInPastWithIncludeAllActiveGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/windowsupdates-product-getknownissuesbytimerange?view=graph-rest-beta
func (m *WindowsUpdatesProductsItemMicrosoftgraphwindowsupdatesgetknownissuesbytimerangewithdaysinpastwithincludeallactiveMicrosoftGraphWindowsUpdatesGetKnownIssuesByTimeRangeWithDaysInPastWithIncludeAllActiveRequestBuilder) GetAsGetKnownIssuesByTimeRangeWithDaysInPastWithIncludeAllActiveGetResponse(ctx context.Context, requestConfiguration *WindowsUpdatesProductsItemMicrosoftgraphwindowsupdatesgetknownissuesbytimerangewithdaysinpastwithincludeallactiveMicrosoftGraphWindowsUpdatesGetKnownIssuesByTimeRangeWithDaysInPastWithIncludeAllActiveRequestBuilderGetRequestConfiguration)(WindowsUpdatesProductsItemMicrosoftgraphwindowsupdatesgetknownissuesbytimerangewithdaysinpastwithincludeallactiveGetKnownIssuesByTimeRangeWithDaysInPastWithIncludeAllActiveGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateWindowsUpdatesProductsItemMicrosoftgraphwindowsupdatesgetknownissuesbytimerangewithdaysinpastwithincludeallactiveGetKnownIssuesByTimeRangeWithDaysInPastWithIncludeAllActiveGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(WindowsUpdatesProductsItemMicrosoftgraphwindowsupdatesgetknownissuesbytimerangewithdaysinpastwithincludeallactiveGetKnownIssuesByTimeRangeWithDaysInPastWithIncludeAllActiveGetResponseable), nil
}
// ToGetRequestInformation get known issues related to a particular product based on a specified timeframe in the past.
// returns a *RequestInformation when successful
func (m *WindowsUpdatesProductsItemMicrosoftgraphwindowsupdatesgetknownissuesbytimerangewithdaysinpastwithincludeallactiveMicrosoftGraphWindowsUpdatesGetKnownIssuesByTimeRangeWithDaysInPastWithIncludeAllActiveRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *WindowsUpdatesProductsItemMicrosoftgraphwindowsupdatesgetknownissuesbytimerangewithdaysinpastwithincludeallactiveMicrosoftGraphWindowsUpdatesGetKnownIssuesByTimeRangeWithDaysInPastWithIncludeAllActiveRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *WindowsUpdatesProductsItemMicrosoftgraphwindowsupdatesgetknownissuesbytimerangewithdaysinpastwithincludeallactiveMicrosoftGraphWindowsUpdatesGetKnownIssuesByTimeRangeWithDaysInPastWithIncludeAllActiveRequestBuilder when successful
func (m *WindowsUpdatesProductsItemMicrosoftgraphwindowsupdatesgetknownissuesbytimerangewithdaysinpastwithincludeallactiveMicrosoftGraphWindowsUpdatesGetKnownIssuesByTimeRangeWithDaysInPastWithIncludeAllActiveRequestBuilder) WithUrl(rawUrl string)(*WindowsUpdatesProductsItemMicrosoftgraphwindowsupdatesgetknownissuesbytimerangewithdaysinpastwithincludeallactiveMicrosoftGraphWindowsUpdatesGetKnownIssuesByTimeRangeWithDaysInPastWithIncludeAllActiveRequestBuilder) {
    return NewWindowsUpdatesProductsItemMicrosoftgraphwindowsupdatesgetknownissuesbytimerangewithdaysinpastwithincludeallactiveMicrosoftGraphWindowsUpdatesGetKnownIssuesByTimeRangeWithDaysInPastWithIncludeAllActiveRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

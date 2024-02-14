package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ManagedEBooksItemUserStateSummaryRequestBuilder provides operations to manage the userStateSummary property of the microsoft.graph.managedEBook entity.
type ManagedEBooksItemUserStateSummaryRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManagedEBooksItemUserStateSummaryRequestBuilderGetQueryParameters the list of installation states for this eBook.
type ManagedEBooksItemUserStateSummaryRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
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
// ManagedEBooksItemUserStateSummaryRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedEBooksItemUserStateSummaryRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ManagedEBooksItemUserStateSummaryRequestBuilderGetQueryParameters
}
// ManagedEBooksItemUserStateSummaryRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedEBooksItemUserStateSummaryRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByUserInstallStateSummaryId provides operations to manage the userStateSummary property of the microsoft.graph.managedEBook entity.
// returns a *ManagedEBooksItemUserStateSummaryUserInstallStateSummaryItemRequestBuilder when successful
func (m *ManagedEBooksItemUserStateSummaryRequestBuilder) ByUserInstallStateSummaryId(userInstallStateSummaryId string)(*ManagedEBooksItemUserStateSummaryUserInstallStateSummaryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if userInstallStateSummaryId != "" {
        urlTplParams["userInstallStateSummary%2Did"] = userInstallStateSummaryId
    }
    return NewManagedEBooksItemUserStateSummaryUserInstallStateSummaryItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewManagedEBooksItemUserStateSummaryRequestBuilderInternal instantiates a new ManagedEBooksItemUserStateSummaryRequestBuilder and sets the default values.
func NewManagedEBooksItemUserStateSummaryRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedEBooksItemUserStateSummaryRequestBuilder) {
    m := &ManagedEBooksItemUserStateSummaryRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/managedEBooks/{managedEBook%2Did}/userStateSummary{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewManagedEBooksItemUserStateSummaryRequestBuilder instantiates a new ManagedEBooksItemUserStateSummaryRequestBuilder and sets the default values.
func NewManagedEBooksItemUserStateSummaryRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedEBooksItemUserStateSummaryRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedEBooksItemUserStateSummaryRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ManagedEBooksItemUserStateSummaryCountRequestBuilder when successful
func (m *ManagedEBooksItemUserStateSummaryRequestBuilder) Count()(*ManagedEBooksItemUserStateSummaryCountRequestBuilder) {
    return NewManagedEBooksItemUserStateSummaryCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the list of installation states for this eBook.
// returns a UserInstallStateSummaryCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManagedEBooksItemUserStateSummaryRequestBuilder) Get(ctx context.Context, requestConfiguration *ManagedEBooksItemUserStateSummaryRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserInstallStateSummaryCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserInstallStateSummaryCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserInstallStateSummaryCollectionResponseable), nil
}
// Post create new navigation property to userStateSummary for deviceAppManagement
// returns a UserInstallStateSummaryable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManagedEBooksItemUserStateSummaryRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserInstallStateSummaryable, requestConfiguration *ManagedEBooksItemUserStateSummaryRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserInstallStateSummaryable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserInstallStateSummaryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserInstallStateSummaryable), nil
}
// ToGetRequestInformation the list of installation states for this eBook.
// returns a *RequestInformation when successful
func (m *ManagedEBooksItemUserStateSummaryRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ManagedEBooksItemUserStateSummaryRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to userStateSummary for deviceAppManagement
// returns a *RequestInformation when successful
func (m *ManagedEBooksItemUserStateSummaryRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserInstallStateSummaryable, requestConfiguration *ManagedEBooksItemUserStateSummaryRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, "{+baseurl}/deviceAppManagement/managedEBooks/{managedEBook%2Did}/userStateSummary", m.BaseRequestBuilder.PathParameters)
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
// returns a *ManagedEBooksItemUserStateSummaryRequestBuilder when successful
func (m *ManagedEBooksItemUserStateSummaryRequestBuilder) WithUrl(rawUrl string)(*ManagedEBooksItemUserStateSummaryRequestBuilder) {
    return NewManagedEBooksItemUserStateSummaryRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

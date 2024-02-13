package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// WindowsInformationProtectionPoliciesItemExemptAppLockerFilesRequestBuilder provides operations to manage the exemptAppLockerFiles property of the microsoft.graph.windowsInformationProtection entity.
type WindowsInformationProtectionPoliciesItemExemptAppLockerFilesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WindowsInformationProtectionPoliciesItemExemptAppLockerFilesRequestBuilderGetQueryParameters another way to input exempt apps through xml files
type WindowsInformationProtectionPoliciesItemExemptAppLockerFilesRequestBuilderGetQueryParameters struct {
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
// WindowsInformationProtectionPoliciesItemExemptAppLockerFilesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsInformationProtectionPoliciesItemExemptAppLockerFilesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *WindowsInformationProtectionPoliciesItemExemptAppLockerFilesRequestBuilderGetQueryParameters
}
// WindowsInformationProtectionPoliciesItemExemptAppLockerFilesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsInformationProtectionPoliciesItemExemptAppLockerFilesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByWindowsInformationProtectionAppLockerFileId provides operations to manage the exemptAppLockerFiles property of the microsoft.graph.windowsInformationProtection entity.
// returns a *WindowsInformationProtectionPoliciesItemExemptAppLockerFilesWindowsInformationProtectionAppLockerFileItemRequestBuilder when successful
func (m *WindowsInformationProtectionPoliciesItemExemptAppLockerFilesRequestBuilder) ByWindowsInformationProtectionAppLockerFileId(windowsInformationProtectionAppLockerFileId string)(*WindowsInformationProtectionPoliciesItemExemptAppLockerFilesWindowsInformationProtectionAppLockerFileItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if windowsInformationProtectionAppLockerFileId != "" {
        urlTplParams["windowsInformationProtectionAppLockerFile%2Did"] = windowsInformationProtectionAppLockerFileId
    }
    return NewWindowsInformationProtectionPoliciesItemExemptAppLockerFilesWindowsInformationProtectionAppLockerFileItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewWindowsInformationProtectionPoliciesItemExemptAppLockerFilesRequestBuilderInternal instantiates a new WindowsInformationProtectionPoliciesItemExemptAppLockerFilesRequestBuilder and sets the default values.
func NewWindowsInformationProtectionPoliciesItemExemptAppLockerFilesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsInformationProtectionPoliciesItemExemptAppLockerFilesRequestBuilder) {
    m := &WindowsInformationProtectionPoliciesItemExemptAppLockerFilesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/windowsInformationProtectionPolicies/{windowsInformationProtectionPolicy%2Did}/exemptAppLockerFiles{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewWindowsInformationProtectionPoliciesItemExemptAppLockerFilesRequestBuilder instantiates a new WindowsInformationProtectionPoliciesItemExemptAppLockerFilesRequestBuilder and sets the default values.
func NewWindowsInformationProtectionPoliciesItemExemptAppLockerFilesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsInformationProtectionPoliciesItemExemptAppLockerFilesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsInformationProtectionPoliciesItemExemptAppLockerFilesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *WindowsInformationProtectionPoliciesItemExemptAppLockerFilesCountRequestBuilder when successful
func (m *WindowsInformationProtectionPoliciesItemExemptAppLockerFilesRequestBuilder) Count()(*WindowsInformationProtectionPoliciesItemExemptAppLockerFilesCountRequestBuilder) {
    return NewWindowsInformationProtectionPoliciesItemExemptAppLockerFilesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get another way to input exempt apps through xml files
// returns a WindowsInformationProtectionAppLockerFileCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsInformationProtectionPoliciesItemExemptAppLockerFilesRequestBuilder) Get(ctx context.Context, requestConfiguration *WindowsInformationProtectionPoliciesItemExemptAppLockerFilesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsInformationProtectionAppLockerFileCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWindowsInformationProtectionAppLockerFileCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsInformationProtectionAppLockerFileCollectionResponseable), nil
}
// Post create new navigation property to exemptAppLockerFiles for deviceAppManagement
// returns a WindowsInformationProtectionAppLockerFileable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsInformationProtectionPoliciesItemExemptAppLockerFilesRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsInformationProtectionAppLockerFileable, requestConfiguration *WindowsInformationProtectionPoliciesItemExemptAppLockerFilesRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsInformationProtectionAppLockerFileable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWindowsInformationProtectionAppLockerFileFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsInformationProtectionAppLockerFileable), nil
}
// ToGetRequestInformation another way to input exempt apps through xml files
// returns a *RequestInformation when successful
func (m *WindowsInformationProtectionPoliciesItemExemptAppLockerFilesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *WindowsInformationProtectionPoliciesItemExemptAppLockerFilesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to exemptAppLockerFiles for deviceAppManagement
// returns a *RequestInformation when successful
func (m *WindowsInformationProtectionPoliciesItemExemptAppLockerFilesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsInformationProtectionAppLockerFileable, requestConfiguration *WindowsInformationProtectionPoliciesItemExemptAppLockerFilesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, "{+baseurl}/deviceAppManagement/windowsInformationProtectionPolicies/{windowsInformationProtectionPolicy%2Did}/exemptAppLockerFiles", m.BaseRequestBuilder.PathParameters)
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
// returns a *WindowsInformationProtectionPoliciesItemExemptAppLockerFilesRequestBuilder when successful
func (m *WindowsInformationProtectionPoliciesItemExemptAppLockerFilesRequestBuilder) WithUrl(rawUrl string)(*WindowsInformationProtectionPoliciesItemExemptAppLockerFilesRequestBuilder) {
    return NewWindowsInformationProtectionPoliciesItemExemptAppLockerFilesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

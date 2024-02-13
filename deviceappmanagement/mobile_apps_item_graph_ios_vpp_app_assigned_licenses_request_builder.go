package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MobileAppsItemGraphIosVppAppAssignedLicensesRequestBuilder provides operations to manage the assignedLicenses property of the microsoft.graph.iosVppApp entity.
type MobileAppsItemGraphIosVppAppAssignedLicensesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MobileAppsItemGraphIosVppAppAssignedLicensesRequestBuilderGetQueryParameters the licenses assigned to this app.
type MobileAppsItemGraphIosVppAppAssignedLicensesRequestBuilderGetQueryParameters struct {
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
// MobileAppsItemGraphIosVppAppAssignedLicensesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileAppsItemGraphIosVppAppAssignedLicensesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MobileAppsItemGraphIosVppAppAssignedLicensesRequestBuilderGetQueryParameters
}
// MobileAppsItemGraphIosVppAppAssignedLicensesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileAppsItemGraphIosVppAppAssignedLicensesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByIosVppAppAssignedLicenseId provides operations to manage the assignedLicenses property of the microsoft.graph.iosVppApp entity.
// returns a *MobileAppsItemGraphIosVppAppAssignedLicensesIosVppAppAssignedLicenseItemRequestBuilder when successful
func (m *MobileAppsItemGraphIosVppAppAssignedLicensesRequestBuilder) ByIosVppAppAssignedLicenseId(iosVppAppAssignedLicenseId string)(*MobileAppsItemGraphIosVppAppAssignedLicensesIosVppAppAssignedLicenseItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if iosVppAppAssignedLicenseId != "" {
        urlTplParams["iosVppAppAssignedLicense%2Did"] = iosVppAppAssignedLicenseId
    }
    return NewMobileAppsItemGraphIosVppAppAssignedLicensesIosVppAppAssignedLicenseItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewMobileAppsItemGraphIosVppAppAssignedLicensesRequestBuilderInternal instantiates a new MobileAppsItemGraphIosVppAppAssignedLicensesRequestBuilder and sets the default values.
func NewMobileAppsItemGraphIosVppAppAssignedLicensesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileAppsItemGraphIosVppAppAssignedLicensesRequestBuilder) {
    m := &MobileAppsItemGraphIosVppAppAssignedLicensesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/mobileApps/{mobileApp%2Did}/graph.iosVppApp/assignedLicenses{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewMobileAppsItemGraphIosVppAppAssignedLicensesRequestBuilder instantiates a new MobileAppsItemGraphIosVppAppAssignedLicensesRequestBuilder and sets the default values.
func NewMobileAppsItemGraphIosVppAppAssignedLicensesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileAppsItemGraphIosVppAppAssignedLicensesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMobileAppsItemGraphIosVppAppAssignedLicensesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *MobileAppsItemGraphIosVppAppAssignedLicensesCountRequestBuilder when successful
func (m *MobileAppsItemGraphIosVppAppAssignedLicensesRequestBuilder) Count()(*MobileAppsItemGraphIosVppAppAssignedLicensesCountRequestBuilder) {
    return NewMobileAppsItemGraphIosVppAppAssignedLicensesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the licenses assigned to this app.
// returns a IosVppAppAssignedLicenseCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MobileAppsItemGraphIosVppAppAssignedLicensesRequestBuilder) Get(ctx context.Context, requestConfiguration *MobileAppsItemGraphIosVppAppAssignedLicensesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IosVppAppAssignedLicenseCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateIosVppAppAssignedLicenseCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IosVppAppAssignedLicenseCollectionResponseable), nil
}
// Post create new navigation property to assignedLicenses for deviceAppManagement
// returns a IosVppAppAssignedLicenseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MobileAppsItemGraphIosVppAppAssignedLicensesRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IosVppAppAssignedLicenseable, requestConfiguration *MobileAppsItemGraphIosVppAppAssignedLicensesRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IosVppAppAssignedLicenseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateIosVppAppAssignedLicenseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IosVppAppAssignedLicenseable), nil
}
// ToGetRequestInformation the licenses assigned to this app.
// returns a *RequestInformation when successful
func (m *MobileAppsItemGraphIosVppAppAssignedLicensesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MobileAppsItemGraphIosVppAppAssignedLicensesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to assignedLicenses for deviceAppManagement
// returns a *RequestInformation when successful
func (m *MobileAppsItemGraphIosVppAppAssignedLicensesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IosVppAppAssignedLicenseable, requestConfiguration *MobileAppsItemGraphIosVppAppAssignedLicensesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, "{+baseurl}/deviceAppManagement/mobileApps/{mobileApp%2Did}/graph.iosVppApp/assignedLicenses", m.BaseRequestBuilder.PathParameters)
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
// returns a *MobileAppsItemGraphIosVppAppAssignedLicensesRequestBuilder when successful
func (m *MobileAppsItemGraphIosVppAppAssignedLicensesRequestBuilder) WithUrl(rawUrl string)(*MobileAppsItemGraphIosVppAppAssignedLicensesRequestBuilder) {
    return NewMobileAppsItemGraphIosVppAppAssignedLicensesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

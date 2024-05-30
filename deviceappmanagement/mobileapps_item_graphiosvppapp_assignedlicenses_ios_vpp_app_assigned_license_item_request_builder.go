package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MobileappsItemGraphiosvppappAssignedlicensesIosVppAppAssignedLicenseItemRequestBuilder provides operations to manage the assignedLicenses property of the microsoft.graph.iosVppApp entity.
type MobileappsItemGraphiosvppappAssignedlicensesIosVppAppAssignedLicenseItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MobileappsItemGraphiosvppappAssignedlicensesIosVppAppAssignedLicenseItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileappsItemGraphiosvppappAssignedlicensesIosVppAppAssignedLicenseItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// MobileappsItemGraphiosvppappAssignedlicensesIosVppAppAssignedLicenseItemRequestBuilderGetQueryParameters the licenses assigned to this app.
type MobileappsItemGraphiosvppappAssignedlicensesIosVppAppAssignedLicenseItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MobileappsItemGraphiosvppappAssignedlicensesIosVppAppAssignedLicenseItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileappsItemGraphiosvppappAssignedlicensesIosVppAppAssignedLicenseItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MobileappsItemGraphiosvppappAssignedlicensesIosVppAppAssignedLicenseItemRequestBuilderGetQueryParameters
}
// MobileappsItemGraphiosvppappAssignedlicensesIosVppAppAssignedLicenseItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileappsItemGraphiosvppappAssignedlicensesIosVppAppAssignedLicenseItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMobileappsItemGraphiosvppappAssignedlicensesIosVppAppAssignedLicenseItemRequestBuilderInternal instantiates a new MobileappsItemGraphiosvppappAssignedlicensesIosVppAppAssignedLicenseItemRequestBuilder and sets the default values.
func NewMobileappsItemGraphiosvppappAssignedlicensesIosVppAppAssignedLicenseItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileappsItemGraphiosvppappAssignedlicensesIosVppAppAssignedLicenseItemRequestBuilder) {
    m := &MobileappsItemGraphiosvppappAssignedlicensesIosVppAppAssignedLicenseItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/mobileApps/{mobileApp%2Did}/graph.iosVppApp/assignedLicenses/{iosVppAppAssignedLicense%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewMobileappsItemGraphiosvppappAssignedlicensesIosVppAppAssignedLicenseItemRequestBuilder instantiates a new MobileappsItemGraphiosvppappAssignedlicensesIosVppAppAssignedLicenseItemRequestBuilder and sets the default values.
func NewMobileappsItemGraphiosvppappAssignedlicensesIosVppAppAssignedLicenseItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileappsItemGraphiosvppappAssignedlicensesIosVppAppAssignedLicenseItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMobileappsItemGraphiosvppappAssignedlicensesIosVppAppAssignedLicenseItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property assignedLicenses for deviceAppManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MobileappsItemGraphiosvppappAssignedlicensesIosVppAppAssignedLicenseItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *MobileappsItemGraphiosvppappAssignedlicensesIosVppAppAssignedLicenseItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the licenses assigned to this app.
// returns a IosVppAppAssignedLicenseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MobileappsItemGraphiosvppappAssignedlicensesIosVppAppAssignedLicenseItemRequestBuilder) Get(ctx context.Context, requestConfiguration *MobileappsItemGraphiosvppappAssignedlicensesIosVppAppAssignedLicenseItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IosVppAppAssignedLicenseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
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
// Patch update the navigation property assignedLicenses in deviceAppManagement
// returns a IosVppAppAssignedLicenseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MobileappsItemGraphiosvppappAssignedlicensesIosVppAppAssignedLicenseItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IosVppAppAssignedLicenseable, requestConfiguration *MobileappsItemGraphiosvppappAssignedlicensesIosVppAppAssignedLicenseItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IosVppAppAssignedLicenseable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
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
// ToDeleteRequestInformation delete navigation property assignedLicenses for deviceAppManagement
// returns a *RequestInformation when successful
func (m *MobileappsItemGraphiosvppappAssignedlicensesIosVppAppAssignedLicenseItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *MobileappsItemGraphiosvppappAssignedlicensesIosVppAppAssignedLicenseItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the licenses assigned to this app.
// returns a *RequestInformation when successful
func (m *MobileappsItemGraphiosvppappAssignedlicensesIosVppAppAssignedLicenseItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MobileappsItemGraphiosvppappAssignedlicensesIosVppAppAssignedLicenseItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property assignedLicenses in deviceAppManagement
// returns a *RequestInformation when successful
func (m *MobileappsItemGraphiosvppappAssignedlicensesIosVppAppAssignedLicenseItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IosVppAppAssignedLicenseable, requestConfiguration *MobileappsItemGraphiosvppappAssignedlicensesIosVppAppAssignedLicenseItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *MobileappsItemGraphiosvppappAssignedlicensesIosVppAppAssignedLicenseItemRequestBuilder when successful
func (m *MobileappsItemGraphiosvppappAssignedlicensesIosVppAppAssignedLicenseItemRequestBuilder) WithUrl(rawUrl string)(*MobileappsItemGraphiosvppappAssignedlicensesIosVppAppAssignedLicenseItemRequestBuilder) {
    return NewMobileappsItemGraphiosvppappAssignedlicensesIosVppAppAssignedLicenseItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

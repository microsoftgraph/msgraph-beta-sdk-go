package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AndroidmanagedstoreaccountenterprisesettingsAndroidManagedStoreAccountEnterpriseSettingsRequestBuilder provides operations to manage the androidManagedStoreAccountEnterpriseSettings property of the microsoft.graph.deviceManagement entity.
type AndroidmanagedstoreaccountenterprisesettingsAndroidManagedStoreAccountEnterpriseSettingsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AndroidmanagedstoreaccountenterprisesettingsAndroidManagedStoreAccountEnterpriseSettingsRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AndroidmanagedstoreaccountenterprisesettingsAndroidManagedStoreAccountEnterpriseSettingsRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AndroidmanagedstoreaccountenterprisesettingsAndroidManagedStoreAccountEnterpriseSettingsRequestBuilderGetQueryParameters the singleton Android managed store account enterprise settings entity.
type AndroidmanagedstoreaccountenterprisesettingsAndroidManagedStoreAccountEnterpriseSettingsRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AndroidmanagedstoreaccountenterprisesettingsAndroidManagedStoreAccountEnterpriseSettingsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AndroidmanagedstoreaccountenterprisesettingsAndroidManagedStoreAccountEnterpriseSettingsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AndroidmanagedstoreaccountenterprisesettingsAndroidManagedStoreAccountEnterpriseSettingsRequestBuilderGetQueryParameters
}
// AndroidmanagedstoreaccountenterprisesettingsAndroidManagedStoreAccountEnterpriseSettingsRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AndroidmanagedstoreaccountenterprisesettingsAndroidManagedStoreAccountEnterpriseSettingsRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AddApps provides operations to call the addApps method.
// returns a *AndroidmanagedstoreaccountenterprisesettingsAddappsAddAppsRequestBuilder when successful
func (m *AndroidmanagedstoreaccountenterprisesettingsAndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) AddApps()(*AndroidmanagedstoreaccountenterprisesettingsAddappsAddAppsRequestBuilder) {
    return NewAndroidmanagedstoreaccountenterprisesettingsAddappsAddAppsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ApproveApps provides operations to call the approveApps method.
// returns a *AndroidmanagedstoreaccountenterprisesettingsApproveappsApproveAppsRequestBuilder when successful
func (m *AndroidmanagedstoreaccountenterprisesettingsAndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) ApproveApps()(*AndroidmanagedstoreaccountenterprisesettingsApproveappsApproveAppsRequestBuilder) {
    return NewAndroidmanagedstoreaccountenterprisesettingsApproveappsApproveAppsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CompleteSignup provides operations to call the completeSignup method.
// returns a *AndroidmanagedstoreaccountenterprisesettingsCompletesignupCompleteSignupRequestBuilder when successful
func (m *AndroidmanagedstoreaccountenterprisesettingsAndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) CompleteSignup()(*AndroidmanagedstoreaccountenterprisesettingsCompletesignupCompleteSignupRequestBuilder) {
    return NewAndroidmanagedstoreaccountenterprisesettingsCompletesignupCompleteSignupRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewAndroidmanagedstoreaccountenterprisesettingsAndroidManagedStoreAccountEnterpriseSettingsRequestBuilderInternal instantiates a new AndroidmanagedstoreaccountenterprisesettingsAndroidManagedStoreAccountEnterpriseSettingsRequestBuilder and sets the default values.
func NewAndroidmanagedstoreaccountenterprisesettingsAndroidManagedStoreAccountEnterpriseSettingsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AndroidmanagedstoreaccountenterprisesettingsAndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) {
    m := &AndroidmanagedstoreaccountenterprisesettingsAndroidManagedStoreAccountEnterpriseSettingsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/androidManagedStoreAccountEnterpriseSettings{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewAndroidmanagedstoreaccountenterprisesettingsAndroidManagedStoreAccountEnterpriseSettingsRequestBuilder instantiates a new AndroidmanagedstoreaccountenterprisesettingsAndroidManagedStoreAccountEnterpriseSettingsRequestBuilder and sets the default values.
func NewAndroidmanagedstoreaccountenterprisesettingsAndroidManagedStoreAccountEnterpriseSettingsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AndroidmanagedstoreaccountenterprisesettingsAndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAndroidmanagedstoreaccountenterprisesettingsAndroidManagedStoreAccountEnterpriseSettingsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGooglePlayWebToken provides operations to call the createGooglePlayWebToken method.
// returns a *AndroidmanagedstoreaccountenterprisesettingsCreategoogleplaywebtokenCreateGooglePlayWebTokenRequestBuilder when successful
func (m *AndroidmanagedstoreaccountenterprisesettingsAndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) CreateGooglePlayWebToken()(*AndroidmanagedstoreaccountenterprisesettingsCreategoogleplaywebtokenCreateGooglePlayWebTokenRequestBuilder) {
    return NewAndroidmanagedstoreaccountenterprisesettingsCreategoogleplaywebtokenCreateGooglePlayWebTokenRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property androidManagedStoreAccountEnterpriseSettings for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AndroidmanagedstoreaccountenterprisesettingsAndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) Delete(ctx context.Context, requestConfiguration *AndroidmanagedstoreaccountenterprisesettingsAndroidManagedStoreAccountEnterpriseSettingsRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the singleton Android managed store account enterprise settings entity.
// returns a AndroidManagedStoreAccountEnterpriseSettingsable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AndroidmanagedstoreaccountenterprisesettingsAndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) Get(ctx context.Context, requestConfiguration *AndroidmanagedstoreaccountenterprisesettingsAndroidManagedStoreAccountEnterpriseSettingsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AndroidManagedStoreAccountEnterpriseSettingsable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAndroidManagedStoreAccountEnterpriseSettingsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AndroidManagedStoreAccountEnterpriseSettingsable), nil
}
// Patch update the navigation property androidManagedStoreAccountEnterpriseSettings in deviceManagement
// returns a AndroidManagedStoreAccountEnterpriseSettingsable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AndroidmanagedstoreaccountenterprisesettingsAndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AndroidManagedStoreAccountEnterpriseSettingsable, requestConfiguration *AndroidmanagedstoreaccountenterprisesettingsAndroidManagedStoreAccountEnterpriseSettingsRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AndroidManagedStoreAccountEnterpriseSettingsable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAndroidManagedStoreAccountEnterpriseSettingsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AndroidManagedStoreAccountEnterpriseSettingsable), nil
}
// RequestSignupUrl provides operations to call the requestSignupUrl method.
// returns a *AndroidmanagedstoreaccountenterprisesettingsRequestsignupurlRequestSignupUrlRequestBuilder when successful
func (m *AndroidmanagedstoreaccountenterprisesettingsAndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) RequestSignupUrl()(*AndroidmanagedstoreaccountenterprisesettingsRequestsignupurlRequestSignupUrlRequestBuilder) {
    return NewAndroidmanagedstoreaccountenterprisesettingsRequestsignupurlRequestSignupUrlRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SetAndroidDeviceOwnerFullyManagedEnrollmentState provides operations to call the setAndroidDeviceOwnerFullyManagedEnrollmentState method.
// returns a *AndroidmanagedstoreaccountenterprisesettingsSetandroiddeviceownerfullymanagedenrollmentstateSetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBuilder when successful
func (m *AndroidmanagedstoreaccountenterprisesettingsAndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) SetAndroidDeviceOwnerFullyManagedEnrollmentState()(*AndroidmanagedstoreaccountenterprisesettingsSetandroiddeviceownerfullymanagedenrollmentstateSetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBuilder) {
    return NewAndroidmanagedstoreaccountenterprisesettingsSetandroiddeviceownerfullymanagedenrollmentstateSetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SyncApps provides operations to call the syncApps method.
// returns a *AndroidmanagedstoreaccountenterprisesettingsSyncappsSyncAppsRequestBuilder when successful
func (m *AndroidmanagedstoreaccountenterprisesettingsAndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) SyncApps()(*AndroidmanagedstoreaccountenterprisesettingsSyncappsSyncAppsRequestBuilder) {
    return NewAndroidmanagedstoreaccountenterprisesettingsSyncappsSyncAppsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property androidManagedStoreAccountEnterpriseSettings for deviceManagement
// returns a *RequestInformation when successful
func (m *AndroidmanagedstoreaccountenterprisesettingsAndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *AndroidmanagedstoreaccountenterprisesettingsAndroidManagedStoreAccountEnterpriseSettingsRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the singleton Android managed store account enterprise settings entity.
// returns a *RequestInformation when successful
func (m *AndroidmanagedstoreaccountenterprisesettingsAndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AndroidmanagedstoreaccountenterprisesettingsAndroidManagedStoreAccountEnterpriseSettingsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property androidManagedStoreAccountEnterpriseSettings in deviceManagement
// returns a *RequestInformation when successful
func (m *AndroidmanagedstoreaccountenterprisesettingsAndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AndroidManagedStoreAccountEnterpriseSettingsable, requestConfiguration *AndroidmanagedstoreaccountenterprisesettingsAndroidManagedStoreAccountEnterpriseSettingsRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Unbind provides operations to call the unbind method.
// returns a *AndroidmanagedstoreaccountenterprisesettingsUnbindRequestBuilder when successful
func (m *AndroidmanagedstoreaccountenterprisesettingsAndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) Unbind()(*AndroidmanagedstoreaccountenterprisesettingsUnbindRequestBuilder) {
    return NewAndroidmanagedstoreaccountenterprisesettingsUnbindRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *AndroidmanagedstoreaccountenterprisesettingsAndroidManagedStoreAccountEnterpriseSettingsRequestBuilder when successful
func (m *AndroidmanagedstoreaccountenterprisesettingsAndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) WithUrl(rawUrl string)(*AndroidmanagedstoreaccountenterprisesettingsAndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) {
    return NewAndroidmanagedstoreaccountenterprisesettingsAndroidManagedStoreAccountEnterpriseSettingsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

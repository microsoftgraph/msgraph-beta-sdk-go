package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DepOnboardingSettingsItemImportedAppleDeviceIdentitiesImportAppleDeviceIdentityListRequestBuilder provides operations to call the importAppleDeviceIdentityList method.
type DepOnboardingSettingsItemImportedAppleDeviceIdentitiesImportAppleDeviceIdentityListRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DepOnboardingSettingsItemImportedAppleDeviceIdentitiesImportAppleDeviceIdentityListRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DepOnboardingSettingsItemImportedAppleDeviceIdentitiesImportAppleDeviceIdentityListRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDepOnboardingSettingsItemImportedAppleDeviceIdentitiesImportAppleDeviceIdentityListRequestBuilderInternal instantiates a new ImportAppleDeviceIdentityListRequestBuilder and sets the default values.
func NewDepOnboardingSettingsItemImportedAppleDeviceIdentitiesImportAppleDeviceIdentityListRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DepOnboardingSettingsItemImportedAppleDeviceIdentitiesImportAppleDeviceIdentityListRequestBuilder) {
    m := &DepOnboardingSettingsItemImportedAppleDeviceIdentitiesImportAppleDeviceIdentityListRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/depOnboardingSettings/{depOnboardingSetting%2Did}/importedAppleDeviceIdentities/importAppleDeviceIdentityList", pathParameters),
    }
    return m
}
// NewDepOnboardingSettingsItemImportedAppleDeviceIdentitiesImportAppleDeviceIdentityListRequestBuilder instantiates a new ImportAppleDeviceIdentityListRequestBuilder and sets the default values.
func NewDepOnboardingSettingsItemImportedAppleDeviceIdentitiesImportAppleDeviceIdentityListRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DepOnboardingSettingsItemImportedAppleDeviceIdentitiesImportAppleDeviceIdentityListRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDepOnboardingSettingsItemImportedAppleDeviceIdentitiesImportAppleDeviceIdentityListRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action importAppleDeviceIdentityList
// Deprecated: This method is obsolete. Use PostAsImportAppleDeviceIdentityListPostResponse instead.
func (m *DepOnboardingSettingsItemImportedAppleDeviceIdentitiesImportAppleDeviceIdentityListRequestBuilder) Post(ctx context.Context, body DepOnboardingSettingsItemImportedAppleDeviceIdentitiesImportAppleDeviceIdentityListPostRequestBodyable, requestConfiguration *DepOnboardingSettingsItemImportedAppleDeviceIdentitiesImportAppleDeviceIdentityListRequestBuilderPostRequestConfiguration)(DepOnboardingSettingsItemImportedAppleDeviceIdentitiesImportAppleDeviceIdentityListResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateDepOnboardingSettingsItemImportedAppleDeviceIdentitiesImportAppleDeviceIdentityListResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(DepOnboardingSettingsItemImportedAppleDeviceIdentitiesImportAppleDeviceIdentityListResponseable), nil
}
// PostAsImportAppleDeviceIdentityListPostResponse invoke action importAppleDeviceIdentityList
func (m *DepOnboardingSettingsItemImportedAppleDeviceIdentitiesImportAppleDeviceIdentityListRequestBuilder) PostAsImportAppleDeviceIdentityListPostResponse(ctx context.Context, body DepOnboardingSettingsItemImportedAppleDeviceIdentitiesImportAppleDeviceIdentityListPostRequestBodyable, requestConfiguration *DepOnboardingSettingsItemImportedAppleDeviceIdentitiesImportAppleDeviceIdentityListRequestBuilderPostRequestConfiguration)(DepOnboardingSettingsItemImportedAppleDeviceIdentitiesImportAppleDeviceIdentityListPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateDepOnboardingSettingsItemImportedAppleDeviceIdentitiesImportAppleDeviceIdentityListPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(DepOnboardingSettingsItemImportedAppleDeviceIdentitiesImportAppleDeviceIdentityListPostResponseable), nil
}
// ToPostRequestInformation invoke action importAppleDeviceIdentityList
func (m *DepOnboardingSettingsItemImportedAppleDeviceIdentitiesImportAppleDeviceIdentityListRequestBuilder) ToPostRequestInformation(ctx context.Context, body DepOnboardingSettingsItemImportedAppleDeviceIdentitiesImportAppleDeviceIdentityListPostRequestBodyable, requestConfiguration *DepOnboardingSettingsItemImportedAppleDeviceIdentitiesImportAppleDeviceIdentityListRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *DepOnboardingSettingsItemImportedAppleDeviceIdentitiesImportAppleDeviceIdentityListRequestBuilder) WithUrl(rawUrl string)(*DepOnboardingSettingsItemImportedAppleDeviceIdentitiesImportAppleDeviceIdentityListRequestBuilder) {
    return NewDepOnboardingSettingsItemImportedAppleDeviceIdentitiesImportAppleDeviceIdentityListRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

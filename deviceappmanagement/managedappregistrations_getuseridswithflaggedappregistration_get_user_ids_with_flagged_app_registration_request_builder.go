package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ManagedappregistrationsGetuseridswithflaggedappregistrationGetUserIdsWithFlaggedAppRegistrationRequestBuilder provides operations to call the getUserIdsWithFlaggedAppRegistration method.
type ManagedappregistrationsGetuseridswithflaggedappregistrationGetUserIdsWithFlaggedAppRegistrationRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManagedappregistrationsGetuseridswithflaggedappregistrationGetUserIdsWithFlaggedAppRegistrationRequestBuilderGetQueryParameters invoke function getUserIdsWithFlaggedAppRegistration
type ManagedappregistrationsGetuseridswithflaggedappregistrationGetUserIdsWithFlaggedAppRegistrationRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// ManagedappregistrationsGetuseridswithflaggedappregistrationGetUserIdsWithFlaggedAppRegistrationRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedappregistrationsGetuseridswithflaggedappregistrationGetUserIdsWithFlaggedAppRegistrationRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ManagedappregistrationsGetuseridswithflaggedappregistrationGetUserIdsWithFlaggedAppRegistrationRequestBuilderGetQueryParameters
}
// NewManagedappregistrationsGetuseridswithflaggedappregistrationGetUserIdsWithFlaggedAppRegistrationRequestBuilderInternal instantiates a new ManagedappregistrationsGetuseridswithflaggedappregistrationGetUserIdsWithFlaggedAppRegistrationRequestBuilder and sets the default values.
func NewManagedappregistrationsGetuseridswithflaggedappregistrationGetUserIdsWithFlaggedAppRegistrationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedappregistrationsGetuseridswithflaggedappregistrationGetUserIdsWithFlaggedAppRegistrationRequestBuilder) {
    m := &ManagedappregistrationsGetuseridswithflaggedappregistrationGetUserIdsWithFlaggedAppRegistrationRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/managedAppRegistrations/getUserIdsWithFlaggedAppRegistration(){?%24count,%24filter,%24search,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewManagedappregistrationsGetuseridswithflaggedappregistrationGetUserIdsWithFlaggedAppRegistrationRequestBuilder instantiates a new ManagedappregistrationsGetuseridswithflaggedappregistrationGetUserIdsWithFlaggedAppRegistrationRequestBuilder and sets the default values.
func NewManagedappregistrationsGetuseridswithflaggedappregistrationGetUserIdsWithFlaggedAppRegistrationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedappregistrationsGetuseridswithflaggedappregistrationGetUserIdsWithFlaggedAppRegistrationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedappregistrationsGetuseridswithflaggedappregistrationGetUserIdsWithFlaggedAppRegistrationRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function getUserIdsWithFlaggedAppRegistration
// Deprecated: This method is obsolete. Use GetAsGetUserIdsWithFlaggedAppRegistrationGetResponse instead.
// returns a ManagedappregistrationsGetuseridswithflaggedappregistrationGetUserIdsWithFlaggedAppRegistrationResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManagedappregistrationsGetuseridswithflaggedappregistrationGetUserIdsWithFlaggedAppRegistrationRequestBuilder) Get(ctx context.Context, requestConfiguration *ManagedappregistrationsGetuseridswithflaggedappregistrationGetUserIdsWithFlaggedAppRegistrationRequestBuilderGetRequestConfiguration)(ManagedappregistrationsGetuseridswithflaggedappregistrationGetUserIdsWithFlaggedAppRegistrationResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateManagedappregistrationsGetuseridswithflaggedappregistrationGetUserIdsWithFlaggedAppRegistrationResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ManagedappregistrationsGetuseridswithflaggedappregistrationGetUserIdsWithFlaggedAppRegistrationResponseable), nil
}
// GetAsGetUserIdsWithFlaggedAppRegistrationGetResponse invoke function getUserIdsWithFlaggedAppRegistration
// returns a ManagedappregistrationsGetuseridswithflaggedappregistrationGetUserIdsWithFlaggedAppRegistrationGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManagedappregistrationsGetuseridswithflaggedappregistrationGetUserIdsWithFlaggedAppRegistrationRequestBuilder) GetAsGetUserIdsWithFlaggedAppRegistrationGetResponse(ctx context.Context, requestConfiguration *ManagedappregistrationsGetuseridswithflaggedappregistrationGetUserIdsWithFlaggedAppRegistrationRequestBuilderGetRequestConfiguration)(ManagedappregistrationsGetuseridswithflaggedappregistrationGetUserIdsWithFlaggedAppRegistrationGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateManagedappregistrationsGetuseridswithflaggedappregistrationGetUserIdsWithFlaggedAppRegistrationGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ManagedappregistrationsGetuseridswithflaggedappregistrationGetUserIdsWithFlaggedAppRegistrationGetResponseable), nil
}
// ToGetRequestInformation invoke function getUserIdsWithFlaggedAppRegistration
// returns a *RequestInformation when successful
func (m *ManagedappregistrationsGetuseridswithflaggedappregistrationGetUserIdsWithFlaggedAppRegistrationRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ManagedappregistrationsGetuseridswithflaggedappregistrationGetUserIdsWithFlaggedAppRegistrationRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ManagedappregistrationsGetuseridswithflaggedappregistrationGetUserIdsWithFlaggedAppRegistrationRequestBuilder when successful
func (m *ManagedappregistrationsGetuseridswithflaggedappregistrationGetUserIdsWithFlaggedAppRegistrationRequestBuilder) WithUrl(rawUrl string)(*ManagedappregistrationsGetuseridswithflaggedappregistrationGetUserIdsWithFlaggedAppRegistrationRequestBuilder) {
    return NewManagedappregistrationsGetuseridswithflaggedappregistrationGetUserIdsWithFlaggedAppRegistrationRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MicrosoftGraphGetCredentialUserRegistrationCountGetCredentialUserRegistrationCountRequestBuilder provides operations to call the getCredentialUserRegistrationCount method.
type MicrosoftGraphGetCredentialUserRegistrationCountGetCredentialUserRegistrationCountRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MicrosoftGraphGetCredentialUserRegistrationCountGetCredentialUserRegistrationCountRequestBuilderGetQueryParameters report the current state of how many users in your organization are registered for self-service password reset and multi-factor authentication (MFA) capabilities.
type MicrosoftGraphGetCredentialUserRegistrationCountGetCredentialUserRegistrationCountRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
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
// MicrosoftGraphGetCredentialUserRegistrationCountGetCredentialUserRegistrationCountRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MicrosoftGraphGetCredentialUserRegistrationCountGetCredentialUserRegistrationCountRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MicrosoftGraphGetCredentialUserRegistrationCountGetCredentialUserRegistrationCountRequestBuilderGetQueryParameters
}
// NewMicrosoftGraphGetCredentialUserRegistrationCountGetCredentialUserRegistrationCountRequestBuilderInternal instantiates a new GetCredentialUserRegistrationCountRequestBuilder and sets the default values.
func NewMicrosoftGraphGetCredentialUserRegistrationCountGetCredentialUserRegistrationCountRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MicrosoftGraphGetCredentialUserRegistrationCountGetCredentialUserRegistrationCountRequestBuilder) {
    m := &MicrosoftGraphGetCredentialUserRegistrationCountGetCredentialUserRegistrationCountRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/reports/microsoft.graph.getCredentialUserRegistrationCount(){?%24top,%24skip,%24search,%24filter,%24count,%24select,%24orderby}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewMicrosoftGraphGetCredentialUserRegistrationCountGetCredentialUserRegistrationCountRequestBuilder instantiates a new GetCredentialUserRegistrationCountRequestBuilder and sets the default values.
func NewMicrosoftGraphGetCredentialUserRegistrationCountGetCredentialUserRegistrationCountRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MicrosoftGraphGetCredentialUserRegistrationCountGetCredentialUserRegistrationCountRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMicrosoftGraphGetCredentialUserRegistrationCountGetCredentialUserRegistrationCountRequestBuilderInternal(urlParams, requestAdapter)
}
// Get report the current state of how many users in your organization are registered for self-service password reset and multi-factor authentication (MFA) capabilities.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/reportroot-getcredentialuserregistrationcount?view=graph-rest-1.0
func (m *MicrosoftGraphGetCredentialUserRegistrationCountGetCredentialUserRegistrationCountRequestBuilder) Get(ctx context.Context, requestConfiguration *MicrosoftGraphGetCredentialUserRegistrationCountGetCredentialUserRegistrationCountRequestBuilderGetRequestConfiguration)(MicrosoftGraphGetCredentialUserRegistrationCountGetCredentialUserRegistrationCountResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateMicrosoftGraphGetCredentialUserRegistrationCountGetCredentialUserRegistrationCountResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(MicrosoftGraphGetCredentialUserRegistrationCountGetCredentialUserRegistrationCountResponseable), nil
}
// ToGetRequestInformation report the current state of how many users in your organization are registered for self-service password reset and multi-factor authentication (MFA) capabilities.
func (m *MicrosoftGraphGetCredentialUserRegistrationCountGetCredentialUserRegistrationCountRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MicrosoftGraphGetCredentialUserRegistrationCountGetCredentialUserRegistrationCountRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
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

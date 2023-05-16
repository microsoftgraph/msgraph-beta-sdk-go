package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CasesEdiscoveryCasesItemSettingsMicrosoftGraphSecurityResetToDefaultRequestBuilder provides operations to call the resetToDefault method.
type CasesEdiscoveryCasesItemSettingsMicrosoftGraphSecurityResetToDefaultRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CasesEdiscoveryCasesItemSettingsMicrosoftGraphSecurityResetToDefaultRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoveryCasesItemSettingsMicrosoftGraphSecurityResetToDefaultRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCasesEdiscoveryCasesItemSettingsMicrosoftGraphSecurityResetToDefaultRequestBuilderInternal instantiates a new MicrosoftGraphSecurityResetToDefaultRequestBuilder and sets the default values.
func NewCasesEdiscoveryCasesItemSettingsMicrosoftGraphSecurityResetToDefaultRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoveryCasesItemSettingsMicrosoftGraphSecurityResetToDefaultRequestBuilder) {
    m := &CasesEdiscoveryCasesItemSettingsMicrosoftGraphSecurityResetToDefaultRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/cases/ediscoveryCases/{ediscoveryCase%2Did}/settings/microsoft.graph.security.resetToDefault", pathParameters),
    }
    return m
}
// NewCasesEdiscoveryCasesItemSettingsMicrosoftGraphSecurityResetToDefaultRequestBuilder instantiates a new MicrosoftGraphSecurityResetToDefaultRequestBuilder and sets the default values.
func NewCasesEdiscoveryCasesItemSettingsMicrosoftGraphSecurityResetToDefaultRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoveryCasesItemSettingsMicrosoftGraphSecurityResetToDefaultRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCasesEdiscoveryCasesItemSettingsMicrosoftGraphSecurityResetToDefaultRequestBuilderInternal(urlParams, requestAdapter)
}
// Post reset a caseSettings object to the default values.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/security-ediscoverycasesettings-resettodefault?view=graph-rest-1.0
func (m *CasesEdiscoveryCasesItemSettingsMicrosoftGraphSecurityResetToDefaultRequestBuilder) Post(ctx context.Context, requestConfiguration *CasesEdiscoveryCasesItemSettingsMicrosoftGraphSecurityResetToDefaultRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation reset a caseSettings object to the default values.
func (m *CasesEdiscoveryCasesItemSettingsMicrosoftGraphSecurityResetToDefaultRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *CasesEdiscoveryCasesItemSettingsMicrosoftGraphSecurityResetToDefaultRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}

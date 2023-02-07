package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EntitlementManagementAccessPackageAssignmentsItemAccessPackageMicrosoftGraphMoveToCatalogRequestBuilder provides operations to call the moveToCatalog method.
type EntitlementManagementAccessPackageAssignmentsItemAccessPackageMicrosoftGraphMoveToCatalogRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// EntitlementManagementAccessPackageAssignmentsItemAccessPackageMicrosoftGraphMoveToCatalogRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementAccessPackageAssignmentsItemAccessPackageMicrosoftGraphMoveToCatalogRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewEntitlementManagementAccessPackageAssignmentsItemAccessPackageMicrosoftGraphMoveToCatalogRequestBuilderInternal instantiates a new MicrosoftGraphMoveToCatalogRequestBuilder and sets the default values.
func NewEntitlementManagementAccessPackageAssignmentsItemAccessPackageMicrosoftGraphMoveToCatalogRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementAccessPackageAssignmentsItemAccessPackageMicrosoftGraphMoveToCatalogRequestBuilder) {
    m := &EntitlementManagementAccessPackageAssignmentsItemAccessPackageMicrosoftGraphMoveToCatalogRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/entitlementManagement/accessPackageAssignments/{accessPackageAssignment%2Did}/accessPackage/microsoft.graph.moveToCatalog";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewEntitlementManagementAccessPackageAssignmentsItemAccessPackageMicrosoftGraphMoveToCatalogRequestBuilder instantiates a new MicrosoftGraphMoveToCatalogRequestBuilder and sets the default values.
func NewEntitlementManagementAccessPackageAssignmentsItemAccessPackageMicrosoftGraphMoveToCatalogRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementAccessPackageAssignmentsItemAccessPackageMicrosoftGraphMoveToCatalogRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementManagementAccessPackageAssignmentsItemAccessPackageMicrosoftGraphMoveToCatalogRequestBuilderInternal(urlParams, requestAdapter)
}
// Post in Azure AD entitlement management, this action moves the accessPackage to a specified target accessPackageCatalog. The resources in the access package must be present in the target catalog.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/accesspackage-movetocatalog?view=graph-rest-1.0
func (m *EntitlementManagementAccessPackageAssignmentsItemAccessPackageMicrosoftGraphMoveToCatalogRequestBuilder) Post(ctx context.Context, body EntitlementManagementAccessPackageAssignmentsItemAccessPackageMicrosoftGraphMoveToCatalogMoveToCatalogPostRequestBodyable, requestConfiguration *EntitlementManagementAccessPackageAssignmentsItemAccessPackageMicrosoftGraphMoveToCatalogRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation in Azure AD entitlement management, this action moves the accessPackage to a specified target accessPackageCatalog. The resources in the access package must be present in the target catalog.
func (m *EntitlementManagementAccessPackageAssignmentsItemAccessPackageMicrosoftGraphMoveToCatalogRequestBuilder) ToPostRequestInformation(ctx context.Context, body EntitlementManagementAccessPackageAssignmentsItemAccessPackageMicrosoftGraphMoveToCatalogMoveToCatalogPostRequestBodyable, requestConfiguration *EntitlementManagementAccessPackageAssignmentsItemAccessPackageMicrosoftGraphMoveToCatalogRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    err := requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}

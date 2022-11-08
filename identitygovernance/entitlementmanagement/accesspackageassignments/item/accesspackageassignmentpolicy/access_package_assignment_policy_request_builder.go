package accesspackageassignmentpolicy

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i400de42849e0a7ccdeb75bfc507a87b90b2ec716cfe648850d153b61b7f48534 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignments/item/accesspackageassignmentpolicy/accesspackagecatalog"
    ibb60fc0712bcf899bd3e6ad5bc54dd0275b26bc20b2995c0596dab995710309f "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignments/item/accesspackageassignmentpolicy/accesspackage"
    ie980aae78a11dcf69f861f2a118cab5a872cafe44fcf59d6604e106460ccc8ce "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignments/item/accesspackageassignmentpolicy/customextensionhandlers"
    i7054a475d07c7c3f753dd11cd7cd0a3c3ff20b9f88829ab23af4a9a9721a3735 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignments/item/accesspackageassignmentpolicy/customextensionhandlers/item"
)

// AccessPackageAssignmentPolicyRequestBuilder provides operations to manage the accessPackageAssignmentPolicy property of the microsoft.graph.accessPackageAssignment entity.
type AccessPackageAssignmentPolicyRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// AccessPackageAssignmentPolicyRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AccessPackageAssignmentPolicyRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AccessPackageAssignmentPolicyRequestBuilderGetQueryParameters read-only. Nullable. Supports $filter (eq) on the id property
type AccessPackageAssignmentPolicyRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AccessPackageAssignmentPolicyRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AccessPackageAssignmentPolicyRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AccessPackageAssignmentPolicyRequestBuilderGetQueryParameters
}
// AccessPackageAssignmentPolicyRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AccessPackageAssignmentPolicyRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AccessPackage provides operations to manage the accessPackage property of the microsoft.graph.accessPackageAssignmentPolicy entity.
func (m *AccessPackageAssignmentPolicyRequestBuilder) AccessPackage()(*ibb60fc0712bcf899bd3e6ad5bc54dd0275b26bc20b2995c0596dab995710309f.AccessPackageRequestBuilder) {
    return ibb60fc0712bcf899bd3e6ad5bc54dd0275b26bc20b2995c0596dab995710309f.NewAccessPackageRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageCatalog provides operations to manage the accessPackageCatalog property of the microsoft.graph.accessPackageAssignmentPolicy entity.
func (m *AccessPackageAssignmentPolicyRequestBuilder) AccessPackageCatalog()(*i400de42849e0a7ccdeb75bfc507a87b90b2ec716cfe648850d153b61b7f48534.AccessPackageCatalogRequestBuilder) {
    return i400de42849e0a7ccdeb75bfc507a87b90b2ec716cfe648850d153b61b7f48534.NewAccessPackageCatalogRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewAccessPackageAssignmentPolicyRequestBuilderInternal instantiates a new AccessPackageAssignmentPolicyRequestBuilder and sets the default values.
func NewAccessPackageAssignmentPolicyRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccessPackageAssignmentPolicyRequestBuilder) {
    m := &AccessPackageAssignmentPolicyRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/entitlementManagement/accessPackageAssignments/{accessPackageAssignment%2Did}/accessPackageAssignmentPolicy{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAccessPackageAssignmentPolicyRequestBuilder instantiates a new AccessPackageAssignmentPolicyRequestBuilder and sets the default values.
func NewAccessPackageAssignmentPolicyRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccessPackageAssignmentPolicyRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAccessPackageAssignmentPolicyRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property accessPackageAssignmentPolicy for identityGovernance
func (m *AccessPackageAssignmentPolicyRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *AccessPackageAssignmentPolicyRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation read-only. Nullable. Supports $filter (eq) on the id property
func (m *AccessPackageAssignmentPolicyRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *AccessPackageAssignmentPolicyRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property accessPackageAssignmentPolicy in identityGovernance
func (m *AccessPackageAssignmentPolicyRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageAssignmentPolicyable, requestConfiguration *AccessPackageAssignmentPolicyRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CustomExtensionHandlers provides operations to manage the customExtensionHandlers property of the microsoft.graph.accessPackageAssignmentPolicy entity.
func (m *AccessPackageAssignmentPolicyRequestBuilder) CustomExtensionHandlers()(*ie980aae78a11dcf69f861f2a118cab5a872cafe44fcf59d6604e106460ccc8ce.CustomExtensionHandlersRequestBuilder) {
    return ie980aae78a11dcf69f861f2a118cab5a872cafe44fcf59d6604e106460ccc8ce.NewCustomExtensionHandlersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CustomExtensionHandlersById provides operations to manage the customExtensionHandlers property of the microsoft.graph.accessPackageAssignmentPolicy entity.
func (m *AccessPackageAssignmentPolicyRequestBuilder) CustomExtensionHandlersById(id string)(*i7054a475d07c7c3f753dd11cd7cd0a3c3ff20b9f88829ab23af4a9a9721a3735.CustomExtensionHandlerItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["customExtensionHandler%2Did"] = id
    }
    return i7054a475d07c7c3f753dd11cd7cd0a3c3ff20b9f88829ab23af4a9a9721a3735.NewCustomExtensionHandlerItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Delete delete navigation property accessPackageAssignmentPolicy for identityGovernance
func (m *AccessPackageAssignmentPolicyRequestBuilder) Delete(ctx context.Context, requestConfiguration *AccessPackageAssignmentPolicyRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get read-only. Nullable. Supports $filter (eq) on the id property
func (m *AccessPackageAssignmentPolicyRequestBuilder) Get(ctx context.Context, requestConfiguration *AccessPackageAssignmentPolicyRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageAssignmentPolicyable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessPackageAssignmentPolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageAssignmentPolicyable), nil
}
// Patch update the navigation property accessPackageAssignmentPolicy in identityGovernance
func (m *AccessPackageAssignmentPolicyRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageAssignmentPolicyable, requestConfiguration *AccessPackageAssignmentPolicyRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageAssignmentPolicyable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessPackageAssignmentPolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageAssignmentPolicyable), nil
}

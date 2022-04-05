package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i34a5e9b6cb103e1786cc9bc0844b427d915cfd9f19ccaf56b5f0f1bab78f824d "github.com/microsoftgraph/msgraph-beta-sdk-go/print/printershares/item/printer"
    i8a20ab890b50d7d004d72ec88971b0682a50e050460da1b743a65b594f56cb56 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/printershares/item/allowedusers"
    id1d036967566825b0af9120492b03a4ec20e8592dbc04e3ef967f4a9ea9fd161 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/printershares/item/allowedgroups"
    i4dfb95df95625fd76c75c0aafc86672d4f8e40be9f132a0da1c2e38ada8c6ed3 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/printershares/item/allowedusers/item"
    i6f9a647c0d781a0a6d5b6ee2240654bf7d6d240213b65921a915d72a25c48568 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/printershares/item/allowedgroups/item"
)

// PrinterShareItemRequestBuilder provides operations to manage the printerShares property of the microsoft.graph.print entity.
type PrinterShareItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// PrinterShareItemRequestBuilderDeleteOptions options for Delete
type PrinterShareItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// PrinterShareItemRequestBuilderGetOptions options for Get
type PrinterShareItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Request query parameters
    QueryParameters *PrinterShareItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// PrinterShareItemRequestBuilderGetQueryParameters get printerShares from print
type PrinterShareItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// PrinterShareItemRequestBuilderPatchOptions options for Patch
type PrinterShareItemRequestBuilderPatchOptions struct {
    // 
    Body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrinterShareable;
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// AllowedGroups the allowedGroups property
func (m *PrinterShareItemRequestBuilder) AllowedGroups()(*id1d036967566825b0af9120492b03a4ec20e8592dbc04e3ef967f4a9ea9fd161.AllowedGroupsRequestBuilder) {
    return id1d036967566825b0af9120492b03a4ec20e8592dbc04e3ef967f4a9ea9fd161.NewAllowedGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AllowedGroupsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.print.printerShares.item.allowedGroups.item collection
func (m *PrinterShareItemRequestBuilder) AllowedGroupsById(id string)(*i6f9a647c0d781a0a6d5b6ee2240654bf7d6d240213b65921a915d72a25c48568.GroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["group_id"] = id
    }
    return i6f9a647c0d781a0a6d5b6ee2240654bf7d6d240213b65921a915d72a25c48568.NewGroupItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AllowedUsers the allowedUsers property
func (m *PrinterShareItemRequestBuilder) AllowedUsers()(*i8a20ab890b50d7d004d72ec88971b0682a50e050460da1b743a65b594f56cb56.AllowedUsersRequestBuilder) {
    return i8a20ab890b50d7d004d72ec88971b0682a50e050460da1b743a65b594f56cb56.NewAllowedUsersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AllowedUsersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.print.printerShares.item.allowedUsers.item collection
func (m *PrinterShareItemRequestBuilder) AllowedUsersById(id string)(*i4dfb95df95625fd76c75c0aafc86672d4f8e40be9f132a0da1c2e38ada8c6ed3.UserItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["user_id"] = id
    }
    return i4dfb95df95625fd76c75c0aafc86672d4f8e40be9f132a0da1c2e38ada8c6ed3.NewUserItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewPrinterShareItemRequestBuilderInternal instantiates a new PrinterShareItemRequestBuilder and sets the default values.
func NewPrinterShareItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrinterShareItemRequestBuilder) {
    m := &PrinterShareItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/print/printerShares/{printerShare_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewPrinterShareItemRequestBuilder instantiates a new PrinterShareItemRequestBuilder and sets the default values.
func NewPrinterShareItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrinterShareItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrinterShareItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property printerShares for print
func (m *PrinterShareItemRequestBuilder) CreateDeleteRequestInformation(options *PrinterShareItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation get printerShares from print
func (m *PrinterShareItemRequestBuilder) CreateGetRequestInformation(options *PrinterShareItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.QueryParameters != nil {
        requestInfo.AddQueryParameters(*(options.QueryParameters))
    }
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property printerShares in print
func (m *PrinterShareItemRequestBuilder) CreatePatchRequestInformation(options *PrinterShareItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete delete navigation property printerShares for print
func (m *PrinterShareItemRequestBuilder) Delete(options *PrinterShareItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get printerShares from print
func (m *PrinterShareItemRequestBuilder) Get(options *PrinterShareItemRequestBuilderGetOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrinterShareable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePrinterShareFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrinterShareable), nil
}
// Patch update the navigation property printerShares in print
func (m *PrinterShareItemRequestBuilder) Patch(options *PrinterShareItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Printer the printer property
func (m *PrinterShareItemRequestBuilder) Printer()(*i34a5e9b6cb103e1786cc9bc0844b427d915cfd9f19ccaf56b5f0f1bab78f824d.PrinterRequestBuilder) {
    return i34a5e9b6cb103e1786cc9bc0844b427d915cfd9f19ccaf56b5f0f1bab78f824d.NewPrinterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}

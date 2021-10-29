package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i005a789b66b49c848e74abb76f962f75f90f0a0145cf80a3da3ca3a26ff5a4ba "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/onenote/sectiongroups/item/parentnotebook/sections/item/pages/item/preview"
    i07a5432dfb3bd794fb35535d0db6f7c9e440efeea131949418329ba82c6e3435 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/onenote/sectiongroups/item/parentnotebook/sections/item/pages/item/onenotepatchcontent"
    i6767e2123d6b3cb8b2b834f50bcf90c82c044a92d4c942a56918f27542ac7960 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/onenote/sectiongroups/item/parentnotebook/sections/item/pages/item/parentnotebook"
    i6cabb51221ef24d9927ae122d57709c10a48f66322d726b3f63a19f5dac993cc "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/onenote/sectiongroups/item/parentnotebook/sections/item/pages/item/parentsection"
    id59ebaf6c207e0e7ea3060daf17f9b3c745f7f48508934af895febb06eaf307b "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/onenote/sectiongroups/item/parentnotebook/sections/item/pages/item/copytosection"
    ie33bf0f9df9e530ed4ba32add5696b5cf4cb45554d6f3933b3686d55abbb1bac "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/onenote/sectiongroups/item/parentnotebook/sections/item/pages/item/content"
)

// Builds and executes requests for operations under \sites\{site-id}\onenote\sectionGroups\{sectionGroup-id}\parentNotebook\sections\{onenoteSection-id}\pages\{onenotePage-id}
type OnenotePageRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Delete
type OnenotePageRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Options for Get
type OnenotePageRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *OnenotePageRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// The collection of pages in the section.  Read-only. Nullable.
type OnenotePageRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Options for Patch
type OnenotePageRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OnenotePage;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Instantiates a new OnenotePageRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewOnenotePageRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OnenotePageRequestBuilder) {
    m := &OnenotePageRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/sites/{site_id}/onenote/sectionGroups/{sectionGroup_id}/parentNotebook/sections/{onenoteSection_id}/pages/{onenotePage_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new OnenotePageRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewOnenotePageRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OnenotePageRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOnenotePageRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *OnenotePageRequestBuilder) Content()(*ie33bf0f9df9e530ed4ba32add5696b5cf4cb45554d6f3933b3686d55abbb1bac.ContentRequestBuilder) {
    return ie33bf0f9df9e530ed4ba32add5696b5cf4cb45554d6f3933b3686d55abbb1bac.NewContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OnenotePageRequestBuilder) CopyToSection()(*id59ebaf6c207e0e7ea3060daf17f9b3c745f7f48508934af895febb06eaf307b.CopyToSectionRequestBuilder) {
    return id59ebaf6c207e0e7ea3060daf17f9b3c745f7f48508934af895febb06eaf307b.NewCopyToSectionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// The collection of pages in the section.  Read-only. Nullable.
// Parameters:
//  - options : Options for the request
func (m *OnenotePageRequestBuilder) CreateDeleteRequestInformation(options *OnenotePageRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// The collection of pages in the section.  Read-only. Nullable.
// Parameters:
//  - options : Options for the request
func (m *OnenotePageRequestBuilder) CreateGetRequestInformation(options *OnenotePageRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        err := options.Q.AddQueryParameters(requestInfo.QueryParameters)
        if err != nil {
            return nil, err
        }
    }
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// The collection of pages in the section.  Read-only. Nullable.
// Parameters:
//  - options : Options for the request
func (m *OnenotePageRequestBuilder) CreatePatchRequestInformation(options *OnenotePageRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// The collection of pages in the section.  Read-only. Nullable.
// Parameters:
//  - options : Options for the request
func (m *OnenotePageRequestBuilder) Delete(options *OnenotePageRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
// The collection of pages in the section.  Read-only. Nullable.
// Parameters:
//  - options : Options for the request
func (m *OnenotePageRequestBuilder) Get(options *OnenotePageRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OnenotePage, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewOnenotePage() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OnenotePage), nil
}
func (m *OnenotePageRequestBuilder) OnenotePatchContent()(*i07a5432dfb3bd794fb35535d0db6f7c9e440efeea131949418329ba82c6e3435.OnenotePatchContentRequestBuilder) {
    return i07a5432dfb3bd794fb35535d0db6f7c9e440efeea131949418329ba82c6e3435.NewOnenotePatchContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OnenotePageRequestBuilder) ParentNotebook()(*i6767e2123d6b3cb8b2b834f50bcf90c82c044a92d4c942a56918f27542ac7960.ParentNotebookRequestBuilder) {
    return i6767e2123d6b3cb8b2b834f50bcf90c82c044a92d4c942a56918f27542ac7960.NewParentNotebookRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OnenotePageRequestBuilder) ParentSection()(*i6cabb51221ef24d9927ae122d57709c10a48f66322d726b3f63a19f5dac993cc.ParentSectionRequestBuilder) {
    return i6cabb51221ef24d9927ae122d57709c10a48f66322d726b3f63a19f5dac993cc.NewParentSectionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// The collection of pages in the section.  Read-only. Nullable.
// Parameters:
//  - options : Options for the request
func (m *OnenotePageRequestBuilder) Patch(options *OnenotePageRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
// Builds and executes requests for operations under \sites\{site-id}\onenote\sectionGroups\{sectionGroup-id}\parentNotebook\sections\{onenoteSection-id}\pages\{onenotePage-id}\microsoft.graph.preview()
func (m *OnenotePageRequestBuilder) Preview()(*i005a789b66b49c848e74abb76f962f75f90f0a0145cf80a3da3ca3a26ff5a4ba.PreviewRequestBuilder) {
    return i005a789b66b49c848e74abb76f962f75f90f0a0145cf80a3da3ca3a26ff5a4ba.NewPreviewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}

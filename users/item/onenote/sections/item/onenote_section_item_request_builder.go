package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i0fc20fd1f7b49561b2eb39e255a258171459e117ee215b32a7876ff924feae04 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/onenote/sections/item/copytosectiongroup"
    i18137a72a87f722fba1eb45b24b36cd64a881c6e2425aabe517212d00eba7447 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/onenote/sections/item/parentnotebook"
    i2d96b518ba6557ded78d4003d310e32c16ad886f3a56814e6efd7c4a6e1b3a44 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/onenote/sections/item/copytonotebook"
    i87cfbcf82b280a4dd6ee3e99ee116db635c37db2036e3eef89898210d9ed0be1 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/onenote/sections/item/pages"
    iee53ef0e675d4704860fcbe6f549da3fb67aa19ef28654d9dc919a22cef5929c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/onenote/sections/item/parentsectiongroup"
    i65099d9dbe634248efcc96947fbd2fe42b23aadf1172db5cc7f079df12bfb18a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/onenote/sections/item/pages/item"
)

// OnenoteSectionItemRequestBuilder builds and executes requests for operations under \users\{user-id}\onenote\sections\{onenoteSection-id}
type OnenoteSectionItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// OnenoteSectionItemRequestBuilderDeleteOptions options for Delete
type OnenoteSectionItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// OnenoteSectionItemRequestBuilderGetOptions options for Get
type OnenoteSectionItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *OnenoteSectionItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// OnenoteSectionItemRequestBuilderGetQueryParameters the sections in all OneNote notebooks that are owned by the user or group.  Read-only. Nullable.
type OnenoteSectionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// OnenoteSectionItemRequestBuilderPatchOptions options for Patch
type OnenoteSectionItemRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OnenoteSection;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewOnenoteSectionItemRequestBuilderInternal instantiates a new OnenoteSectionItemRequestBuilder and sets the default values.
func NewOnenoteSectionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OnenoteSectionItemRequestBuilder) {
    m := &OnenoteSectionItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/onenote/sections/{onenoteSection_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewOnenoteSectionItemRequestBuilder instantiates a new OnenoteSectionItemRequestBuilder and sets the default values.
func NewOnenoteSectionItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OnenoteSectionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOnenoteSectionItemRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *OnenoteSectionItemRequestBuilder) CopyToNotebook()(*i2d96b518ba6557ded78d4003d310e32c16ad886f3a56814e6efd7c4a6e1b3a44.CopyToNotebookRequestBuilder) {
    return i2d96b518ba6557ded78d4003d310e32c16ad886f3a56814e6efd7c4a6e1b3a44.NewCopyToNotebookRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OnenoteSectionItemRequestBuilder) CopyToSectionGroup()(*i0fc20fd1f7b49561b2eb39e255a258171459e117ee215b32a7876ff924feae04.CopyToSectionGroupRequestBuilder) {
    return i0fc20fd1f7b49561b2eb39e255a258171459e117ee215b32a7876ff924feae04.NewCopyToSectionGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation the sections in all OneNote notebooks that are owned by the user or group.  Read-only. Nullable.
func (m *OnenoteSectionItemRequestBuilder) CreateDeleteRequestInformation(options *OnenoteSectionItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the sections in all OneNote notebooks that are owned by the user or group.  Read-only. Nullable.
func (m *OnenoteSectionItemRequestBuilder) CreateGetRequestInformation(options *OnenoteSectionItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
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
// CreatePatchRequestInformation the sections in all OneNote notebooks that are owned by the user or group.  Read-only. Nullable.
func (m *OnenoteSectionItemRequestBuilder) CreatePatchRequestInformation(options *OnenoteSectionItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete the sections in all OneNote notebooks that are owned by the user or group.  Read-only. Nullable.
func (m *OnenoteSectionItemRequestBuilder) Delete(options *OnenoteSectionItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
// Get the sections in all OneNote notebooks that are owned by the user or group.  Read-only. Nullable.
func (m *OnenoteSectionItemRequestBuilder) Get(options *OnenoteSectionItemRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OnenoteSection, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewOnenoteSection() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OnenoteSection), nil
}
func (m *OnenoteSectionItemRequestBuilder) Pages()(*i87cfbcf82b280a4dd6ee3e99ee116db635c37db2036e3eef89898210d9ed0be1.PagesRequestBuilder) {
    return i87cfbcf82b280a4dd6ee3e99ee116db635c37db2036e3eef89898210d9ed0be1.NewPagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PagesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.onenote.sections.item.pages.item collection
func (m *OnenoteSectionItemRequestBuilder) PagesById(id string)(*i65099d9dbe634248efcc96947fbd2fe42b23aadf1172db5cc7f079df12bfb18a.OnenotePageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenotePage_id"] = id
    }
    return i65099d9dbe634248efcc96947fbd2fe42b23aadf1172db5cc7f079df12bfb18a.NewOnenotePageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *OnenoteSectionItemRequestBuilder) ParentNotebook()(*i18137a72a87f722fba1eb45b24b36cd64a881c6e2425aabe517212d00eba7447.ParentNotebookRequestBuilder) {
    return i18137a72a87f722fba1eb45b24b36cd64a881c6e2425aabe517212d00eba7447.NewParentNotebookRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OnenoteSectionItemRequestBuilder) ParentSectionGroup()(*iee53ef0e675d4704860fcbe6f549da3fb67aa19ef28654d9dc919a22cef5929c.ParentSectionGroupRequestBuilder) {
    return iee53ef0e675d4704860fcbe6f549da3fb67aa19ef28654d9dc919a22cef5929c.NewParentSectionGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch the sections in all OneNote notebooks that are owned by the user or group.  Read-only. Nullable.
func (m *OnenoteSectionItemRequestBuilder) Patch(options *OnenoteSectionItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
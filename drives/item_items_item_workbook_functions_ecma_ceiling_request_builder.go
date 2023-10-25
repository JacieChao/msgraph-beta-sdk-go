package drives

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemItemsItemWorkbookFunctionsEcma_CeilingRequestBuilder provides operations to call the ecma_Ceiling method.
type ItemItemsItemWorkbookFunctionsEcma_CeilingRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemsItemWorkbookFunctionsEcma_CeilingRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookFunctionsEcma_CeilingRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemsItemWorkbookFunctionsEcma_CeilingRequestBuilderInternal instantiates a new Ecma_CeilingRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookFunctionsEcma_CeilingRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookFunctionsEcma_CeilingRequestBuilder) {
    m := &ItemItemsItemWorkbookFunctionsEcma_CeilingRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/drives/{drive%2Did}/items/{driveItem%2Did}/workbook/functions/ecma_Ceiling", pathParameters),
    }
    return m
}
// NewItemItemsItemWorkbookFunctionsEcma_CeilingRequestBuilder instantiates a new Ecma_CeilingRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookFunctionsEcma_CeilingRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookFunctionsEcma_CeilingRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemsItemWorkbookFunctionsEcma_CeilingRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action ecma_Ceiling
func (m *ItemItemsItemWorkbookFunctionsEcma_CeilingRequestBuilder) Post(ctx context.Context, body ItemItemsItemWorkbookFunctionsEcma_CeilingPostRequestBodyable, requestConfiguration *ItemItemsItemWorkbookFunctionsEcma_CeilingRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookFunctionResultable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWorkbookFunctionResultFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookFunctionResultable), nil
}
// ToPostRequestInformation invoke action ecma_Ceiling
func (m *ItemItemsItemWorkbookFunctionsEcma_CeilingRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemItemsItemWorkbookFunctionsEcma_CeilingPostRequestBodyable, requestConfiguration *ItemItemsItemWorkbookFunctionsEcma_CeilingRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemItemsItemWorkbookFunctionsEcma_CeilingRequestBuilder) WithUrl(rawUrl string)(*ItemItemsItemWorkbookFunctionsEcma_CeilingRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsEcma_CeilingRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

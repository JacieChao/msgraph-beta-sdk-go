package drives

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemItemsItemAssignSensitivityLabelRequestBuilder provides operations to call the assignSensitivityLabel method.
type ItemItemsItemAssignSensitivityLabelRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemItemsItemAssignSensitivityLabelRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemAssignSensitivityLabelRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemsItemAssignSensitivityLabelRequestBuilderInternal instantiates a new AssignSensitivityLabelRequestBuilder and sets the default values.
func NewItemItemsItemAssignSensitivityLabelRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemAssignSensitivityLabelRequestBuilder) {
    m := &ItemItemsItemAssignSensitivityLabelRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/drives/{drive%2Did}/items/{driveItem%2Did}/microsoft.graph.assignSensitivityLabel";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewItemItemsItemAssignSensitivityLabelRequestBuilder instantiates a new AssignSensitivityLabelRequestBuilder and sets the default values.
func NewItemItemsItemAssignSensitivityLabelRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemAssignSensitivityLabelRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemsItemAssignSensitivityLabelRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action assignSensitivityLabel
func (m *ItemItemsItemAssignSensitivityLabelRequestBuilder) Post(ctx context.Context, body ItemItemsItemAssignSensitivityLabelPostRequestBodyable, requestConfiguration *ItemItemsItemAssignSensitivityLabelRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation invoke action assignSensitivityLabel
func (m *ItemItemsItemAssignSensitivityLabelRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemItemsItemAssignSensitivityLabelPostRequestBodyable, requestConfiguration *ItemItemsItemAssignSensitivityLabelRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
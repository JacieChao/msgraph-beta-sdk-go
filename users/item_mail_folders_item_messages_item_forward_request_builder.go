package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemMailFoldersItemMessagesItemForwardRequestBuilder provides operations to call the forward method.
type ItemMailFoldersItemMessagesItemForwardRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemMailFoldersItemMessagesItemForwardRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMailFoldersItemMessagesItemForwardRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemMailFoldersItemMessagesItemForwardRequestBuilderInternal instantiates a new ForwardRequestBuilder and sets the default values.
func NewItemMailFoldersItemMessagesItemForwardRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMailFoldersItemMessagesItemForwardRequestBuilder) {
    m := &ItemMailFoldersItemMessagesItemForwardRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/mailFolders/{mailFolder%2Did}/messages/{message%2Did}/microsoft.graph.forward";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewItemMailFoldersItemMessagesItemForwardRequestBuilder instantiates a new ForwardRequestBuilder and sets the default values.
func NewItemMailFoldersItemMessagesItemForwardRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMailFoldersItemMessagesItemForwardRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemMailFoldersItemMessagesItemForwardRequestBuilderInternal(urlParams, requestAdapter)
}
// Post forward a message using either JSON or MIME format. When using JSON format, you can:- Specify either a comment or the **body** property of the `message` parameter. Specifying both will return an HTTP 400 Bad Request error.- Specify either the `toRecipients` parameter or the **toRecipients** property of the `message` parameter. Specifying both or specifying neither will return an HTTP 400 Bad Request error. When using MIME format:- Provide the applicable Internet message headers and the MIME content, all encoded in **base64** format in the request body.- Add any attachments and S/MIME properties to the MIME content. This method saves the message in the **Sent Items** folder. Alternatively, create a draft to forward a message, and send it later.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/message-forward?view=graph-rest-1.0
func (m *ItemMailFoldersItemMessagesItemForwardRequestBuilder) Post(ctx context.Context, body ItemMailFoldersItemMessagesItemForwardPostRequestBodyable, requestConfiguration *ItemMailFoldersItemMessagesItemForwardRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation forward a message using either JSON or MIME format. When using JSON format, you can:- Specify either a comment or the **body** property of the `message` parameter. Specifying both will return an HTTP 400 Bad Request error.- Specify either the `toRecipients` parameter or the **toRecipients** property of the `message` parameter. Specifying both or specifying neither will return an HTTP 400 Bad Request error. When using MIME format:- Provide the applicable Internet message headers and the MIME content, all encoded in **base64** format in the request body.- Add any attachments and S/MIME properties to the MIME content. This method saves the message in the **Sent Items** folder. Alternatively, create a draft to forward a message, and send it later.
func (m *ItemMailFoldersItemMessagesItemForwardRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemMailFoldersItemMessagesItemForwardPostRequestBodyable, requestConfiguration *ItemMailFoldersItemMessagesItemForwardRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
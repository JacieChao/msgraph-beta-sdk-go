package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemEventsItemExceptionOccurrencesItemInstancesItemTentativelyAcceptRequestBuilder provides operations to call the tentativelyAccept method.
type ItemEventsItemExceptionOccurrencesItemInstancesItemTentativelyAcceptRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemEventsItemExceptionOccurrencesItemInstancesItemTentativelyAcceptRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemEventsItemExceptionOccurrencesItemInstancesItemTentativelyAcceptRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemEventsItemExceptionOccurrencesItemInstancesItemTentativelyAcceptRequestBuilderInternal instantiates a new TentativelyAcceptRequestBuilder and sets the default values.
func NewItemEventsItemExceptionOccurrencesItemInstancesItemTentativelyAcceptRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemEventsItemExceptionOccurrencesItemInstancesItemTentativelyAcceptRequestBuilder) {
    m := &ItemEventsItemExceptionOccurrencesItemInstancesItemTentativelyAcceptRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/events/{event%2Did}/exceptionOccurrences/{event%2Did1}/instances/{event%2Did2}/tentativelyAccept", pathParameters),
    }
    return m
}
// NewItemEventsItemExceptionOccurrencesItemInstancesItemTentativelyAcceptRequestBuilder instantiates a new TentativelyAcceptRequestBuilder and sets the default values.
func NewItemEventsItemExceptionOccurrencesItemInstancesItemTentativelyAcceptRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemEventsItemExceptionOccurrencesItemInstancesItemTentativelyAcceptRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemEventsItemExceptionOccurrencesItemInstancesItemTentativelyAcceptRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action tentativelyAccept
func (m *ItemEventsItemExceptionOccurrencesItemInstancesItemTentativelyAcceptRequestBuilder) Post(ctx context.Context, body ItemEventsItemExceptionOccurrencesItemInstancesItemTentativelyAcceptPostRequestBodyable, requestConfiguration *ItemEventsItemExceptionOccurrencesItemInstancesItemTentativelyAcceptRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToPostRequestInformation invoke action tentativelyAccept
func (m *ItemEventsItemExceptionOccurrencesItemInstancesItemTentativelyAcceptRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemEventsItemExceptionOccurrencesItemInstancesItemTentativelyAcceptPostRequestBodyable, requestConfiguration *ItemEventsItemExceptionOccurrencesItemInstancesItemTentativelyAcceptRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}

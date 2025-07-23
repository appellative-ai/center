package module

// Modal URI's

// Resource - representation and context, get and put
//
// Path : /resource/representation?n=core:resiliency:agent/traffic/http/limiter#1.2.3
//        /resource/context?n=core:resiliency:agent/traffic/http/limiter

// Notification - message, advice, put and get, trace only put
//
// Path : /notification/message
//        /notification/advice
//        /notification/trace

// Namespace - retrieval and request, post over the namespace. Retrieval needs to support a simple get with
//             query args
//
// Path : GET  /namespace/retrieval
//        POST /namespace/retrieval
//        POST /namespace/request
//        POST /namespace/connect

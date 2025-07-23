package module

// Modal URI's

// Resource - representation and context, get and put
//
// Path : /resource/representation?n=core:resiliency:agent/traffic/http/limiter#1.2.3
//        /resource/context?n=core:resiliency:agent/traffic/http/limiter

// Notification - message, advice, put and get, trace only put
//
// Path : /notification/message?n=core:resiliency:agent/traffic/http/limiter#1.2.3
//        /notification/advice/join?n=core:resiliency:agent/traffic/http/limiter
//        /notification/trace/join?n=core:resiliency:agent/traffic/http/limiter

// Namespace - thing and join, get and put
//
// Path : /namespace/thing?n=core:resiliency:agent/traffic/http/limiter#1.2.3
//        /namespace/join?n=core:resiliency:agent/traffic/http/limiter

// Namespace - retrieval and request, post over the namespace
//
// Path : /namespace/retrieval?n=core:resiliency:agent/traffic/http/limiter#1.2.3
//        /namespace/request?n=core:resiliency:agent/traffic/http/limiter

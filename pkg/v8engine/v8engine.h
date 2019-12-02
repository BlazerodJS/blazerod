#ifndef V8ENGINE_H
#define V8ENGINE_H
#ifdef __cplusplus
extern "C" {
#endif

typedef void* ContextPtr;
typedef void* IsolatePtr;
typedef void* ValuePtr;

typedef struct {
  const char* msg;
  const char* location;
  const char* stack;
} RtnError;

typedef struct {
  ValuePtr value;
  RtnError error;
} RtnValue;

// Initialize V8
extern void InitV8();

// Isolates
extern IsolatePtr NewIsolate();
extern void DisposeIsolate(IsolatePtr isolate);

// Contexts
extern ContextPtr NewContext(IsolatePtr isolate);
extern RtnValue Run(ContextPtr context, const char* source, const char* origin);
extern void DisposeContext(ContextPtr context);

// Values
const char* ValueToString(ValuePtr ptr);
extern void DisposeValue(ValuePtr value);

// V8 version
const char* Version();

#ifdef __cplusplus
}
#endif
#endif

function uintToString(data) {
  return String.fromCharCode.apply(null, new Uint8Array(data))
}

V8Engine.cb(msg => {
  V8Engine.log('Got a message from Go!')
  V8Engine.log(msg)
  V8Engine.log(uintToString(msg))
})

V8Engine.log('hello from JS')

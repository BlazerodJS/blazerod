declare module V8Engine {
  function log(msg: string): void
}

export async function log(...args: unknown[]) {
  V8Engine.log(args.join(' '))
}

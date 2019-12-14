import ts from '@wessberg/rollup-plugin-ts'
import resolve from 'rollup-plugin-node-resolve'

const extensions = ['.js', '.ts', '.mjs']

export default [
  // std library
  {
    input: 'std/log.ts',
    output: {
      file: 'std/dist/std-bundle.js',
      format: 'esm',
    },
    external: ['@test'],
    plugins: [resolve({extensions}), ts({tsconfig: 'std/tsconfig.json'})],
  },
]

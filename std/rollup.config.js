import resolve from 'rollup-plugin-node-resolve'
import babel from 'rollup-plugin-babel'

const extensions = ['.js', '.ts', '.mjs']

export default {
  input: 'log.ts',
  output: {
    file: 'std-bundle.js',
    format: 'esm',
  },
  plugins: [
    resolve({extensions}),
    babel({
      extensions,
      exclude: 'node_modules/**', // only transpile our source code
    }),
  ],
}

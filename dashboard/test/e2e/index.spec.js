import { resolve } from 'path'
import test from 'ava'
import { Nuxt, Builder } from 'nuxt'

// We keep the nuxt and server instance
// So we can close them at the end of the test
let nuxt = null

// Init Nuxt.js and create a server listening on localhost:4000
test.before(async () => {
  const config = {
    dev: false,
    rootDir: resolve(__dirname, '../../'),
  }
  nuxt = new Nuxt(config)
  await nuxt.ready()
  await new Builder(nuxt).build().catch(console.error)
  await nuxt.server.listen(4000, 'localhost')
}, 30000)

// Example of testing only generated html
test('Route / exits and render HTML', async (t) => {
  const { html } = await nuxt.renderRoute('/', {}).catch(console.error)
  t.true(html.includes('Mattrax Dashboard'))
})

// Close server and ask nuxt to stop listening to file changes
test.after('Closing server and nuxt.js', () => {
  nuxt.close()
})

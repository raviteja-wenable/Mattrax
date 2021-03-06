export function errorForStatus(
  context: any,
  res: any,
  catchErr: string
): Error {
  console.error('Error on request: ' + res.headers.get('x-request-id'))
  context.commit('dashboard/setErrorTraceID', res.headers.get('x-request-id'), {
    root: true,
  })

  if (res.status === 401) {
    const err = new Error('Unauthorised access to API')
    err.name = 'AuthError'
    return err
  } else if (res.status === 403) {
    const err = new Error('You do not have permission to access resource')
    return err
  } else if (res.status === 404) {
    return new Error('Resource not found')
  } else if (res.status === 422) {
    return new Error('Resource is not unique')
  } else if (res.status === 500) {
    return new Error('Internal server error')
  } else {
    return new Error(catchErr)
  }
}

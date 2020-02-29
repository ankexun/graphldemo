import {ApolloClient} from 'apollo-client'
import {HttpLink} from 'apollo-link-http'
import {InMemoryCache, IntrospectionFragmentMatcher} from 'apollo-cache-inmemory'
import {ApolloLink} from 'apollo-link'
import {onError} from 'apollo-link-error'

const httpLink = new HttpLink({
    uri: 'http://127.0.0.1:9090/graphql',    //请求路径
    credentials: 'include'        // 请求需要带入cookie则配置
  })

const middlewareLink = new ApolloLink((operation, forward) => {
    operation.setContext({
      headers: {
        'token': 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwicGFzc3dvcmQiOiIxMjM0IiwiZXhwIjoxNTgzNDk5OTE1LCJpc3MiOiJkZW1vIn0.5tXTOiLHTlRM1Uf7WHpTNyA1BaClaDz3QnfYJsHauF',
        // 'token': sessionStorage.getItem('token') || ${token} || null
      }
    })  //request拦截器
  
    return forward(operation).map(response => {
      return response
    })  //response拦截器
  })
  
  // 错误响应拦截器
  const errorLink = onError(({networkError, response}) => {
    let errorMsg = ''
    if (!!response && response.errors !== undefined && response.errors.length) {
      errorMsg = !response.errors[0].message ? '服务器错误' : response.errors[0].message
    }
    if (networkError) {
      errorMsg = networkError.message
      if (networkError.result !== undefined) {
        errorMsg = networkError.result.success === false ? networkError.result.message : networkError.result.error
      }
    }
    if (errorMsg) {
      console.log('apollo client error: ' + errorMsg)
    }
  })

  const authLink = middlewareLink.concat(httpLink)
  
  const defaultOptions = {
    watchQuery: {
      fetchPolicy: 'network-only',
      errorPolicy: 'ignore'
    },
    query: {
      fetchPolicy: 'network-only',
      errorPolicy: 'all'
    }
  }
  
  // 支持联合查询 
  const fragmentMatcher = new IntrospectionFragmentMatcher({
    introspectionQueryResultData: {
      __schema: {
        types: [
          {
            kind: 'INTERFACE',
            name: 'Document',
            possibleTypes: [
              {name: 'MyInterface1'},
              {name: 'SomeInterface2'}
            ]
          }
        ]
      }
    }
  })

  // 需要添加请求头
  export const apolloClient = new ApolloClient({
    link: errorLink.concat(authLink),
    cache: new InMemoryCache({fragmentMatcher}),
    connectToDevTools: true,
    defaultOptions: defaultOptions
  })
  
  // 不需要添加请求头
  export const baseClient = new ApolloClient({
    link: httpLink,
    cache: new InMemoryCache({fragmentMatcher}),
    connectToDevTools: true,
    defaultOptions: defaultOptions
  })

  // export default apolloClient
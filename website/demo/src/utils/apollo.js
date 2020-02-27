import ApolloClient from 'apollo-boost';
const apolloClient = new ApolloClient({
  // 这里指向后端的地址、端口和URL
  uri: 'http://127.0.0.1:9090/graphql'
})
export default apolloClient;
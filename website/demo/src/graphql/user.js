import gql from 'graphql-tag'
import {apolloClient,baseClient} from '../utils/apollo'

// 登录
export function login(params) {
 return baseClient.query({  //不需要带上token
  query: gql `query ($username : String!, $password : String!) {
   login(username: $username, password: $password) {
     token
   }
  }`,
  variables: params
 })
}

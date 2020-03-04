import gql from 'graphql-tag'
import {apolloClient,baseClient} from '../utils/apollo'

// todo list
export function queryTodolists(params) {
    return apolloClient.query({ //需要带上token
      query: gql `query ($user_id: Int, $page_num: Int, $page_size: Int) {
        queryTodolists(user_id: $user_id, page_num: $page_num, page_size: $page_size) {
            id
            content
            user_id
            status
            created_at
            updated_at
        }
      }`,
      variables: params
    })
  }

import gql from 'graphql-tag'
import {apolloClient,baseClient} from '../utils/apollo'

// 文章列表
export function getArticles(params) {
 return baseClient.query({  //不需要带上token
  query: gql `{
   articles{
    id
    title
    content
   }
  }`,
  variables: params
 })
}

// 单篇文章详情
export function getArticle(params) {
  return apolloClient.query({ //需要带上token
    query: gql `query ($id : Int) {
      article(id: $id) {
        id
        title
        content
      }
    }`,
    variables: params
  })
}

// 添加新文章
export function createArticle(params) {
 return apolloClient.mutate({
  mutation: gql `mutation ($title: String, $content: String) {
   add(title: $title, content: $content){
    id
    title
    content
   }
  }`,
  variables: params
 })
}

// 编辑文章
export function editArticle(params) {
  return apolloClient.mutate({
   mutation: gql `mutation ($id: Int, $title: String, $content: String) {
    update(id: $id, title: $title, content: $content){
     id
     title
     content
    }
   }`,
   variables: params
  })
 }

// 删除文章
export function deleteArticle(params) {
  return apolloClient.mutate({
   mutation: gql `mutation ($id: Int) {
    delete(id: $id){
     id
     title
     content
    }
   }`,
   variables: params
  })
 }
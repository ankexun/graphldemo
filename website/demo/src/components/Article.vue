<template>
  <div id="article">
    <div class="list">
      <h1>{{ msg }}</h1>
      <ul>
        <li v-for="(v, k) of list" :key="k">
          文章名称: {{ v.id }}----------------({{ v.title }})
          <button @click="getArticle(v.id)">详情</button>
          <button @click="deleteArticle(v.id)">删除</button>
        </li>
      </ul>
    </div>
    <div v-if="article.id > 0">
      <div>文章id:{{ article.id }}</div>
      标题：<input v-model="article.title" type="text"><br>
      文章内容: <textarea v-model="article.content" name="" id="" cols="30" rows="10"></textarea><br>
      <button @click="editArticle">编辑</button><button @click="article={}">取消</button>
    </div>
    <div class="form">
      <h1>添加文章</h1>
      标题：<input v-model="formData.title" type="text"><br>
      文章内容: <textarea v-model="formData.content" name="" id="" cols="30" rows="10"></textarea><br>
      <button @click="createArticle">添加</button>
    </div>
  </div>
  
</template>

<script>
import { getArticles,getArticle,createArticle,deleteArticle,editArticle } from '../graphql/article'

export default {
  name: 'Article',
  props: {
    msg: String
  },
  data() {
    return {
      list: [],
      formData: {
        title: '',
        content: ''
      },
      article: {
        id: 0,
        title: '',
        content: ''
      }
    }
  },
  methods: {
    initData() {
      getArticles()
      .then(res=>{
        console.log('request success')
        this.list = res.data.articles
      })
      .catch(err=>{
        console.log(err)
      })
    },
    getArticle(id) {
      getArticle({id:id})
      .then(res =>{
        this.article = res.data.article
      })
      .catch(err =>{
        console.log(err)
      })
    },
    createArticle() {
      createArticle(this.formData)
      .then(()=>{
        this.initData()
      })
      .catch(err=>{
        console.log(err)
      })
    },
    deleteArticle(id) {
      deleteArticle({id: id})
      .then(() =>{
        this.initData()
      })
      .catch(err=>{
        console.log(err)
      })
    },
    editArticle() {
      editArticle(this.article)
      .then(() =>{
        this.initData()
      })
      .catch(err=>{
        console.log(err)
      })
    }
  },
  mounted() {
    this.initData()
  }
}
</script>

<style>

</style>
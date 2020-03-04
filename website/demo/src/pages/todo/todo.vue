<template>
  <div id="todolist">
    <div class="content">
      <uni-list style="width: 98vw"
        v-model="loading"
        :finished="finished"
        finished-text="没有更多了"
        @load="onLoad"
      >
        <uni-list-item v-for="item in list" :key="item.id" :title="item.content" :note="item.updated_at" @click="popup(item.id)" />
      </uni-list>
    </div>
  </div>
</template>

<script>
  import { uniList,uniListItem } from '@dcloudio/uni-ui'
  import { queryTodolists } from '../../graphql/todo'

  export default {
    components: {
      uniList,
      uniListItem
    },
    data() {
      return {
        list: [],
        pageSize: 10,
        pageNum: 0, //起始页必须从0开始
        pageTotal: 0,
        loading: false,
        finished: false
      }
    },
    methods: {
      getTodolist(pageSize, pageNum) {
        queryTodolists({user_id: 1, page_num: pageNum, page_size: pageSize})
          .then(res => {
            this.list = res.data.queryTodolists
          })
          .catch(err => {
            console.log(err)
          })
      },
      popup(id) {
        console.log(id)
      },
      addTodolist() {
        //第一个' '里必须是空格，否则不显示input
        MessageBox.prompt(' ', '请输入新的待办事项')
          .then(({ value }) => {
            if (value.trim().length < 1) {
              MessageBox('提示', '亲,输入点正常的数据呗')
            } else {
              this.axios.post('/todolist/save.html', {
                  Content: value
                })
                .then(function (response) {
                  if (response.data.data.Code == 200) {
                    this.getTodolist(this.pageSize, 0)
                    this.bottomAllLoaded = false
                    this.pageNum = 0
                  } else {
                    MessageBox('添加失败', '添加的数据保存失败')
                  }
                }.bind(this))
                .catch(function (error) {
                  console.log("错误：" + error);
                })
            }
          });
      },
      updateTodolist(id, status) {
        this.axios.post('/todolist/update.html', {
            Id: id,
            Status: status
          })
          .then(function (response) {
            if (response.data.data.Code == 200) {
              this.getTodolist(this.pageSize, 0)
              this.bottomAllLoaded = false
              this.pageNum = 0
            } else {
              MessageBox('修改失败', '修改数据失败')
            }
          }.bind(this))
          .catch(function (error) {
            console.log("错误：" + error);
          })
      },
      onLoad() {
      // 异步更新数据
      // setTimeout 仅做示例，真实场景中一般为 ajax 请求
      setTimeout(() => {
        for (let i = 0; i < 10; i++) {
          // this.list.push(this.list.length + 1);
        }

        // 加载状态结束
        this.loading = false;

        // 数据全部加载完成
        if (this.list.length >= 40) {
          this.finished = true;
        }
      }, 1000);
    }
    },
    onReady() {
      this.getTodolist(this.pageSize, this.pageNum);
    }
  }
</script>

<style scoped>
     
</style>

<template>
 <div class="mt-3">
       <b-pagination
      v-model="curPage"
      :total-rows="rows"
      :per-page="pageSize"
      align="center"
      v-if="!totalRows == 0"
      @change="onPageChanged"
    ></b-pagination>
    </div>
   
</template>

<script lang="ts">
    import Vue from 'vue';

export default Vue.extend({
    name: "pagination",
    data() {
        return {
            curPage: this.currentPage
        }
    },
    props:["currentPage", "pageSize", "totalRows"],
    methods: {
        updatePage(pageNumber:number){
            this.$emit("page:update", pageNumber)
        }, 
        onPageChanged(page: number) {
            this.updatePage(page);
        }
    },
    computed: {
        rows() {
            return this.totalRows
        },
        
    }

})
</script>

<style>

.pagination-btn {
    cursor:pointer,
}

</style>
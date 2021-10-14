<template>
  <b-container class="mt-4">
    <b-row>
        <b-col sm align-h="center">
          <product-deck :productMatrix="productRows">
          </product-deck>
        </b-col>  
    </b-row>
    <!-- <b-row sm align-h="center">
      <b-form-group label="Large:" label-cols-sm="2" label-size="md">
        <b-form-file ref="upload" id="file-large" size="md" @change="onChangeFileUpload"></b-form-file>
        <b-button variant="success" @click="submitForm">Button</b-button>
      </b-form-group>
    </b-row> -->
    <b-row>
      <b-col sm align-h="center">
      <pagination :totalRows="products.length" align-h="center" :currentPage="currentPage" :pageSize="pageSize"
        v-on:page:update="updatePage"></pagination>
      </b-col>
    </b-row>
  </b-container>
</template>

<script lang="ts">
import Vue from 'vue';
import ProductDeck from '@/components/ProductDeck.vue';
import Pagination from '@/components/Pagination.vue'
import axios from 'axios'

export default Vue.extend({
  name: 'Home',
  data() {
    return {
      products: [],
      visibleProducts: [] as any,
      image: "",
      productRows: [] as any,
      currentPage: 1,
      pageSize: 4,
      rowSize: 4,
    }
  },
  components: {
    'product-deck': ProductDeck,
    'pagination': Pagination,
  },
  methods: {
    submitForm() {
      let formData = new FormData();
      formData.append('file', this.image);

      axios.post("/api/products/6154d2870fdacd4793d28f6c/image",
          formData,
        ).then(function (data) {
          console.log(data);
        })
        .catch(function (error) {
          if (error.response) {
            console.log(error.response);
          } else if (error.request) {
            console.log(error.request);
          } else {
            console.log(error.message);
          }
          console.log(error.config);
        });

    },

    onChangeFileUpload(event: any) {
      this.image = event.target.files[0]
    },

    updatePage(page: number) {
      this.currentPage = page
      this.updateVisibleProducts()
    },

    updateVisibleProducts() {
      this.visibleProducts = this.products.slice((this.currentPage - 1) * this.pageSize, this.pageSize * (this.currentPage))
      this.productRows = []

      this.fillEmptySpaces()

      while (this.visibleProducts.length) this.productRows.push(this.visibleProducts.splice(0, this.rowSize));

    },

    fillEmptySpaces() {
      if (this.visibleProducts.length % this.rowSize != 0) {
        let productsToAddCount = Math.abs(this.visibleProducts.length - this.pageSize)

        let prodLength = this.visibleProducts.length
        for (let i = 0; i < productsToAddCount; i++) {
          var invisibleProd = Object.assign({}, this.visibleProducts[0])
          invisibleProd.invisible = true
          this.visibleProducts[prodLength + i] = invisibleProd
        }
      }
    },

    async fetchProducts() {
      fetch("api/products")
        .then(res => res.json())
        .then(data => this.products = data)
        .then(data => this.updateVisibleProducts())
        .catch(err => console.log(err.message))
    }
  },

  mounted() {
    this.fetchProducts()
  }


});
</script>

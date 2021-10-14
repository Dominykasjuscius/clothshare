<template>
  <div id="app">
    <div id="nav">
    <my-nav></my-nav>
    </div>
    <router-view></router-view>
  </div>
</template>

<script>
import Navbar from '@/components/NavBar.vue'

export default {
  components:{
    "my-nav": Navbar,
  },
  methods: {
        updateVisibleProducts() {
      this.visibleProducts = this.products.slice((this.currentPage - 1) * this.pageSize, this.pageSize * (this.currentPage))
      this.productRows = []

console.log(this.pageSize - this.visibleProducts.length)
                if (this.visibleProducts.length % this.rowSize != 0) {
                    let productsToAddCount = Math.abs(this.visibleProducts.length -  this.pageSize)

                    let prodLength = this.visibleProducts.length
                    for (let i = 0; i < productsToAddCount; i++) {
                        var invisibleProd = Object.assign({}, this.visibleProducts[0])
                        invisibleProd.invisible = true
                        invisibleProd.index = i
                        this.visibleProducts[prodLength + i]= invisibleProd
                                                console.log(i ,this.visibleProducts.length)

                    }   
                }
     while (this.visibleProducts.length) this.productRows.push(this.visibleProducts.splice(0, this.rowSize));

      // if (this.visibleProducts.length == 0 && this.currentPage > 0) {
      //   this.updatePage(this.currentPage - 1)
      // }
    },
  }
}

</script>
<style lang="scss">

$primary: #D2E3C8;
$secondary: #2F4447;


@import "/Users/dominykasjuscius/Projektai/clothshare/webapp/node_modules/bootstrap/scss/bootstrap";
@import "/Users/dominykasjuscius/Projektai/clothshare/webapp/node_modules/bootstrap-vue/src/index.scss";
#app {
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: right;
  color: #2c3e50;
}

html *
{
   font-family: Arial !important;
   font-size: 100%
}
</style>
